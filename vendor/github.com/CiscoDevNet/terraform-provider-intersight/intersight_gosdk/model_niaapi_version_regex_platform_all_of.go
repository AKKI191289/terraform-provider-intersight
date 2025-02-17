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

// NiaapiVersionRegexPlatformAllOf Definition of the list of properties defined in 'niaapi.VersionRegexPlatform', excluding properties defined in parent classes.
type NiaapiVersionRegexPlatformAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// All long live version Regex pattern.
	Anyllregex           *string                     `json:"Anyllregex,omitempty"`
	Currentlltrain       NullableNiaapiSoftwareRegex `json:"Currentlltrain,omitempty"`
	Latestsltrain        NullableNiaapiSoftwareRegex `json:"Latestsltrain,omitempty"`
	Sltrain              []NiaapiSoftwareRegex       `json:"Sltrain,omitempty"`
	Upcominglltrain      NullableNiaapiSoftwareRegex `json:"Upcominglltrain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegexPlatformAllOf NiaapiVersionRegexPlatformAllOf

// NewNiaapiVersionRegexPlatformAllOf instantiates a new NiaapiVersionRegexPlatformAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegexPlatformAllOf(classId string, objectType string) *NiaapiVersionRegexPlatformAllOf {
	this := NiaapiVersionRegexPlatformAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiVersionRegexPlatformAllOfWithDefaults instantiates a new NiaapiVersionRegexPlatformAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexPlatformAllOfWithDefaults() *NiaapiVersionRegexPlatformAllOf {
	this := NiaapiVersionRegexPlatformAllOf{}
	var classId string = "niaapi.VersionRegexPlatform"
	this.ClassId = classId
	var objectType string = "niaapi.VersionRegexPlatform"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiVersionRegexPlatformAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiVersionRegexPlatformAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiVersionRegexPlatformAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiVersionRegexPlatformAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnyllregex returns the Anyllregex field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregex() string {
	if o == nil || o.Anyllregex == nil {
		var ret string
		return ret
	}
	return *o.Anyllregex
}

// GetAnyllregexOk returns a tuple with the Anyllregex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregexOk() (*string, bool) {
	if o == nil || o.Anyllregex == nil {
		return nil, false
	}
	return o.Anyllregex, true
}

// HasAnyllregex returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasAnyllregex() bool {
	if o != nil && o.Anyllregex != nil {
		return true
	}

	return false
}

// SetAnyllregex gets a reference to the given string and assigns it to the Anyllregex field.
func (o *NiaapiVersionRegexPlatformAllOf) SetAnyllregex(v string) {
	o.Anyllregex = &v
}

// GetCurrentlltrain returns the Currentlltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrain() NiaapiSoftwareRegex {
	if o == nil || o.Currentlltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Currentlltrain.Get()
}

// GetCurrentlltrainOk returns a tuple with the Currentlltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currentlltrain.Get(), o.Currentlltrain.IsSet()
}

// HasCurrentlltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasCurrentlltrain() bool {
	if o != nil && o.Currentlltrain.IsSet() {
		return true
	}

	return false
}

// SetCurrentlltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Currentlltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetCurrentlltrain(v NiaapiSoftwareRegex) {
	o.Currentlltrain.Set(&v)
}

// SetCurrentlltrainNil sets the value for Currentlltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) SetCurrentlltrainNil() {
	o.Currentlltrain.Set(nil)
}

// UnsetCurrentlltrain ensures that no value is present for Currentlltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) UnsetCurrentlltrain() {
	o.Currentlltrain.Unset()
}

// GetLatestsltrain returns the Latestsltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrain() NiaapiSoftwareRegex {
	if o == nil || o.Latestsltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Latestsltrain.Get()
}

// GetLatestsltrainOk returns a tuple with the Latestsltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latestsltrain.Get(), o.Latestsltrain.IsSet()
}

// HasLatestsltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasLatestsltrain() bool {
	if o != nil && o.Latestsltrain.IsSet() {
		return true
	}

	return false
}

// SetLatestsltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Latestsltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetLatestsltrain(v NiaapiSoftwareRegex) {
	o.Latestsltrain.Set(&v)
}

// SetLatestsltrainNil sets the value for Latestsltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) SetLatestsltrainNil() {
	o.Latestsltrain.Set(nil)
}

// UnsetLatestsltrain ensures that no value is present for Latestsltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) UnsetLatestsltrain() {
	o.Latestsltrain.Unset()
}

// GetSltrain returns the Sltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatformAllOf) GetSltrain() []NiaapiSoftwareRegex {
	if o == nil {
		var ret []NiaapiSoftwareRegex
		return ret
	}
	return o.Sltrain
}

// GetSltrainOk returns a tuple with the Sltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatformAllOf) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool) {
	if o == nil || o.Sltrain == nil {
		return nil, false
	}
	return &o.Sltrain, true
}

// HasSltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasSltrain() bool {
	if o != nil && o.Sltrain != nil {
		return true
	}

	return false
}

// SetSltrain gets a reference to the given []NiaapiSoftwareRegex and assigns it to the Sltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetSltrain(v []NiaapiSoftwareRegex) {
	o.Sltrain = v
}

// GetUpcominglltrain returns the Upcominglltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrain() NiaapiSoftwareRegex {
	if o == nil || o.Upcominglltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Upcominglltrain.Get()
}

// GetUpcominglltrainOk returns a tuple with the Upcominglltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil {
		return nil, false
	}
	return o.Upcominglltrain.Get(), o.Upcominglltrain.IsSet()
}

// HasUpcominglltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasUpcominglltrain() bool {
	if o != nil && o.Upcominglltrain.IsSet() {
		return true
	}

	return false
}

// SetUpcominglltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Upcominglltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetUpcominglltrain(v NiaapiSoftwareRegex) {
	o.Upcominglltrain.Set(&v)
}

// SetUpcominglltrainNil sets the value for Upcominglltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) SetUpcominglltrainNil() {
	o.Upcominglltrain.Set(nil)
}

// UnsetUpcominglltrain ensures that no value is present for Upcominglltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatformAllOf) UnsetUpcominglltrain() {
	o.Upcominglltrain.Unset()
}

func (o NiaapiVersionRegexPlatformAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Anyllregex != nil {
		toSerialize["Anyllregex"] = o.Anyllregex
	}
	if o.Currentlltrain.IsSet() {
		toSerialize["Currentlltrain"] = o.Currentlltrain.Get()
	}
	if o.Latestsltrain.IsSet() {
		toSerialize["Latestsltrain"] = o.Latestsltrain.Get()
	}
	if o.Sltrain != nil {
		toSerialize["Sltrain"] = o.Sltrain
	}
	if o.Upcominglltrain.IsSet() {
		toSerialize["Upcominglltrain"] = o.Upcominglltrain.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiVersionRegexPlatformAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiVersionRegexPlatformAllOf := _NiaapiVersionRegexPlatformAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiVersionRegexPlatformAllOf); err == nil {
		*o = NiaapiVersionRegexPlatformAllOf(varNiaapiVersionRegexPlatformAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Anyllregex")
		delete(additionalProperties, "Currentlltrain")
		delete(additionalProperties, "Latestsltrain")
		delete(additionalProperties, "Sltrain")
		delete(additionalProperties, "Upcominglltrain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiVersionRegexPlatformAllOf struct {
	value *NiaapiVersionRegexPlatformAllOf
	isSet bool
}

func (v NullableNiaapiVersionRegexPlatformAllOf) Get() *NiaapiVersionRegexPlatformAllOf {
	return v.value
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) Set(val *NiaapiVersionRegexPlatformAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegexPlatformAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegexPlatformAllOf(val *NiaapiVersionRegexPlatformAllOf) *NullableNiaapiVersionRegexPlatformAllOf {
	return &NullableNiaapiVersionRegexPlatformAllOf{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegexPlatformAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
