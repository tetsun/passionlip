package redis

import "testing"

func TestPublisher(t *testing.T) {

	// Create Publisher
	p := MakePublisher("127.0.0.1:6379", 0, 3, "testchannel")

	// Test
	ps := p.Client.Subscribe("testchannel")
	defer ps.Close()

	if exp := ps.Ping(); exp != nil {
		t.Errorf("PubSub Ping error %s\n", exp)
	}

	if err := p.Pub("testmessage"); err != nil {
		t.Errorf("Publisher Pub error %s\n", err)
	}

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
