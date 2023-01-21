package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port string
}

var config Config

func Init() {
	port := os.Getenv("PORT")

	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	c := Config{
		Port: port,
	}

	config = c
}

func Get() Config {
	return config
}
