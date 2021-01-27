// Code generated by aws/endpoints/v3model_codegen.go. DO NOT EDIT.

package endpoints

import (
	"regexp"
)

// Partition identifiers
const (
	AwsPartitionID      = "aws"        // AWS Standard partition.
	AwsCnPartitionID    = "aws-cn"     // AWS China partition.
	AwsUsGovPartitionID = "aws-us-gov" // AWS GovCloud (US) partition.
	AwsIsoPartitionID   = "aws-iso"    // AWS ISO (US) partition.
	AwsIsoBPartitionID  = "aws-iso-b"  // AWS ISOB (US) partition.
)

// AWS Standard partition's regions.
const (
	AfSouth1RegionID     = "af-south-1"     // Africa (Cape Town).
	ApEast1RegionID      = "ap-east-1"      // Asia Pacific (Hong Kong).
	ApNortheast1RegionID = "ap-northeast-1" // Asia Pacific (Tokyo).
	ApNortheast2RegionID = "ap-northeast-2" // Asia Pacific (Seoul).
	ApSouth1RegionID     = "ap-south-1"     // Asia Pacific (Mumbai).
	ApSoutheast1RegionID = "ap-southeast-1" // Asia Pacific (Singapore).
	ApSoutheast2RegionID = "ap-southeast-2" // Asia Pacific (Sydney).
	CaCentral1RegionID   = "ca-central-1"   // Canada (Central).
	EuCentral1RegionID   = "eu-central-1"   // Europe (Frankfurt).
	EuNorth1RegionID     = "eu-north-1"     // Europe (Stockholm).
	EuSouth1RegionID     = "eu-south-1"     // Europe (Milan).
	EuWest1RegionID      = "eu-west-1"      // Europe (Ireland).
	EuWest2RegionID      = "eu-west-2"      // Europe (London).
	EuWest3RegionID      = "eu-west-3"      // Europe (Paris).
	MeSouth1RegionID     = "me-south-1"     // Middle East (Bahrain).
	SaEast1RegionID      = "sa-east-1"      // South America (Sao Paulo).
	UsEast1RegionID      = "us-east-1"      // US East (N. Virginia).
	UsEast2RegionID      = "us-east-2"      // US East (Ohio).
	UsWest1RegionID      = "us-west-1"      // US West (N. California).
	UsWest2RegionID      = "us-west-2"      // US West (Oregon).
)

// AWS China partition's regions.
const (
	CnNorth1RegionID     = "cn-north-1"     // China (Beijing).
	CnNorthwest1RegionID = "cn-northwest-1" // China (Ningxia).
)

// AWS GovCloud (US) partition's regions.
const (
	UsGovEast1RegionID = "us-gov-east-1" // AWS GovCloud (US-East).
	UsGovWest1RegionID = "us-gov-west-1" // AWS GovCloud (US-West).
)

// AWS ISO (US) partition's regions.
const (
	UsIsoEast1RegionID = "us-iso-east-1" // US ISO East.
)

// AWS ISOB (US) partition's regions.
const (
	UsIsobEast1RegionID = "us-isob-east-1" // US ISOB East (Ohio).
)

// DefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: AWS Standard, AWS China, AWS GovCloud (US), AWS ISO (US), and AWS ISOB (US).
//
// Use DefaultPartitions() to get the list of the default partitions.
func DefaultResolver() Resolver {
	return defaultPartitions
}

// DefaultPartitions returns a list of the partitions the SDK is bundled
// with. The available partitions are: AWS Standard, AWS China, AWS GovCloud (US), AWS ISO (US), and AWS ISOB (US).
//
//    partitions := endpoints.DefaultPartitions
//    for _, p := range partitions {
//        // ... inspect partitions
//    }
func DefaultPartitions() []Partition {
	return defaultPartitions.Partitions()
}

var defaultPartitions = partitions{
	awsPartition,
}

// AwsPartition returns the Resolver for AWS Standard.
func AwsPartition() Partition {
	return awsPartition.Partition()
}

var awsPartition = partition{
	ID:        "aws",
	Name:      "AWS Standard",
	DNSSuffix: "amazonaws.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{service}.{region}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"v4"},
	},
	Regions: regions{
		"af-south-1": region{
			Description: "Africa (Cape Town)",
		},
		"ap-east-1": region{
			Description: "Asia Pacific (Hong Kong)",
		},
		"ap-northeast-1": region{
			Description: "Asia Pacific (Tokyo)",
		},
		"ap-northeast-2": region{
			Description: "Asia Pacific (Seoul)",
		},
		"ap-south-1": region{
			Description: "Asia Pacific (Mumbai)",
		},
		"ap-southeast-1": region{
			Description: "Asia Pacific (Singapore)",
		},
		"ap-southeast-2": region{
			Description: "Asia Pacific (Sydney)",
		},
		"ca-central-1": region{
			Description: "Canada (Central)",
		},
		"eu-central-1": region{
			Description: "Europe (Frankfurt)",
		},
		"eu-north-1": region{
			Description: "Europe (Stockholm)",
		},
		"eu-south-1": region{
			Description: "Europe (Milan)",
		},
		"eu-west-1": region{
			Description: "Europe (Ireland)",
		},
		"eu-west-2": region{
			Description: "Europe (London)",
		},
		"eu-west-3": region{
			Description: "Europe (Paris)",
		},
		"me-south-1": region{
			Description: "Middle East (Bahrain)",
		},
		"sa-east-1": region{
			Description: "South America (Sao Paulo)",
		},
		"us-east-1": region{
			Description: "US East (N. Virginia)",
		},
		"us-east-2": region{
			Description: "US East (Ohio)",
		},
		"us-west-1": region{
			Description: "US West (N. California)",
		},
		"us-west-2": region{
			Description: "US West (Oregon)",
		},
	},
	Services: services{
		"appsync": service{
			Endpoints: endpoints{
				"ap-east-1":      endpoint{},
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-north-1":     endpoint{},
				"eu-south-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"me-south-1":     endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cognito-identity": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-north-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"fips-us-east-1": endpoint{
					Hostname: "cognito-identity-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"fips-us-east-2": endpoint{
					Hostname: "cognito-identity-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"fips-us-west-2": endpoint{
					Hostname: "cognito-identity-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
				"sa-east-1": endpoint{},
				"us-east-1": endpoint{},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"cognito-idp": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-north-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"fips-us-east-1": endpoint{
					Hostname: "cognito-idp-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"fips-us-east-2": endpoint{
					Hostname: "cognito-idp-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"fips-us-west-2": endpoint{
					Hostname: "cognito-idp-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
				"sa-east-1": endpoint{},
				"us-east-1": endpoint{},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"cognito-sync": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"dynamodb": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"af-south-1":     endpoint{},
				"ap-east-1":      endpoint{},
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"ca-central-1-fips": endpoint{
					Hostname: "dynamodb-fips.ca-central-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "ca-central-1",
					},
				},
				"eu-central-1": endpoint{},
				"eu-north-1":   endpoint{},
				"eu-south-1":   endpoint{},
				"eu-west-1":    endpoint{},
				"eu-west-2":    endpoint{},
				"eu-west-3":    endpoint{},
				"local": endpoint{
					Hostname:  "localhost:8000",
					Protocols: []string{"http"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"me-south-1": endpoint{},
				"sa-east-1":  endpoint{},
				"us-east-1":  endpoint{},
				"us-east-1-fips": endpoint{
					Hostname: "dynamodb-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-east-2": endpoint{},
				"us-east-2-fips": endpoint{
					Hostname: "dynamodb-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"us-west-1": endpoint{},
				"us-west-1-fips": endpoint{
					Hostname: "dynamodb-fips.us-west-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"us-west-2": endpoint{},
				"us-west-2-fips": endpoint{
					Hostname: "dynamodb-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
			},
		},
		"identitystore": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-north-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"lambda": service{

			Endpoints: endpoints{
				"af-south-1":     endpoint{},
				"ap-east-1":      endpoint{},
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-north-1":     endpoint{},
				"eu-south-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"fips-us-east-1": endpoint{
					Hostname: "lambda-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"fips-us-east-2": endpoint{
					Hostname: "lambda-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"fips-us-west-1": endpoint{
					Hostname: "lambda-fips.us-west-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"fips-us-west-2": endpoint{
					Hostname: "lambda-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
				"me-south-1": endpoint{},
				"sa-east-1":  endpoint{},
				"us-east-1":  endpoint{},
				"us-east-2":  endpoint{},
				"us-west-1":  endpoint{},
				"us-west-2":  endpoint{},
			},
		},
		"s3": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedTrue,
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},

				HasDualStack:      boxedTrue,
				DualStackHostname: "{service}.dualstack.{region}.{dnsSuffix}",
			},
			Endpoints: endpoints{
				"af-south-1": endpoint{},
				"ap-east-1":  endpoint{},
				"ap-northeast-1": endpoint{
					Hostname:          "s3.ap-northeast-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{
					Hostname:          "s3.ap-southeast-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ap-southeast-2": endpoint{
					Hostname:          "s3.ap-southeast-2.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"aws-global": endpoint{
					Hostname:          "s3.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"ca-central-1": endpoint{},
				"eu-central-1": endpoint{},
				"eu-north-1":   endpoint{},
				"eu-south-1":   endpoint{},
				"eu-west-1": endpoint{
					Hostname:          "s3.eu-west-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"eu-west-2":  endpoint{},
				"eu-west-3":  endpoint{},
				"me-south-1": endpoint{},
				"s3-external-1": endpoint{
					Hostname:          "s3-external-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"sa-east-1": endpoint{
					Hostname:          "s3.sa-east-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-east-1": endpoint{
					Hostname:          "s3.us-east-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{
					Hostname:          "s3.us-west-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-west-2": endpoint{
					Hostname:          "s3.us-west-2.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
			},
		},
		"s3-control": service{
			Defaults: endpoint{
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3v4"},

				HasDualStack:      boxedTrue,
				DualStackHostname: "{service}.dualstack.{region}.{dnsSuffix}",
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{
					Hostname:          "s3-control.ap-northeast-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ap-northeast-1",
					},
				},
				"ap-northeast-2": endpoint{
					Hostname:          "s3-control.ap-northeast-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ap-northeast-2",
					},
				},
				"ap-south-1": endpoint{
					Hostname:          "s3-control.ap-south-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ap-south-1",
					},
				},
				"ap-southeast-1": endpoint{
					Hostname:          "s3-control.ap-southeast-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ap-southeast-1",
					},
				},
				"ap-southeast-2": endpoint{
					Hostname:          "s3-control.ap-southeast-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ap-southeast-2",
					},
				},
				"ca-central-1": endpoint{
					Hostname:          "s3-control.ca-central-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ca-central-1",
					},
				},
				"ca-central-1-fips": endpoint{
					Hostname:          "s3-control-fips.ca-central-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "ca-central-1",
					},
				},
				"eu-central-1": endpoint{
					Hostname:          "s3-control.eu-central-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "eu-central-1",
					},
				},
				"eu-north-1": endpoint{
					Hostname:          "s3-control.eu-north-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "eu-north-1",
					},
				},
				"eu-west-1": endpoint{
					Hostname:          "s3-control.eu-west-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "eu-west-1",
					},
				},
				"eu-west-2": endpoint{
					Hostname:          "s3-control.eu-west-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "eu-west-2",
					},
				},
				"eu-west-3": endpoint{
					Hostname:          "s3-control.eu-west-3.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "eu-west-3",
					},
				},
				"sa-east-1": endpoint{
					Hostname:          "s3-control.sa-east-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "sa-east-1",
					},
				},
				"us-east-1": endpoint{
					Hostname:          "s3-control.us-east-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-east-1-fips": endpoint{
					Hostname:          "s3-control-fips.us-east-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-east-2": endpoint{
					Hostname:          "s3-control.us-east-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"us-east-2-fips": endpoint{
					Hostname:          "s3-control-fips.us-east-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"us-west-1": endpoint{
					Hostname:          "s3-control.us-west-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"us-west-1-fips": endpoint{
					Hostname:          "s3-control-fips.us-west-1.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"us-west-2": endpoint{
					Hostname:          "s3-control.us-west-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
				"us-west-2-fips": endpoint{
					Hostname:          "s3-control-fips.us-west-2.amazonaws.com",
					SignatureVersions: []string{"s3v4"},
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
			},
		},
		"session.qldb": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"streams.dynamodb": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
				CredentialScope: credentialScope{
					Service: "dynamodb",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"ca-central-1-fips": endpoint{
					Hostname: "dynamodb-fips.ca-central-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "ca-central-1",
					},
				},
				"eu-central-1": endpoint{},
				"eu-north-1":   endpoint{},
				"eu-west-1":    endpoint{},
				"eu-west-2":    endpoint{},
				"eu-west-3":    endpoint{},
				"local": endpoint{
					Hostname:  "localhost:8000",
					Protocols: []string{"http"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"me-south-1": endpoint{},
				"sa-east-1":  endpoint{},
				"us-east-1":  endpoint{},
				"us-east-1-fips": endpoint{
					Hostname: "dynamodb-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-east-2": endpoint{},
				"us-east-2-fips": endpoint{
					Hostname: "dynamodb-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"us-west-1": endpoint{},
				"us-west-1-fips": endpoint{
					Hostname: "dynamodb-fips.us-west-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"us-west-2": endpoint{},
				"us-west-2-fips": endpoint{
					Hostname: "dynamodb-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
			},
		},
	},
}
