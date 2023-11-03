package controller

import (
	"sync"
	"web/model"
)

type Config struct {
	Wg     *sync.WaitGroup
	Models model.Models
}

var Cfg *Config

func InitCfg(c *Config) {
	Cfg = c
}
