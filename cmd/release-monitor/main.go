package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"release-monitor/app"
	"release-monitor/app_context"
	"release-monitor/config"
	"release-monitor/model"
)

func main() {
	configPath := flag.String("config", "config.json", "path to config file")
	flag.StringVar(configPath, "c", "config.json", "short for --config")
	githubToken := flag.String("github-token", "", "GitHub access token")
	verbose := flag.Bool("v", false, "verbose output")
	onlyUpdates := flag.Bool("only-updates", false, "show only apps with updates")
	flag.BoolVar(onlyUpdates, "u", false, "short for --only-updates")
	timeout := flag.Int("timeout", 10, "http timeout in seconds")
	retries := flag.Int("retries", 2, "number of retries")
	format := flag.String("format", "text", "output format: text|json")

	flag.Parse()

	token := *githubToken

	if token == "" {
		token = os.Getenv("GH_TOKEN")
	}

	ctx := app_context.Context{
		GitHubToken: token,
		Verbose:     *verbose,
		Timeout:     time.Duration(*timeout) * time.Second,
		Retries:     *retries,
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	hasUpdates := false
	hasErrors := false

	var results []model.Result

	for _, a := range cfg.Apps {
		result := app.Process(ctx, a)

		results = append(results, result)

		if result.Changed {
			hasUpdates = true
		}

		if result.Err != "" {
			hasErrors = true
		}

		if *format == "text" {
			if *onlyUpdates && !result.Changed && result.Err == "" {
				continue
			}

			fmt.Println(app.Format(result))
		}
	}

	if *format == "json" {
		output := results

		if *onlyUpdates {
			var filtered []model.Result
			for _, r := range results {
				if r.Changed || r.Err != "" {
					filtered = append(filtered, r)
				}
			}
			output = filtered
		}

		data, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	}

	if hasErrors {
		os.Exit(2)
	}

	if hasUpdates {
		os.Exit(1)
	}

	os.Exit(0)

}
