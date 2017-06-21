package redis

import (
	"github.com/go-redis/redis"
	"github.com/tetsun/passionlip/config"
)

/*
Publisher structure
*/
type Publisher struct {
	Channel string
	Client  *redis.Client
}

/*
Pub publishes a message
*/
func (p *Publisher) Pub(msg string) error {
	return p.Client.Publish(p.Channel, msg).Err()
}

/*
MakePublisher creates a Publisher
*/
func MakePublisher(addr string, db int, maxretries int, channel string) *Publisher {

	var p Publisher

	// Set config
	p.Channel = channel

	// Set redis client
	p.Client = redis.NewClient(&redis.Options{
		Addr:       addr,
		DB:         db,
		MaxRetries: maxretries,
	})

	return &p
}

/*
NewPublisher creates a new Publisher
*/
func NewPublisher(cfg *config.Config) *Publisher {
	return MakePublisher(
		cfg.Redis.Addr,
		cfg.Redis.DB,
		cfg.Redis.MaxRetries,
		cfg.Redis.PubChannel,
	)
}
