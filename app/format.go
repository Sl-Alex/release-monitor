package app

import (
	"release-monitor/model"
)

func Format(r model.Result) string {
	if r.Err != "" {
		return "> " + r.Name + ": error (" + r.Err + ")"
	}

	if r.Changed {
		return "> " + r.Name + ": update available (" + r.CurrentVersion + " → " + r.NewVersion + ")"
	}

	return "  " + r.Name + ": up to date (" + r.CurrentVersion + ")"
}
