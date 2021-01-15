/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// RecoveryConfigResultEntry An entry that describes the result of a Backup Profile state on the end device.
type RecoveryConfigResultEntry struct {
	PolicyAbstractConfigResultEntry
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                            `json:"ObjectType"`
	ConfigResult         *RecoveryConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryConfigResultEntry RecoveryConfigResultEntry

// NewRecoveryConfigResultEntry instantiates a new RecoveryConfigResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryConfigResultEntry(classId string, objectType string) *RecoveryConfigResultEntry {
	this := RecoveryConfigResultEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryConfigResultEntryWithDefaults instantiates a new RecoveryConfigResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryConfigResultEntryWithDefaults() *RecoveryConfigResultEntry {
	this := RecoveryConfigResultEntry{}
	var classId string = "recovery.ConfigResultEntry"
	this.ClassId = classId
	var objectType string = "recovery.ConfigResultEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryConfigResultEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryConfigResultEntry) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryConfigResultEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryConfigResultEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryConfigResultEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryConfigResultEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *RecoveryConfigResultEntry) GetConfigResult() RecoveryConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret RecoveryConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryConfigResultEntry) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *RecoveryConfigResultEntry) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given RecoveryConfigResultRelationship and assigns it to the ConfigResult field.
func (o *RecoveryConfigResultEntry) SetConfigResult(v RecoveryConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o RecoveryConfigResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResultEntry, errPolicyAbstractConfigResultEntry := json.Marshal(o.PolicyAbstractConfigResultEntry)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	errPolicyAbstractConfigResultEntry = json.Unmarshal([]byte(serializedPolicyAbstractConfigResultEntry), &toSerialize)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryConfigResultEntry) UnmarshalJSON(bytes []byte) (err error) {
	type RecoveryConfigResultEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                            `json:"ObjectType"`
		ConfigResult *RecoveryConfigResultRelationship `json:"ConfigResult,omitempty"`
	}

	varRecoveryConfigResultEntryWithoutEmbeddedStruct := RecoveryConfigResultEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecoveryConfigResultEntryWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryConfigResultEntry := _RecoveryConfigResultEntry{}
		varRecoveryConfigResultEntry.ClassId = varRecoveryConfigResultEntryWithoutEmbeddedStruct.ClassId
		varRecoveryConfigResultEntry.ObjectType = varRecoveryConfigResultEntryWithoutEmbeddedStruct.ObjectType
		varRecoveryConfigResultEntry.ConfigResult = varRecoveryConfigResultEntryWithoutEmbeddedStruct.ConfigResult
		*o = RecoveryConfigResultEntry(varRecoveryConfigResultEntry)
	} else {
		return err
	}

	varRecoveryConfigResultEntry := _RecoveryConfigResultEntry{}

	err = json.Unmarshal(bytes, &varRecoveryConfigResultEntry)
	if err == nil {
		o.PolicyAbstractConfigResultEntry = varRecoveryConfigResultEntry.PolicyAbstractConfigResultEntry
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigResult")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResultEntry := reflect.ValueOf(o.PolicyAbstractConfigResultEntry)
		for i := 0; i < reflectPolicyAbstractConfigResultEntry.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResultEntry.Type().Field(i)

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

type NullableRecoveryConfigResultEntry struct {
	value *RecoveryConfigResultEntry
	isSet bool
}

func (v NullableRecoveryConfigResultEntry) Get() *RecoveryConfigResultEntry {
	return v.value
}

func (v *NullableRecoveryConfigResultEntry) Set(val *RecoveryConfigResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryConfigResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryConfigResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryConfigResultEntry(val *RecoveryConfigResultEntry) *NullableRecoveryConfigResultEntry {
	return &NullableRecoveryConfigResultEntry{value: val, isSet: true}
}

func (v NullableRecoveryConfigResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryConfigResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
