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
)

// ContentTextParameterAllOf Definition of the list of properties defined in 'content.TextParameter', excluding properties defined in parent classes.
type ContentTextParameterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data.
	IsDelimiter *bool `json:"IsDelimiter,omitempty"`
	// Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match.
	IsNextCaptureOnSameLine *bool `json:"IsNextCaptureOnSameLine,omitempty"`
	// Regular expression of the line containing the data to be extracted from text content.
	RegexLine            *string `json:"RegexLine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentTextParameterAllOf ContentTextParameterAllOf

// NewContentTextParameterAllOf instantiates a new ContentTextParameterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentTextParameterAllOf(classId string, objectType string) *ContentTextParameterAllOf {
	this := ContentTextParameterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isDelimiter bool = false
	this.IsDelimiter = &isDelimiter
	var isNextCaptureOnSameLine bool = false
	this.IsNextCaptureOnSameLine = &isNextCaptureOnSameLine
	return &this
}

// NewContentTextParameterAllOfWithDefaults instantiates a new ContentTextParameterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTextParameterAllOfWithDefaults() *ContentTextParameterAllOf {
	this := ContentTextParameterAllOf{}
	var classId string = "content.TextParameter"
	this.ClassId = classId
	var objectType string = "content.TextParameter"
	this.ObjectType = objectType
	var isDelimiter bool = false
	this.IsDelimiter = &isDelimiter
	var isNextCaptureOnSameLine bool = false
	this.IsNextCaptureOnSameLine = &isNextCaptureOnSameLine
	return &this
}

// GetClassId returns the ClassId field value
func (o *ContentTextParameterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ContentTextParameterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ContentTextParameterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ContentTextParameterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContentTextParameterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContentTextParameterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsDelimiter returns the IsDelimiter field value if set, zero value otherwise.
func (o *ContentTextParameterAllOf) GetIsDelimiter() bool {
	if o == nil || o.IsDelimiter == nil {
		var ret bool
		return ret
	}
	return *o.IsDelimiter
}

// GetIsDelimiterOk returns a tuple with the IsDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameterAllOf) GetIsDelimiterOk() (*bool, bool) {
	if o == nil || o.IsDelimiter == nil {
		return nil, false
	}
	return o.IsDelimiter, true
}

// HasIsDelimiter returns a boolean if a field has been set.
func (o *ContentTextParameterAllOf) HasIsDelimiter() bool {
	if o != nil && o.IsDelimiter != nil {
		return true
	}

	return false
}

// SetIsDelimiter gets a reference to the given bool and assigns it to the IsDelimiter field.
func (o *ContentTextParameterAllOf) SetIsDelimiter(v bool) {
	o.IsDelimiter = &v
}

// GetIsNextCaptureOnSameLine returns the IsNextCaptureOnSameLine field value if set, zero value otherwise.
func (o *ContentTextParameterAllOf) GetIsNextCaptureOnSameLine() bool {
	if o == nil || o.IsNextCaptureOnSameLine == nil {
		var ret bool
		return ret
	}
	return *o.IsNextCaptureOnSameLine
}

// GetIsNextCaptureOnSameLineOk returns a tuple with the IsNextCaptureOnSameLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameterAllOf) GetIsNextCaptureOnSameLineOk() (*bool, bool) {
	if o == nil || o.IsNextCaptureOnSameLine == nil {
		return nil, false
	}
	return o.IsNextCaptureOnSameLine, true
}

// HasIsNextCaptureOnSameLine returns a boolean if a field has been set.
func (o *ContentTextParameterAllOf) HasIsNextCaptureOnSameLine() bool {
	if o != nil && o.IsNextCaptureOnSameLine != nil {
		return true
	}

	return false
}

// SetIsNextCaptureOnSameLine gets a reference to the given bool and assigns it to the IsNextCaptureOnSameLine field.
func (o *ContentTextParameterAllOf) SetIsNextCaptureOnSameLine(v bool) {
	o.IsNextCaptureOnSameLine = &v
}

// GetRegexLine returns the RegexLine field value if set, zero value otherwise.
func (o *ContentTextParameterAllOf) GetRegexLine() string {
	if o == nil || o.RegexLine == nil {
		var ret string
		return ret
	}
	return *o.RegexLine
}

// GetRegexLineOk returns a tuple with the RegexLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameterAllOf) GetRegexLineOk() (*string, bool) {
	if o == nil || o.RegexLine == nil {
		return nil, false
	}
	return o.RegexLine, true
}

// HasRegexLine returns a boolean if a field has been set.
func (o *ContentTextParameterAllOf) HasRegexLine() bool {
	if o != nil && o.RegexLine != nil {
		return true
	}

	return false
}

// SetRegexLine gets a reference to the given string and assigns it to the RegexLine field.
func (o *ContentTextParameterAllOf) SetRegexLine(v string) {
	o.RegexLine = &v
}

func (o ContentTextParameterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsDelimiter != nil {
		toSerialize["IsDelimiter"] = o.IsDelimiter
	}
	if o.IsNextCaptureOnSameLine != nil {
		toSerialize["IsNextCaptureOnSameLine"] = o.IsNextCaptureOnSameLine
	}
	if o.RegexLine != nil {
		toSerialize["RegexLine"] = o.RegexLine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentTextParameterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varContentTextParameterAllOf := _ContentTextParameterAllOf{}

	if err = json.Unmarshal(bytes, &varContentTextParameterAllOf); err == nil {
		*o = ContentTextParameterAllOf(varContentTextParameterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsDelimiter")
		delete(additionalProperties, "IsNextCaptureOnSameLine")
		delete(additionalProperties, "RegexLine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentTextParameterAllOf struct {
	value *ContentTextParameterAllOf
	isSet bool
}

func (v NullableContentTextParameterAllOf) Get() *ContentTextParameterAllOf {
	return v.value
}

func (v *NullableContentTextParameterAllOf) Set(val *ContentTextParameterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentTextParameterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentTextParameterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentTextParameterAllOf(val *ContentTextParameterAllOf) *NullableContentTextParameterAllOf {
	return &NullableContentTextParameterAllOf{value: val, isSet: true}
}

func (v NullableContentTextParameterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentTextParameterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
