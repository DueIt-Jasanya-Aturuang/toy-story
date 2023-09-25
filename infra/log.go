package infra

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.With().Caller().Logger()

	log.Info().Msgf("successfully init logger")
}
