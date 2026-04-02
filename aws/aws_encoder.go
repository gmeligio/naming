package aws

import (
	"fmt"
	"strings"
)

// AWSEncoder is responsible for encoding AWS names.
type AWSEncoder struct {
	// PrefixSegments are the default prefix segments used in AWS names.
	PrefixSegments []string
	// Separator is the default separator used in AWS names.
	Separator string
	// SSMSeparator is the default separator used in AWS SSM names.
	SSMSeparator string
	// SuffixLength is the length of the suffix to be added to AWS names.
	SuffixLength int
	// UseShortRegion indicates whether to use the short region name.
	UseShortRegion bool
}

// AWSEncoderOptions holds the options for generating AWS names.
type AWSEncoderOptions struct {
	// PrefixSegments are the prefix segments for the AWS name.
	PrefixSegments []string
	// Region is the region for the AWS name.
	Region *string
	// Separator is the separator for the AWS name.
	Separator *string
	// ShortName is the short name for the AWS name.
	ShortName *string
	// StartsWithSeparator indicates whether the name should start with the separator.
	StartsWithSeparator bool
}

// NewAWSEncoder creates a new AWSEncoder with default values.
func NewAWSEncoder() *AWSEncoder {
	return &AWSEncoder{
		PrefixSegments: []string{},
		SSMSeparator:   "/",
		Separator:      "-",
		SuffixLength:   4,
		UseShortRegion: true,
	}
}

// WithSeparator sets the default separator for the AWSEncoder.
func (e *AWSEncoder) WithSeparator(separator string) *AWSEncoder {
	e.Separator = separator
	return e
}

// WithSSMSeparator sets the default SSM separator for the AWSEncoder.
func (e *AWSEncoder) WithSSMSeparator(separator string) *AWSEncoder {
	e.SSMSeparator = separator
	return e
}

// WithPrefixSegments sets the default prefix segments for the AWSEncoder.
func (e *AWSEncoder) WithPrefixSegments(prefixSegments []string) *AWSEncoder {
	e.PrefixSegments = prefixSegments
	return e
}

// WithSuffixLength sets the default suffix length for the AWSEncoder.
func (e *AWSEncoder) WithSuffixLength(length int) *AWSEncoder {
	e.SuffixLength = length
	return e
}

// WithUseShortRegion sets whether to use the short region name for the AWSEncoder.
func (e *AWSEncoder) WithUseShortRegion(useShortRegion bool) *AWSEncoder {
	e.UseShortRegion = useShortRegion
	return e
}

// Encode generates a name based on the provided options.
func (e *AWSEncoder) Encode(options *AWSEncoderOptions) (string, error) {
	var segments []string

	// If region is provided, it will be the first segment.
	if options.Region != nil {
		// Initialize regionSegment with the full region name
		regionSegment := *options.Region

		if e.UseShortRegion {
			if foundRegion, ok := ShortRegions[*options.Region]; ok {
				regionSegment = foundRegion
			} else {
				return "", NewErrNonSupportedRegion(*options.Region, nil)
			}
		}

		segments = append(segments, regionSegment)
	}

	prefixSegments := e.PrefixSegments
	if options.PrefixSegments != nil {
		prefixSegments = options.PrefixSegments
	}

	segments = append(segments, prefixSegments...)
	segments = append(segments, *options.ShortName)

	separator := e.Separator
	if options.Separator != nil {
		separator = *options.Separator
	}

	joinedSegments := strings.Join(segments, separator)

	if options.StartsWithSeparator {
		return fmt.Sprintf("%s%s", separator, joinedSegments), nil
	}

	return joinedSegments, nil
}
