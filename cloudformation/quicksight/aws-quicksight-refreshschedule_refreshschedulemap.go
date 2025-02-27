// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// RefreshSchedule_RefreshScheduleMap AWS CloudFormation Resource (AWS::QuickSight::RefreshSchedule.RefreshScheduleMap)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html
type RefreshSchedule_RefreshScheduleMap struct {

	// RefreshType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-refreshtype
	RefreshType *string `json:"RefreshType,omitempty"`

	// ScheduleFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-schedulefrequency
	ScheduleFrequency *RefreshSchedule_ScheduleFrequency `json:"ScheduleFrequency,omitempty"`

	// ScheduleId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-scheduleid
	ScheduleId *string `json:"ScheduleId,omitempty"`

	// StartAfterDateTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-startafterdatetime
	StartAfterDateTime *string `json:"StartAfterDateTime,omitempty"`

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
func (r *RefreshSchedule_RefreshScheduleMap) AWSCloudFormationType() string {
	return "AWS::QuickSight::RefreshSchedule.RefreshScheduleMap"
}
