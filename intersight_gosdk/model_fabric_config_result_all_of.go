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

// FabricConfigResultAllOf Definition of the list of properties defined in 'fabric.ConfigResult', excluding properties defined in parent classes.
type FabricConfigResultAllOf struct {
	Profile *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	// An array of relationships to fabricConfigResultEntry resources.
	ResultEntries        []FabricConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricConfigResultAllOf FabricConfigResultAllOf

// NewFabricConfigResultAllOf instantiates a new FabricConfigResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConfigResultAllOf() *FabricConfigResultAllOf {
	this := FabricConfigResultAllOf{}
	return &this
}

// NewFabricConfigResultAllOfWithDefaults instantiates a new FabricConfigResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConfigResultAllOfWithDefaults() *FabricConfigResultAllOf {
	this := FabricConfigResultAllOf{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *FabricConfigResultAllOf) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConfigResultAllOf) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricConfigResultAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given FabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricConfigResultAllOf) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricConfigResultAllOf) GetResultEntries() []FabricConfigResultEntryRelationship {
	if o == nil {
		var ret []FabricConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricConfigResultAllOf) GetResultEntriesOk() (*[]FabricConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return &o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *FabricConfigResultAllOf) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []FabricConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *FabricConfigResultAllOf) SetResultEntries(v []FabricConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o FabricConfigResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricConfigResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricConfigResultAllOf := _FabricConfigResultAllOf{}

	if err = json.Unmarshal(bytes, &varFabricConfigResultAllOf); err == nil {
		*o = FabricConfigResultAllOf(varFabricConfigResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ResultEntries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricConfigResultAllOf struct {
	value *FabricConfigResultAllOf
	isSet bool
}

func (v NullableFabricConfigResultAllOf) Get() *FabricConfigResultAllOf {
	return v.value
}

func (v *NullableFabricConfigResultAllOf) Set(val *FabricConfigResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConfigResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConfigResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConfigResultAllOf(val *FabricConfigResultAllOf) *NullableFabricConfigResultAllOf {
	return &NullableFabricConfigResultAllOf{value: val, isSet: true}
}

func (v NullableFabricConfigResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConfigResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
