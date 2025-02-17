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

// ForecastCatalogAllOf Definition of the list of properties defined in 'forecast.Catalog', excluding properties defined in parent classes.
type ForecastCatalogAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time at which the regression model needs to run for all the metrics specified in catalog.
	SchedTime *string `json:"SchedTime,omitempty"`
	// The catalog version used in forecast configuration service.
	Version *string `json:"Version,omitempty"`
	// An array of relationships to forecastDefinition resources.
	Definition           []ForecastDefinitionRelationship `json:"Definition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastCatalogAllOf ForecastCatalogAllOf

// NewForecastCatalogAllOf instantiates a new ForecastCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastCatalogAllOf(classId string, objectType string) *ForecastCatalogAllOf {
	this := ForecastCatalogAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewForecastCatalogAllOfWithDefaults instantiates a new ForecastCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastCatalogAllOfWithDefaults() *ForecastCatalogAllOf {
	this := ForecastCatalogAllOf{}
	var classId string = "forecast.Catalog"
	this.ClassId = classId
	var objectType string = "forecast.Catalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastCatalogAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastCatalogAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastCatalogAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ForecastCatalogAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastCatalogAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastCatalogAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSchedTime returns the SchedTime field value if set, zero value otherwise.
func (o *ForecastCatalogAllOf) GetSchedTime() string {
	if o == nil || o.SchedTime == nil {
		var ret string
		return ret
	}
	return *o.SchedTime
}

// GetSchedTimeOk returns a tuple with the SchedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastCatalogAllOf) GetSchedTimeOk() (*string, bool) {
	if o == nil || o.SchedTime == nil {
		return nil, false
	}
	return o.SchedTime, true
}

// HasSchedTime returns a boolean if a field has been set.
func (o *ForecastCatalogAllOf) HasSchedTime() bool {
	if o != nil && o.SchedTime != nil {
		return true
	}

	return false
}

// SetSchedTime gets a reference to the given string and assigns it to the SchedTime field.
func (o *ForecastCatalogAllOf) SetSchedTime(v string) {
	o.SchedTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ForecastCatalogAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastCatalogAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ForecastCatalogAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ForecastCatalogAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastCatalogAllOf) GetDefinition() []ForecastDefinitionRelationship {
	if o == nil {
		var ret []ForecastDefinitionRelationship
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastCatalogAllOf) GetDefinitionOk() (*[]ForecastDefinitionRelationship, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return &o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ForecastCatalogAllOf) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given []ForecastDefinitionRelationship and assigns it to the Definition field.
func (o *ForecastCatalogAllOf) SetDefinition(v []ForecastDefinitionRelationship) {
	o.Definition = v
}

func (o ForecastCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SchedTime != nil {
		toSerialize["SchedTime"] = o.SchedTime
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Definition != nil {
		toSerialize["Definition"] = o.Definition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForecastCatalogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varForecastCatalogAllOf := _ForecastCatalogAllOf{}

	if err = json.Unmarshal(bytes, &varForecastCatalogAllOf); err == nil {
		*o = ForecastCatalogAllOf(varForecastCatalogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SchedTime")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Definition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForecastCatalogAllOf struct {
	value *ForecastCatalogAllOf
	isSet bool
}

func (v NullableForecastCatalogAllOf) Get() *ForecastCatalogAllOf {
	return v.value
}

func (v *NullableForecastCatalogAllOf) Set(val *ForecastCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastCatalogAllOf(val *ForecastCatalogAllOf) *NullableForecastCatalogAllOf {
	return &NullableForecastCatalogAllOf{value: val, isSet: true}
}

func (v NullableForecastCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
