package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tico88612/devstats-card/models"
	"github.com/tico88612/devstats-card/pkg/devstats"
	"github.com/tico88612/devstats-card/svg"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET("/score", func(c *gin.Context) {
		githubID := c.Query("username")
		if githubID == "" {
			c.String(http.StatusBadRequest, "Missing username")
			return
		}

		a := devstats.NewDevStats("")
		u := models.User{
			Username:     githubID,
			Contribution: -1,
			PRCount:      -1,
			Rank:         -1,
		}

		a.FetchContribute(&u)
		a.FetchPRCount(&u)

		card := svg.GenerateSVG(models.CardData{
			Username:      githubID,
			Score:         u.Contribution,
			PRs:           u.PRCount,
			Rank:          u.Rank,
			Background:    "#0d1117",
			TitleColor:    "#0086FF",
			TextColor:     "#555555",
			Radius:        10,
			TitleFontSize: 24,
			TextFontSize:  18,
		})

		c.Header("Content-Type", "image/svg+xml")
		c.String(http.StatusOK, card)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
