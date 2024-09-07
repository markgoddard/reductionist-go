package config

import (
	"flag"
)

type Config struct {
	Port int
}

func Parse() Config {
	var conf Config
	conf.flags()
	flag.Parse()
	return conf
}

func (conf *Config) flags() {
	flag.IntVar(&conf.Port, "port", 8080, "Port to bind to")
}
