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
	"reflect"
	"strings"
)

// ApplianceGroupStatusList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type ApplianceGroupStatusList struct {
	MoBaseResponse
	// The total number of 'appliance.GroupStatus' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'appliance.GroupStatus' resources matching the request.
	Results              []ApplianceGroupStatus `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceGroupStatusList ApplianceGroupStatusList

// NewApplianceGroupStatusList instantiates a new ApplianceGroupStatusList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceGroupStatusList() *ApplianceGroupStatusList {
	this := ApplianceGroupStatusList{}
	return &this
}

// NewApplianceGroupStatusListWithDefaults instantiates a new ApplianceGroupStatusList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceGroupStatusListWithDefaults() *ApplianceGroupStatusList {
	this := ApplianceGroupStatusList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApplianceGroupStatusList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatusList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApplianceGroupStatusList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ApplianceGroupStatusList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceGroupStatusList) GetResults() []ApplianceGroupStatus {
	if o == nil {
		var ret []ApplianceGroupStatus
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceGroupStatusList) GetResultsOk() (*[]ApplianceGroupStatus, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApplianceGroupStatusList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApplianceGroupStatus and assigns it to the Results field.
func (o *ApplianceGroupStatusList) SetResults(v []ApplianceGroupStatus) {
	o.Results = v
}

func (o ApplianceGroupStatusList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
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

func (o *ApplianceGroupStatusList) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceGroupStatusListWithoutEmbeddedStruct struct {
		// The total number of 'appliance.GroupStatus' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'appliance.GroupStatus' resources matching the request.
		Results []ApplianceGroupStatus `json:"Results,omitempty"`
	}

	varApplianceGroupStatusListWithoutEmbeddedStruct := ApplianceGroupStatusListWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceGroupStatusListWithoutEmbeddedStruct)
	if err == nil {
		varApplianceGroupStatusList := _ApplianceGroupStatusList{}
		varApplianceGroupStatusList.Count = varApplianceGroupStatusListWithoutEmbeddedStruct.Count
		varApplianceGroupStatusList.Results = varApplianceGroupStatusListWithoutEmbeddedStruct.Results
		*o = ApplianceGroupStatusList(varApplianceGroupStatusList)
	} else {
		return err
	}

	varApplianceGroupStatusList := _ApplianceGroupStatusList{}

	err = json.Unmarshal(bytes, &varApplianceGroupStatusList)
	if err == nil {
		o.MoBaseResponse = varApplianceGroupStatusList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableApplianceGroupStatusList struct {
	value *ApplianceGroupStatusList
	isSet bool
}

func (v NullableApplianceGroupStatusList) Get() *ApplianceGroupStatusList {
	return v.value
}

func (v *NullableApplianceGroupStatusList) Set(val *ApplianceGroupStatusList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceGroupStatusList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceGroupStatusList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceGroupStatusList(val *ApplianceGroupStatusList) *NullableApplianceGroupStatusList {
	return &NullableApplianceGroupStatusList{value: val, isSet: true}
}

func (v NullableApplianceGroupStatusList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceGroupStatusList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
