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

// KubernetesVirtualMachineNodeProfileAllOf Definition of the list of properties defined in 'kubernetes.VirtualMachineNodeProfile', excluding properties defined in parent classes.
type KubernetesVirtualMachineNodeProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to ippoolIpLease resources.
	IpAddresses          []IppoolIpLeaseRelationship               `json:"IpAddresses,omitempty"`
	NodeIp               *IppoolIpLeaseRelationship                `json:"NodeIp,omitempty"`
	VirtualMachine       *VirtualizationVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineNodeProfileAllOf KubernetesVirtualMachineNodeProfileAllOf

// NewKubernetesVirtualMachineNodeProfileAllOf instantiates a new KubernetesVirtualMachineNodeProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineNodeProfileAllOf(classId string, objectType string) *KubernetesVirtualMachineNodeProfileAllOf {
	this := KubernetesVirtualMachineNodeProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVirtualMachineNodeProfileAllOfWithDefaults instantiates a new KubernetesVirtualMachineNodeProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineNodeProfileAllOfWithDefaults() *KubernetesVirtualMachineNodeProfileAllOf {
	this := KubernetesVirtualMachineNodeProfileAllOf{}
	var classId string = "kubernetes.VirtualMachineNodeProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineNodeProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineNodeProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineNodeProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetIpAddresses() []IppoolIpLeaseRelationship {
	if o == nil {
		var ret []IppoolIpLeaseRelationship
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetIpAddressesOk() (*[]IppoolIpLeaseRelationship, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IppoolIpLeaseRelationship and assigns it to the IpAddresses field.
func (o *KubernetesVirtualMachineNodeProfileAllOf) SetIpAddresses(v []IppoolIpLeaseRelationship) {
	o.IpAddresses = v
}

// GetNodeIp returns the NodeIp field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetNodeIp() IppoolIpLeaseRelationship {
	if o == nil || o.NodeIp == nil {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.NodeIp
}

// GetNodeIpOk returns a tuple with the NodeIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetNodeIpOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil || o.NodeIp == nil {
		return nil, false
	}
	return o.NodeIp, true
}

// HasNodeIp returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) HasNodeIp() bool {
	if o != nil && o.NodeIp != nil {
		return true
	}

	return false
}

// SetNodeIp gets a reference to the given IppoolIpLeaseRelationship and assigns it to the NodeIp field.
func (o *KubernetesVirtualMachineNodeProfileAllOf) SetNodeIp(v IppoolIpLeaseRelationship) {
	o.NodeIp = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetVirtualMachine() VirtualizationVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) GetVirtualMachineOk() (*VirtualizationVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfileAllOf) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *KubernetesVirtualMachineNodeProfileAllOf) SetVirtualMachine(v VirtualizationVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o KubernetesVirtualMachineNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}
	if o.NodeIp != nil {
		toSerialize["NodeIp"] = o.NodeIp
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineNodeProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesVirtualMachineNodeProfileAllOf := _KubernetesVirtualMachineNodeProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesVirtualMachineNodeProfileAllOf); err == nil {
		*o = KubernetesVirtualMachineNodeProfileAllOf(varKubernetesVirtualMachineNodeProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpAddresses")
		delete(additionalProperties, "NodeIp")
		delete(additionalProperties, "VirtualMachine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesVirtualMachineNodeProfileAllOf struct {
	value *KubernetesVirtualMachineNodeProfileAllOf
	isSet bool
}

func (v NullableKubernetesVirtualMachineNodeProfileAllOf) Get() *KubernetesVirtualMachineNodeProfileAllOf {
	return v.value
}

func (v *NullableKubernetesVirtualMachineNodeProfileAllOf) Set(val *KubernetesVirtualMachineNodeProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineNodeProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineNodeProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineNodeProfileAllOf(val *KubernetesVirtualMachineNodeProfileAllOf) *NullableKubernetesVirtualMachineNodeProfileAllOf {
	return &NullableKubernetesVirtualMachineNodeProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineNodeProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
