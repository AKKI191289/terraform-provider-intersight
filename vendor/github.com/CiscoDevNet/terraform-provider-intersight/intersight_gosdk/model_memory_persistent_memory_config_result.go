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

// MemoryPersistentMemoryConfigResult Result of a previously applied Persistent Memory configuration on a server.
type MemoryPersistentMemoryConfigResult struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error in the result of a previously applied Persistent Memory configuration on a server.
	ConfigErrorDesc *string `json:"ConfigErrorDesc,omitempty"`
	// Result of a previously applied Persistent Memory configuration on a server.
	ConfigResult *string `json:"ConfigResult,omitempty"`
	// Sequence number of a previously applied Persistent Memory configuration on a server.
	ConfigSequenceNo *int64 `json:"ConfigSequenceNo,omitempty"`
	// State of a previously applied Persistent Memory configuration on a server.
	ConfigState                         *string                                          `json:"ConfigState,omitempty"`
	InventoryDeviceInfo                 *InventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
	// An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources.
	PersistentMemoryNamespaceConfigResults []MemoryPersistentMemoryNamespaceConfigResultRelationship `json:"PersistentMemoryNamespaceConfigResults,omitempty"`
	RegisteredDevice                       *AssetDeviceRegistrationRelationship                      `json:"RegisteredDevice,omitempty"`
	AdditionalProperties                   map[string]interface{}
}

type _MemoryPersistentMemoryConfigResult MemoryPersistentMemoryConfigResult

// NewMemoryPersistentMemoryConfigResult instantiates a new MemoryPersistentMemoryConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryConfigResult(classId string, objectType string) *MemoryPersistentMemoryConfigResult {
	this := MemoryPersistentMemoryConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryConfigResultWithDefaults instantiates a new MemoryPersistentMemoryConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryConfigResultWithDefaults() *MemoryPersistentMemoryConfigResult {
	this := MemoryPersistentMemoryConfigResult{}
	var classId string = "memory.PersistentMemoryConfigResult"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigErrorDesc returns the ConfigErrorDesc field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDesc() string {
	if o == nil || o.ConfigErrorDesc == nil {
		var ret string
		return ret
	}
	return *o.ConfigErrorDesc
}

// GetConfigErrorDescOk returns a tuple with the ConfigErrorDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDescOk() (*string, bool) {
	if o == nil || o.ConfigErrorDesc == nil {
		return nil, false
	}
	return o.ConfigErrorDesc, true
}

// HasConfigErrorDesc returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigErrorDesc() bool {
	if o != nil && o.ConfigErrorDesc != nil {
		return true
	}

	return false
}

// SetConfigErrorDesc gets a reference to the given string and assigns it to the ConfigErrorDesc field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigErrorDesc(v string) {
	o.ConfigErrorDesc = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigResult() string {
	if o == nil || o.ConfigResult == nil {
		var ret string
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigResultOk() (*string, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given string and assigns it to the ConfigResult field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigResult(v string) {
	o.ConfigResult = &v
}

// GetConfigSequenceNo returns the ConfigSequenceNo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNo() int64 {
	if o == nil || o.ConfigSequenceNo == nil {
		var ret int64
		return ret
	}
	return *o.ConfigSequenceNo
}

// GetConfigSequenceNoOk returns a tuple with the ConfigSequenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNoOk() (*int64, bool) {
	if o == nil || o.ConfigSequenceNo == nil {
		return nil, false
	}
	return o.ConfigSequenceNo, true
}

// HasConfigSequenceNo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigSequenceNo() bool {
	if o != nil && o.ConfigSequenceNo != nil {
		return true
	}

	return false
}

// SetConfigSequenceNo gets a reference to the given int64 and assigns it to the ConfigSequenceNo field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigSequenceNo(v int64) {
	o.ConfigSequenceNo = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		var ret MemoryPersistentMemoryConfigurationRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfiguration
}

// GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool) {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfiguration, true
}

// HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasMemoryPersistentMemoryConfiguration() bool {
	if o != nil && o.MemoryPersistentMemoryConfiguration != nil {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfiguration gets a reference to the given MemoryPersistentMemoryConfigurationRelationship and assigns it to the MemoryPersistentMemoryConfiguration field.
func (o *MemoryPersistentMemoryConfigResult) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship) {
	o.MemoryPersistentMemoryConfiguration = &v
}

// GetPersistentMemoryNamespaceConfigResults returns the PersistentMemoryNamespaceConfigResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResults() []MemoryPersistentMemoryNamespaceConfigResultRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryNamespaceConfigResultRelationship
		return ret
	}
	return o.PersistentMemoryNamespaceConfigResults
}

// GetPersistentMemoryNamespaceConfigResultsOk returns a tuple with the PersistentMemoryNamespaceConfigResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResultsOk() (*[]MemoryPersistentMemoryNamespaceConfigResultRelationship, bool) {
	if o == nil || o.PersistentMemoryNamespaceConfigResults == nil {
		return nil, false
	}
	return &o.PersistentMemoryNamespaceConfigResults, true
}

// HasPersistentMemoryNamespaceConfigResults returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasPersistentMemoryNamespaceConfigResults() bool {
	if o != nil && o.PersistentMemoryNamespaceConfigResults != nil {
		return true
	}

	return false
}

// SetPersistentMemoryNamespaceConfigResults gets a reference to the given []MemoryPersistentMemoryNamespaceConfigResultRelationship and assigns it to the PersistentMemoryNamespaceConfigResults field.
func (o *MemoryPersistentMemoryConfigResult) SetPersistentMemoryNamespaceConfigResults(v []MemoryPersistentMemoryNamespaceConfigResultRelationship) {
	o.PersistentMemoryNamespaceConfigResults = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigErrorDesc != nil {
		toSerialize["ConfigErrorDesc"] = o.ConfigErrorDesc
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.ConfigSequenceNo != nil {
		toSerialize["ConfigSequenceNo"] = o.ConfigSequenceNo
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryPersistentMemoryConfiguration != nil {
		toSerialize["MemoryPersistentMemoryConfiguration"] = o.MemoryPersistentMemoryConfiguration
	}
	if o.PersistentMemoryNamespaceConfigResults != nil {
		toSerialize["PersistentMemoryNamespaceConfigResults"] = o.PersistentMemoryNamespaceConfigResults
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Error in the result of a previously applied Persistent Memory configuration on a server.
		ConfigErrorDesc *string `json:"ConfigErrorDesc,omitempty"`
		// Result of a previously applied Persistent Memory configuration on a server.
		ConfigResult *string `json:"ConfigResult,omitempty"`
		// Sequence number of a previously applied Persistent Memory configuration on a server.
		ConfigSequenceNo *int64 `json:"ConfigSequenceNo,omitempty"`
		// State of a previously applied Persistent Memory configuration on a server.
		ConfigState                         *string                                          `json:"ConfigState,omitempty"`
		InventoryDeviceInfo                 *InventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
		MemoryPersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
		// An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources.
		PersistentMemoryNamespaceConfigResults []MemoryPersistentMemoryNamespaceConfigResultRelationship `json:"PersistentMemoryNamespaceConfigResults,omitempty"`
		RegisteredDevice                       *AssetDeviceRegistrationRelationship                      `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct := MemoryPersistentMemoryConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryConfigResult := _MemoryPersistentMemoryConfigResult{}
		varMemoryPersistentMemoryConfigResult.ClassId = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryConfigResult.ObjectType = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryConfigResult.ConfigErrorDesc = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigErrorDesc
		varMemoryPersistentMemoryConfigResult.ConfigResult = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigResult
		varMemoryPersistentMemoryConfigResult.ConfigSequenceNo = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigSequenceNo
		varMemoryPersistentMemoryConfigResult.ConfigState = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigState
		varMemoryPersistentMemoryConfigResult.InventoryDeviceInfo = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryConfigResult.MemoryPersistentMemoryConfiguration = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.MemoryPersistentMemoryConfiguration
		varMemoryPersistentMemoryConfigResult.PersistentMemoryNamespaceConfigResults = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.PersistentMemoryNamespaceConfigResults
		varMemoryPersistentMemoryConfigResult.RegisteredDevice = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryConfigResult(varMemoryPersistentMemoryConfigResult)
	} else {
		return err
	}

	varMemoryPersistentMemoryConfigResult := _MemoryPersistentMemoryConfigResult{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryConfigResult)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryConfigResult.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigErrorDesc")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "ConfigSequenceNo")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfiguration")
		delete(additionalProperties, "PersistentMemoryNamespaceConfigResults")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableMemoryPersistentMemoryConfigResult struct {
	value *MemoryPersistentMemoryConfigResult
	isSet bool
}

func (v NullableMemoryPersistentMemoryConfigResult) Get() *MemoryPersistentMemoryConfigResult {
	return v.value
}

func (v *NullableMemoryPersistentMemoryConfigResult) Set(val *MemoryPersistentMemoryConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryConfigResult(val *MemoryPersistentMemoryConfigResult) *NullableMemoryPersistentMemoryConfigResult {
	return &NullableMemoryPersistentMemoryConfigResult{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
