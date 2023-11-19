package lib

import (
	"sync"
	"web/model"
)

type Config struct {
	Wg            *sync.WaitGroup
	Models        model.Models
	Mailer        *Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}

var Cfg *Config

func InitCfg(c *Config) {
	Cfg = c
}
