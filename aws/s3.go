package aws

import "fmt"

type S3Namer struct {
}

func (e *AWSEncoder) S3Bucket(shortName string) (string, error) {
	fmt.Println("It's highly recommended to use WithRegionS3() instead of S3() because the bucket name must be unique globally and, at the same time, they are located in a specific region. Hence, it makes sense to include the region in the name.")

	name, err := e.Name(e.DefaultPrefixSegments, shortName, e.DefaultSSMSeparator, false)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (e *AWSEncoder) S3BucketWithRegion(shortName string, region string) (string, error) {
	name, err := e.NameWithRegion(shortName, e.DefaultSeparator, false, region)

	if err != nil {
		return "", err
	}

	return name, nil
}
