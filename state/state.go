package state

import (
	"github.com/faisal-porag/gin-template-backend-golang/config"
	"github.com/faisal-porag/gin-template-backend-golang/db_repository"
	"github.com/rs/zerolog/log"
)

type State struct {
	Cfg            *config.Config
	PgDbRepository *db_repository.PgRepository
}

func NewState(cfg *config.Config) *State {
	pgDB, err := db_repository.NewPgRepository(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("pg connection error")
	} else {
		log.Info().Msg("pg connection successful")
	}

	return &State{
		Cfg:            cfg,
		PgDbRepository: pgDB,
	}
}
