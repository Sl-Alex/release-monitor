package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"release-monitor/app_context"
	"release-monitor/model"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func fetchGitHub(ctx app_context.Context, cfg *model.GitHubConfig) (string, error) {
	if cfg == nil {
		return "", fmt.Errorf("nil github config")
	}

	repo := strings.TrimSpace(cfg.Repo)
	if repo == "" {
		return "", fmt.Errorf("empty github repository")
	}

	if len(strings.Split(repo, "/")) != 2 {
		return "", fmt.Errorf("invalid repository format (expected owner/repo): %s", repo)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// GitHub API requires User-Agent
	req.Header.Set("User-Agent", "release-monitor")
	if ctx.GitHubToken != "" {
		req.Header.Set("Authorization", "Bearer "+ctx.GitHubToken)
	}

    app_context.Debug(ctx, "github request: %s", url)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github api error: %s", resp.Status)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

    app_context.Debug(ctx, "github response tag: %s", release.TagName)

	if release.TagName == "" {
		return "", fmt.Errorf("empty tag_name in response")
	}

	return release.TagName, nil
}
