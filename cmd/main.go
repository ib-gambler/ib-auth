package main

import (
	"github.com/ib-gambler/ib-auth/pkg"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	c := cron.New(
		cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)),
	)
	_, err := c.AddFunc("* * * * *", func() {
		err := pkg.ValidateSSO()
		if err != nil {
			log.Error().Err(err).Msgf("ValidateSSO() failed")
		}
	})
	if err != nil {
		log.Error().Err(err).Msgf("add func ValidateSSO() failed")
	}

	c.Run()
}
