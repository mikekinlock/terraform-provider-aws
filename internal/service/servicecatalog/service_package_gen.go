// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	servicecatalog_sdkv1 "github.com/aws/aws-sdk-go/service/servicecatalog"
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
			Factory:  DataSourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
		},
		{
			Factory:  DataSourceLaunchPaths,
			TypeName: "aws_servicecatalog_launch_paths",
		},
		{
			Factory:  DataSourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
		},
		{
			Factory:  DataSourcePortfolioConstraints,
			TypeName: "aws_servicecatalog_portfolio_constraints",
		},
		{
			Factory:  DataSourceProduct,
			TypeName: "aws_servicecatalog_product",
		},
		{
			Factory:  DataSourceProvisioningArtifacts,
			TypeName: "aws_servicecatalog_provisioning_artifacts",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBudgetResourceAssociation,
			TypeName: "aws_servicecatalog_budget_resource_association",
		},
		{
			Factory:  ResourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
		},
		{
			Factory:  ResourceOrganizationsAccess,
			TypeName: "aws_servicecatalog_organizations_access",
		},
		{
			Factory:  ResourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourcePortfolioShare,
			TypeName: "aws_servicecatalog_portfolio_share",
		},
		{
			Factory:  ResourcePrincipalPortfolioAssociation,
			TypeName: "aws_servicecatalog_principal_portfolio_association",
		},
		{
			Factory:  ResourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceProductPortfolioAssociation,
			TypeName: "aws_servicecatalog_product_portfolio_association",
		},
		{
			Factory:  ResourceProvisionedProduct,
			TypeName: "aws_servicecatalog_provisioned_product",
			Name:     "Provisioned Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceProvisioningArtifact,
			TypeName: "aws_servicecatalog_provisioning_artifact",
		},
		{
			Factory:  ResourceServiceAction,
			TypeName: "aws_servicecatalog_service_action",
		},
		{
			Factory:  ResourceTagOption,
			TypeName: "aws_servicecatalog_tag_option",
		},
		{
			Factory:  ResourceTagOptionResourceAssociation,
			TypeName: "aws_servicecatalog_tag_option_resource_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceCatalog
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *servicecatalog_sdkv1.ServiceCatalog {
	return servicecatalog_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
