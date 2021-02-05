package pkg

import (
	"time"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var Config = struct {
	Log struct {
		Level   int  `default:"0" binding:"gte=-1,lte=7" env:"LOG_LEVEL"`
		NoColor bool `default:"false" env:"LOG_NOCOLOR"`
	}
	Host string `env:"HOST"`
}{}

func init() {
	err := configor.Load(&Config)
	if err != nil {
		log.Fatal().Err(err).Msgf("init config failed")
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	var out = zerolog.NewConsoleWriter()
	out.TimeFormat = time.RFC3339Nano
	out.NoColor = Config.Log.NoColor
	log.Logger = log.Output(out)
	// set log level
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.Level(Config.Log.Level))
}
