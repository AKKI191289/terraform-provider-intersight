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

// ServerProfileAllOf Definition of the list of properties defined in 'server.Profile', excluding properties defined in parent classes.
type ServerProfileAllOf struct {
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

type _ServerProfileAllOf ServerProfileAllOf

// NewServerProfileAllOf instantiates a new ServerProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProfileAllOf(classId string, objectType string) *ServerProfileAllOf {
	this := ServerProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPmcDeployedSecurePassphraseSet bool = false
	this.IsPmcDeployedSecurePassphraseSet = &isPmcDeployedSecurePassphraseSet
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewServerProfileAllOfWithDefaults instantiates a new ServerProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProfileAllOfWithDefaults() *ServerProfileAllOf {
	this := ServerProfileAllOf{}
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
func (o *ServerProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChangeContext returns the ConfigChangeContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChangeContext() PolicyConfigChangeContext {
	if o == nil || o.ConfigChangeContext.Get() == nil {
		var ret PolicyConfigChangeContext
		return ret
	}
	return *o.ConfigChangeContext.Get()
}

// GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChangeContext.Get(), o.ConfigChangeContext.IsSet()
}

// HasConfigChangeContext returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChangeContext() bool {
	if o != nil && o.ConfigChangeContext.IsSet() {
		return true
	}

	return false
}

// SetConfigChangeContext gets a reference to the given NullablePolicyConfigChangeContext and assigns it to the ConfigChangeContext field.
func (o *ServerProfileAllOf) SetConfigChangeContext(v PolicyConfigChangeContext) {
	o.ConfigChangeContext.Set(&v)
}

// SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil
func (o *ServerProfileAllOf) SetConfigChangeContextNil() {
	o.ConfigChangeContext.Set(nil)
}

// UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
func (o *ServerProfileAllOf) UnsetConfigChangeContext() {
	o.ConfigChangeContext.Unset()
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ServerProfileAllOf) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ServerProfileAllOf) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ServerProfileAllOf) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetIsPmcDeployedSecurePassphraseSet returns the IsPmcDeployedSecurePassphraseSet field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSet() bool {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPmcDeployedSecurePassphraseSet
}

// GetIsPmcDeployedSecurePassphraseSetOk returns a tuple with the IsPmcDeployedSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsPmcDeployedSecurePassphraseSet, true
}

// HasIsPmcDeployedSecurePassphraseSet returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasIsPmcDeployedSecurePassphraseSet() bool {
	if o != nil && o.IsPmcDeployedSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPmcDeployedSecurePassphraseSet gets a reference to the given bool and assigns it to the IsPmcDeployedSecurePassphraseSet field.
func (o *ServerProfileAllOf) SetIsPmcDeployedSecurePassphraseSet(v bool) {
	o.IsPmcDeployedSecurePassphraseSet = &v
}

// GetPmcDeployedSecurePassphrase returns the PmcDeployedSecurePassphrase field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphrase() string {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.PmcDeployedSecurePassphrase
}

// GetPmcDeployedSecurePassphraseOk returns a tuple with the PmcDeployedSecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphraseOk() (*string, bool) {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		return nil, false
	}
	return o.PmcDeployedSecurePassphrase, true
}

// HasPmcDeployedSecurePassphrase returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasPmcDeployedSecurePassphrase() bool {
	if o != nil && o.PmcDeployedSecurePassphrase != nil {
		return true
	}

	return false
}

// SetPmcDeployedSecurePassphrase gets a reference to the given string and assigns it to the PmcDeployedSecurePassphrase field.
func (o *ServerProfileAllOf) SetPmcDeployedSecurePassphrase(v string) {
	o.PmcDeployedSecurePassphrase = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ServerProfileAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetAssignedServer returns the AssignedServer field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetAssignedServer() ComputePhysicalRelationship {
	if o == nil || o.AssignedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssignedServer
}

// GetAssignedServerOk returns a tuple with the AssignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetAssignedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssignedServer == nil {
		return nil, false
	}
	return o.AssignedServer, true
}

// HasAssignedServer returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasAssignedServer() bool {
	if o != nil && o.AssignedServer != nil {
		return true
	}

	return false
}

// SetAssignedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssignedServer field.
func (o *ServerProfileAllOf) SetAssignedServer(v ComputePhysicalRelationship) {
	o.AssignedServer = &v
}

// GetAssociatedServer returns the AssociatedServer field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetAssociatedServer() ComputePhysicalRelationship {
	if o == nil || o.AssociatedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssociatedServer
}

// GetAssociatedServerOk returns a tuple with the AssociatedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetAssociatedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssociatedServer == nil {
		return nil, false
	}
	return o.AssociatedServer, true
}

// HasAssociatedServer returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasAssociatedServer() bool {
	if o != nil && o.AssociatedServer != nil {
		return true
	}

	return false
}

// SetAssociatedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssociatedServer field.
func (o *ServerProfileAllOf) SetAssociatedServer(v ComputePhysicalRelationship) {
	o.AssociatedServer = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChangeDetails() []ServerConfigChangeDetailRelationship {
	if o == nil {
		var ret []ServerConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangeDetailsOk() (*[]ServerConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ServerConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ServerProfileAllOf) SetConfigChangeDetails(v []ServerConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetConfigResult() ServerConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ServerConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetConfigResultOk() (*ServerConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ServerConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ServerProfileAllOf) SetConfigResult(v ServerConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ServerProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

func (o ServerProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ServerProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerProfileAllOf := _ServerProfileAllOf{}

	if err = json.Unmarshal(bytes, &varServerProfileAllOf); err == nil {
		*o = ServerProfileAllOf(varServerProfileAllOf)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerProfileAllOf struct {
	value *ServerProfileAllOf
	isSet bool
}

func (v NullableServerProfileAllOf) Get() *ServerProfileAllOf {
	return v.value
}

func (v *NullableServerProfileAllOf) Set(val *ServerProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProfileAllOf(val *ServerProfileAllOf) *NullableServerProfileAllOf {
	return &NullableServerProfileAllOf{value: val, isSet: true}
}

func (v NullableServerProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
