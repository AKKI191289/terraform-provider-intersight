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

// AssetAddressInformation Type for saving the address information. It is used in asset.DeviceContractInformation object to save customer address.
type AssetAddressInformation struct {
	MoBaseComplexType
	// Address Line one of the address information. example \"PO BOX 641570\".
	Address1 *string `json:"Address1,omitempty"`
	// Address Line two of the address information. example \"Cisco Systems\".
	Address2 *string `json:"Address2,omitempty"`
	// City in which the address resides. example \"San Jose\".
	City *string `json:"City,omitempty"`
	// Country in which the address resides. example \"US\".
	Country *string `json:"Country,omitempty"`
	// Location in which the address resides. example \"14852\".
	Location *string `json:"Location,omitempty"`
	// Name of the user whose address is being populated.
	Name *string `json:"Name,omitempty"`
	// Postal Code in which the address resides. example \"95164-1570\".
	PostalCode *string `json:"PostalCode,omitempty"`
	// State in which the address resides. example \"CA\".
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetAddressInformation AssetAddressInformation

// NewAssetAddressInformation instantiates a new AssetAddressInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetAddressInformation() *AssetAddressInformation {
	this := AssetAddressInformation{}
	return &this
}

// NewAssetAddressInformationWithDefaults instantiates a new AssetAddressInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetAddressInformationWithDefaults() *AssetAddressInformation {
	this := AssetAddressInformation{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *AssetAddressInformation) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *AssetAddressInformation) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AssetAddressInformation) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AssetAddressInformation) SetCountry(v string) {
	o.Country = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AssetAddressInformation) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetAddressInformation) SetName(v string) {
	o.Name = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AssetAddressInformation) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AssetAddressInformation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AssetAddressInformation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AssetAddressInformation) SetState(v string) {
	o.State = &v
}

func (o AssetAddressInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Address1 != nil {
		toSerialize["Address1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["Address2"] = o.Address2
	}
	if o.City != nil {
		toSerialize["City"] = o.City
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PostalCode != nil {
		toSerialize["PostalCode"] = o.PostalCode
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetAddressInformation) UnmarshalJSON(bytes []byte) (err error) {
	type AssetAddressInformationWithoutEmbeddedStruct struct {
		// Address Line one of the address information. example \"PO BOX 641570\".
		Address1 *string `json:"Address1,omitempty"`
		// Address Line two of the address information. example \"Cisco Systems\".
		Address2 *string `json:"Address2,omitempty"`
		// City in which the address resides. example \"San Jose\".
		City *string `json:"City,omitempty"`
		// Country in which the address resides. example \"US\".
		Country *string `json:"Country,omitempty"`
		// Location in which the address resides. example \"14852\".
		Location *string `json:"Location,omitempty"`
		// Name of the user whose address is being populated.
		Name *string `json:"Name,omitempty"`
		// Postal Code in which the address resides. example \"95164-1570\".
		PostalCode *string `json:"PostalCode,omitempty"`
		// State in which the address resides. example \"CA\".
		State *string `json:"State,omitempty"`
	}

	varAssetAddressInformationWithoutEmbeddedStruct := AssetAddressInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetAddressInformationWithoutEmbeddedStruct)
	if err == nil {
		varAssetAddressInformation := _AssetAddressInformation{}
		varAssetAddressInformation.Address1 = varAssetAddressInformationWithoutEmbeddedStruct.Address1
		varAssetAddressInformation.Address2 = varAssetAddressInformationWithoutEmbeddedStruct.Address2
		varAssetAddressInformation.City = varAssetAddressInformationWithoutEmbeddedStruct.City
		varAssetAddressInformation.Country = varAssetAddressInformationWithoutEmbeddedStruct.Country
		varAssetAddressInformation.Location = varAssetAddressInformationWithoutEmbeddedStruct.Location
		varAssetAddressInformation.Name = varAssetAddressInformationWithoutEmbeddedStruct.Name
		varAssetAddressInformation.PostalCode = varAssetAddressInformationWithoutEmbeddedStruct.PostalCode
		varAssetAddressInformation.State = varAssetAddressInformationWithoutEmbeddedStruct.State
		*o = AssetAddressInformation(varAssetAddressInformation)
	} else {
		return err
	}

	varAssetAddressInformation := _AssetAddressInformation{}

	err = json.Unmarshal(bytes, &varAssetAddressInformation)
	if err == nil {
		o.MoBaseComplexType = varAssetAddressInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Address1")
		delete(additionalProperties, "Address2")
		delete(additionalProperties, "City")
		delete(additionalProperties, "Country")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PostalCode")
		delete(additionalProperties, "State")

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

type NullableAssetAddressInformation struct {
	value *AssetAddressInformation
	isSet bool
}

func (v NullableAssetAddressInformation) Get() *AssetAddressInformation {
	return v.value
}

func (v *NullableAssetAddressInformation) Set(val *AssetAddressInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetAddressInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetAddressInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetAddressInformation(val *AssetAddressInformation) *NullableAssetAddressInformation {
	return &NullableAssetAddressInformation{value: val, isSet: true}
}

func (v NullableAssetAddressInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetAddressInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
