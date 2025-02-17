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

// ApplianceCertificateSettingAllOf Definition of the list of properties defined in 'appliance.CertificateSetting', excluding properties defined in parent classes.
type ApplianceCertificateSettingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                      `json:"ObjectType"`
	Account              *IamAccountRelationship     `json:"Account,omitempty"`
	Certificate          *IamCertificateRelationship `json:"Certificate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceCertificateSettingAllOf ApplianceCertificateSettingAllOf

// NewApplianceCertificateSettingAllOf instantiates a new ApplianceCertificateSettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceCertificateSettingAllOf(classId string, objectType string) *ApplianceCertificateSettingAllOf {
	this := ApplianceCertificateSettingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceCertificateSettingAllOfWithDefaults instantiates a new ApplianceCertificateSettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceCertificateSettingAllOfWithDefaults() *ApplianceCertificateSettingAllOf {
	this := ApplianceCertificateSettingAllOf{}
	var classId string = "appliance.CertificateSetting"
	this.ClassId = classId
	var objectType string = "appliance.CertificateSetting"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceCertificateSettingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSettingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceCertificateSettingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceCertificateSettingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSettingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceCertificateSettingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceCertificateSettingAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSettingAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceCertificateSettingAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceCertificateSettingAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *ApplianceCertificateSettingAllOf) GetCertificate() IamCertificateRelationship {
	if o == nil || o.Certificate == nil {
		var ret IamCertificateRelationship
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSettingAllOf) GetCertificateOk() (*IamCertificateRelationship, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ApplianceCertificateSettingAllOf) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given IamCertificateRelationship and assigns it to the Certificate field.
func (o *ApplianceCertificateSettingAllOf) SetCertificate(v IamCertificateRelationship) {
	o.Certificate = &v
}

func (o ApplianceCertificateSettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceCertificateSettingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceCertificateSettingAllOf := _ApplianceCertificateSettingAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceCertificateSettingAllOf); err == nil {
		*o = ApplianceCertificateSettingAllOf(varApplianceCertificateSettingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceCertificateSettingAllOf struct {
	value *ApplianceCertificateSettingAllOf
	isSet bool
}

func (v NullableApplianceCertificateSettingAllOf) Get() *ApplianceCertificateSettingAllOf {
	return v.value
}

func (v *NullableApplianceCertificateSettingAllOf) Set(val *ApplianceCertificateSettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceCertificateSettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceCertificateSettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceCertificateSettingAllOf(val *ApplianceCertificateSettingAllOf) *NullableApplianceCertificateSettingAllOf {
	return &NullableApplianceCertificateSettingAllOf{value: val, isSet: true}
}

func (v NullableApplianceCertificateSettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceCertificateSettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
