// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kinesis

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	kinesis_sdkv1 "github.com/aws/aws-sdk-go/service/kinesis"
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
			Factory:  DataSourceStream,
			TypeName: "aws_kinesis_stream",
		},
		{
			Factory:  DataSourceStreamConsumer,
			TypeName: "aws_kinesis_stream_consumer",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceStream,
			TypeName: "aws_kinesis_stream",
			Name:     "Stream",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "name",
			},
		},
		{
			Factory:  ResourceStreamConsumer,
			TypeName: "aws_kinesis_stream_consumer",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kinesis
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *kinesis_sdkv1.Kinesis {
	return kinesis_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
