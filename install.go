package granny

import (
	"errors"
)

var (
	// ErrInstallLocationEmpty occurs whenever the passed location of the
	// Install method is empty.
	ErrInstallLocationEmpty = errors.New("install: location must not be empty")
)

// InstallOptions holds optional configuration options for the Install
// method.
type InstallOptions struct {
	RunOptions

	// StartsWithWin indicates whether AnyDesk should be started with Windows.
	StartWithWin bool `json:"startWithWin,omitempty"`
}

// defaultInstallationOptions are the default InstallOptions which are used
// if none are given.
var defaultInstallationOptions = InstallOptions{
	RunOptions: defaultRunOptions,
}

// Install installs AnyDesk into the target location. The location
// must not be empty.
// See: https://support.anydesk.com/Command_Line_Interface
func Install(location string, opts ...InstallOptions) error {
	if location == "" {
		return ErrInstallLocationEmpty
	}
	p := pickInstallOptions(opts...)
	return Run(buildInstallArgs(location, p), p.RunOptions)
}

// pickInstallOptions picks the first index of the passed opts or the defaultInstallationOptions
// if none are given.
func pickInstallOptions(opts ...InstallOptions) InstallOptions {
	if len(opts) == 0 {
		return defaultInstallationOptions
	}
	return opts[0]
}

// buildInstallArgs builds the cli arguments for the Install method from the given
// location and opts.
func buildInstallArgs(location string, opts InstallOptions) (args []string) {
	args = append(args, "--install", location)
	if opts.StartWithWin {
		args = append(args, "--start-with-win")
	}
	return args
}
