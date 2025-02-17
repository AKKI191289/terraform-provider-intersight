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

// KubernetesObjectMetaAllOf Definition of the list of properties defined in 'kubernetes.ObjectMeta', excluding properties defined in parent classes.
type KubernetesObjectMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists.
	CreationTimestamp *string `json:"CreationTimestamp,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Read-only.
	Name *string `json:"Name,omitempty"`
	// Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Read-only.
	Namespace *string `json:"Namespace,omitempty"`
	// Specific resourceVersion to which this reference is made, if any.
	ResourceVersion *string `json:"ResourceVersion,omitempty"`
	// UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesObjectMetaAllOf KubernetesObjectMetaAllOf

// NewKubernetesObjectMetaAllOf instantiates a new KubernetesObjectMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesObjectMetaAllOf(classId string, objectType string) *KubernetesObjectMetaAllOf {
	this := KubernetesObjectMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesObjectMetaAllOfWithDefaults instantiates a new KubernetesObjectMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesObjectMetaAllOfWithDefaults() *KubernetesObjectMetaAllOf {
	this := KubernetesObjectMetaAllOf{}
	var classId string = "kubernetes.ObjectMeta"
	this.ClassId = classId
	var objectType string = "kubernetes.ObjectMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesObjectMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesObjectMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesObjectMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesObjectMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *KubernetesObjectMetaAllOf) GetCreationTimestamp() string {
	if o == nil || o.CreationTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetCreationTimestampOk() (*string, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *KubernetesObjectMetaAllOf) HasCreationTimestamp() bool {
	if o != nil && o.CreationTimestamp != nil {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *KubernetesObjectMetaAllOf) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesObjectMetaAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesObjectMetaAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesObjectMetaAllOf) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *KubernetesObjectMetaAllOf) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubernetesObjectMetaAllOf) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *KubernetesObjectMetaAllOf) SetNamespace(v string) {
	o.Namespace = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *KubernetesObjectMetaAllOf) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *KubernetesObjectMetaAllOf) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *KubernetesObjectMetaAllOf) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *KubernetesObjectMetaAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMetaAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *KubernetesObjectMetaAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *KubernetesObjectMetaAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o KubernetesObjectMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CreationTimestamp != nil {
		toSerialize["CreationTimestamp"] = o.CreationTimestamp
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.ResourceVersion != nil {
		toSerialize["ResourceVersion"] = o.ResourceVersion
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesObjectMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesObjectMetaAllOf := _KubernetesObjectMetaAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesObjectMetaAllOf); err == nil {
		*o = KubernetesObjectMetaAllOf(varKubernetesObjectMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreationTimestamp")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "ResourceVersion")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesObjectMetaAllOf struct {
	value *KubernetesObjectMetaAllOf
	isSet bool
}

func (v NullableKubernetesObjectMetaAllOf) Get() *KubernetesObjectMetaAllOf {
	return v.value
}

func (v *NullableKubernetesObjectMetaAllOf) Set(val *KubernetesObjectMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesObjectMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesObjectMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesObjectMetaAllOf(val *KubernetesObjectMetaAllOf) *NullableKubernetesObjectMetaAllOf {
	return &NullableKubernetesObjectMetaAllOf{value: val, isSet: true}
}

func (v NullableKubernetesObjectMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesObjectMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
