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

// HyperflexVmSnapshotInfoAllOf Definition of the list of properties defined in 'hyperflex.VmSnapshotInfo', excluding properties defined in parent classes.
type HyperflexVmSnapshotInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                                   `json:"ObjectType"`
	ClusterIdSnapMap []HyperflexMapClusterIdToStSnapshotPoint `json:"ClusterIdSnapMap,omitempty"`
	ErrorStack       NullableHyperflexErrorStack              `json:"ErrorStack,omitempty"`
	// The name of the Virtual Machine and the time stamp of the snapshot.
	Label *string `json:"Label,omitempty"`
	// Quiesce Mode for snapshot taken on Hyperflex cluster. * `NONE` - The snapshot quiesce mode is none. * `CRASH` - The snapshot quiesce mode is crash. * `VMTOOLS` - The snapshot quiesce mode is VMTOOLS. * `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.
	Mode           *string                          `json:"Mode,omitempty"`
	ParentSnapshot NullableHyperflexEntityReference `json:"ParentSnapshot,omitempty"`
	// Replication status of the least successful target cluster. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
	ReplicationStatus *string `json:"ReplicationStatus,omitempty"`
	// Snapshot status of the source cluster. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none.
	SnapshotStatus *string `json:"SnapshotStatus,omitempty"`
	// Timestamp the snapshot was created on the source cluster.
	SourceTimestamp           *int64                             `json:"SourceTimestamp,omitempty"`
	VmEntityReference         NullableHyperflexEntityReference   `json:"VmEntityReference,omitempty"`
	VmSnapshotEntityReference NullableHyperflexEntityReference   `json:"VmSnapshotEntityReference,omitempty"`
	SrcCluster                *HyperflexClusterRelationship      `json:"SrcCluster,omitempty"`
	TgtCluster                *HyperflexClusterRelationship      `json:"TgtCluster,omitempty"`
	VmBackupInfo              *HyperflexVmBackupInfoRelationship `json:"VmBackupInfo,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _HyperflexVmSnapshotInfoAllOf HyperflexVmSnapshotInfoAllOf

// NewHyperflexVmSnapshotInfoAllOf instantiates a new HyperflexVmSnapshotInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmSnapshotInfoAllOf(classId string, objectType string) *HyperflexVmSnapshotInfoAllOf {
	this := HyperflexVmSnapshotInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "NONE"
	this.Mode = &mode
	var replicationStatus string = "NONE"
	this.ReplicationStatus = &replicationStatus
	var snapshotStatus string = "SUCCESS"
	this.SnapshotStatus = &snapshotStatus
	return &this
}

// NewHyperflexVmSnapshotInfoAllOfWithDefaults instantiates a new HyperflexVmSnapshotInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmSnapshotInfoAllOfWithDefaults() *HyperflexVmSnapshotInfoAllOf {
	this := HyperflexVmSnapshotInfoAllOf{}
	var classId string = "hyperflex.VmSnapshotInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.VmSnapshotInfo"
	this.ObjectType = objectType
	var mode string = "NONE"
	this.Mode = &mode
	var replicationStatus string = "NONE"
	this.ReplicationStatus = &replicationStatus
	var snapshotStatus string = "SUCCESS"
	this.SnapshotStatus = &snapshotStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmSnapshotInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmSnapshotInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmSnapshotInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmSnapshotInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterIdSnapMap returns the ClusterIdSnapMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfoAllOf) GetClusterIdSnapMap() []HyperflexMapClusterIdToStSnapshotPoint {
	if o == nil {
		var ret []HyperflexMapClusterIdToStSnapshotPoint
		return ret
	}
	return o.ClusterIdSnapMap
}

// GetClusterIdSnapMapOk returns a tuple with the ClusterIdSnapMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfoAllOf) GetClusterIdSnapMapOk() (*[]HyperflexMapClusterIdToStSnapshotPoint, bool) {
	if o == nil || o.ClusterIdSnapMap == nil {
		return nil, false
	}
	return &o.ClusterIdSnapMap, true
}

// HasClusterIdSnapMap returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasClusterIdSnapMap() bool {
	if o != nil && o.ClusterIdSnapMap != nil {
		return true
	}

	return false
}

// SetClusterIdSnapMap gets a reference to the given []HyperflexMapClusterIdToStSnapshotPoint and assigns it to the ClusterIdSnapMap field.
func (o *HyperflexVmSnapshotInfoAllOf) SetClusterIdSnapMap(v []HyperflexMapClusterIdToStSnapshotPoint) {
	o.ClusterIdSnapMap = v
}

// GetErrorStack returns the ErrorStack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfoAllOf) GetErrorStack() HyperflexErrorStack {
	if o == nil || o.ErrorStack.Get() == nil {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.ErrorStack.Get()
}

// GetErrorStackOk returns a tuple with the ErrorStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfoAllOf) GetErrorStackOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorStack.Get(), o.ErrorStack.IsSet()
}

// HasErrorStack returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasErrorStack() bool {
	if o != nil && o.ErrorStack.IsSet() {
		return true
	}

	return false
}

// SetErrorStack gets a reference to the given NullableHyperflexErrorStack and assigns it to the ErrorStack field.
func (o *HyperflexVmSnapshotInfoAllOf) SetErrorStack(v HyperflexErrorStack) {
	o.ErrorStack.Set(&v)
}

// SetErrorStackNil sets the value for ErrorStack to be an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) SetErrorStackNil() {
	o.ErrorStack.Set(nil)
}

// UnsetErrorStack ensures that no value is present for ErrorStack, not even an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) UnsetErrorStack() {
	o.ErrorStack.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *HyperflexVmSnapshotInfoAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *HyperflexVmSnapshotInfoAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetParentSnapshot returns the ParentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfoAllOf) GetParentSnapshot() HyperflexEntityReference {
	if o == nil || o.ParentSnapshot.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.ParentSnapshot.Get()
}

// GetParentSnapshotOk returns a tuple with the ParentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfoAllOf) GetParentSnapshotOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSnapshot.Get(), o.ParentSnapshot.IsSet()
}

// HasParentSnapshot returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasParentSnapshot() bool {
	if o != nil && o.ParentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetParentSnapshot gets a reference to the given NullableHyperflexEntityReference and assigns it to the ParentSnapshot field.
func (o *HyperflexVmSnapshotInfoAllOf) SetParentSnapshot(v HyperflexEntityReference) {
	o.ParentSnapshot.Set(&v)
}

// SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) SetParentSnapshotNil() {
	o.ParentSnapshot.Set(nil)
}

// UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) UnsetParentSnapshot() {
	o.ParentSnapshot.Unset()
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetReplicationStatus() string {
	if o == nil || o.ReplicationStatus == nil {
		var ret string
		return ret
	}
	return *o.ReplicationStatus
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetReplicationStatusOk() (*string, bool) {
	if o == nil || o.ReplicationStatus == nil {
		return nil, false
	}
	return o.ReplicationStatus, true
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasReplicationStatus() bool {
	if o != nil && o.ReplicationStatus != nil {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given string and assigns it to the ReplicationStatus field.
func (o *HyperflexVmSnapshotInfoAllOf) SetReplicationStatus(v string) {
	o.ReplicationStatus = &v
}

// GetSnapshotStatus returns the SnapshotStatus field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotStatus() string {
	if o == nil || o.SnapshotStatus == nil {
		var ret string
		return ret
	}
	return *o.SnapshotStatus
}

// GetSnapshotStatusOk returns a tuple with the SnapshotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotStatusOk() (*string, bool) {
	if o == nil || o.SnapshotStatus == nil {
		return nil, false
	}
	return o.SnapshotStatus, true
}

// HasSnapshotStatus returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasSnapshotStatus() bool {
	if o != nil && o.SnapshotStatus != nil {
		return true
	}

	return false
}

// SetSnapshotStatus gets a reference to the given string and assigns it to the SnapshotStatus field.
func (o *HyperflexVmSnapshotInfoAllOf) SetSnapshotStatus(v string) {
	o.SnapshotStatus = &v
}

// GetSourceTimestamp returns the SourceTimestamp field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetSourceTimestamp() int64 {
	if o == nil || o.SourceTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.SourceTimestamp
}

// GetSourceTimestampOk returns a tuple with the SourceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetSourceTimestampOk() (*int64, bool) {
	if o == nil || o.SourceTimestamp == nil {
		return nil, false
	}
	return o.SourceTimestamp, true
}

// HasSourceTimestamp returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasSourceTimestamp() bool {
	if o != nil && o.SourceTimestamp != nil {
		return true
	}

	return false
}

// SetSourceTimestamp gets a reference to the given int64 and assigns it to the SourceTimestamp field.
func (o *HyperflexVmSnapshotInfoAllOf) SetSourceTimestamp(v int64) {
	o.SourceTimestamp = &v
}

// GetVmEntityReference returns the VmEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfoAllOf) GetVmEntityReference() HyperflexEntityReference {
	if o == nil || o.VmEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmEntityReference.Get()
}

// GetVmEntityReferenceOk returns a tuple with the VmEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfoAllOf) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmEntityReference.Get(), o.VmEntityReference.IsSet()
}

// HasVmEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasVmEntityReference() bool {
	if o != nil && o.VmEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmEntityReference field.
func (o *HyperflexVmSnapshotInfoAllOf) SetVmEntityReference(v HyperflexEntityReference) {
	o.VmEntityReference.Set(&v)
}

// SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) SetVmEntityReferenceNil() {
	o.VmEntityReference.Set(nil)
}

// UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) UnsetVmEntityReference() {
	o.VmEntityReference.Unset()
}

// GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfoAllOf) GetVmSnapshotEntityReference() HyperflexEntityReference {
	if o == nil || o.VmSnapshotEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmSnapshotEntityReference.Get()
}

// GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfoAllOf) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSnapshotEntityReference.Get(), o.VmSnapshotEntityReference.IsSet()
}

// HasVmSnapshotEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasVmSnapshotEntityReference() bool {
	if o != nil && o.VmSnapshotEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmSnapshotEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmSnapshotEntityReference field.
func (o *HyperflexVmSnapshotInfoAllOf) SetVmSnapshotEntityReference(v HyperflexEntityReference) {
	o.VmSnapshotEntityReference.Set(&v)
}

// SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) SetVmSnapshotEntityReferenceNil() {
	o.VmSnapshotEntityReference.Set(nil)
}

// UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil
func (o *HyperflexVmSnapshotInfoAllOf) UnsetVmSnapshotEntityReference() {
	o.VmSnapshotEntityReference.Unset()
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexVmSnapshotInfoAllOf) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexVmSnapshotInfoAllOf) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

// GetVmBackupInfo returns the VmBackupInfo field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfoAllOf) GetVmBackupInfo() HyperflexVmBackupInfoRelationship {
	if o == nil || o.VmBackupInfo == nil {
		var ret HyperflexVmBackupInfoRelationship
		return ret
	}
	return *o.VmBackupInfo
}

// GetVmBackupInfoOk returns a tuple with the VmBackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfoAllOf) GetVmBackupInfoOk() (*HyperflexVmBackupInfoRelationship, bool) {
	if o == nil || o.VmBackupInfo == nil {
		return nil, false
	}
	return o.VmBackupInfo, true
}

// HasVmBackupInfo returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfoAllOf) HasVmBackupInfo() bool {
	if o != nil && o.VmBackupInfo != nil {
		return true
	}

	return false
}

// SetVmBackupInfo gets a reference to the given HyperflexVmBackupInfoRelationship and assigns it to the VmBackupInfo field.
func (o *HyperflexVmSnapshotInfoAllOf) SetVmBackupInfo(v HyperflexVmBackupInfoRelationship) {
	o.VmBackupInfo = &v
}

func (o HyperflexVmSnapshotInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterIdSnapMap != nil {
		toSerialize["ClusterIdSnapMap"] = o.ClusterIdSnapMap
	}
	if o.ErrorStack.IsSet() {
		toSerialize["ErrorStack"] = o.ErrorStack.Get()
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.ParentSnapshot.IsSet() {
		toSerialize["ParentSnapshot"] = o.ParentSnapshot.Get()
	}
	if o.ReplicationStatus != nil {
		toSerialize["ReplicationStatus"] = o.ReplicationStatus
	}
	if o.SnapshotStatus != nil {
		toSerialize["SnapshotStatus"] = o.SnapshotStatus
	}
	if o.SourceTimestamp != nil {
		toSerialize["SourceTimestamp"] = o.SourceTimestamp
	}
	if o.VmEntityReference.IsSet() {
		toSerialize["VmEntityReference"] = o.VmEntityReference.Get()
	}
	if o.VmSnapshotEntityReference.IsSet() {
		toSerialize["VmSnapshotEntityReference"] = o.VmSnapshotEntityReference.Get()
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}
	if o.VmBackupInfo != nil {
		toSerialize["VmBackupInfo"] = o.VmBackupInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVmSnapshotInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVmSnapshotInfoAllOf := _HyperflexVmSnapshotInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVmSnapshotInfoAllOf); err == nil {
		*o = HyperflexVmSnapshotInfoAllOf(varHyperflexVmSnapshotInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterIdSnapMap")
		delete(additionalProperties, "ErrorStack")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "ParentSnapshot")
		delete(additionalProperties, "ReplicationStatus")
		delete(additionalProperties, "SnapshotStatus")
		delete(additionalProperties, "SourceTimestamp")
		delete(additionalProperties, "VmEntityReference")
		delete(additionalProperties, "VmSnapshotEntityReference")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")
		delete(additionalProperties, "VmBackupInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVmSnapshotInfoAllOf struct {
	value *HyperflexVmSnapshotInfoAllOf
	isSet bool
}

func (v NullableHyperflexVmSnapshotInfoAllOf) Get() *HyperflexVmSnapshotInfoAllOf {
	return v.value
}

func (v *NullableHyperflexVmSnapshotInfoAllOf) Set(val *HyperflexVmSnapshotInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmSnapshotInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmSnapshotInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmSnapshotInfoAllOf(val *HyperflexVmSnapshotInfoAllOf) *NullableHyperflexVmSnapshotInfoAllOf {
	return &NullableHyperflexVmSnapshotInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVmSnapshotInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmSnapshotInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
