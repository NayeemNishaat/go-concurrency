package lib

import (
	"log"
	"os"
	"strconv"
)

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

func (cfg *Config) ListenForMail() {
	for { // Note: Listening continuously from channels
		select {
		case msg := <-cfg.Mailer.MailerChan:
			go cfg.Mailer.SendMail(msg, cfg.Mailer.ErrorChan)
		case err := <-cfg.Mailer.ErrorChan:
			log.Println(err)
		case <-cfg.Mailer.DoneChan:
			return // Remark: Break out the loop and exit the go routine and stop listening
		}
	}
}

func (cfg *Config) PostMail(msg Message) {
	cfg.Wg.Add(1)
	cfg.Mailer.MailerChan <- msg // Important: If a chan is not created it will be nil and go will hang when it tries to send a msg to a nil chan and must need to listen from the chan (specially incase of unbuffered channel) else it will also cause go to hang. If the chan is unbuffered go will not return until the sent message is received/channel is empty.
}
