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

// CondHclStatusDetailListAllOf struct for CondHclStatusDetailListAllOf
type CondHclStatusDetailListAllOf struct {
	// The total number of 'cond.HclStatusDetail' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'cond.HclStatusDetail' resources matching the request.
	Results              []CondHclStatusDetail `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondHclStatusDetailListAllOf CondHclStatusDetailListAllOf

// NewCondHclStatusDetailListAllOf instantiates a new CondHclStatusDetailListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondHclStatusDetailListAllOf() *CondHclStatusDetailListAllOf {
	this := CondHclStatusDetailListAllOf{}
	return &this
}

// NewCondHclStatusDetailListAllOfWithDefaults instantiates a new CondHclStatusDetailListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondHclStatusDetailListAllOfWithDefaults() *CondHclStatusDetailListAllOf {
	this := CondHclStatusDetailListAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CondHclStatusDetailListAllOf) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailListAllOf) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CondHclStatusDetailListAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CondHclStatusDetailListAllOf) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondHclStatusDetailListAllOf) GetResults() []CondHclStatusDetail {
	if o == nil {
		var ret []CondHclStatusDetail
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondHclStatusDetailListAllOf) GetResultsOk() (*[]CondHclStatusDetail, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CondHclStatusDetailListAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CondHclStatusDetail and assigns it to the Results field.
func (o *CondHclStatusDetailListAllOf) SetResults(v []CondHclStatusDetail) {
	o.Results = v
}

func (o CondHclStatusDetailListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondHclStatusDetailListAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondHclStatusDetailListAllOf := _CondHclStatusDetailListAllOf{}

	if err = json.Unmarshal(bytes, &varCondHclStatusDetailListAllOf); err == nil {
		*o = CondHclStatusDetailListAllOf(varCondHclStatusDetailListAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondHclStatusDetailListAllOf struct {
	value *CondHclStatusDetailListAllOf
	isSet bool
}

func (v NullableCondHclStatusDetailListAllOf) Get() *CondHclStatusDetailListAllOf {
	return v.value
}

func (v *NullableCondHclStatusDetailListAllOf) Set(val *CondHclStatusDetailListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondHclStatusDetailListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondHclStatusDetailListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondHclStatusDetailListAllOf(val *CondHclStatusDetailListAllOf) *NullableCondHclStatusDetailListAllOf {
	return &NullableCondHclStatusDetailListAllOf{value: val, isSet: true}
}

func (v NullableCondHclStatusDetailListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondHclStatusDetailListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
