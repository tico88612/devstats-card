package devstats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tico88612/devstats-card/models"
)

var URL = "https://devstats.cncf.io/api/v1"

type DevStats struct {
	DevStatsURL string
}

type DevStatsInterface interface {
	FetchContribute(user *models.User) error
}

func NewDevStats(serverURL string) DevStatsInterface {
	if serverURL == "" {
		serverURL = URL
	}
	return &DevStats{DevStatsURL: serverURL}
}

func (ds *DevStats) FetchContribute(user *models.User) error {
	data := DevStatsRequest{
		API: "GithubIDContributions",
		Payload: DevStatsPayload{
			GitHubID: user.Username,
		},
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("JSON Marshal error: %v", err)
		return err
	}

	resp, err := http.Post(ds.DevStatsURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read response error: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp struct {
			Error string `json:"error"`
		}
		if jsonErr := json.Unmarshal(bodyBytes, &errorResp); jsonErr == nil {
			log.Printf("API error: %s", errorResp.Error)
			return fmt.Errorf("API error: %s", errorResp.Error)
		} else {
			log.Printf("Server returned error: %d %s", resp.StatusCode, resp.Status)
			return fmt.Errorf("server error: %s", resp.Status)
		}
	}

	var result DevStatsContributionsResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return err
	}

	user.Contribution = result.Contributions
	user.IssueCount = result.Issues
	user.PRCount = result.PRs
	return nil
}
