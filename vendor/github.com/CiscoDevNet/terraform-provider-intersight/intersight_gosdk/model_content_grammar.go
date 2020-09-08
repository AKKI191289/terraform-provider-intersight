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

// ContentGrammar Content handler framework supports extraction of required values from API/device responses. These responses may be of various content types such as XML, JSON, etc. The values of importance are modeled as parameters in the content handler framework. The parameters can be of a scalar value type or a collection of values. A group of related parameters can be modeled as a single complex type parameter. These complex types will be very useful to extract a set of repeating group of related parameters. A grammar specification defines the set of parameters that need to be extracted from the content. The grammar specification allows complex type definitions to be defined for any complex parameters.
type ContentGrammar struct {
	MoBaseComplexType
	ErrorParameters      *[]ContentBaseParameter `json:"ErrorParameters,omitempty"`
	Parameters           *[]ContentBaseParameter `json:"Parameters,omitempty"`
	Types                *[]ContentComplexType   `json:"Types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentGrammar ContentGrammar

// NewContentGrammar instantiates a new ContentGrammar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentGrammar() *ContentGrammar {
	this := ContentGrammar{}
	return &this
}

// NewContentGrammarWithDefaults instantiates a new ContentGrammar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentGrammarWithDefaults() *ContentGrammar {
	this := ContentGrammar{}
	return &this
}

// GetErrorParameters returns the ErrorParameters field value if set, zero value otherwise.
func (o *ContentGrammar) GetErrorParameters() []ContentBaseParameter {
	if o == nil || o.ErrorParameters == nil {
		var ret []ContentBaseParameter
		return ret
	}
	return *o.ErrorParameters
}

// GetErrorParametersOk returns a tuple with the ErrorParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammar) GetErrorParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.ErrorParameters == nil {
		return nil, false
	}
	return o.ErrorParameters, true
}

// HasErrorParameters returns a boolean if a field has been set.
func (o *ContentGrammar) HasErrorParameters() bool {
	if o != nil && o.ErrorParameters != nil {
		return true
	}

	return false
}

// SetErrorParameters gets a reference to the given []ContentBaseParameter and assigns it to the ErrorParameters field.
func (o *ContentGrammar) SetErrorParameters(v []ContentBaseParameter) {
	o.ErrorParameters = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ContentGrammar) GetParameters() []ContentBaseParameter {
	if o == nil || o.Parameters == nil {
		var ret []ContentBaseParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammar) GetParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentGrammar) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ContentBaseParameter and assigns it to the Parameters field.
func (o *ContentGrammar) SetParameters(v []ContentBaseParameter) {
	o.Parameters = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *ContentGrammar) GetTypes() []ContentComplexType {
	if o == nil || o.Types == nil {
		var ret []ContentComplexType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammar) GetTypesOk() (*[]ContentComplexType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ContentGrammar) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ContentComplexType and assigns it to the Types field.
func (o *ContentGrammar) SetTypes(v []ContentComplexType) {
	o.Types = &v
}

func (o ContentGrammar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.ErrorParameters != nil {
		toSerialize["ErrorParameters"] = o.ErrorParameters
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.Types != nil {
		toSerialize["Types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentGrammar) UnmarshalJSON(bytes []byte) (err error) {
	type ContentGrammarWithoutEmbeddedStruct struct {
		ErrorParameters *[]ContentBaseParameter `json:"ErrorParameters,omitempty"`
		Parameters      *[]ContentBaseParameter `json:"Parameters,omitempty"`
		Types           *[]ContentComplexType   `json:"Types,omitempty"`
	}

	varContentGrammarWithoutEmbeddedStruct := ContentGrammarWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varContentGrammarWithoutEmbeddedStruct)
	if err == nil {
		varContentGrammar := _ContentGrammar{}
		varContentGrammar.ErrorParameters = varContentGrammarWithoutEmbeddedStruct.ErrorParameters
		varContentGrammar.Parameters = varContentGrammarWithoutEmbeddedStruct.Parameters
		varContentGrammar.Types = varContentGrammarWithoutEmbeddedStruct.Types
		*o = ContentGrammar(varContentGrammar)
	} else {
		return err
	}

	varContentGrammar := _ContentGrammar{}

	err = json.Unmarshal(bytes, &varContentGrammar)
	if err == nil {
		o.MoBaseComplexType = varContentGrammar.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ErrorParameters")
		delete(additionalProperties, "Parameters")
		delete(additionalProperties, "Types")

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

type NullableContentGrammar struct {
	value *ContentGrammar
	isSet bool
}

func (v NullableContentGrammar) Get() *ContentGrammar {
	return v.value
}

func (v *NullableContentGrammar) Set(val *ContentGrammar) {
	v.value = val
	v.isSet = true
}

func (v NullableContentGrammar) IsSet() bool {
	return v.isSet
}

func (v *NullableContentGrammar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentGrammar(val *ContentGrammar) *NullableContentGrammar {
	return &NullableContentGrammar{value: val, isSet: true}
}

func (v NullableContentGrammar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentGrammar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
