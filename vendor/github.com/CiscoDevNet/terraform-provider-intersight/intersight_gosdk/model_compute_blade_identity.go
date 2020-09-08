/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ComputeBladeIdentity Identity object that uniquely represents a blade server object under a DR.
type ComputeBladeIdentity struct {
	EquipmentPhysicalIdentity
	// Chassis Identifier of a blade server.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// FI Device registration Mo ID.
	DeviceMoId *string `json:"DeviceMoId,omitempty"`
	// Indicates pending discovery flag.
	PendingDiscovery *string `json:"PendingDiscovery,omitempty"`
	// Chassis slot number of a blade server.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBladeIdentity ComputeBladeIdentity

// NewComputeBladeIdentity instantiates a new ComputeBladeIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBladeIdentity() *ComputeBladeIdentity {
	this := ComputeBladeIdentity{}
	return &this
}

// NewComputeBladeIdentityWithDefaults instantiates a new ComputeBladeIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBladeIdentityWithDefaults() *ComputeBladeIdentity {
	this := ComputeBladeIdentity{}
	return &this
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *ComputeBladeIdentity) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *ComputeBladeIdentity) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

// GetPendingDiscovery returns the PendingDiscovery field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetPendingDiscovery() string {
	if o == nil || o.PendingDiscovery == nil {
		var ret string
		return ret
	}
	return *o.PendingDiscovery
}

// GetPendingDiscoveryOk returns a tuple with the PendingDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetPendingDiscoveryOk() (*string, bool) {
	if o == nil || o.PendingDiscovery == nil {
		return nil, false
	}
	return o.PendingDiscovery, true
}

// HasPendingDiscovery returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasPendingDiscovery() bool {
	if o != nil && o.PendingDiscovery != nil {
		return true
	}

	return false
}

// SetPendingDiscovery gets a reference to the given string and assigns it to the PendingDiscovery field.
func (o *ComputeBladeIdentity) SetPendingDiscovery(v string) {
	o.PendingDiscovery = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ComputeBladeIdentity) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o ComputeBladeIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}
	if o.PendingDiscovery != nil {
		toSerialize["PendingDiscovery"] = o.PendingDiscovery
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBladeIdentity) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeBladeIdentityWithoutEmbeddedStruct struct {
		// Chassis Identifier of a blade server.
		ChassisId *int64 `json:"ChassisId,omitempty"`
		// FI Device registration Mo ID.
		DeviceMoId *string `json:"DeviceMoId,omitempty"`
		// Indicates pending discovery flag.
		PendingDiscovery *string `json:"PendingDiscovery,omitempty"`
		// Chassis slot number of a blade server.
		SlotId *int64 `json:"SlotId,omitempty"`
	}

	varComputeBladeIdentityWithoutEmbeddedStruct := ComputeBladeIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeBladeIdentityWithoutEmbeddedStruct)
	if err == nil {
		varComputeBladeIdentity := _ComputeBladeIdentity{}
		varComputeBladeIdentity.ChassisId = varComputeBladeIdentityWithoutEmbeddedStruct.ChassisId
		varComputeBladeIdentity.DeviceMoId = varComputeBladeIdentityWithoutEmbeddedStruct.DeviceMoId
		varComputeBladeIdentity.PendingDiscovery = varComputeBladeIdentityWithoutEmbeddedStruct.PendingDiscovery
		varComputeBladeIdentity.SlotId = varComputeBladeIdentityWithoutEmbeddedStruct.SlotId
		*o = ComputeBladeIdentity(varComputeBladeIdentity)
	} else {
		return err
	}

	varComputeBladeIdentity := _ComputeBladeIdentity{}

	err = json.Unmarshal(bytes, &varComputeBladeIdentity)
	if err == nil {
		o.EquipmentPhysicalIdentity = varComputeBladeIdentity.EquipmentPhysicalIdentity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "DeviceMoId")
		delete(additionalProperties, "PendingDiscovery")
		delete(additionalProperties, "SlotId")

		// remove fields from embedded structs
		reflectEquipmentPhysicalIdentity := reflect.ValueOf(o.EquipmentPhysicalIdentity)
		for i := 0; i < reflectEquipmentPhysicalIdentity.Type().NumField(); i++ {
			t := reflectEquipmentPhysicalIdentity.Type().Field(i)

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

type NullableComputeBladeIdentity struct {
	value *ComputeBladeIdentity
	isSet bool
}

func (v NullableComputeBladeIdentity) Get() *ComputeBladeIdentity {
	return v.value
}

func (v *NullableComputeBladeIdentity) Set(val *ComputeBladeIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeIdentity(val *ComputeBladeIdentity) *NullableComputeBladeIdentity {
	return &NullableComputeBladeIdentity{value: val, isSet: true}
}

func (v NullableComputeBladeIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
