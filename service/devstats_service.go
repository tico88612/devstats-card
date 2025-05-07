package service

import (
	"github.com/tico88612/devstats-card/models"

	"github.com/tico88612/devstats-card/pkg/devstats"
)

type DevStatsService struct {
	client devstats.DevStatsInterface
}

func NewDevStatsService() *DevStatsService {
	return &DevStatsService{
		client: devstats.NewDevStats(""),
	}
}

func (s *DevStatsService) GetUserStats(username string) *models.User {
	user := &models.User{
		Username:     username,
		Contribution: -1,
		PRCount:      -1,
		Rank:         -1,
	}

	s.client.FetchContribute(user)
	s.client.FetchPRCount(user)

	return user
}
