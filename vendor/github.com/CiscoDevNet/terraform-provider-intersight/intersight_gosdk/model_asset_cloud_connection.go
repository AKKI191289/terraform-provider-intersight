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

// AssetCloudConnection CloudConnection provides the necessary details for Intersight to connect to and authenticate with a target at a well-known service address. The service address is inferred based upon the target type. For example Amazon Web Services.
type AssetCloudConnection struct {
	AssetConnection
	AdditionalProperties map[string]interface{}
}

type _AssetCloudConnection AssetCloudConnection

// NewAssetCloudConnection instantiates a new AssetCloudConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCloudConnection() *AssetCloudConnection {
	this := AssetCloudConnection{}
	return &this
}

// NewAssetCloudConnectionWithDefaults instantiates a new AssetCloudConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCloudConnectionWithDefaults() *AssetCloudConnection {
	this := AssetCloudConnection{}
	return &this
}

func (o AssetCloudConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetConnection, errAssetConnection := json.Marshal(o.AssetConnection)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}
	errAssetConnection = json.Unmarshal([]byte(serializedAssetConnection), &toSerialize)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetCloudConnection) UnmarshalJSON(bytes []byte) (err error) {
	type AssetCloudConnectionWithoutEmbeddedStruct struct {
	}

	varAssetCloudConnectionWithoutEmbeddedStruct := AssetCloudConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetCloudConnectionWithoutEmbeddedStruct)
	if err == nil {
		varAssetCloudConnection := _AssetCloudConnection{}
		*o = AssetCloudConnection(varAssetCloudConnection)
	} else {
		return err
	}

	varAssetCloudConnection := _AssetCloudConnection{}

	err = json.Unmarshal(bytes, &varAssetCloudConnection)
	if err == nil {
		o.AssetConnection = varAssetCloudConnection.AssetConnection
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAssetConnection := reflect.ValueOf(o.AssetConnection)
		for i := 0; i < reflectAssetConnection.Type().NumField(); i++ {
			t := reflectAssetConnection.Type().Field(i)

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

type NullableAssetCloudConnection struct {
	value *AssetCloudConnection
	isSet bool
}

func (v NullableAssetCloudConnection) Get() *AssetCloudConnection {
	return v.value
}

func (v *NullableAssetCloudConnection) Set(val *AssetCloudConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCloudConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCloudConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCloudConnection(val *AssetCloudConnection) *NullableAssetCloudConnection {
	return &NullableAssetCloudConnection{value: val, isSet: true}
}

func (v NullableAssetCloudConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCloudConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
