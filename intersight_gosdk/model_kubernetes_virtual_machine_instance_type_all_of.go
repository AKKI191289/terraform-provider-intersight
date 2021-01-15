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

// KubernetesVirtualMachineInstanceTypeAllOf Definition of the list of properties defined in 'kubernetes.VirtualMachineInstanceType', excluding properties defined in parent classes.
type KubernetesVirtualMachineInstanceTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of CPUs allocated to virtual machine.
	Cpu *int64 `json:"Cpu,omitempty"`
	// Ephemeral disk capacity to be provided with units example - 10Gi.
	DiskSize *string `json:"DiskSize,omitempty"`
	// Virtual machine memory defined in mebibytes (MiB).
	Memory       *int64                                `json:"Memory,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources.
	Profiles             []KubernetesVirtualMachineInfrastructureProviderRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineInstanceTypeAllOf KubernetesVirtualMachineInstanceTypeAllOf

// NewKubernetesVirtualMachineInstanceTypeAllOf instantiates a new KubernetesVirtualMachineInstanceTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineInstanceTypeAllOf(classId string, objectType string) *KubernetesVirtualMachineInstanceTypeAllOf {
	this := KubernetesVirtualMachineInstanceTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cpu int64 = 4
	this.Cpu = &cpu
	var memory int64 = 16384
	this.Memory = &memory
	return &this
}

// NewKubernetesVirtualMachineInstanceTypeAllOfWithDefaults instantiates a new KubernetesVirtualMachineInstanceTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineInstanceTypeAllOfWithDefaults() *KubernetesVirtualMachineInstanceTypeAllOf {
	this := KubernetesVirtualMachineInstanceTypeAllOf{}
	var classId string = "kubernetes.VirtualMachineInstanceType"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInstanceType"
	this.ObjectType = objectType
	var cpu int64 = 4
	this.Cpu = &cpu
	var memory int64 = 16384
	this.Memory = &memory
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetCpu() int64 {
	if o == nil || o.Cpu == nil {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetCpuOk() (*int64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetCpu(v int64) {
	o.Cpu = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetDiskSize() string {
	if o == nil || o.DiskSize == nil {
		var ret string
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetDiskSizeOk() (*string, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given string and assigns it to the DiskSize field.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetDiskSize(v string) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetMemory(v int64) {
	o.Memory = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship {
	if o == nil {
		var ret []KubernetesVirtualMachineInfrastructureProviderRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInstanceTypeAllOf) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesVirtualMachineInfrastructureProviderRelationship and assigns it to the Profiles field.
func (o *KubernetesVirtualMachineInstanceTypeAllOf) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship) {
	o.Profiles = v
}

func (o KubernetesVirtualMachineInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.DiskSize != nil {
		toSerialize["DiskSize"] = o.DiskSize
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineInstanceTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesVirtualMachineInstanceTypeAllOf := _KubernetesVirtualMachineInstanceTypeAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInstanceTypeAllOf); err == nil {
		*o = KubernetesVirtualMachineInstanceTypeAllOf(varKubernetesVirtualMachineInstanceTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cpu")
		delete(additionalProperties, "DiskSize")
		delete(additionalProperties, "Memory")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesVirtualMachineInstanceTypeAllOf struct {
	value *KubernetesVirtualMachineInstanceTypeAllOf
	isSet bool
}

func (v NullableKubernetesVirtualMachineInstanceTypeAllOf) Get() *KubernetesVirtualMachineInstanceTypeAllOf {
	return v.value
}

func (v *NullableKubernetesVirtualMachineInstanceTypeAllOf) Set(val *KubernetesVirtualMachineInstanceTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineInstanceTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineInstanceTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineInstanceTypeAllOf(val *KubernetesVirtualMachineInstanceTypeAllOf) *NullableKubernetesVirtualMachineInstanceTypeAllOf {
	return &NullableKubernetesVirtualMachineInstanceTypeAllOf{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineInstanceTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
