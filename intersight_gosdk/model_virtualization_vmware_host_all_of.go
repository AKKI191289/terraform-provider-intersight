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
	"time"
)

// VirtualizationVmwareHostAllOf Definition of the list of properties defined in 'virtualization.VmwareHost', excluding properties defined in parent classes.
type VirtualizationVmwareHostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time when this host booted up.
	BootTime *time.Time `json:"BootTime,omitempty"`
	// Indicates if the host is connected to the vCenter. Values are connected, not connected.
	ConnectionState *string `json:"ConnectionState,omitempty"`
	// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweredOn` - The entity is powered on. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state.
	HwPowerState *string `json:"HwPowerState,omitempty"`
	// The count of all network adapters attached to this host.
	NetworkAdapterCount *int64                                          `json:"NetworkAdapterCount,omitempty"`
	ResourceConsumed    NullableVirtualizationVmwareResourceConsumption `json:"ResourceConsumed,omitempty"`
	// The count of all storage adapters attached to this host.
	StorageAdapterCount *int64 `json:"StorageAdapterCount,omitempty"`
	// The identity of this host within vCenter (optional).
	VcenterHostId *string                                     `json:"VcenterHostId,omitempty"`
	Cluster       *VirtualizationVmwareClusterRelationship    `json:"Cluster,omitempty"`
	Datacenter    *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	// An array of relationships to virtualizationVmwareDatastore resources.
	Datastores           []VirtualizationVmwareDatastoreRelationship `json:"Datastores,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareHostAllOf VirtualizationVmwareHostAllOf

// NewVirtualizationVmwareHostAllOf instantiates a new VirtualizationVmwareHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareHostAllOf(classId string, objectType string) *VirtualizationVmwareHostAllOf {
	this := VirtualizationVmwareHostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// NewVirtualizationVmwareHostAllOfWithDefaults instantiates a new VirtualizationVmwareHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareHostAllOfWithDefaults() *VirtualizationVmwareHostAllOf {
	this := VirtualizationVmwareHostAllOf{}
	var classId string = "virtualization.VmwareHost"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareHost"
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareHostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareHostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareHostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareHostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootTime returns the BootTime field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetBootTime() time.Time {
	if o == nil || o.BootTime == nil {
		var ret time.Time
		return ret
	}
	return *o.BootTime
}

// GetBootTimeOk returns a tuple with the BootTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetBootTimeOk() (*time.Time, bool) {
	if o == nil || o.BootTime == nil {
		return nil, false
	}
	return o.BootTime, true
}

// HasBootTime returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasBootTime() bool {
	if o != nil && o.BootTime != nil {
		return true
	}

	return false
}

// SetBootTime gets a reference to the given time.Time and assigns it to the BootTime field.
func (o *VirtualizationVmwareHostAllOf) SetBootTime(v time.Time) {
	o.BootTime = &v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetConnectionState() string {
	if o == nil || o.ConnectionState == nil {
		var ret string
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetConnectionStateOk() (*string, bool) {
	if o == nil || o.ConnectionState == nil {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasConnectionState() bool {
	if o != nil && o.ConnectionState != nil {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given string and assigns it to the ConnectionState field.
func (o *VirtualizationVmwareHostAllOf) SetConnectionState(v string) {
	o.ConnectionState = &v
}

// GetHwPowerState returns the HwPowerState field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetHwPowerState() string {
	if o == nil || o.HwPowerState == nil {
		var ret string
		return ret
	}
	return *o.HwPowerState
}

// GetHwPowerStateOk returns a tuple with the HwPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetHwPowerStateOk() (*string, bool) {
	if o == nil || o.HwPowerState == nil {
		return nil, false
	}
	return o.HwPowerState, true
}

// HasHwPowerState returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasHwPowerState() bool {
	if o != nil && o.HwPowerState != nil {
		return true
	}

	return false
}

// SetHwPowerState gets a reference to the given string and assigns it to the HwPowerState field.
func (o *VirtualizationVmwareHostAllOf) SetHwPowerState(v string) {
	o.HwPowerState = &v
}

// GetNetworkAdapterCount returns the NetworkAdapterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetNetworkAdapterCount() int64 {
	if o == nil || o.NetworkAdapterCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkAdapterCount
}

// GetNetworkAdapterCountOk returns a tuple with the NetworkAdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetNetworkAdapterCountOk() (*int64, bool) {
	if o == nil || o.NetworkAdapterCount == nil {
		return nil, false
	}
	return o.NetworkAdapterCount, true
}

// HasNetworkAdapterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasNetworkAdapterCount() bool {
	if o != nil && o.NetworkAdapterCount != nil {
		return true
	}

	return false
}

// SetNetworkAdapterCount gets a reference to the given int64 and assigns it to the NetworkAdapterCount field.
func (o *VirtualizationVmwareHostAllOf) SetNetworkAdapterCount(v int64) {
	o.NetworkAdapterCount = &v
}

// GetResourceConsumed returns the ResourceConsumed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHostAllOf) GetResourceConsumed() VirtualizationVmwareResourceConsumption {
	if o == nil || o.ResourceConsumed.Get() == nil {
		var ret VirtualizationVmwareResourceConsumption
		return ret
	}
	return *o.ResourceConsumed.Get()
}

// GetResourceConsumedOk returns a tuple with the ResourceConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHostAllOf) GetResourceConsumedOk() (*VirtualizationVmwareResourceConsumption, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceConsumed.Get(), o.ResourceConsumed.IsSet()
}

// HasResourceConsumed returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasResourceConsumed() bool {
	if o != nil && o.ResourceConsumed.IsSet() {
		return true
	}

	return false
}

// SetResourceConsumed gets a reference to the given NullableVirtualizationVmwareResourceConsumption and assigns it to the ResourceConsumed field.
func (o *VirtualizationVmwareHostAllOf) SetResourceConsumed(v VirtualizationVmwareResourceConsumption) {
	o.ResourceConsumed.Set(&v)
}

// SetResourceConsumedNil sets the value for ResourceConsumed to be an explicit nil
func (o *VirtualizationVmwareHostAllOf) SetResourceConsumedNil() {
	o.ResourceConsumed.Set(nil)
}

// UnsetResourceConsumed ensures that no value is present for ResourceConsumed, not even an explicit nil
func (o *VirtualizationVmwareHostAllOf) UnsetResourceConsumed() {
	o.ResourceConsumed.Unset()
}

// GetStorageAdapterCount returns the StorageAdapterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetStorageAdapterCount() int64 {
	if o == nil || o.StorageAdapterCount == nil {
		var ret int64
		return ret
	}
	return *o.StorageAdapterCount
}

// GetStorageAdapterCountOk returns a tuple with the StorageAdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetStorageAdapterCountOk() (*int64, bool) {
	if o == nil || o.StorageAdapterCount == nil {
		return nil, false
	}
	return o.StorageAdapterCount, true
}

// HasStorageAdapterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasStorageAdapterCount() bool {
	if o != nil && o.StorageAdapterCount != nil {
		return true
	}

	return false
}

// SetStorageAdapterCount gets a reference to the given int64 and assigns it to the StorageAdapterCount field.
func (o *VirtualizationVmwareHostAllOf) SetStorageAdapterCount(v int64) {
	o.StorageAdapterCount = &v
}

// GetVcenterHostId returns the VcenterHostId field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetVcenterHostId() string {
	if o == nil || o.VcenterHostId == nil {
		var ret string
		return ret
	}
	return *o.VcenterHostId
}

// GetVcenterHostIdOk returns a tuple with the VcenterHostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetVcenterHostIdOk() (*string, bool) {
	if o == nil || o.VcenterHostId == nil {
		return nil, false
	}
	return o.VcenterHostId, true
}

// HasVcenterHostId returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasVcenterHostId() bool {
	if o != nil && o.VcenterHostId != nil {
		return true
	}

	return false
}

// SetVcenterHostId gets a reference to the given string and assigns it to the VcenterHostId field.
func (o *VirtualizationVmwareHostAllOf) SetVcenterHostId(v string) {
	o.VcenterHostId = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetCluster() VirtualizationVmwareClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationVmwareClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationVmwareClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVmwareHostAllOf) SetCluster(v VirtualizationVmwareClusterRelationship) {
	o.Cluster = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareHostAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHostAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareHostAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHostAllOf) GetDatastores() []VirtualizationVmwareDatastoreRelationship {
	if o == nil {
		var ret []VirtualizationVmwareDatastoreRelationship
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHostAllOf) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return &o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VirtualizationVmwareHostAllOf) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []VirtualizationVmwareDatastoreRelationship and assigns it to the Datastores field.
func (o *VirtualizationVmwareHostAllOf) SetDatastores(v []VirtualizationVmwareDatastoreRelationship) {
	o.Datastores = v
}

func (o VirtualizationVmwareHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootTime != nil {
		toSerialize["BootTime"] = o.BootTime
	}
	if o.ConnectionState != nil {
		toSerialize["ConnectionState"] = o.ConnectionState
	}
	if o.HwPowerState != nil {
		toSerialize["HwPowerState"] = o.HwPowerState
	}
	if o.NetworkAdapterCount != nil {
		toSerialize["NetworkAdapterCount"] = o.NetworkAdapterCount
	}
	if o.ResourceConsumed.IsSet() {
		toSerialize["ResourceConsumed"] = o.ResourceConsumed.Get()
	}
	if o.StorageAdapterCount != nil {
		toSerialize["StorageAdapterCount"] = o.StorageAdapterCount
	}
	if o.VcenterHostId != nil {
		toSerialize["VcenterHostId"] = o.VcenterHostId
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.Datastores != nil {
		toSerialize["Datastores"] = o.Datastores
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareHostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareHostAllOf := _VirtualizationVmwareHostAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareHostAllOf); err == nil {
		*o = VirtualizationVmwareHostAllOf(varVirtualizationVmwareHostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootTime")
		delete(additionalProperties, "ConnectionState")
		delete(additionalProperties, "HwPowerState")
		delete(additionalProperties, "NetworkAdapterCount")
		delete(additionalProperties, "ResourceConsumed")
		delete(additionalProperties, "StorageAdapterCount")
		delete(additionalProperties, "VcenterHostId")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "Datastores")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareHostAllOf struct {
	value *VirtualizationVmwareHostAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareHostAllOf) Get() *VirtualizationVmwareHostAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareHostAllOf) Set(val *VirtualizationVmwareHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareHostAllOf(val *VirtualizationVmwareHostAllOf) *NullableVirtualizationVmwareHostAllOf {
	return &NullableVirtualizationVmwareHostAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
