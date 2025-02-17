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

// FirmwareComponentDescriptor Descriptor to uniquely identify each component.
type FirmwareComponentDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The brand string of the endpoint for which this capability information is applicable.
	BrandString *string `json:"BrandString,omitempty"`
	// The label type for the component.
	Label *string `json:"Label,omitempty"`
	// The revision for the component.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentDescriptor FirmwareComponentDescriptor

// NewFirmwareComponentDescriptor instantiates a new FirmwareComponentDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentDescriptor(classId string, objectType string) *FirmwareComponentDescriptor {
	this := FirmwareComponentDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareComponentDescriptorWithDefaults instantiates a new FirmwareComponentDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentDescriptorWithDefaults() *FirmwareComponentDescriptor {
	this := FirmwareComponentDescriptor{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBrandString returns the BrandString field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetBrandString() string {
	if o == nil || o.BrandString == nil {
		var ret string
		return ret
	}
	return *o.BrandString
}

// GetBrandStringOk returns a tuple with the BrandString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetBrandStringOk() (*string, bool) {
	if o == nil || o.BrandString == nil {
		return nil, false
	}
	return o.BrandString, true
}

// HasBrandString returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasBrandString() bool {
	if o != nil && o.BrandString != nil {
		return true
	}

	return false
}

// SetBrandString gets a reference to the given string and assigns it to the BrandString field.
func (o *FirmwareComponentDescriptor) SetBrandString(v string) {
	o.BrandString = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FirmwareComponentDescriptor) SetLabel(v string) {
	o.Label = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *FirmwareComponentDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o FirmwareComponentDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BrandString != nil {
		toSerialize["BrandString"] = o.BrandString
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareComponentDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareComponentDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The brand string of the endpoint for which this capability information is applicable.
		BrandString *string `json:"BrandString,omitempty"`
		// The label type for the component.
		Label *string `json:"Label,omitempty"`
		// The revision for the component.
		Revision *string `json:"Revision,omitempty"`
	}

	varFirmwareComponentDescriptorWithoutEmbeddedStruct := FirmwareComponentDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareComponentDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareComponentDescriptor := _FirmwareComponentDescriptor{}
		varFirmwareComponentDescriptor.ClassId = varFirmwareComponentDescriptorWithoutEmbeddedStruct.ClassId
		varFirmwareComponentDescriptor.ObjectType = varFirmwareComponentDescriptorWithoutEmbeddedStruct.ObjectType
		varFirmwareComponentDescriptor.BrandString = varFirmwareComponentDescriptorWithoutEmbeddedStruct.BrandString
		varFirmwareComponentDescriptor.Label = varFirmwareComponentDescriptorWithoutEmbeddedStruct.Label
		varFirmwareComponentDescriptor.Revision = varFirmwareComponentDescriptorWithoutEmbeddedStruct.Revision
		*o = FirmwareComponentDescriptor(varFirmwareComponentDescriptor)
	} else {
		return err
	}

	varFirmwareComponentDescriptor := _FirmwareComponentDescriptor{}

	err = json.Unmarshal(bytes, &varFirmwareComponentDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varFirmwareComponentDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BrandString")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Revision")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableFirmwareComponentDescriptor struct {
	value *FirmwareComponentDescriptor
	isSet bool
}

func (v NullableFirmwareComponentDescriptor) Get() *FirmwareComponentDescriptor {
	return v.value
}

func (v *NullableFirmwareComponentDescriptor) Set(val *FirmwareComponentDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentDescriptor(val *FirmwareComponentDescriptor) *NullableFirmwareComponentDescriptor {
	return &NullableFirmwareComponentDescriptor{value: val, isSet: true}
}

func (v NullableFirmwareComponentDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
