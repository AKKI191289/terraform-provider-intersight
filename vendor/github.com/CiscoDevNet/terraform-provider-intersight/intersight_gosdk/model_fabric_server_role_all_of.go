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

// FabricServerRoleAllOf Definition of the list of properties defined in 'fabric.ServerRole', excluding properties defined in parent classes.
type FabricServerRoleAllOf struct {
	// Forward error correction configuration for the port. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
	Fec                  *string `json:"Fec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricServerRoleAllOf FabricServerRoleAllOf

// NewFabricServerRoleAllOf instantiates a new FabricServerRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricServerRoleAllOf() *FabricServerRoleAllOf {
	this := FabricServerRoleAllOf{}
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// NewFabricServerRoleAllOfWithDefaults instantiates a new FabricServerRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricServerRoleAllOfWithDefaults() *FabricServerRoleAllOf {
	this := FabricServerRoleAllOf{}
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// GetFec returns the Fec field value if set, zero value otherwise.
func (o *FabricServerRoleAllOf) GetFec() string {
	if o == nil || o.Fec == nil {
		var ret string
		return ret
	}
	return *o.Fec
}

// GetFecOk returns a tuple with the Fec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricServerRoleAllOf) GetFecOk() (*string, bool) {
	if o == nil || o.Fec == nil {
		return nil, false
	}
	return o.Fec, true
}

// HasFec returns a boolean if a field has been set.
func (o *FabricServerRoleAllOf) HasFec() bool {
	if o != nil && o.Fec != nil {
		return true
	}

	return false
}

// SetFec gets a reference to the given string and assigns it to the Fec field.
func (o *FabricServerRoleAllOf) SetFec(v string) {
	o.Fec = &v
}

func (o FabricServerRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fec != nil {
		toSerialize["Fec"] = o.Fec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricServerRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricServerRoleAllOf := _FabricServerRoleAllOf{}

	if err = json.Unmarshal(bytes, &varFabricServerRoleAllOf); err == nil {
		*o = FabricServerRoleAllOf(varFabricServerRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Fec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricServerRoleAllOf struct {
	value *FabricServerRoleAllOf
	isSet bool
}

func (v NullableFabricServerRoleAllOf) Get() *FabricServerRoleAllOf {
	return v.value
}

func (v *NullableFabricServerRoleAllOf) Set(val *FabricServerRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricServerRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricServerRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricServerRoleAllOf(val *FabricServerRoleAllOf) *NullableFabricServerRoleAllOf {
	return &NullableFabricServerRoleAllOf{value: val, isSet: true}
}

func (v NullableFabricServerRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricServerRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
