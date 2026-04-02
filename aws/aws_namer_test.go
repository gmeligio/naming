package aws_test

import (
	"testing"

	"github.com/gmeligio/naming/aws"
	"github.com/stretchr/testify/require"
)

func TestAWSNamer_WithRegion(t *testing.T) {
	_, err := aws.NewAWSNamer().WithRegion("")
	require.Equal(t, aws.ErrEmptyRegion, err)
}

func TestAWSNamer_WithShortName(t *testing.T) {
	_, err := aws.NewAWSNamer().WithShortName("")
	require.Equal(t, aws.ErrEmptyShortName, err)
}
