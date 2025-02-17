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

// TelemetryDruidSelectorFilterAllOf struct for TelemetryDruidSelectorFilterAllOf
type TelemetryDruidSelectorFilterAllOf struct {
	// The name of a dimension.
	Dimension string `json:"dimension"`
	// The value of a dimension.
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidSelectorFilterAllOf TelemetryDruidSelectorFilterAllOf

// NewTelemetryDruidSelectorFilterAllOf instantiates a new TelemetryDruidSelectorFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSelectorFilterAllOf(dimension string, value string) *TelemetryDruidSelectorFilterAllOf {
	this := TelemetryDruidSelectorFilterAllOf{}
	this.Dimension = dimension
	this.Value = value
	return &this
}

// NewTelemetryDruidSelectorFilterAllOfWithDefaults instantiates a new TelemetryDruidSelectorFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSelectorFilterAllOfWithDefaults() *TelemetryDruidSelectorFilterAllOf {
	this := TelemetryDruidSelectorFilterAllOf{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidSelectorFilterAllOf) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSelectorFilterAllOf) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidSelectorFilterAllOf) SetDimension(v string) {
	o.Dimension = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidSelectorFilterAllOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSelectorFilterAllOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidSelectorFilterAllOf) SetValue(v string) {
	o.Value = v
}

func (o TelemetryDruidSelectorFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidSelectorFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidSelectorFilterAllOf := _TelemetryDruidSelectorFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidSelectorFilterAllOf); err == nil {
		*o = TelemetryDruidSelectorFilterAllOf(varTelemetryDruidSelectorFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidSelectorFilterAllOf struct {
	value *TelemetryDruidSelectorFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidSelectorFilterAllOf) Get() *TelemetryDruidSelectorFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidSelectorFilterAllOf) Set(val *TelemetryDruidSelectorFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSelectorFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSelectorFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSelectorFilterAllOf(val *TelemetryDruidSelectorFilterAllOf) *NullableTelemetryDruidSelectorFilterAllOf {
	return &NullableTelemetryDruidSelectorFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidSelectorFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSelectorFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
