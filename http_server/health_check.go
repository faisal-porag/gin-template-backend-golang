package http_server

import (
	"github.com/faisal-porag/gin-template-backend-golang/state"
	"github.com/faisal-porag/gin-template-backend-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func healthCheck(s *state.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		language := utils.GetLanguageFromRequest(c.Request)

		logger := log.With().
			Str("handler", "healthCheck").
			Logger()

		isServiceDbOk, dErr := s.PgDbRepository.GetIsDatabaseUpAndRunning()
		if dErr != nil || !isServiceDbOk {
			if dErr != nil {
				logger.Error().Err(dErr).Msgf("Service Health Check: Debug Mode: %v", s.Cfg.DebugMode)
			}
			utils.ShowCommonSuccessResponse(
				c,
				"Service db is not running",
				nil,
				language,
			)
			return
		}

		// SUCCESS RESPONSE
		utils.ShowCommonSuccessResponse(
			c,
			"Service up & running",
			nil,
			language,
		)
		return
	}
}
