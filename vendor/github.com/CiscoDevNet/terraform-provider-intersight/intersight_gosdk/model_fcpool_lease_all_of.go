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

// FcpoolLeaseAllOf Definition of the list of properties defined in 'fcpool.Lease', excluding properties defined in parent classes.
type FcpoolLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Purpose of this WWN pool.
	PoolPurpose *string `json:"PoolPurpose,omitempty"`
	// WWN ID allocated for pool based allocation.
	WwnId                *string                       `json:"WwnId,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship         `json:"AssignedToEntity,omitempty"`
	Pool                 *FcpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *FcpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *FcpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolLeaseAllOf FcpoolLeaseAllOf

// NewFcpoolLeaseAllOf instantiates a new FcpoolLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolLeaseAllOf(classId string, objectType string) *FcpoolLeaseAllOf {
	this := FcpoolLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolLeaseAllOfWithDefaults instantiates a new FcpoolLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolLeaseAllOfWithDefaults() *FcpoolLeaseAllOf {
	this := FcpoolLeaseAllOf{}
	var classId string = "fcpool.Lease"
	this.ClassId = classId
	var objectType string = "fcpool.Lease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPoolPurpose returns the PoolPurpose field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetPoolPurpose() string {
	if o == nil || o.PoolPurpose == nil {
		var ret string
		return ret
	}
	return *o.PoolPurpose
}

// GetPoolPurposeOk returns a tuple with the PoolPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetPoolPurposeOk() (*string, bool) {
	if o == nil || o.PoolPurpose == nil {
		return nil, false
	}
	return o.PoolPurpose, true
}

// HasPoolPurpose returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasPoolPurpose() bool {
	if o != nil && o.PoolPurpose != nil {
		return true
	}

	return false
}

// SetPoolPurpose gets a reference to the given string and assigns it to the PoolPurpose field.
func (o *FcpoolLeaseAllOf) SetPoolPurpose(v string) {
	o.PoolPurpose = &v
}

// GetWwnId returns the WwnId field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetWwnId() string {
	if o == nil || o.WwnId == nil {
		var ret string
		return ret
	}
	return *o.WwnId
}

// GetWwnIdOk returns a tuple with the WwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetWwnIdOk() (*string, bool) {
	if o == nil || o.WwnId == nil {
		return nil, false
	}
	return o.WwnId, true
}

// HasWwnId returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasWwnId() bool {
	if o != nil && o.WwnId != nil {
		return true
	}

	return false
}

// SetWwnId gets a reference to the given string and assigns it to the WwnId field.
func (o *FcpoolLeaseAllOf) SetWwnId(v string) {
	o.WwnId = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *FcpoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolLeaseAllOf) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetPoolMember() FcpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret FcpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given FcpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *FcpoolLeaseAllOf) SetPoolMember(v FcpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *FcpoolLeaseAllOf) GetUniverse() FcpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret FcpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLeaseAllOf) GetUniverseOk() (*FcpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *FcpoolLeaseAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given FcpoolUniverseRelationship and assigns it to the Universe field.
func (o *FcpoolLeaseAllOf) SetUniverse(v FcpoolUniverseRelationship) {
	o.Universe = &v
}

func (o FcpoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PoolPurpose != nil {
		toSerialize["PoolPurpose"] = o.PoolPurpose
	}
	if o.WwnId != nil {
		toSerialize["WwnId"] = o.WwnId
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcpoolLeaseAllOf := _FcpoolLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varFcpoolLeaseAllOf); err == nil {
		*o = FcpoolLeaseAllOf(varFcpoolLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PoolPurpose")
		delete(additionalProperties, "WwnId")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcpoolLeaseAllOf struct {
	value *FcpoolLeaseAllOf
	isSet bool
}

func (v NullableFcpoolLeaseAllOf) Get() *FcpoolLeaseAllOf {
	return v.value
}

func (v *NullableFcpoolLeaseAllOf) Set(val *FcpoolLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolLeaseAllOf(val *FcpoolLeaseAllOf) *NullableFcpoolLeaseAllOf {
	return &NullableFcpoolLeaseAllOf{value: val, isSet: true}
}

func (v NullableFcpoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
