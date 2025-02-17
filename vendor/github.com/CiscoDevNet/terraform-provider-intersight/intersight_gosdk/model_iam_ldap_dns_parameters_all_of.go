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

// IamLdapDnsParametersAllOf Definition of the list of properties defined in 'iam.LdapDnsParameters', excluding properties defined in parent classes.
type IamLdapDnsParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Domain name that acts as a source for a DNS query.
	SearchDomain *string `json:"SearchDomain,omitempty"`
	// Forest name that acts as a source for a DNS query.
	SearchForest *string `json:"SearchForest,omitempty"`
	// Source of the domain name used for the DNS SRV request. * `Extracted` - The domain name extracted-domain from the login ID. * `Configured` - The configured-search domain. * `ConfiguredExtracted` - The domain name extracted from the login ID than the configured-search domain.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLdapDnsParametersAllOf IamLdapDnsParametersAllOf

// NewIamLdapDnsParametersAllOf instantiates a new IamLdapDnsParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapDnsParametersAllOf(classId string, objectType string) *IamLdapDnsParametersAllOf {
	this := IamLdapDnsParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var source string = "Extracted"
	this.Source = &source
	return &this
}

// NewIamLdapDnsParametersAllOfWithDefaults instantiates a new IamLdapDnsParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapDnsParametersAllOfWithDefaults() *IamLdapDnsParametersAllOf {
	this := IamLdapDnsParametersAllOf{}
	var classId string = "iam.LdapDnsParameters"
	this.ClassId = classId
	var objectType string = "iam.LdapDnsParameters"
	this.ObjectType = objectType
	var source string = "Extracted"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLdapDnsParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLdapDnsParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLdapDnsParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLdapDnsParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLdapDnsParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLdapDnsParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSearchDomain returns the SearchDomain field value if set, zero value otherwise.
func (o *IamLdapDnsParametersAllOf) GetSearchDomain() string {
	if o == nil || o.SearchDomain == nil {
		var ret string
		return ret
	}
	return *o.SearchDomain
}

// GetSearchDomainOk returns a tuple with the SearchDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapDnsParametersAllOf) GetSearchDomainOk() (*string, bool) {
	if o == nil || o.SearchDomain == nil {
		return nil, false
	}
	return o.SearchDomain, true
}

// HasSearchDomain returns a boolean if a field has been set.
func (o *IamLdapDnsParametersAllOf) HasSearchDomain() bool {
	if o != nil && o.SearchDomain != nil {
		return true
	}

	return false
}

// SetSearchDomain gets a reference to the given string and assigns it to the SearchDomain field.
func (o *IamLdapDnsParametersAllOf) SetSearchDomain(v string) {
	o.SearchDomain = &v
}

// GetSearchForest returns the SearchForest field value if set, zero value otherwise.
func (o *IamLdapDnsParametersAllOf) GetSearchForest() string {
	if o == nil || o.SearchForest == nil {
		var ret string
		return ret
	}
	return *o.SearchForest
}

// GetSearchForestOk returns a tuple with the SearchForest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapDnsParametersAllOf) GetSearchForestOk() (*string, bool) {
	if o == nil || o.SearchForest == nil {
		return nil, false
	}
	return o.SearchForest, true
}

// HasSearchForest returns a boolean if a field has been set.
func (o *IamLdapDnsParametersAllOf) HasSearchForest() bool {
	if o != nil && o.SearchForest != nil {
		return true
	}

	return false
}

// SetSearchForest gets a reference to the given string and assigns it to the SearchForest field.
func (o *IamLdapDnsParametersAllOf) SetSearchForest(v string) {
	o.SearchForest = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IamLdapDnsParametersAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapDnsParametersAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IamLdapDnsParametersAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *IamLdapDnsParametersAllOf) SetSource(v string) {
	o.Source = &v
}

func (o IamLdapDnsParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SearchDomain != nil {
		toSerialize["SearchDomain"] = o.SearchDomain
	}
	if o.SearchForest != nil {
		toSerialize["SearchForest"] = o.SearchForest
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLdapDnsParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamLdapDnsParametersAllOf := _IamLdapDnsParametersAllOf{}

	if err = json.Unmarshal(bytes, &varIamLdapDnsParametersAllOf); err == nil {
		*o = IamLdapDnsParametersAllOf(varIamLdapDnsParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SearchDomain")
		delete(additionalProperties, "SearchForest")
		delete(additionalProperties, "Source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamLdapDnsParametersAllOf struct {
	value *IamLdapDnsParametersAllOf
	isSet bool
}

func (v NullableIamLdapDnsParametersAllOf) Get() *IamLdapDnsParametersAllOf {
	return v.value
}

func (v *NullableIamLdapDnsParametersAllOf) Set(val *IamLdapDnsParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapDnsParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapDnsParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapDnsParametersAllOf(val *IamLdapDnsParametersAllOf) *NullableIamLdapDnsParametersAllOf {
	return &NullableIamLdapDnsParametersAllOf{value: val, isSet: true}
}

func (v NullableIamLdapDnsParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapDnsParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
