package granny

// SettingsPage represents a settings page of AnyDesk.
type SettingsPage string

const (
	UISettings           SettingsPage = "ui"
	SecuritySettings     SettingsPage = "security"
	AliasSettings        SettingsPage = "alias"
	PrivacySettings      SettingsPage = "privacy"
	VideoSettings        SettingsPage = "video"
	CaptureSettings      SettingsPage = "capture"
	AudioSettings        SettingsPage = "audio"
	ConnectionSettings   SettingsPage = "connection"
	FileTransferSettings SettingsPage = "filetransfer"
	RecordingSettings    SettingsPage = "recording"
	PrinterSettings      SettingsPage = "printer"
	WakeOnLANSettings    SettingsPage = "wol"
	LicenseSettings      SettingsPage = "license"
	AboutSettings        SettingsPage = "about"
)

// SettingsOptions holds optional configuration options for the
// OpenSettings method.
type SettingsOptions struct {
	RunOptions

	// Plain specifies whether to open a plain AnyDesk window.
	Plain bool `json:"plain,omitempty"`

	// Disclaimer specifies whether to show a custom disclaimer.
	Disclaimer bool `json:"disclaimer,omitempty"`

	// Advertisement specifies whether to show the advertisement page of
	// AnyDesk.
	Advertisement bool `json:"advertisement,omitempty"`
}

// defaultSettingsOptions are the default OpenSettings which are used
// if none are given.
var defaultSettingsOptions = SettingsOptions{
	RunOptions: defaultRunOptions,
}

// OpenSettings opens a settings page window in AnyDesk.
// See: https://support.anydesk.com/Command_Line_Interface
func OpenSettings(page SettingsPage, opts ...SettingsOptions) error {
	p := pickSettingsOptions(opts...)
	return Run(buildSettingsArgs(page, p), p.RunOptions)
}

// pickSettingsOptions picks the first index of the passed opts or the
// defaultSettingsOptions if none are given.
func pickSettingsOptions(opts ...SettingsOptions) SettingsOptions {
	if len(opts) == 0 {
		return defaultSettingsOptions
	}
	return opts[0]
}

// buildSettingsArgs builds the cli arguments for the OpenSettings method from the given
// page and opts.
func buildSettingsArgs(page SettingsPage, opts SettingsOptions) (args []string) {
	args = append(args, "--settings", string(page))
	if opts.Plain {
		args = append(args, "--plain")
	}
	if opts.Disclaimer {
		args = append(args, "--disclaimer")
	}
	if opts.Advertisement {
		args = append(args, "--show-advert")
	}
	return args
}
