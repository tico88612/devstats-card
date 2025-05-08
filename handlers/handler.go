package handlers

import (
	"net/http"

	"github.com/tico88612/devstats-card/models"
	"github.com/tico88612/devstats-card/service"
	"github.com/tico88612/devstats-card/svg"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, devStatsService *service.DevStatsService) {
	router.GET("/health", HealthHandler)
	router.GET("/score", ScoreHandler(devStatsService))
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func ScoreHandler(devStatsService *service.DevStatsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		githubID := c.Query("username")
		if githubID == "" {
			c.String(http.StatusBadRequest, "Missing username")
			return
		}

		user, err := devStatsService.GetUserStats(githubID)
		if err != nil {
			c.String(http.StatusBadRequest, "Username not found.")
			return
		}

		card := svg.GenerateSVG(models.CardData{
			Username:      githubID,
			Score:         user.Contribution,
			PRs:           user.PRCount,
			Rank:          user.Rank,
			Background:    "#0d1117",
			TitleColor:    "#0086FF",
			TextColor:     "#555555",
			Radius:        10,
			TitleFontSize: 24,
			TextFontSize:  18,
		})

		c.Header("Content-Type", "image/svg+xml")
		c.String(http.StatusOK, card)
	}
}
