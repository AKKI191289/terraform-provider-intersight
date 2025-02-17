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
)

// AssetAddressInformationAllOf Definition of the list of properties defined in 'asset.AddressInformation', excluding properties defined in parent classes.
type AssetAddressInformationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Address Line one of the address information. example \"PO BOX 641570\".
	Address1 *string `json:"Address1,omitempty"`
	// Address Line two of the address information. example \"Cisco Systems\".
	Address2 *string `json:"Address2,omitempty"`
	// Address Line three of the address information. example \"Cisco Systems\".
	Address3 *string `json:"Address3,omitempty"`
	// City in which the address resides. example \"San Jose\".
	City *string `json:"City,omitempty"`
	// Country in which the address resides. example \"US\".
	Country *string `json:"Country,omitempty"`
	// County in which the address resides. example \"Washington County\".
	County *string `json:"County,omitempty"`
	// Location in which the address resides. example \"14852\".
	Location *string `json:"Location,omitempty"`
	// Name of the user whose address is being populated.
	Name *string `json:"Name,omitempty"`
	// Postal Code in which the address resides. example \"95164-1570\".
	PostalCode *string `json:"PostalCode,omitempty"`
	// Province in which the address resides. example \"AB\".
	Province *string `json:"Province,omitempty"`
	// State in which the address resides. example \"CA\".
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetAddressInformationAllOf AssetAddressInformationAllOf

// NewAssetAddressInformationAllOf instantiates a new AssetAddressInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetAddressInformationAllOf(classId string, objectType string) *AssetAddressInformationAllOf {
	this := AssetAddressInformationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetAddressInformationAllOfWithDefaults instantiates a new AssetAddressInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetAddressInformationAllOfWithDefaults() *AssetAddressInformationAllOf {
	this := AssetAddressInformationAllOf{}
	var classId string = "asset.AddressInformation"
	this.ClassId = classId
	var objectType string = "asset.AddressInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetAddressInformationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetAddressInformationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetAddressInformationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetAddressInformationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *AssetAddressInformationAllOf) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *AssetAddressInformationAllOf) SetAddress2(v string) {
	o.Address2 = &v
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetAddress3() string {
	if o == nil || o.Address3 == nil {
		var ret string
		return ret
	}
	return *o.Address3
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetAddress3Ok() (*string, bool) {
	if o == nil || o.Address3 == nil {
		return nil, false
	}
	return o.Address3, true
}

// HasAddress3 returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasAddress3() bool {
	if o != nil && o.Address3 != nil {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given string and assigns it to the Address3 field.
func (o *AssetAddressInformationAllOf) SetAddress3(v string) {
	o.Address3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AssetAddressInformationAllOf) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AssetAddressInformationAllOf) SetCountry(v string) {
	o.Country = &v
}

// GetCounty returns the County field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetCounty() string {
	if o == nil || o.County == nil {
		var ret string
		return ret
	}
	return *o.County
}

// GetCountyOk returns a tuple with the County field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetCountyOk() (*string, bool) {
	if o == nil || o.County == nil {
		return nil, false
	}
	return o.County, true
}

// HasCounty returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasCounty() bool {
	if o != nil && o.County != nil {
		return true
	}

	return false
}

// SetCounty gets a reference to the given string and assigns it to the County field.
func (o *AssetAddressInformationAllOf) SetCounty(v string) {
	o.County = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AssetAddressInformationAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetAddressInformationAllOf) SetName(v string) {
	o.Name = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AssetAddressInformationAllOf) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetProvince() string {
	if o == nil || o.Province == nil {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetProvinceOk() (*string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *AssetAddressInformationAllOf) SetProvince(v string) {
	o.Province = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AssetAddressInformationAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAddressInformationAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AssetAddressInformationAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AssetAddressInformationAllOf) SetState(v string) {
	o.State = &v
}

func (o AssetAddressInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Address1 != nil {
		toSerialize["Address1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["Address2"] = o.Address2
	}
	if o.Address3 != nil {
		toSerialize["Address3"] = o.Address3
	}
	if o.City != nil {
		toSerialize["City"] = o.City
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.County != nil {
		toSerialize["County"] = o.County
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
	if o.Province != nil {
		toSerialize["Province"] = o.Province
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetAddressInformationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetAddressInformationAllOf := _AssetAddressInformationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetAddressInformationAllOf); err == nil {
		*o = AssetAddressInformationAllOf(varAssetAddressInformationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address1")
		delete(additionalProperties, "Address2")
		delete(additionalProperties, "Address3")
		delete(additionalProperties, "City")
		delete(additionalProperties, "Country")
		delete(additionalProperties, "County")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PostalCode")
		delete(additionalProperties, "Province")
		delete(additionalProperties, "State")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetAddressInformationAllOf struct {
	value *AssetAddressInformationAllOf
	isSet bool
}

func (v NullableAssetAddressInformationAllOf) Get() *AssetAddressInformationAllOf {
	return v.value
}

func (v *NullableAssetAddressInformationAllOf) Set(val *AssetAddressInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetAddressInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetAddressInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetAddressInformationAllOf(val *AssetAddressInformationAllOf) *NullableAssetAddressInformationAllOf {
	return &NullableAssetAddressInformationAllOf{value: val, isSet: true}
}

func (v NullableAssetAddressInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetAddressInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
