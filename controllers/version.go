package controllers

import (
	"fmt"
	"strings"

	goversion "github.com/hashicorp/go-version"
)

// Version is a wrapper around github.com/hashicorp/go-version that adds additional
// functions for developer's usability
type Version struct {
	version *goversion.Version
}

// NewVersion creates a new version from given string
func NewVersion(v string) (*Version, error) {
	version, err := goversion.NewVersion(v)
	if err != nil {
		return nil, err
	}
	return &Version{version: version}, nil
}

// String returns version string
func (v *Version) String() string {
	return v.version.String()
}

// ToCRVersion returns version usable as CRversion parameter
func (v *Version) ToCRVersion() string {
	return strings.ReplaceAll(v.String(), "v", "")
}

// ToSemver returns version is semver format
func (v *Version) ToSemver() string {
	return fmt.Sprintf("v%s", v.String())
}

// ToAPIVersion returns version that can be used as K8s APIVersion parameter
func (v *Version) ToAPIVersion(apiRoot string) string {
	ver, _ := goversion.NewVersion("v1.12.0")
	if v.version.GreaterThan(ver) {
		return fmt.Sprintf("%s/v1", apiRoot)
	}
	return fmt.Sprintf("%s/v%s", apiRoot, strings.ReplaceAll(v.String(), ".", "-"))
}
