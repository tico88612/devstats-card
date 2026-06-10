package handlers

import (
	"net/http"
	"regexp"
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
	router.HEAD("/", HeadHandler)
	router.GET("/health", HealthHandler)
	router.HEAD("/health", HeadHandler)
}

// HeadHandler responds to HEAD requests with 200 and no body, so external
// uptime monitors can health-check the service cheaply.
func HeadHandler(c *gin.Context) {
	c.Status(http.StatusOK)
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

var hexColorPattern = regexp.MustCompile(`^[0-9a-fA-F]{3,8}$`)

// resolveTheme builds the card palette from the `theme` query parameter, then
// applies any per-color overrides (hex without the leading `#`, e.g.
// `&title_color=ff79c6`). Invalid values are ignored to keep the SVG safe.
func resolveTheme(c *gin.Context) models.Theme {
	theme := models.GetTheme(c.Query("theme"))

	overrides := []struct {
		param string
		dest  *string
	}{
		{"bg_color", &theme.Background},
		{"border_color", &theme.Border},
		{"title_color", &theme.TitleColor},
		{"text_color", &theme.TextColor},
		{"icon_color", &theme.IconColor},
	}
	for _, o := range overrides {
		if v := c.Query(o.param); hexColorPattern.MatchString(v) {
			*o.dest = "#" + v
		}
	}
	return theme
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

		theme := resolveTheme(c)
		card := svg.GenerateSVG(models.CardData{
			Username:      githubID,
			Score:         user.Contribution,
			PRs:           user.PRCount,
			Issues:        user.IssueCount,
			Rank:          user.Rank,
			Background:    theme.Background,
			Border:        theme.Border,
			TitleColor:    theme.TitleColor,
			TextColor:     theme.TextColor,
			IconColor:     theme.IconColor,
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
