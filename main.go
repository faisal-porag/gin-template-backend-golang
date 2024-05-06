package main

import (
	"github.com/faisal-porag/gin-template-backend-golang/config"
	"github.com/faisal-porag/gin-template-backend-golang/http_server"
	"github.com/faisal-porag/gin-template-backend-golang/state"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal().Err(envLoadErr).Msg("env load error")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Config parsing failed")
	}

	appState := state.NewState(cfg)

	// Start the HTTP server
	go http_server.Server(appState)

	// Block the main goroutine to prevent the program from exiting
	select {}
}
