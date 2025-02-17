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

// RecoveryRestore Triggers a restore operation on the target endpoint.
type RecoveryRestore struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	ConfigParams         NullableRecoveryConfigParams            `json:"ConfigParams,omitempty"`
	BackupInfo           *RecoveryAbstractBackupInfoRelationship `json:"BackupInfo,omitempty"`
	Device               *AssetDeviceRegistrationRelationship    `json:"Device,omitempty"`
	Organization         *OrganizationOrganizationRelationship   `json:"Organization,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship       `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryRestore RecoveryRestore

// NewRecoveryRestore instantiates a new RecoveryRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryRestore(classId string, objectType string) *RecoveryRestore {
	this := RecoveryRestore{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryRestoreWithDefaults instantiates a new RecoveryRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryRestoreWithDefaults() *RecoveryRestore {
	this := RecoveryRestore{}
	var classId string = "recovery.Restore"
	this.ClassId = classId
	var objectType string = "recovery.Restore"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryRestore) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryRestore) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryRestore) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryRestore) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetConfigParams() RecoveryConfigParams {
	if o == nil || o.ConfigParams.Get() == nil {
		var ret RecoveryConfigParams
		return ret
	}
	return *o.ConfigParams.Get()
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetConfigParamsOk() (*RecoveryConfigParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigParams.Get(), o.ConfigParams.IsSet()
}

// HasConfigParams returns a boolean if a field has been set.
func (o *RecoveryRestore) HasConfigParams() bool {
	if o != nil && o.ConfigParams.IsSet() {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given NullableRecoveryConfigParams and assigns it to the ConfigParams field.
func (o *RecoveryRestore) SetConfigParams(v RecoveryConfigParams) {
	o.ConfigParams.Set(&v)
}

// SetConfigParamsNil sets the value for ConfigParams to be an explicit nil
func (o *RecoveryRestore) SetConfigParamsNil() {
	o.ConfigParams.Set(nil)
}

// UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
func (o *RecoveryRestore) UnsetConfigParams() {
	o.ConfigParams.Unset()
}

// GetBackupInfo returns the BackupInfo field value if set, zero value otherwise.
func (o *RecoveryRestore) GetBackupInfo() RecoveryAbstractBackupInfoRelationship {
	if o == nil || o.BackupInfo == nil {
		var ret RecoveryAbstractBackupInfoRelationship
		return ret
	}
	return *o.BackupInfo
}

// GetBackupInfoOk returns a tuple with the BackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetBackupInfoOk() (*RecoveryAbstractBackupInfoRelationship, bool) {
	if o == nil || o.BackupInfo == nil {
		return nil, false
	}
	return o.BackupInfo, true
}

// HasBackupInfo returns a boolean if a field has been set.
func (o *RecoveryRestore) HasBackupInfo() bool {
	if o != nil && o.BackupInfo != nil {
		return true
	}

	return false
}

// SetBackupInfo gets a reference to the given RecoveryAbstractBackupInfoRelationship and assigns it to the BackupInfo field.
func (o *RecoveryRestore) SetBackupInfo(v RecoveryAbstractBackupInfoRelationship) {
	o.BackupInfo = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *RecoveryRestore) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *RecoveryRestore) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *RecoveryRestore) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryRestore) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryRestore) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryRestore) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *RecoveryRestore) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *RecoveryRestore) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *RecoveryRestore) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o RecoveryRestore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigParams.IsSet() {
		toSerialize["ConfigParams"] = o.ConfigParams.Get()
	}
	if o.BackupInfo != nil {
		toSerialize["BackupInfo"] = o.BackupInfo
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryRestore) UnmarshalJSON(bytes []byte) (err error) {
	type RecoveryRestoreWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                  `json:"ObjectType"`
		ConfigParams NullableRecoveryConfigParams            `json:"ConfigParams,omitempty"`
		BackupInfo   *RecoveryAbstractBackupInfoRelationship `json:"BackupInfo,omitempty"`
		Device       *AssetDeviceRegistrationRelationship    `json:"Device,omitempty"`
		Organization *OrganizationOrganizationRelationship   `json:"Organization,omitempty"`
		Workflow     *WorkflowWorkflowInfoRelationship       `json:"Workflow,omitempty"`
	}

	varRecoveryRestoreWithoutEmbeddedStruct := RecoveryRestoreWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecoveryRestoreWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryRestore := _RecoveryRestore{}
		varRecoveryRestore.ClassId = varRecoveryRestoreWithoutEmbeddedStruct.ClassId
		varRecoveryRestore.ObjectType = varRecoveryRestoreWithoutEmbeddedStruct.ObjectType
		varRecoveryRestore.ConfigParams = varRecoveryRestoreWithoutEmbeddedStruct.ConfigParams
		varRecoveryRestore.BackupInfo = varRecoveryRestoreWithoutEmbeddedStruct.BackupInfo
		varRecoveryRestore.Device = varRecoveryRestoreWithoutEmbeddedStruct.Device
		varRecoveryRestore.Organization = varRecoveryRestoreWithoutEmbeddedStruct.Organization
		varRecoveryRestore.Workflow = varRecoveryRestoreWithoutEmbeddedStruct.Workflow
		*o = RecoveryRestore(varRecoveryRestore)
	} else {
		return err
	}

	varRecoveryRestore := _RecoveryRestore{}

	err = json.Unmarshal(bytes, &varRecoveryRestore)
	if err == nil {
		o.MoBaseMo = varRecoveryRestore.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigParams")
		delete(additionalProperties, "BackupInfo")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Workflow")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableRecoveryRestore struct {
	value *RecoveryRestore
	isSet bool
}

func (v NullableRecoveryRestore) Get() *RecoveryRestore {
	return v.value
}

func (v *NullableRecoveryRestore) Set(val *RecoveryRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryRestore(val *RecoveryRestore) *NullableRecoveryRestore {
	return &NullableRecoveryRestore{value: val, isSet: true}
}

func (v NullableRecoveryRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
