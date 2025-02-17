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

// FabricFcoeUplinkRole Configuration object sent by user to create a fcoe uplink port.
type FabricFcoeUplinkRole struct {
	FabricTransceiverRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured state for UDLD for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	UdldAdminState       *string `json:"UdldAdminState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcoeUplinkRole FabricFcoeUplinkRole

// NewFabricFcoeUplinkRole instantiates a new FabricFcoeUplinkRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcoeUplinkRole(classId string, objectType string) *FabricFcoeUplinkRole {
	this := FabricFcoeUplinkRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// NewFabricFcoeUplinkRoleWithDefaults instantiates a new FabricFcoeUplinkRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcoeUplinkRoleWithDefaults() *FabricFcoeUplinkRole {
	this := FabricFcoeUplinkRole{}
	var classId string = "fabric.FcoeUplinkRole"
	this.ClassId = classId
	var objectType string = "fabric.FcoeUplinkRole"
	this.ObjectType = objectType
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcoeUplinkRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcoeUplinkRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcoeUplinkRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcoeUplinkRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUdldAdminState returns the UdldAdminState field value if set, zero value otherwise.
func (o *FabricFcoeUplinkRole) GetUdldAdminState() string {
	if o == nil || o.UdldAdminState == nil {
		var ret string
		return ret
	}
	return *o.UdldAdminState
}

// GetUdldAdminStateOk returns a tuple with the UdldAdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetUdldAdminStateOk() (*string, bool) {
	if o == nil || o.UdldAdminState == nil {
		return nil, false
	}
	return o.UdldAdminState, true
}

// HasUdldAdminState returns a boolean if a field has been set.
func (o *FabricFcoeUplinkRole) HasUdldAdminState() bool {
	if o != nil && o.UdldAdminState != nil {
		return true
	}

	return false
}

// SetUdldAdminState gets a reference to the given string and assigns it to the UdldAdminState field.
func (o *FabricFcoeUplinkRole) SetUdldAdminState(v string) {
	o.UdldAdminState = &v
}

func (o FabricFcoeUplinkRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricTransceiverRole, errFabricTransceiverRole := json.Marshal(o.FabricTransceiverRole)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	errFabricTransceiverRole = json.Unmarshal([]byte(serializedFabricTransceiverRole), &toSerialize)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UdldAdminState != nil {
		toSerialize["UdldAdminState"] = o.UdldAdminState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcoeUplinkRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcoeUplinkRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured state for UDLD for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		UdldAdminState *string `json:"UdldAdminState,omitempty"`
	}

	varFabricFcoeUplinkRoleWithoutEmbeddedStruct := FabricFcoeUplinkRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcoeUplinkRole := _FabricFcoeUplinkRole{}
		varFabricFcoeUplinkRole.ClassId = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.ClassId
		varFabricFcoeUplinkRole.ObjectType = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.ObjectType
		varFabricFcoeUplinkRole.UdldAdminState = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.UdldAdminState
		*o = FabricFcoeUplinkRole(varFabricFcoeUplinkRole)
	} else {
		return err
	}

	varFabricFcoeUplinkRole := _FabricFcoeUplinkRole{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkRole)
	if err == nil {
		o.FabricTransceiverRole = varFabricFcoeUplinkRole.FabricTransceiverRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UdldAdminState")

		// remove fields from embedded structs
		reflectFabricTransceiverRole := reflect.ValueOf(o.FabricTransceiverRole)
		for i := 0; i < reflectFabricTransceiverRole.Type().NumField(); i++ {
			t := reflectFabricTransceiverRole.Type().Field(i)

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

type NullableFabricFcoeUplinkRole struct {
	value *FabricFcoeUplinkRole
	isSet bool
}

func (v NullableFabricFcoeUplinkRole) Get() *FabricFcoeUplinkRole {
	return v.value
}

func (v *NullableFabricFcoeUplinkRole) Set(val *FabricFcoeUplinkRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcoeUplinkRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcoeUplinkRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcoeUplinkRole(val *FabricFcoeUplinkRole) *NullableFabricFcoeUplinkRole {
	return &NullableFabricFcoeUplinkRole{value: val, isSet: true}
}

func (v NullableFabricFcoeUplinkRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcoeUplinkRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
