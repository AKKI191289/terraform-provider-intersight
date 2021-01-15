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
)

// StorageBaseReplicationBlackoutAllOf Definition of the list of properties defined in 'storage.BaseReplicationBlackout', excluding properties defined in parent classes.
type StorageBaseReplicationBlackoutAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The end time of day for replication blackout window. Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
	End *string `json:"End,omitempty"`
	// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication. Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds. The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
	Start                *string `json:"Start,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseReplicationBlackoutAllOf StorageBaseReplicationBlackoutAllOf

// NewStorageBaseReplicationBlackoutAllOf instantiates a new StorageBaseReplicationBlackoutAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseReplicationBlackoutAllOf(classId string, objectType string) *StorageBaseReplicationBlackoutAllOf {
	this := StorageBaseReplicationBlackoutAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseReplicationBlackoutAllOfWithDefaults instantiates a new StorageBaseReplicationBlackoutAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseReplicationBlackoutAllOfWithDefaults() *StorageBaseReplicationBlackoutAllOf {
	this := StorageBaseReplicationBlackoutAllOf{}
	var classId string = "storage.PureReplicationBlackout"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationBlackout"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseReplicationBlackoutAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackoutAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseReplicationBlackoutAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseReplicationBlackoutAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackoutAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseReplicationBlackoutAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *StorageBaseReplicationBlackoutAllOf) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackoutAllOf) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *StorageBaseReplicationBlackoutAllOf) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *StorageBaseReplicationBlackoutAllOf) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *StorageBaseReplicationBlackoutAllOf) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackoutAllOf) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *StorageBaseReplicationBlackoutAllOf) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *StorageBaseReplicationBlackoutAllOf) SetStart(v string) {
	o.Start = &v
}

func (o StorageBaseReplicationBlackoutAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.End != nil {
		toSerialize["End"] = o.End
	}
	if o.Start != nil {
		toSerialize["Start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseReplicationBlackoutAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseReplicationBlackoutAllOf := _StorageBaseReplicationBlackoutAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseReplicationBlackoutAllOf); err == nil {
		*o = StorageBaseReplicationBlackoutAllOf(varStorageBaseReplicationBlackoutAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "End")
		delete(additionalProperties, "Start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseReplicationBlackoutAllOf struct {
	value *StorageBaseReplicationBlackoutAllOf
	isSet bool
}

func (v NullableStorageBaseReplicationBlackoutAllOf) Get() *StorageBaseReplicationBlackoutAllOf {
	return v.value
}

func (v *NullableStorageBaseReplicationBlackoutAllOf) Set(val *StorageBaseReplicationBlackoutAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseReplicationBlackoutAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseReplicationBlackoutAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseReplicationBlackoutAllOf(val *StorageBaseReplicationBlackoutAllOf) *NullableStorageBaseReplicationBlackoutAllOf {
	return &NullableStorageBaseReplicationBlackoutAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseReplicationBlackoutAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseReplicationBlackoutAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
