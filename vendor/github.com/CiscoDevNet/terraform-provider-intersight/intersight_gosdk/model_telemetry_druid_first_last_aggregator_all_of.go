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

// TelemetryDruidFirstLastAggregatorAllOf struct for TelemetryDruidFirstLastAggregatorAllOf
type TelemetryDruidFirstLastAggregatorAllOf struct {
	// Output name for the first/last value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidFirstLastAggregatorAllOf TelemetryDruidFirstLastAggregatorAllOf

// NewTelemetryDruidFirstLastAggregatorAllOf instantiates a new TelemetryDruidFirstLastAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidFirstLastAggregatorAllOf(name string, fieldName string) *TelemetryDruidFirstLastAggregatorAllOf {
	this := TelemetryDruidFirstLastAggregatorAllOf{}
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidFirstLastAggregatorAllOfWithDefaults instantiates a new TelemetryDruidFirstLastAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidFirstLastAggregatorAllOfWithDefaults() *TelemetryDruidFirstLastAggregatorAllOf {
	this := TelemetryDruidFirstLastAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *TelemetryDruidFirstLastAggregatorAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFirstLastAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidFirstLastAggregatorAllOf) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidFirstLastAggregatorAllOf) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFirstLastAggregatorAllOf) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidFirstLastAggregatorAllOf) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidFirstLastAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidFirstLastAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidFirstLastAggregatorAllOf := _TelemetryDruidFirstLastAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidFirstLastAggregatorAllOf); err == nil {
		*o = TelemetryDruidFirstLastAggregatorAllOf(varTelemetryDruidFirstLastAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidFirstLastAggregatorAllOf struct {
	value *TelemetryDruidFirstLastAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidFirstLastAggregatorAllOf) Get() *TelemetryDruidFirstLastAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidFirstLastAggregatorAllOf) Set(val *TelemetryDruidFirstLastAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFirstLastAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFirstLastAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFirstLastAggregatorAllOf(val *TelemetryDruidFirstLastAggregatorAllOf) *NullableTelemetryDruidFirstLastAggregatorAllOf {
	return &NullableTelemetryDruidFirstLastAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidFirstLastAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFirstLastAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
