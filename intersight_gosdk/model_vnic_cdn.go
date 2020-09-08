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

// VnicCdn Consistent Device Naming configuration for the virtual NIC.
type VnicCdn struct {
	MoBaseComplexType
	// Source of the CDN. It can either be user specified or be the same as the vNIC name. * `vnic` - Source of the CDN is the same as the vNIC name. * `user` - Source of the CDN is specified by the user.
	Source *string `json:"Source,omitempty"`
	// The CDN value entered in case of user defined mode.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicCdn VnicCdn

// NewVnicCdn instantiates a new VnicCdn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicCdn() *VnicCdn {
	this := VnicCdn{}
	var source string = "vnic"
	this.Source = &source
	return &this
}

// NewVnicCdnWithDefaults instantiates a new VnicCdn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicCdnWithDefaults() *VnicCdn {
	this := VnicCdn{}
	var source string = "vnic"
	this.Source = &source
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *VnicCdn) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCdn) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *VnicCdn) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *VnicCdn) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VnicCdn) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCdn) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VnicCdn) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VnicCdn) SetValue(v string) {
	o.Value = &v
}

func (o VnicCdn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicCdn) UnmarshalJSON(bytes []byte) (err error) {
	type VnicCdnWithoutEmbeddedStruct struct {
		// Source of the CDN. It can either be user specified or be the same as the vNIC name. * `vnic` - Source of the CDN is the same as the vNIC name. * `user` - Source of the CDN is specified by the user.
		Source *string `json:"Source,omitempty"`
		// The CDN value entered in case of user defined mode.
		Value *string `json:"Value,omitempty"`
	}

	varVnicCdnWithoutEmbeddedStruct := VnicCdnWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicCdnWithoutEmbeddedStruct)
	if err == nil {
		varVnicCdn := _VnicCdn{}
		varVnicCdn.Source = varVnicCdnWithoutEmbeddedStruct.Source
		varVnicCdn.Value = varVnicCdnWithoutEmbeddedStruct.Value
		*o = VnicCdn(varVnicCdn)
	} else {
		return err
	}

	varVnicCdn := _VnicCdn{}

	err = json.Unmarshal(bytes, &varVnicCdn)
	if err == nil {
		o.MoBaseComplexType = varVnicCdn.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Value")

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

type NullableVnicCdn struct {
	value *VnicCdn
	isSet bool
}

func (v NullableVnicCdn) Get() *VnicCdn {
	return v.value
}

func (v *NullableVnicCdn) Set(val *VnicCdn) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicCdn) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicCdn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicCdn(val *VnicCdn) *NullableVnicCdn {
	return &NullableVnicCdn{value: val, isSet: true}
}

func (v NullableVnicCdn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicCdn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
