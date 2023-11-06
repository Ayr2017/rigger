package initvar

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
}

func (conf *Config) _init(){ 
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	conf.AppName = os.Getenv("APP_NAME")
}

func Init() (*Config, error) { 
	conf := &Config{}
    conf._init()
    return conf, nil
}
