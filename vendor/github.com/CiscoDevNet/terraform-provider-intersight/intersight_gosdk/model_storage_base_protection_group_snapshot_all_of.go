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
	"time"
)

// StorageBaseProtectionGroupSnapshotAllOf Definition of the list of properties defined in 'storage.BaseProtectionGroupSnapshot', excluding properties defined in parent classes.
type StorageBaseProtectionGroupSnapshotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Protection group snapshot creation time.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Protection group snapshot name which represents point-in-time copy of all members in protection group.
	Name *string `json:"Name,omitempty"`
	// Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set.
	Size *int64 `json:"Size,omitempty"`
	// Source protection group name on which the snapshot is created.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseProtectionGroupSnapshotAllOf StorageBaseProtectionGroupSnapshotAllOf

// NewStorageBaseProtectionGroupSnapshotAllOf instantiates a new StorageBaseProtectionGroupSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseProtectionGroupSnapshotAllOf(classId string, objectType string) *StorageBaseProtectionGroupSnapshotAllOf {
	this := StorageBaseProtectionGroupSnapshotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseProtectionGroupSnapshotAllOfWithDefaults instantiates a new StorageBaseProtectionGroupSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseProtectionGroupSnapshotAllOfWithDefaults() *StorageBaseProtectionGroupSnapshotAllOf {
	this := StorageBaseProtectionGroupSnapshotAllOf{}
	var classId string = "storage.PureProtectionGroupSnapshot"
	this.ClassId = classId
	var objectType string = "storage.PureProtectionGroupSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroupSnapshotAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageBaseProtectionGroupSnapshotAllOf) SetSource(v string) {
	o.Source = &v
}

func (o StorageBaseProtectionGroupSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CreatedTime != nil {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseProtectionGroupSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseProtectionGroupSnapshotAllOf := _StorageBaseProtectionGroupSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseProtectionGroupSnapshotAllOf); err == nil {
		*o = StorageBaseProtectionGroupSnapshotAllOf(varStorageBaseProtectionGroupSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseProtectionGroupSnapshotAllOf struct {
	value *StorageBaseProtectionGroupSnapshotAllOf
	isSet bool
}

func (v NullableStorageBaseProtectionGroupSnapshotAllOf) Get() *StorageBaseProtectionGroupSnapshotAllOf {
	return v.value
}

func (v *NullableStorageBaseProtectionGroupSnapshotAllOf) Set(val *StorageBaseProtectionGroupSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseProtectionGroupSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseProtectionGroupSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseProtectionGroupSnapshotAllOf(val *StorageBaseProtectionGroupSnapshotAllOf) *NullableStorageBaseProtectionGroupSnapshotAllOf {
	return &NullableStorageBaseProtectionGroupSnapshotAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseProtectionGroupSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseProtectionGroupSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
