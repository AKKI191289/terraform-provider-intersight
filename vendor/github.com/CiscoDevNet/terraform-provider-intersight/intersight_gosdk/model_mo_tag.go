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

// MoTag An arbitrary key and value pair that can be used to tag REST resources and organize managed objects by assigning meta-data tags to any object.
type MoTag struct {
	// The string representation of a tag key.
	Key *string `json:"Key,omitempty"`
	// The string representation of a tag value.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoTag MoTag

// NewMoTag instantiates a new MoTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoTag() *MoTag {
	this := MoTag{}
	return &this
}

// NewMoTagWithDefaults instantiates a new MoTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoTagWithDefaults() *MoTag {
	this := MoTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MoTag) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTag) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MoTag) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MoTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MoTag) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTag) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MoTag) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MoTag) SetValue(v string) {
	o.Value = &v
}

func (o MoTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoTag) UnmarshalJSON(bytes []byte) (err error) {
	varMoTag := _MoTag{}

	if err = json.Unmarshal(bytes, &varMoTag); err == nil {
		*o = MoTag(varMoTag)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoTag struct {
	value *MoTag
	isSet bool
}

func (v NullableMoTag) Get() *MoTag {
	return v.value
}

func (v *NullableMoTag) Set(val *MoTag) {
	v.value = val
	v.isSet = true
}

func (v NullableMoTag) IsSet() bool {
	return v.isSet
}

func (v *NullableMoTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoTag(val *MoTag) *NullableMoTag {
	return &NullableMoTag{value: val, isSet: true}
}

func (v NullableMoTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
