package source

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"release-monitor/app_context"
	"release-monitor/model"

	"github.com/PuerkitoBio/goquery"
)

func fetchHTML(ctx app_context.Context, cfg *model.HTMLConfig) (string, error) {
	if cfg == nil {
		return "", fmt.Errorf("nil html config")
	}

	url := strings.TrimSpace(cfg.URL)
	if url == "" {
		return "", fmt.Errorf("empty html url")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "release-monitor")

	app_context.Debug(ctx, "fetching url: %s", cfg.URL)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("html request error: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	selection := doc.Find(cfg.Selector)
	if selection.Length() == 0 {
		return "", fmt.Errorf("selector not found: %s", cfg.Selector)
	}

	// Text of the first match
	text := strings.TrimSpace(selection.First().Text())

	app_context.Debug(ctx, "raw html text: %s", text)

	if text == "" {
		return "", fmt.Errorf("empty text for selector: %s", cfg.Selector)
	}

	return text, nil
}
