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

// SnmpPolicyAllOf Definition of the list of properties defined in 'snmp.Policy', excluding properties defined in parent classes.
type SnmpPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.
	AccessCommunityString *string `json:"AccessCommunityString,omitempty"`
	// Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users. * `Disabled` - Blocks access to the information in the inventory tables. * `Limited` - Partial access to read the information in the inventory tables. * `Full` - Full access to read the information in the inventory tables.
	CommunityAccess *string `json:"CommunityAccess,omitempty"`
	// State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host.
	Enabled *bool `json:"Enabled,omitempty"`
	// User-defined unique identification of the static engine.
	EngineId *string `json:"EngineId,omitempty"`
	// Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).
	SnmpPort  *int64     `json:"SnmpPort,omitempty"`
	SnmpTraps []SnmpTrap `json:"SnmpTraps,omitempty"`
	SnmpUsers []SnmpUser `json:"SnmpUsers,omitempty"`
	// Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number.
	SysContact *string `json:"SysContact,omitempty"`
	// Location of host on which the SNMP agent (server) runs.
	SysLocation *string `json:"SysLocation,omitempty"`
	// SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.
	TrapCommunity *string                               `json:"TrapCommunity,omitempty"`
	Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpPolicyAllOf SnmpPolicyAllOf

// NewSnmpPolicyAllOf instantiates a new SnmpPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpPolicyAllOf(classId string, objectType string) *SnmpPolicyAllOf {
	this := SnmpPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var communityAccess string = "Disabled"
	this.CommunityAccess = &communityAccess
	var enabled bool = true
	this.Enabled = &enabled
	var snmpPort int64 = 161
	this.SnmpPort = &snmpPort
	return &this
}

// NewSnmpPolicyAllOfWithDefaults instantiates a new SnmpPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpPolicyAllOfWithDefaults() *SnmpPolicyAllOf {
	this := SnmpPolicyAllOf{}
	var classId string = "snmp.Policy"
	this.ClassId = classId
	var objectType string = "snmp.Policy"
	this.ObjectType = objectType
	var communityAccess string = "Disabled"
	this.CommunityAccess = &communityAccess
	var enabled bool = true
	this.Enabled = &enabled
	var snmpPort int64 = 161
	this.SnmpPort = &snmpPort
	return &this
}

// GetClassId returns the ClassId field value
func (o *SnmpPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SnmpPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SnmpPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SnmpPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessCommunityString returns the AccessCommunityString field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetAccessCommunityString() string {
	if o == nil || o.AccessCommunityString == nil {
		var ret string
		return ret
	}
	return *o.AccessCommunityString
}

// GetAccessCommunityStringOk returns a tuple with the AccessCommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetAccessCommunityStringOk() (*string, bool) {
	if o == nil || o.AccessCommunityString == nil {
		return nil, false
	}
	return o.AccessCommunityString, true
}

// HasAccessCommunityString returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasAccessCommunityString() bool {
	if o != nil && o.AccessCommunityString != nil {
		return true
	}

	return false
}

// SetAccessCommunityString gets a reference to the given string and assigns it to the AccessCommunityString field.
func (o *SnmpPolicyAllOf) SetAccessCommunityString(v string) {
	o.AccessCommunityString = &v
}

// GetCommunityAccess returns the CommunityAccess field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetCommunityAccess() string {
	if o == nil || o.CommunityAccess == nil {
		var ret string
		return ret
	}
	return *o.CommunityAccess
}

// GetCommunityAccessOk returns a tuple with the CommunityAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetCommunityAccessOk() (*string, bool) {
	if o == nil || o.CommunityAccess == nil {
		return nil, false
	}
	return o.CommunityAccess, true
}

// HasCommunityAccess returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasCommunityAccess() bool {
	if o != nil && o.CommunityAccess != nil {
		return true
	}

	return false
}

// SetCommunityAccess gets a reference to the given string and assigns it to the CommunityAccess field.
func (o *SnmpPolicyAllOf) SetCommunityAccess(v string) {
	o.CommunityAccess = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SnmpPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetEngineId() string {
	if o == nil || o.EngineId == nil {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetEngineIdOk() (*string, bool) {
	if o == nil || o.EngineId == nil {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasEngineId() bool {
	if o != nil && o.EngineId != nil {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *SnmpPolicyAllOf) SetEngineId(v string) {
	o.EngineId = &v
}

// GetSnmpPort returns the SnmpPort field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetSnmpPort() int64 {
	if o == nil || o.SnmpPort == nil {
		var ret int64
		return ret
	}
	return *o.SnmpPort
}

// GetSnmpPortOk returns a tuple with the SnmpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetSnmpPortOk() (*int64, bool) {
	if o == nil || o.SnmpPort == nil {
		return nil, false
	}
	return o.SnmpPort, true
}

// HasSnmpPort returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasSnmpPort() bool {
	if o != nil && o.SnmpPort != nil {
		return true
	}

	return false
}

// SetSnmpPort gets a reference to the given int64 and assigns it to the SnmpPort field.
func (o *SnmpPolicyAllOf) SetSnmpPort(v int64) {
	o.SnmpPort = &v
}

// GetSnmpTraps returns the SnmpTraps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpPolicyAllOf) GetSnmpTraps() []SnmpTrap {
	if o == nil {
		var ret []SnmpTrap
		return ret
	}
	return o.SnmpTraps
}

// GetSnmpTrapsOk returns a tuple with the SnmpTraps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpPolicyAllOf) GetSnmpTrapsOk() (*[]SnmpTrap, bool) {
	if o == nil || o.SnmpTraps == nil {
		return nil, false
	}
	return &o.SnmpTraps, true
}

// HasSnmpTraps returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasSnmpTraps() bool {
	if o != nil && o.SnmpTraps != nil {
		return true
	}

	return false
}

// SetSnmpTraps gets a reference to the given []SnmpTrap and assigns it to the SnmpTraps field.
func (o *SnmpPolicyAllOf) SetSnmpTraps(v []SnmpTrap) {
	o.SnmpTraps = v
}

// GetSnmpUsers returns the SnmpUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpPolicyAllOf) GetSnmpUsers() []SnmpUser {
	if o == nil {
		var ret []SnmpUser
		return ret
	}
	return o.SnmpUsers
}

// GetSnmpUsersOk returns a tuple with the SnmpUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpPolicyAllOf) GetSnmpUsersOk() (*[]SnmpUser, bool) {
	if o == nil || o.SnmpUsers == nil {
		return nil, false
	}
	return &o.SnmpUsers, true
}

// HasSnmpUsers returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasSnmpUsers() bool {
	if o != nil && o.SnmpUsers != nil {
		return true
	}

	return false
}

// SetSnmpUsers gets a reference to the given []SnmpUser and assigns it to the SnmpUsers field.
func (o *SnmpPolicyAllOf) SetSnmpUsers(v []SnmpUser) {
	o.SnmpUsers = v
}

// GetSysContact returns the SysContact field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetSysContact() string {
	if o == nil || o.SysContact == nil {
		var ret string
		return ret
	}
	return *o.SysContact
}

// GetSysContactOk returns a tuple with the SysContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetSysContactOk() (*string, bool) {
	if o == nil || o.SysContact == nil {
		return nil, false
	}
	return o.SysContact, true
}

// HasSysContact returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasSysContact() bool {
	if o != nil && o.SysContact != nil {
		return true
	}

	return false
}

// SetSysContact gets a reference to the given string and assigns it to the SysContact field.
func (o *SnmpPolicyAllOf) SetSysContact(v string) {
	o.SysContact = &v
}

// GetSysLocation returns the SysLocation field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetSysLocation() string {
	if o == nil || o.SysLocation == nil {
		var ret string
		return ret
	}
	return *o.SysLocation
}

// GetSysLocationOk returns a tuple with the SysLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetSysLocationOk() (*string, bool) {
	if o == nil || o.SysLocation == nil {
		return nil, false
	}
	return o.SysLocation, true
}

// HasSysLocation returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasSysLocation() bool {
	if o != nil && o.SysLocation != nil {
		return true
	}

	return false
}

// SetSysLocation gets a reference to the given string and assigns it to the SysLocation field.
func (o *SnmpPolicyAllOf) SetSysLocation(v string) {
	o.SysLocation = &v
}

// GetTrapCommunity returns the TrapCommunity field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetTrapCommunity() string {
	if o == nil || o.TrapCommunity == nil {
		var ret string
		return ret
	}
	return *o.TrapCommunity
}

// GetTrapCommunityOk returns a tuple with the TrapCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetTrapCommunityOk() (*string, bool) {
	if o == nil || o.TrapCommunity == nil {
		return nil, false
	}
	return o.TrapCommunity, true
}

// HasTrapCommunity returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasTrapCommunity() bool {
	if o != nil && o.TrapCommunity != nil {
		return true
	}

	return false
}

// SetTrapCommunity gets a reference to the given string and assigns it to the TrapCommunity field.
func (o *SnmpPolicyAllOf) SetTrapCommunity(v string) {
	o.TrapCommunity = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SnmpPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SnmpPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SnmpPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SnmpPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SnmpPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessCommunityString != nil {
		toSerialize["AccessCommunityString"] = o.AccessCommunityString
	}
	if o.CommunityAccess != nil {
		toSerialize["CommunityAccess"] = o.CommunityAccess
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.EngineId != nil {
		toSerialize["EngineId"] = o.EngineId
	}
	if o.SnmpPort != nil {
		toSerialize["SnmpPort"] = o.SnmpPort
	}
	if o.SnmpTraps != nil {
		toSerialize["SnmpTraps"] = o.SnmpTraps
	}
	if o.SnmpUsers != nil {
		toSerialize["SnmpUsers"] = o.SnmpUsers
	}
	if o.SysContact != nil {
		toSerialize["SysContact"] = o.SysContact
	}
	if o.SysLocation != nil {
		toSerialize["SysLocation"] = o.SysLocation
	}
	if o.TrapCommunity != nil {
		toSerialize["TrapCommunity"] = o.TrapCommunity
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

func (o *SnmpPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnmpPolicyAllOf := _SnmpPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varSnmpPolicyAllOf); err == nil {
		*o = SnmpPolicyAllOf(varSnmpPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessCommunityString")
		delete(additionalProperties, "CommunityAccess")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EngineId")
		delete(additionalProperties, "SnmpPort")
		delete(additionalProperties, "SnmpTraps")
		delete(additionalProperties, "SnmpUsers")
		delete(additionalProperties, "SysContact")
		delete(additionalProperties, "SysLocation")
		delete(additionalProperties, "TrapCommunity")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpPolicyAllOf struct {
	value *SnmpPolicyAllOf
	isSet bool
}

func (v NullableSnmpPolicyAllOf) Get() *SnmpPolicyAllOf {
	return v.value
}

func (v *NullableSnmpPolicyAllOf) Set(val *SnmpPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpPolicyAllOf(val *SnmpPolicyAllOf) *NullableSnmpPolicyAllOf {
	return &NullableSnmpPolicyAllOf{value: val, isSet: true}
}

func (v NullableSnmpPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
