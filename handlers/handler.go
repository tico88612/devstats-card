package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/tico88612/devstats-card/models"
	"github.com/tico88612/devstats-card/service"
	"github.com/tico88612/devstats-card/svg"
	"github.com/tico88612/devstats-card/web"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, devStatsService *service.DevStatsService) {
	router.GET("/", RootHandler(devStatsService))
	router.GET("/health", HealthHandler)
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// RootHandler serves the SVG card when a username is provided (used by README
// embeds), and otherwise serves the preview frontend page.
func RootHandler(devStatsService *service.DevStatsService) gin.HandlerFunc {
	scoreHandler := ScoreHandler(devStatsService)
	return func(c *gin.Context) {
		if c.Query("username") == "" {
			c.Header("Cache-Control", "public, max-age=3600")
			c.Data(http.StatusOK, "text/html; charset=utf-8", web.IndexPage())
			return
		}
		scoreHandler(c)
	}
}

func ScoreHandler(devStatsService *service.DevStatsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		githubID := c.Query("username")
		if githubID == "" {
			c.String(http.StatusBadRequest, "Missing username")
			return
		}

		githubID = strings.ToLower(githubID)

		user, err := devStatsService.GetUserStats(githubID)
		if err != nil {
			c.String(http.StatusBadRequest, "Username not found.")
			return
		}

		card := svg.GenerateSVG(models.CardData{
			Username:      githubID,
			Score:         user.Contribution,
			PRs:           user.PRCount,
			Issues:        user.IssueCount,
			Rank:          user.Rank,
			Background:    "#0d1117",
			TitleColor:    "#0086FF",
			TextColor:     "#555555",
			Radius:        10,
			TitleFontSize: 24,
			TextFontSize:  18,
		})
		c.Header("Cache-Control", "public, max-age=7200")

		expiresTime := time.Now().Add(2 * time.Hour).Format(time.RFC1123)
		c.Header("Expires", expiresTime)

		c.Header("Content-Type", "image/svg+xml")
		c.String(http.StatusOK, card)
	}
}
