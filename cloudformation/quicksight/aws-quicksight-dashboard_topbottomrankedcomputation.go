// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_TopBottomRankedComputation AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TopBottomRankedComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html
type Dashboard_TopBottomRankedComputation struct {

	// Category AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-category
	Category *Dashboard_DimensionField `json:"Category"`

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-computationid
	ComputationId string `json:"ComputationId"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-name
	Name *string `json:"Name,omitempty"`

	// ResultSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-resultsize
	ResultSize *float64 `json:"ResultSize,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topbottomrankedcomputation.html#cfn-quicksight-dashboard-topbottomrankedcomputation-value
	Value *Dashboard_MeasureField `json:"Value,omitempty"`

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
func (r *Dashboard_TopBottomRankedComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TopBottomRankedComputation"
}
