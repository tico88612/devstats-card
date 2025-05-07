package devstats

type DevStatsRequest struct {
	API     string          `json:"api"`
	Payload DevStatsPayload `json:"payload"`
}

type DevStatsPayload struct {
	Project         string `json:"project"`
	Range           string `json:"range"`
	Metric          string `json:"metric"`
	RepositoryGroup string `json:"repository_group"`
	Country         string `json:"country"`
	GitHubID        string `json:"github_id"`
	BG              string `json:"bg"`
}
type DevStatsResponse struct {
	Project         string   `json:"project"`
	DBName          string   `json:"db_name"`
	Range           string   `json:"range"`
	Metric          string   `json:"metric"`
	RepositoryGroup string   `json:"repository_group"`
	Country         string   `json:"country"`
	GitHubID        string   `json:"github_id"`
	Filter          string   `json:"filter"`
	Rank            []int    `json:"rank"`
	Login           []string `json:"login"`
	Number          []int    `json:"number"`
}
