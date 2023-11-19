package lib

import "log"

func (cfg *Config) ListenForErrors() {
	for {
		select {
		case err := <-cfg.ErrorChan:
			log.Println(err)
		case <-cfg.ErrorChanDone:
			return
		}
	}
}
