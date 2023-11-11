package controller

import (
	"os"
	"strconv"
	"sync"
	"web/lib"
	"web/model"
)

type Config struct {
	Wg     *sync.WaitGroup
	Models model.Models
	Mailer *lib.Mail
}

var Cfg *Config

func InitCfg(c *Config) {
	Cfg = c
}

func (cfg *Config) CreateMailer() *lib.Mail {
	errorChan := make(chan error)
	mailerChan := make(chan lib.Message, 100)
	mailerDoneChan := make(chan bool)

	port, err := strconv.ParseInt(os.Getenv("MAIL_PORT"), 10, 32)

	if err != nil {
		port = 25
	}

	m := lib.Mail{
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
