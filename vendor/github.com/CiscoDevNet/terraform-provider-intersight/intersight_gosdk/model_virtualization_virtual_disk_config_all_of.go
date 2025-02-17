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

// VirtualizationVirtualDiskConfigAllOf Definition of the list of properties defined in 'virtualization.VirtualDiskConfig', excluding properties defined in parent classes.
type VirtualizationVirtualDiskConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk capacity to be provided with units example - 10Gi.
	Capacity *string `json:"Capacity,omitempty"`
	// File mode of the disk, example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk.
	Mode *string `json:"Mode,omitempty"`
	// Base64 encoded CA certificates of the https source to check against.
	SourceCerts *string `json:"SourceCerts,omitempty"`
	// Source disk name from where the clone is done.
	SourceDiskToClone *string `json:"SourceDiskToClone,omitempty"`
	// Disk image source for the virtual machine.
	SourceFilePath       *string `json:"SourceFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualDiskConfigAllOf VirtualizationVirtualDiskConfigAllOf

// NewVirtualizationVirtualDiskConfigAllOf instantiates a new VirtualizationVirtualDiskConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualDiskConfigAllOf(classId string, objectType string) *VirtualizationVirtualDiskConfigAllOf {
	this := VirtualizationVirtualDiskConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewVirtualizationVirtualDiskConfigAllOfWithDefaults instantiates a new VirtualizationVirtualDiskConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualDiskConfigAllOfWithDefaults() *VirtualizationVirtualDiskConfigAllOf {
	this := VirtualizationVirtualDiskConfigAllOf{}
	var classId string = "virtualization.VirtualDiskConfig"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualDiskConfig"
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualDiskConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualDiskConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualDiskConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualDiskConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfigAllOf) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *VirtualizationVirtualDiskConfigAllOf) SetCapacity(v string) {
	o.Capacity = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfigAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VirtualizationVirtualDiskConfigAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetSourceCerts returns the SourceCerts field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceCerts() string {
	if o == nil || o.SourceCerts == nil {
		var ret string
		return ret
	}
	return *o.SourceCerts
}

// GetSourceCertsOk returns a tuple with the SourceCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceCertsOk() (*string, bool) {
	if o == nil || o.SourceCerts == nil {
		return nil, false
	}
	return o.SourceCerts, true
}

// HasSourceCerts returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) HasSourceCerts() bool {
	if o != nil && o.SourceCerts != nil {
		return true
	}

	return false
}

// SetSourceCerts gets a reference to the given string and assigns it to the SourceCerts field.
func (o *VirtualizationVirtualDiskConfigAllOf) SetSourceCerts(v string) {
	o.SourceCerts = &v
}

// GetSourceDiskToClone returns the SourceDiskToClone field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceDiskToClone() string {
	if o == nil || o.SourceDiskToClone == nil {
		var ret string
		return ret
	}
	return *o.SourceDiskToClone
}

// GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceDiskToCloneOk() (*string, bool) {
	if o == nil || o.SourceDiskToClone == nil {
		return nil, false
	}
	return o.SourceDiskToClone, true
}

// HasSourceDiskToClone returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) HasSourceDiskToClone() bool {
	if o != nil && o.SourceDiskToClone != nil {
		return true
	}

	return false
}

// SetSourceDiskToClone gets a reference to the given string and assigns it to the SourceDiskToClone field.
func (o *VirtualizationVirtualDiskConfigAllOf) SetSourceDiskToClone(v string) {
	o.SourceDiskToClone = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfigAllOf) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *VirtualizationVirtualDiskConfigAllOf) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

func (o VirtualizationVirtualDiskConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.SourceCerts != nil {
		toSerialize["SourceCerts"] = o.SourceCerts
	}
	if o.SourceDiskToClone != nil {
		toSerialize["SourceDiskToClone"] = o.SourceDiskToClone
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualDiskConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVirtualDiskConfigAllOf := _VirtualizationVirtualDiskConfigAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVirtualDiskConfigAllOf); err == nil {
		*o = VirtualizationVirtualDiskConfigAllOf(varVirtualizationVirtualDiskConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "SourceCerts")
		delete(additionalProperties, "SourceDiskToClone")
		delete(additionalProperties, "SourceFilePath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVirtualDiskConfigAllOf struct {
	value *VirtualizationVirtualDiskConfigAllOf
	isSet bool
}

func (v NullableVirtualizationVirtualDiskConfigAllOf) Get() *VirtualizationVirtualDiskConfigAllOf {
	return v.value
}

func (v *NullableVirtualizationVirtualDiskConfigAllOf) Set(val *VirtualizationVirtualDiskConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualDiskConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualDiskConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualDiskConfigAllOf(val *VirtualizationVirtualDiskConfigAllOf) *NullableVirtualizationVirtualDiskConfigAllOf {
	return &NullableVirtualizationVirtualDiskConfigAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualDiskConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualDiskConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
