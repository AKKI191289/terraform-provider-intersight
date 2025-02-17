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

// KvmPolicy Policy to configure KVM Launch settings.
type KvmPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, displays KVM session on any monitor attached to the server.
	EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
	// If enabled, encrypts all video information sent through KVM.
	EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
	// State of the vKVM service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The maximum number of concurrent KVM sessions allowed.
	MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
	// The port used for KVM communication.
	RemotePort   *int64                                `json:"RemotePort,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmPolicy KvmPolicy

// NewKvmPolicy instantiates a new KvmPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmPolicy(classId string, objectType string) *KvmPolicy {
	this := KvmPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	return &this
}

// NewKvmPolicyWithDefaults instantiates a new KvmPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmPolicyWithDefaults() *KvmPolicy {
	this := KvmPolicy{}
	var classId string = "kvm.Policy"
	this.ClassId = classId
	var objectType string = "kvm.Policy"
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableLocalServerVideo returns the EnableLocalServerVideo field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnableLocalServerVideo() bool {
	if o == nil || o.EnableLocalServerVideo == nil {
		var ret bool
		return ret
	}
	return *o.EnableLocalServerVideo
}

// GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnableLocalServerVideoOk() (*bool, bool) {
	if o == nil || o.EnableLocalServerVideo == nil {
		return nil, false
	}
	return o.EnableLocalServerVideo, true
}

// HasEnableLocalServerVideo returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnableLocalServerVideo() bool {
	if o != nil && o.EnableLocalServerVideo != nil {
		return true
	}

	return false
}

// SetEnableLocalServerVideo gets a reference to the given bool and assigns it to the EnableLocalServerVideo field.
func (o *KvmPolicy) SetEnableLocalServerVideo(v bool) {
	o.EnableLocalServerVideo = &v
}

// GetEnableVideoEncryption returns the EnableVideoEncryption field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnableVideoEncryption() bool {
	if o == nil || o.EnableVideoEncryption == nil {
		var ret bool
		return ret
	}
	return *o.EnableVideoEncryption
}

// GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnableVideoEncryptionOk() (*bool, bool) {
	if o == nil || o.EnableVideoEncryption == nil {
		return nil, false
	}
	return o.EnableVideoEncryption, true
}

// HasEnableVideoEncryption returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnableVideoEncryption() bool {
	if o != nil && o.EnableVideoEncryption != nil {
		return true
	}

	return false
}

// SetEnableVideoEncryption gets a reference to the given bool and assigns it to the EnableVideoEncryption field.
func (o *KvmPolicy) SetEnableVideoEncryption(v bool) {
	o.EnableVideoEncryption = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KvmPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumSessions returns the MaximumSessions field value if set, zero value otherwise.
func (o *KvmPolicy) GetMaximumSessions() int64 {
	if o == nil || o.MaximumSessions == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSessions
}

// GetMaximumSessionsOk returns a tuple with the MaximumSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetMaximumSessionsOk() (*int64, bool) {
	if o == nil || o.MaximumSessions == nil {
		return nil, false
	}
	return o.MaximumSessions, true
}

// HasMaximumSessions returns a boolean if a field has been set.
func (o *KvmPolicy) HasMaximumSessions() bool {
	if o != nil && o.MaximumSessions != nil {
		return true
	}

	return false
}

// SetMaximumSessions gets a reference to the given int64 and assigns it to the MaximumSessions field.
func (o *KvmPolicy) SetMaximumSessions(v int64) {
	o.MaximumSessions = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *KvmPolicy) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *KvmPolicy) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *KvmPolicy) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KvmPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KvmPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KvmPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KvmPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *KvmPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o KvmPolicy) MarshalJSON() ([]byte, error) {
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
	if o.EnableLocalServerVideo != nil {
		toSerialize["EnableLocalServerVideo"] = o.EnableLocalServerVideo
	}
	if o.EnableVideoEncryption != nil {
		toSerialize["EnableVideoEncryption"] = o.EnableVideoEncryption
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MaximumSessions != nil {
		toSerialize["MaximumSessions"] = o.MaximumSessions
	}
	if o.RemotePort != nil {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KvmPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If enabled, displays KVM session on any monitor attached to the server.
		EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
		// If enabled, encrypts all video information sent through KVM.
		EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
		// State of the vKVM service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// The maximum number of concurrent KVM sessions allowed.
		MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
		// The port used for KVM communication.
		RemotePort   *int64                                `json:"RemotePort,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varKvmPolicyWithoutEmbeddedStruct := KvmPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKvmPolicy := _KvmPolicy{}
		varKvmPolicy.ClassId = varKvmPolicyWithoutEmbeddedStruct.ClassId
		varKvmPolicy.ObjectType = varKvmPolicyWithoutEmbeddedStruct.ObjectType
		varKvmPolicy.EnableLocalServerVideo = varKvmPolicyWithoutEmbeddedStruct.EnableLocalServerVideo
		varKvmPolicy.EnableVideoEncryption = varKvmPolicyWithoutEmbeddedStruct.EnableVideoEncryption
		varKvmPolicy.Enabled = varKvmPolicyWithoutEmbeddedStruct.Enabled
		varKvmPolicy.MaximumSessions = varKvmPolicyWithoutEmbeddedStruct.MaximumSessions
		varKvmPolicy.RemotePort = varKvmPolicyWithoutEmbeddedStruct.RemotePort
		varKvmPolicy.Organization = varKvmPolicyWithoutEmbeddedStruct.Organization
		varKvmPolicy.Profiles = varKvmPolicyWithoutEmbeddedStruct.Profiles
		*o = KvmPolicy(varKvmPolicy)
	} else {
		return err
	}

	varKvmPolicy := _KvmPolicy{}

	err = json.Unmarshal(bytes, &varKvmPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKvmPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableLocalServerVideo")
		delete(additionalProperties, "EnableVideoEncryption")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MaximumSessions")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableKvmPolicy struct {
	value *KvmPolicy
	isSet bool
}

func (v NullableKvmPolicy) Get() *KvmPolicy {
	return v.value
}

func (v *NullableKvmPolicy) Set(val *KvmPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmPolicy(val *KvmPolicy) *NullableKvmPolicy {
	return &NullableKvmPolicy{value: val, isSet: true}
}

func (v NullableKvmPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
