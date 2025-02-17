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

// PkixEddsaKeySpecAllOf Definition of the list of properties defined in 'pkix.EddsaKeySpec', excluding properties defined in parent classes.
type PkixEddsaKeySpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The EdDSA algorithm, as defined in RFC 8032. * `Ed25519` - The edwards25519 curve, as defined in RFC 8032 section 5.1. * `Ed25519ph` - The edwards25519 curve for the PureEdDSA variant. * `Ed25519ctx` - The edwards25519 curve for the HashEdDSA variant.
	Algorithm            *string `json:"Algorithm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixEddsaKeySpecAllOf PkixEddsaKeySpecAllOf

// NewPkixEddsaKeySpecAllOf instantiates a new PkixEddsaKeySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixEddsaKeySpecAllOf(classId string, objectType string) *PkixEddsaKeySpecAllOf {
	this := PkixEddsaKeySpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// NewPkixEddsaKeySpecAllOfWithDefaults instantiates a new PkixEddsaKeySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixEddsaKeySpecAllOfWithDefaults() *PkixEddsaKeySpecAllOf {
	this := PkixEddsaKeySpecAllOf{}
	var classId string = "pkix.EddsaKeySpec"
	this.ClassId = classId
	var objectType string = "pkix.EddsaKeySpec"
	this.ObjectType = objectType
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixEddsaKeySpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixEddsaKeySpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixEddsaKeySpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixEddsaKeySpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixEddsaKeySpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixEddsaKeySpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PkixEddsaKeySpecAllOf) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixEddsaKeySpecAllOf) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PkixEddsaKeySpecAllOf) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *PkixEddsaKeySpecAllOf) SetAlgorithm(v string) {
	o.Algorithm = &v
}

func (o PkixEddsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Algorithm != nil {
		toSerialize["Algorithm"] = o.Algorithm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixEddsaKeySpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPkixEddsaKeySpecAllOf := _PkixEddsaKeySpecAllOf{}

	if err = json.Unmarshal(bytes, &varPkixEddsaKeySpecAllOf); err == nil {
		*o = PkixEddsaKeySpecAllOf(varPkixEddsaKeySpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Algorithm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePkixEddsaKeySpecAllOf struct {
	value *PkixEddsaKeySpecAllOf
	isSet bool
}

func (v NullablePkixEddsaKeySpecAllOf) Get() *PkixEddsaKeySpecAllOf {
	return v.value
}

func (v *NullablePkixEddsaKeySpecAllOf) Set(val *PkixEddsaKeySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixEddsaKeySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixEddsaKeySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixEddsaKeySpecAllOf(val *PkixEddsaKeySpecAllOf) *NullablePkixEddsaKeySpecAllOf {
	return &NullablePkixEddsaKeySpecAllOf{value: val, isSet: true}
}

func (v NullablePkixEddsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixEddsaKeySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
