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
)

// BootLocalDiskAllOf Definition of the list of properties defined in 'boot.LocalDisk', excluding properties defined in parent classes.
type BootLocalDiskAllOf struct {
	Bootloader *BootBootloader `json:"Bootloader,omitempty"`
	// The slot id of the local disk device. Supported values are (1-255, \"M\", \"HBA\", \"SAS\", \"RAID\", \"MRAID\", \"MSTOR-RAID\").
	Slot                 *string `json:"Slot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootLocalDiskAllOf BootLocalDiskAllOf

// NewBootLocalDiskAllOf instantiates a new BootLocalDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootLocalDiskAllOf() *BootLocalDiskAllOf {
	this := BootLocalDiskAllOf{}
	return &this
}

// NewBootLocalDiskAllOfWithDefaults instantiates a new BootLocalDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootLocalDiskAllOfWithDefaults() *BootLocalDiskAllOf {
	this := BootLocalDiskAllOf{}
	return &this
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise.
func (o *BootLocalDiskAllOf) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootLocalDiskAllOf) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil || o.Bootloader == nil {
		return nil, false
	}
	return o.Bootloader, true
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootLocalDiskAllOf) HasBootloader() bool {
	if o != nil && o.Bootloader != nil {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given BootBootloader and assigns it to the Bootloader field.
func (o *BootLocalDiskAllOf) SetBootloader(v BootBootloader) {
	o.Bootloader = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootLocalDiskAllOf) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootLocalDiskAllOf) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootLocalDiskAllOf) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootLocalDiskAllOf) SetSlot(v string) {
	o.Slot = &v
}

func (o BootLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bootloader != nil {
		toSerialize["Bootloader"] = o.Bootloader
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootLocalDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootLocalDiskAllOf := _BootLocalDiskAllOf{}

	if err = json.Unmarshal(bytes, &varBootLocalDiskAllOf); err == nil {
		*o = BootLocalDiskAllOf(varBootLocalDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Bootloader")
		delete(additionalProperties, "Slot")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootLocalDiskAllOf struct {
	value *BootLocalDiskAllOf
	isSet bool
}

func (v NullableBootLocalDiskAllOf) Get() *BootLocalDiskAllOf {
	return v.value
}

func (v *NullableBootLocalDiskAllOf) Set(val *BootLocalDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootLocalDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootLocalDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootLocalDiskAllOf(val *BootLocalDiskAllOf) *NullableBootLocalDiskAllOf {
	return &NullableBootLocalDiskAllOf{value: val, isSet: true}
}

func (v NullableBootLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootLocalDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
