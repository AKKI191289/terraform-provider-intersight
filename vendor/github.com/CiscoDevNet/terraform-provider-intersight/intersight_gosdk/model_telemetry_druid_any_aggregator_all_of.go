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

// TelemetryDruidAnyAggregatorAllOf struct for TelemetryDruidAnyAggregatorAllOf
type TelemetryDruidAnyAggregatorAllOf struct {
	// Output name for the 'any' value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidAnyAggregatorAllOf TelemetryDruidAnyAggregatorAllOf

// NewTelemetryDruidAnyAggregatorAllOf instantiates a new TelemetryDruidAnyAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidAnyAggregatorAllOf(name string, fieldName string) *TelemetryDruidAnyAggregatorAllOf {
	this := TelemetryDruidAnyAggregatorAllOf{}
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidAnyAggregatorAllOfWithDefaults instantiates a new TelemetryDruidAnyAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidAnyAggregatorAllOfWithDefaults() *TelemetryDruidAnyAggregatorAllOf {
	this := TelemetryDruidAnyAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *TelemetryDruidAnyAggregatorAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidAnyAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidAnyAggregatorAllOf) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidAnyAggregatorAllOf) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidAnyAggregatorAllOf) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidAnyAggregatorAllOf) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidAnyAggregatorAllOf) MarshalJSON() ([]byte, error) {
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

func (o *TelemetryDruidAnyAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidAnyAggregatorAllOf := _TelemetryDruidAnyAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidAnyAggregatorAllOf); err == nil {
		*o = TelemetryDruidAnyAggregatorAllOf(varTelemetryDruidAnyAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidAnyAggregatorAllOf struct {
	value *TelemetryDruidAnyAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidAnyAggregatorAllOf) Get() *TelemetryDruidAnyAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidAnyAggregatorAllOf) Set(val *TelemetryDruidAnyAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidAnyAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidAnyAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidAnyAggregatorAllOf(val *TelemetryDruidAnyAggregatorAllOf) *NullableTelemetryDruidAnyAggregatorAllOf {
	return &NullableTelemetryDruidAnyAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidAnyAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidAnyAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
