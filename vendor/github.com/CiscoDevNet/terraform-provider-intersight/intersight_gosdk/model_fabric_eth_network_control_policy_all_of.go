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

// FabricEthNetworkControlPolicyAllOf Definition of the list of properties defined in 'fabric.EthNetworkControlPolicy', excluding properties defined in parent classes.
type FabricEthNetworkControlPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables the CDP on an interface.
	CdpEnabled *bool `json:"CdpEnabled,omitempty"`
	// Determines if the MAC forging is allowed or denied on an interface. * `allow` - Allows mac forging on an interface. * `deny` - Denies mac forging on an interface.
	ForgeMac     *string                    `json:"ForgeMac,omitempty"`
	LldpSettings NullableFabricLldpSettings `json:"LldpSettings,omitempty"`
	// Determines the MAC addresses that have to be registered with the switch. * `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN. * `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs.
	MacRegistrationMode *string `json:"MacRegistrationMode,omitempty"`
	// Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * `linkDown` - The vethernet will go down in case a suitable uplink is not pinned. * `warning` - The vethernet will remain up even if a suitable uplink is not pinned.
	UplinkFailAction *string `json:"UplinkFailAction,omitempty"`
	// An array of relationships to vnicEthNetworkPolicy resources.
	NetworkPolicy        []VnicEthNetworkPolicyRelationship    `json:"NetworkPolicy,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkControlPolicyAllOf FabricEthNetworkControlPolicyAllOf

// NewFabricEthNetworkControlPolicyAllOf instantiates a new FabricEthNetworkControlPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkControlPolicyAllOf(classId string, objectType string) *FabricEthNetworkControlPolicyAllOf {
	this := FabricEthNetworkControlPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cdpEnabled bool = false
	this.CdpEnabled = &cdpEnabled
	var forgeMac string = "allow"
	this.ForgeMac = &forgeMac
	var macRegistrationMode string = "nativeVlanOnly"
	this.MacRegistrationMode = &macRegistrationMode
	var uplinkFailAction string = "linkDown"
	this.UplinkFailAction = &uplinkFailAction
	return &this
}

// NewFabricEthNetworkControlPolicyAllOfWithDefaults instantiates a new FabricEthNetworkControlPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkControlPolicyAllOfWithDefaults() *FabricEthNetworkControlPolicyAllOf {
	this := FabricEthNetworkControlPolicyAllOf{}
	var classId string = "fabric.EthNetworkControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.EthNetworkControlPolicy"
	this.ObjectType = objectType
	var cdpEnabled bool = false
	this.CdpEnabled = &cdpEnabled
	var forgeMac string = "allow"
	this.ForgeMac = &forgeMac
	var macRegistrationMode string = "nativeVlanOnly"
	this.MacRegistrationMode = &macRegistrationMode
	var uplinkFailAction string = "linkDown"
	this.UplinkFailAction = &uplinkFailAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEthNetworkControlPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEthNetworkControlPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricEthNetworkControlPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEthNetworkControlPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCdpEnabled returns the CdpEnabled field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabled() bool {
	if o == nil || o.CdpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CdpEnabled
}

// GetCdpEnabledOk returns a tuple with the CdpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabledOk() (*bool, bool) {
	if o == nil || o.CdpEnabled == nil {
		return nil, false
	}
	return o.CdpEnabled, true
}

// HasCdpEnabled returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasCdpEnabled() bool {
	if o != nil && o.CdpEnabled != nil {
		return true
	}

	return false
}

// SetCdpEnabled gets a reference to the given bool and assigns it to the CdpEnabled field.
func (o *FabricEthNetworkControlPolicyAllOf) SetCdpEnabled(v bool) {
	o.CdpEnabled = &v
}

// GetForgeMac returns the ForgeMac field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMac() string {
	if o == nil || o.ForgeMac == nil {
		var ret string
		return ret
	}
	return *o.ForgeMac
}

// GetForgeMacOk returns a tuple with the ForgeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMacOk() (*string, bool) {
	if o == nil || o.ForgeMac == nil {
		return nil, false
	}
	return o.ForgeMac, true
}

// HasForgeMac returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasForgeMac() bool {
	if o != nil && o.ForgeMac != nil {
		return true
	}

	return false
}

// SetForgeMac gets a reference to the given string and assigns it to the ForgeMac field.
func (o *FabricEthNetworkControlPolicyAllOf) SetForgeMac(v string) {
	o.ForgeMac = &v
}

// GetLldpSettings returns the LldpSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettings() FabricLldpSettings {
	if o == nil || o.LldpSettings.Get() == nil {
		var ret FabricLldpSettings
		return ret
	}
	return *o.LldpSettings.Get()
}

// GetLldpSettingsOk returns a tuple with the LldpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettingsOk() (*FabricLldpSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.LldpSettings.Get(), o.LldpSettings.IsSet()
}

// HasLldpSettings returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasLldpSettings() bool {
	if o != nil && o.LldpSettings.IsSet() {
		return true
	}

	return false
}

// SetLldpSettings gets a reference to the given NullableFabricLldpSettings and assigns it to the LldpSettings field.
func (o *FabricEthNetworkControlPolicyAllOf) SetLldpSettings(v FabricLldpSettings) {
	o.LldpSettings.Set(&v)
}

// SetLldpSettingsNil sets the value for LldpSettings to be an explicit nil
func (o *FabricEthNetworkControlPolicyAllOf) SetLldpSettingsNil() {
	o.LldpSettings.Set(nil)
}

// UnsetLldpSettings ensures that no value is present for LldpSettings, not even an explicit nil
func (o *FabricEthNetworkControlPolicyAllOf) UnsetLldpSettings() {
	o.LldpSettings.Unset()
}

// GetMacRegistrationMode returns the MacRegistrationMode field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegistrationMode() string {
	if o == nil || o.MacRegistrationMode == nil {
		var ret string
		return ret
	}
	return *o.MacRegistrationMode
}

// GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegistrationModeOk() (*string, bool) {
	if o == nil || o.MacRegistrationMode == nil {
		return nil, false
	}
	return o.MacRegistrationMode, true
}

// HasMacRegistrationMode returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasMacRegistrationMode() bool {
	if o != nil && o.MacRegistrationMode != nil {
		return true
	}

	return false
}

// SetMacRegistrationMode gets a reference to the given string and assigns it to the MacRegistrationMode field.
func (o *FabricEthNetworkControlPolicyAllOf) SetMacRegistrationMode(v string) {
	o.MacRegistrationMode = &v
}

// GetUplinkFailAction returns the UplinkFailAction field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailAction() string {
	if o == nil || o.UplinkFailAction == nil {
		var ret string
		return ret
	}
	return *o.UplinkFailAction
}

// GetUplinkFailActionOk returns a tuple with the UplinkFailAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailActionOk() (*string, bool) {
	if o == nil || o.UplinkFailAction == nil {
		return nil, false
	}
	return o.UplinkFailAction, true
}

// HasUplinkFailAction returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasUplinkFailAction() bool {
	if o != nil && o.UplinkFailAction != nil {
		return true
	}

	return false
}

// SetUplinkFailAction gets a reference to the given string and assigns it to the UplinkFailAction field.
func (o *FabricEthNetworkControlPolicyAllOf) SetUplinkFailAction(v string) {
	o.UplinkFailAction = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicy() []VnicEthNetworkPolicyRelationship {
	if o == nil {
		var ret []VnicEthNetworkPolicyRelationship
		return ret
	}
	return o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyRelationship, bool) {
	if o == nil || o.NetworkPolicy == nil {
		return nil, false
	}
	return &o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasNetworkPolicy() bool {
	if o != nil && o.NetworkPolicy != nil {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given []VnicEthNetworkPolicyRelationship and assigns it to the NetworkPolicy field.
func (o *FabricEthNetworkControlPolicyAllOf) SetNetworkPolicy(v []VnicEthNetworkPolicyRelationship) {
	o.NetworkPolicy = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricEthNetworkControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricEthNetworkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CdpEnabled != nil {
		toSerialize["CdpEnabled"] = o.CdpEnabled
	}
	if o.ForgeMac != nil {
		toSerialize["ForgeMac"] = o.ForgeMac
	}
	if o.LldpSettings.IsSet() {
		toSerialize["LldpSettings"] = o.LldpSettings.Get()
	}
	if o.MacRegistrationMode != nil {
		toSerialize["MacRegistrationMode"] = o.MacRegistrationMode
	}
	if o.UplinkFailAction != nil {
		toSerialize["UplinkFailAction"] = o.UplinkFailAction
	}
	if o.NetworkPolicy != nil {
		toSerialize["NetworkPolicy"] = o.NetworkPolicy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEthNetworkControlPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricEthNetworkControlPolicyAllOf := _FabricEthNetworkControlPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricEthNetworkControlPolicyAllOf); err == nil {
		*o = FabricEthNetworkControlPolicyAllOf(varFabricEthNetworkControlPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CdpEnabled")
		delete(additionalProperties, "ForgeMac")
		delete(additionalProperties, "LldpSettings")
		delete(additionalProperties, "MacRegistrationMode")
		delete(additionalProperties, "UplinkFailAction")
		delete(additionalProperties, "NetworkPolicy")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricEthNetworkControlPolicyAllOf struct {
	value *FabricEthNetworkControlPolicyAllOf
	isSet bool
}

func (v NullableFabricEthNetworkControlPolicyAllOf) Get() *FabricEthNetworkControlPolicyAllOf {
	return v.value
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) Set(val *FabricEthNetworkControlPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkControlPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkControlPolicyAllOf(val *FabricEthNetworkControlPolicyAllOf) *NullableFabricEthNetworkControlPolicyAllOf {
	return &NullableFabricEthNetworkControlPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricEthNetworkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
