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

// HyperflexClusterBackupPolicyDeploymentAllOf Definition of the list of properties defined in 'hyperflex.ClusterBackupPolicyDeployment', excluding properties defined in parent classes.
type HyperflexClusterBackupPolicyDeploymentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up.
	BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
	// Replication data store size in backupDataStoreSizeUnit.
	BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
	// Replication data store size.
	BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
	// Description from corresponding ClusterBackupPolicy.
	Description *string `json:"Description,omitempty"`
	// True if record created by discovery on HyperFlex cluster.
	Discovered *bool `json:"Discovered,omitempty"`
	// Name from corresponding ClusterBackupPolicy.
	Name *string `json:"Name,omitempty"`
	// Deployed cluster policy moid.
	PolicyMoid *string `json:"PolicyMoid,omitempty"`
	// Deployed cluster profile moid.
	ProfileMoid *string `json:"ProfileMoid,omitempty"`
	// Replication cluster pairing name prefix.
	ReplicationPairNamePrefix *string                              `json:"ReplicationPairNamePrefix,omitempty"`
	ReplicationSchedule       NullableHyperflexReplicationSchedule `json:"ReplicationSchedule,omitempty"`
	// Number of snapshots that will be retained as part of the Multi Point in Time support.
	SnapshotRetentionCount *int64 `json:"SnapshotRetentionCount,omitempty"`
	// True if policy was detached from source Hyperflex Cluster.
	SourceDetached *bool `json:"SourceDetached,omitempty"`
	// Unique source cluster request ID allowing retry of the same logical request following a transient communication failure.
	SourceRequestId *string `json:"SourceRequestId,omitempty"`
	// Uuid of the source Hyperflex Cluster.
	SourceUuid *string `json:"SourceUuid,omitempty"`
	// True if policy was detached from target Hyperflex Cluster.
	TargetDetached *bool `json:"TargetDetached,omitempty"`
	// Unique target cluster request ID allowing retry of the same logical request following a transient communication failure.
	TargetRequestId *string `json:"TargetRequestId,omitempty"`
	// Uuid of the target Hyperflex Cluster.
	TargetUuid           *string                               `json:"TargetUuid,omitempty"`
	BackupTarget         *HyperflexClusterRelationship         `json:"BackupTarget,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	SourceCluster        *HyperflexClusterRelationship         `json:"SourceCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterBackupPolicyDeploymentAllOf HyperflexClusterBackupPolicyDeploymentAllOf

// NewHyperflexClusterBackupPolicyDeploymentAllOf instantiates a new HyperflexClusterBackupPolicyDeploymentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterBackupPolicyDeploymentAllOf(classId string, objectType string) *HyperflexClusterBackupPolicyDeploymentAllOf {
	this := HyperflexClusterBackupPolicyDeploymentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// NewHyperflexClusterBackupPolicyDeploymentAllOfWithDefaults instantiates a new HyperflexClusterBackupPolicyDeploymentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterBackupPolicyDeploymentAllOfWithDefaults() *HyperflexClusterBackupPolicyDeploymentAllOf {
	this := HyperflexClusterBackupPolicyDeploymentAllOf{}
	var classId string = "hyperflex.ClusterBackupPolicyDeployment"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterBackupPolicyDeployment"
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupDataStoreName returns the BackupDataStoreName field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreName() string {
	if o == nil || o.BackupDataStoreName == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreName
}

// GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreNameOk() (*string, bool) {
	if o == nil || o.BackupDataStoreName == nil {
		return nil, false
	}
	return o.BackupDataStoreName, true
}

// HasBackupDataStoreName returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasBackupDataStoreName() bool {
	if o != nil && o.BackupDataStoreName != nil {
		return true
	}

	return false
}

// SetBackupDataStoreName gets a reference to the given string and assigns it to the BackupDataStoreName field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetBackupDataStoreName(v string) {
	o.BackupDataStoreName = &v
}

// GetBackupDataStoreSize returns the BackupDataStoreSize field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreSize() int64 {
	if o == nil || o.BackupDataStoreSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupDataStoreSize
}

// GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreSizeOk() (*int64, bool) {
	if o == nil || o.BackupDataStoreSize == nil {
		return nil, false
	}
	return o.BackupDataStoreSize, true
}

// HasBackupDataStoreSize returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasBackupDataStoreSize() bool {
	if o != nil && o.BackupDataStoreSize != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSize gets a reference to the given int64 and assigns it to the BackupDataStoreSize field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetBackupDataStoreSize(v int64) {
	o.BackupDataStoreSize = &v
}

// GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreSizeUnit() string {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreSizeUnit
}

// GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupDataStoreSizeUnitOk() (*string, bool) {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		return nil, false
	}
	return o.BackupDataStoreSizeUnit, true
}

// HasBackupDataStoreSizeUnit returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasBackupDataStoreSizeUnit() bool {
	if o != nil && o.BackupDataStoreSizeUnit != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSizeUnit gets a reference to the given string and assigns it to the BackupDataStoreSizeUnit field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetBackupDataStoreSizeUnit(v string) {
	o.BackupDataStoreSizeUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetName(v string) {
	o.Name = &v
}

// GetPolicyMoid returns the PolicyMoid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetPolicyMoid() string {
	if o == nil || o.PolicyMoid == nil {
		var ret string
		return ret
	}
	return *o.PolicyMoid
}

// GetPolicyMoidOk returns a tuple with the PolicyMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetPolicyMoidOk() (*string, bool) {
	if o == nil || o.PolicyMoid == nil {
		return nil, false
	}
	return o.PolicyMoid, true
}

// HasPolicyMoid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasPolicyMoid() bool {
	if o != nil && o.PolicyMoid != nil {
		return true
	}

	return false
}

// SetPolicyMoid gets a reference to the given string and assigns it to the PolicyMoid field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetPolicyMoid(v string) {
	o.PolicyMoid = &v
}

// GetProfileMoid returns the ProfileMoid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetProfileMoid() string {
	if o == nil || o.ProfileMoid == nil {
		var ret string
		return ret
	}
	return *o.ProfileMoid
}

// GetProfileMoidOk returns a tuple with the ProfileMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetProfileMoidOk() (*string, bool) {
	if o == nil || o.ProfileMoid == nil {
		return nil, false
	}
	return o.ProfileMoid, true
}

// HasProfileMoid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasProfileMoid() bool {
	if o != nil && o.ProfileMoid != nil {
		return true
	}

	return false
}

// SetProfileMoid gets a reference to the given string and assigns it to the ProfileMoid field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetProfileMoid(v string) {
	o.ProfileMoid = &v
}

// GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetReplicationPairNamePrefix() string {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.ReplicationPairNamePrefix
}

// GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetReplicationPairNamePrefixOk() (*string, bool) {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		return nil, false
	}
	return o.ReplicationPairNamePrefix, true
}

// HasReplicationPairNamePrefix returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasReplicationPairNamePrefix() bool {
	if o != nil && o.ReplicationPairNamePrefix != nil {
		return true
	}

	return false
}

// SetReplicationPairNamePrefix gets a reference to the given string and assigns it to the ReplicationPairNamePrefix field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetReplicationPairNamePrefix(v string) {
	o.ReplicationPairNamePrefix = &v
}

// GetReplicationSchedule returns the ReplicationSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetReplicationSchedule() HyperflexReplicationSchedule {
	if o == nil || o.ReplicationSchedule.Get() == nil {
		var ret HyperflexReplicationSchedule
		return ret
	}
	return *o.ReplicationSchedule.Get()
}

// GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetReplicationScheduleOk() (*HyperflexReplicationSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationSchedule.Get(), o.ReplicationSchedule.IsSet()
}

// HasReplicationSchedule returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasReplicationSchedule() bool {
	if o != nil && o.ReplicationSchedule.IsSet() {
		return true
	}

	return false
}

// SetReplicationSchedule gets a reference to the given NullableHyperflexReplicationSchedule and assigns it to the ReplicationSchedule field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetReplicationSchedule(v HyperflexReplicationSchedule) {
	o.ReplicationSchedule.Set(&v)
}

// SetReplicationScheduleNil sets the value for ReplicationSchedule to be an explicit nil
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetReplicationScheduleNil() {
	o.ReplicationSchedule.Set(nil)
}

// UnsetReplicationSchedule ensures that no value is present for ReplicationSchedule, not even an explicit nil
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) UnsetReplicationSchedule() {
	o.ReplicationSchedule.Unset()
}

// GetSnapshotRetentionCount returns the SnapshotRetentionCount field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSnapshotRetentionCount() int64 {
	if o == nil || o.SnapshotRetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotRetentionCount
}

// GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSnapshotRetentionCountOk() (*int64, bool) {
	if o == nil || o.SnapshotRetentionCount == nil {
		return nil, false
	}
	return o.SnapshotRetentionCount, true
}

// HasSnapshotRetentionCount returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasSnapshotRetentionCount() bool {
	if o != nil && o.SnapshotRetentionCount != nil {
		return true
	}

	return false
}

// SetSnapshotRetentionCount gets a reference to the given int64 and assigns it to the SnapshotRetentionCount field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetSnapshotRetentionCount(v int64) {
	o.SnapshotRetentionCount = &v
}

// GetSourceDetached returns the SourceDetached field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceDetached() bool {
	if o == nil || o.SourceDetached == nil {
		var ret bool
		return ret
	}
	return *o.SourceDetached
}

// GetSourceDetachedOk returns a tuple with the SourceDetached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceDetachedOk() (*bool, bool) {
	if o == nil || o.SourceDetached == nil {
		return nil, false
	}
	return o.SourceDetached, true
}

// HasSourceDetached returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasSourceDetached() bool {
	if o != nil && o.SourceDetached != nil {
		return true
	}

	return false
}

// SetSourceDetached gets a reference to the given bool and assigns it to the SourceDetached field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetSourceDetached(v bool) {
	o.SourceDetached = &v
}

// GetSourceRequestId returns the SourceRequestId field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceRequestId() string {
	if o == nil || o.SourceRequestId == nil {
		var ret string
		return ret
	}
	return *o.SourceRequestId
}

// GetSourceRequestIdOk returns a tuple with the SourceRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceRequestIdOk() (*string, bool) {
	if o == nil || o.SourceRequestId == nil {
		return nil, false
	}
	return o.SourceRequestId, true
}

// HasSourceRequestId returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasSourceRequestId() bool {
	if o != nil && o.SourceRequestId != nil {
		return true
	}

	return false
}

// SetSourceRequestId gets a reference to the given string and assigns it to the SourceRequestId field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetSourceRequestId(v string) {
	o.SourceRequestId = &v
}

// GetSourceUuid returns the SourceUuid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceUuid() string {
	if o == nil || o.SourceUuid == nil {
		var ret string
		return ret
	}
	return *o.SourceUuid
}

// GetSourceUuidOk returns a tuple with the SourceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceUuidOk() (*string, bool) {
	if o == nil || o.SourceUuid == nil {
		return nil, false
	}
	return o.SourceUuid, true
}

// HasSourceUuid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasSourceUuid() bool {
	if o != nil && o.SourceUuid != nil {
		return true
	}

	return false
}

// SetSourceUuid gets a reference to the given string and assigns it to the SourceUuid field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetSourceUuid(v string) {
	o.SourceUuid = &v
}

// GetTargetDetached returns the TargetDetached field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetDetached() bool {
	if o == nil || o.TargetDetached == nil {
		var ret bool
		return ret
	}
	return *o.TargetDetached
}

// GetTargetDetachedOk returns a tuple with the TargetDetached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetDetachedOk() (*bool, bool) {
	if o == nil || o.TargetDetached == nil {
		return nil, false
	}
	return o.TargetDetached, true
}

// HasTargetDetached returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasTargetDetached() bool {
	if o != nil && o.TargetDetached != nil {
		return true
	}

	return false
}

// SetTargetDetached gets a reference to the given bool and assigns it to the TargetDetached field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetTargetDetached(v bool) {
	o.TargetDetached = &v
}

// GetTargetRequestId returns the TargetRequestId field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetRequestId() string {
	if o == nil || o.TargetRequestId == nil {
		var ret string
		return ret
	}
	return *o.TargetRequestId
}

// GetTargetRequestIdOk returns a tuple with the TargetRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetRequestIdOk() (*string, bool) {
	if o == nil || o.TargetRequestId == nil {
		return nil, false
	}
	return o.TargetRequestId, true
}

// HasTargetRequestId returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasTargetRequestId() bool {
	if o != nil && o.TargetRequestId != nil {
		return true
	}

	return false
}

// SetTargetRequestId gets a reference to the given string and assigns it to the TargetRequestId field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetTargetRequestId(v string) {
	o.TargetRequestId = &v
}

// GetTargetUuid returns the TargetUuid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetUuid() string {
	if o == nil || o.TargetUuid == nil {
		var ret string
		return ret
	}
	return *o.TargetUuid
}

// GetTargetUuidOk returns a tuple with the TargetUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetTargetUuidOk() (*string, bool) {
	if o == nil || o.TargetUuid == nil {
		return nil, false
	}
	return o.TargetUuid, true
}

// HasTargetUuid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasTargetUuid() bool {
	if o != nil && o.TargetUuid != nil {
		return true
	}

	return false
}

// SetTargetUuid gets a reference to the given string and assigns it to the TargetUuid field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetTargetUuid(v string) {
	o.TargetUuid = &v
}

// GetBackupTarget returns the BackupTarget field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupTarget() HyperflexClusterRelationship {
	if o == nil || o.BackupTarget == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.BackupTarget
}

// GetBackupTargetOk returns a tuple with the BackupTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetBackupTargetOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.BackupTarget == nil {
		return nil, false
	}
	return o.BackupTarget, true
}

// HasBackupTarget returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasBackupTarget() bool {
	if o != nil && o.BackupTarget != nil {
		return true
	}

	return false
}

// SetBackupTarget gets a reference to the given HyperflexClusterRelationship and assigns it to the BackupTarget field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetBackupTarget(v HyperflexClusterRelationship) {
	o.BackupTarget = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSourceCluster returns the SourceCluster field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceCluster() HyperflexClusterRelationship {
	if o == nil || o.SourceCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SourceCluster
}

// GetSourceClusterOk returns a tuple with the SourceCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) GetSourceClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SourceCluster == nil {
		return nil, false
	}
	return o.SourceCluster, true
}

// HasSourceCluster returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) HasSourceCluster() bool {
	if o != nil && o.SourceCluster != nil {
		return true
	}

	return false
}

// SetSourceCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SourceCluster field.
func (o *HyperflexClusterBackupPolicyDeploymentAllOf) SetSourceCluster(v HyperflexClusterRelationship) {
	o.SourceCluster = &v
}

func (o HyperflexClusterBackupPolicyDeploymentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupDataStoreName != nil {
		toSerialize["BackupDataStoreName"] = o.BackupDataStoreName
	}
	if o.BackupDataStoreSize != nil {
		toSerialize["BackupDataStoreSize"] = o.BackupDataStoreSize
	}
	if o.BackupDataStoreSizeUnit != nil {
		toSerialize["BackupDataStoreSizeUnit"] = o.BackupDataStoreSizeUnit
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PolicyMoid != nil {
		toSerialize["PolicyMoid"] = o.PolicyMoid
	}
	if o.ProfileMoid != nil {
		toSerialize["ProfileMoid"] = o.ProfileMoid
	}
	if o.ReplicationPairNamePrefix != nil {
		toSerialize["ReplicationPairNamePrefix"] = o.ReplicationPairNamePrefix
	}
	if o.ReplicationSchedule.IsSet() {
		toSerialize["ReplicationSchedule"] = o.ReplicationSchedule.Get()
	}
	if o.SnapshotRetentionCount != nil {
		toSerialize["SnapshotRetentionCount"] = o.SnapshotRetentionCount
	}
	if o.SourceDetached != nil {
		toSerialize["SourceDetached"] = o.SourceDetached
	}
	if o.SourceRequestId != nil {
		toSerialize["SourceRequestId"] = o.SourceRequestId
	}
	if o.SourceUuid != nil {
		toSerialize["SourceUuid"] = o.SourceUuid
	}
	if o.TargetDetached != nil {
		toSerialize["TargetDetached"] = o.TargetDetached
	}
	if o.TargetRequestId != nil {
		toSerialize["TargetRequestId"] = o.TargetRequestId
	}
	if o.TargetUuid != nil {
		toSerialize["TargetUuid"] = o.TargetUuid
	}
	if o.BackupTarget != nil {
		toSerialize["BackupTarget"] = o.BackupTarget
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.SourceCluster != nil {
		toSerialize["SourceCluster"] = o.SourceCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterBackupPolicyDeploymentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterBackupPolicyDeploymentAllOf := _HyperflexClusterBackupPolicyDeploymentAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicyDeploymentAllOf); err == nil {
		*o = HyperflexClusterBackupPolicyDeploymentAllOf(varHyperflexClusterBackupPolicyDeploymentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupDataStoreName")
		delete(additionalProperties, "BackupDataStoreSize")
		delete(additionalProperties, "BackupDataStoreSizeUnit")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Discovered")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PolicyMoid")
		delete(additionalProperties, "ProfileMoid")
		delete(additionalProperties, "ReplicationPairNamePrefix")
		delete(additionalProperties, "ReplicationSchedule")
		delete(additionalProperties, "SnapshotRetentionCount")
		delete(additionalProperties, "SourceDetached")
		delete(additionalProperties, "SourceRequestId")
		delete(additionalProperties, "SourceUuid")
		delete(additionalProperties, "TargetDetached")
		delete(additionalProperties, "TargetRequestId")
		delete(additionalProperties, "TargetUuid")
		delete(additionalProperties, "BackupTarget")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "SourceCluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterBackupPolicyDeploymentAllOf struct {
	value *HyperflexClusterBackupPolicyDeploymentAllOf
	isSet bool
}

func (v NullableHyperflexClusterBackupPolicyDeploymentAllOf) Get() *HyperflexClusterBackupPolicyDeploymentAllOf {
	return v.value
}

func (v *NullableHyperflexClusterBackupPolicyDeploymentAllOf) Set(val *HyperflexClusterBackupPolicyDeploymentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterBackupPolicyDeploymentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterBackupPolicyDeploymentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterBackupPolicyDeploymentAllOf(val *HyperflexClusterBackupPolicyDeploymentAllOf) *NullableHyperflexClusterBackupPolicyDeploymentAllOf {
	return &NullableHyperflexClusterBackupPolicyDeploymentAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterBackupPolicyDeploymentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterBackupPolicyDeploymentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
