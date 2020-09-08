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

// ServerConfigChangeDetail The configuration change details are captured here.
type ServerConfigChangeDetail struct {
	PolicyAbstractConfigChangeDetail
	Profile              *ServerProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerConfigChangeDetail ServerConfigChangeDetail

// NewServerConfigChangeDetail instantiates a new ServerConfigChangeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigChangeDetail() *ServerConfigChangeDetail {
	this := ServerConfigChangeDetail{}
	return &this
}

// NewServerConfigChangeDetailWithDefaults instantiates a new ServerConfigChangeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigChangeDetailWithDefaults() *ServerConfigChangeDetail {
	this := ServerConfigChangeDetail{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ServerConfigChangeDetail) GetProfile() ServerProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ServerProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigChangeDetail) GetProfileOk() (*ServerProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ServerConfigChangeDetail) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ServerProfileRelationship and assigns it to the Profile field.
func (o *ServerConfigChangeDetail) SetProfile(v ServerProfileRelationship) {
	o.Profile = &v
}

func (o ServerConfigChangeDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigChangeDetail, errPolicyAbstractConfigChangeDetail := json.Marshal(o.PolicyAbstractConfigChangeDetail)
	if errPolicyAbstractConfigChangeDetail != nil {
		return []byte{}, errPolicyAbstractConfigChangeDetail
	}
	errPolicyAbstractConfigChangeDetail = json.Unmarshal([]byte(serializedPolicyAbstractConfigChangeDetail), &toSerialize)
	if errPolicyAbstractConfigChangeDetail != nil {
		return []byte{}, errPolicyAbstractConfigChangeDetail
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerConfigChangeDetail) UnmarshalJSON(bytes []byte) (err error) {
	type ServerConfigChangeDetailWithoutEmbeddedStruct struct {
		Profile *ServerProfileRelationship `json:"Profile,omitempty"`
	}

	varServerConfigChangeDetailWithoutEmbeddedStruct := ServerConfigChangeDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServerConfigChangeDetailWithoutEmbeddedStruct)
	if err == nil {
		varServerConfigChangeDetail := _ServerConfigChangeDetail{}
		varServerConfigChangeDetail.Profile = varServerConfigChangeDetailWithoutEmbeddedStruct.Profile
		*o = ServerConfigChangeDetail(varServerConfigChangeDetail)
	} else {
		return err
	}

	varServerConfigChangeDetail := _ServerConfigChangeDetail{}

	err = json.Unmarshal(bytes, &varServerConfigChangeDetail)
	if err == nil {
		o.PolicyAbstractConfigChangeDetail = varServerConfigChangeDetail.PolicyAbstractConfigChangeDetail
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Profile")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigChangeDetail := reflect.ValueOf(o.PolicyAbstractConfigChangeDetail)
		for i := 0; i < reflectPolicyAbstractConfigChangeDetail.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigChangeDetail.Type().Field(i)

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

type NullableServerConfigChangeDetail struct {
	value *ServerConfigChangeDetail
	isSet bool
}

func (v NullableServerConfigChangeDetail) Get() *ServerConfigChangeDetail {
	return v.value
}

func (v *NullableServerConfigChangeDetail) Set(val *ServerConfigChangeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigChangeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigChangeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigChangeDetail(val *ServerConfigChangeDetail) *NullableServerConfigChangeDetail {
	return &NullableServerConfigChangeDetail{value: val, isSet: true}
}

func (v NullableServerConfigChangeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigChangeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
