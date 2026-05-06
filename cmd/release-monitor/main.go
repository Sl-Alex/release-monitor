package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"release-monitor/app"
	"release-monitor/app_context"
	"release-monitor/config"
)

func main() {
	configPath := flag.String("config", "config.json", "path to config file")
	flag.StringVar(configPath, "c", "config.json", "short for --config")
	githubToken := flag.String("github-token", "", "GitHub access token")
	verbose := flag.Bool("v", false, "verbose output")
	onlyUpdates := flag.Bool("only-updates", false, "show only apps with updates")
	flag.BoolVar(onlyUpdates, "u", false, "short for --only-updates")

	flag.Parse()

	token := *githubToken

	if token == "" {
		token = os.Getenv("GH_TOKEN")
	}

	ctx := app_context.Context{
		GitHubToken: token,
		Verbose:     *verbose,
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	hasUpdates := false
	hasErrors := false

	for _, a := range cfg.Apps {
		result := app.Process(ctx, a)

		if result.Changed {
			hasUpdates = true
		}

		if result.Err != "" {
			hasErrors = true
		}

		if *onlyUpdates && !result.Changed && result.Err == "" {
			continue
		}

		fmt.Println(app.Format(result))
	}

	if hasErrors {
		os.Exit(2)
	}

	if hasUpdates {
		os.Exit(1)
	}

	os.Exit(0)

}
