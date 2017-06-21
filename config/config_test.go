package config

import "testing"

func TestConfig(t *testing.T) {

	// Get config
	cfg := loadConfig("../passionlip.cfg.sample")

	// Test server config
	if exp := "[::]:8080"; cfg.Server.Listen != exp {
		t.Errorf("cfg.Server.Listen should be %s\n", exp)
	}

	// Test redis config
	if exp := "127.0.0.1:6379"; cfg.Redis.Addr != exp {
		t.Errorf("cfg.Redis.Addr should be %s\n", exp)
	}

	if exp := 0; cfg.Redis.DB != exp {
		t.Errorf("cfg.Redis.DB should be %d\n", exp)
	}

	if exp := 3; cfg.Redis.MaxRetries != exp {
		t.Errorf("cfg.Redis.MaxRetries should be %d\n", exp)
	}

	if exp := "channel01"; cfg.Redis.PubChannel != exp {
		t.Errorf("cfg.Redis.PubChannel should be %s\n", exp)
	}
}
