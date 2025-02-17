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

// ConnectorAuthMessageAllOf Definition of the list of properties defined in 'connector.AuthMessage', excluding properties defined in parent classes.
type ConnectorAuthMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The platform locale to assign user. A locale defines one or more organizations (domains) the user is allowed access, and access is limited to the organizations specified in the locale.
	RemoteUserLocale *string `json:"RemoteUserLocale,omitempty"`
	// The user name passed to the platform for use in platform audit logs.
	RemoteUserName *string `json:"RemoteUserName,omitempty"`
	// The list of roles to pass to the platform to validate the action against.
	RemoteUserRoles *string `json:"RemoteUserRoles,omitempty"`
	// The session Id passed to the platform for use in platforms auditing.
	RemoteUserSessionId  *string `json:"RemoteUserSessionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorAuthMessageAllOf ConnectorAuthMessageAllOf

// NewConnectorAuthMessageAllOf instantiates a new ConnectorAuthMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorAuthMessageAllOf(classId string, objectType string) *ConnectorAuthMessageAllOf {
	this := ConnectorAuthMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorAuthMessageAllOfWithDefaults instantiates a new ConnectorAuthMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorAuthMessageAllOfWithDefaults() *ConnectorAuthMessageAllOf {
	this := ConnectorAuthMessageAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorAuthMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorAuthMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorAuthMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorAuthMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRemoteUserLocale returns the RemoteUserLocale field value if set, zero value otherwise.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserLocale() string {
	if o == nil || o.RemoteUserLocale == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserLocale
}

// GetRemoteUserLocaleOk returns a tuple with the RemoteUserLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserLocaleOk() (*string, bool) {
	if o == nil || o.RemoteUserLocale == nil {
		return nil, false
	}
	return o.RemoteUserLocale, true
}

// HasRemoteUserLocale returns a boolean if a field has been set.
func (o *ConnectorAuthMessageAllOf) HasRemoteUserLocale() bool {
	if o != nil && o.RemoteUserLocale != nil {
		return true
	}

	return false
}

// SetRemoteUserLocale gets a reference to the given string and assigns it to the RemoteUserLocale field.
func (o *ConnectorAuthMessageAllOf) SetRemoteUserLocale(v string) {
	o.RemoteUserLocale = &v
}

// GetRemoteUserName returns the RemoteUserName field value if set, zero value otherwise.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserName() string {
	if o == nil || o.RemoteUserName == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserName
}

// GetRemoteUserNameOk returns a tuple with the RemoteUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserNameOk() (*string, bool) {
	if o == nil || o.RemoteUserName == nil {
		return nil, false
	}
	return o.RemoteUserName, true
}

// HasRemoteUserName returns a boolean if a field has been set.
func (o *ConnectorAuthMessageAllOf) HasRemoteUserName() bool {
	if o != nil && o.RemoteUserName != nil {
		return true
	}

	return false
}

// SetRemoteUserName gets a reference to the given string and assigns it to the RemoteUserName field.
func (o *ConnectorAuthMessageAllOf) SetRemoteUserName(v string) {
	o.RemoteUserName = &v
}

// GetRemoteUserRoles returns the RemoteUserRoles field value if set, zero value otherwise.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserRoles() string {
	if o == nil || o.RemoteUserRoles == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserRoles
}

// GetRemoteUserRolesOk returns a tuple with the RemoteUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserRolesOk() (*string, bool) {
	if o == nil || o.RemoteUserRoles == nil {
		return nil, false
	}
	return o.RemoteUserRoles, true
}

// HasRemoteUserRoles returns a boolean if a field has been set.
func (o *ConnectorAuthMessageAllOf) HasRemoteUserRoles() bool {
	if o != nil && o.RemoteUserRoles != nil {
		return true
	}

	return false
}

// SetRemoteUserRoles gets a reference to the given string and assigns it to the RemoteUserRoles field.
func (o *ConnectorAuthMessageAllOf) SetRemoteUserRoles(v string) {
	o.RemoteUserRoles = &v
}

// GetRemoteUserSessionId returns the RemoteUserSessionId field value if set, zero value otherwise.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserSessionId() string {
	if o == nil || o.RemoteUserSessionId == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserSessionId
}

// GetRemoteUserSessionIdOk returns a tuple with the RemoteUserSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessageAllOf) GetRemoteUserSessionIdOk() (*string, bool) {
	if o == nil || o.RemoteUserSessionId == nil {
		return nil, false
	}
	return o.RemoteUserSessionId, true
}

// HasRemoteUserSessionId returns a boolean if a field has been set.
func (o *ConnectorAuthMessageAllOf) HasRemoteUserSessionId() bool {
	if o != nil && o.RemoteUserSessionId != nil {
		return true
	}

	return false
}

// SetRemoteUserSessionId gets a reference to the given string and assigns it to the RemoteUserSessionId field.
func (o *ConnectorAuthMessageAllOf) SetRemoteUserSessionId(v string) {
	o.RemoteUserSessionId = &v
}

func (o ConnectorAuthMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RemoteUserLocale != nil {
		toSerialize["RemoteUserLocale"] = o.RemoteUserLocale
	}
	if o.RemoteUserName != nil {
		toSerialize["RemoteUserName"] = o.RemoteUserName
	}
	if o.RemoteUserRoles != nil {
		toSerialize["RemoteUserRoles"] = o.RemoteUserRoles
	}
	if o.RemoteUserSessionId != nil {
		toSerialize["RemoteUserSessionId"] = o.RemoteUserSessionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorAuthMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorAuthMessageAllOf := _ConnectorAuthMessageAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorAuthMessageAllOf); err == nil {
		*o = ConnectorAuthMessageAllOf(varConnectorAuthMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RemoteUserLocale")
		delete(additionalProperties, "RemoteUserName")
		delete(additionalProperties, "RemoteUserRoles")
		delete(additionalProperties, "RemoteUserSessionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorAuthMessageAllOf struct {
	value *ConnectorAuthMessageAllOf
	isSet bool
}

func (v NullableConnectorAuthMessageAllOf) Get() *ConnectorAuthMessageAllOf {
	return v.value
}

func (v *NullableConnectorAuthMessageAllOf) Set(val *ConnectorAuthMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorAuthMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorAuthMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorAuthMessageAllOf(val *ConnectorAuthMessageAllOf) *NullableConnectorAuthMessageAllOf {
	return &NullableConnectorAuthMessageAllOf{value: val, isSet: true}
}

func (v NullableConnectorAuthMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorAuthMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
