package source

import (
	"fmt"
	"net/http"
	"time"

	"release-monitor/app_context"
)

func doRequestWithRetry(
	ctx app_context.Context,
	client *http.Client,
	req *http.Request,
) (*http.Response, error) {

	var (
		resp    *http.Response
		err     error
		lastErr error
	)

	for i := 0; i <= ctx.Retries; i++ {
		resp, err = client.Do(req)

		if err != nil {
			lastErr = err
			app_context.Debug(ctx, "retry %d failed (network): %v", i+1, err)
		} else if shouldRetry(resp.StatusCode) {
			lastErr = fmt.Errorf("http error: %s", resp.Status)
			app_context.Debug(ctx, "retry %d failed (status): %s", i+1, resp.Status)

			resp.Body.Close()
		} else {
			// success or non-retryable error (e.g. 4xx)
			return resp, nil
		}

		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return nil, lastErr
}

func shouldRetry(status int) bool {
	return status >= 500 || status == 429
}
