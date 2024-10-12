// Unit tests for the AWS package

package aws_test

import (
	"errors"
	"testing"

	"github.com/gmeligio/naming/aws"
	"github.com/stretchr/testify/require"
)

func TestNamingAws_ShortRegionsUnique(t *testing.T) {
	t.Parallel()

	require.Len(t, aws.ShortRegions, 33)

	hasUniqueValues := containsUniqueValues(aws.ShortRegions)
	require.True(t, hasUniqueValues)
}

func TestNamingAws_Name(t *testing.T) {
	testCases := []struct {
		name           string
		shortName      string
		separator      string
		expected       string
		expectedError  error
		prefixSegments []string
	}{
		{
			name:           "Simple name with no prefix",
			prefixSegments: []string{},
			shortName:      "service",
			separator:      "-",
			expected:       "service",
			expectedError:  nil,
		},
		{
			name:           "Multiple prefixes with simple name",
			prefixSegments: []string{"prod", "app"},
			shortName:      "service",
			separator:      "-",
			expected:       "prod-app-service",
			expectedError:  nil,
		},
		{
			name:           "Special characters in names",
			prefixSegments: []string{"prod$", "app#"},
			shortName:      "service*",
			separator:      "_",
			expected:       "prod$_app#_service*",
			expectedError:  nil,
		},
		{
			name:           "Empty short name",
			prefixSegments: []string{"prod", "app"},
			shortName:      "",
			separator:      "-",
			expected:       "prod-app-",
			expectedError:  errors.New("Passed an empty shortName but that's not supported. shortName is required to be non-empty."),
		},
		{
			name:           "Different separators",
			prefixSegments: []string{"prod", "app"},
			shortName:      "service",
			separator:      ".",
			expected:       "prod.app.service",
			expectedError:  nil,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			result, err := naming.Name(tc.prefixSegments, tc.shortName, tc.separator)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestNamingAws_Default(t *testing.T) {
	testCases := []struct {
		name           string
		prefixSegments []string
		shortName      string
		expected       string
	}{
		{
			name:           "Default naming with no prefix",
			prefixSegments: []string{},
			shortName:      "service",
			expected:       "service",
		},
		{
			name:           "Default naming with single prefix",
			prefixSegments: []string{"prod"},
			shortName:      "service",
			expected:       "prod-service",
		},
		{
			name:           "Default naming with multiple prefixes",
			prefixSegments: []string{"prod", "app"},
			shortName:      "service",
			expected:       "prod-app-service",
		},
		{
			name:           "Default naming with special characters",
			prefixSegments: []string{"prod$", "app#"},
			shortName:      "service*",
			expected:       "prod$-app#-service*",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.PrefixSegments = tc.prefixSegments
			result, err := naming.Default(tc.shortName)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestNamingAws_S3Bucket(t *testing.T) {
	testCases := []struct {
		name      string
		shortName string
		expected  string
	}{
		{
			name:      "Simple S3 name",
			shortName: "bucket",
			expected:  "bucket",
		},
		{
			name:      "S3 name with special characters",
			shortName: "bucket-name*",
			expected:  "bucket-name*",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			result, err := naming.S3Bucket(tc.shortName)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestNamingAws_SSMParameter(t *testing.T) {
	testCases := []struct {
		name           string
		prefixSegments []string
		shortName      string
		expected       string
	}{
		{
			name:           "Simple name with no prefix",
			prefixSegments: []string{},
			shortName:      "parameter",
			expected:       "parameter",
		},
		{
			name:           "Multiple prefixes with simple name",
			prefixSegments: []string{"prod", "app"},
			shortName:      "parameter",
			expected:       "prod/app/parameter",
		},
		{
			name:           "Special characters in names",
			prefixSegments: []string{"prod$", "app#"},
			shortName:      "parameter*",
			expected:       "prod$/app#/parameter*",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.PrefixSegments = tc.prefixSegments
			result, err := naming.SSMParameter(tc.shortName)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestNamingAws_WithRegionName(t *testing.T) {
	testCases := []struct {
		name           string
		shortName      string
		separator      string
		region         string
		expected       string
		useShortRegion bool
		expectedError  error
	}{
		{
			name:           "With full region name and default separator",
			shortName:      "service",
			separator:      "-",
			region:         "us-west-2",
			useShortRegion: false,
			expected:       "us-west-2-service",
			expectedError:  nil,
		},
		{
			name:           "With short region name and custom separator",
			shortName:      "service",
			separator:      "_",
			region:         "us-west-2",
			useShortRegion: true,
			expected:       "usw2_service",
			expectedError:  nil,
		},
		{
			name:           "With full region name and SSM delimiter",
			shortName:      "parameter",
			separator:      "/",
			region:         "eu-central-1",
			useShortRegion: false,
			expected:       "eu-central-1/parameter",
			expectedError:  nil,
		},
		{
			name:           "With short region name and no separator",
			shortName:      "config",
			separator:      "",
			region:         "ap-southeast-1",
			useShortRegion: true,
			expected:       "apse1config",
			expectedError:  nil,
		},
		{
			name:           "Default naming with non-existent region",
			shortName:      "service",
			region:         "non-existent",
			useShortRegion: true,
			expected:       "-service",
			expectedError:  errors.New("Region non-existent is not supported. Please create a new issue if it's a region that is supported by AWS."),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.UseShortRegion = tc.useShortRegion

			result, err := naming.WithRegionName(tc.shortName, tc.separator, tc.region)

			if tc.expectedError != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestNamingAws_WithRegionDefault(t *testing.T) {
	testCases := []struct {
		name           string
		shortName      string
		region         string
		useShortRegion bool
		expected       string
	}{
		{
			name:           "Default naming with region and short region name",
			shortName:      "service",
			region:         "us-west-2",
			useShortRegion: true,
			expected:       "usw2-service",
		},
		{
			name:           "Default naming with region and full region name",
			shortName:      "service",
			region:         "us-west-2",
			useShortRegion: false,
			expected:       "us-west-2-service",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.UseShortRegion = tc.useShortRegion
			result, err := naming.WithRegionDefault(tc.shortName, tc.region)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestNamingAws_WithRegionSSMParemeter(t *testing.T) {
	testCases := []struct {
		name           string
		shortName      string
		region         string
		useShortRegion bool
		expected       string
	}{
		{
			name:           "Simple name with region",
			shortName:      "service",
			region:         "us-west-2",
			useShortRegion: false,
			expected:       "us-west-2/service",
		},
		{
			name:           "Name with prefix and region",
			shortName:      "service",
			region:         "eu-central-1",
			useShortRegion: false,
			expected:       "eu-central-1/service",
		},
		{
			name:           "Name with short region",
			shortName:      "service",
			region:         "eu-central-1",
			useShortRegion: true,
			expected:       "euc1/service",
		},
		{
			name:           "Empty region",
			shortName:      "service",
			region:         "",
			useShortRegion: false,
			expected:       "/service",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.UseShortRegion = tc.useShortRegion

			result, err := naming.WithRegionSSMParameter(tc.shortName, tc.region)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestNamingAws_WithRegionS3Bucket(t *testing.T) {
	testCases := []struct {
		name           string
		shortName      string
		region         string
		useShortRegion bool
		expected       string
	}{
		{
			name:           "With short region",
			shortName:      "service",
			region:         "us-west-2",
			useShortRegion: true,
			expected:       "usw2-service",
		},
		{
			name:           "Without short region",
			shortName:      "service",
			region:         "us-west-2",
			useShortRegion: false,
			expected:       "us-west-2-service",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			naming := aws.NewNamingAws()
			naming.UseShortRegion = tc.useShortRegion

			result, err := naming.WithRegionS3Bucket(tc.shortName, tc.region)

			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

// containsUniqueValues checks if a map contains unique values
func containsUniqueValues(m map[string]string) bool {
	// Set to track seen values
	seen := make(map[string]bool)

	for _, value := range m {
		if _, ok := seen[value]; ok {
			// Value already seen, not unique
			return false
		}
		// Mark this value as seen
		seen[value] = true
	}
	return true
}
