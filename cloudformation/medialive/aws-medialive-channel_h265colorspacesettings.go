// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_H265ColorSpaceSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.H265ColorSpaceSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html
type Channel_H265ColorSpaceSettings struct {

	// ColorSpacePassthroughSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-colorspacepassthroughsettings
	ColorSpacePassthroughSettings *Channel_ColorSpacePassthroughSettings `json:"ColorSpacePassthroughSettings,omitempty"`

	// DolbyVision81Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-dolbyvision81settings
	DolbyVision81Settings *Channel_DolbyVision81Settings `json:"DolbyVision81Settings,omitempty"`

	// Hdr10Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-hdr10settings
	Hdr10Settings *Channel_Hdr10Settings `json:"Hdr10Settings,omitempty"`

	// Rec601Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-rec601settings
	Rec601Settings *Channel_Rec601Settings `json:"Rec601Settings,omitempty"`

	// Rec709Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-rec709settings
	Rec709Settings *Channel_Rec709Settings `json:"Rec709Settings,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Channel_H265ColorSpaceSettings) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.H265ColorSpaceSettings"
}
