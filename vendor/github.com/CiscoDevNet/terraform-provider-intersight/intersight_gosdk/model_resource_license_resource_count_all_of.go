/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ResourceLicenseResourceCountAllOf Definition of the list of properties defined in 'resource.LicenseResourceCount', excluding properties defined in parent classes.
type ResourceLicenseResourceCountAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of licensing defined for this resource group. Used for licensing group. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type.
	LicenseType *string `json:"LicenseType,omitempty"`
	// The number of resource belongs to this licensing tier.
	ResourceCount *int64                  `json:"ResourceCount,omitempty"`
	Account       *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to resourceGroup resources.
	LicenseGroups        []ResourceGroupRelationship `json:"LicenseGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceLicenseResourceCountAllOf ResourceLicenseResourceCountAllOf

// NewResourceLicenseResourceCountAllOf instantiates a new ResourceLicenseResourceCountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLicenseResourceCountAllOf(classId string, objectType string) *ResourceLicenseResourceCountAllOf {
	this := ResourceLicenseResourceCountAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var licenseType string = "Base"
	this.LicenseType = &licenseType
	return &this
}

// NewResourceLicenseResourceCountAllOfWithDefaults instantiates a new ResourceLicenseResourceCountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLicenseResourceCountAllOfWithDefaults() *ResourceLicenseResourceCountAllOf {
	this := ResourceLicenseResourceCountAllOf{}
	var classId string = "resource.LicenseResourceCount"
	this.ClassId = classId
	var objectType string = "resource.LicenseResourceCount"
	this.ObjectType = objectType
	var licenseType string = "Base"
	this.LicenseType = &licenseType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceLicenseResourceCountAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceLicenseResourceCountAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceLicenseResourceCountAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceLicenseResourceCountAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceLicenseResourceCountAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceLicenseResourceCountAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *ResourceLicenseResourceCountAllOf) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLicenseResourceCountAllOf) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *ResourceLicenseResourceCountAllOf) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *ResourceLicenseResourceCountAllOf) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetResourceCount returns the ResourceCount field value if set, zero value otherwise.
func (o *ResourceLicenseResourceCountAllOf) GetResourceCount() int64 {
	if o == nil || o.ResourceCount == nil {
		var ret int64
		return ret
	}
	return *o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLicenseResourceCountAllOf) GetResourceCountOk() (*int64, bool) {
	if o == nil || o.ResourceCount == nil {
		return nil, false
	}
	return o.ResourceCount, true
}

// HasResourceCount returns a boolean if a field has been set.
func (o *ResourceLicenseResourceCountAllOf) HasResourceCount() bool {
	if o != nil && o.ResourceCount != nil {
		return true
	}

	return false
}

// SetResourceCount gets a reference to the given int64 and assigns it to the ResourceCount field.
func (o *ResourceLicenseResourceCountAllOf) SetResourceCount(v int64) {
	o.ResourceCount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ResourceLicenseResourceCountAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLicenseResourceCountAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ResourceLicenseResourceCountAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ResourceLicenseResourceCountAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetLicenseGroups returns the LicenseGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceLicenseResourceCountAllOf) GetLicenseGroups() []ResourceGroupRelationship {
	if o == nil {
		var ret []ResourceGroupRelationship
		return ret
	}
	return o.LicenseGroups
}

// GetLicenseGroupsOk returns a tuple with the LicenseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceLicenseResourceCountAllOf) GetLicenseGroupsOk() (*[]ResourceGroupRelationship, bool) {
	if o == nil || o.LicenseGroups == nil {
		return nil, false
	}
	return &o.LicenseGroups, true
}

// HasLicenseGroups returns a boolean if a field has been set.
func (o *ResourceLicenseResourceCountAllOf) HasLicenseGroups() bool {
	if o != nil && o.LicenseGroups != nil {
		return true
	}

	return false
}

// SetLicenseGroups gets a reference to the given []ResourceGroupRelationship and assigns it to the LicenseGroups field.
func (o *ResourceLicenseResourceCountAllOf) SetLicenseGroups(v []ResourceGroupRelationship) {
	o.LicenseGroups = v
}

func (o ResourceLicenseResourceCountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.ResourceCount != nil {
		toSerialize["ResourceCount"] = o.ResourceCount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.LicenseGroups != nil {
		toSerialize["LicenseGroups"] = o.LicenseGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceLicenseResourceCountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceLicenseResourceCountAllOf := _ResourceLicenseResourceCountAllOf{}

	if err = json.Unmarshal(bytes, &varResourceLicenseResourceCountAllOf); err == nil {
		*o = ResourceLicenseResourceCountAllOf(varResourceLicenseResourceCountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "ResourceCount")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "LicenseGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceLicenseResourceCountAllOf struct {
	value *ResourceLicenseResourceCountAllOf
	isSet bool
}

func (v NullableResourceLicenseResourceCountAllOf) Get() *ResourceLicenseResourceCountAllOf {
	return v.value
}

func (v *NullableResourceLicenseResourceCountAllOf) Set(val *ResourceLicenseResourceCountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLicenseResourceCountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLicenseResourceCountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLicenseResourceCountAllOf(val *ResourceLicenseResourceCountAllOf) *NullableResourceLicenseResourceCountAllOf {
	return &NullableResourceLicenseResourceCountAllOf{value: val, isSet: true}
}

func (v NullableResourceLicenseResourceCountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLicenseResourceCountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
