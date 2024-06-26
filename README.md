<h1 align="center">AtomicGo | ghissue</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Fghissue&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/ghissue/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/ghissue?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/ghissue/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/ghissue?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-0-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/ghissue" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/ghissue?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/ghissue#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/ghissue</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ghissue

```go
import "atomicgo.dev/ghissue"
```

Package ghissue enables your users to submit issues to GitHub directly.

\!\[Demo Video\]\(https://raw.githubusercontent.com/atomicgo/ghissue/main/demo.gif\)

Example:

```
repo := ghissue.NewRepository("atomicgo", "ghissue")
// [...]
err := errors.New("This is an error")
repo.CreateErrorReport(err) // Only creates an error report if the error is not nil
```

## Index

- [Variables](<#variables>)
- [type Issue](<#Issue>)
  - [func NewIssue\(repo Repository, title, body string\) Issue](<#NewIssue>)
  - [func \(issue Issue\) GetCreateURL\(\) string](<#Issue.GetCreateURL>)
  - [func \(issue Issue\) Open\(\) error](<#Issue.Open>)
- [type Repository](<#Repository>)
  - [func NewRepository\(owner, name string\) Repository](<#NewRepository>)
  - [func \(repo Repository\) CreateErrorReport\(err error\) error](<#Repository.CreateErrorReport>)
  - [func \(repo Repository\) NewIssue\(title, body string\) Issue](<#Repository.NewIssue>)
  - [func \(repo Repository\) String\(\) string](<#Repository.String>)


## Variables

<a name="ErrOpenBrowser"></a>ErrOpenBrowser is the error returned when opening the browser fails.

```go
var ErrOpenBrowser = errors.New("failed to open browser")
```

<a name="Issue"></a>
## type [Issue](<https://github.com/atomicgo/ghissue/blob/main/issue.go#L11-L15>)

Issue is a GitHub issue.

```go
type Issue struct {
    Repository Repository
    Title      string
    Body       string
}
```

<a name="NewIssue"></a>
### func [NewIssue](<https://github.com/atomicgo/ghissue/blob/main/issue.go#L18>)

```go
func NewIssue(repo Repository, title, body string) Issue
```

NewIssue creates a new issue.

<a name="Issue.GetCreateURL"></a>
### func \(Issue\) [GetCreateURL](<https://github.com/atomicgo/ghissue/blob/main/issue.go#L27>)

```go
func (issue Issue) GetCreateURL() string
```

GetCreateURL returns the URL to create an issue.

<a name="Issue.Open"></a>
### func \(Issue\) [Open](<https://github.com/atomicgo/ghissue/blob/main/issue.go#L32>)

```go
func (issue Issue) Open() error
```

Open opens the "create issue" menu on GitHub in the browser.

<a name="Repository"></a>
## type [Repository](<https://github.com/atomicgo/ghissue/blob/main/repository.go#L10-L13>)

Repository is a GitHub repository.

```go
type Repository struct {
    Owner string
    Name  string
}
```

<a name="NewRepository"></a>
### func [NewRepository](<https://github.com/atomicgo/ghissue/blob/main/repository.go#L16>)

```go
func NewRepository(owner, name string) Repository
```

NewRepository creates a new Repository from an owner and repository name.

<a name="Repository.CreateErrorReport"></a>
### func \(Repository\) [CreateErrorReport](<https://github.com/atomicgo/ghissue/blob/main/repository.go#L41>)

```go
func (repo Repository) CreateErrorReport(err error) error
```

CreateErrorReport creates a new issue on GitHub with a detailed error report including the stack trace.

Example:

```
repo := ghissue.NewRepository("atomicgo", "ghissue")
     // [...]
     err := errors.New("This is an error")
		repo.CreateErrorReport(err)
```

<a name="Repository.NewIssue"></a>
### func \(Repository\) [NewIssue](<https://github.com/atomicgo/ghissue/blob/main/repository.go#L29>)

```go
func (repo Repository) NewIssue(title, body string) Issue
```

NewIssue creates a new issue with a title and body.

<a name="Repository.String"></a>
### func \(Repository\) [String](<https://github.com/atomicgo/ghissue/blob/main/repository.go#L24>)

```go
func (repo Repository) String() string
```

String returns the string representation of the repository.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
