package lib

import (
	"log"
)

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
	cfg.Mailer.MailerChan <- msg
}
