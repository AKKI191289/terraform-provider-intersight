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

// FirmwareFabricUpgradeImpactAllOf Definition of the list of properties defined in 'firmware.FabricUpgradeImpact', excluding properties defined in parent classes.
type FirmwareFabricUpgradeImpactAllOf struct {
	ImpactDetail *[]FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
	// Details for the Fabric Interconnect that will be impacted by the upgrade.
	Serial               *string `json:"Serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareFabricUpgradeImpactAllOf FirmwareFabricUpgradeImpactAllOf

// NewFirmwareFabricUpgradeImpactAllOf instantiates a new FirmwareFabricUpgradeImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFabricUpgradeImpactAllOf() *FirmwareFabricUpgradeImpactAllOf {
	this := FirmwareFabricUpgradeImpactAllOf{}
	return &this
}

// NewFirmwareFabricUpgradeImpactAllOfWithDefaults instantiates a new FirmwareFabricUpgradeImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFabricUpgradeImpactAllOfWithDefaults() *FirmwareFabricUpgradeImpactAllOf {
	this := FirmwareFabricUpgradeImpactAllOf{}
	return &this
}

// GetImpactDetail returns the ImpactDetail field value if set, zero value otherwise.
func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact {
	if o == nil || o.ImpactDetail == nil {
		var ret []FirmwareComponentImpact
		return ret
	}
	return *o.ImpactDetail
}

// GetImpactDetailOk returns a tuple with the ImpactDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool) {
	if o == nil || o.ImpactDetail == nil {
		return nil, false
	}
	return o.ImpactDetail, true
}

// HasImpactDetail returns a boolean if a field has been set.
func (o *FirmwareFabricUpgradeImpactAllOf) HasImpactDetail() bool {
	if o != nil && o.ImpactDetail != nil {
		return true
	}

	return false
}

// SetImpactDetail gets a reference to the given []FirmwareComponentImpact and assigns it to the ImpactDetail field.
func (o *FirmwareFabricUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact) {
	o.ImpactDetail = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *FirmwareFabricUpgradeImpactAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFabricUpgradeImpactAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *FirmwareFabricUpgradeImpactAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *FirmwareFabricUpgradeImpactAllOf) SetSerial(v string) {
	o.Serial = &v
}

func (o FirmwareFabricUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImpactDetail != nil {
		toSerialize["ImpactDetail"] = o.ImpactDetail
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareFabricUpgradeImpactAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareFabricUpgradeImpactAllOf := _FirmwareFabricUpgradeImpactAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareFabricUpgradeImpactAllOf); err == nil {
		*o = FirmwareFabricUpgradeImpactAllOf(varFirmwareFabricUpgradeImpactAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ImpactDetail")
		delete(additionalProperties, "Serial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareFabricUpgradeImpactAllOf struct {
	value *FirmwareFabricUpgradeImpactAllOf
	isSet bool
}

func (v NullableFirmwareFabricUpgradeImpactAllOf) Get() *FirmwareFabricUpgradeImpactAllOf {
	return v.value
}

func (v *NullableFirmwareFabricUpgradeImpactAllOf) Set(val *FirmwareFabricUpgradeImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFabricUpgradeImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFabricUpgradeImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFabricUpgradeImpactAllOf(val *FirmwareFabricUpgradeImpactAllOf) *NullableFirmwareFabricUpgradeImpactAllOf {
	return &NullableFirmwareFabricUpgradeImpactAllOf{value: val, isSet: true}
}

func (v NullableFirmwareFabricUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFabricUpgradeImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
