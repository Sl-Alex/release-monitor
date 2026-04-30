package model

type AppConfig struct {
	Name    string
	Current string

	Source SourceConfig

	Transform []Transform
}

type SourceConfig struct {
	Type string // github | html

	GitHub *GitHubConfig
	HTML   *HTMLConfig
}

type GitHubConfig struct {
	Repo string
}

type HTMLConfig struct {
	URL      string
	Selector string
}

type Transform struct {
	Type   string
	Params []string
}

type Result struct {
	Name           string
	CurrentVersion string
	NewVersion     string
	Changed        bool
	Err            string
}
