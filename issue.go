package ghissue

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/pterm/pterm"
)

// Issue is a GitHub issue.
type Issue struct {
	Repository Repository
	Title      string
	Body       string
}

// NewIssue creates a new issue.
func NewIssue(repo Repository, title, body string) Issue {
	return Issue{
		Repository: repo,
		Title:      title,
		Body:       body,
	}
}

// GetCreateURL returns the URL to create an issue.
func (issue Issue) GetCreateURL() string {
	return getURL(issue.Repository.Owner, issue.Repository.Name, issue.Title, issue.Body)
}

// Open opens the "create issue" menu in the browser.
func (issue Issue) Open() error {
	pterm.DefaultInteractiveConfirm.Show("Oh no! An error occurred. Would you like to create an issue on GitHub?")
	err := browser.OpenURL(issue.GetCreateURL())
	if err != nil {
		return fmt.Errorf("%w: %s", ErrOpenBrowser, err)
	}

	return nil
}
