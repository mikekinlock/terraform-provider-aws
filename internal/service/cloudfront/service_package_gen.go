// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	cloudfront_sdkv1 "github.com/aws/aws-sdk-go/service/cloudfront"
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
			Factory:  DataSourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
		},
		{
			Factory:  DataSourceDistribution,
			TypeName: "aws_cloudfront_distribution",
		},
		{
			Factory:  DataSourceFunction,
			TypeName: "aws_cloudfront_function",
		},
		{
			Factory:  DataSourceLogDeliveryCanonicalUserID,
			TypeName: "aws_cloudfront_log_delivery_canonical_user_id",
		},
		{
			Factory:  DataSourceOriginAccessIdentities,
			TypeName: "aws_cloudfront_origin_access_identities",
		},
		{
			Factory:  DataSourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
		},
		{
			Factory:  DataSourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
		},
		{
			Factory:  DataSourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
		},
		{
			Factory:  DataSourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
		},
		{
			Factory:  ResourceDistribution,
			TypeName: "aws_cloudfront_distribution",
			Name:     "Distribution",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFieldLevelEncryptionConfig,
			TypeName: "aws_cloudfront_field_level_encryption_config",
		},
		{
			Factory:  ResourceFieldLevelEncryptionProfile,
			TypeName: "aws_cloudfront_field_level_encryption_profile",
		},
		{
			Factory:  ResourceFunction,
			TypeName: "aws_cloudfront_function",
		},
		{
			Factory:  ResourceKeyGroup,
			TypeName: "aws_cloudfront_key_group",
		},
		{
			Factory:  ResourceMonitoringSubscription,
			TypeName: "aws_cloudfront_monitoring_subscription",
		},
		{
			Factory:  ResourceOriginAccessControl,
			TypeName: "aws_cloudfront_origin_access_control",
		},
		{
			Factory:  ResourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
		},
		{
			Factory:  ResourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
		},
		{
			Factory:  ResourcePublicKey,
			TypeName: "aws_cloudfront_public_key",
		},
		{
			Factory:  ResourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
		},
		{
			Factory:  ResourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudFront
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *cloudfront_sdkv1.CloudFront {
	return cloudfront_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
