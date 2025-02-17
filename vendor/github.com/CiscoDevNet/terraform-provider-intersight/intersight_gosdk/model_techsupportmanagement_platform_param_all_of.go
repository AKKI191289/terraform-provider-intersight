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

// TechsupportmanagementPlatformParamAllOf Definition of the list of properties defined in 'techsupportmanagement.PlatformParam', excluding properties defined in parent classes.
type TechsupportmanagementPlatformParamAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CollectionType specifies if basic or detailed techsupport needs to be collected. * `1` - Collect basic techsupport. * `2` - Collect detailed techsupport.
	CollectionType *int32 `json:"CollectionType,omitempty"`
	// Controls whether to include core file in the techsupport bundle.
	IncludeCore          *bool `json:"IncludeCore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementPlatformParamAllOf TechsupportmanagementPlatformParamAllOf

// NewTechsupportmanagementPlatformParamAllOf instantiates a new TechsupportmanagementPlatformParamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementPlatformParamAllOf(classId string, objectType string) *TechsupportmanagementPlatformParamAllOf {
	this := TechsupportmanagementPlatformParamAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var collectionType int32 = 1
	this.CollectionType = &collectionType
	return &this
}

// NewTechsupportmanagementPlatformParamAllOfWithDefaults instantiates a new TechsupportmanagementPlatformParamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementPlatformParamAllOfWithDefaults() *TechsupportmanagementPlatformParamAllOf {
	this := TechsupportmanagementPlatformParamAllOf{}
	var classId string = "techsupportmanagement.PlatformParam"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.PlatformParam"
	this.ObjectType = objectType
	var collectionType int32 = 1
	this.CollectionType = &collectionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementPlatformParamAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementPlatformParamAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementPlatformParamAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementPlatformParamAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementPlatformParamAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementPlatformParamAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCollectionType returns the CollectionType field value if set, zero value otherwise.
func (o *TechsupportmanagementPlatformParamAllOf) GetCollectionType() int32 {
	if o == nil || o.CollectionType == nil {
		var ret int32
		return ret
	}
	return *o.CollectionType
}

// GetCollectionTypeOk returns a tuple with the CollectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementPlatformParamAllOf) GetCollectionTypeOk() (*int32, bool) {
	if o == nil || o.CollectionType == nil {
		return nil, false
	}
	return o.CollectionType, true
}

// HasCollectionType returns a boolean if a field has been set.
func (o *TechsupportmanagementPlatformParamAllOf) HasCollectionType() bool {
	if o != nil && o.CollectionType != nil {
		return true
	}

	return false
}

// SetCollectionType gets a reference to the given int32 and assigns it to the CollectionType field.
func (o *TechsupportmanagementPlatformParamAllOf) SetCollectionType(v int32) {
	o.CollectionType = &v
}

// GetIncludeCore returns the IncludeCore field value if set, zero value otherwise.
func (o *TechsupportmanagementPlatformParamAllOf) GetIncludeCore() bool {
	if o == nil || o.IncludeCore == nil {
		var ret bool
		return ret
	}
	return *o.IncludeCore
}

// GetIncludeCoreOk returns a tuple with the IncludeCore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementPlatformParamAllOf) GetIncludeCoreOk() (*bool, bool) {
	if o == nil || o.IncludeCore == nil {
		return nil, false
	}
	return o.IncludeCore, true
}

// HasIncludeCore returns a boolean if a field has been set.
func (o *TechsupportmanagementPlatformParamAllOf) HasIncludeCore() bool {
	if o != nil && o.IncludeCore != nil {
		return true
	}

	return false
}

// SetIncludeCore gets a reference to the given bool and assigns it to the IncludeCore field.
func (o *TechsupportmanagementPlatformParamAllOf) SetIncludeCore(v bool) {
	o.IncludeCore = &v
}

func (o TechsupportmanagementPlatformParamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CollectionType != nil {
		toSerialize["CollectionType"] = o.CollectionType
	}
	if o.IncludeCore != nil {
		toSerialize["IncludeCore"] = o.IncludeCore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementPlatformParamAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTechsupportmanagementPlatformParamAllOf := _TechsupportmanagementPlatformParamAllOf{}

	if err = json.Unmarshal(bytes, &varTechsupportmanagementPlatformParamAllOf); err == nil {
		*o = TechsupportmanagementPlatformParamAllOf(varTechsupportmanagementPlatformParamAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CollectionType")
		delete(additionalProperties, "IncludeCore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTechsupportmanagementPlatformParamAllOf struct {
	value *TechsupportmanagementPlatformParamAllOf
	isSet bool
}

func (v NullableTechsupportmanagementPlatformParamAllOf) Get() *TechsupportmanagementPlatformParamAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementPlatformParamAllOf) Set(val *TechsupportmanagementPlatformParamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementPlatformParamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementPlatformParamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementPlatformParamAllOf(val *TechsupportmanagementPlatformParamAllOf) *NullableTechsupportmanagementPlatformParamAllOf {
	return &NullableTechsupportmanagementPlatformParamAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementPlatformParamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementPlatformParamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
