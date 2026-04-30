package config

import (
	"errors"
	"strconv"
	"strings"

	"release-monitor/model"
)

func validate(cfg Config) error {
	if len(cfg.Apps) == 0 {
		return errors.New("no apps defined")
	}

	for i, app := range cfg.Apps {
		if app.Name == "" {
			return errors.New("app name is empty at index " + strconv.Itoa(i))
		}

		if app.Current == "" {
			return errors.New("app current version is empty: " + app.Name)
		}

		// source validation
		if err := validateSource(app); err != nil {
			return err
		}

		// transform validation
		if err := validateTransforms(app); err != nil {
			return err
		}
	}

	return nil
}

func validateSource(app model.AppConfig) error {
	if app.Source.Type == "" {
		return errors.New("source type is empty for app: " + app.Name)
	}

	switch app.Source.Type {
	case "github":
		if app.Source.GitHub == nil {
			return errors.New("github config missing for app: " + app.Name)
		}
		if app.Source.GitHub.Repo == "" {
			return errors.New("invalid github config for app: " + app.Name)
		}
		if !strings.Contains(app.Source.GitHub.Repo, "/") {
			return errors.New("github repository must be owner/repo")
		}
	case "html":
		if app.Source.HTML == nil {
			return errors.New("html config missing for app: " + app.Name)
		}
		if app.Source.HTML.URL == "" || app.Source.HTML.Selector == "" {
			return errors.New("invalid html config for app: " + app.Name)
		}

	default:
		return errors.New("unknown source type: " + app.Source.Type)
	}

	return nil
}

func validateTransforms(app model.AppConfig) error {
	for _, t := range app.Transform {
		if t.Type == "" {
			return errors.New("empty transform type in app: " + app.Name)
		}

		switch t.Type {
		case "regex":
			if len(t.Params) != 1 {
				return errors.New("regex transform must have 1 params in app: " + app.Name)
			}

		case "split":
			if len(t.Params) != 2 {
				return errors.New("split transform must have 2 params in app: " + app.Name)
			}

		default:
			return errors.New("unknown transform type: " + t.Type + " in app: " + app.Name)
		}
	}

	return nil
}
