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

// AdapterPortChannelSettingsAllOf Definition of the list of properties defined in 'adapter.PortChannelSettings', excluding properties defined in parent classes.
type AdapterPortChannelSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters.
	Enabled              *bool `json:"Enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterPortChannelSettingsAllOf AdapterPortChannelSettingsAllOf

// NewAdapterPortChannelSettingsAllOf instantiates a new AdapterPortChannelSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterPortChannelSettingsAllOf(classId string, objectType string) *AdapterPortChannelSettingsAllOf {
	this := AdapterPortChannelSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewAdapterPortChannelSettingsAllOfWithDefaults instantiates a new AdapterPortChannelSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterPortChannelSettingsAllOfWithDefaults() *AdapterPortChannelSettingsAllOf {
	this := AdapterPortChannelSettingsAllOf{}
	var classId string = "adapter.PortChannelSettings"
	this.ClassId = classId
	var objectType string = "adapter.PortChannelSettings"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterPortChannelSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterPortChannelSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterPortChannelSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterPortChannelSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AdapterPortChannelSettingsAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettingsAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AdapterPortChannelSettingsAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AdapterPortChannelSettingsAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AdapterPortChannelSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterPortChannelSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterPortChannelSettingsAllOf := _AdapterPortChannelSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterPortChannelSettingsAllOf); err == nil {
		*o = AdapterPortChannelSettingsAllOf(varAdapterPortChannelSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterPortChannelSettingsAllOf struct {
	value *AdapterPortChannelSettingsAllOf
	isSet bool
}

func (v NullableAdapterPortChannelSettingsAllOf) Get() *AdapterPortChannelSettingsAllOf {
	return v.value
}

func (v *NullableAdapterPortChannelSettingsAllOf) Set(val *AdapterPortChannelSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterPortChannelSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterPortChannelSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterPortChannelSettingsAllOf(val *AdapterPortChannelSettingsAllOf) *NullableAdapterPortChannelSettingsAllOf {
	return &NullableAdapterPortChannelSettingsAllOf{value: val, isSet: true}
}

func (v NullableAdapterPortChannelSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterPortChannelSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
