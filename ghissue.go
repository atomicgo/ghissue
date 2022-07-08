package ghissue

import (
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strings"
)

// ErrOpenBrowser is the error returned when opening the browser fails.
var ErrOpenBrowser = errors.New("failed to open browser")

const urlTemplate = "https://github.com/%s/%s/issues/new?title=%s&body=%s"

func getURL(owner, repo, title, body string) string {
	title = url.QueryEscape(title)
	body = url.QueryEscape(body)

	return fmt.Sprintf(urlTemplate, owner, repo, title, body)
}

func getStackTrace() string {
	b := make([]byte, 1024)
	n := runtime.Stack(b, false)

	trace := string(b[:n])

	// cut of first 5 lines of stack trace
	lines := strings.Split(trace, "\n")
	lines = lines[5:]
	trace = strings.Join(lines, "\n")
	trace = strings.TrimRight(trace, "\n")

	return trace
}
