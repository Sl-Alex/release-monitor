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

	for _, a := range cfg.Apps {
		result := app.Process(ctx, a)
		fmt.Println(app.Format(result))
	}
}
