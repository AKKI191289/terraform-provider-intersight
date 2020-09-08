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

// PkixSubjectAlternateName The additional hostnames to be protected by a given X.509 certificate.
type PkixSubjectAlternateName struct {
	MoBaseComplexType
	DnsName              *[]string `json:"DnsName,omitempty"`
	EmailAddress         *[]string `json:"EmailAddress,omitempty"`
	IpAddress            *[]string `json:"IpAddress,omitempty"`
	Uri                  *[]string `json:"Uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixSubjectAlternateName PkixSubjectAlternateName

// NewPkixSubjectAlternateName instantiates a new PkixSubjectAlternateName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixSubjectAlternateName() *PkixSubjectAlternateName {
	this := PkixSubjectAlternateName{}
	return &this
}

// NewPkixSubjectAlternateNameWithDefaults instantiates a new PkixSubjectAlternateName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixSubjectAlternateNameWithDefaults() *PkixSubjectAlternateName {
	this := PkixSubjectAlternateName{}
	return &this
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *PkixSubjectAlternateName) GetDnsName() []string {
	if o == nil || o.DnsName == nil {
		var ret []string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateName) GetDnsNameOk() (*[]string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *PkixSubjectAlternateName) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given []string and assigns it to the DnsName field.
func (o *PkixSubjectAlternateName) SetDnsName(v []string) {
	o.DnsName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *PkixSubjectAlternateName) GetEmailAddress() []string {
	if o == nil || o.EmailAddress == nil {
		var ret []string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateName) GetEmailAddressOk() (*[]string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *PkixSubjectAlternateName) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given []string and assigns it to the EmailAddress field.
func (o *PkixSubjectAlternateName) SetEmailAddress(v []string) {
	o.EmailAddress = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *PkixSubjectAlternateName) GetIpAddress() []string {
	if o == nil || o.IpAddress == nil {
		var ret []string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateName) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PkixSubjectAlternateName) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *PkixSubjectAlternateName) SetIpAddress(v []string) {
	o.IpAddress = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PkixSubjectAlternateName) GetUri() []string {
	if o == nil || o.Uri == nil {
		var ret []string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateName) GetUriOk() (*[]string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PkixSubjectAlternateName) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given []string and assigns it to the Uri field.
func (o *PkixSubjectAlternateName) SetUri(v []string) {
	o.Uri = &v
}

func (o PkixSubjectAlternateName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DnsName != nil {
		toSerialize["DnsName"] = o.DnsName
	}
	if o.EmailAddress != nil {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixSubjectAlternateName) UnmarshalJSON(bytes []byte) (err error) {
	type PkixSubjectAlternateNameWithoutEmbeddedStruct struct {
		DnsName      *[]string `json:"DnsName,omitempty"`
		EmailAddress *[]string `json:"EmailAddress,omitempty"`
		IpAddress    *[]string `json:"IpAddress,omitempty"`
		Uri          *[]string `json:"Uri,omitempty"`
	}

	varPkixSubjectAlternateNameWithoutEmbeddedStruct := PkixSubjectAlternateNameWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPkixSubjectAlternateNameWithoutEmbeddedStruct)
	if err == nil {
		varPkixSubjectAlternateName := _PkixSubjectAlternateName{}
		varPkixSubjectAlternateName.DnsName = varPkixSubjectAlternateNameWithoutEmbeddedStruct.DnsName
		varPkixSubjectAlternateName.EmailAddress = varPkixSubjectAlternateNameWithoutEmbeddedStruct.EmailAddress
		varPkixSubjectAlternateName.IpAddress = varPkixSubjectAlternateNameWithoutEmbeddedStruct.IpAddress
		varPkixSubjectAlternateName.Uri = varPkixSubjectAlternateNameWithoutEmbeddedStruct.Uri
		*o = PkixSubjectAlternateName(varPkixSubjectAlternateName)
	} else {
		return err
	}

	varPkixSubjectAlternateName := _PkixSubjectAlternateName{}

	err = json.Unmarshal(bytes, &varPkixSubjectAlternateName)
	if err == nil {
		o.MoBaseComplexType = varPkixSubjectAlternateName.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DnsName")
		delete(additionalProperties, "EmailAddress")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Uri")

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

type NullablePkixSubjectAlternateName struct {
	value *PkixSubjectAlternateName
	isSet bool
}

func (v NullablePkixSubjectAlternateName) Get() *PkixSubjectAlternateName {
	return v.value
}

func (v *NullablePkixSubjectAlternateName) Set(val *PkixSubjectAlternateName) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixSubjectAlternateName) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixSubjectAlternateName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixSubjectAlternateName(val *PkixSubjectAlternateName) *NullablePkixSubjectAlternateName {
	return &NullablePkixSubjectAlternateName{value: val, isSet: true}
}

func (v NullablePkixSubjectAlternateName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixSubjectAlternateName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
