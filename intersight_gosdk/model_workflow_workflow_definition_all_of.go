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

// WorkflowWorkflowDefinitionAllOf Definition of the list of properties defined in 'workflow.WorkflowDefinition', excluding properties defined in parent classes.
type WorkflowWorkflowDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty"`
	// The description for this workflow.
	Description       *string                `json:"Description,omitempty"`
	InputDefinition   []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	InputParameterSet []WorkflowParameterSet `json:"InputParameterSet,omitempty"`
	// A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type.
	LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
	// The maximum number of tasks that can be executed on this workflow.
	MaxTaskCount *int64 `json:"MaxTaskCount,omitempty"`
	// The maximum number of external (worker) tasks that can be executed on this workflow.
	MaxWorkerTaskCount *int64 `json:"MaxWorkerTaskCount,omitempty"`
	// The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
	Name             *string                `json:"Name,omitempty"`
	OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output.
	OutputParameters interface{}                        `json:"OutputParameters,omitempty"`
	Properties       NullableWorkflowWorkflowProperties `json:"Properties,omitempty"`
	Tasks            []WorkflowWorkflowTask             `json:"Tasks,omitempty"`
	UiInputFilters   []WorkflowUiInputFilter            `json:"UiInputFilters,omitempty"`
	// This will hold the data needed for workflow to be rendered in the user interface.
	UiRenderingData       interface{}                           `json:"UiRenderingData,omitempty"`
	ValidationInformation NullableWorkflowValidationInformation `json:"ValidationInformation,omitempty"`
	// The version of the workflow to support multiple versions.
	Version              *int64                                `json:"Version,omitempty"`
	Catalog              *WorkflowCatalogRelationship          `json:"Catalog,omitempty"`
	WorkflowMetadata     *WorkflowWorkflowMetadataRelationship `json:"WorkflowMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowDefinitionAllOf WorkflowWorkflowDefinitionAllOf

// NewWorkflowWorkflowDefinitionAllOf instantiates a new WorkflowWorkflowDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowDefinitionAllOf(classId string, objectType string) *WorkflowWorkflowDefinitionAllOf {
	this := WorkflowWorkflowDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	var version int64 = 1
	this.Version = &version
	return &this
}

// NewWorkflowWorkflowDefinitionAllOfWithDefaults instantiates a new WorkflowWorkflowDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowDefinitionAllOfWithDefaults() *WorkflowWorkflowDefinitionAllOf {
	this := WorkflowWorkflowDefinitionAllOf{}
	var classId string = "workflow.WorkflowDefinition"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowDefinition"
	this.ObjectType = objectType
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	var version int64 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultVersion returns the DefaultVersion field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetDefaultVersion() bool {
	if o == nil || o.DefaultVersion == nil {
		var ret bool
		return ret
	}
	return *o.DefaultVersion
}

// GetDefaultVersionOk returns a tuple with the DefaultVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetDefaultVersionOk() (*bool, bool) {
	if o == nil || o.DefaultVersion == nil {
		return nil, false
	}
	return o.DefaultVersion, true
}

// HasDefaultVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasDefaultVersion() bool {
	if o != nil && o.DefaultVersion != nil {
		return true
	}

	return false
}

// SetDefaultVersion gets a reference to the given bool and assigns it to the DefaultVersion field.
func (o *WorkflowWorkflowDefinitionAllOf) SetDefaultVersion(v bool) {
	o.DefaultVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWorkflowDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return &o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowWorkflowDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetInputParameterSet returns the InputParameterSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetInputParameterSet() []WorkflowParameterSet {
	if o == nil {
		var ret []WorkflowParameterSet
		return ret
	}
	return o.InputParameterSet
}

// GetInputParameterSetOk returns a tuple with the InputParameterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetInputParameterSetOk() (*[]WorkflowParameterSet, bool) {
	if o == nil || o.InputParameterSet == nil {
		return nil, false
	}
	return &o.InputParameterSet, true
}

// HasInputParameterSet returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasInputParameterSet() bool {
	if o != nil && o.InputParameterSet != nil {
		return true
	}

	return false
}

// SetInputParameterSet gets a reference to the given []WorkflowParameterSet and assigns it to the InputParameterSet field.
func (o *WorkflowWorkflowDefinitionAllOf) SetInputParameterSet(v []WorkflowParameterSet) {
	o.InputParameterSet = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWorkflowDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetLicenseEntitlement() string {
	if o == nil || o.LicenseEntitlement == nil {
		var ret string
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool) {
	if o == nil || o.LicenseEntitlement == nil {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasLicenseEntitlement() bool {
	if o != nil && o.LicenseEntitlement != nil {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given string and assigns it to the LicenseEntitlement field.
func (o *WorkflowWorkflowDefinitionAllOf) SetLicenseEntitlement(v string) {
	o.LicenseEntitlement = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetMaxTaskCount() int64 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetMaxTaskCountOk() (*int64, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int64 and assigns it to the MaxTaskCount field.
func (o *WorkflowWorkflowDefinitionAllOf) SetMaxTaskCount(v int64) {
	o.MaxTaskCount = &v
}

// GetMaxWorkerTaskCount returns the MaxWorkerTaskCount field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetMaxWorkerTaskCount() int64 {
	if o == nil || o.MaxWorkerTaskCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxWorkerTaskCount
}

// GetMaxWorkerTaskCountOk returns a tuple with the MaxWorkerTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetMaxWorkerTaskCountOk() (*int64, bool) {
	if o == nil || o.MaxWorkerTaskCount == nil {
		return nil, false
	}
	return o.MaxWorkerTaskCount, true
}

// HasMaxWorkerTaskCount returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasMaxWorkerTaskCount() bool {
	if o != nil && o.MaxWorkerTaskCount != nil {
		return true
	}

	return false
}

// SetMaxWorkerTaskCount gets a reference to the given int64 and assigns it to the MaxWorkerTaskCount field.
func (o *WorkflowWorkflowDefinitionAllOf) SetMaxWorkerTaskCount(v int64) {
	o.MaxWorkerTaskCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWorkflowDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return &o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowWorkflowDefinitionAllOf) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || o.OutputParameters == nil {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasOutputParameters() bool {
	if o != nil && o.OutputParameters != nil {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *WorkflowWorkflowDefinitionAllOf) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetProperties() WorkflowWorkflowProperties {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowWorkflowProperties
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetPropertiesOk() (*WorkflowWorkflowProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowWorkflowProperties and assigns it to the Properties field.
func (o *WorkflowWorkflowDefinitionAllOf) SetProperties(v WorkflowWorkflowProperties) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowWorkflowDefinitionAllOf) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowWorkflowDefinitionAllOf) UnsetProperties() {
	o.Properties.Unset()
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetTasks() []WorkflowWorkflowTask {
	if o == nil {
		var ret []WorkflowWorkflowTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetTasksOk() (*[]WorkflowWorkflowTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []WorkflowWorkflowTask and assigns it to the Tasks field.
func (o *WorkflowWorkflowDefinitionAllOf) SetTasks(v []WorkflowWorkflowTask) {
	o.Tasks = v
}

// GetUiInputFilters returns the UiInputFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetUiInputFilters() []WorkflowUiInputFilter {
	if o == nil {
		var ret []WorkflowUiInputFilter
		return ret
	}
	return o.UiInputFilters
}

// GetUiInputFiltersOk returns a tuple with the UiInputFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetUiInputFiltersOk() (*[]WorkflowUiInputFilter, bool) {
	if o == nil || o.UiInputFilters == nil {
		return nil, false
	}
	return &o.UiInputFilters, true
}

// HasUiInputFilters returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasUiInputFilters() bool {
	if o != nil && o.UiInputFilters != nil {
		return true
	}

	return false
}

// SetUiInputFilters gets a reference to the given []WorkflowUiInputFilter and assigns it to the UiInputFilters field.
func (o *WorkflowWorkflowDefinitionAllOf) SetUiInputFilters(v []WorkflowUiInputFilter) {
	o.UiInputFilters = v
}

// GetUiRenderingData returns the UiRenderingData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetUiRenderingData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiRenderingData
}

// GetUiRenderingDataOk returns a tuple with the UiRenderingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetUiRenderingDataOk() (*interface{}, bool) {
	if o == nil || o.UiRenderingData == nil {
		return nil, false
	}
	return &o.UiRenderingData, true
}

// HasUiRenderingData returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasUiRenderingData() bool {
	if o != nil && o.UiRenderingData != nil {
		return true
	}

	return false
}

// SetUiRenderingData gets a reference to the given interface{} and assigns it to the UiRenderingData field.
func (o *WorkflowWorkflowDefinitionAllOf) SetUiRenderingData(v interface{}) {
	o.UiRenderingData = v
}

// GetValidationInformation returns the ValidationInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation {
	if o == nil || o.ValidationInformation.Get() == nil {
		var ret WorkflowValidationInformation
		return ret
	}
	return *o.ValidationInformation.Get()
}

// GetValidationInformationOk returns a tuple with the ValidationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationInformation.Get(), o.ValidationInformation.IsSet()
}

// HasValidationInformation returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasValidationInformation() bool {
	if o != nil && o.ValidationInformation.IsSet() {
		return true
	}

	return false
}

// SetValidationInformation gets a reference to the given NullableWorkflowValidationInformation and assigns it to the ValidationInformation field.
func (o *WorkflowWorkflowDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation) {
	o.ValidationInformation.Set(&v)
}

// SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil
func (o *WorkflowWorkflowDefinitionAllOf) SetValidationInformationNil() {
	o.ValidationInformation.Set(nil)
}

// UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
func (o *WorkflowWorkflowDefinitionAllOf) UnsetValidationInformation() {
	o.ValidationInformation.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowWorkflowDefinitionAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowWorkflowDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

// GetWorkflowMetadata returns the WorkflowMetadata field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinitionAllOf) GetWorkflowMetadata() WorkflowWorkflowMetadataRelationship {
	if o == nil || o.WorkflowMetadata == nil {
		var ret WorkflowWorkflowMetadataRelationship
		return ret
	}
	return *o.WorkflowMetadata
}

// GetWorkflowMetadataOk returns a tuple with the WorkflowMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinitionAllOf) GetWorkflowMetadataOk() (*WorkflowWorkflowMetadataRelationship, bool) {
	if o == nil || o.WorkflowMetadata == nil {
		return nil, false
	}
	return o.WorkflowMetadata, true
}

// HasWorkflowMetadata returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinitionAllOf) HasWorkflowMetadata() bool {
	if o != nil && o.WorkflowMetadata != nil {
		return true
	}

	return false
}

// SetWorkflowMetadata gets a reference to the given WorkflowWorkflowMetadataRelationship and assigns it to the WorkflowMetadata field.
func (o *WorkflowWorkflowDefinitionAllOf) SetWorkflowMetadata(v WorkflowWorkflowMetadataRelationship) {
	o.WorkflowMetadata = &v
}

func (o WorkflowWorkflowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DefaultVersion != nil {
		toSerialize["DefaultVersion"] = o.DefaultVersion
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.InputParameterSet != nil {
		toSerialize["InputParameterSet"] = o.InputParameterSet
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.LicenseEntitlement != nil {
		toSerialize["LicenseEntitlement"] = o.LicenseEntitlement
	}
	if o.MaxTaskCount != nil {
		toSerialize["MaxTaskCount"] = o.MaxTaskCount
	}
	if o.MaxWorkerTaskCount != nil {
		toSerialize["MaxWorkerTaskCount"] = o.MaxWorkerTaskCount
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.OutputParameters != nil {
		toSerialize["OutputParameters"] = o.OutputParameters
	}
	if o.Properties.IsSet() {
		toSerialize["Properties"] = o.Properties.Get()
	}
	if o.Tasks != nil {
		toSerialize["Tasks"] = o.Tasks
	}
	if o.UiInputFilters != nil {
		toSerialize["UiInputFilters"] = o.UiInputFilters
	}
	if o.UiRenderingData != nil {
		toSerialize["UiRenderingData"] = o.UiRenderingData
	}
	if o.ValidationInformation.IsSet() {
		toSerialize["ValidationInformation"] = o.ValidationInformation.Get()
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.WorkflowMetadata != nil {
		toSerialize["WorkflowMetadata"] = o.WorkflowMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWorkflowDefinitionAllOf := _WorkflowWorkflowDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWorkflowDefinitionAllOf); err == nil {
		*o = WorkflowWorkflowDefinitionAllOf(varWorkflowWorkflowDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultVersion")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "InputParameterSet")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "LicenseEntitlement")
		delete(additionalProperties, "MaxTaskCount")
		delete(additionalProperties, "MaxWorkerTaskCount")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "OutputParameters")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "Tasks")
		delete(additionalProperties, "UiInputFilters")
		delete(additionalProperties, "UiRenderingData")
		delete(additionalProperties, "ValidationInformation")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "WorkflowMetadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWorkflowDefinitionAllOf struct {
	value *WorkflowWorkflowDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowWorkflowDefinitionAllOf) Get() *WorkflowWorkflowDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowWorkflowDefinitionAllOf) Set(val *WorkflowWorkflowDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowDefinitionAllOf(val *WorkflowWorkflowDefinitionAllOf) *NullableWorkflowWorkflowDefinitionAllOf {
	return &NullableWorkflowWorkflowDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
