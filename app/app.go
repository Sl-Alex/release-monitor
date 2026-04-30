package app

import (
	"release-monitor/app_context"
	"release-monitor/model"
	"release-monitor/source"
	"release-monitor/transform"
)

func Process(ctx app_context.Context, app model.AppConfig) model.Result {
	var result model.Result
	result.Name = app.Name
	result.CurrentVersion = app.Current

	// Fetch the source (html/github/...)
	raw, err := source.Fetch(ctx, app)
	if err != nil {
		result.Err = err.Error()
		result.NewVersion = app.Current
		return result
	}

	// Apply the transformation to the raw result (split/regex/...)
	normalized, err := transform.Apply(raw, app.Transform)
	if err != nil {
		result.Err = err.Error()
		result.NewVersion = app.Current
		return result
	}

	result.NewVersion = normalized
	result.Changed = normalized != app.Current

	return result
}
