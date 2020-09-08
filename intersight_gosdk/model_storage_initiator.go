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

// StorageInitiator An initiator is the consumer of storage, typically a server with an adapter card in it called a Host Bus Adapter (HBA). The initiator \"initiates\" a connection over the fabric to one or more ports on storage system target ports.
type StorageInitiator struct {
	StorageBaseInitiator
	AdditionalProperties map[string]interface{}
}

type _StorageInitiator StorageInitiator

// NewStorageInitiator instantiates a new StorageInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageInitiator() *StorageInitiator {
	this := StorageInitiator{}
	return &this
}

// NewStorageInitiatorWithDefaults instantiates a new StorageInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageInitiatorWithDefaults() *StorageInitiator {
	this := StorageInitiator{}
	return &this
}

func (o StorageInitiator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseInitiator, errStorageBaseInitiator := json.Marshal(o.StorageBaseInitiator)
	if errStorageBaseInitiator != nil {
		return []byte{}, errStorageBaseInitiator
	}
	errStorageBaseInitiator = json.Unmarshal([]byte(serializedStorageBaseInitiator), &toSerialize)
	if errStorageBaseInitiator != nil {
		return []byte{}, errStorageBaseInitiator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageInitiator) UnmarshalJSON(bytes []byte) (err error) {
	type StorageInitiatorWithoutEmbeddedStruct struct {
	}

	varStorageInitiatorWithoutEmbeddedStruct := StorageInitiatorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageInitiatorWithoutEmbeddedStruct)
	if err == nil {
		varStorageInitiator := _StorageInitiator{}
		*o = StorageInitiator(varStorageInitiator)
	} else {
		return err
	}

	varStorageInitiator := _StorageInitiator{}

	err = json.Unmarshal(bytes, &varStorageInitiator)
	if err == nil {
		o.StorageBaseInitiator = varStorageInitiator.StorageBaseInitiator
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectStorageBaseInitiator := reflect.ValueOf(o.StorageBaseInitiator)
		for i := 0; i < reflectStorageBaseInitiator.Type().NumField(); i++ {
			t := reflectStorageBaseInitiator.Type().Field(i)

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

type NullableStorageInitiator struct {
	value *StorageInitiator
	isSet bool
}

func (v NullableStorageInitiator) Get() *StorageInitiator {
	return v.value
}

func (v *NullableStorageInitiator) Set(val *StorageInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageInitiator(val *StorageInitiator) *NullableStorageInitiator {
	return &NullableStorageInitiator{value: val, isSet: true}
}

func (v NullableStorageInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
