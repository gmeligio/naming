# naming

Create easily infrastructure names using organization conventions and following cloud provider recommendations.

## Supported Conventions

- [ ] Custom
- [x] S3 bucket
- [x] SSM parameter
- [ ] LB balancer
- [ ] CloudWatch log group

## TODO

- [ ] Use a name generator for the region to simplify the generation algorithm used. See <https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/create/naming.go>

## Contributing

See [Contributing](docs/contributing).

## Acknowledgements

- <https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/core/lib/names.ts>
- <https://github.com/hashicorp/terraform-plugin-sdk/blob/main/helper/id/id.go>
- <https://github.com/cawcaw253/terraform-aws-namer>
- <https://github.com/cloudposse/terraform-null-label>
- <https://www.pulumi.com/docs/concepts/resources/names>
- <https://docs.aws.amazon.com/cdk/v2/guide/identifiers.html>
- <https://github.com/clouddrove/terraform-aws-labels>
