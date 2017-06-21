package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tetsun/passionlip/config"
	"github.com/tetsun/passionlip/server"
)

func main() {

	// Create server
	cfg := config.NewConfig()
	srv := server.NewDav(cfg)

	// Handle signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	exitChan := make(chan int)

	go func() {
		for {
			s := <-sigChan

			switch s {
			// graceful restart
			case syscall.SIGHUP:
				log.Println("SIGHUP")
				cfg = config.NewConfig()
				srv.Shutdown(nil)
				srv = server.NewDav(cfg)

			// stop
			case syscall.SIGINT:
				log.Println("SIGINT")
				srv.Shutdown(nil)
				exitChan <- 0

			// kill
			case syscall.SIGTERM:
				log.Println("SIGTERM")
				exitChan <- 0

			default:
				log.Println("Unknown Signal")
				exitChan <- 1

			}
		}
	}()

	exitCode := <-exitChan
	os.Exit(exitCode)

}
