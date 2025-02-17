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

// PoolAbstractBlockAllOf Definition of the list of properties defined in 'pool.AbstractBlock', excluding properties defined in parent classes.
type PoolAbstractBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Free IDs that can be allocated in this block.
	FreeBlockCount *int64 `json:"FreeBlockCount,omitempty"`
	// Moving counter to allocate the next identifier.
	NextIdAllocator      *int64 `json:"NextIdAllocator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractBlockAllOf PoolAbstractBlockAllOf

// NewPoolAbstractBlockAllOf instantiates a new PoolAbstractBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractBlockAllOf(classId string, objectType string) *PoolAbstractBlockAllOf {
	this := PoolAbstractBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPoolAbstractBlockAllOfWithDefaults instantiates a new PoolAbstractBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractBlockAllOfWithDefaults() *PoolAbstractBlockAllOf {
	this := PoolAbstractBlockAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFreeBlockCount returns the FreeBlockCount field value if set, zero value otherwise.
func (o *PoolAbstractBlockAllOf) GetFreeBlockCount() int64 {
	if o == nil || o.FreeBlockCount == nil {
		var ret int64
		return ret
	}
	return *o.FreeBlockCount
}

// GetFreeBlockCountOk returns a tuple with the FreeBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockAllOf) GetFreeBlockCountOk() (*int64, bool) {
	if o == nil || o.FreeBlockCount == nil {
		return nil, false
	}
	return o.FreeBlockCount, true
}

// HasFreeBlockCount returns a boolean if a field has been set.
func (o *PoolAbstractBlockAllOf) HasFreeBlockCount() bool {
	if o != nil && o.FreeBlockCount != nil {
		return true
	}

	return false
}

// SetFreeBlockCount gets a reference to the given int64 and assigns it to the FreeBlockCount field.
func (o *PoolAbstractBlockAllOf) SetFreeBlockCount(v int64) {
	o.FreeBlockCount = &v
}

// GetNextIdAllocator returns the NextIdAllocator field value if set, zero value otherwise.
func (o *PoolAbstractBlockAllOf) GetNextIdAllocator() int64 {
	if o == nil || o.NextIdAllocator == nil {
		var ret int64
		return ret
	}
	return *o.NextIdAllocator
}

// GetNextIdAllocatorOk returns a tuple with the NextIdAllocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockAllOf) GetNextIdAllocatorOk() (*int64, bool) {
	if o == nil || o.NextIdAllocator == nil {
		return nil, false
	}
	return o.NextIdAllocator, true
}

// HasNextIdAllocator returns a boolean if a field has been set.
func (o *PoolAbstractBlockAllOf) HasNextIdAllocator() bool {
	if o != nil && o.NextIdAllocator != nil {
		return true
	}

	return false
}

// SetNextIdAllocator gets a reference to the given int64 and assigns it to the NextIdAllocator field.
func (o *PoolAbstractBlockAllOf) SetNextIdAllocator(v int64) {
	o.NextIdAllocator = &v
}

func (o PoolAbstractBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FreeBlockCount != nil {
		toSerialize["FreeBlockCount"] = o.FreeBlockCount
	}
	if o.NextIdAllocator != nil {
		toSerialize["NextIdAllocator"] = o.NextIdAllocator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractBlockAllOf := _PoolAbstractBlockAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractBlockAllOf); err == nil {
		*o = PoolAbstractBlockAllOf(varPoolAbstractBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeBlockCount")
		delete(additionalProperties, "NextIdAllocator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractBlockAllOf struct {
	value *PoolAbstractBlockAllOf
	isSet bool
}

func (v NullablePoolAbstractBlockAllOf) Get() *PoolAbstractBlockAllOf {
	return v.value
}

func (v *NullablePoolAbstractBlockAllOf) Set(val *PoolAbstractBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractBlockAllOf(val *PoolAbstractBlockAllOf) *NullablePoolAbstractBlockAllOf {
	return &NullablePoolAbstractBlockAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
