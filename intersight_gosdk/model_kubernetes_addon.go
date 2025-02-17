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

// KubernetesAddon An Addon that can be deployed to a Kubernetes cluster.
type KubernetesAddon struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                               `json:"ObjectType"`
	AddonConfiguration NullableKubernetesAddonConfiguration `json:"AddonConfiguration,omitempty"`
	AddonPolicy        *MoMoRef                             `json:"AddonPolicy,omitempty"`
	// Name of addon to be installed on a Kubernetes cluster.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddon KubernetesAddon

// NewKubernetesAddon instantiates a new KubernetesAddon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddon(classId string, objectType string) *KubernetesAddon {
	this := KubernetesAddon{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonWithDefaults instantiates a new KubernetesAddon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonWithDefaults() *KubernetesAddon {
	this := KubernetesAddon{}
	var classId string = "kubernetes.Addon"
	this.ClassId = classId
	var objectType string = "kubernetes.Addon"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddon) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddon) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddon) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddon) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddon) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddon) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddonConfiguration returns the AddonConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddon) GetAddonConfiguration() KubernetesAddonConfiguration {
	if o == nil || o.AddonConfiguration.Get() == nil {
		var ret KubernetesAddonConfiguration
		return ret
	}
	return *o.AddonConfiguration.Get()
}

// GetAddonConfigurationOk returns a tuple with the AddonConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddon) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddonConfiguration.Get(), o.AddonConfiguration.IsSet()
}

// HasAddonConfiguration returns a boolean if a field has been set.
func (o *KubernetesAddon) HasAddonConfiguration() bool {
	if o != nil && o.AddonConfiguration.IsSet() {
		return true
	}

	return false
}

// SetAddonConfiguration gets a reference to the given NullableKubernetesAddonConfiguration and assigns it to the AddonConfiguration field.
func (o *KubernetesAddon) SetAddonConfiguration(v KubernetesAddonConfiguration) {
	o.AddonConfiguration.Set(&v)
}

// SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil
func (o *KubernetesAddon) SetAddonConfigurationNil() {
	o.AddonConfiguration.Set(nil)
}

// UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
func (o *KubernetesAddon) UnsetAddonConfiguration() {
	o.AddonConfiguration.Unset()
}

// GetAddonPolicy returns the AddonPolicy field value if set, zero value otherwise.
func (o *KubernetesAddon) GetAddonPolicy() MoMoRef {
	if o == nil || o.AddonPolicy == nil {
		var ret MoMoRef
		return ret
	}
	return *o.AddonPolicy
}

// GetAddonPolicyOk returns a tuple with the AddonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddon) GetAddonPolicyOk() (*MoMoRef, bool) {
	if o == nil || o.AddonPolicy == nil {
		return nil, false
	}
	return o.AddonPolicy, true
}

// HasAddonPolicy returns a boolean if a field has been set.
func (o *KubernetesAddon) HasAddonPolicy() bool {
	if o != nil && o.AddonPolicy != nil {
		return true
	}

	return false
}

// SetAddonPolicy gets a reference to the given MoMoRef and assigns it to the AddonPolicy field.
func (o *KubernetesAddon) SetAddonPolicy(v MoMoRef) {
	o.AddonPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesAddon) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddon) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesAddon) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesAddon) SetName(v string) {
	o.Name = &v
}

func (o KubernetesAddon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddonConfiguration.IsSet() {
		toSerialize["AddonConfiguration"] = o.AddonConfiguration.Get()
	}
	if o.AddonPolicy != nil {
		toSerialize["AddonPolicy"] = o.AddonPolicy
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddon) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAddonWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                               `json:"ObjectType"`
		AddonConfiguration NullableKubernetesAddonConfiguration `json:"AddonConfiguration,omitempty"`
		AddonPolicy        *MoMoRef                             `json:"AddonPolicy,omitempty"`
		// Name of addon to be installed on a Kubernetes cluster.
		Name *string `json:"Name,omitempty"`
	}

	varKubernetesAddonWithoutEmbeddedStruct := KubernetesAddonWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAddonWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddon := _KubernetesAddon{}
		varKubernetesAddon.ClassId = varKubernetesAddonWithoutEmbeddedStruct.ClassId
		varKubernetesAddon.ObjectType = varKubernetesAddonWithoutEmbeddedStruct.ObjectType
		varKubernetesAddon.AddonConfiguration = varKubernetesAddonWithoutEmbeddedStruct.AddonConfiguration
		varKubernetesAddon.AddonPolicy = varKubernetesAddonWithoutEmbeddedStruct.AddonPolicy
		varKubernetesAddon.Name = varKubernetesAddonWithoutEmbeddedStruct.Name
		*o = KubernetesAddon(varKubernetesAddon)
	} else {
		return err
	}

	varKubernetesAddon := _KubernetesAddon{}

	err = json.Unmarshal(bytes, &varKubernetesAddon)
	if err == nil {
		o.MoBaseComplexType = varKubernetesAddon.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddonConfiguration")
		delete(additionalProperties, "AddonPolicy")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableKubernetesAddon struct {
	value *KubernetesAddon
	isSet bool
}

func (v NullableKubernetesAddon) Get() *KubernetesAddon {
	return v.value
}

func (v *NullableKubernetesAddon) Set(val *KubernetesAddon) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddon) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddon(val *KubernetesAddon) *NullableKubernetesAddon {
	return &NullableKubernetesAddon{value: val, isSet: true}
}

func (v NullableKubernetesAddon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
