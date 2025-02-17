/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// PolicyAbstractConfigResultEntryAllOf Definition of the list of properties defined in 'policy.AbstractConfigResultEntry', excluding properties defined in parent classes.
type PolicyAbstractConfigResultEntryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The completed time of the task in installer.
	CompletedTime *string                           `json:"CompletedTime,omitempty"`
	Context       NullablePolicyConfigResultContext `json:"Context,omitempty"`
	// Localized message based on the locale setting of the user's context.
	Message *string `json:"Message,omitempty"`
	// The identifier of the object that owns the result message. The owner ID is used to correlate a given result entry to a task or entity. For example, a config result entry that describes the result of a workflow task may have the task's instance ID as the owner.
	OwnerId *string `json:"OwnerId,omitempty"`
	// Values  -- Ok, Ok-with-warning, Errored.
	State *string `json:"State,omitempty"`
	// Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigResultEntryAllOf PolicyAbstractConfigResultEntryAllOf

// NewPolicyAbstractConfigResultEntryAllOf instantiates a new PolicyAbstractConfigResultEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigResultEntryAllOf(classId string, objectType string) *PolicyAbstractConfigResultEntryAllOf {
	this := PolicyAbstractConfigResultEntryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyAbstractConfigResultEntryAllOfWithDefaults instantiates a new PolicyAbstractConfigResultEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigResultEntryAllOfWithDefaults() *PolicyAbstractConfigResultEntryAllOf {
	this := PolicyAbstractConfigResultEntryAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractConfigResultEntryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractConfigResultEntryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractConfigResultEntryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractConfigResultEntryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompletedTime returns the CompletedTime field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultEntryAllOf) GetCompletedTime() string {
	if o == nil || o.CompletedTime == nil {
		var ret string
		return ret
	}
	return *o.CompletedTime
}

// GetCompletedTimeOk returns a tuple with the CompletedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetCompletedTimeOk() (*string, bool) {
	if o == nil || o.CompletedTime == nil {
		return nil, false
	}
	return o.CompletedTime, true
}

// HasCompletedTime returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasCompletedTime() bool {
	if o != nil && o.CompletedTime != nil {
		return true
	}

	return false
}

// SetCompletedTime gets a reference to the given string and assigns it to the CompletedTime field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetCompletedTime(v string) {
	o.CompletedTime = &v
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigResultEntryAllOf) GetContext() PolicyConfigResultContext {
	if o == nil || o.Context.Get() == nil {
		var ret PolicyConfigResultContext
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigResultEntryAllOf) GetContextOk() (*PolicyConfigResultContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasContext() bool {
	if o != nil && o.Context.IsSet() {
		return true
	}

	return false
}

// SetContext gets a reference to the given NullablePolicyConfigResultContext and assigns it to the Context field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetContext(v PolicyConfigResultContext) {
	o.Context.Set(&v)
}

// SetContextNil sets the value for Context to be an explicit nil
func (o *PolicyAbstractConfigResultEntryAllOf) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil
func (o *PolicyAbstractConfigResultEntryAllOf) UnsetContext() {
	o.Context.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultEntryAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultEntryAllOf) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultEntryAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultEntryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultEntryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyAbstractConfigResultEntryAllOf) SetType(v string) {
	o.Type = &v
}

func (o PolicyAbstractConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompletedTime != nil {
		toSerialize["CompletedTime"] = o.CompletedTime
	}
	if o.Context.IsSet() {
		toSerialize["Context"] = o.Context.Get()
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.OwnerId != nil {
		toSerialize["OwnerId"] = o.OwnerId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractConfigResultEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAbstractConfigResultEntryAllOf := _PolicyAbstractConfigResultEntryAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyAbstractConfigResultEntryAllOf); err == nil {
		*o = PolicyAbstractConfigResultEntryAllOf(varPolicyAbstractConfigResultEntryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompletedTime")
		delete(additionalProperties, "Context")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "OwnerId")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAbstractConfigResultEntryAllOf struct {
	value *PolicyAbstractConfigResultEntryAllOf
	isSet bool
}

func (v NullablePolicyAbstractConfigResultEntryAllOf) Get() *PolicyAbstractConfigResultEntryAllOf {
	return v.value
}

func (v *NullablePolicyAbstractConfigResultEntryAllOf) Set(val *PolicyAbstractConfigResultEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigResultEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigResultEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigResultEntryAllOf(val *PolicyAbstractConfigResultEntryAllOf) *NullablePolicyAbstractConfigResultEntryAllOf {
	return &NullablePolicyAbstractConfigResultEntryAllOf{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigResultEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
