package controller

import "web/lib"

// type Cfg lib.Config

type Config struct {
	*lib.Config
}

var Cfg *Config

func InitCfg(c *lib.Config) {
	Cfg = &Config{c}
}
