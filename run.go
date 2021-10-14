package granny

import (
	"io"
	"os/exec"
)

// RunOptions holds optional configuration options for the Run method.
type RunOptions struct {
	// Runnable is the name or path to the AnyDesk runnable.
	// e.g.: 'anydesk.exe'
	Runnable string `json:"runnable,omitempty"`

	// In specifies the process' standard input.
	In io.Reader

	// Out specifies the process' standard output and error.
	Out io.Writer
}

// defaultRunOptions are the default RunOptions which are used if none are given.
var defaultRunOptions = RunOptions{
	Runnable: "anydesk",
}

// Run runs the runnable with the given args. The Runnable and more can be specified
// in the opts.
func Run(args []string, opts ...RunOptions) error {
	return createRunCommand(args, pickRunOptions(opts...)).Run()
}

// pickRunOptions picks the first index of the passed opts or the defaultRunOptions
// if none are given.
func pickRunOptions(opts ...RunOptions) RunOptions {
	if len(opts) == 0 {
		return defaultRunOptions
	}
	return opts[0]
}

// createRunCommand creates a new exec command from the given args and opts and
// returns it.
func createRunCommand(args []string, opts RunOptions) *exec.Cmd {
	cmd := exec.Command(opts.Runnable, args...)
	cmd.Stdin = opts.In
	cmd.Stdout = opts.Out
	return cmd
}
