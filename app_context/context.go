package app_context

import (
	"fmt"
	"os"
)

type Context struct {
	GitHubToken string
	Verbose     bool
}

func Debug(ctx Context, format string, args ...any) {
	if !ctx.Verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "[DEBUG] "+format+"\n", args...)
}
