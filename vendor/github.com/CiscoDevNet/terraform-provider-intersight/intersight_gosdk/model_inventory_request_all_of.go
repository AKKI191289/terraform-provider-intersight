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

// InventoryRequestAllOf Definition of the list of properties defined in 'inventory.Request', excluding properties defined in parent classes.
type InventoryRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Mos                  []InventoryInventoryMo               `json:"Mos,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryRequestAllOf InventoryRequestAllOf

// NewInventoryRequestAllOf instantiates a new InventoryRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryRequestAllOf(classId string, objectType string) *InventoryRequestAllOf {
	this := InventoryRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryRequestAllOfWithDefaults instantiates a new InventoryRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryRequestAllOfWithDefaults() *InventoryRequestAllOf {
	this := InventoryRequestAllOf{}
	var classId string = "inventory.Request"
	this.ClassId = classId
	var objectType string = "inventory.Request"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMos returns the Mos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryRequestAllOf) GetMos() []InventoryInventoryMo {
	if o == nil {
		var ret []InventoryInventoryMo
		return ret
	}
	return o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryRequestAllOf) GetMosOk() (*[]InventoryInventoryMo, bool) {
	if o == nil || o.Mos == nil {
		return nil, false
	}
	return &o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *InventoryRequestAllOf) HasMos() bool {
	if o != nil && o.Mos != nil {
		return true
	}

	return false
}

// SetMos gets a reference to the given []InventoryInventoryMo and assigns it to the Mos field.
func (o *InventoryRequestAllOf) SetMos(v []InventoryInventoryMo) {
	o.Mos = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InventoryRequestAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequestAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InventoryRequestAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *InventoryRequestAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o InventoryRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Mos != nil {
		toSerialize["Mos"] = o.Mos
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryRequestAllOf := _InventoryRequestAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryRequestAllOf); err == nil {
		*o = InventoryRequestAllOf(varInventoryRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mos")
		delete(additionalProperties, "Device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryRequestAllOf struct {
	value *InventoryRequestAllOf
	isSet bool
}

func (v NullableInventoryRequestAllOf) Get() *InventoryRequestAllOf {
	return v.value
}

func (v *NullableInventoryRequestAllOf) Set(val *InventoryRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryRequestAllOf(val *InventoryRequestAllOf) *NullableInventoryRequestAllOf {
	return &NullableInventoryRequestAllOf{value: val, isSet: true}
}

func (v NullableInventoryRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
