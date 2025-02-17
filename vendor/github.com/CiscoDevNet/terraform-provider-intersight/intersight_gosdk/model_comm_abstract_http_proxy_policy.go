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

// CommAbstractHttpProxyPolicy An Abstract policy specifying the HTTP proxy settings.
type CommAbstractHttpProxyPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// HTTP Proxy server FQDN or IP.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for the HTTP Proxy.
	Password *string `json:"Password,omitempty"`
	// The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive.
	Port *int64 `json:"Port,omitempty"`
	// The username for the HTTP Proxy.
	Username             *string                               `json:"Username,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommAbstractHttpProxyPolicy CommAbstractHttpProxyPolicy

// NewCommAbstractHttpProxyPolicy instantiates a new CommAbstractHttpProxyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommAbstractHttpProxyPolicy(classId string, objectType string) *CommAbstractHttpProxyPolicy {
	this := CommAbstractHttpProxyPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewCommAbstractHttpProxyPolicyWithDefaults instantiates a new CommAbstractHttpProxyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommAbstractHttpProxyPolicyWithDefaults() *CommAbstractHttpProxyPolicy {
	this := CommAbstractHttpProxyPolicy{}
	var classId string = "comm.HttpProxyPolicy"
	this.ClassId = classId
	var objectType string = "comm.HttpProxyPolicy"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommAbstractHttpProxyPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommAbstractHttpProxyPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommAbstractHttpProxyPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommAbstractHttpProxyPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CommAbstractHttpProxyPolicy) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *CommAbstractHttpProxyPolicy) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CommAbstractHttpProxyPolicy) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *CommAbstractHttpProxyPolicy) SetPort(v int64) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CommAbstractHttpProxyPolicy) SetUsername(v string) {
	o.Username = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *CommAbstractHttpProxyPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o CommAbstractHttpProxyPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommAbstractHttpProxyPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type CommAbstractHttpProxyPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// HTTP Proxy server FQDN or IP.
		Hostname *string `json:"Hostname,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password for the HTTP Proxy.
		Password *string `json:"Password,omitempty"`
		// The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive.
		Port *int64 `json:"Port,omitempty"`
		// The username for the HTTP Proxy.
		Username     *string                               `json:"Username,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct := CommAbstractHttpProxyPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct)
	if err == nil {
		varCommAbstractHttpProxyPolicy := _CommAbstractHttpProxyPolicy{}
		varCommAbstractHttpProxyPolicy.ClassId = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.ClassId
		varCommAbstractHttpProxyPolicy.ObjectType = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.ObjectType
		varCommAbstractHttpProxyPolicy.Hostname = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.Hostname
		varCommAbstractHttpProxyPolicy.IsPasswordSet = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.IsPasswordSet
		varCommAbstractHttpProxyPolicy.Password = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.Password
		varCommAbstractHttpProxyPolicy.Port = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.Port
		varCommAbstractHttpProxyPolicy.Username = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.Username
		varCommAbstractHttpProxyPolicy.Organization = varCommAbstractHttpProxyPolicyWithoutEmbeddedStruct.Organization
		*o = CommAbstractHttpProxyPolicy(varCommAbstractHttpProxyPolicy)
	} else {
		return err
	}

	varCommAbstractHttpProxyPolicy := _CommAbstractHttpProxyPolicy{}

	err = json.Unmarshal(bytes, &varCommAbstractHttpProxyPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varCommAbstractHttpProxyPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableCommAbstractHttpProxyPolicy struct {
	value *CommAbstractHttpProxyPolicy
	isSet bool
}

func (v NullableCommAbstractHttpProxyPolicy) Get() *CommAbstractHttpProxyPolicy {
	return v.value
}

func (v *NullableCommAbstractHttpProxyPolicy) Set(val *CommAbstractHttpProxyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCommAbstractHttpProxyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCommAbstractHttpProxyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommAbstractHttpProxyPolicy(val *CommAbstractHttpProxyPolicy) *NullableCommAbstractHttpProxyPolicy {
	return &NullableCommAbstractHttpProxyPolicy{value: val, isSet: true}
}

func (v NullableCommAbstractHttpProxyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommAbstractHttpProxyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
