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

// FcpoolPoolAllOf Definition of the list of properties defined in 'fcpool.Pool', excluding properties defined in parent classes.
type FcpoolPoolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string        `json:"ObjectType"`
	IdBlocks   []FcpoolBlock `json:"IdBlocks,omitempty"`
	// Purpose of this WWN pool.
	PoolPurpose *string `json:"PoolPurpose,omitempty"`
	// An array of relationships to fcpoolFcBlock resources.
	BlockHeads           []FcpoolFcBlockRelationship           `json:"BlockHeads,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolPoolAllOf FcpoolPoolAllOf

// NewFcpoolPoolAllOf instantiates a new FcpoolPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolPoolAllOf(classId string, objectType string) *FcpoolPoolAllOf {
	this := FcpoolPoolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolPoolAllOfWithDefaults instantiates a new FcpoolPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolPoolAllOfWithDefaults() *FcpoolPoolAllOf {
	this := FcpoolPoolAllOf{}
	var classId string = "fcpool.Pool"
	this.ClassId = classId
	var objectType string = "fcpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolPoolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolPoolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolPoolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolPoolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolPoolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolPoolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdBlocks returns the IdBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolAllOf) GetIdBlocks() []FcpoolBlock {
	if o == nil {
		var ret []FcpoolBlock
		return ret
	}
	return o.IdBlocks
}

// GetIdBlocksOk returns a tuple with the IdBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolAllOf) GetIdBlocksOk() (*[]FcpoolBlock, bool) {
	if o == nil || o.IdBlocks == nil {
		return nil, false
	}
	return &o.IdBlocks, true
}

// HasIdBlocks returns a boolean if a field has been set.
func (o *FcpoolPoolAllOf) HasIdBlocks() bool {
	if o != nil && o.IdBlocks != nil {
		return true
	}

	return false
}

// SetIdBlocks gets a reference to the given []FcpoolBlock and assigns it to the IdBlocks field.
func (o *FcpoolPoolAllOf) SetIdBlocks(v []FcpoolBlock) {
	o.IdBlocks = v
}

// GetPoolPurpose returns the PoolPurpose field value if set, zero value otherwise.
func (o *FcpoolPoolAllOf) GetPoolPurpose() string {
	if o == nil || o.PoolPurpose == nil {
		var ret string
		return ret
	}
	return *o.PoolPurpose
}

// GetPoolPurposeOk returns a tuple with the PoolPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolAllOf) GetPoolPurposeOk() (*string, bool) {
	if o == nil || o.PoolPurpose == nil {
		return nil, false
	}
	return o.PoolPurpose, true
}

// HasPoolPurpose returns a boolean if a field has been set.
func (o *FcpoolPoolAllOf) HasPoolPurpose() bool {
	if o != nil && o.PoolPurpose != nil {
		return true
	}

	return false
}

// SetPoolPurpose gets a reference to the given string and assigns it to the PoolPurpose field.
func (o *FcpoolPoolAllOf) SetPoolPurpose(v string) {
	o.PoolPurpose = &v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPoolAllOf) GetBlockHeads() []FcpoolFcBlockRelationship {
	if o == nil {
		var ret []FcpoolFcBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPoolAllOf) GetBlockHeadsOk() (*[]FcpoolFcBlockRelationship, bool) {
	if o == nil || o.BlockHeads == nil {
		return nil, false
	}
	return &o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *FcpoolPoolAllOf) HasBlockHeads() bool {
	if o != nil && o.BlockHeads != nil {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []FcpoolFcBlockRelationship and assigns it to the BlockHeads field.
func (o *FcpoolPoolAllOf) SetBlockHeads(v []FcpoolFcBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FcpoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FcpoolPoolAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FcpoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FcpoolPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdBlocks != nil {
		toSerialize["IdBlocks"] = o.IdBlocks
	}
	if o.PoolPurpose != nil {
		toSerialize["PoolPurpose"] = o.PoolPurpose
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolPoolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcpoolPoolAllOf := _FcpoolPoolAllOf{}

	if err = json.Unmarshal(bytes, &varFcpoolPoolAllOf); err == nil {
		*o = FcpoolPoolAllOf(varFcpoolPoolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdBlocks")
		delete(additionalProperties, "PoolPurpose")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcpoolPoolAllOf struct {
	value *FcpoolPoolAllOf
	isSet bool
}

func (v NullableFcpoolPoolAllOf) Get() *FcpoolPoolAllOf {
	return v.value
}

func (v *NullableFcpoolPoolAllOf) Set(val *FcpoolPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPoolAllOf(val *FcpoolPoolAllOf) *NullableFcpoolPoolAllOf {
	return &NullableFcpoolPoolAllOf{value: val, isSet: true}
}

func (v NullableFcpoolPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
