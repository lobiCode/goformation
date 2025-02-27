// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_SectionBasedLayoutPaperCanvasSizeOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.SectionBasedLayoutPaperCanvasSizeOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions.html
type Analysis_SectionBasedLayoutPaperCanvasSizeOptions struct {

	// PaperMargin AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions-papermargin
	PaperMargin *Analysis_Spacing `json:"PaperMargin,omitempty"`

	// PaperOrientation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions-paperorientation
	PaperOrientation *string `json:"PaperOrientation,omitempty"`

	// PaperSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-analysis-sectionbasedlayoutpapercanvassizeoptions-papersize
	PaperSize *string `json:"PaperSize,omitempty"`

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
func (r *Analysis_SectionBasedLayoutPaperCanvasSizeOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.SectionBasedLayoutPaperCanvasSizeOptions"
}
