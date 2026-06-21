// Package httpclient provides a shared HTTP client with automatic retry
// and exponential backoff for resilient API calls.
package httpclient

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"
)

const (
	// maxRetries is the maximum number of retry attempts.
	maxRetries = 3

	// baseDelay is the initial delay before the first retry.
	baseDelay = 1 * time.Second

	// requestTimeout is the per-request timeout.
	requestTimeout = 30 * time.Second
)

// retryableStatusCodes defines HTTP status codes that warrant a retry.
var retryableStatusCodes = map[int]bool{
	http.StatusTooManyRequests:     true, // 429
	http.StatusInternalServerError: true, // 500
	http.StatusBadGateway:          true, // 502
	http.StatusServiceUnavailable:  true, // 503
	http.StatusGatewayTimeout:      true, // 504
}

// client is the shared HTTP client instance.
var client = &http.Client{
	Timeout: requestTimeout,
}

// Do performs an HTTP request with automatic retry and exponential backoff.
// It retries on network errors and retryable HTTP status codes (429, 5xx).
//
// Usage:
//
//	resp, err := httpclient.Do(ctx, "GET", url)
//	if err != nil { ... }
//	defer resp.Body.Close()
func Do(ctx context.Context, method, url string) (*http.Response, error) {
	var lastErr error

	for attempt := range maxRetries {
		// Wait before retry (skip wait on first attempt).
		if attempt > 0 {
			delay := backoffDelay(attempt)
			fmt.Printf("      ↻ Retry %d/%d in %s...\n", attempt, maxRetries-1, delay)

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}
		req.Header.Set("User-Agent", "free-games-tracker/0.2.0")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed (attempt %d): %w", attempt+1, err)
			continue
		}

		// Success — return immediately.
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		// Check if the status code is retryable.
		if retryableStatusCodes[resp.StatusCode] {
			resp.Body.Close()
			lastErr = fmt.Errorf("server returned %d (attempt %d)", resp.StatusCode, attempt+1)
			continue
		}

		// Non-retryable error — return immediately.
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", maxRetries, lastErr)
}

// backoffDelay calculates the exponential backoff delay for a given attempt.
// attempt 1 → 1s, attempt 2 → 2s, attempt 3 → 4s, etc.
func backoffDelay(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt-1))) * baseDelay
}
