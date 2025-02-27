// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_AudioDolbyEDecode AWS CloudFormation Resource (AWS::MediaLive::Channel.AudioDolbyEDecode)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodolbyedecode.html
type Channel_AudioDolbyEDecode struct {

	// ProgramSelection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodolbyedecode.html#cfn-medialive-channel-audiodolbyedecode-programselection
	ProgramSelection *string `json:"ProgramSelection,omitempty"`

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
func (r *Channel_AudioDolbyEDecode) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.AudioDolbyEDecode"
}
