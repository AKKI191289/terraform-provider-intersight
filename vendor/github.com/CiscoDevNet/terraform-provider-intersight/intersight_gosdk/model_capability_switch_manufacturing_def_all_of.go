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

// CapabilitySwitchManufacturingDefAllOf Definition of the list of properties defined in 'capability.SwitchManufacturingDef', excluding properties defined in parent classes.
type CapabilitySwitchManufacturingDefAllOf struct {
	// Caption for Switch/Fabric-Interconnect.
	Caption *string `json:"Caption,omitempty"`
	// Description for Switch/Fabric-Interconnect.
	Description *string `json:"Description,omitempty"`
	// Part Number for Switch/Fabric-Interconnect.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Product Name for Switch/Fabric-Interconnect.
	ProductName          *string `json:"ProductName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchManufacturingDefAllOf CapabilitySwitchManufacturingDefAllOf

// NewCapabilitySwitchManufacturingDefAllOf instantiates a new CapabilitySwitchManufacturingDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchManufacturingDefAllOf() *CapabilitySwitchManufacturingDefAllOf {
	this := CapabilitySwitchManufacturingDefAllOf{}
	return &this
}

// NewCapabilitySwitchManufacturingDefAllOfWithDefaults instantiates a new CapabilitySwitchManufacturingDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchManufacturingDefAllOfWithDefaults() *CapabilitySwitchManufacturingDefAllOf {
	this := CapabilitySwitchManufacturingDefAllOf{}
	return &this
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDefAllOf) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *CapabilitySwitchManufacturingDefAllOf) SetCaption(v string) {
	o.Caption = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDefAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilitySwitchManufacturingDefAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDefAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *CapabilitySwitchManufacturingDefAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDefAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDefAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *CapabilitySwitchManufacturingDefAllOf) SetProductName(v string) {
	o.ProductName = &v
}

func (o CapabilitySwitchManufacturingDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Caption != nil {
		toSerialize["Caption"] = o.Caption
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchManufacturingDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySwitchManufacturingDefAllOf := _CapabilitySwitchManufacturingDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySwitchManufacturingDefAllOf); err == nil {
		*o = CapabilitySwitchManufacturingDefAllOf(varCapabilitySwitchManufacturingDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Caption")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "ProductName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchManufacturingDefAllOf struct {
	value *CapabilitySwitchManufacturingDefAllOf
	isSet bool
}

func (v NullableCapabilitySwitchManufacturingDefAllOf) Get() *CapabilitySwitchManufacturingDefAllOf {
	return v.value
}

func (v *NullableCapabilitySwitchManufacturingDefAllOf) Set(val *CapabilitySwitchManufacturingDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchManufacturingDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchManufacturingDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchManufacturingDefAllOf(val *CapabilitySwitchManufacturingDefAllOf) *NullableCapabilitySwitchManufacturingDefAllOf {
	return &NullableCapabilitySwitchManufacturingDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySwitchManufacturingDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchManufacturingDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
