package source

import (
	"release-monitor/app_context"
	"release-monitor/model"
)

func Fetch(ctx app_context.Context, app model.AppConfig) (string, error) {
	switch app.Source.Type {
	case "github":
		return fetchGitHub(ctx, app.Source.GitHub)
	case "html":
		return fetchHTML(ctx, app.Source.HTML)
	default:
		return "", nil
	}
}
