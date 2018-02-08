package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func Load() Type {
	file, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var cfg Type
	_, err = toml.Decode(string(dataBytes), &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
