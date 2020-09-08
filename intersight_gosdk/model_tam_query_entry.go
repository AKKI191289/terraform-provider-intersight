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

// TamQueryEntry Contains a set of queries, each with an integer priority. the queries are executed in the order of specified priority. The result of each query is used as an input to the query next in priority order.
type TamQueryEntry struct {
	MoBaseComplexType
	// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
	Name *string `json:"Name,omitempty"`
	// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
	Priority *int64 `json:"Priority,omitempty"`
	// A SparkSQL query to be used on a given data source.
	Query                *string `json:"Query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamQueryEntry TamQueryEntry

// NewTamQueryEntry instantiates a new TamQueryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamQueryEntry() *TamQueryEntry {
	this := TamQueryEntry{}
	return &this
}

// NewTamQueryEntryWithDefaults instantiates a new TamQueryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamQueryEntryWithDefaults() *TamQueryEntry {
	this := TamQueryEntry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamQueryEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamQueryEntry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamQueryEntry) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TamQueryEntry) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntry) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TamQueryEntry) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *TamQueryEntry) SetPriority(v int64) {
	o.Priority = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TamQueryEntry) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntry) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TamQueryEntry) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TamQueryEntry) SetQuery(v string) {
	o.Query = &v
}

func (o TamQueryEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamQueryEntry) UnmarshalJSON(bytes []byte) (err error) {
	type TamQueryEntryWithoutEmbeddedStruct struct {
		// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
		Name *string `json:"Name,omitempty"`
		// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
		Priority *int64 `json:"Priority,omitempty"`
		// A SparkSQL query to be used on a given data source.
		Query *string `json:"Query,omitempty"`
	}

	varTamQueryEntryWithoutEmbeddedStruct := TamQueryEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamQueryEntryWithoutEmbeddedStruct)
	if err == nil {
		varTamQueryEntry := _TamQueryEntry{}
		varTamQueryEntry.Name = varTamQueryEntryWithoutEmbeddedStruct.Name
		varTamQueryEntry.Priority = varTamQueryEntryWithoutEmbeddedStruct.Priority
		varTamQueryEntry.Query = varTamQueryEntryWithoutEmbeddedStruct.Query
		*o = TamQueryEntry(varTamQueryEntry)
	} else {
		return err
	}

	varTamQueryEntry := _TamQueryEntry{}

	err = json.Unmarshal(bytes, &varTamQueryEntry)
	if err == nil {
		o.MoBaseComplexType = varTamQueryEntry.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "Query")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableTamQueryEntry struct {
	value *TamQueryEntry
	isSet bool
}

func (v NullableTamQueryEntry) Get() *TamQueryEntry {
	return v.value
}

func (v *NullableTamQueryEntry) Set(val *TamQueryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTamQueryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTamQueryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamQueryEntry(val *TamQueryEntry) *NullableTamQueryEntry {
	return &NullableTamQueryEntry{value: val, isSet: true}
}

func (v NullableTamQueryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamQueryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
