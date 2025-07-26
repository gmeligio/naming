package aws

// AWSNamer is responsible for validating and generating AWS names.
type AWSNamer struct {
	// Encoder is the AWSEncoder used to generate names.
	Encoder *AWSEncoder
	// Options holds the options for generating AWS names.
	Options *AWSEncoderOptions
}

// NewAWSNamer creates a new AWSNamer with default values.
func NewAWSNamer() *AWSNamer {
	return &AWSNamer{
		Encoder: NewAWSEncoder(),
		Options: &AWSEncoderOptions{},
	}
}

// WithEncoder sets the encoder for the AWSNamer.
func (n *AWSNamer) WithEncoder(encoder *AWSEncoder) *AWSNamer {
	n.Encoder = encoder
	return n
}

// WithPrefixSegments sets the prefix segments for the AWSNamer.
func (n *AWSNamer) WithPrefixSegments(prefixSegments []string) *AWSNamer {
	n.Options.PrefixSegments = prefixSegments
	return n
}

// WithRegion sets the region for the AWSNamer.
func (n *AWSNamer) WithRegion(region string) (*AWSNamer, error) {
	if region == "" {
		return nil, ErrEmptyRegion
	}

	n.Options.Region = &region

	return n, nil
}

// WithShortName sets the short name for the AWSNamer.
func (n *AWSNamer) WithShortName(shortName string) (*AWSNamer, error) {
	if shortName == "" {
		return nil, ErrEmptyShortName
	}

	n.Options.ShortName = &shortName
	return n, nil
}

// WithSeparator sets the separator for the AWSNamer.
func (n *AWSNamer) WithSeparator(separator string) *AWSNamer {
	n.Options.Separator = &separator

	return n
}

// WithStartsWithSeparator sets whether the name should start with the separator.
func (n *AWSNamer) WithStartsWithSeparator(startsWithSeparator bool) *AWSNamer {
	n.Options.StartsWithSeparator = startsWithSeparator
	return n
}

// Name generates the AWS name based on the provided options.
func (n *AWSNamer) Name() (string, error) {
	return n.Encoder.Encode(n.Options)
}
