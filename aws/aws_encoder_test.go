// Unit tests for the AWS package

package aws_test

import (
	"testing"

	"github.com/gmeligio/naming/aws"
	"github.com/gmeligio/naming/name"
	"github.com/stretchr/testify/require"
)

func TestAWSEncoder_Encode_DefaultOptions(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                 "basic service",
			shortName:            "service",
			expected:             "service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: nil,
			expectedError:        nil,
		},
		{
			name:                 "with region uses short region by default",
			shortName:            "service",
			expected:             "usw2-service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("us-west-2"),
			startsWithSeparator:  false,
			encoderModifications: nil,
			expectedError:        nil,
		},
		{
			name:                 "with prefixes",
			shortName:            "service",
			expected:             "prod-app-service",
			prefixSegments:       []string{"prod", "app"},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: nil,
			expectedError:        nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestAWSEncoder_Encode_UseShortRegion(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                 "UseShortRegion false",
			shortName:            "service",
			expected:             "us-west-2-service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("us-west-2"),
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithUseShortRegion(false) },
			expectedError:        nil,
		},
		{
			name:                 "UseShortRegion true",
			shortName:            "service",
			expected:             "usw2-service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("us-west-2"),
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithUseShortRegion(true) },
			expectedError:        nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestAWSEncoder_Encode_Separator(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                 "Separator underscore",
			shortName:            "service",
			expected:             "prod_app_service",
			prefixSegments:       []string{"prod", "app"},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithSeparator("_") },
			expectedError:        nil,
		},
		{
			name:                 "Separator with region",
			shortName:            "service",
			expected:             "usw2_service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("us-west-2"),
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithSeparator("_") },
			expectedError:        nil,
		},
		{
			name:                 "Separator overridden by options",
			shortName:            "service",
			expected:             "prod/service",
			prefixSegments:       []string{"prod"},
			separator:            name.StringPtr("/"),
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithSeparator("_") },
			expectedError:        nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestAWSEncoder_Encode_PrefixSegments(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                "PrefixSegments",
			shortName:           "service",
			expected:            "global-shared-service",
			prefixSegments:      nil,
			separator:           nil,
			region:              nil,
			startsWithSeparator: false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder {
				return e.WithPrefixSegments([]string{"global", "shared"})
			},
			expectedError: nil,
		},
		{
			name:                "PrefixSegments overridden by options",
			shortName:           "service",
			expected:            "prod-app-service",
			prefixSegments:      []string{"prod", "app"},
			separator:           nil,
			region:              nil,
			startsWithSeparator: false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder {
				return e.WithPrefixSegments([]string{"global", "shared"})
			},
			expectedError: nil,
		},
		{
			name:                "PrefixSegments with region",
			shortName:           "service",
			expected:            "usw2-global-shared-service",
			prefixSegments:      nil,
			separator:           nil,
			region:              name.StringPtr("us-west-2"),
			startsWithSeparator: false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder {
				return e.WithPrefixSegments([]string{"global", "shared"})
			},
			expectedError: nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestAWSEncoder_Encode_Advanced(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                "Multiple options",
			shortName:           "service",
			expected:            "us-west-2_global_shared_service",
			prefixSegments:      nil,
			separator:           nil,
			region:              name.StringPtr("us-west-2"),
			startsWithSeparator: false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder {
				return e.WithUseShortRegion(false).WithSeparator("_").WithPrefixSegments([]string{"global", "shared"})
			},
			expectedError: nil,
		},
		{
			name:                 "StartsWithSeparator true",
			shortName:            "service",
			expected:             "_prod_service",
			prefixSegments:       []string{"prod"},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  true,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithSeparator("_") },
			expectedError:        nil,
		},
		{
			name:                 "Special characters",
			shortName:            "service*",
			expected:             "prod$-app#-service*",
			prefixSegments:       []string{"prod$", "app#"},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: nil,
			expectedError:        nil,
		},
		{
			name:                 "Empty separator",
			shortName:            "service",
			expected:             "prodappservice",
			prefixSegments:       []string{"prod", "app"},
			separator:            nil,
			region:               nil,
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithSeparator("") },
			expectedError:        nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestAWSEncoder_Encode_Errors(t *testing.T) {
	testCases := []struct {
		name                 string
		shortName            string
		expected             string
		prefixSegments       []string
		separator            *string
		region               *string
		startsWithSeparator  bool
		encoderModifications func(*aws.AWSEncoder) *aws.AWSEncoder
		expectedError        error
	}{
		{
			name:                 "Non-existent region",
			shortName:            "service",
			expected:             "",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("non-existent"),
			startsWithSeparator:  false,
			encoderModifications: nil, // Uses default UseShortRegion=true, which will cause error for invalid region
			expectedError:        aws.NewErrNonSupportedRegion("non-existent", nil),
		},
		{
			name:                 "Non-existent region with UseShortRegion false",
			shortName:            "service",
			expected:             "non-existent-service",
			prefixSegments:       []string{},
			separator:            nil,
			region:               name.StringPtr("non-existent"),
			startsWithSeparator:  false,
			encoderModifications: func(e *aws.AWSEncoder) *aws.AWSEncoder { return e.WithUseShortRegion(false) },
			expectedError:        nil, // Should not error when UseShortRegion=false
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := aws.NewAWSEncoder()

			// Apply encoder modifications if specified
			if tc.encoderModifications != nil {
				encoder = tc.encoderModifications(encoder)
			}

			options := &aws.AWSEncoderOptions{
				PrefixSegments:      tc.prefixSegments,
				ShortName:           &tc.shortName,
				Separator:           tc.separator,
				Region:              tc.region,
				StartsWithSeparator: tc.startsWithSeparator,
			}

			result, err := encoder.Encode(options)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}
