// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Connection struct {
		Sidekiq struct {
			Password string `config:"password"`
			Host     string `config:"host"`
			Port     string `config:"port"`
			Type     string `config:"type"`
		} `config:"sidekiq"`
	} `config:"connection"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
}
