package app_context

import (
	"fmt"
	"os"
	"time"
)

type Context struct {
	GitHubToken string
	Verbose     bool
	Timeout     time.Duration
	Retries     int
}

func Debug(ctx Context, format string, args ...any) {
	if !ctx.Verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "[DEBUG] "+format+"\n", args...)
}
