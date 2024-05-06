package http_server

import (
	"fmt"
	"github.com/faisal-porag/gin-template-backend-golang/route_manager"
	"github.com/faisal-porag/gin-template-backend-golang/state"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Server(s *state.State) {
	if s.Cfg.AppReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a Gin router instance
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{
		"Accept",
		"Authorization",
		"Content-Type",
		"X-CSRF-Token",
		"Accept-Language",
	}
	router.Use(cors.New(config))

	// Define routes
	v1 := router.Group("/api/v1")
	{
		// SERVICE HEALTH CHECK
		v1.GET("/db/health-check", healthCheck(s))

		// Example
		route_manager.MountExampleRoutes(v1.Group("/examples"), s)

	}

	// Start the Gin server
	port := s.Cfg.ApplicationPort
	log.Info().Int("port", port).Msg("Starting backend service http server")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("router.Run err")
	}
}
