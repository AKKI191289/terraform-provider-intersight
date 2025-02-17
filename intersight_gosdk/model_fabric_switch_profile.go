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
	"reflect"
	"strings"
)

// FabricSwitchProfile This specifies configuration policies for a managed network switch.
type FabricSwitchProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                      `json:"ObjectType"`
	ConfigChanges    NullablePolicyConfigChange  `json:"ConfigChanges,omitempty"`
	AssignedSwitch   *NetworkElementRelationship `json:"AssignedSwitch,omitempty"`
	AssociatedSwitch *NetworkElementRelationship `json:"AssociatedSwitch,omitempty"`
	// An array of relationships to fabricConfigChangeDetail resources.
	ConfigChangeDetails []FabricConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	ConfigResult        *FabricConfigResultRelationship        `json:"ConfigResult,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship      `json:"RunningWorkflows,omitempty"`
	SwitchClusterProfile *FabricSwitchClusterProfileRelationship `json:"SwitchClusterProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSwitchProfile FabricSwitchProfile

// NewFabricSwitchProfile instantiates a new FabricSwitchProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchProfile(classId string, objectType string) *FabricSwitchProfile {
	this := FabricSwitchProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricSwitchProfileWithDefaults instantiates a new FabricSwitchProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchProfileWithDefaults() *FabricSwitchProfile {
	this := FabricSwitchProfile{}
	var classId string = "fabric.SwitchProfile"
	this.ClassId = classId
	var objectType string = "fabric.SwitchProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfile) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfile) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *FabricSwitchProfile) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *FabricSwitchProfile) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *FabricSwitchProfile) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetAssignedSwitch returns the AssignedSwitch field value if set, zero value otherwise.
func (o *FabricSwitchProfile) GetAssignedSwitch() NetworkElementRelationship {
	if o == nil || o.AssignedSwitch == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.AssignedSwitch
}

// GetAssignedSwitchOk returns a tuple with the AssignedSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetAssignedSwitchOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.AssignedSwitch == nil {
		return nil, false
	}
	return o.AssignedSwitch, true
}

// HasAssignedSwitch returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasAssignedSwitch() bool {
	if o != nil && o.AssignedSwitch != nil {
		return true
	}

	return false
}

// SetAssignedSwitch gets a reference to the given NetworkElementRelationship and assigns it to the AssignedSwitch field.
func (o *FabricSwitchProfile) SetAssignedSwitch(v NetworkElementRelationship) {
	o.AssignedSwitch = &v
}

// GetAssociatedSwitch returns the AssociatedSwitch field value if set, zero value otherwise.
func (o *FabricSwitchProfile) GetAssociatedSwitch() NetworkElementRelationship {
	if o == nil || o.AssociatedSwitch == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.AssociatedSwitch
}

// GetAssociatedSwitchOk returns a tuple with the AssociatedSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetAssociatedSwitchOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.AssociatedSwitch == nil {
		return nil, false
	}
	return o.AssociatedSwitch, true
}

// HasAssociatedSwitch returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasAssociatedSwitch() bool {
	if o != nil && o.AssociatedSwitch != nil {
		return true
	}

	return false
}

// SetAssociatedSwitch gets a reference to the given NetworkElementRelationship and assigns it to the AssociatedSwitch field.
func (o *FabricSwitchProfile) SetAssociatedSwitch(v NetworkElementRelationship) {
	o.AssociatedSwitch = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfile) GetConfigChangeDetails() []FabricConfigChangeDetailRelationship {
	if o == nil {
		var ret []FabricConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfile) GetConfigChangeDetailsOk() (*[]FabricConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []FabricConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *FabricSwitchProfile) SetConfigChangeDetails(v []FabricConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *FabricSwitchProfile) GetConfigResult() FabricConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret FabricConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetConfigResultOk() (*FabricConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given FabricConfigResultRelationship and assigns it to the ConfigResult field.
func (o *FabricSwitchProfile) SetConfigResult(v FabricConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *FabricSwitchProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

// GetSwitchClusterProfile returns the SwitchClusterProfile field value if set, zero value otherwise.
func (o *FabricSwitchProfile) GetSwitchClusterProfile() FabricSwitchClusterProfileRelationship {
	if o == nil || o.SwitchClusterProfile == nil {
		var ret FabricSwitchClusterProfileRelationship
		return ret
	}
	return *o.SwitchClusterProfile
}

// GetSwitchClusterProfileOk returns a tuple with the SwitchClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfile) GetSwitchClusterProfileOk() (*FabricSwitchClusterProfileRelationship, bool) {
	if o == nil || o.SwitchClusterProfile == nil {
		return nil, false
	}
	return o.SwitchClusterProfile, true
}

// HasSwitchClusterProfile returns a boolean if a field has been set.
func (o *FabricSwitchProfile) HasSwitchClusterProfile() bool {
	if o != nil && o.SwitchClusterProfile != nil {
		return true
	}

	return false
}

// SetSwitchClusterProfile gets a reference to the given FabricSwitchClusterProfileRelationship and assigns it to the SwitchClusterProfile field.
func (o *FabricSwitchProfile) SetSwitchClusterProfile(v FabricSwitchClusterProfileRelationship) {
	o.SwitchClusterProfile = &v
}

func (o FabricSwitchProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.AssignedSwitch != nil {
		toSerialize["AssignedSwitch"] = o.AssignedSwitch
	}
	if o.AssociatedSwitch != nil {
		toSerialize["AssociatedSwitch"] = o.AssociatedSwitch
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}
	if o.SwitchClusterProfile != nil {
		toSerialize["SwitchClusterProfile"] = o.SwitchClusterProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSwitchProfile) UnmarshalJSON(bytes []byte) (err error) {
	type FabricSwitchProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                      `json:"ObjectType"`
		ConfigChanges    NullablePolicyConfigChange  `json:"ConfigChanges,omitempty"`
		AssignedSwitch   *NetworkElementRelationship `json:"AssignedSwitch,omitempty"`
		AssociatedSwitch *NetworkElementRelationship `json:"AssociatedSwitch,omitempty"`
		// An array of relationships to fabricConfigChangeDetail resources.
		ConfigChangeDetails []FabricConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
		ConfigResult        *FabricConfigResultRelationship        `json:"ConfigResult,omitempty"`
		// An array of relationships to workflowWorkflowInfo resources.
		RunningWorkflows     []WorkflowWorkflowInfoRelationship      `json:"RunningWorkflows,omitempty"`
		SwitchClusterProfile *FabricSwitchClusterProfileRelationship `json:"SwitchClusterProfile,omitempty"`
	}

	varFabricSwitchProfileWithoutEmbeddedStruct := FabricSwitchProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricSwitchProfileWithoutEmbeddedStruct)
	if err == nil {
		varFabricSwitchProfile := _FabricSwitchProfile{}
		varFabricSwitchProfile.ClassId = varFabricSwitchProfileWithoutEmbeddedStruct.ClassId
		varFabricSwitchProfile.ObjectType = varFabricSwitchProfileWithoutEmbeddedStruct.ObjectType
		varFabricSwitchProfile.ConfigChanges = varFabricSwitchProfileWithoutEmbeddedStruct.ConfigChanges
		varFabricSwitchProfile.AssignedSwitch = varFabricSwitchProfileWithoutEmbeddedStruct.AssignedSwitch
		varFabricSwitchProfile.AssociatedSwitch = varFabricSwitchProfileWithoutEmbeddedStruct.AssociatedSwitch
		varFabricSwitchProfile.ConfigChangeDetails = varFabricSwitchProfileWithoutEmbeddedStruct.ConfigChangeDetails
		varFabricSwitchProfile.ConfigResult = varFabricSwitchProfileWithoutEmbeddedStruct.ConfigResult
		varFabricSwitchProfile.RunningWorkflows = varFabricSwitchProfileWithoutEmbeddedStruct.RunningWorkflows
		varFabricSwitchProfile.SwitchClusterProfile = varFabricSwitchProfileWithoutEmbeddedStruct.SwitchClusterProfile
		*o = FabricSwitchProfile(varFabricSwitchProfile)
	} else {
		return err
	}

	varFabricSwitchProfile := _FabricSwitchProfile{}

	err = json.Unmarshal(bytes, &varFabricSwitchProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varFabricSwitchProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "AssignedSwitch")
		delete(additionalProperties, "AssociatedSwitch")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "RunningWorkflows")
		delete(additionalProperties, "SwitchClusterProfile")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableFabricSwitchProfile struct {
	value *FabricSwitchProfile
	isSet bool
}

func (v NullableFabricSwitchProfile) Get() *FabricSwitchProfile {
	return v.value
}

func (v *NullableFabricSwitchProfile) Set(val *FabricSwitchProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchProfile(val *FabricSwitchProfile) *NullableFabricSwitchProfile {
	return &NullableFabricSwitchProfile{value: val, isSet: true}
}

func (v NullableFabricSwitchProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
