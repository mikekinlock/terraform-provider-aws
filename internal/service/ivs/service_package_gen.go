// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ivs

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	ivs_sdkv1 "github.com/aws/aws-sdk-go/service/ivs"
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
			Factory:  DataSourceStreamKey,
			TypeName: "aws_ivs_stream_key",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceChannel,
			TypeName: "aws_ivs_channel",
			Name:     "Channel",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourcePlaybackKeyPair,
			TypeName: "aws_ivs_playback_key_pair",
			Name:     "Playback Key Pair",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceRecordingConfiguration,
			TypeName: "aws_ivs_recording_configuration",
			Name:     "Recording Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IVS
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) *ivs_sdkv1.IVS {
	return ivs_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)}))
}

var ServicePackage = &servicePackage{}
