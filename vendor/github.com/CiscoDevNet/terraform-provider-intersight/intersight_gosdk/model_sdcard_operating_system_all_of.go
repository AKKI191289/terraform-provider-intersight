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

// SdcardOperatingSystemAllOf Definition of the list of properties defined in 'sdcard.OperatingSystem', excluding properties defined in parent classes.
type SdcardOperatingSystemAllOf struct {
	// Name of virtual drive for operating system partition.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdcardOperatingSystemAllOf SdcardOperatingSystemAllOf

// NewSdcardOperatingSystemAllOf instantiates a new SdcardOperatingSystemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardOperatingSystemAllOf() *SdcardOperatingSystemAllOf {
	this := SdcardOperatingSystemAllOf{}
	return &this
}

// NewSdcardOperatingSystemAllOfWithDefaults instantiates a new SdcardOperatingSystemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardOperatingSystemAllOfWithDefaults() *SdcardOperatingSystemAllOf {
	this := SdcardOperatingSystemAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdcardOperatingSystemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardOperatingSystemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdcardOperatingSystemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdcardOperatingSystemAllOf) SetName(v string) {
	o.Name = &v
}

func (o SdcardOperatingSystemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdcardOperatingSystemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdcardOperatingSystemAllOf := _SdcardOperatingSystemAllOf{}

	if err = json.Unmarshal(bytes, &varSdcardOperatingSystemAllOf); err == nil {
		*o = SdcardOperatingSystemAllOf(varSdcardOperatingSystemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdcardOperatingSystemAllOf struct {
	value *SdcardOperatingSystemAllOf
	isSet bool
}

func (v NullableSdcardOperatingSystemAllOf) Get() *SdcardOperatingSystemAllOf {
	return v.value
}

func (v *NullableSdcardOperatingSystemAllOf) Set(val *SdcardOperatingSystemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardOperatingSystemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardOperatingSystemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardOperatingSystemAllOf(val *SdcardOperatingSystemAllOf) *NullableSdcardOperatingSystemAllOf {
	return &NullableSdcardOperatingSystemAllOf{value: val, isSet: true}
}

func (v NullableSdcardOperatingSystemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardOperatingSystemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
