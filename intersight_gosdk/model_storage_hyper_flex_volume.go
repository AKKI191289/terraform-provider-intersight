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

// StorageHyperFlexVolume A HyperFlex Volume entity.
type StorageHyperFlexVolume struct {
	StorageBaseVolume
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Provisioned Capacity of the Storage container in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Serial number of the volume.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Uuid of the Datastore/Storage Container.
	Uuid                 *string                                       `json:"Uuid,omitempty"`
	Cluster              *HyperflexClusterRelationship                 `json:"Cluster,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	StorageContainer     *StorageHyperFlexStorageContainerRelationship `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHyperFlexVolume StorageHyperFlexVolume

// NewStorageHyperFlexVolume instantiates a new StorageHyperFlexVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHyperFlexVolume(classId string, objectType string) *StorageHyperFlexVolume {
	this := StorageHyperFlexVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHyperFlexVolumeWithDefaults instantiates a new StorageHyperFlexVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHyperFlexVolumeWithDefaults() *StorageHyperFlexVolume {
	this := StorageHyperFlexVolume{}
	var classId string = "storage.HyperFlexVolume"
	this.ClassId = classId
	var objectType string = "storage.HyperFlexVolume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHyperFlexVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHyperFlexVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHyperFlexVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHyperFlexVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *StorageHyperFlexVolume) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *StorageHyperFlexVolume) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageHyperFlexVolume) SetUuid(v string) {
	o.Uuid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *StorageHyperFlexVolume) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHyperFlexVolume) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageHyperFlexVolume) GetStorageContainer() StorageHyperFlexStorageContainerRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageHyperFlexStorageContainerRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexVolume) GetStorageContainerOk() (*StorageHyperFlexStorageContainerRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageHyperFlexVolume) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageHyperFlexStorageContainerRelationship and assigns it to the StorageContainer field.
func (o *StorageHyperFlexVolume) SetStorageContainer(v StorageHyperFlexStorageContainerRelationship) {
	o.StorageContainer = &v
}

func (o StorageHyperFlexVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseVolume, errStorageBaseVolume := json.Marshal(o.StorageBaseVolume)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	errStorageBaseVolume = json.Unmarshal([]byte(serializedStorageBaseVolume), &toSerialize)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHyperFlexVolume) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHyperFlexVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Provisioned Capacity of the Storage container in bytes.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Serial number of the volume.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Uuid of the Datastore/Storage Container.
		Uuid             *string                                       `json:"Uuid,omitempty"`
		Cluster          *HyperflexClusterRelationship                 `json:"Cluster,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
		StorageContainer *StorageHyperFlexStorageContainerRelationship `json:"StorageContainer,omitempty"`
	}

	varStorageHyperFlexVolumeWithoutEmbeddedStruct := StorageHyperFlexVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHyperFlexVolumeWithoutEmbeddedStruct)
	if err == nil {
		varStorageHyperFlexVolume := _StorageHyperFlexVolume{}
		varStorageHyperFlexVolume.ClassId = varStorageHyperFlexVolumeWithoutEmbeddedStruct.ClassId
		varStorageHyperFlexVolume.ObjectType = varStorageHyperFlexVolumeWithoutEmbeddedStruct.ObjectType
		varStorageHyperFlexVolume.Capacity = varStorageHyperFlexVolumeWithoutEmbeddedStruct.Capacity
		varStorageHyperFlexVolume.SerialNumber = varStorageHyperFlexVolumeWithoutEmbeddedStruct.SerialNumber
		varStorageHyperFlexVolume.Uuid = varStorageHyperFlexVolumeWithoutEmbeddedStruct.Uuid
		varStorageHyperFlexVolume.Cluster = varStorageHyperFlexVolumeWithoutEmbeddedStruct.Cluster
		varStorageHyperFlexVolume.RegisteredDevice = varStorageHyperFlexVolumeWithoutEmbeddedStruct.RegisteredDevice
		varStorageHyperFlexVolume.StorageContainer = varStorageHyperFlexVolumeWithoutEmbeddedStruct.StorageContainer
		*o = StorageHyperFlexVolume(varStorageHyperFlexVolume)
	} else {
		return err
	}

	varStorageHyperFlexVolume := _StorageHyperFlexVolume{}

	err = json.Unmarshal(bytes, &varStorageHyperFlexVolume)
	if err == nil {
		o.StorageBaseVolume = varStorageHyperFlexVolume.StorageBaseVolume
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageContainer")

		// remove fields from embedded structs
		reflectStorageBaseVolume := reflect.ValueOf(o.StorageBaseVolume)
		for i := 0; i < reflectStorageBaseVolume.Type().NumField(); i++ {
			t := reflectStorageBaseVolume.Type().Field(i)

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

type NullableStorageHyperFlexVolume struct {
	value *StorageHyperFlexVolume
	isSet bool
}

func (v NullableStorageHyperFlexVolume) Get() *StorageHyperFlexVolume {
	return v.value
}

func (v *NullableStorageHyperFlexVolume) Set(val *StorageHyperFlexVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHyperFlexVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHyperFlexVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHyperFlexVolume(val *StorageHyperFlexVolume) *NullableStorageHyperFlexVolume {
	return &NullableStorageHyperFlexVolume{value: val, isSet: true}
}

func (v NullableStorageHyperFlexVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHyperFlexVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
