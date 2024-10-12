package aws

import (
	"fmt"
	"slices"
	"strings"
)

const (
	s3BucketLimit  = 63
	lbBalancerLimit = 32
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

type NamingAws struct {
	DefaultDelimiter string
	SSMDelimiter     string
	PrefixSegments   []string
	SuffixLength     int
	UseShortRegion   bool
}

func NewNamingAws() *NamingAws {
	return &NamingAws{
		PrefixSegments:   []string{},
		DefaultDelimiter: "-",
		SSMDelimiter:     "/",
		SuffixLength:     4,
	}
}

func (n *NamingAws) Name(prefixSegments []string, shortName string, separator string) (string, error) {
	if shortName == "" {
		return "", fmt.Errorf("Passed an empty shortName but that's not supported. shortName is required to be non-empty.")
	}

	var segments []string
	segments = append(segments, prefixSegments...)
	segments = append(segments, shortName)

	return strings.Join(segments, separator), nil
}

func (n *NamingAws) Default(shortName string) (string, error) {
	name, err := n.Name(n.PrefixSegments, shortName, n.DefaultDelimiter)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) S3Bucket(shortName string) (string, error) {
	fmt.Println("It's highly recommended to use WithRegionS3() instead of S3() because the bucket name must be unique globally and, at the same time, they are located in a specific region. Hence, it makes sense to include the region in the name.")

	name, err := n.Name(n.PrefixSegments, shortName, n.SSMDelimiter)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) SSMParameter(shortName string) (string, error) {
	name, err := n.Name(n.PrefixSegments, shortName, n.SSMDelimiter)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) WithRegionName(shortName string, separator string, region string) (string, error) {
	regionSegment := region
	if n.UseShortRegion {
		if foundRegion, ok := ShortRegions[region]; ok {
			regionSegment = foundRegion
		} else {
			return "", fmt.Errorf("Region %s is not supported. Please create a new issue if it's a region that is supported by AWS.", region)
		}
	}

	segments := slices.Insert(n.PrefixSegments, 0, regionSegment)

	name, err := n.Name(segments, shortName, separator)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) WithRegionDefault(shortName string, region string) (string, error) {
	name, err := n.WithRegionName(shortName, n.DefaultDelimiter, region)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) WithRegionSSMParameter(shortName string, region string) (string, error) {
	name, err := n.WithRegionName(shortName, n.SSMDelimiter, region)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (n *NamingAws) WithRegionS3Bucket(shortName string, region string) (string, error) {
	name, err := n.WithRegionName(shortName, n.DefaultDelimiter, region)

	if err != nil {
		return "", err
	}

	return name, nil
}
