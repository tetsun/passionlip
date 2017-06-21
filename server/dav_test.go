package server

import (
	"net/http"
	"strings"
	"testing"

	"github.com/tetsun/passionlip/redis"
)

func TestDav(t *testing.T) {

	// Create dav server
	p := redis.MakePublisher("127.0.0.1:6379", 0, 3, "testchannel")
	srv := MakeDav("127.0.0.1:8080", p)
	defer srv.Shutdown(nil)

	// Subscriber
	ps := p.Client.Subscribe("testchannel")
	defer ps.Close()

	// 405
	res, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		t.Errorf("http.Get error %s\n", err)
	}
	if exp := 405; res.StatusCode != exp {
		t.Errorf("res.StatusCode should be %d\n", exp)
	}

	// 201
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://127.0.0.1:8080", strings.NewReader("testmessage"))
	req.ContentLength = int64(len("testmessage"))
	res, err = client.Do(req)
	if err != nil {
		t.Errorf("client.Do error %s\n", err)
	}
	if exp := 201; res.StatusCode != exp {
		t.Errorf("res.StatusCode should be %d\n", exp)
	}

	// message check
	msg, err := ps.ReceiveMessage()

	if err != nil {
		t.Errorf("Pubsub ReceiveMessage error %s\n", err)
	}

	if exp := "testchannel"; msg.Channel != exp {
		t.Errorf("msg.Channel should be %s\n", exp)
	}

	if exp := "testmessage"; msg.Payload != exp {
		t.Errorf("msg.Payload should be %s\n", exp)
	}

}
