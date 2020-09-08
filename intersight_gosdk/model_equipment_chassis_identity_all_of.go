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

// EquipmentChassisIdentityAllOf Definition of the list of properties defined in 'equipment.ChassisIdentity', excluding properties defined in parent classes.
type EquipmentChassisIdentityAllOf struct {
	IoCardIdentityList   *[]EquipmentIoCardIdentity    `json:"IoCardIdentityList,omitempty"`
	Chassis              *EquipmentChassisRelationship `json:"Chassis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentChassisIdentityAllOf EquipmentChassisIdentityAllOf

// NewEquipmentChassisIdentityAllOf instantiates a new EquipmentChassisIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentChassisIdentityAllOf() *EquipmentChassisIdentityAllOf {
	this := EquipmentChassisIdentityAllOf{}
	return &this
}

// NewEquipmentChassisIdentityAllOfWithDefaults instantiates a new EquipmentChassisIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentChassisIdentityAllOfWithDefaults() *EquipmentChassisIdentityAllOf {
	this := EquipmentChassisIdentityAllOf{}
	return &this
}

// GetIoCardIdentityList returns the IoCardIdentityList field value if set, zero value otherwise.
func (o *EquipmentChassisIdentityAllOf) GetIoCardIdentityList() []EquipmentIoCardIdentity {
	if o == nil || o.IoCardIdentityList == nil {
		var ret []EquipmentIoCardIdentity
		return ret
	}
	return *o.IoCardIdentityList
}

// GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisIdentityAllOf) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool) {
	if o == nil || o.IoCardIdentityList == nil {
		return nil, false
	}
	return o.IoCardIdentityList, true
}

// HasIoCardIdentityList returns a boolean if a field has been set.
func (o *EquipmentChassisIdentityAllOf) HasIoCardIdentityList() bool {
	if o != nil && o.IoCardIdentityList != nil {
		return true
	}

	return false
}

// SetIoCardIdentityList gets a reference to the given []EquipmentIoCardIdentity and assigns it to the IoCardIdentityList field.
func (o *EquipmentChassisIdentityAllOf) SetIoCardIdentityList(v []EquipmentIoCardIdentity) {
	o.IoCardIdentityList = &v
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *EquipmentChassisIdentityAllOf) GetChassis() EquipmentChassisRelationship {
	if o == nil || o.Chassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisIdentityAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *EquipmentChassisIdentityAllOf) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *EquipmentChassisIdentityAllOf) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis = &v
}

func (o EquipmentChassisIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IoCardIdentityList != nil {
		toSerialize["IoCardIdentityList"] = o.IoCardIdentityList
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentChassisIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentChassisIdentityAllOf := _EquipmentChassisIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentChassisIdentityAllOf); err == nil {
		*o = EquipmentChassisIdentityAllOf(varEquipmentChassisIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IoCardIdentityList")
		delete(additionalProperties, "Chassis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentChassisIdentityAllOf struct {
	value *EquipmentChassisIdentityAllOf
	isSet bool
}

func (v NullableEquipmentChassisIdentityAllOf) Get() *EquipmentChassisIdentityAllOf {
	return v.value
}

func (v *NullableEquipmentChassisIdentityAllOf) Set(val *EquipmentChassisIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentChassisIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentChassisIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentChassisIdentityAllOf(val *EquipmentChassisIdentityAllOf) *NullableEquipmentChassisIdentityAllOf {
	return &NullableEquipmentChassisIdentityAllOf{value: val, isSet: true}
}

func (v NullableEquipmentChassisIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentChassisIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
