package aws_test

import (
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
