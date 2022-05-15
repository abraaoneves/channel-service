package infra

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title string
	DB    Database `toml:"database"`
}

type Database struct {
	Address  string
	Password string
}

func configure() Config {
	// validate if file exist
	file := "./configuration.toml"
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
		panic(err)
	}

	var confing Config
	_, err := toml.DecodeFile(file, &confing)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return confing
}
