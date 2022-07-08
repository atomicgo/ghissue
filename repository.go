package ghissue

import (
	"fmt"
	"runtime"
	"strings"
)

// Repository is a GitHub repository.
type Repository struct {
	Owner string
	Name  string
}

// NewRepository creates a new Repository from an owner and repository name.
func NewRepository(owner, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// String returns the string representation of the repository.
func (repo Repository) String() string {
	return fmt.Sprintf("%s/%s", repo.Owner, repo.Name)
}

// NewIssue creates a new issue with a title and body.
func (repo Repository) NewIssue(title, body string) Issue {
	return NewIssue(repo, title, body)
}

// CreateErrorReport creates a new issue on GitHub with a detailed error report including the stack trace.
//
// Example:
// 		repo := ghissue.NewRepository("atomicgo", "ghissue")
//      // [...]
//      err := errors.New("This is an error")
// 		repo.CreateErrorReport(err)
func (repo Repository) CreateErrorReport(err error) error {
	if err == nil {
		return nil
	}

	title := fmt.Sprintf("[Error Report] `%s`", err.Error())

	var report []string

	report = append(report, "# Automatic Error Report")
	report = append(report, "")
	report = append(report, "## Error")
	report = append(report, "")
	report = append(report, fmt.Sprintf("```\n%s\n```", err))
	report = append(report, "")
	report = append(report, "## Operating System")
	report = append(report, "")
	report = append(report, "| Key | Value |")
	report = append(report, "| --- | --- |")
	report = append(report, fmt.Sprintf("| %s | %s |", "Operating System", runtime.GOOS))
	report = append(report, fmt.Sprintf("| %s | %s |", "Architecture", runtime.GOARCH))
	report = append(report, fmt.Sprintf("| %s | %s |", "Go Version", runtime.Version()))
	report = append(report, "")
	report = append(report, "## Stack Trace")
	report = append(report, "")
	report = append(report, "```")
	report = append(report, getStackTrace())
	report = append(report, "```")

	body := strings.Join(report, "\n")

	return repo.NewIssue(title, body).Open()
}
