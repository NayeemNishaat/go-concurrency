package lib

import (
	"os"
	"strconv"
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

func (cfg *Config) CreateMailer() *Mail {
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	port, err := strconv.ParseInt(os.Getenv("MAIL_PORT"), 10, 32)

	if err != nil {
		port = 25
	}

	m := Mail{
		Host:       os.Getenv("MAIL_HOST"),
		Port:       int(port),
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
		From:       os.Getenv("MAIL_FROM"),
		FromName:   os.Getenv("MAIL_FROM_NAME"),
		Wait:       cfg.Wg,
		ErrorChan:  errorChan,
		MailerChan: mailerChan,
		DoneChan:   mailerDoneChan,
	}

	return &m
}
