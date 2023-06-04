// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ecr

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	ecr_sdkv1 "github.com/aws/aws-sdk-go/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	endpoint string
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAuthorizationToken,
			TypeName: "aws_ecr_authorization_token",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_ecr_image",
		},
		{
			Factory:  DataSourcePullThroughCacheRule,
			TypeName: "aws_ecr_pull_through_cache_rule",
		},
		{
			Factory:  DataSourceRepository,
			TypeName: "aws_ecr_repository",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLifecyclePolicy,
			TypeName: "aws_ecr_lifecycle_policy",
		},
		{
			Factory:  ResourcePullThroughCacheRule,
			TypeName: "aws_ecr_pull_through_cache_rule",
		},
		{
			Factory:  ResourceRegistryPolicy,
			TypeName: "aws_ecr_registry_policy",
		},
		{
			Factory:  ResourceRegistryScanningConfiguration,
			TypeName: "aws_ecr_registry_scanning_configuration",
		},
		{
			Factory:  ResourceReplicationConfiguration,
			TypeName: "aws_ecr_replication_configuration",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_ecr_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRepositoryPolicy,
			TypeName: "aws_ecr_repository_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECR
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *ecr_sdkv1.ECR {
	return ecr_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
