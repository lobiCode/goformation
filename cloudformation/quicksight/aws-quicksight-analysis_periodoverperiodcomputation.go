// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_PeriodOverPeriodComputation AWS CloudFormation Resource (AWS::QuickSight::Analysis.PeriodOverPeriodComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html
type Analysis_PeriodOverPeriodComputation struct {

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-computationid
	ComputationId string `json:"ComputationId"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-name
	Name *string `json:"Name,omitempty"`

	// Time AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-time
	Time *Analysis_DimensionField `json:"Time"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-value
	Value *Analysis_MeasureField `json:"Value,omitempty"`

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
func (r *Analysis_PeriodOverPeriodComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PeriodOverPeriodComputation"
}
