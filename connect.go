package granny

import "errors"

var (
	// ErrConnectDestinationEmpty occurs whenever the passed location of the
	// Connect method is empty.
	ErrConnectDestinationEmpty = errors.New("connect: destination must not be empty")
)

// ConnectOptions holds optional configuration options for the Connect method.
type ConnectOptions struct {
	RunOptions

	// FileTransfer specifies whether a file transfer session should be started.
	FileTransfer bool `json:"fileTransfer,omitempty"`

	// FullScreen specifies whether to start the session in fullscreen mode.
	FullScreen bool `json:"fullScreen,omitempty"`

	// Plain specifies whether to start a session without window title and toolbar.
	Plain bool `json:"plain,omitempty"`
}

// defaultConnectOptions are the default ConnectOptions which are used if none
// are given.
var defaultConnectOptions = ConnectOptions{
	RunOptions: defaultRunOptions,
}

// Connect opens a session to the destination. The destination should be an AnyDesk
// address, such as 'foo@ad' or '871813768'.
// See: https://support.anydesk.com/Command_Line_Interface
func Connect(destination string, opts ...ConnectOptions) error {
	if destination == "" {
		return ErrConnectDestinationEmpty
	}
	p := pickConnectOptions(opts...)
	return Run(buildConnectArgs(destination, p), p.RunOptions)
}

// pickConnectOptions picks the first index of the passed opts or the defaultConnectOptions
// if none are given.
func pickConnectOptions(opts ...ConnectOptions) ConnectOptions {
	if len(opts) == 0 {
		return defaultConnectOptions
	}
	return opts[0]
}

// buildConnectArgs builds the cli arguments for the Connect method from the given
// location and opts.
func buildConnectArgs(destination string, opts ConnectOptions) (args []string) {
	args = append(args, destination)
	if opts.FileTransfer {
		args = append(args, "--file-transfer")
	}
	if opts.FullScreen {
		args = append(args, "--fullscreen")
	}
	if opts.Plain {
		args = append(args, "--plain")
	}
	return args
}
