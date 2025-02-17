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

// PoolAbstractBlockType Abstract base class for a block of identifiers.
type PoolAbstractBlockType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Number of identifiers this block can hold.
	Size                 *int64 `json:"Size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractBlockType PoolAbstractBlockType

// NewPoolAbstractBlockType instantiates a new PoolAbstractBlockType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractBlockType(classId string, objectType string) *PoolAbstractBlockType {
	this := PoolAbstractBlockType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPoolAbstractBlockTypeWithDefaults instantiates a new PoolAbstractBlockType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractBlockTypeWithDefaults() *PoolAbstractBlockType {
	this := PoolAbstractBlockType{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractBlockType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractBlockType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractBlockType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractBlockType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PoolAbstractBlockType) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockType) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PoolAbstractBlockType) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *PoolAbstractBlockType) SetSize(v int64) {
	o.Size = &v
}

func (o PoolAbstractBlockType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractBlockType) UnmarshalJSON(bytes []byte) (err error) {
	type PoolAbstractBlockTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Number of identifiers this block can hold.
		Size *int64 `json:"Size,omitempty"`
	}

	varPoolAbstractBlockTypeWithoutEmbeddedStruct := PoolAbstractBlockTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPoolAbstractBlockTypeWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractBlockType := _PoolAbstractBlockType{}
		varPoolAbstractBlockType.ClassId = varPoolAbstractBlockTypeWithoutEmbeddedStruct.ClassId
		varPoolAbstractBlockType.ObjectType = varPoolAbstractBlockTypeWithoutEmbeddedStruct.ObjectType
		varPoolAbstractBlockType.Size = varPoolAbstractBlockTypeWithoutEmbeddedStruct.Size
		*o = PoolAbstractBlockType(varPoolAbstractBlockType)
	} else {
		return err
	}

	varPoolAbstractBlockType := _PoolAbstractBlockType{}

	err = json.Unmarshal(bytes, &varPoolAbstractBlockType)
	if err == nil {
		o.MoBaseComplexType = varPoolAbstractBlockType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Size")

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

type NullablePoolAbstractBlockType struct {
	value *PoolAbstractBlockType
	isSet bool
}

func (v NullablePoolAbstractBlockType) Get() *PoolAbstractBlockType {
	return v.value
}

func (v *NullablePoolAbstractBlockType) Set(val *PoolAbstractBlockType) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractBlockType) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractBlockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractBlockType(val *PoolAbstractBlockType) *NullablePoolAbstractBlockType {
	return &NullablePoolAbstractBlockType{value: val, isSet: true}
}

func (v NullablePoolAbstractBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractBlockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
