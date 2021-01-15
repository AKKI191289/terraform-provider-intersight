/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexVmRestoreOperationAllOf Definition of the list of properties defined in 'hyperflex.VmRestoreOperation', excluding properties defined in parent classes.
type HyperflexVmRestoreOperationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// New name for the Virtual Machine after recovery.
	NewName *string `json:"NewName,omitempty"`
	// Power on the Virtual Machine after recovery.
	PowerOn                *bool                                 `json:"PowerOn,omitempty"`
	Organization           *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	RestoreEdgeClusterMoid *HyperflexClusterRelationship         `json:"RestoreEdgeClusterMoid,omitempty"`
	VmBackupInfoMoid       *HyperflexVmBackupInfoRelationship    `json:"VmBackupInfoMoid,omitempty"`
	VmSnapshotInfoMoid     *HyperflexVmSnapshotInfoRelationship  `json:"VmSnapshotInfoMoid,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _HyperflexVmRestoreOperationAllOf HyperflexVmRestoreOperationAllOf

// NewHyperflexVmRestoreOperationAllOf instantiates a new HyperflexVmRestoreOperationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmRestoreOperationAllOf(classId string, objectType string) *HyperflexVmRestoreOperationAllOf {
	this := HyperflexVmRestoreOperationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var powerOn bool = true
	this.PowerOn = &powerOn
	return &this
}

// NewHyperflexVmRestoreOperationAllOfWithDefaults instantiates a new HyperflexVmRestoreOperationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmRestoreOperationAllOfWithDefaults() *HyperflexVmRestoreOperationAllOf {
	this := HyperflexVmRestoreOperationAllOf{}
	var classId string = "hyperflex.VmRestoreOperation"
	this.ClassId = classId
	var objectType string = "hyperflex.VmRestoreOperation"
	this.ObjectType = objectType
	var powerOn bool = true
	this.PowerOn = &powerOn
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmRestoreOperationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmRestoreOperationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmRestoreOperationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmRestoreOperationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *HyperflexVmRestoreOperationAllOf) SetNewName(v string) {
	o.NewName = &v
}

// GetPowerOn returns the PowerOn field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetPowerOn() bool {
	if o == nil || o.PowerOn == nil {
		var ret bool
		return ret
	}
	return *o.PowerOn
}

// GetPowerOnOk returns a tuple with the PowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetPowerOnOk() (*bool, bool) {
	if o == nil || o.PowerOn == nil {
		return nil, false
	}
	return o.PowerOn, true
}

// HasPowerOn returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasPowerOn() bool {
	if o != nil && o.PowerOn != nil {
		return true
	}

	return false
}

// SetPowerOn gets a reference to the given bool and assigns it to the PowerOn field.
func (o *HyperflexVmRestoreOperationAllOf) SetPowerOn(v bool) {
	o.PowerOn = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexVmRestoreOperationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRestoreEdgeClusterMoid returns the RestoreEdgeClusterMoid field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetRestoreEdgeClusterMoid() HyperflexClusterRelationship {
	if o == nil || o.RestoreEdgeClusterMoid == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.RestoreEdgeClusterMoid
}

// GetRestoreEdgeClusterMoidOk returns a tuple with the RestoreEdgeClusterMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetRestoreEdgeClusterMoidOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.RestoreEdgeClusterMoid == nil {
		return nil, false
	}
	return o.RestoreEdgeClusterMoid, true
}

// HasRestoreEdgeClusterMoid returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasRestoreEdgeClusterMoid() bool {
	if o != nil && o.RestoreEdgeClusterMoid != nil {
		return true
	}

	return false
}

// SetRestoreEdgeClusterMoid gets a reference to the given HyperflexClusterRelationship and assigns it to the RestoreEdgeClusterMoid field.
func (o *HyperflexVmRestoreOperationAllOf) SetRestoreEdgeClusterMoid(v HyperflexClusterRelationship) {
	o.RestoreEdgeClusterMoid = &v
}

// GetVmBackupInfoMoid returns the VmBackupInfoMoid field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetVmBackupInfoMoid() HyperflexVmBackupInfoRelationship {
	if o == nil || o.VmBackupInfoMoid == nil {
		var ret HyperflexVmBackupInfoRelationship
		return ret
	}
	return *o.VmBackupInfoMoid
}

// GetVmBackupInfoMoidOk returns a tuple with the VmBackupInfoMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetVmBackupInfoMoidOk() (*HyperflexVmBackupInfoRelationship, bool) {
	if o == nil || o.VmBackupInfoMoid == nil {
		return nil, false
	}
	return o.VmBackupInfoMoid, true
}

// HasVmBackupInfoMoid returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasVmBackupInfoMoid() bool {
	if o != nil && o.VmBackupInfoMoid != nil {
		return true
	}

	return false
}

// SetVmBackupInfoMoid gets a reference to the given HyperflexVmBackupInfoRelationship and assigns it to the VmBackupInfoMoid field.
func (o *HyperflexVmRestoreOperationAllOf) SetVmBackupInfoMoid(v HyperflexVmBackupInfoRelationship) {
	o.VmBackupInfoMoid = &v
}

// GetVmSnapshotInfoMoid returns the VmSnapshotInfoMoid field value if set, zero value otherwise.
func (o *HyperflexVmRestoreOperationAllOf) GetVmSnapshotInfoMoid() HyperflexVmSnapshotInfoRelationship {
	if o == nil || o.VmSnapshotInfoMoid == nil {
		var ret HyperflexVmSnapshotInfoRelationship
		return ret
	}
	return *o.VmSnapshotInfoMoid
}

// GetVmSnapshotInfoMoidOk returns a tuple with the VmSnapshotInfoMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmRestoreOperationAllOf) GetVmSnapshotInfoMoidOk() (*HyperflexVmSnapshotInfoRelationship, bool) {
	if o == nil || o.VmSnapshotInfoMoid == nil {
		return nil, false
	}
	return o.VmSnapshotInfoMoid, true
}

// HasVmSnapshotInfoMoid returns a boolean if a field has been set.
func (o *HyperflexVmRestoreOperationAllOf) HasVmSnapshotInfoMoid() bool {
	if o != nil && o.VmSnapshotInfoMoid != nil {
		return true
	}

	return false
}

// SetVmSnapshotInfoMoid gets a reference to the given HyperflexVmSnapshotInfoRelationship and assigns it to the VmSnapshotInfoMoid field.
func (o *HyperflexVmRestoreOperationAllOf) SetVmSnapshotInfoMoid(v HyperflexVmSnapshotInfoRelationship) {
	o.VmSnapshotInfoMoid = &v
}

func (o HyperflexVmRestoreOperationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NewName != nil {
		toSerialize["NewName"] = o.NewName
	}
	if o.PowerOn != nil {
		toSerialize["PowerOn"] = o.PowerOn
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RestoreEdgeClusterMoid != nil {
		toSerialize["RestoreEdgeClusterMoid"] = o.RestoreEdgeClusterMoid
	}
	if o.VmBackupInfoMoid != nil {
		toSerialize["VmBackupInfoMoid"] = o.VmBackupInfoMoid
	}
	if o.VmSnapshotInfoMoid != nil {
		toSerialize["VmSnapshotInfoMoid"] = o.VmSnapshotInfoMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVmRestoreOperationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVmRestoreOperationAllOf := _HyperflexVmRestoreOperationAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVmRestoreOperationAllOf); err == nil {
		*o = HyperflexVmRestoreOperationAllOf(varHyperflexVmRestoreOperationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NewName")
		delete(additionalProperties, "PowerOn")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RestoreEdgeClusterMoid")
		delete(additionalProperties, "VmBackupInfoMoid")
		delete(additionalProperties, "VmSnapshotInfoMoid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVmRestoreOperationAllOf struct {
	value *HyperflexVmRestoreOperationAllOf
	isSet bool
}

func (v NullableHyperflexVmRestoreOperationAllOf) Get() *HyperflexVmRestoreOperationAllOf {
	return v.value
}

func (v *NullableHyperflexVmRestoreOperationAllOf) Set(val *HyperflexVmRestoreOperationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmRestoreOperationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmRestoreOperationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmRestoreOperationAllOf(val *HyperflexVmRestoreOperationAllOf) *NullableHyperflexVmRestoreOperationAllOf {
	return &NullableHyperflexVmRestoreOperationAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVmRestoreOperationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmRestoreOperationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
