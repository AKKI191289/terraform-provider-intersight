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
)

// WorkflowApiAllOf Definition of the list of properties defined in 'workflow.Api', excluding properties defined in parent classes.
type WorkflowApiAllOf struct {
	// The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters.
	Body *string `json:"Body,omitempty"`
	// Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML and JSON. The type of the content that gets passed as payload and response in this API. * `json` - The type of content to be parsed, API response or device response, is inJSON format. * `xml` - The type of content to be parsed, API response or device response,is in XML format. * `text` - The type of content to be parsed, API response or device response, is inTEXT format.
	ContentType *string `json:"ContentType,omitempty"`
	// A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes     interface{}     `json:"Outcomes,omitempty"`
	ResponseSpec *ContentGrammar `json:"ResponseSpec,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
	// The delay in seconds after which the API needs to be executed. By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution. Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out.
	StartDelay *int64 `json:"StartDelay,omitempty"`
	// The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowApiAllOf WorkflowApiAllOf

// NewWorkflowApiAllOf instantiates a new WorkflowApiAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowApiAllOf() *WorkflowApiAllOf {
	this := WorkflowApiAllOf{}
	var contentType string = "json"
	this.ContentType = &contentType
	return &this
}

// NewWorkflowApiAllOfWithDefaults instantiates a new WorkflowApiAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowApiAllOfWithDefaults() *WorkflowApiAllOf {
	this := WorkflowApiAllOf{}
	var contentType string = "json"
	this.ContentType = &contentType
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *WorkflowApiAllOf) SetBody(v string) {
	o.Body = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *WorkflowApiAllOf) SetContentType(v string) {
	o.ContentType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowApiAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowApiAllOf) GetOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowApiAllOf) GetOutcomesOk() (*interface{}, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return &o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given interface{} and assigns it to the Outcomes field.
func (o *WorkflowApiAllOf) SetOutcomes(v interface{}) {
	o.Outcomes = v
}

// GetResponseSpec returns the ResponseSpec field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetResponseSpec() ContentGrammar {
	if o == nil || o.ResponseSpec == nil {
		var ret ContentGrammar
		return ret
	}
	return *o.ResponseSpec
}

// GetResponseSpecOk returns a tuple with the ResponseSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetResponseSpecOk() (*ContentGrammar, bool) {
	if o == nil || o.ResponseSpec == nil {
		return nil, false
	}
	return o.ResponseSpec, true
}

// HasResponseSpec returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasResponseSpec() bool {
	if o != nil && o.ResponseSpec != nil {
		return true
	}

	return false
}

// SetResponseSpec gets a reference to the given ContentGrammar and assigns it to the ResponseSpec field.
func (o *WorkflowApiAllOf) SetResponseSpec(v ContentGrammar) {
	o.ResponseSpec = &v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetSkipOnCondition() string {
	if o == nil || o.SkipOnCondition == nil {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || o.SkipOnCondition == nil {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasSkipOnCondition() bool {
	if o != nil && o.SkipOnCondition != nil {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowApiAllOf) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetStartDelay returns the StartDelay field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetStartDelay() int64 {
	if o == nil || o.StartDelay == nil {
		var ret int64
		return ret
	}
	return *o.StartDelay
}

// GetStartDelayOk returns a tuple with the StartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetStartDelayOk() (*int64, bool) {
	if o == nil || o.StartDelay == nil {
		return nil, false
	}
	return o.StartDelay, true
}

// HasStartDelay returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasStartDelay() bool {
	if o != nil && o.StartDelay != nil {
		return true
	}

	return false
}

// SetStartDelay gets a reference to the given int64 and assigns it to the StartDelay field.
func (o *WorkflowApiAllOf) SetStartDelay(v int64) {
	o.StartDelay = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *WorkflowApiAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApiAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *WorkflowApiAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *WorkflowApiAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o WorkflowApiAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.ResponseSpec != nil {
		toSerialize["ResponseSpec"] = o.ResponseSpec
	}
	if o.SkipOnCondition != nil {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.StartDelay != nil {
		toSerialize["StartDelay"] = o.StartDelay
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowApiAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowApiAllOf := _WorkflowApiAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowApiAllOf); err == nil {
		*o = WorkflowApiAllOf(varWorkflowApiAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Body")
		delete(additionalProperties, "ContentType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outcomes")
		delete(additionalProperties, "ResponseSpec")
		delete(additionalProperties, "SkipOnCondition")
		delete(additionalProperties, "StartDelay")
		delete(additionalProperties, "Timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowApiAllOf struct {
	value *WorkflowApiAllOf
	isSet bool
}

func (v NullableWorkflowApiAllOf) Get() *WorkflowApiAllOf {
	return v.value
}

func (v *NullableWorkflowApiAllOf) Set(val *WorkflowApiAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowApiAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowApiAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowApiAllOf(val *WorkflowApiAllOf) *NullableWorkflowApiAllOf {
	return &NullableWorkflowApiAllOf{value: val, isSet: true}
}

func (v NullableWorkflowApiAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowApiAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
