// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_FilterControl AWS CloudFormation Resource (AWS::QuickSight::Template.FilterControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html
type Template_FilterControl struct {

	// DateTimePicker AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-datetimepicker
	DateTimePicker *Template_FilterDateTimePickerControl `json:"DateTimePicker,omitempty"`

	// Dropdown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-dropdown
	Dropdown *Template_FilterDropDownControl `json:"Dropdown,omitempty"`

	// List AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-list
	List *Template_FilterListControl `json:"List,omitempty"`

	// RelativeDateTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-relativedatetime
	RelativeDateTime *Template_FilterRelativeDateTimeControl `json:"RelativeDateTime,omitempty"`

	// Slider AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-slider
	Slider *Template_FilterSliderControl `json:"Slider,omitempty"`

	// TextArea AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-textarea
	TextArea *Template_FilterTextAreaControl `json:"TextArea,omitempty"`

	// TextField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filtercontrol.html#cfn-quicksight-template-filtercontrol-textfield
	TextField *Template_FilterTextFieldControl `json:"TextField,omitempty"`

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
func (r *Template_FilterControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FilterControl"
}
