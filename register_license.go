package granny

import (
	"errors"
	"strings"
)

var (
	// ErrLicenseKeyEmpty occurs whenever the passed key of the
	// RegisterLicense method is empty.
	ErrLicenseKeyEmpty = errors.New("register license: key must not be empty")
)

// RegisterLicense registers a license with the current AnyDesk
// installation. The passed key is the license key to use. It must not
// be empty.
func RegisterLicense(key string) error {
	if key == "" {
		return ErrLicenseKeyEmpty
	}
	return Run([]string{"--register--license"}, RunOptions{
		In: strings.NewReader(key),
	})
}
