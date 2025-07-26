// Package aws provides functionality for generating standardized AWS resource names
// following naming conventions for different AWS services such as S3 buckets and SSM parameters.
// It supports region-aware naming with options for short region codes, custom separators,
// and various naming patterns to ensure consistency across AWS resources.
package aws

import (
	"errors"
	"fmt"
)

const (
// s3BucketLimit   = 63
// lbBalancerLimit = 32
)

var (
	ErrEmptyRegion    = errors.New("found an empty region. Region must be non-empty")
	ErrEmptyShortName = errors.New("received an empty shortName but that's not supported. shortName is required to be non-empty")
)

var ShortRegions = map[string]string{
	"af-south-1":     "afs1",
	"ap-east-1":      "ape1",
	"ap-northeast-1": "apne1",
	"ap-northeast-2": "apne2",
	"ap-northeast-3": "apne3",
	"ap-south-1":     "aps1",
	"ap-south-2":     "aps2",
	"ap-southeast-1": "apse1",
	"ap-southeast-2": "apse2",
	"ap-southeast-3": "apse3",
	"ap-southeast-4": "apse4",
	"ca-central-1":   "cac1",
	"ca-west-1":      "caw1",
	"cn-north-1":     "cnn1",
	"cn-northwest-1": "cnnw1",
	"eu-central-1":   "euc1",
	"eu-central-2":   "euc2",
	"eu-north-1":     "eun1",
	"eu-south-1":     "eus1",
	"eu-south-2":     "eus2",
	"eu-west-1":      "euw1",
	"eu-west-2":      "euw2",
	"eu-west-3":      "euw3",
	"il-central-1":   "ilc1",
	"me-central-1":   "mec1",
	"me-south-1":     "mes1",
	"sa-east-1":      "sae1",
	"us-east-1":      "use1",
	"us-east-2":      "use2",
	"us-gov-east-1":  "usge1",
	"us-gov-west-1":  "usgw1",
	"us-west-1":      "usw1",
	"us-west-2":      "usw2",
}

type ErrNonSupportedRegion struct {
	Region string
	Err    error
}

func (e ErrNonSupportedRegion) Unwrap() error {
	return e.Err
}

func (e ErrNonSupportedRegion) Is(target error) bool {
	t, ok := target.(*ErrNonSupportedRegion)
	if !ok {
		return false
	}

	return errors.Is(e.Err, t.Err)
}

func (e ErrNonSupportedRegion) Error() string {
	return fmt.Sprintf("region %q is not supported. Please create a new issue if it's a region that is supported by AWS", e.Region)
}

func NewErrNonSupportedRegion(region string, err error) *ErrNonSupportedRegion {
	return &ErrNonSupportedRegion{
		Region: region,
		Err:    err,
	}
}
