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

// VirtualizationProductInfo The details of the product offering.
type VirtualizationProductInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Commercial product name. For example, VMware ESXi.
	ProductName *string `json:"ProductName,omitempty"`
	// Product name provided by the vendor. For example, embeddedEsx.
	ProductType *string `json:"ProductType,omitempty"`
	// Commercial vendor name. For example, VMware Inc.
	ProductVendor *string `json:"ProductVendor,omitempty"`
	// Hypervisor version running on the system.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationProductInfo VirtualizationProductInfo

// NewVirtualizationProductInfo instantiates a new VirtualizationProductInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationProductInfo(classId string, objectType string) *VirtualizationProductInfo {
	this := VirtualizationProductInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationProductInfoWithDefaults instantiates a new VirtualizationProductInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationProductInfoWithDefaults() *VirtualizationProductInfo {
	this := VirtualizationProductInfo{}
	var classId string = "virtualization.ProductInfo"
	this.ClassId = classId
	var objectType string = "virtualization.ProductInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationProductInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationProductInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationProductInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationProductInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *VirtualizationProductInfo) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *VirtualizationProductInfo) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *VirtualizationProductInfo) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *VirtualizationProductInfo) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *VirtualizationProductInfo) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *VirtualizationProductInfo) SetProductType(v string) {
	o.ProductType = &v
}

// GetProductVendor returns the ProductVendor field value if set, zero value otherwise.
func (o *VirtualizationProductInfo) GetProductVendor() string {
	if o == nil || o.ProductVendor == nil {
		var ret string
		return ret
	}
	return *o.ProductVendor
}

// GetProductVendorOk returns a tuple with the ProductVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetProductVendorOk() (*string, bool) {
	if o == nil || o.ProductVendor == nil {
		return nil, false
	}
	return o.ProductVendor, true
}

// HasProductVendor returns a boolean if a field has been set.
func (o *VirtualizationProductInfo) HasProductVendor() bool {
	if o != nil && o.ProductVendor != nil {
		return true
	}

	return false
}

// SetProductVendor gets a reference to the given string and assigns it to the ProductVendor field.
func (o *VirtualizationProductInfo) SetProductVendor(v string) {
	o.ProductVendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationProductInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationProductInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationProductInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationProductInfo) SetVersion(v string) {
	o.Version = &v
}

func (o VirtualizationProductInfo) MarshalJSON() ([]byte, error) {
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
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.ProductType != nil {
		toSerialize["ProductType"] = o.ProductType
	}
	if o.ProductVendor != nil {
		toSerialize["ProductVendor"] = o.ProductVendor
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationProductInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationProductInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Commercial product name. For example, VMware ESXi.
		ProductName *string `json:"ProductName,omitempty"`
		// Product name provided by the vendor. For example, embeddedEsx.
		ProductType *string `json:"ProductType,omitempty"`
		// Commercial vendor name. For example, VMware Inc.
		ProductVendor *string `json:"ProductVendor,omitempty"`
		// Hypervisor version running on the system.
		Version *string `json:"Version,omitempty"`
	}

	varVirtualizationProductInfoWithoutEmbeddedStruct := VirtualizationProductInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationProductInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationProductInfo := _VirtualizationProductInfo{}
		varVirtualizationProductInfo.ClassId = varVirtualizationProductInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationProductInfo.ObjectType = varVirtualizationProductInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationProductInfo.ProductName = varVirtualizationProductInfoWithoutEmbeddedStruct.ProductName
		varVirtualizationProductInfo.ProductType = varVirtualizationProductInfoWithoutEmbeddedStruct.ProductType
		varVirtualizationProductInfo.ProductVendor = varVirtualizationProductInfoWithoutEmbeddedStruct.ProductVendor
		varVirtualizationProductInfo.Version = varVirtualizationProductInfoWithoutEmbeddedStruct.Version
		*o = VirtualizationProductInfo(varVirtualizationProductInfo)
	} else {
		return err
	}

	varVirtualizationProductInfo := _VirtualizationProductInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationProductInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationProductInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "ProductVendor")
		delete(additionalProperties, "Version")

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

type NullableVirtualizationProductInfo struct {
	value *VirtualizationProductInfo
	isSet bool
}

func (v NullableVirtualizationProductInfo) Get() *VirtualizationProductInfo {
	return v.value
}

func (v *NullableVirtualizationProductInfo) Set(val *VirtualizationProductInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationProductInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationProductInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationProductInfo(val *VirtualizationProductInfo) *NullableVirtualizationProductInfo {
	return &NullableVirtualizationProductInfo{value: val, isSet: true}
}

func (v NullableVirtualizationProductInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationProductInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
