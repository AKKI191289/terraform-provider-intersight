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

// EquipmentIdentitySummary Consolidated view of all equipment identities.
type EquipmentIdentitySummary struct {
	ViewsView
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Serial Identifier of an adapter participating in SWM.
	AdapterSerial *string `json:"AdapterSerial,omitempty"`
	// Updated by UI/API to trigger specific chassis action type. * `None` - No operation value for maintenance actions on an equipment. * `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight. * `Recommission` - Recommission the equipment. * `Reack` - Reacknowledge the equipment and discover it again. * `Remove` - Remove the equipment permanently from Intersight management.
	AdminAction *string `json:"AdminAction,omitempty"`
	// The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	AdminActionState *string `json:"AdminActionState,omitempty"`
	// Chassis Identifier of a blade server.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// Describes whether the running CIMC version supports Intersight managed mode. * `Unknown` - The running firmware version is unknown. * `Supported` - The running firmware version is known and supports IMM mode. * `NotSupported` - The running firmware version is known and does not support IMM mode.
	FirmwareSupportability *string `json:"FirmwareSupportability,omitempty"`
	// Numeric Identifier assigned by the management system to the equipment.
	Identifier         *int64                    `json:"Identifier,omitempty"`
	IoCardIdentityList []EquipmentIoCardIdentity `json:"IoCardIdentityList,omitempty"`
	// The equipment's lifecycle status. * `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * `Active` - Default Lifecycle State for a physical entity. * `Decommissioned` - Decommission Lifecycle state. * `DecommissionInProgress` - Decommission Inprogress Lifecycle state. * `RecommissionInProgress` - Recommission Inprogress Lifecycle state. * `OperationFailed` - Failed Operation Lifecycle state. * `ReackInProgress` - ReackInProgress Lifecycle state. * `RemoveInProgress` - RemoveInProgress Lifecycle state. * `Discovered` - Discovered Lifecycle state. * `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state. * `DiscoveryFailed` - DiscoveryFailed Lifecycle state. * `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity. * `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.
	Lifecycle *string `json:"Lifecycle,omitempty"`
	// The vendor provided model name for the equipment.
	Model *string `json:"Model,omitempty"`
	// The presence state of the blade server. * `Unknown` - The default presence state. * `Equipped` - The server is equipped in the slot. * `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot. * `Missing` - The server is not present in the given slot.
	Presence *string `json:"Presence,omitempty"`
	// The serial number of the equipment.
	Serial *string `json:"Serial,omitempty"`
	// Chassis slot number of a blade server.
	SlotId *int64 `json:"SlotId,omitempty"`
	// The source object type of this view MO.
	SourceObjectType *string `json:"SourceObjectType,omitempty"`
	// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
	SwitchId *int64 `json:"SwitchId,omitempty"`
	// The manufacturer of the equipment.
	Vendor               *string                              `json:"Vendor,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIdentitySummary EquipmentIdentitySummary

// NewEquipmentIdentitySummary instantiates a new EquipmentIdentitySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIdentitySummary(classId string, objectType string) *EquipmentIdentitySummary {
	this := EquipmentIdentitySummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var adminActionState string = "None"
	this.AdminActionState = &adminActionState
	var firmwareSupportability string = "Unknown"
	this.FirmwareSupportability = &firmwareSupportability
	var lifecycle string = "None"
	this.Lifecycle = &lifecycle
	var presence string = "Unknown"
	this.Presence = &presence
	return &this
}

// NewEquipmentIdentitySummaryWithDefaults instantiates a new EquipmentIdentitySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIdentitySummaryWithDefaults() *EquipmentIdentitySummary {
	this := EquipmentIdentitySummary{}
	var classId string = "equipment.IdentitySummary"
	this.ClassId = classId
	var objectType string = "equipment.IdentitySummary"
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var adminActionState string = "None"
	this.AdminActionState = &adminActionState
	var firmwareSupportability string = "Unknown"
	this.FirmwareSupportability = &firmwareSupportability
	var lifecycle string = "None"
	this.Lifecycle = &lifecycle
	var presence string = "Unknown"
	this.Presence = &presence
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIdentitySummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIdentitySummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIdentitySummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIdentitySummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterSerial returns the AdapterSerial field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetAdapterSerial() string {
	if o == nil || o.AdapterSerial == nil {
		var ret string
		return ret
	}
	return *o.AdapterSerial
}

// GetAdapterSerialOk returns a tuple with the AdapterSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetAdapterSerialOk() (*string, bool) {
	if o == nil || o.AdapterSerial == nil {
		return nil, false
	}
	return o.AdapterSerial, true
}

// HasAdapterSerial returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasAdapterSerial() bool {
	if o != nil && o.AdapterSerial != nil {
		return true
	}

	return false
}

// SetAdapterSerial gets a reference to the given string and assigns it to the AdapterSerial field.
func (o *EquipmentIdentitySummary) SetAdapterSerial(v string) {
	o.AdapterSerial = &v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *EquipmentIdentitySummary) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetAdminActionState returns the AdminActionState field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetAdminActionState() string {
	if o == nil || o.AdminActionState == nil {
		var ret string
		return ret
	}
	return *o.AdminActionState
}

// GetAdminActionStateOk returns a tuple with the AdminActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetAdminActionStateOk() (*string, bool) {
	if o == nil || o.AdminActionState == nil {
		return nil, false
	}
	return o.AdminActionState, true
}

// HasAdminActionState returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasAdminActionState() bool {
	if o != nil && o.AdminActionState != nil {
		return true
	}

	return false
}

// SetAdminActionState gets a reference to the given string and assigns it to the AdminActionState field.
func (o *EquipmentIdentitySummary) SetAdminActionState(v string) {
	o.AdminActionState = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *EquipmentIdentitySummary) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetFirmwareSupportability returns the FirmwareSupportability field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetFirmwareSupportability() string {
	if o == nil || o.FirmwareSupportability == nil {
		var ret string
		return ret
	}
	return *o.FirmwareSupportability
}

// GetFirmwareSupportabilityOk returns a tuple with the FirmwareSupportability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetFirmwareSupportabilityOk() (*string, bool) {
	if o == nil || o.FirmwareSupportability == nil {
		return nil, false
	}
	return o.FirmwareSupportability, true
}

// HasFirmwareSupportability returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasFirmwareSupportability() bool {
	if o != nil && o.FirmwareSupportability != nil {
		return true
	}

	return false
}

// SetFirmwareSupportability gets a reference to the given string and assigns it to the FirmwareSupportability field.
func (o *EquipmentIdentitySummary) SetFirmwareSupportability(v string) {
	o.FirmwareSupportability = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetIdentifier() int64 {
	if o == nil || o.Identifier == nil {
		var ret int64
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetIdentifierOk() (*int64, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given int64 and assigns it to the Identifier field.
func (o *EquipmentIdentitySummary) SetIdentifier(v int64) {
	o.Identifier = &v
}

// GetIoCardIdentityList returns the IoCardIdentityList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIdentitySummary) GetIoCardIdentityList() []EquipmentIoCardIdentity {
	if o == nil {
		var ret []EquipmentIoCardIdentity
		return ret
	}
	return o.IoCardIdentityList
}

// GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIdentitySummary) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool) {
	if o == nil || o.IoCardIdentityList == nil {
		return nil, false
	}
	return &o.IoCardIdentityList, true
}

// HasIoCardIdentityList returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasIoCardIdentityList() bool {
	if o != nil && o.IoCardIdentityList != nil {
		return true
	}

	return false
}

// SetIoCardIdentityList gets a reference to the given []EquipmentIoCardIdentity and assigns it to the IoCardIdentityList field.
func (o *EquipmentIdentitySummary) SetIoCardIdentityList(v []EquipmentIoCardIdentity) {
	o.IoCardIdentityList = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EquipmentIdentitySummary) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentIdentitySummary) SetModel(v string) {
	o.Model = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentIdentitySummary) SetPresence(v string) {
	o.Presence = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentIdentitySummary) SetSerial(v string) {
	o.Serial = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EquipmentIdentitySummary) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSourceObjectType returns the SourceObjectType field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetSourceObjectType() string {
	if o == nil || o.SourceObjectType == nil {
		var ret string
		return ret
	}
	return *o.SourceObjectType
}

// GetSourceObjectTypeOk returns a tuple with the SourceObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetSourceObjectTypeOk() (*string, bool) {
	if o == nil || o.SourceObjectType == nil {
		return nil, false
	}
	return o.SourceObjectType, true
}

// HasSourceObjectType returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasSourceObjectType() bool {
	if o != nil && o.SourceObjectType != nil {
		return true
	}

	return false
}

// SetSourceObjectType gets a reference to the given string and assigns it to the SourceObjectType field.
func (o *EquipmentIdentitySummary) SetSourceObjectType(v string) {
	o.SourceObjectType = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetSwitchId() int64 {
	if o == nil || o.SwitchId == nil {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetSwitchIdOk() (*int64, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentIdentitySummary) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentIdentitySummary) SetVendor(v string) {
	o.Vendor = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIdentitySummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIdentitySummary) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIdentitySummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIdentitySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedViewsView, errViewsView := json.Marshal(o.ViewsView)
	if errViewsView != nil {
		return []byte{}, errViewsView
	}
	errViewsView = json.Unmarshal([]byte(serializedViewsView), &toSerialize)
	if errViewsView != nil {
		return []byte{}, errViewsView
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterSerial != nil {
		toSerialize["AdapterSerial"] = o.AdapterSerial
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.AdminActionState != nil {
		toSerialize["AdminActionState"] = o.AdminActionState
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.FirmwareSupportability != nil {
		toSerialize["FirmwareSupportability"] = o.FirmwareSupportability
	}
	if o.Identifier != nil {
		toSerialize["Identifier"] = o.Identifier
	}
	if o.IoCardIdentityList != nil {
		toSerialize["IoCardIdentityList"] = o.IoCardIdentityList
	}
	if o.Lifecycle != nil {
		toSerialize["Lifecycle"] = o.Lifecycle
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.SourceObjectType != nil {
		toSerialize["SourceObjectType"] = o.SourceObjectType
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIdentitySummary) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentIdentitySummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Serial Identifier of an adapter participating in SWM.
		AdapterSerial *string `json:"AdapterSerial,omitempty"`
		// Updated by UI/API to trigger specific chassis action type. * `None` - No operation value for maintenance actions on an equipment. * `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight. * `Recommission` - Recommission the equipment. * `Reack` - Reacknowledge the equipment and discover it again. * `Remove` - Remove the equipment permanently from Intersight management.
		AdminAction *string `json:"AdminAction,omitempty"`
		// The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		AdminActionState *string `json:"AdminActionState,omitempty"`
		// Chassis Identifier of a blade server.
		ChassisId *int64 `json:"ChassisId,omitempty"`
		// Describes whether the running CIMC version supports Intersight managed mode. * `Unknown` - The running firmware version is unknown. * `Supported` - The running firmware version is known and supports IMM mode. * `NotSupported` - The running firmware version is known and does not support IMM mode.
		FirmwareSupportability *string `json:"FirmwareSupportability,omitempty"`
		// Numeric Identifier assigned by the management system to the equipment.
		Identifier         *int64                    `json:"Identifier,omitempty"`
		IoCardIdentityList []EquipmentIoCardIdentity `json:"IoCardIdentityList,omitempty"`
		// The equipment's lifecycle status. * `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * `Active` - Default Lifecycle State for a physical entity. * `Decommissioned` - Decommission Lifecycle state. * `DecommissionInProgress` - Decommission Inprogress Lifecycle state. * `RecommissionInProgress` - Recommission Inprogress Lifecycle state. * `OperationFailed` - Failed Operation Lifecycle state. * `ReackInProgress` - ReackInProgress Lifecycle state. * `RemoveInProgress` - RemoveInProgress Lifecycle state. * `Discovered` - Discovered Lifecycle state. * `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state. * `DiscoveryFailed` - DiscoveryFailed Lifecycle state. * `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity. * `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.
		Lifecycle *string `json:"Lifecycle,omitempty"`
		// The vendor provided model name for the equipment.
		Model *string `json:"Model,omitempty"`
		// The presence state of the blade server. * `Unknown` - The default presence state. * `Equipped` - The server is equipped in the slot. * `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot. * `Missing` - The server is not present in the given slot.
		Presence *string `json:"Presence,omitempty"`
		// The serial number of the equipment.
		Serial *string `json:"Serial,omitempty"`
		// Chassis slot number of a blade server.
		SlotId *int64 `json:"SlotId,omitempty"`
		// The source object type of this view MO.
		SourceObjectType *string `json:"SourceObjectType,omitempty"`
		// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
		SwitchId *int64 `json:"SwitchId,omitempty"`
		// The manufacturer of the equipment.
		Vendor           *string                              `json:"Vendor,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentIdentitySummaryWithoutEmbeddedStruct := EquipmentIdentitySummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentIdentitySummaryWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIdentitySummary := _EquipmentIdentitySummary{}
		varEquipmentIdentitySummary.ClassId = varEquipmentIdentitySummaryWithoutEmbeddedStruct.ClassId
		varEquipmentIdentitySummary.ObjectType = varEquipmentIdentitySummaryWithoutEmbeddedStruct.ObjectType
		varEquipmentIdentitySummary.AdapterSerial = varEquipmentIdentitySummaryWithoutEmbeddedStruct.AdapterSerial
		varEquipmentIdentitySummary.AdminAction = varEquipmentIdentitySummaryWithoutEmbeddedStruct.AdminAction
		varEquipmentIdentitySummary.AdminActionState = varEquipmentIdentitySummaryWithoutEmbeddedStruct.AdminActionState
		varEquipmentIdentitySummary.ChassisId = varEquipmentIdentitySummaryWithoutEmbeddedStruct.ChassisId
		varEquipmentIdentitySummary.FirmwareSupportability = varEquipmentIdentitySummaryWithoutEmbeddedStruct.FirmwareSupportability
		varEquipmentIdentitySummary.Identifier = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Identifier
		varEquipmentIdentitySummary.IoCardIdentityList = varEquipmentIdentitySummaryWithoutEmbeddedStruct.IoCardIdentityList
		varEquipmentIdentitySummary.Lifecycle = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Lifecycle
		varEquipmentIdentitySummary.Model = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Model
		varEquipmentIdentitySummary.Presence = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Presence
		varEquipmentIdentitySummary.Serial = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Serial
		varEquipmentIdentitySummary.SlotId = varEquipmentIdentitySummaryWithoutEmbeddedStruct.SlotId
		varEquipmentIdentitySummary.SourceObjectType = varEquipmentIdentitySummaryWithoutEmbeddedStruct.SourceObjectType
		varEquipmentIdentitySummary.SwitchId = varEquipmentIdentitySummaryWithoutEmbeddedStruct.SwitchId
		varEquipmentIdentitySummary.Vendor = varEquipmentIdentitySummaryWithoutEmbeddedStruct.Vendor
		varEquipmentIdentitySummary.RegisteredDevice = varEquipmentIdentitySummaryWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentIdentitySummary(varEquipmentIdentitySummary)
	} else {
		return err
	}

	varEquipmentIdentitySummary := _EquipmentIdentitySummary{}

	err = json.Unmarshal(bytes, &varEquipmentIdentitySummary)
	if err == nil {
		o.ViewsView = varEquipmentIdentitySummary.ViewsView
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterSerial")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "AdminActionState")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "FirmwareSupportability")
		delete(additionalProperties, "Identifier")
		delete(additionalProperties, "IoCardIdentityList")
		delete(additionalProperties, "Lifecycle")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SourceObjectType")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectViewsView := reflect.ValueOf(o.ViewsView)
		for i := 0; i < reflectViewsView.Type().NumField(); i++ {
			t := reflectViewsView.Type().Field(i)

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

type NullableEquipmentIdentitySummary struct {
	value *EquipmentIdentitySummary
	isSet bool
}

func (v NullableEquipmentIdentitySummary) Get() *EquipmentIdentitySummary {
	return v.value
}

func (v *NullableEquipmentIdentitySummary) Set(val *EquipmentIdentitySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIdentitySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIdentitySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIdentitySummary(val *EquipmentIdentitySummary) *NullableEquipmentIdentitySummary {
	return &NullableEquipmentIdentitySummary{value: val, isSet: true}
}

func (v NullableEquipmentIdentitySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIdentitySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
