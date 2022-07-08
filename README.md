<h1 align="center">AtomicGo | ghissue</h1>

<p align="center">

<a href="https://github.com/atomicgo/ghissue/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/ghissue?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue" target="_blank">
<img src="https://img.shields.io/github/workflow/status/atomicgo/ghissue/Go?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/ghissue?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/ghissue">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-0-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://github.com/atomicgo/ghissue/issues">
<img src="https://img.shields.io/github/issues/atomicgo/ghissue.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

---

<p align="center">
<strong><a href="#install">Get The Module</a></strong>
|
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
<td align="center">
<img width="2000" height="0"><br>
  -----------------------------------------------------------------------------------------------------
<img width="2000" height="0">
</td>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/ghissue</pre></h3>
<p align="center">
<table>
<tbody>
<td align="center">
<img width="2000" height="0"><br>
   -----------------------------------------------------------------------------------------------------
<img width="2000" height="0">
</td>
</tbody>
</table>
</p>

## Description

Package ghissue enables your users to submit issues to GitHub directly.

![Demo Video](https://raw.githubusercontent.com/atomicgo/ghissue/main/demo.gif)

Example:

    repo := ghissue.NewRepository("atomicgo", "ghissue")
    // [...]
    err := errors.New("This is an error")
    repo.CreateErrorReport(err) // Only creates an error report if the error is not nil


## Usage

```go
var ErrOpenBrowser = errors.New("failed to open browser")
```
ErrOpenBrowser is the error returned when opening the browser fails.

#### type Issue

```go
type Issue struct {
	Repository Repository
	Title      string
	Body       string
}
```

Issue is a GitHub issue.

#### func  NewIssue

```go
func NewIssue(repo Repository, title, body string) Issue
```
NewIssue creates a new issue.

#### func (Issue) GetCreateURL

```go
func (issue Issue) GetCreateURL() string
```
GetCreateURL returns the URL to create an issue.

#### func (Issue) Open

```go
func (issue Issue) Open() error
```
Open opens the "create issue" menu on GitHub in the browser.

#### type Repository

```go
type Repository struct {
	Owner string
	Name  string
}
```

Repository is a GitHub repository.

#### func  NewRepository

```go
func NewRepository(owner, name string) Repository
```
NewRepository creates a new Repository from an owner and repository name.

#### func (Repository) CreateErrorReport

```go
func (repo Repository) CreateErrorReport(err error) error
```
CreateErrorReport creates a new issue on GitHub with a detailed error report
including the stack trace.

Example:

    		repo := ghissue.NewRepository("atomicgo", "ghissue")
         // [...]
         err := errors.New("This is an error")
    		repo.CreateErrorReport(err)

#### func (Repository) NewIssue

```go
func (repo Repository) NewIssue(title, body string) Issue
```
NewIssue creates a new issue with a title and body.

#### func (Repository) String

```go
func (repo Repository) String() string
```
String returns the string representation of the repository.

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
