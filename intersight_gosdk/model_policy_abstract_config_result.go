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
	"reflect"
	"strings"
)

// PolicyAbstractConfigResult The results with the overall state and detailed result messages.
type PolicyAbstractConfigResult struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The current running stage of the configuration or workflow.
	ConfigStage *string `json:"ConfigStage,omitempty"`
	// Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored.
	ValidationState      *string `json:"ValidationState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigResult PolicyAbstractConfigResult

// NewPolicyAbstractConfigResult instantiates a new PolicyAbstractConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigResult(classId string, objectType string) *PolicyAbstractConfigResult {
	this := PolicyAbstractConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyAbstractConfigResultWithDefaults instantiates a new PolicyAbstractConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigResultWithDefaults() *PolicyAbstractConfigResult {
	this := PolicyAbstractConfigResult{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigStage returns the ConfigStage field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResult) GetConfigStage() string {
	if o == nil || o.ConfigStage == nil {
		var ret string
		return ret
	}
	return *o.ConfigStage
}

// GetConfigStageOk returns a tuple with the ConfigStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResult) GetConfigStageOk() (*string, bool) {
	if o == nil || o.ConfigStage == nil {
		return nil, false
	}
	return o.ConfigStage, true
}

// HasConfigStage returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResult) HasConfigStage() bool {
	if o != nil && o.ConfigStage != nil {
		return true
	}

	return false
}

// SetConfigStage gets a reference to the given string and assigns it to the ConfigStage field.
func (o *PolicyAbstractConfigResult) SetConfigStage(v string) {
	o.ConfigStage = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResult) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResult) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResult) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *PolicyAbstractConfigResult) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetValidationState returns the ValidationState field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResult) GetValidationState() string {
	if o == nil || o.ValidationState == nil {
		var ret string
		return ret
	}
	return *o.ValidationState
}

// GetValidationStateOk returns a tuple with the ValidationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResult) GetValidationStateOk() (*string, bool) {
	if o == nil || o.ValidationState == nil {
		return nil, false
	}
	return o.ValidationState, true
}

// HasValidationState returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResult) HasValidationState() bool {
	if o != nil && o.ValidationState != nil {
		return true
	}

	return false
}

// SetValidationState gets a reference to the given string and assigns it to the ValidationState field.
func (o *PolicyAbstractConfigResult) SetValidationState(v string) {
	o.ValidationState = &v
}

func (o PolicyAbstractConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigStage != nil {
		toSerialize["ConfigStage"] = o.ConfigStage
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.ValidationState != nil {
		toSerialize["ValidationState"] = o.ValidationState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyAbstractConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The current running stage of the configuration or workflow.
		ConfigStage *string `json:"ConfigStage,omitempty"`
		// Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored.
		ConfigState *string `json:"ConfigState,omitempty"`
		// Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored.
		ValidationState *string `json:"ValidationState,omitempty"`
	}

	varPolicyAbstractConfigResultWithoutEmbeddedStruct := PolicyAbstractConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractConfigResult := _PolicyAbstractConfigResult{}
		varPolicyAbstractConfigResult.ClassId = varPolicyAbstractConfigResultWithoutEmbeddedStruct.ClassId
		varPolicyAbstractConfigResult.ObjectType = varPolicyAbstractConfigResultWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractConfigResult.ConfigStage = varPolicyAbstractConfigResultWithoutEmbeddedStruct.ConfigStage
		varPolicyAbstractConfigResult.ConfigState = varPolicyAbstractConfigResultWithoutEmbeddedStruct.ConfigState
		varPolicyAbstractConfigResult.ValidationState = varPolicyAbstractConfigResultWithoutEmbeddedStruct.ValidationState
		*o = PolicyAbstractConfigResult(varPolicyAbstractConfigResult)
	} else {
		return err
	}

	varPolicyAbstractConfigResult := _PolicyAbstractConfigResult{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigResult)
	if err == nil {
		o.MoBaseMo = varPolicyAbstractConfigResult.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigStage")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "ValidationState")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAbstractConfigResult struct {
	value *PolicyAbstractConfigResult
	isSet bool
}

func (v NullablePolicyAbstractConfigResult) Get() *PolicyAbstractConfigResult {
	return v.value
}

func (v *NullablePolicyAbstractConfigResult) Set(val *PolicyAbstractConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigResult(val *PolicyAbstractConfigResult) *NullablePolicyAbstractConfigResult {
	return &NullablePolicyAbstractConfigResult{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
