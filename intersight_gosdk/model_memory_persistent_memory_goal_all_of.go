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

// MemoryPersistentMemoryGoalAllOf Definition of the list of properties defined in 'memory.PersistentMemoryGoal', excluding properties defined in parent classes.
type MemoryPersistentMemoryGoalAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Volatile memory percentage.
	MemoryModePercentage *int64 `json:"MemoryModePercentage,omitempty"`
	// Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not. * `app-direct` - The App Direct interleaved Persistent Memory type. * `app-direct-non-interleaved` - The App Direct non-interleaved Persistent Memory type.
	PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
	// CPU Socket ID to which this goal will be applied. * `All Sockets` - All the CPU socket IDs in a server.
	SocketId             *string `json:"SocketId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPersistentMemoryGoalAllOf MemoryPersistentMemoryGoalAllOf

// NewMemoryPersistentMemoryGoalAllOf instantiates a new MemoryPersistentMemoryGoalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryGoalAllOf(classId string, objectType string) *MemoryPersistentMemoryGoalAllOf {
	this := MemoryPersistentMemoryGoalAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// NewMemoryPersistentMemoryGoalAllOfWithDefaults instantiates a new MemoryPersistentMemoryGoalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryGoalAllOfWithDefaults() *MemoryPersistentMemoryGoalAllOf {
	this := MemoryPersistentMemoryGoalAllOf{}
	var classId string = "memory.PersistentMemoryGoal"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryGoal"
	this.ObjectType = objectType
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryGoalAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoalAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryGoalAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryGoalAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoalAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryGoalAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemoryModePercentage returns the MemoryModePercentage field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoalAllOf) GetMemoryModePercentage() int64 {
	if o == nil || o.MemoryModePercentage == nil {
		var ret int64
		return ret
	}
	return *o.MemoryModePercentage
}

// GetMemoryModePercentageOk returns a tuple with the MemoryModePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoalAllOf) GetMemoryModePercentageOk() (*int64, bool) {
	if o == nil || o.MemoryModePercentage == nil {
		return nil, false
	}
	return o.MemoryModePercentage, true
}

// HasMemoryModePercentage returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoalAllOf) HasMemoryModePercentage() bool {
	if o != nil && o.MemoryModePercentage != nil {
		return true
	}

	return false
}

// SetMemoryModePercentage gets a reference to the given int64 and assigns it to the MemoryModePercentage field.
func (o *MemoryPersistentMemoryGoalAllOf) SetMemoryModePercentage(v int64) {
	o.MemoryModePercentage = &v
}

// GetPersistentMemoryType returns the PersistentMemoryType field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoalAllOf) GetPersistentMemoryType() string {
	if o == nil || o.PersistentMemoryType == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryType
}

// GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoalAllOf) GetPersistentMemoryTypeOk() (*string, bool) {
	if o == nil || o.PersistentMemoryType == nil {
		return nil, false
	}
	return o.PersistentMemoryType, true
}

// HasPersistentMemoryType returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoalAllOf) HasPersistentMemoryType() bool {
	if o != nil && o.PersistentMemoryType != nil {
		return true
	}

	return false
}

// SetPersistentMemoryType gets a reference to the given string and assigns it to the PersistentMemoryType field.
func (o *MemoryPersistentMemoryGoalAllOf) SetPersistentMemoryType(v string) {
	o.PersistentMemoryType = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoalAllOf) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoalAllOf) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoalAllOf) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryGoalAllOf) SetSocketId(v string) {
	o.SocketId = &v
}

func (o MemoryPersistentMemoryGoalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MemoryModePercentage != nil {
		toSerialize["MemoryModePercentage"] = o.MemoryModePercentage
	}
	if o.PersistentMemoryType != nil {
		toSerialize["PersistentMemoryType"] = o.PersistentMemoryType
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryGoalAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryPersistentMemoryGoalAllOf := _MemoryPersistentMemoryGoalAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryPersistentMemoryGoalAllOf); err == nil {
		*o = MemoryPersistentMemoryGoalAllOf(varMemoryPersistentMemoryGoalAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemoryModePercentage")
		delete(additionalProperties, "PersistentMemoryType")
		delete(additionalProperties, "SocketId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryPersistentMemoryGoalAllOf struct {
	value *MemoryPersistentMemoryGoalAllOf
	isSet bool
}

func (v NullableMemoryPersistentMemoryGoalAllOf) Get() *MemoryPersistentMemoryGoalAllOf {
	return v.value
}

func (v *NullableMemoryPersistentMemoryGoalAllOf) Set(val *MemoryPersistentMemoryGoalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryGoalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryGoalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryGoalAllOf(val *MemoryPersistentMemoryGoalAllOf) *NullableMemoryPersistentMemoryGoalAllOf {
	return &NullableMemoryPersistentMemoryGoalAllOf{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryGoalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryGoalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
