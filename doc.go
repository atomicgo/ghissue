/*
Package ghissue enables your users to submit issues to GitHub directly.

![Demo Video](https://raw.githubusercontent.com/atomicgo/ghissue/main/demo.gif)

Example:

	repo := ghissue.NewRepository("atomicgo", "ghissue")
	// [...]
	err := errors.New("This is an error")
	repo.CreateErrorReport(err) // Only creates an error report if the error is not nil
*/
package ghissue
