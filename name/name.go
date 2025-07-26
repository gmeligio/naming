// Package name provides interfaces and utilities for generating standardized
// resource names with customizable prefixes, segments, and separators.
package name

type Namer interface {
	Name(prefixSegments []string, shortName string, separator string, startsWithSeparator bool) (string, error)
}

type Encoder interface {
	Encode(name string) (string, error)
}

// StringPtr is a helper function to get a pointer to a string.
func StringPtr(s string) *string { return &s }
