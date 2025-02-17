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

// CapabilitySwitchCapabilityDefAllOf Definition of the list of properties defined in 'capability.SwitchCapabilityDef', excluding properties defined in parent classes.
type CapabilitySwitchCapabilityDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Product Identifier for a Switch/Fabric-Interconnect. * `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports. * `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * `unknown` - Unknown device type, usage is TBD.
	Pid *string `json:"Pid,omitempty"`
	// SKU information for Switch/Fabric-Interconnect.
	Sku *string `json:"Sku,omitempty"`
	// VID information for Switch/Fabric-Interconnect.
	Vid                  *string `json:"Vid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchCapabilityDefAllOf CapabilitySwitchCapabilityDefAllOf

// NewCapabilitySwitchCapabilityDefAllOf instantiates a new CapabilitySwitchCapabilityDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchCapabilityDefAllOf(classId string, objectType string) *CapabilitySwitchCapabilityDefAllOf {
	this := CapabilitySwitchCapabilityDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var pid string = "UCS-FI-6454"
	this.Pid = &pid
	return &this
}

// NewCapabilitySwitchCapabilityDefAllOfWithDefaults instantiates a new CapabilitySwitchCapabilityDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchCapabilityDefAllOfWithDefaults() *CapabilitySwitchCapabilityDefAllOf {
	this := CapabilitySwitchCapabilityDefAllOf{}
	var pid string = "UCS-FI-6454"
	this.Pid = &pid
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchCapabilityDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchCapabilityDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchCapabilityDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchCapabilityDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDefAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *CapabilitySwitchCapabilityDefAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDefAllOf) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *CapabilitySwitchCapabilityDefAllOf) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *CapabilitySwitchCapabilityDefAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *CapabilitySwitchCapabilityDefAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *CapabilitySwitchCapabilityDefAllOf) SetVid(v string) {
	o.Vid = &v
}

func (o CapabilitySwitchCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchCapabilityDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySwitchCapabilityDefAllOf := _CapabilitySwitchCapabilityDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySwitchCapabilityDefAllOf); err == nil {
		*o = CapabilitySwitchCapabilityDefAllOf(varCapabilitySwitchCapabilityDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchCapabilityDefAllOf struct {
	value *CapabilitySwitchCapabilityDefAllOf
	isSet bool
}

func (v NullableCapabilitySwitchCapabilityDefAllOf) Get() *CapabilitySwitchCapabilityDefAllOf {
	return v.value
}

func (v *NullableCapabilitySwitchCapabilityDefAllOf) Set(val *CapabilitySwitchCapabilityDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchCapabilityDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchCapabilityDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchCapabilityDefAllOf(val *CapabilitySwitchCapabilityDefAllOf) *NullableCapabilitySwitchCapabilityDefAllOf {
	return &NullableCapabilitySwitchCapabilityDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySwitchCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchCapabilityDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
