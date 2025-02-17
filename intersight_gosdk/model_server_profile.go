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

// ServerProfile A profile specifying configuration settings for a physical server.
type ServerProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType          string                            `json:"ObjectType"`
	ConfigChangeContext NullablePolicyConfigChangeContext `json:"ConfigChangeContext,omitempty"`
	ConfigChanges       NullablePolicyConfigChange        `json:"ConfigChanges,omitempty"`
	// Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set.
	IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`
	// Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy.
	PmcDeployedSecurePassphrase *string `json:"PmcDeployedSecurePassphrase,omitempty"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform   *string                      `json:"TargetPlatform,omitempty"`
	AssignedServer   *ComputePhysicalRelationship `json:"AssignedServer,omitempty"`
	AssociatedServer *ComputePhysicalRelationship `json:"AssociatedServer,omitempty"`
	// An array of relationships to serverConfigChangeDetail resources.
	ConfigChangeDetails []ServerConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	ConfigResult        *ServerConfigResultRelationship        `json:"ConfigResult,omitempty"`
	Organization        *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerProfile ServerProfile

// NewServerProfile instantiates a new ServerProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProfile(classId string, objectType string) *ServerProfile {
	this := ServerProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPmcDeployedSecurePassphraseSet bool = false
	this.IsPmcDeployedSecurePassphraseSet = &isPmcDeployedSecurePassphraseSet
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewServerProfileWithDefaults instantiates a new ServerProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProfileWithDefaults() *ServerProfile {
	this := ServerProfile{}
	var classId string = "server.Profile"
	this.ClassId = classId
	var objectType string = "server.Profile"
	this.ObjectType = objectType
	var isPmcDeployedSecurePassphraseSet bool = false
	this.IsPmcDeployedSecurePassphraseSet = &isPmcDeployedSecurePassphraseSet
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChangeContext returns the ConfigChangeContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfile) GetConfigChangeContext() PolicyConfigChangeContext {
	if o == nil || o.ConfigChangeContext.Get() == nil {
		var ret PolicyConfigChangeContext
		return ret
	}
	return *o.ConfigChangeContext.Get()
}

// GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfile) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChangeContext.Get(), o.ConfigChangeContext.IsSet()
}

// HasConfigChangeContext returns a boolean if a field has been set.
func (o *ServerProfile) HasConfigChangeContext() bool {
	if o != nil && o.ConfigChangeContext.IsSet() {
		return true
	}

	return false
}

// SetConfigChangeContext gets a reference to the given NullablePolicyConfigChangeContext and assigns it to the ConfigChangeContext field.
func (o *ServerProfile) SetConfigChangeContext(v PolicyConfigChangeContext) {
	o.ConfigChangeContext.Set(&v)
}

// SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil
func (o *ServerProfile) SetConfigChangeContextNil() {
	o.ConfigChangeContext.Set(nil)
}

// UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
func (o *ServerProfile) UnsetConfigChangeContext() {
	o.ConfigChangeContext.Unset()
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfile) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfile) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ServerProfile) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ServerProfile) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ServerProfile) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ServerProfile) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetIsPmcDeployedSecurePassphraseSet returns the IsPmcDeployedSecurePassphraseSet field value if set, zero value otherwise.
func (o *ServerProfile) GetIsPmcDeployedSecurePassphraseSet() bool {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPmcDeployedSecurePassphraseSet
}

// GetIsPmcDeployedSecurePassphraseSetOk returns a tuple with the IsPmcDeployedSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetIsPmcDeployedSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsPmcDeployedSecurePassphraseSet, true
}

// HasIsPmcDeployedSecurePassphraseSet returns a boolean if a field has been set.
func (o *ServerProfile) HasIsPmcDeployedSecurePassphraseSet() bool {
	if o != nil && o.IsPmcDeployedSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPmcDeployedSecurePassphraseSet gets a reference to the given bool and assigns it to the IsPmcDeployedSecurePassphraseSet field.
func (o *ServerProfile) SetIsPmcDeployedSecurePassphraseSet(v bool) {
	o.IsPmcDeployedSecurePassphraseSet = &v
}

// GetPmcDeployedSecurePassphrase returns the PmcDeployedSecurePassphrase field value if set, zero value otherwise.
func (o *ServerProfile) GetPmcDeployedSecurePassphrase() string {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.PmcDeployedSecurePassphrase
}

// GetPmcDeployedSecurePassphraseOk returns a tuple with the PmcDeployedSecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetPmcDeployedSecurePassphraseOk() (*string, bool) {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		return nil, false
	}
	return o.PmcDeployedSecurePassphrase, true
}

// HasPmcDeployedSecurePassphrase returns a boolean if a field has been set.
func (o *ServerProfile) HasPmcDeployedSecurePassphrase() bool {
	if o != nil && o.PmcDeployedSecurePassphrase != nil {
		return true
	}

	return false
}

// SetPmcDeployedSecurePassphrase gets a reference to the given string and assigns it to the PmcDeployedSecurePassphrase field.
func (o *ServerProfile) SetPmcDeployedSecurePassphrase(v string) {
	o.PmcDeployedSecurePassphrase = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ServerProfile) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ServerProfile) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ServerProfile) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetAssignedServer returns the AssignedServer field value if set, zero value otherwise.
func (o *ServerProfile) GetAssignedServer() ComputePhysicalRelationship {
	if o == nil || o.AssignedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssignedServer
}

// GetAssignedServerOk returns a tuple with the AssignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetAssignedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssignedServer == nil {
		return nil, false
	}
	return o.AssignedServer, true
}

// HasAssignedServer returns a boolean if a field has been set.
func (o *ServerProfile) HasAssignedServer() bool {
	if o != nil && o.AssignedServer != nil {
		return true
	}

	return false
}

// SetAssignedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssignedServer field.
func (o *ServerProfile) SetAssignedServer(v ComputePhysicalRelationship) {
	o.AssignedServer = &v
}

// GetAssociatedServer returns the AssociatedServer field value if set, zero value otherwise.
func (o *ServerProfile) GetAssociatedServer() ComputePhysicalRelationship {
	if o == nil || o.AssociatedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssociatedServer
}

// GetAssociatedServerOk returns a tuple with the AssociatedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetAssociatedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssociatedServer == nil {
		return nil, false
	}
	return o.AssociatedServer, true
}

// HasAssociatedServer returns a boolean if a field has been set.
func (o *ServerProfile) HasAssociatedServer() bool {
	if o != nil && o.AssociatedServer != nil {
		return true
	}

	return false
}

// SetAssociatedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssociatedServer field.
func (o *ServerProfile) SetAssociatedServer(v ComputePhysicalRelationship) {
	o.AssociatedServer = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfile) GetConfigChangeDetails() []ServerConfigChangeDetailRelationship {
	if o == nil {
		var ret []ServerConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfile) GetConfigChangeDetailsOk() (*[]ServerConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ServerProfile) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ServerConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ServerProfile) SetConfigChangeDetails(v []ServerConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ServerProfile) GetConfigResult() ServerConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ServerConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetConfigResultOk() (*ServerConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ServerProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ServerConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ServerProfile) SetConfigResult(v ServerConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ServerProfile) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ServerProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

func (o ServerProfile) MarshalJSON() ([]byte, error) {
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
	if o.ConfigChangeContext.IsSet() {
		toSerialize["ConfigChangeContext"] = o.ConfigChangeContext.Get()
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.IsPmcDeployedSecurePassphraseSet != nil {
		toSerialize["IsPmcDeployedSecurePassphraseSet"] = o.IsPmcDeployedSecurePassphraseSet
	}
	if o.PmcDeployedSecurePassphrase != nil {
		toSerialize["PmcDeployedSecurePassphrase"] = o.PmcDeployedSecurePassphrase
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.AssignedServer != nil {
		toSerialize["AssignedServer"] = o.AssignedServer
	}
	if o.AssociatedServer != nil {
		toSerialize["AssociatedServer"] = o.AssociatedServer
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerProfile) UnmarshalJSON(bytes []byte) (err error) {
	type ServerProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                            `json:"ObjectType"`
		ConfigChangeContext NullablePolicyConfigChangeContext `json:"ConfigChangeContext,omitempty"`
		ConfigChanges       NullablePolicyConfigChange        `json:"ConfigChanges,omitempty"`
		// Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set.
		IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`
		// Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy.
		PmcDeployedSecurePassphrase *string `json:"PmcDeployedSecurePassphrase,omitempty"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform   *string                      `json:"TargetPlatform,omitempty"`
		AssignedServer   *ComputePhysicalRelationship `json:"AssignedServer,omitempty"`
		AssociatedServer *ComputePhysicalRelationship `json:"AssociatedServer,omitempty"`
		// An array of relationships to serverConfigChangeDetail resources.
		ConfigChangeDetails []ServerConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
		ConfigResult        *ServerConfigResultRelationship        `json:"ConfigResult,omitempty"`
		Organization        *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
		// An array of relationships to workflowWorkflowInfo resources.
		RunningWorkflows []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	}

	varServerProfileWithoutEmbeddedStruct := ServerProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServerProfileWithoutEmbeddedStruct)
	if err == nil {
		varServerProfile := _ServerProfile{}
		varServerProfile.ClassId = varServerProfileWithoutEmbeddedStruct.ClassId
		varServerProfile.ObjectType = varServerProfileWithoutEmbeddedStruct.ObjectType
		varServerProfile.ConfigChangeContext = varServerProfileWithoutEmbeddedStruct.ConfigChangeContext
		varServerProfile.ConfigChanges = varServerProfileWithoutEmbeddedStruct.ConfigChanges
		varServerProfile.IsPmcDeployedSecurePassphraseSet = varServerProfileWithoutEmbeddedStruct.IsPmcDeployedSecurePassphraseSet
		varServerProfile.PmcDeployedSecurePassphrase = varServerProfileWithoutEmbeddedStruct.PmcDeployedSecurePassphrase
		varServerProfile.TargetPlatform = varServerProfileWithoutEmbeddedStruct.TargetPlatform
		varServerProfile.AssignedServer = varServerProfileWithoutEmbeddedStruct.AssignedServer
		varServerProfile.AssociatedServer = varServerProfileWithoutEmbeddedStruct.AssociatedServer
		varServerProfile.ConfigChangeDetails = varServerProfileWithoutEmbeddedStruct.ConfigChangeDetails
		varServerProfile.ConfigResult = varServerProfileWithoutEmbeddedStruct.ConfigResult
		varServerProfile.Organization = varServerProfileWithoutEmbeddedStruct.Organization
		varServerProfile.RunningWorkflows = varServerProfileWithoutEmbeddedStruct.RunningWorkflows
		*o = ServerProfile(varServerProfile)
	} else {
		return err
	}

	varServerProfile := _ServerProfile{}

	err = json.Unmarshal(bytes, &varServerProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varServerProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChangeContext")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "IsPmcDeployedSecurePassphraseSet")
		delete(additionalProperties, "PmcDeployedSecurePassphrase")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "AssignedServer")
		delete(additionalProperties, "AssociatedServer")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RunningWorkflows")

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

type NullableServerProfile struct {
	value *ServerProfile
	isSet bool
}

func (v NullableServerProfile) Get() *ServerProfile {
	return v.value
}

func (v *NullableServerProfile) Set(val *ServerProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProfile(val *ServerProfile) *NullableServerProfile {
	return &NullableServerProfile{value: val, isSet: true}
}

func (v NullableServerProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
