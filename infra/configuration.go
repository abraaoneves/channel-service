package infra

import (
	"log"
	"path"
	"runtime"

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
	// Get current filename execute caller
	_, filename, _, ok := runtime.Caller(1)

	if ok != true {
		log.Fatal(ok)
	}

	// Join file caller name to create configuration path
	filepath := path.Join(path.Dir(filename), "../configuration.toml")

	var confing Config
	_, err := toml.DecodeFile(filepath, &confing)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return confing
}
