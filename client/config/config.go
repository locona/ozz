package config

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Client struct {
	Redis Redis
}

type Redis struct {
	Host    string `required:"true"`
	Port    string `required:"true"`
	MaxConn int    `required:"true"`
}

type config struct {
	Redis
}

var Config *config

func Init() {
	client := Client{}
	if err := envconfig.Process("client", &client); err != nil {
		log.Fatal(err)
	}

	Config = &config{
		Redis: client.Redis,
	}
}
