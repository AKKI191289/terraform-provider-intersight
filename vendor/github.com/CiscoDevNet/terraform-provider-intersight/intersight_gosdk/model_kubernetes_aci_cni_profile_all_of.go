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

// KubernetesAciCniProfileAllOf Definition of the list of properties defined in 'kubernetes.AciCniProfile', excluding properties defined in parent classes.
type KubernetesAciCniProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy.
	AaepName *string `json:"AaepName,omitempty"`
	// Start of range of IP subnets for external services with dynamic IP allocation for use by Kubernetes clusters using this ACI CNI policy.
	ExtSvcDynSubnetStart *string `json:"ExtSvcDynSubnetStart,omitempty"`
	// Start of range of IP subnets for external services with static IP allocation for use by Kubernetes clusters using this ACI CNI policy.
	ExtSvcStaticSubnetStart *string `json:"ExtSvcStaticSubnetStart,omitempty"`
	// Value of ACI infrastructuere VLAN ID for the ACI fabric.
	InfraVlanId *int64 `json:"InfraVlanId,omitempty"`
	// Name of ACI L3Out network to be used for all Kubernetes clusters using this policy.
	L3OutNetworkName *string `json:"L3OutNetworkName,omitempty"`
	// Name of ACI L3Out policy to be used for all Kubernetes clusters using this policy.
	L3OutPolicyName *string `json:"L3OutPolicyName,omitempty"`
	// Tenant in ACI used by this L3Out and Common VRF.
	L3OutTenant *string `json:"L3OutTenant,omitempty"`
	// VMM domain within which Kubernetes clusters using this policy are nested.
	NestedVmmDomain *string `json:"NestedVmmDomain,omitempty"`
	// Start of range of ACI Node Service IP subnets to use by Kubernetes clusters using this ACI CNI policy This is used for the service graph which is used for ACI PBR based load balancing.
	NodeSvcSubnetStart *string `json:"NodeSvcSubnetStart,omitempty"`
	// Ending value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.
	NodeVlanRangeEnd *int64 `json:"NodeVlanRangeEnd,omitempty"`
	// Starting value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.
	NodeVlanRangeStart *int64 `json:"NodeVlanRangeStart,omitempty"`
	// Number of k8s clusters currently using this ACI CNI profile.
	NumberOfKubernetesClusters *int64 `json:"NumberOfKubernetesClusters,omitempty"`
	// Range of IP Multicast addresses to be used by the Opflex protocol for Kubernetes clusters using this policy.
	OpflexMulticastAddressRange *string `json:"OpflexMulticastAddressRange,omitempty"`
	// Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile.
	PodSubnetStart *string `json:"PodSubnetStart,omitempty"`
	// Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only.
	SvcSubnetStart *string `json:"SvcSubnetStart,omitempty"`
	// VRF (Virtual Routing and Forwarding) domain to be used within ACI fabric by all k8s clusters using this policy.
	Vrf *string `json:"Vrf,omitempty"`
	// An array of relationships to kubernetesAciCniTenantClusterAllocation resources.
	ClusterAciAllocations []KubernetesAciCniTenantClusterAllocationRelationship `json:"ClusterAciAllocations,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniProfileAllOf KubernetesAciCniProfileAllOf

// NewKubernetesAciCniProfileAllOf instantiates a new KubernetesAciCniProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniProfileAllOf(classId string, objectType string) *KubernetesAciCniProfileAllOf {
	this := KubernetesAciCniProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniProfileAllOfWithDefaults instantiates a new KubernetesAciCniProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniProfileAllOfWithDefaults() *KubernetesAciCniProfileAllOf {
	this := KubernetesAciCniProfileAllOf{}
	var classId string = "kubernetes.AciCniProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAaepName returns the AaepName field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetAaepName() string {
	if o == nil || o.AaepName == nil {
		var ret string
		return ret
	}
	return *o.AaepName
}

// GetAaepNameOk returns a tuple with the AaepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetAaepNameOk() (*string, bool) {
	if o == nil || o.AaepName == nil {
		return nil, false
	}
	return o.AaepName, true
}

// HasAaepName returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasAaepName() bool {
	if o != nil && o.AaepName != nil {
		return true
	}

	return false
}

// SetAaepName gets a reference to the given string and assigns it to the AaepName field.
func (o *KubernetesAciCniProfileAllOf) SetAaepName(v string) {
	o.AaepName = &v
}

// GetExtSvcDynSubnetStart returns the ExtSvcDynSubnetStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetExtSvcDynSubnetStart() string {
	if o == nil || o.ExtSvcDynSubnetStart == nil {
		var ret string
		return ret
	}
	return *o.ExtSvcDynSubnetStart
}

// GetExtSvcDynSubnetStartOk returns a tuple with the ExtSvcDynSubnetStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetExtSvcDynSubnetStartOk() (*string, bool) {
	if o == nil || o.ExtSvcDynSubnetStart == nil {
		return nil, false
	}
	return o.ExtSvcDynSubnetStart, true
}

// HasExtSvcDynSubnetStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasExtSvcDynSubnetStart() bool {
	if o != nil && o.ExtSvcDynSubnetStart != nil {
		return true
	}

	return false
}

// SetExtSvcDynSubnetStart gets a reference to the given string and assigns it to the ExtSvcDynSubnetStart field.
func (o *KubernetesAciCniProfileAllOf) SetExtSvcDynSubnetStart(v string) {
	o.ExtSvcDynSubnetStart = &v
}

// GetExtSvcStaticSubnetStart returns the ExtSvcStaticSubnetStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetExtSvcStaticSubnetStart() string {
	if o == nil || o.ExtSvcStaticSubnetStart == nil {
		var ret string
		return ret
	}
	return *o.ExtSvcStaticSubnetStart
}

// GetExtSvcStaticSubnetStartOk returns a tuple with the ExtSvcStaticSubnetStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetExtSvcStaticSubnetStartOk() (*string, bool) {
	if o == nil || o.ExtSvcStaticSubnetStart == nil {
		return nil, false
	}
	return o.ExtSvcStaticSubnetStart, true
}

// HasExtSvcStaticSubnetStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasExtSvcStaticSubnetStart() bool {
	if o != nil && o.ExtSvcStaticSubnetStart != nil {
		return true
	}

	return false
}

// SetExtSvcStaticSubnetStart gets a reference to the given string and assigns it to the ExtSvcStaticSubnetStart field.
func (o *KubernetesAciCniProfileAllOf) SetExtSvcStaticSubnetStart(v string) {
	o.ExtSvcStaticSubnetStart = &v
}

// GetInfraVlanId returns the InfraVlanId field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetInfraVlanId() int64 {
	if o == nil || o.InfraVlanId == nil {
		var ret int64
		return ret
	}
	return *o.InfraVlanId
}

// GetInfraVlanIdOk returns a tuple with the InfraVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetInfraVlanIdOk() (*int64, bool) {
	if o == nil || o.InfraVlanId == nil {
		return nil, false
	}
	return o.InfraVlanId, true
}

// HasInfraVlanId returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasInfraVlanId() bool {
	if o != nil && o.InfraVlanId != nil {
		return true
	}

	return false
}

// SetInfraVlanId gets a reference to the given int64 and assigns it to the InfraVlanId field.
func (o *KubernetesAciCniProfileAllOf) SetInfraVlanId(v int64) {
	o.InfraVlanId = &v
}

// GetL3OutNetworkName returns the L3OutNetworkName field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetL3OutNetworkName() string {
	if o == nil || o.L3OutNetworkName == nil {
		var ret string
		return ret
	}
	return *o.L3OutNetworkName
}

// GetL3OutNetworkNameOk returns a tuple with the L3OutNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetL3OutNetworkNameOk() (*string, bool) {
	if o == nil || o.L3OutNetworkName == nil {
		return nil, false
	}
	return o.L3OutNetworkName, true
}

// HasL3OutNetworkName returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasL3OutNetworkName() bool {
	if o != nil && o.L3OutNetworkName != nil {
		return true
	}

	return false
}

// SetL3OutNetworkName gets a reference to the given string and assigns it to the L3OutNetworkName field.
func (o *KubernetesAciCniProfileAllOf) SetL3OutNetworkName(v string) {
	o.L3OutNetworkName = &v
}

// GetL3OutPolicyName returns the L3OutPolicyName field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetL3OutPolicyName() string {
	if o == nil || o.L3OutPolicyName == nil {
		var ret string
		return ret
	}
	return *o.L3OutPolicyName
}

// GetL3OutPolicyNameOk returns a tuple with the L3OutPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetL3OutPolicyNameOk() (*string, bool) {
	if o == nil || o.L3OutPolicyName == nil {
		return nil, false
	}
	return o.L3OutPolicyName, true
}

// HasL3OutPolicyName returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasL3OutPolicyName() bool {
	if o != nil && o.L3OutPolicyName != nil {
		return true
	}

	return false
}

// SetL3OutPolicyName gets a reference to the given string and assigns it to the L3OutPolicyName field.
func (o *KubernetesAciCniProfileAllOf) SetL3OutPolicyName(v string) {
	o.L3OutPolicyName = &v
}

// GetL3OutTenant returns the L3OutTenant field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetL3OutTenant() string {
	if o == nil || o.L3OutTenant == nil {
		var ret string
		return ret
	}
	return *o.L3OutTenant
}

// GetL3OutTenantOk returns a tuple with the L3OutTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetL3OutTenantOk() (*string, bool) {
	if o == nil || o.L3OutTenant == nil {
		return nil, false
	}
	return o.L3OutTenant, true
}

// HasL3OutTenant returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasL3OutTenant() bool {
	if o != nil && o.L3OutTenant != nil {
		return true
	}

	return false
}

// SetL3OutTenant gets a reference to the given string and assigns it to the L3OutTenant field.
func (o *KubernetesAciCniProfileAllOf) SetL3OutTenant(v string) {
	o.L3OutTenant = &v
}

// GetNestedVmmDomain returns the NestedVmmDomain field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetNestedVmmDomain() string {
	if o == nil || o.NestedVmmDomain == nil {
		var ret string
		return ret
	}
	return *o.NestedVmmDomain
}

// GetNestedVmmDomainOk returns a tuple with the NestedVmmDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetNestedVmmDomainOk() (*string, bool) {
	if o == nil || o.NestedVmmDomain == nil {
		return nil, false
	}
	return o.NestedVmmDomain, true
}

// HasNestedVmmDomain returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasNestedVmmDomain() bool {
	if o != nil && o.NestedVmmDomain != nil {
		return true
	}

	return false
}

// SetNestedVmmDomain gets a reference to the given string and assigns it to the NestedVmmDomain field.
func (o *KubernetesAciCniProfileAllOf) SetNestedVmmDomain(v string) {
	o.NestedVmmDomain = &v
}

// GetNodeSvcSubnetStart returns the NodeSvcSubnetStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetNodeSvcSubnetStart() string {
	if o == nil || o.NodeSvcSubnetStart == nil {
		var ret string
		return ret
	}
	return *o.NodeSvcSubnetStart
}

// GetNodeSvcSubnetStartOk returns a tuple with the NodeSvcSubnetStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetNodeSvcSubnetStartOk() (*string, bool) {
	if o == nil || o.NodeSvcSubnetStart == nil {
		return nil, false
	}
	return o.NodeSvcSubnetStart, true
}

// HasNodeSvcSubnetStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasNodeSvcSubnetStart() bool {
	if o != nil && o.NodeSvcSubnetStart != nil {
		return true
	}

	return false
}

// SetNodeSvcSubnetStart gets a reference to the given string and assigns it to the NodeSvcSubnetStart field.
func (o *KubernetesAciCniProfileAllOf) SetNodeSvcSubnetStart(v string) {
	o.NodeSvcSubnetStart = &v
}

// GetNodeVlanRangeEnd returns the NodeVlanRangeEnd field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeEnd() int64 {
	if o == nil || o.NodeVlanRangeEnd == nil {
		var ret int64
		return ret
	}
	return *o.NodeVlanRangeEnd
}

// GetNodeVlanRangeEndOk returns a tuple with the NodeVlanRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeEndOk() (*int64, bool) {
	if o == nil || o.NodeVlanRangeEnd == nil {
		return nil, false
	}
	return o.NodeVlanRangeEnd, true
}

// HasNodeVlanRangeEnd returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasNodeVlanRangeEnd() bool {
	if o != nil && o.NodeVlanRangeEnd != nil {
		return true
	}

	return false
}

// SetNodeVlanRangeEnd gets a reference to the given int64 and assigns it to the NodeVlanRangeEnd field.
func (o *KubernetesAciCniProfileAllOf) SetNodeVlanRangeEnd(v int64) {
	o.NodeVlanRangeEnd = &v
}

// GetNodeVlanRangeStart returns the NodeVlanRangeStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeStart() int64 {
	if o == nil || o.NodeVlanRangeStart == nil {
		var ret int64
		return ret
	}
	return *o.NodeVlanRangeStart
}

// GetNodeVlanRangeStartOk returns a tuple with the NodeVlanRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeStartOk() (*int64, bool) {
	if o == nil || o.NodeVlanRangeStart == nil {
		return nil, false
	}
	return o.NodeVlanRangeStart, true
}

// HasNodeVlanRangeStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasNodeVlanRangeStart() bool {
	if o != nil && o.NodeVlanRangeStart != nil {
		return true
	}

	return false
}

// SetNodeVlanRangeStart gets a reference to the given int64 and assigns it to the NodeVlanRangeStart field.
func (o *KubernetesAciCniProfileAllOf) SetNodeVlanRangeStart(v int64) {
	o.NodeVlanRangeStart = &v
}

// GetNumberOfKubernetesClusters returns the NumberOfKubernetesClusters field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetNumberOfKubernetesClusters() int64 {
	if o == nil || o.NumberOfKubernetesClusters == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfKubernetesClusters
}

// GetNumberOfKubernetesClustersOk returns a tuple with the NumberOfKubernetesClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetNumberOfKubernetesClustersOk() (*int64, bool) {
	if o == nil || o.NumberOfKubernetesClusters == nil {
		return nil, false
	}
	return o.NumberOfKubernetesClusters, true
}

// HasNumberOfKubernetesClusters returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasNumberOfKubernetesClusters() bool {
	if o != nil && o.NumberOfKubernetesClusters != nil {
		return true
	}

	return false
}

// SetNumberOfKubernetesClusters gets a reference to the given int64 and assigns it to the NumberOfKubernetesClusters field.
func (o *KubernetesAciCniProfileAllOf) SetNumberOfKubernetesClusters(v int64) {
	o.NumberOfKubernetesClusters = &v
}

// GetOpflexMulticastAddressRange returns the OpflexMulticastAddressRange field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetOpflexMulticastAddressRange() string {
	if o == nil || o.OpflexMulticastAddressRange == nil {
		var ret string
		return ret
	}
	return *o.OpflexMulticastAddressRange
}

// GetOpflexMulticastAddressRangeOk returns a tuple with the OpflexMulticastAddressRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetOpflexMulticastAddressRangeOk() (*string, bool) {
	if o == nil || o.OpflexMulticastAddressRange == nil {
		return nil, false
	}
	return o.OpflexMulticastAddressRange, true
}

// HasOpflexMulticastAddressRange returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasOpflexMulticastAddressRange() bool {
	if o != nil && o.OpflexMulticastAddressRange != nil {
		return true
	}

	return false
}

// SetOpflexMulticastAddressRange gets a reference to the given string and assigns it to the OpflexMulticastAddressRange field.
func (o *KubernetesAciCniProfileAllOf) SetOpflexMulticastAddressRange(v string) {
	o.OpflexMulticastAddressRange = &v
}

// GetPodSubnetStart returns the PodSubnetStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetPodSubnetStart() string {
	if o == nil || o.PodSubnetStart == nil {
		var ret string
		return ret
	}
	return *o.PodSubnetStart
}

// GetPodSubnetStartOk returns a tuple with the PodSubnetStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetPodSubnetStartOk() (*string, bool) {
	if o == nil || o.PodSubnetStart == nil {
		return nil, false
	}
	return o.PodSubnetStart, true
}

// HasPodSubnetStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasPodSubnetStart() bool {
	if o != nil && o.PodSubnetStart != nil {
		return true
	}

	return false
}

// SetPodSubnetStart gets a reference to the given string and assigns it to the PodSubnetStart field.
func (o *KubernetesAciCniProfileAllOf) SetPodSubnetStart(v string) {
	o.PodSubnetStart = &v
}

// GetSvcSubnetStart returns the SvcSubnetStart field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetSvcSubnetStart() string {
	if o == nil || o.SvcSubnetStart == nil {
		var ret string
		return ret
	}
	return *o.SvcSubnetStart
}

// GetSvcSubnetStartOk returns a tuple with the SvcSubnetStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetSvcSubnetStartOk() (*string, bool) {
	if o == nil || o.SvcSubnetStart == nil {
		return nil, false
	}
	return o.SvcSubnetStart, true
}

// HasSvcSubnetStart returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasSvcSubnetStart() bool {
	if o != nil && o.SvcSubnetStart != nil {
		return true
	}

	return false
}

// SetSvcSubnetStart gets a reference to the given string and assigns it to the SvcSubnetStart field.
func (o *KubernetesAciCniProfileAllOf) SetSvcSubnetStart(v string) {
	o.SvcSubnetStart = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetVrf() string {
	if o == nil || o.Vrf == nil {
		var ret string
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetVrfOk() (*string, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given string and assigns it to the Vrf field.
func (o *KubernetesAciCniProfileAllOf) SetVrf(v string) {
	o.Vrf = &v
}

// GetClusterAciAllocations returns the ClusterAciAllocations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniProfileAllOf) GetClusterAciAllocations() []KubernetesAciCniTenantClusterAllocationRelationship {
	if o == nil {
		var ret []KubernetesAciCniTenantClusterAllocationRelationship
		return ret
	}
	return o.ClusterAciAllocations
}

// GetClusterAciAllocationsOk returns a tuple with the ClusterAciAllocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniProfileAllOf) GetClusterAciAllocationsOk() (*[]KubernetesAciCniTenantClusterAllocationRelationship, bool) {
	if o == nil || o.ClusterAciAllocations == nil {
		return nil, false
	}
	return &o.ClusterAciAllocations, true
}

// HasClusterAciAllocations returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasClusterAciAllocations() bool {
	if o != nil && o.ClusterAciAllocations != nil {
		return true
	}

	return false
}

// SetClusterAciAllocations gets a reference to the given []KubernetesAciCniTenantClusterAllocationRelationship and assigns it to the ClusterAciAllocations field.
func (o *KubernetesAciCniProfileAllOf) SetClusterAciAllocations(v []KubernetesAciCniTenantClusterAllocationRelationship) {
	o.ClusterAciAllocations = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniProfileAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniProfileAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesAciCniProfileAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesAciCniProfileAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniProfileAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesAciCniProfileAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesAciCniProfileAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesAciCniProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AaepName != nil {
		toSerialize["AaepName"] = o.AaepName
	}
	if o.ExtSvcDynSubnetStart != nil {
		toSerialize["ExtSvcDynSubnetStart"] = o.ExtSvcDynSubnetStart
	}
	if o.ExtSvcStaticSubnetStart != nil {
		toSerialize["ExtSvcStaticSubnetStart"] = o.ExtSvcStaticSubnetStart
	}
	if o.InfraVlanId != nil {
		toSerialize["InfraVlanId"] = o.InfraVlanId
	}
	if o.L3OutNetworkName != nil {
		toSerialize["L3OutNetworkName"] = o.L3OutNetworkName
	}
	if o.L3OutPolicyName != nil {
		toSerialize["L3OutPolicyName"] = o.L3OutPolicyName
	}
	if o.L3OutTenant != nil {
		toSerialize["L3OutTenant"] = o.L3OutTenant
	}
	if o.NestedVmmDomain != nil {
		toSerialize["NestedVmmDomain"] = o.NestedVmmDomain
	}
	if o.NodeSvcSubnetStart != nil {
		toSerialize["NodeSvcSubnetStart"] = o.NodeSvcSubnetStart
	}
	if o.NodeVlanRangeEnd != nil {
		toSerialize["NodeVlanRangeEnd"] = o.NodeVlanRangeEnd
	}
	if o.NodeVlanRangeStart != nil {
		toSerialize["NodeVlanRangeStart"] = o.NodeVlanRangeStart
	}
	if o.NumberOfKubernetesClusters != nil {
		toSerialize["NumberOfKubernetesClusters"] = o.NumberOfKubernetesClusters
	}
	if o.OpflexMulticastAddressRange != nil {
		toSerialize["OpflexMulticastAddressRange"] = o.OpflexMulticastAddressRange
	}
	if o.PodSubnetStart != nil {
		toSerialize["PodSubnetStart"] = o.PodSubnetStart
	}
	if o.SvcSubnetStart != nil {
		toSerialize["SvcSubnetStart"] = o.SvcSubnetStart
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}
	if o.ClusterAciAllocations != nil {
		toSerialize["ClusterAciAllocations"] = o.ClusterAciAllocations
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAciCniProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAciCniProfileAllOf := _KubernetesAciCniProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAciCniProfileAllOf); err == nil {
		*o = KubernetesAciCniProfileAllOf(varKubernetesAciCniProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AaepName")
		delete(additionalProperties, "ExtSvcDynSubnetStart")
		delete(additionalProperties, "ExtSvcStaticSubnetStart")
		delete(additionalProperties, "InfraVlanId")
		delete(additionalProperties, "L3OutNetworkName")
		delete(additionalProperties, "L3OutPolicyName")
		delete(additionalProperties, "L3OutTenant")
		delete(additionalProperties, "NestedVmmDomain")
		delete(additionalProperties, "NodeSvcSubnetStart")
		delete(additionalProperties, "NodeVlanRangeEnd")
		delete(additionalProperties, "NodeVlanRangeStart")
		delete(additionalProperties, "NumberOfKubernetesClusters")
		delete(additionalProperties, "OpflexMulticastAddressRange")
		delete(additionalProperties, "PodSubnetStart")
		delete(additionalProperties, "SvcSubnetStart")
		delete(additionalProperties, "Vrf")
		delete(additionalProperties, "ClusterAciAllocations")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAciCniProfileAllOf struct {
	value *KubernetesAciCniProfileAllOf
	isSet bool
}

func (v NullableKubernetesAciCniProfileAllOf) Get() *KubernetesAciCniProfileAllOf {
	return v.value
}

func (v *NullableKubernetesAciCniProfileAllOf) Set(val *KubernetesAciCniProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniProfileAllOf(val *KubernetesAciCniProfileAllOf) *NullableKubernetesAciCniProfileAllOf {
	return &NullableKubernetesAciCniProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAciCniProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
