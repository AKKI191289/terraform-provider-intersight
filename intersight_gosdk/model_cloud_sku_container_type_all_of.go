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

// CloudSkuContainerTypeAllOf Definition of the list of properties defined in 'cloud.SkuContainerType', excluding properties defined in parent classes.
type CloudSkuContainerTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The CpuUnit. Will be MILLI_CPU for containers. * `VIRTUAL_CPU` - The CPU unit used for virtual machines. * `MILLI_CPU` - The CPU unit used by containers.
	CpuUnit *string `json:"CpuUnit,omitempty"`
	// The number of CPUs in this instance type.
	NumOfCpus            *int64 `json:"NumOfCpus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuContainerTypeAllOf CloudSkuContainerTypeAllOf

// NewCloudSkuContainerTypeAllOf instantiates a new CloudSkuContainerTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuContainerTypeAllOf(classId string, objectType string) *CloudSkuContainerTypeAllOf {
	this := CloudSkuContainerTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	return &this
}

// NewCloudSkuContainerTypeAllOfWithDefaults instantiates a new CloudSkuContainerTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuContainerTypeAllOfWithDefaults() *CloudSkuContainerTypeAllOf {
	this := CloudSkuContainerTypeAllOf{}
	var classId string = "cloud.SkuContainerType"
	this.ClassId = classId
	var objectType string = "cloud.SkuContainerType"
	this.ObjectType = objectType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuContainerTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuContainerTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuContainerTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuContainerTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuContainerTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuContainerTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuUnit returns the CpuUnit field value if set, zero value otherwise.
func (o *CloudSkuContainerTypeAllOf) GetCpuUnit() string {
	if o == nil || o.CpuUnit == nil {
		var ret string
		return ret
	}
	return *o.CpuUnit
}

// GetCpuUnitOk returns a tuple with the CpuUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuContainerTypeAllOf) GetCpuUnitOk() (*string, bool) {
	if o == nil || o.CpuUnit == nil {
		return nil, false
	}
	return o.CpuUnit, true
}

// HasCpuUnit returns a boolean if a field has been set.
func (o *CloudSkuContainerTypeAllOf) HasCpuUnit() bool {
	if o != nil && o.CpuUnit != nil {
		return true
	}

	return false
}

// SetCpuUnit gets a reference to the given string and assigns it to the CpuUnit field.
func (o *CloudSkuContainerTypeAllOf) SetCpuUnit(v string) {
	o.CpuUnit = &v
}

// GetNumOfCpus returns the NumOfCpus field value if set, zero value otherwise.
func (o *CloudSkuContainerTypeAllOf) GetNumOfCpus() int64 {
	if o == nil || o.NumOfCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumOfCpus
}

// GetNumOfCpusOk returns a tuple with the NumOfCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuContainerTypeAllOf) GetNumOfCpusOk() (*int64, bool) {
	if o == nil || o.NumOfCpus == nil {
		return nil, false
	}
	return o.NumOfCpus, true
}

// HasNumOfCpus returns a boolean if a field has been set.
func (o *CloudSkuContainerTypeAllOf) HasNumOfCpus() bool {
	if o != nil && o.NumOfCpus != nil {
		return true
	}

	return false
}

// SetNumOfCpus gets a reference to the given int64 and assigns it to the NumOfCpus field.
func (o *CloudSkuContainerTypeAllOf) SetNumOfCpus(v int64) {
	o.NumOfCpus = &v
}

func (o CloudSkuContainerTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuUnit != nil {
		toSerialize["CpuUnit"] = o.CpuUnit
	}
	if o.NumOfCpus != nil {
		toSerialize["NumOfCpus"] = o.NumOfCpus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuContainerTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudSkuContainerTypeAllOf := _CloudSkuContainerTypeAllOf{}

	if err = json.Unmarshal(bytes, &varCloudSkuContainerTypeAllOf); err == nil {
		*o = CloudSkuContainerTypeAllOf(varCloudSkuContainerTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuUnit")
		delete(additionalProperties, "NumOfCpus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSkuContainerTypeAllOf struct {
	value *CloudSkuContainerTypeAllOf
	isSet bool
}

func (v NullableCloudSkuContainerTypeAllOf) Get() *CloudSkuContainerTypeAllOf {
	return v.value
}

func (v *NullableCloudSkuContainerTypeAllOf) Set(val *CloudSkuContainerTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuContainerTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuContainerTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuContainerTypeAllOf(val *CloudSkuContainerTypeAllOf) *NullableCloudSkuContainerTypeAllOf {
	return &NullableCloudSkuContainerTypeAllOf{value: val, isSet: true}
}

func (v NullableCloudSkuContainerTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuContainerTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
