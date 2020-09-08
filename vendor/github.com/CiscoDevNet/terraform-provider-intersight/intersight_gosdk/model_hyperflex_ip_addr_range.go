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
	"reflect"
	"strings"
)

// HyperflexIpAddrRange A range of IPv4 addresses. The range is inclusive and comprised of a start IP address, an end IP address, netmask, and default gateway.
type HyperflexIpAddrRange struct {
	MoBaseComplexType
	// The end IPv4 address of the range.
	EndAddr *string `json:"EndAddr,omitempty"`
	// The default gateway for the start and end IPv4 addresses.
	Gateway *string `json:"Gateway,omitempty"`
	// The netmask specified in dot decimal notation. The start address, end address, and gateway must all be within the network specified by this netmask.
	Netmask *string `json:"Netmask,omitempty"`
	// The start IPv4 address of the range.
	StartAddr            *string `json:"StartAddr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexIpAddrRange HyperflexIpAddrRange

// NewHyperflexIpAddrRange instantiates a new HyperflexIpAddrRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexIpAddrRange() *HyperflexIpAddrRange {
	this := HyperflexIpAddrRange{}
	return &this
}

// NewHyperflexIpAddrRangeWithDefaults instantiates a new HyperflexIpAddrRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexIpAddrRangeWithDefaults() *HyperflexIpAddrRange {
	this := HyperflexIpAddrRange{}
	return &this
}

// GetEndAddr returns the EndAddr field value if set, zero value otherwise.
func (o *HyperflexIpAddrRange) GetEndAddr() string {
	if o == nil || o.EndAddr == nil {
		var ret string
		return ret
	}
	return *o.EndAddr
}

// GetEndAddrOk returns a tuple with the EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRange) GetEndAddrOk() (*string, bool) {
	if o == nil || o.EndAddr == nil {
		return nil, false
	}
	return o.EndAddr, true
}

// HasEndAddr returns a boolean if a field has been set.
func (o *HyperflexIpAddrRange) HasEndAddr() bool {
	if o != nil && o.EndAddr != nil {
		return true
	}

	return false
}

// SetEndAddr gets a reference to the given string and assigns it to the EndAddr field.
func (o *HyperflexIpAddrRange) SetEndAddr(v string) {
	o.EndAddr = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *HyperflexIpAddrRange) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRange) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *HyperflexIpAddrRange) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *HyperflexIpAddrRange) SetGateway(v string) {
	o.Gateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *HyperflexIpAddrRange) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRange) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *HyperflexIpAddrRange) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *HyperflexIpAddrRange) SetNetmask(v string) {
	o.Netmask = &v
}

// GetStartAddr returns the StartAddr field value if set, zero value otherwise.
func (o *HyperflexIpAddrRange) GetStartAddr() string {
	if o == nil || o.StartAddr == nil {
		var ret string
		return ret
	}
	return *o.StartAddr
}

// GetStartAddrOk returns a tuple with the StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRange) GetStartAddrOk() (*string, bool) {
	if o == nil || o.StartAddr == nil {
		return nil, false
	}
	return o.StartAddr, true
}

// HasStartAddr returns a boolean if a field has been set.
func (o *HyperflexIpAddrRange) HasStartAddr() bool {
	if o != nil && o.StartAddr != nil {
		return true
	}

	return false
}

// SetStartAddr gets a reference to the given string and assigns it to the StartAddr field.
func (o *HyperflexIpAddrRange) SetStartAddr(v string) {
	o.StartAddr = &v
}

func (o HyperflexIpAddrRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.EndAddr != nil {
		toSerialize["EndAddr"] = o.EndAddr
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.StartAddr != nil {
		toSerialize["StartAddr"] = o.StartAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexIpAddrRange) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexIpAddrRangeWithoutEmbeddedStruct struct {
		// The end IPv4 address of the range.
		EndAddr *string `json:"EndAddr,omitempty"`
		// The default gateway for the start and end IPv4 addresses.
		Gateway *string `json:"Gateway,omitempty"`
		// The netmask specified in dot decimal notation. The start address, end address, and gateway must all be within the network specified by this netmask.
		Netmask *string `json:"Netmask,omitempty"`
		// The start IPv4 address of the range.
		StartAddr *string `json:"StartAddr,omitempty"`
	}

	varHyperflexIpAddrRangeWithoutEmbeddedStruct := HyperflexIpAddrRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexIpAddrRangeWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexIpAddrRange := _HyperflexIpAddrRange{}
		varHyperflexIpAddrRange.EndAddr = varHyperflexIpAddrRangeWithoutEmbeddedStruct.EndAddr
		varHyperflexIpAddrRange.Gateway = varHyperflexIpAddrRangeWithoutEmbeddedStruct.Gateway
		varHyperflexIpAddrRange.Netmask = varHyperflexIpAddrRangeWithoutEmbeddedStruct.Netmask
		varHyperflexIpAddrRange.StartAddr = varHyperflexIpAddrRangeWithoutEmbeddedStruct.StartAddr
		*o = HyperflexIpAddrRange(varHyperflexIpAddrRange)
	} else {
		return err
	}

	varHyperflexIpAddrRange := _HyperflexIpAddrRange{}

	err = json.Unmarshal(bytes, &varHyperflexIpAddrRange)
	if err == nil {
		o.MoBaseComplexType = varHyperflexIpAddrRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "EndAddr")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "StartAddr")

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

type NullableHyperflexIpAddrRange struct {
	value *HyperflexIpAddrRange
	isSet bool
}

func (v NullableHyperflexIpAddrRange) Get() *HyperflexIpAddrRange {
	return v.value
}

func (v *NullableHyperflexIpAddrRange) Set(val *HyperflexIpAddrRange) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexIpAddrRange) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexIpAddrRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexIpAddrRange(val *HyperflexIpAddrRange) *NullableHyperflexIpAddrRange {
	return &NullableHyperflexIpAddrRange{value: val, isSet: true}
}

func (v NullableHyperflexIpAddrRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexIpAddrRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
