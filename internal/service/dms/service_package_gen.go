// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package dms

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	databasemigrationservice_sdkv1 "github.com/aws/aws-sdk-go/service/databasemigrationservice"
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
			Factory:  DataSourceCertificate,
			TypeName: "aws_dms_certificate",
		},
		{
			Factory:  DataSourceEndpoint,
			TypeName: "aws_dms_endpoint",
		},
		{
			Factory:  DataSourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
		},
		{
			Factory:  DataSourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
		},
		{
			Factory:  DataSourceReplicationTask,
			TypeName: "aws_dms_replication_task",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_dms_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "certificate_arn",
			},
		},
		{
			Factory:  ResourceEndpoint,
			TypeName: "aws_dms_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
		},
		{
			Factory:  ResourceEventSubscription,
			TypeName: "aws_dms_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
			Name:     "Replication Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_instance_arn",
			},
		},
		{
			Factory:  ResourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
			Name:     "Replication Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_subnet_group_arn",
			},
		},
		{
			Factory:  ResourceReplicationTask,
			TypeName: "aws_dms_replication_task",
			Name:     "Replication Task",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_task_arn",
			},
		},
		{
			Factory:  ResourceS3Endpoint,
			TypeName: "aws_dms_s3_endpoint",
			Name:     "S3 Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DMS
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *databasemigrationservice_sdkv1.DatabaseMigrationService {
	return databasemigrationservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
