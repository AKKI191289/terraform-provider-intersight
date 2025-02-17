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

// HyperflexHxapEventAllOf Definition of the list of properties defined in 'hyperflex.HxapEvent', excluding properties defined in parent classes.
type HyperflexHxapEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// First timestamp of the event occurrence.
	FirstTime *string `json:"FirstTime,omitempty"`
	// Internally generated identity (UUID) of this event.
	Identity *string `json:"Identity,omitempty"`
	// Last timestamp of the event occurrence.
	LastTime *string `json:"LastTime,omitempty"`
	// Full description of the event.
	Message *string `json:"Message,omitempty"`
	// Name of the owner with which event is associated.
	OwnerName *string `json:"OwnerName,omitempty"`
	// Type of the object with which event is associated (Host, Cluster, VM). * `Unknown` - Value is Unknown if the target type is unidentified. * `Cluster` - Cluster refers to HyperFlex AP Cluster. * `Host` - Host refers to server node which is part of HyperFlex AP Cluster. * `VM` - VM refers to Virtual machine available on a HyperFlex AP Node. * `Disk` - Disk refers to Virtual Disk available on a HyperFlex AP Cluster.
	OwnerType *string `json:"OwnerType,omitempty"`
	// UUID of the owner with which event is associated.
	OwnerUuid *string `json:"OwnerUuid,omitempty"`
	// Severity level of the event (Info/Warning/Critical). * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity             *string                           `json:"Severity,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapEventAllOf HyperflexHxapEventAllOf

// NewHyperflexHxapEventAllOf instantiates a new HyperflexHxapEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapEventAllOf(classId string, objectType string) *HyperflexHxapEventAllOf {
	this := HyperflexHxapEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ownerType string = "Unknown"
	this.OwnerType = &ownerType
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewHyperflexHxapEventAllOfWithDefaults instantiates a new HyperflexHxapEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapEventAllOfWithDefaults() *HyperflexHxapEventAllOf {
	this := HyperflexHxapEventAllOf{}
	var classId string = "hyperflex.HxapEvent"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapEvent"
	this.ObjectType = objectType
	var ownerType string = "Unknown"
	this.OwnerType = &ownerType
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFirstTime returns the FirstTime field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetFirstTime() string {
	if o == nil || o.FirstTime == nil {
		var ret string
		return ret
	}
	return *o.FirstTime
}

// GetFirstTimeOk returns a tuple with the FirstTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetFirstTimeOk() (*string, bool) {
	if o == nil || o.FirstTime == nil {
		return nil, false
	}
	return o.FirstTime, true
}

// HasFirstTime returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasFirstTime() bool {
	if o != nil && o.FirstTime != nil {
		return true
	}

	return false
}

// SetFirstTime gets a reference to the given string and assigns it to the FirstTime field.
func (o *HyperflexHxapEventAllOf) SetFirstTime(v string) {
	o.FirstTime = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *HyperflexHxapEventAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetLastTime returns the LastTime field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetLastTime() string {
	if o == nil || o.LastTime == nil {
		var ret string
		return ret
	}
	return *o.LastTime
}

// GetLastTimeOk returns a tuple with the LastTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetLastTimeOk() (*string, bool) {
	if o == nil || o.LastTime == nil {
		return nil, false
	}
	return o.LastTime, true
}

// HasLastTime returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasLastTime() bool {
	if o != nil && o.LastTime != nil {
		return true
	}

	return false
}

// SetLastTime gets a reference to the given string and assigns it to the LastTime field.
func (o *HyperflexHxapEventAllOf) SetLastTime(v string) {
	o.LastTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HyperflexHxapEventAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasOwnerName() bool {
	if o != nil && o.OwnerName != nil {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *HyperflexHxapEventAllOf) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetOwnerType() string {
	if o == nil || o.OwnerType == nil {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetOwnerTypeOk() (*string, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *HyperflexHxapEventAllOf) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetOwnerUuid returns the OwnerUuid field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetOwnerUuid() string {
	if o == nil || o.OwnerUuid == nil {
		var ret string
		return ret
	}
	return *o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetOwnerUuidOk() (*string, bool) {
	if o == nil || o.OwnerUuid == nil {
		return nil, false
	}
	return o.OwnerUuid, true
}

// HasOwnerUuid returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasOwnerUuid() bool {
	if o != nil && o.OwnerUuid != nil {
		return true
	}

	return false
}

// SetOwnerUuid gets a reference to the given string and assigns it to the OwnerUuid field.
func (o *HyperflexHxapEventAllOf) SetOwnerUuid(v string) {
	o.OwnerUuid = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *HyperflexHxapEventAllOf) SetSeverity(v string) {
	o.Severity = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapEventAllOf) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapEventAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapEventAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapEventAllOf) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

func (o HyperflexHxapEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FirstTime != nil {
		toSerialize["FirstTime"] = o.FirstTime
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.LastTime != nil {
		toSerialize["LastTime"] = o.LastTime
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.OwnerName != nil {
		toSerialize["OwnerName"] = o.OwnerName
	}
	if o.OwnerType != nil {
		toSerialize["OwnerType"] = o.OwnerType
	}
	if o.OwnerUuid != nil {
		toSerialize["OwnerUuid"] = o.OwnerUuid
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapEventAllOf := _HyperflexHxapEventAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapEventAllOf); err == nil {
		*o = HyperflexHxapEventAllOf(varHyperflexHxapEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FirstTime")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "LastTime")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "OwnerName")
		delete(additionalProperties, "OwnerType")
		delete(additionalProperties, "OwnerUuid")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "Cluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapEventAllOf struct {
	value *HyperflexHxapEventAllOf
	isSet bool
}

func (v NullableHyperflexHxapEventAllOf) Get() *HyperflexHxapEventAllOf {
	return v.value
}

func (v *NullableHyperflexHxapEventAllOf) Set(val *HyperflexHxapEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapEventAllOf(val *HyperflexHxapEventAllOf) *NullableHyperflexHxapEventAllOf {
	return &NullableHyperflexHxapEventAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
