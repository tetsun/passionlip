package server

import (
	"io/ioutil"
	"net/http"

	"github.com/tetsun/passionlip/config"
	"github.com/tetsun/passionlip/logger"
	"github.com/tetsun/passionlip/redis"
)

/*
MakeDav creates a dav server
*/
func MakeDav(listen string, p *redis.Publisher) *http.Server {

	srv := &http.Server{Addr: listen}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Allow only PUT request
		if r.Method != "PUT" {
			http.Error(w, "Method Not Allowed", 405)
			logger.DavLog(405, r.Method, r.RemoteAddr, "Method not allowed")
			return
		}

		// Read and publish the body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			logger.DavLog(500, r.Method, r.RemoteAddr, err.Error())
			return
		}

		if err := p.Pub(string(b)); err != nil {
			http.Error(w, "Internal Server Error", 500)
			logger.DavLog(500, r.Method, r.RemoteAddr, err.Error())
			return
		}

		logger.DavLog(201, r.Method, r.RemoteAddr, "PUT success")
		w.WriteHeader(201)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	return srv
}

/*
NewDav creates a new dav server
*/
func NewDav(cfg *config.Config) *http.Server {
	return MakeDav(cfg.Server.Listen, redis.NewPublisher(cfg))
}
