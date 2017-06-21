package config

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/tetsun/passionlip/logger"
)

const defaultServerListen = "127.0.0.1:8080"
const defaultRedisAddr = "127.0.0.1:6379"
const defaultRedisDB = 0
const defaultRedisMaxRetries = 3
const defaultRedisPubChannel = "channel01"

/*
Config structure
*/
type Config struct {
	Server serverConfig
	Redis  redisConfig
}

type serverConfig struct {
	Listen string `toml:"listen"`
}

type redisConfig struct {
	Addr       string `toml:"addr"`
	DB         int    `toml:"db"`
	MaxRetries int    `toml:"maxretries"`
	PubChannel string `toml:"pubchannel"`
}

func defaultConfig() *Config {

	var cfg Config
	cfg.Server.Listen = defaultServerListen
	cfg.Redis.Addr = defaultRedisAddr
	cfg.Redis.DB = defaultRedisDB
	cfg.Redis.MaxRetries = defaultRedisMaxRetries
	cfg.Redis.PubChannel = defaultRedisPubChannel

	return &cfg
}

func loadConfig(path string) *Config {

	// Decode config
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Set default values
	if cfg.Server.Listen == "" {
		cfg.Server.Listen = defaultServerListen
	}

	if cfg.Redis.Addr == "" {
		cfg.Redis.Addr = defaultRedisAddr
	}

	if cfg.Redis.DB == 0 {
		cfg.Redis.DB = defaultRedisDB
	}

	if cfg.Redis.MaxRetries == 0 {
		cfg.Redis.MaxRetries = defaultRedisMaxRetries
	}

	if cfg.Redis.PubChannel == "" {
		cfg.Redis.PubChannel = defaultRedisPubChannel
	}

	return &cfg
}

/*
NewConfig creates new config
*/
func NewConfig() *Config {

	var cfg *Config

	// Config path
	cfgPath := flag.String("c", "", "config file path")
	flag.Parse()

	if *cfgPath == "" {
		cfg = defaultConfig()
	} else {
		cfg = loadConfig(*cfgPath)
	}

	return cfg
}
