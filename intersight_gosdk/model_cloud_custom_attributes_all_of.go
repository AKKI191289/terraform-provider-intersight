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

// CloudCustomAttributesAllOf Definition of the list of properties defined in 'cloud.CustomAttributes', excluding properties defined in parent classes.
type CloudCustomAttributesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of an attribute. If used as a key-value pair then this field represents the key.
	AttributeName *string `json:"AttributeName,omitempty"`
	// The data type for attributeValue. For e.g. string, int, float.
	AttributeType *string `json:"AttributeType,omitempty"`
	// The attribute value. If used as a key-value pair then this field represents the value.
	AttributeValue       *string `json:"AttributeValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudCustomAttributesAllOf CloudCustomAttributesAllOf

// NewCloudCustomAttributesAllOf instantiates a new CloudCustomAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCustomAttributesAllOf(classId string, objectType string) *CloudCustomAttributesAllOf {
	this := CloudCustomAttributesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudCustomAttributesAllOfWithDefaults instantiates a new CloudCustomAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCustomAttributesAllOfWithDefaults() *CloudCustomAttributesAllOf {
	this := CloudCustomAttributesAllOf{}
	var classId string = "cloud.CustomAttributes"
	this.ClassId = classId
	var objectType string = "cloud.CustomAttributes"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudCustomAttributesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudCustomAttributesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudCustomAttributesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudCustomAttributesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudCustomAttributesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudCustomAttributesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *CloudCustomAttributesAllOf) GetAttributeName() string {
	if o == nil || o.AttributeName == nil {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCustomAttributesAllOf) GetAttributeNameOk() (*string, bool) {
	if o == nil || o.AttributeName == nil {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *CloudCustomAttributesAllOf) HasAttributeName() bool {
	if o != nil && o.AttributeName != nil {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *CloudCustomAttributesAllOf) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeType returns the AttributeType field value if set, zero value otherwise.
func (o *CloudCustomAttributesAllOf) GetAttributeType() string {
	if o == nil || o.AttributeType == nil {
		var ret string
		return ret
	}
	return *o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCustomAttributesAllOf) GetAttributeTypeOk() (*string, bool) {
	if o == nil || o.AttributeType == nil {
		return nil, false
	}
	return o.AttributeType, true
}

// HasAttributeType returns a boolean if a field has been set.
func (o *CloudCustomAttributesAllOf) HasAttributeType() bool {
	if o != nil && o.AttributeType != nil {
		return true
	}

	return false
}

// SetAttributeType gets a reference to the given string and assigns it to the AttributeType field.
func (o *CloudCustomAttributesAllOf) SetAttributeType(v string) {
	o.AttributeType = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *CloudCustomAttributesAllOf) GetAttributeValue() string {
	if o == nil || o.AttributeValue == nil {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCustomAttributesAllOf) GetAttributeValueOk() (*string, bool) {
	if o == nil || o.AttributeValue == nil {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *CloudCustomAttributesAllOf) HasAttributeValue() bool {
	if o != nil && o.AttributeValue != nil {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *CloudCustomAttributesAllOf) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o CloudCustomAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AttributeName != nil {
		toSerialize["AttributeName"] = o.AttributeName
	}
	if o.AttributeType != nil {
		toSerialize["AttributeType"] = o.AttributeType
	}
	if o.AttributeValue != nil {
		toSerialize["AttributeValue"] = o.AttributeValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudCustomAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudCustomAttributesAllOf := _CloudCustomAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varCloudCustomAttributesAllOf); err == nil {
		*o = CloudCustomAttributesAllOf(varCloudCustomAttributesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttributeName")
		delete(additionalProperties, "AttributeType")
		delete(additionalProperties, "AttributeValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudCustomAttributesAllOf struct {
	value *CloudCustomAttributesAllOf
	isSet bool
}

func (v NullableCloudCustomAttributesAllOf) Get() *CloudCustomAttributesAllOf {
	return v.value
}

func (v *NullableCloudCustomAttributesAllOf) Set(val *CloudCustomAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCustomAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCustomAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCustomAttributesAllOf(val *CloudCustomAttributesAllOf) *NullableCloudCustomAttributesAllOf {
	return &NullableCloudCustomAttributesAllOf{value: val, isSet: true}
}

func (v NullableCloudCustomAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCustomAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
