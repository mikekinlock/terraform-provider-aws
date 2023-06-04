// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kafka

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	kafka_sdkv1 "github.com/aws/aws-sdk-go/service/kafka"
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
			Factory:  DataSourceBrokerNodes,
			TypeName: "aws_msk_broker_nodes",
		},
		{
			Factory:  DataSourceCluster,
			TypeName: "aws_msk_cluster",
		},
		{
			Factory:  DataSourceConfiguration,
			TypeName: "aws_msk_configuration",
		},
		{
			Factory:  DataSourceVersion,
			TypeName: "aws_msk_kafka_version",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCluster,
			TypeName: "aws_msk_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceConfiguration,
			TypeName: "aws_msk_configuration",
		},
		{
			Factory:  ResourceScramSecretAssociation,
			TypeName: "aws_msk_scram_secret_association",
		},
		{
			Factory:  ResourceServerlessCluster,
			TypeName: "aws_msk_serverless_cluster",
			Name:     "Serverless Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kafka
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *kafka_sdkv1.Kafka {
	return kafka_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
