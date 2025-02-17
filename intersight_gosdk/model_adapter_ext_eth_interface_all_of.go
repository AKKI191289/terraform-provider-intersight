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

// AdapterExtEthInterfaceAllOf Definition of the list of properties defined in 'adapter.ExtEthInterface', excluding properties defined in parent classes.
type AdapterExtEthInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured state of an External Ethernet Interface.
	AdminState *string `json:"AdminState,omitempty"`
	// Endpoint Config DN of an External Ethernet Interface.
	EpDn *string `json:"EpDn,omitempty"`
	// Unique Identifier for an External Ethernet Interface within the adapter object.
	ExtEthInterfaceId *string `json:"ExtEthInterfaceId,omitempty"`
	// Type of an External Ethernet Interface.
	InterfaceType *string `json:"InterfaceType,omitempty"`
	// MAC address of an External Ethernet Interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Peer Aggregate Port Id attached to an External Ethernet Interface.
	PeerAggrPortId *int64 `json:"PeerAggrPortId,omitempty"`
	// DN of peer end-point attached to an External Ethernet Interface.
	PeerDn *string `json:"PeerDn,omitempty"`
	// Peer Port Id attached to an External Ethernet Interface.
	PeerPortId *int64 `json:"PeerPortId,omitempty"`
	// Peer Slot Id attached to an External Ethernet Interface.
	PeerSlotId *int64 `json:"PeerSlotId,omitempty"`
	// SwitchId attached to an External Ethernet Interface.
	SwitchId             *string                              `json:"SwitchId,omitempty"`
	AdapterUnit          *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterExtEthInterfaceAllOf AdapterExtEthInterfaceAllOf

// NewAdapterExtEthInterfaceAllOf instantiates a new AdapterExtEthInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterExtEthInterfaceAllOf(classId string, objectType string) *AdapterExtEthInterfaceAllOf {
	this := AdapterExtEthInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterExtEthInterfaceAllOfWithDefaults instantiates a new AdapterExtEthInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterExtEthInterfaceAllOfWithDefaults() *AdapterExtEthInterfaceAllOf {
	this := AdapterExtEthInterfaceAllOf{}
	var classId string = "adapter.ExtEthInterface"
	this.ClassId = classId
	var objectType string = "adapter.ExtEthInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterExtEthInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterExtEthInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterExtEthInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterExtEthInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *AdapterExtEthInterfaceAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetEpDn returns the EpDn field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetEpDn() string {
	if o == nil || o.EpDn == nil {
		var ret string
		return ret
	}
	return *o.EpDn
}

// GetEpDnOk returns a tuple with the EpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetEpDnOk() (*string, bool) {
	if o == nil || o.EpDn == nil {
		return nil, false
	}
	return o.EpDn, true
}

// HasEpDn returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasEpDn() bool {
	if o != nil && o.EpDn != nil {
		return true
	}

	return false
}

// SetEpDn gets a reference to the given string and assigns it to the EpDn field.
func (o *AdapterExtEthInterfaceAllOf) SetEpDn(v string) {
	o.EpDn = &v
}

// GetExtEthInterfaceId returns the ExtEthInterfaceId field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetExtEthInterfaceId() string {
	if o == nil || o.ExtEthInterfaceId == nil {
		var ret string
		return ret
	}
	return *o.ExtEthInterfaceId
}

// GetExtEthInterfaceIdOk returns a tuple with the ExtEthInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetExtEthInterfaceIdOk() (*string, bool) {
	if o == nil || o.ExtEthInterfaceId == nil {
		return nil, false
	}
	return o.ExtEthInterfaceId, true
}

// HasExtEthInterfaceId returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasExtEthInterfaceId() bool {
	if o != nil && o.ExtEthInterfaceId != nil {
		return true
	}

	return false
}

// SetExtEthInterfaceId gets a reference to the given string and assigns it to the ExtEthInterfaceId field.
func (o *AdapterExtEthInterfaceAllOf) SetExtEthInterfaceId(v string) {
	o.ExtEthInterfaceId = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *AdapterExtEthInterfaceAllOf) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *AdapterExtEthInterfaceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPeerAggrPortId returns the PeerAggrPortId field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetPeerAggrPortId() int64 {
	if o == nil || o.PeerAggrPortId == nil {
		var ret int64
		return ret
	}
	return *o.PeerAggrPortId
}

// GetPeerAggrPortIdOk returns a tuple with the PeerAggrPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetPeerAggrPortIdOk() (*int64, bool) {
	if o == nil || o.PeerAggrPortId == nil {
		return nil, false
	}
	return o.PeerAggrPortId, true
}

// HasPeerAggrPortId returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasPeerAggrPortId() bool {
	if o != nil && o.PeerAggrPortId != nil {
		return true
	}

	return false
}

// SetPeerAggrPortId gets a reference to the given int64 and assigns it to the PeerAggrPortId field.
func (o *AdapterExtEthInterfaceAllOf) SetPeerAggrPortId(v int64) {
	o.PeerAggrPortId = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *AdapterExtEthInterfaceAllOf) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetPeerPortId returns the PeerPortId field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetPeerPortId() int64 {
	if o == nil || o.PeerPortId == nil {
		var ret int64
		return ret
	}
	return *o.PeerPortId
}

// GetPeerPortIdOk returns a tuple with the PeerPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetPeerPortIdOk() (*int64, bool) {
	if o == nil || o.PeerPortId == nil {
		return nil, false
	}
	return o.PeerPortId, true
}

// HasPeerPortId returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasPeerPortId() bool {
	if o != nil && o.PeerPortId != nil {
		return true
	}

	return false
}

// SetPeerPortId gets a reference to the given int64 and assigns it to the PeerPortId field.
func (o *AdapterExtEthInterfaceAllOf) SetPeerPortId(v int64) {
	o.PeerPortId = &v
}

// GetPeerSlotId returns the PeerSlotId field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetPeerSlotId() int64 {
	if o == nil || o.PeerSlotId == nil {
		var ret int64
		return ret
	}
	return *o.PeerSlotId
}

// GetPeerSlotIdOk returns a tuple with the PeerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetPeerSlotIdOk() (*int64, bool) {
	if o == nil || o.PeerSlotId == nil {
		return nil, false
	}
	return o.PeerSlotId, true
}

// HasPeerSlotId returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasPeerSlotId() bool {
	if o != nil && o.PeerSlotId != nil {
		return true
	}

	return false
}

// SetPeerSlotId gets a reference to the given int64 and assigns it to the PeerSlotId field.
func (o *AdapterExtEthInterfaceAllOf) SetPeerSlotId(v int64) {
	o.PeerSlotId = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *AdapterExtEthInterfaceAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *AdapterExtEthInterfaceAllOf) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterExtEthInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterExtEthInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExtEthInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterExtEthInterfaceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterExtEthInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterExtEthInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.EpDn != nil {
		toSerialize["EpDn"] = o.EpDn
	}
	if o.ExtEthInterfaceId != nil {
		toSerialize["ExtEthInterfaceId"] = o.ExtEthInterfaceId
	}
	if o.InterfaceType != nil {
		toSerialize["InterfaceType"] = o.InterfaceType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.PeerAggrPortId != nil {
		toSerialize["PeerAggrPortId"] = o.PeerAggrPortId
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.PeerPortId != nil {
		toSerialize["PeerPortId"] = o.PeerPortId
	}
	if o.PeerSlotId != nil {
		toSerialize["PeerSlotId"] = o.PeerSlotId
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterExtEthInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterExtEthInterfaceAllOf := _AdapterExtEthInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterExtEthInterfaceAllOf); err == nil {
		*o = AdapterExtEthInterfaceAllOf(varAdapterExtEthInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "EpDn")
		delete(additionalProperties, "ExtEthInterfaceId")
		delete(additionalProperties, "InterfaceType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "PeerAggrPortId")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "PeerPortId")
		delete(additionalProperties, "PeerSlotId")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterExtEthInterfaceAllOf struct {
	value *AdapterExtEthInterfaceAllOf
	isSet bool
}

func (v NullableAdapterExtEthInterfaceAllOf) Get() *AdapterExtEthInterfaceAllOf {
	return v.value
}

func (v *NullableAdapterExtEthInterfaceAllOf) Set(val *AdapterExtEthInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterExtEthInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterExtEthInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterExtEthInterfaceAllOf(val *AdapterExtEthInterfaceAllOf) *NullableAdapterExtEthInterfaceAllOf {
	return &NullableAdapterExtEthInterfaceAllOf{value: val, isSet: true}
}

func (v NullableAdapterExtEthInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterExtEthInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
