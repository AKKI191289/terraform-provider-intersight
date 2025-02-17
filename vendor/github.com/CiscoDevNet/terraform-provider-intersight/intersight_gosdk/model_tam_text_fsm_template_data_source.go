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
	"reflect"
	"strings"
)

// TamTextFsmTemplateDataSource TextFsmTemplate based data source used for data collection from the managed devices.
type TamTextFsmTemplateDataSource struct {
	TamBaseDataSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Command used to gather data needed to evaluate field notice applicability.
	Cmd                  *string `json:"Cmd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamTextFsmTemplateDataSource TamTextFsmTemplateDataSource

// NewTamTextFsmTemplateDataSource instantiates a new TamTextFsmTemplateDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamTextFsmTemplateDataSource(classId string, objectType string) *TamTextFsmTemplateDataSource {
	this := TamTextFsmTemplateDataSource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamTextFsmTemplateDataSourceWithDefaults instantiates a new TamTextFsmTemplateDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamTextFsmTemplateDataSourceWithDefaults() *TamTextFsmTemplateDataSource {
	this := TamTextFsmTemplateDataSource{}
	var classId string = "tam.TextFsmTemplateDataSource"
	this.ClassId = classId
	var objectType string = "tam.TextFsmTemplateDataSource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamTextFsmTemplateDataSource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamTextFsmTemplateDataSource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamTextFsmTemplateDataSource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamTextFsmTemplateDataSource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *TamTextFsmTemplateDataSource) GetCmd() string {
	if o == nil || o.Cmd == nil {
		var ret string
		return ret
	}
	return *o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSource) GetCmdOk() (*string, bool) {
	if o == nil || o.Cmd == nil {
		return nil, false
	}
	return o.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *TamTextFsmTemplateDataSource) HasCmd() bool {
	if o != nil && o.Cmd != nil {
		return true
	}

	return false
}

// SetCmd gets a reference to the given string and assigns it to the Cmd field.
func (o *TamTextFsmTemplateDataSource) SetCmd(v string) {
	o.Cmd = &v
}

func (o TamTextFsmTemplateDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseDataSource, errTamBaseDataSource := json.Marshal(o.TamBaseDataSource)
	if errTamBaseDataSource != nil {
		return []byte{}, errTamBaseDataSource
	}
	errTamBaseDataSource = json.Unmarshal([]byte(serializedTamBaseDataSource), &toSerialize)
	if errTamBaseDataSource != nil {
		return []byte{}, errTamBaseDataSource
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cmd != nil {
		toSerialize["Cmd"] = o.Cmd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamTextFsmTemplateDataSource) UnmarshalJSON(bytes []byte) (err error) {
	type TamTextFsmTemplateDataSourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Command used to gather data needed to evaluate field notice applicability.
		Cmd *string `json:"Cmd,omitempty"`
	}

	varTamTextFsmTemplateDataSourceWithoutEmbeddedStruct := TamTextFsmTemplateDataSourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamTextFsmTemplateDataSourceWithoutEmbeddedStruct)
	if err == nil {
		varTamTextFsmTemplateDataSource := _TamTextFsmTemplateDataSource{}
		varTamTextFsmTemplateDataSource.ClassId = varTamTextFsmTemplateDataSourceWithoutEmbeddedStruct.ClassId
		varTamTextFsmTemplateDataSource.ObjectType = varTamTextFsmTemplateDataSourceWithoutEmbeddedStruct.ObjectType
		varTamTextFsmTemplateDataSource.Cmd = varTamTextFsmTemplateDataSourceWithoutEmbeddedStruct.Cmd
		*o = TamTextFsmTemplateDataSource(varTamTextFsmTemplateDataSource)
	} else {
		return err
	}

	varTamTextFsmTemplateDataSource := _TamTextFsmTemplateDataSource{}

	err = json.Unmarshal(bytes, &varTamTextFsmTemplateDataSource)
	if err == nil {
		o.TamBaseDataSource = varTamTextFsmTemplateDataSource.TamBaseDataSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cmd")

		// remove fields from embedded structs
		reflectTamBaseDataSource := reflect.ValueOf(o.TamBaseDataSource)
		for i := 0; i < reflectTamBaseDataSource.Type().NumField(); i++ {
			t := reflectTamBaseDataSource.Type().Field(i)

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

type NullableTamTextFsmTemplateDataSource struct {
	value *TamTextFsmTemplateDataSource
	isSet bool
}

func (v NullableTamTextFsmTemplateDataSource) Get() *TamTextFsmTemplateDataSource {
	return v.value
}

func (v *NullableTamTextFsmTemplateDataSource) Set(val *TamTextFsmTemplateDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTamTextFsmTemplateDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTamTextFsmTemplateDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamTextFsmTemplateDataSource(val *TamTextFsmTemplateDataSource) *NullableTamTextFsmTemplateDataSource {
	return &NullableTamTextFsmTemplateDataSource{value: val, isSet: true}
}

func (v NullableTamTextFsmTemplateDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamTextFsmTemplateDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
