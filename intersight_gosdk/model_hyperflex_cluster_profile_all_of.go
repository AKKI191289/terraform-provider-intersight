/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// HyperflexClusterProfileAllOf Definition of the list of properties defined in 'hyperflex.ClusterProfile', excluding properties defined in parent classes.
type HyperflexClusterProfileAllOf struct {
	// The storage data IP address for the HyperFlex cluster.
	DataIpAddress *string `json:"DataIpAddress,omitempty"`
	// The hypervisor type for the HyperFlex cluster. * `ESXi` - ESXi hypervisor as specified by the user. * `HYPERV` - Hyperv hypervisor as specified by the user. * `KVM` - KVM hypervisor as specified by the user.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The MAC address prefix in the form of 00:25:B5:XX.
	MacAddressPrefix *string `json:"MacAddressPrefix,omitempty"`
	// The management IP address for the HyperFlex cluster.
	MgmtIpAddress *string `json:"MgmtIpAddress,omitempty"`
	// The management platform for the HyperFlex cluster. * `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * `EDGE` - The host servers used in the cluster deployment are standalone severs.
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`
	// The number of copies of each data block written.
	Replication     *int64              `json:"Replication,omitempty"`
	StorageDataVlan *HyperflexNamedVlan `json:"StorageDataVlan,omitempty"`
	// The WWxN prefix in the form of 20:00:00:25:B5:XX.
	WwxnPrefix        *string                                     `json:"WwxnPrefix,omitempty"`
	AssociatedCluster *HyperflexClusterRelationship               `json:"AssociatedCluster,omitempty"`
	AutoSupport       *HyperflexAutoSupportPolicyRelationship     `json:"AutoSupport,omitempty"`
	ClusterNetwork    *HyperflexClusterNetworkPolicyRelationship  `json:"ClusterNetwork,omitempty"`
	ClusterStorage    *HyperflexClusterStoragePolicyRelationship  `json:"ClusterStorage,omitempty"`
	ConfigResult      *HyperflexConfigResultRelationship          `json:"ConfigResult,omitempty"`
	ExtFcStorage      *HyperflexExtFcStoragePolicyRelationship    `json:"ExtFcStorage,omitempty"`
	ExtIscsiStorage   *HyperflexExtIscsiStoragePolicyRelationship `json:"ExtIscsiStorage,omitempty"`
	LocalCredential   *HyperflexLocalCredentialPolicyRelationship `json:"LocalCredential,omitempty"`
	NodeConfig        *HyperflexNodeConfigPolicyRelationship      `json:"NodeConfig,omitempty"`
	// An array of relationships to hyperflexNodeProfile resources.
	NodeProfileConfig []HyperflexNodeProfileRelationship       `json:"NodeProfileConfig,omitempty"`
	Organization      *OrganizationOrganizationRelationship    `json:"Organization,omitempty"`
	ProxySetting      *HyperflexProxySettingPolicyRelationship `json:"ProxySetting,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship          `json:"RunningWorkflows,omitempty"`
	SoftwareVersion      *HyperflexSoftwareVersionPolicyRelationship `json:"SoftwareVersion,omitempty"`
	SysConfig            *HyperflexSysConfigPolicyRelationship       `json:"SysConfig,omitempty"`
	UcsmConfig           *HyperflexUcsmConfigPolicyRelationship      `json:"UcsmConfig,omitempty"`
	VcenterConfig        *HyperflexVcenterConfigPolicyRelationship   `json:"VcenterConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterProfileAllOf HyperflexClusterProfileAllOf

// NewHyperflexClusterProfileAllOf instantiates a new HyperflexClusterProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterProfileAllOf() *HyperflexClusterProfileAllOf {
	this := HyperflexClusterProfileAllOf{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// NewHyperflexClusterProfileAllOfWithDefaults instantiates a new HyperflexClusterProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterProfileAllOfWithDefaults() *HyperflexClusterProfileAllOf {
	this := HyperflexClusterProfileAllOf{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// GetDataIpAddress returns the DataIpAddress field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetDataIpAddress() string {
	if o == nil || o.DataIpAddress == nil {
		var ret string
		return ret
	}
	return *o.DataIpAddress
}

// GetDataIpAddressOk returns a tuple with the DataIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetDataIpAddressOk() (*string, bool) {
	if o == nil || o.DataIpAddress == nil {
		return nil, false
	}
	return o.DataIpAddress, true
}

// HasDataIpAddress returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasDataIpAddress() bool {
	if o != nil && o.DataIpAddress != nil {
		return true
	}

	return false
}

// SetDataIpAddress gets a reference to the given string and assigns it to the DataIpAddress field.
func (o *HyperflexClusterProfileAllOf) SetDataIpAddress(v string) {
	o.DataIpAddress = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexClusterProfileAllOf) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetMacAddressPrefix returns the MacAddressPrefix field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetMacAddressPrefix() string {
	if o == nil || o.MacAddressPrefix == nil {
		var ret string
		return ret
	}
	return *o.MacAddressPrefix
}

// GetMacAddressPrefixOk returns a tuple with the MacAddressPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetMacAddressPrefixOk() (*string, bool) {
	if o == nil || o.MacAddressPrefix == nil {
		return nil, false
	}
	return o.MacAddressPrefix, true
}

// HasMacAddressPrefix returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasMacAddressPrefix() bool {
	if o != nil && o.MacAddressPrefix != nil {
		return true
	}

	return false
}

// SetMacAddressPrefix gets a reference to the given string and assigns it to the MacAddressPrefix field.
func (o *HyperflexClusterProfileAllOf) SetMacAddressPrefix(v string) {
	o.MacAddressPrefix = &v
}

// GetMgmtIpAddress returns the MgmtIpAddress field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetMgmtIpAddress() string {
	if o == nil || o.MgmtIpAddress == nil {
		var ret string
		return ret
	}
	return *o.MgmtIpAddress
}

// GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetMgmtIpAddressOk() (*string, bool) {
	if o == nil || o.MgmtIpAddress == nil {
		return nil, false
	}
	return o.MgmtIpAddress, true
}

// HasMgmtIpAddress returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasMgmtIpAddress() bool {
	if o != nil && o.MgmtIpAddress != nil {
		return true
	}

	return false
}

// SetMgmtIpAddress gets a reference to the given string and assigns it to the MgmtIpAddress field.
func (o *HyperflexClusterProfileAllOf) SetMgmtIpAddress(v string) {
	o.MgmtIpAddress = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetMgmtPlatform() string {
	if o == nil || o.MgmtPlatform == nil {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || o.MgmtPlatform == nil {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasMgmtPlatform() bool {
	if o != nil && o.MgmtPlatform != nil {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *HyperflexClusterProfileAllOf) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetReplication() int64 {
	if o == nil || o.Replication == nil {
		var ret int64
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetReplicationOk() (*int64, bool) {
	if o == nil || o.Replication == nil {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasReplication() bool {
	if o != nil && o.Replication != nil {
		return true
	}

	return false
}

// SetReplication gets a reference to the given int64 and assigns it to the Replication field.
func (o *HyperflexClusterProfileAllOf) SetReplication(v int64) {
	o.Replication = &v
}

// GetStorageDataVlan returns the StorageDataVlan field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetStorageDataVlan() HyperflexNamedVlan {
	if o == nil || o.StorageDataVlan == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.StorageDataVlan
}

// GetStorageDataVlanOk returns a tuple with the StorageDataVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetStorageDataVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil || o.StorageDataVlan == nil {
		return nil, false
	}
	return o.StorageDataVlan, true
}

// HasStorageDataVlan returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasStorageDataVlan() bool {
	if o != nil && o.StorageDataVlan != nil {
		return true
	}

	return false
}

// SetStorageDataVlan gets a reference to the given HyperflexNamedVlan and assigns it to the StorageDataVlan field.
func (o *HyperflexClusterProfileAllOf) SetStorageDataVlan(v HyperflexNamedVlan) {
	o.StorageDataVlan = &v
}

// GetWwxnPrefix returns the WwxnPrefix field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetWwxnPrefix() string {
	if o == nil || o.WwxnPrefix == nil {
		var ret string
		return ret
	}
	return *o.WwxnPrefix
}

// GetWwxnPrefixOk returns a tuple with the WwxnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetWwxnPrefixOk() (*string, bool) {
	if o == nil || o.WwxnPrefix == nil {
		return nil, false
	}
	return o.WwxnPrefix, true
}

// HasWwxnPrefix returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasWwxnPrefix() bool {
	if o != nil && o.WwxnPrefix != nil {
		return true
	}

	return false
}

// SetWwxnPrefix gets a reference to the given string and assigns it to the WwxnPrefix field.
func (o *HyperflexClusterProfileAllOf) SetWwxnPrefix(v string) {
	o.WwxnPrefix = &v
}

// GetAssociatedCluster returns the AssociatedCluster field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetAssociatedCluster() HyperflexClusterRelationship {
	if o == nil || o.AssociatedCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.AssociatedCluster
}

// GetAssociatedClusterOk returns a tuple with the AssociatedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetAssociatedClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.AssociatedCluster == nil {
		return nil, false
	}
	return o.AssociatedCluster, true
}

// HasAssociatedCluster returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasAssociatedCluster() bool {
	if o != nil && o.AssociatedCluster != nil {
		return true
	}

	return false
}

// SetAssociatedCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the AssociatedCluster field.
func (o *HyperflexClusterProfileAllOf) SetAssociatedCluster(v HyperflexClusterRelationship) {
	o.AssociatedCluster = &v
}

// GetAutoSupport returns the AutoSupport field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetAutoSupport() HyperflexAutoSupportPolicyRelationship {
	if o == nil || o.AutoSupport == nil {
		var ret HyperflexAutoSupportPolicyRelationship
		return ret
	}
	return *o.AutoSupport
}

// GetAutoSupportOk returns a tuple with the AutoSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetAutoSupportOk() (*HyperflexAutoSupportPolicyRelationship, bool) {
	if o == nil || o.AutoSupport == nil {
		return nil, false
	}
	return o.AutoSupport, true
}

// HasAutoSupport returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasAutoSupport() bool {
	if o != nil && o.AutoSupport != nil {
		return true
	}

	return false
}

// SetAutoSupport gets a reference to the given HyperflexAutoSupportPolicyRelationship and assigns it to the AutoSupport field.
func (o *HyperflexClusterProfileAllOf) SetAutoSupport(v HyperflexAutoSupportPolicyRelationship) {
	o.AutoSupport = &v
}

// GetClusterNetwork returns the ClusterNetwork field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetClusterNetwork() HyperflexClusterNetworkPolicyRelationship {
	if o == nil || o.ClusterNetwork == nil {
		var ret HyperflexClusterNetworkPolicyRelationship
		return ret
	}
	return *o.ClusterNetwork
}

// GetClusterNetworkOk returns a tuple with the ClusterNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetClusterNetworkOk() (*HyperflexClusterNetworkPolicyRelationship, bool) {
	if o == nil || o.ClusterNetwork == nil {
		return nil, false
	}
	return o.ClusterNetwork, true
}

// HasClusterNetwork returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasClusterNetwork() bool {
	if o != nil && o.ClusterNetwork != nil {
		return true
	}

	return false
}

// SetClusterNetwork gets a reference to the given HyperflexClusterNetworkPolicyRelationship and assigns it to the ClusterNetwork field.
func (o *HyperflexClusterProfileAllOf) SetClusterNetwork(v HyperflexClusterNetworkPolicyRelationship) {
	o.ClusterNetwork = &v
}

// GetClusterStorage returns the ClusterStorage field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetClusterStorage() HyperflexClusterStoragePolicyRelationship {
	if o == nil || o.ClusterStorage == nil {
		var ret HyperflexClusterStoragePolicyRelationship
		return ret
	}
	return *o.ClusterStorage
}

// GetClusterStorageOk returns a tuple with the ClusterStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetClusterStorageOk() (*HyperflexClusterStoragePolicyRelationship, bool) {
	if o == nil || o.ClusterStorage == nil {
		return nil, false
	}
	return o.ClusterStorage, true
}

// HasClusterStorage returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasClusterStorage() bool {
	if o != nil && o.ClusterStorage != nil {
		return true
	}

	return false
}

// SetClusterStorage gets a reference to the given HyperflexClusterStoragePolicyRelationship and assigns it to the ClusterStorage field.
func (o *HyperflexClusterProfileAllOf) SetClusterStorage(v HyperflexClusterStoragePolicyRelationship) {
	o.ClusterStorage = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetConfigResult() HyperflexConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret HyperflexConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetConfigResultOk() (*HyperflexConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given HyperflexConfigResultRelationship and assigns it to the ConfigResult field.
func (o *HyperflexClusterProfileAllOf) SetConfigResult(v HyperflexConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetExtFcStorage returns the ExtFcStorage field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetExtFcStorage() HyperflexExtFcStoragePolicyRelationship {
	if o == nil || o.ExtFcStorage == nil {
		var ret HyperflexExtFcStoragePolicyRelationship
		return ret
	}
	return *o.ExtFcStorage
}

// GetExtFcStorageOk returns a tuple with the ExtFcStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetExtFcStorageOk() (*HyperflexExtFcStoragePolicyRelationship, bool) {
	if o == nil || o.ExtFcStorage == nil {
		return nil, false
	}
	return o.ExtFcStorage, true
}

// HasExtFcStorage returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasExtFcStorage() bool {
	if o != nil && o.ExtFcStorage != nil {
		return true
	}

	return false
}

// SetExtFcStorage gets a reference to the given HyperflexExtFcStoragePolicyRelationship and assigns it to the ExtFcStorage field.
func (o *HyperflexClusterProfileAllOf) SetExtFcStorage(v HyperflexExtFcStoragePolicyRelationship) {
	o.ExtFcStorage = &v
}

// GetExtIscsiStorage returns the ExtIscsiStorage field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetExtIscsiStorage() HyperflexExtIscsiStoragePolicyRelationship {
	if o == nil || o.ExtIscsiStorage == nil {
		var ret HyperflexExtIscsiStoragePolicyRelationship
		return ret
	}
	return *o.ExtIscsiStorage
}

// GetExtIscsiStorageOk returns a tuple with the ExtIscsiStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetExtIscsiStorageOk() (*HyperflexExtIscsiStoragePolicyRelationship, bool) {
	if o == nil || o.ExtIscsiStorage == nil {
		return nil, false
	}
	return o.ExtIscsiStorage, true
}

// HasExtIscsiStorage returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasExtIscsiStorage() bool {
	if o != nil && o.ExtIscsiStorage != nil {
		return true
	}

	return false
}

// SetExtIscsiStorage gets a reference to the given HyperflexExtIscsiStoragePolicyRelationship and assigns it to the ExtIscsiStorage field.
func (o *HyperflexClusterProfileAllOf) SetExtIscsiStorage(v HyperflexExtIscsiStoragePolicyRelationship) {
	o.ExtIscsiStorage = &v
}

// GetLocalCredential returns the LocalCredential field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetLocalCredential() HyperflexLocalCredentialPolicyRelationship {
	if o == nil || o.LocalCredential == nil {
		var ret HyperflexLocalCredentialPolicyRelationship
		return ret
	}
	return *o.LocalCredential
}

// GetLocalCredentialOk returns a tuple with the LocalCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetLocalCredentialOk() (*HyperflexLocalCredentialPolicyRelationship, bool) {
	if o == nil || o.LocalCredential == nil {
		return nil, false
	}
	return o.LocalCredential, true
}

// HasLocalCredential returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasLocalCredential() bool {
	if o != nil && o.LocalCredential != nil {
		return true
	}

	return false
}

// SetLocalCredential gets a reference to the given HyperflexLocalCredentialPolicyRelationship and assigns it to the LocalCredential field.
func (o *HyperflexClusterProfileAllOf) SetLocalCredential(v HyperflexLocalCredentialPolicyRelationship) {
	o.LocalCredential = &v
}

// GetNodeConfig returns the NodeConfig field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetNodeConfig() HyperflexNodeConfigPolicyRelationship {
	if o == nil || o.NodeConfig == nil {
		var ret HyperflexNodeConfigPolicyRelationship
		return ret
	}
	return *o.NodeConfig
}

// GetNodeConfigOk returns a tuple with the NodeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetNodeConfigOk() (*HyperflexNodeConfigPolicyRelationship, bool) {
	if o == nil || o.NodeConfig == nil {
		return nil, false
	}
	return o.NodeConfig, true
}

// HasNodeConfig returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasNodeConfig() bool {
	if o != nil && o.NodeConfig != nil {
		return true
	}

	return false
}

// SetNodeConfig gets a reference to the given HyperflexNodeConfigPolicyRelationship and assigns it to the NodeConfig field.
func (o *HyperflexClusterProfileAllOf) SetNodeConfig(v HyperflexNodeConfigPolicyRelationship) {
	o.NodeConfig = &v
}

// GetNodeProfileConfig returns the NodeProfileConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterProfileAllOf) GetNodeProfileConfig() []HyperflexNodeProfileRelationship {
	if o == nil {
		var ret []HyperflexNodeProfileRelationship
		return ret
	}
	return o.NodeProfileConfig
}

// GetNodeProfileConfigOk returns a tuple with the NodeProfileConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterProfileAllOf) GetNodeProfileConfigOk() (*[]HyperflexNodeProfileRelationship, bool) {
	if o == nil || o.NodeProfileConfig == nil {
		return nil, false
	}
	return &o.NodeProfileConfig, true
}

// HasNodeProfileConfig returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasNodeProfileConfig() bool {
	if o != nil && o.NodeProfileConfig != nil {
		return true
	}

	return false
}

// SetNodeProfileConfig gets a reference to the given []HyperflexNodeProfileRelationship and assigns it to the NodeProfileConfig field.
func (o *HyperflexClusterProfileAllOf) SetNodeProfileConfig(v []HyperflexNodeProfileRelationship) {
	o.NodeProfileConfig = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProxySetting returns the ProxySetting field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetProxySetting() HyperflexProxySettingPolicyRelationship {
	if o == nil || o.ProxySetting == nil {
		var ret HyperflexProxySettingPolicyRelationship
		return ret
	}
	return *o.ProxySetting
}

// GetProxySettingOk returns a tuple with the ProxySetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetProxySettingOk() (*HyperflexProxySettingPolicyRelationship, bool) {
	if o == nil || o.ProxySetting == nil {
		return nil, false
	}
	return o.ProxySetting, true
}

// HasProxySetting returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasProxySetting() bool {
	if o != nil && o.ProxySetting != nil {
		return true
	}

	return false
}

// SetProxySetting gets a reference to the given HyperflexProxySettingPolicyRelationship and assigns it to the ProxySetting field.
func (o *HyperflexClusterProfileAllOf) SetProxySetting(v HyperflexProxySettingPolicyRelationship) {
	o.ProxySetting = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *HyperflexClusterProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetSoftwareVersion() HyperflexSoftwareVersionPolicyRelationship {
	if o == nil || o.SoftwareVersion == nil {
		var ret HyperflexSoftwareVersionPolicyRelationship
		return ret
	}
	return *o.SoftwareVersion
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetSoftwareVersionOk() (*HyperflexSoftwareVersionPolicyRelationship, bool) {
	if o == nil || o.SoftwareVersion == nil {
		return nil, false
	}
	return o.SoftwareVersion, true
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion != nil {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given HyperflexSoftwareVersionPolicyRelationship and assigns it to the SoftwareVersion field.
func (o *HyperflexClusterProfileAllOf) SetSoftwareVersion(v HyperflexSoftwareVersionPolicyRelationship) {
	o.SoftwareVersion = &v
}

// GetSysConfig returns the SysConfig field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetSysConfig() HyperflexSysConfigPolicyRelationship {
	if o == nil || o.SysConfig == nil {
		var ret HyperflexSysConfigPolicyRelationship
		return ret
	}
	return *o.SysConfig
}

// GetSysConfigOk returns a tuple with the SysConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetSysConfigOk() (*HyperflexSysConfigPolicyRelationship, bool) {
	if o == nil || o.SysConfig == nil {
		return nil, false
	}
	return o.SysConfig, true
}

// HasSysConfig returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasSysConfig() bool {
	if o != nil && o.SysConfig != nil {
		return true
	}

	return false
}

// SetSysConfig gets a reference to the given HyperflexSysConfigPolicyRelationship and assigns it to the SysConfig field.
func (o *HyperflexClusterProfileAllOf) SetSysConfig(v HyperflexSysConfigPolicyRelationship) {
	o.SysConfig = &v
}

// GetUcsmConfig returns the UcsmConfig field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetUcsmConfig() HyperflexUcsmConfigPolicyRelationship {
	if o == nil || o.UcsmConfig == nil {
		var ret HyperflexUcsmConfigPolicyRelationship
		return ret
	}
	return *o.UcsmConfig
}

// GetUcsmConfigOk returns a tuple with the UcsmConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetUcsmConfigOk() (*HyperflexUcsmConfigPolicyRelationship, bool) {
	if o == nil || o.UcsmConfig == nil {
		return nil, false
	}
	return o.UcsmConfig, true
}

// HasUcsmConfig returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasUcsmConfig() bool {
	if o != nil && o.UcsmConfig != nil {
		return true
	}

	return false
}

// SetUcsmConfig gets a reference to the given HyperflexUcsmConfigPolicyRelationship and assigns it to the UcsmConfig field.
func (o *HyperflexClusterProfileAllOf) SetUcsmConfig(v HyperflexUcsmConfigPolicyRelationship) {
	o.UcsmConfig = &v
}

// GetVcenterConfig returns the VcenterConfig field value if set, zero value otherwise.
func (o *HyperflexClusterProfileAllOf) GetVcenterConfig() HyperflexVcenterConfigPolicyRelationship {
	if o == nil || o.VcenterConfig == nil {
		var ret HyperflexVcenterConfigPolicyRelationship
		return ret
	}
	return *o.VcenterConfig
}

// GetVcenterConfigOk returns a tuple with the VcenterConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterProfileAllOf) GetVcenterConfigOk() (*HyperflexVcenterConfigPolicyRelationship, bool) {
	if o == nil || o.VcenterConfig == nil {
		return nil, false
	}
	return o.VcenterConfig, true
}

// HasVcenterConfig returns a boolean if a field has been set.
func (o *HyperflexClusterProfileAllOf) HasVcenterConfig() bool {
	if o != nil && o.VcenterConfig != nil {
		return true
	}

	return false
}

// SetVcenterConfig gets a reference to the given HyperflexVcenterConfigPolicyRelationship and assigns it to the VcenterConfig field.
func (o *HyperflexClusterProfileAllOf) SetVcenterConfig(v HyperflexVcenterConfigPolicyRelationship) {
	o.VcenterConfig = &v
}

func (o HyperflexClusterProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataIpAddress != nil {
		toSerialize["DataIpAddress"] = o.DataIpAddress
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.MacAddressPrefix != nil {
		toSerialize["MacAddressPrefix"] = o.MacAddressPrefix
	}
	if o.MgmtIpAddress != nil {
		toSerialize["MgmtIpAddress"] = o.MgmtIpAddress
	}
	if o.MgmtPlatform != nil {
		toSerialize["MgmtPlatform"] = o.MgmtPlatform
	}
	if o.Replication != nil {
		toSerialize["Replication"] = o.Replication
	}
	if o.StorageDataVlan != nil {
		toSerialize["StorageDataVlan"] = o.StorageDataVlan
	}
	if o.WwxnPrefix != nil {
		toSerialize["WwxnPrefix"] = o.WwxnPrefix
	}
	if o.AssociatedCluster != nil {
		toSerialize["AssociatedCluster"] = o.AssociatedCluster
	}
	if o.AutoSupport != nil {
		toSerialize["AutoSupport"] = o.AutoSupport
	}
	if o.ClusterNetwork != nil {
		toSerialize["ClusterNetwork"] = o.ClusterNetwork
	}
	if o.ClusterStorage != nil {
		toSerialize["ClusterStorage"] = o.ClusterStorage
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.ExtFcStorage != nil {
		toSerialize["ExtFcStorage"] = o.ExtFcStorage
	}
	if o.ExtIscsiStorage != nil {
		toSerialize["ExtIscsiStorage"] = o.ExtIscsiStorage
	}
	if o.LocalCredential != nil {
		toSerialize["LocalCredential"] = o.LocalCredential
	}
	if o.NodeConfig != nil {
		toSerialize["NodeConfig"] = o.NodeConfig
	}
	if o.NodeProfileConfig != nil {
		toSerialize["NodeProfileConfig"] = o.NodeProfileConfig
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ProxySetting != nil {
		toSerialize["ProxySetting"] = o.ProxySetting
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}
	if o.SoftwareVersion != nil {
		toSerialize["SoftwareVersion"] = o.SoftwareVersion
	}
	if o.SysConfig != nil {
		toSerialize["SysConfig"] = o.SysConfig
	}
	if o.UcsmConfig != nil {
		toSerialize["UcsmConfig"] = o.UcsmConfig
	}
	if o.VcenterConfig != nil {
		toSerialize["VcenterConfig"] = o.VcenterConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterProfileAllOf := _HyperflexClusterProfileAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterProfileAllOf); err == nil {
		*o = HyperflexClusterProfileAllOf(varHyperflexClusterProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DataIpAddress")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "MacAddressPrefix")
		delete(additionalProperties, "MgmtIpAddress")
		delete(additionalProperties, "MgmtPlatform")
		delete(additionalProperties, "Replication")
		delete(additionalProperties, "StorageDataVlan")
		delete(additionalProperties, "WwxnPrefix")
		delete(additionalProperties, "AssociatedCluster")
		delete(additionalProperties, "AutoSupport")
		delete(additionalProperties, "ClusterNetwork")
		delete(additionalProperties, "ClusterStorage")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "ExtFcStorage")
		delete(additionalProperties, "ExtIscsiStorage")
		delete(additionalProperties, "LocalCredential")
		delete(additionalProperties, "NodeConfig")
		delete(additionalProperties, "NodeProfileConfig")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "ProxySetting")
		delete(additionalProperties, "RunningWorkflows")
		delete(additionalProperties, "SoftwareVersion")
		delete(additionalProperties, "SysConfig")
		delete(additionalProperties, "UcsmConfig")
		delete(additionalProperties, "VcenterConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterProfileAllOf struct {
	value *HyperflexClusterProfileAllOf
	isSet bool
}

func (v NullableHyperflexClusterProfileAllOf) Get() *HyperflexClusterProfileAllOf {
	return v.value
}

func (v *NullableHyperflexClusterProfileAllOf) Set(val *HyperflexClusterProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterProfileAllOf(val *HyperflexClusterProfileAllOf) *NullableHyperflexClusterProfileAllOf {
	return &NullableHyperflexClusterProfileAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
