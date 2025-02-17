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

// HyperflexHxapHostAllOf Definition of the list of properties defined in 'hyperflex.HxapHost', excluding properties defined in parent classes.
type HyperflexHxapHostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Denotes age or life time of the Host in nano seconds.
	Age *string `json:"Age,omitempty"`
	// The UUID of the cluster to which this Host belongs to.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// Reason of the failure when host is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweredOn` - The entity is powered on. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state.
	HwPowerState *string `json:"HwPowerState,omitempty"`
	// Internal IP Address of the Host.
	InternalIpAddress *string `json:"InternalIpAddress,omitempty"`
	// Management IP Address of the Host.
	ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
	// Is the role of this host is master in the cluster? true or false.
	MasterRole *bool `json:"MasterRole,omitempty"`
	// Product version of the Host.
	Version              *string                           `json:"Version,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	ClusterMember        *AssetClusterMemberRelationship   `json:"ClusterMember,omitempty"`
	PhysicalServer       *ComputePhysicalRelationship      `json:"PhysicalServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapHostAllOf HyperflexHxapHostAllOf

// NewHyperflexHxapHostAllOf instantiates a new HyperflexHxapHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapHostAllOf(classId string, objectType string) *HyperflexHxapHostAllOf {
	this := HyperflexHxapHostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// NewHyperflexHxapHostAllOfWithDefaults instantiates a new HyperflexHxapHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapHostAllOfWithDefaults() *HyperflexHxapHostAllOf {
	this := HyperflexHxapHostAllOf{}
	var classId string = "hyperflex.HxapHost"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapHost"
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapHostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapHostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapHostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapHostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *HyperflexHxapHostAllOf) SetAge(v string) {
	o.Age = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexHxapHostAllOf) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *HyperflexHxapHostAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHwPowerState returns the HwPowerState field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetHwPowerState() string {
	if o == nil || o.HwPowerState == nil {
		var ret string
		return ret
	}
	return *o.HwPowerState
}

// GetHwPowerStateOk returns a tuple with the HwPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetHwPowerStateOk() (*string, bool) {
	if o == nil || o.HwPowerState == nil {
		return nil, false
	}
	return o.HwPowerState, true
}

// HasHwPowerState returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasHwPowerState() bool {
	if o != nil && o.HwPowerState != nil {
		return true
	}

	return false
}

// SetHwPowerState gets a reference to the given string and assigns it to the HwPowerState field.
func (o *HyperflexHxapHostAllOf) SetHwPowerState(v string) {
	o.HwPowerState = &v
}

// GetInternalIpAddress returns the InternalIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetInternalIpAddress() string {
	if o == nil || o.InternalIpAddress == nil {
		var ret string
		return ret
	}
	return *o.InternalIpAddress
}

// GetInternalIpAddressOk returns a tuple with the InternalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetInternalIpAddressOk() (*string, bool) {
	if o == nil || o.InternalIpAddress == nil {
		return nil, false
	}
	return o.InternalIpAddress, true
}

// HasInternalIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasInternalIpAddress() bool {
	if o != nil && o.InternalIpAddress != nil {
		return true
	}

	return false
}

// SetInternalIpAddress gets a reference to the given string and assigns it to the InternalIpAddress field.
func (o *HyperflexHxapHostAllOf) SetInternalIpAddress(v string) {
	o.InternalIpAddress = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *HyperflexHxapHostAllOf) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetMasterRole returns the MasterRole field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetMasterRole() bool {
	if o == nil || o.MasterRole == nil {
		var ret bool
		return ret
	}
	return *o.MasterRole
}

// GetMasterRoleOk returns a tuple with the MasterRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetMasterRoleOk() (*bool, bool) {
	if o == nil || o.MasterRole == nil {
		return nil, false
	}
	return o.MasterRole, true
}

// HasMasterRole returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasMasterRole() bool {
	if o != nil && o.MasterRole != nil {
		return true
	}

	return false
}

// SetMasterRole gets a reference to the given bool and assigns it to the MasterRole field.
func (o *HyperflexHxapHostAllOf) SetMasterRole(v bool) {
	o.MasterRole = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHxapHostAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapHostAllOf) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *HyperflexHxapHostAllOf) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *HyperflexHxapHostAllOf) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *HyperflexHxapHostAllOf) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *HyperflexHxapHostAllOf) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o HyperflexHxapHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.HwPowerState != nil {
		toSerialize["HwPowerState"] = o.HwPowerState
	}
	if o.InternalIpAddress != nil {
		toSerialize["InternalIpAddress"] = o.InternalIpAddress
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.MasterRole != nil {
		toSerialize["MasterRole"] = o.MasterRole
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapHostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapHostAllOf := _HyperflexHxapHostAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapHostAllOf); err == nil {
		*o = HyperflexHxapHostAllOf(varHyperflexHxapHostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Age")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "HwPowerState")
		delete(additionalProperties, "InternalIpAddress")
		delete(additionalProperties, "ManagementIpAddress")
		delete(additionalProperties, "MasterRole")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "PhysicalServer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapHostAllOf struct {
	value *HyperflexHxapHostAllOf
	isSet bool
}

func (v NullableHyperflexHxapHostAllOf) Get() *HyperflexHxapHostAllOf {
	return v.value
}

func (v *NullableHyperflexHxapHostAllOf) Set(val *HyperflexHxapHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapHostAllOf(val *HyperflexHxapHostAllOf) *NullableHyperflexHxapHostAllOf {
	return &NullableHyperflexHxapHostAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
