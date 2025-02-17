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

// PolicyAbstractProfileAllOf Definition of the list of properties defined in 'policy.AbstractProfile', excluding properties defined in parent classes.
type PolicyAbstractProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Description of the profile.
	Description *string `json:"Description,omitempty"`
	// Name of the concrete profile.
	Name *string `json:"Name,omitempty"`
	// Defines the type of the profile. Accepted value is instance. * `instance` - The profile defines the configuration for a specific instance of a target.
	Type                 *string                            `json:"Type,omitempty"`
	SrcTemplate          *PolicyAbstractProfileRelationship `json:"SrcTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractProfileAllOf PolicyAbstractProfileAllOf

// NewPolicyAbstractProfileAllOf instantiates a new PolicyAbstractProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractProfileAllOf(classId string, objectType string) *PolicyAbstractProfileAllOf {
	this := PolicyAbstractProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// NewPolicyAbstractProfileAllOfWithDefaults instantiates a new PolicyAbstractProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractProfileAllOfWithDefaults() *PolicyAbstractProfileAllOf {
	this := PolicyAbstractProfileAllOf{}
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyAbstractProfileAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyAbstractProfileAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyAbstractProfileAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyAbstractProfileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyAbstractProfileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyAbstractProfileAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyAbstractProfileAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyAbstractProfileAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyAbstractProfileAllOf) SetType(v string) {
	o.Type = &v
}

// GetSrcTemplate returns the SrcTemplate field value if set, zero value otherwise.
func (o *PolicyAbstractProfileAllOf) GetSrcTemplate() PolicyAbstractProfileRelationship {
	if o == nil || o.SrcTemplate == nil {
		var ret PolicyAbstractProfileRelationship
		return ret
	}
	return *o.SrcTemplate
}

// GetSrcTemplateOk returns a tuple with the SrcTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfileAllOf) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool) {
	if o == nil || o.SrcTemplate == nil {
		return nil, false
	}
	return o.SrcTemplate, true
}

// HasSrcTemplate returns a boolean if a field has been set.
func (o *PolicyAbstractProfileAllOf) HasSrcTemplate() bool {
	if o != nil && o.SrcTemplate != nil {
		return true
	}

	return false
}

// SetSrcTemplate gets a reference to the given PolicyAbstractProfileRelationship and assigns it to the SrcTemplate field.
func (o *PolicyAbstractProfileAllOf) SetSrcTemplate(v PolicyAbstractProfileRelationship) {
	o.SrcTemplate = &v
}

func (o PolicyAbstractProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.SrcTemplate != nil {
		toSerialize["SrcTemplate"] = o.SrcTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAbstractProfileAllOf := _PolicyAbstractProfileAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyAbstractProfileAllOf); err == nil {
		*o = PolicyAbstractProfileAllOf(varPolicyAbstractProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "SrcTemplate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAbstractProfileAllOf struct {
	value *PolicyAbstractProfileAllOf
	isSet bool
}

func (v NullablePolicyAbstractProfileAllOf) Get() *PolicyAbstractProfileAllOf {
	return v.value
}

func (v *NullablePolicyAbstractProfileAllOf) Set(val *PolicyAbstractProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractProfileAllOf(val *PolicyAbstractProfileAllOf) *NullablePolicyAbstractProfileAllOf {
	return &NullablePolicyAbstractProfileAllOf{value: val, isSet: true}
}

func (v NullablePolicyAbstractProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
