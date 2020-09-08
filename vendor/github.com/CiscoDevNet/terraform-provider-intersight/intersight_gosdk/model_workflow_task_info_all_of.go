/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"time"
)

// WorkflowTaskInfoAllOf Definition of the list of properties defined in 'workflow.TaskInfo', excluding properties defined in parent classes.
type WorkflowTaskInfoAllOf struct {
	// The task description and this is the description that was added when the task was included into the workflow.
	Description *string `json:"Description,omitempty"`
	// The time stamp when the task reached a final state.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Description of the reason why the task failed.
	FailureReason *string `json:"FailureReason,omitempty"`
	// The input data that was sent to the task at the start of execution.
	Input interface{} `json:"Input,omitempty"`
	// The instance ID of the task running in the workflow engine.
	InstId *string `json:"InstId,omitempty"`
	// Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.
	Internal *bool `json:"Internal,omitempty"`
	// User friendly short label to describe this task instance in the workflow.
	Label   *string            `json:"Label,omitempty"`
	Message *[]WorkflowMessage `json:"Message,omitempty"`
	// Task definition name which specifies the task type.
	Name *string `json:"Name,omitempty"`
	// The output data that was generated by this task.
	Output interface{} `json:"Output,omitempty"`
	// The task reference name to ensure its unique inside a workflow.
	RefName *string `json:"RefName,omitempty"`
	// A counter for number of retries.
	RetryCount *int64 `json:"RetryCount,omitempty"`
	// The instance ID of the task that is currently being executed. When retrying a workflow with failed tasks, the task in workflow engine will have a new instance ID, but the task may still be in-progress. In this case, the task instId reflects the instance ID in the workflow engine, while runningInstId reflects the instance ID of the instance that is currently being executed.
	RunningInstId *string `json:"RunningInstId,omitempty"`
	// The time stamp when the task started execution.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The status of the task and this will specify if the task is running or has reached a final state.
	Status               *string                             `json:"Status,omitempty"`
	TaskInstIdList       *[]WorkflowTaskRetryInfo            `json:"TaskInstIdList,omitempty"`
	SubWorkflowInfo      *WorkflowWorkflowInfoRelationship   `json:"SubWorkflowInfo,omitempty"`
	TaskDefinition       *WorkflowTaskDefinitionRelationship `json:"TaskDefinition,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship   `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskInfoAllOf WorkflowTaskInfoAllOf

// NewWorkflowTaskInfoAllOf instantiates a new WorkflowTaskInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskInfoAllOf() *WorkflowTaskInfoAllOf {
	this := WorkflowTaskInfoAllOf{}
	return &this
}

// NewWorkflowTaskInfoAllOfWithDefaults instantiates a new WorkflowTaskInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskInfoAllOfWithDefaults() *WorkflowTaskInfoAllOf {
	this := WorkflowTaskInfoAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowTaskInfoAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WorkflowTaskInfoAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *WorkflowTaskInfoAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskInfoAllOf) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskInfoAllOf) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowTaskInfoAllOf) SetInput(v interface{}) {
	o.Input = v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetInstId() string {
	if o == nil || o.InstId == nil {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetInstIdOk() (*string, bool) {
	if o == nil || o.InstId == nil {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasInstId() bool {
	if o != nil && o.InstId != nil {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *WorkflowTaskInfoAllOf) SetInstId(v string) {
	o.InstId = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *WorkflowTaskInfoAllOf) SetInternal(v bool) {
	o.Internal = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowTaskInfoAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetMessage() []WorkflowMessage {
	if o == nil || o.Message == nil {
		var ret []WorkflowMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetMessageOk() (*[]WorkflowMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given []WorkflowMessage and assigns it to the Message field.
func (o *WorkflowTaskInfoAllOf) SetMessage(v []WorkflowMessage) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskInfoAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskInfoAllOf) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskInfoAllOf) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowTaskInfoAllOf) SetOutput(v interface{}) {
	o.Output = v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetRefName() string {
	if o == nil || o.RefName == nil {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetRefNameOk() (*string, bool) {
	if o == nil || o.RefName == nil {
		return nil, false
	}
	return o.RefName, true
}

// HasRefName returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasRefName() bool {
	if o != nil && o.RefName != nil {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *WorkflowTaskInfoAllOf) SetRefName(v string) {
	o.RefName = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *WorkflowTaskInfoAllOf) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRunningInstId returns the RunningInstId field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetRunningInstId() string {
	if o == nil || o.RunningInstId == nil {
		var ret string
		return ret
	}
	return *o.RunningInstId
}

// GetRunningInstIdOk returns a tuple with the RunningInstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetRunningInstIdOk() (*string, bool) {
	if o == nil || o.RunningInstId == nil {
		return nil, false
	}
	return o.RunningInstId, true
}

// HasRunningInstId returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasRunningInstId() bool {
	if o != nil && o.RunningInstId != nil {
		return true
	}

	return false
}

// SetRunningInstId gets a reference to the given string and assigns it to the RunningInstId field.
func (o *WorkflowTaskInfoAllOf) SetRunningInstId(v string) {
	o.RunningInstId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WorkflowTaskInfoAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowTaskInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskInstIdList returns the TaskInstIdList field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetTaskInstIdList() []WorkflowTaskRetryInfo {
	if o == nil || o.TaskInstIdList == nil {
		var ret []WorkflowTaskRetryInfo
		return ret
	}
	return *o.TaskInstIdList
}

// GetTaskInstIdListOk returns a tuple with the TaskInstIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetTaskInstIdListOk() (*[]WorkflowTaskRetryInfo, bool) {
	if o == nil || o.TaskInstIdList == nil {
		return nil, false
	}
	return o.TaskInstIdList, true
}

// HasTaskInstIdList returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasTaskInstIdList() bool {
	if o != nil && o.TaskInstIdList != nil {
		return true
	}

	return false
}

// SetTaskInstIdList gets a reference to the given []WorkflowTaskRetryInfo and assigns it to the TaskInstIdList field.
func (o *WorkflowTaskInfoAllOf) SetTaskInstIdList(v []WorkflowTaskRetryInfo) {
	o.TaskInstIdList = &v
}

// GetSubWorkflowInfo returns the SubWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetSubWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.SubWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.SubWorkflowInfo
}

// GetSubWorkflowInfoOk returns a tuple with the SubWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetSubWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.SubWorkflowInfo == nil {
		return nil, false
	}
	return o.SubWorkflowInfo, true
}

// HasSubWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasSubWorkflowInfo() bool {
	if o != nil && o.SubWorkflowInfo != nil {
		return true
	}

	return false
}

// SetSubWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the SubWorkflowInfo field.
func (o *WorkflowTaskInfoAllOf) SetSubWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.SubWorkflowInfo = &v
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || o.TaskDefinition == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.TaskDefinition == nil {
		return nil, false
	}
	return o.TaskDefinition, true
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition != nil {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowTaskInfoAllOf) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowTaskInfoAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowTaskInfoAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowTaskInfoAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowTaskInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.InstId != nil {
		toSerialize["InstId"] = o.InstId
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.RefName != nil {
		toSerialize["RefName"] = o.RefName
	}
	if o.RetryCount != nil {
		toSerialize["RetryCount"] = o.RetryCount
	}
	if o.RunningInstId != nil {
		toSerialize["RunningInstId"] = o.RunningInstId
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskInstIdList != nil {
		toSerialize["TaskInstIdList"] = o.TaskInstIdList
	}
	if o.SubWorkflowInfo != nil {
		toSerialize["SubWorkflowInfo"] = o.SubWorkflowInfo
	}
	if o.TaskDefinition != nil {
		toSerialize["TaskDefinition"] = o.TaskDefinition
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTaskInfoAllOf := _WorkflowTaskInfoAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTaskInfoAllOf); err == nil {
		*o = WorkflowTaskInfoAllOf(varWorkflowTaskInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "InstId")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "RefName")
		delete(additionalProperties, "RetryCount")
		delete(additionalProperties, "RunningInstId")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskInstIdList")
		delete(additionalProperties, "SubWorkflowInfo")
		delete(additionalProperties, "TaskDefinition")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskInfoAllOf struct {
	value *WorkflowTaskInfoAllOf
	isSet bool
}

func (v NullableWorkflowTaskInfoAllOf) Get() *WorkflowTaskInfoAllOf {
	return v.value
}

func (v *NullableWorkflowTaskInfoAllOf) Set(val *WorkflowTaskInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskInfoAllOf(val *WorkflowTaskInfoAllOf) *NullableWorkflowTaskInfoAllOf {
	return &NullableWorkflowTaskInfoAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
