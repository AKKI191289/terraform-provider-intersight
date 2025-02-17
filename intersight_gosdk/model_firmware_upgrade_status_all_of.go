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

// FirmwareUpgradeStatusAllOf Definition of the list of properties defined in 'firmware.UpgradeStatus', excluding properties defined in parent classes.
type FirmwareUpgradeStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The message from the endpoint during the download.
	DownloadMessage *string `json:"DownloadMessage,omitempty"`
	// The percentage of the image downloaded in the endpoint.
	DownloadPercentage *int64 `json:"DownloadPercentage,omitempty"`
	// The image download stages. Example:downloading, flashing.
	DownloadStage *string `json:"DownloadStage,omitempty"`
	// The server power status after the upgrade request is submitted in the endpoint. * `none` - Server power status is none. * `powered on` - Server power status is powered on. * `powered off` - Server power status is powered off.
	EpPowerStatus *string `json:"EpPowerStatus,omitempty"`
	// The reason for the operation failure.
	OverallError *string `json:"OverallError,omitempty"`
	// The overall percentage of the operation.
	OverallPercentage *int64 `json:"OverallPercentage,omitempty"`
	// The overall status of the operation. * `none` - Upgrade stage is no upgrade stage. * `started` - Upgrade stage is started. * `prepare initiating` - Upgrade configuration is being prepared. * `prepare initiated` - Upgrade configuration is initiated. * `prepared` - Upgrade configuration is prepared. * `download initiating` - Upgrade stage is download initiating. * `download initiated` - Upgrade stage is download initiated. * `downloading` - Upgrade stage is downloading. * `downloaded` - Upgrade stage is downloaded. * `upgrade initiating on fabric A` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * `upgrade initiated on fabric A` - Upgrade stage is in upgrade initiated when the upgrade has started in endpoint. * `upgrading fabric A` - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * `rebooting fabric A` - Upgrade is in rebooting when the endpoint is being rebooted. * `upgraded fabric A` - Upgrade stage is in upgraded when the corresponding endpoint has completed. * `upgrade initiating on fabric B` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * `upgrade initiated on fabric B` - Upgrade stage is in upgrade initiated when upgrade has started in endpoint. * `upgrading fabric B` - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * `rebooting fabric B` - Upgrade is in rebooting when the endpoint is being rebooted. * `upgraded fabric B` - Upgrade stage is in upgraded when the corresponding endpoint has completed. * `upgrade initiating` - Upgrade stage is upgrade initiating. * `upgrade initiated` - Upgrade stage is upgrade initiated. * `upgrading` - Upgrade stage is upgrading. * `oob images staging` - Out-of-band component images staging. * `oob images staged` - Out-of-band component images staged. * `rebooting` - Upgrade is rebooting the endpoint. * `upgraded` - Upgrade stage is upgraded. * `success` - Upgrade stage is success. * `failed` - Upgrade stage is upgrade failed. * `terminated` - Upgrade stage is terminated. * `pending` - Upgrade stage is pending. * `ReadyForCache` - The image is ready to be cached into the Intersight Appliance. * `Caching` - The image will be cached into Intersight Appliance or an endpoint cache. * `Cached` - The image has been cached into the Intersight Appliance or endpoint cache. * `CachingFailed` - The image caching into the Intersight Appliance failed or endpoint cache.
	Overallstatus *string `json:"Overallstatus,omitempty"`
	// Pending reason for the upgrade waiting. * `none` - Upgrade pending reason is none. * `pending for next reboot` - Upgrade pending reason is pending for next reboot.
	PendingType          *string                           `json:"PendingType,omitempty"`
	Upgrade              *FirmwareUpgradeBaseRelationship  `json:"Upgrade,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeStatusAllOf FirmwareUpgradeStatusAllOf

// NewFirmwareUpgradeStatusAllOf instantiates a new FirmwareUpgradeStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeStatusAllOf(classId string, objectType string) *FirmwareUpgradeStatusAllOf {
	this := FirmwareUpgradeStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// NewFirmwareUpgradeStatusAllOfWithDefaults instantiates a new FirmwareUpgradeStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeStatusAllOfWithDefaults() *FirmwareUpgradeStatusAllOf {
	this := FirmwareUpgradeStatusAllOf{}
	var classId string = "firmware.UpgradeStatus"
	this.ClassId = classId
	var objectType string = "firmware.UpgradeStatus"
	this.ObjectType = objectType
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadMessage returns the DownloadMessage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessage() string {
	if o == nil || o.DownloadMessage == nil {
		var ret string
		return ret
	}
	return *o.DownloadMessage
}

// GetDownloadMessageOk returns a tuple with the DownloadMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessageOk() (*string, bool) {
	if o == nil || o.DownloadMessage == nil {
		return nil, false
	}
	return o.DownloadMessage, true
}

// HasDownloadMessage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadMessage() bool {
	if o != nil && o.DownloadMessage != nil {
		return true
	}

	return false
}

// SetDownloadMessage gets a reference to the given string and assigns it to the DownloadMessage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadMessage(v string) {
	o.DownloadMessage = &v
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentage() int64 {
	if o == nil || o.DownloadPercentage == nil {
		var ret int64
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentageOk() (*int64, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given int64 and assigns it to the DownloadPercentage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadPercentage(v int64) {
	o.DownloadPercentage = &v
}

// GetDownloadStage returns the DownloadStage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadStage() string {
	if o == nil || o.DownloadStage == nil {
		var ret string
		return ret
	}
	return *o.DownloadStage
}

// GetDownloadStageOk returns a tuple with the DownloadStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadStageOk() (*string, bool) {
	if o == nil || o.DownloadStage == nil {
		return nil, false
	}
	return o.DownloadStage, true
}

// HasDownloadStage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadStage() bool {
	if o != nil && o.DownloadStage != nil {
		return true
	}

	return false
}

// SetDownloadStage gets a reference to the given string and assigns it to the DownloadStage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadStage(v string) {
	o.DownloadStage = &v
}

// GetEpPowerStatus returns the EpPowerStatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatus() string {
	if o == nil || o.EpPowerStatus == nil {
		var ret string
		return ret
	}
	return *o.EpPowerStatus
}

// GetEpPowerStatusOk returns a tuple with the EpPowerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatusOk() (*string, bool) {
	if o == nil || o.EpPowerStatus == nil {
		return nil, false
	}
	return o.EpPowerStatus, true
}

// HasEpPowerStatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasEpPowerStatus() bool {
	if o != nil && o.EpPowerStatus != nil {
		return true
	}

	return false
}

// SetEpPowerStatus gets a reference to the given string and assigns it to the EpPowerStatus field.
func (o *FirmwareUpgradeStatusAllOf) SetEpPowerStatus(v string) {
	o.EpPowerStatus = &v
}

// GetOverallError returns the OverallError field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallError() string {
	if o == nil || o.OverallError == nil {
		var ret string
		return ret
	}
	return *o.OverallError
}

// GetOverallErrorOk returns a tuple with the OverallError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallErrorOk() (*string, bool) {
	if o == nil || o.OverallError == nil {
		return nil, false
	}
	return o.OverallError, true
}

// HasOverallError returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallError() bool {
	if o != nil && o.OverallError != nil {
		return true
	}

	return false
}

// SetOverallError gets a reference to the given string and assigns it to the OverallError field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallError(v string) {
	o.OverallError = &v
}

// GetOverallPercentage returns the OverallPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentage() int64 {
	if o == nil || o.OverallPercentage == nil {
		var ret int64
		return ret
	}
	return *o.OverallPercentage
}

// GetOverallPercentageOk returns a tuple with the OverallPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentageOk() (*int64, bool) {
	if o == nil || o.OverallPercentage == nil {
		return nil, false
	}
	return o.OverallPercentage, true
}

// HasOverallPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallPercentage() bool {
	if o != nil && o.OverallPercentage != nil {
		return true
	}

	return false
}

// SetOverallPercentage gets a reference to the given int64 and assigns it to the OverallPercentage field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallPercentage(v int64) {
	o.OverallPercentage = &v
}

// GetOverallstatus returns the Overallstatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallstatus() string {
	if o == nil || o.Overallstatus == nil {
		var ret string
		return ret
	}
	return *o.Overallstatus
}

// GetOverallstatusOk returns a tuple with the Overallstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallstatusOk() (*string, bool) {
	if o == nil || o.Overallstatus == nil {
		return nil, false
	}
	return o.Overallstatus, true
}

// HasOverallstatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallstatus() bool {
	if o != nil && o.Overallstatus != nil {
		return true
	}

	return false
}

// SetOverallstatus gets a reference to the given string and assigns it to the Overallstatus field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallstatus(v string) {
	o.Overallstatus = &v
}

// GetPendingType returns the PendingType field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetPendingType() string {
	if o == nil || o.PendingType == nil {
		var ret string
		return ret
	}
	return *o.PendingType
}

// GetPendingTypeOk returns a tuple with the PendingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetPendingTypeOk() (*string, bool) {
	if o == nil || o.PendingType == nil {
		return nil, false
	}
	return o.PendingType, true
}

// HasPendingType returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasPendingType() bool {
	if o != nil && o.PendingType != nil {
		return true
	}

	return false
}

// SetPendingType gets a reference to the given string and assigns it to the PendingType field.
func (o *FirmwareUpgradeStatusAllOf) SetPendingType(v string) {
	o.PendingType = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetUpgrade() FirmwareUpgradeBaseRelationship {
	if o == nil || o.Upgrade == nil {
		var ret FirmwareUpgradeBaseRelationship
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given FirmwareUpgradeBaseRelationship and assigns it to the Upgrade field.
func (o *FirmwareUpgradeStatusAllOf) SetUpgrade(v FirmwareUpgradeBaseRelationship) {
	o.Upgrade = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *FirmwareUpgradeStatusAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o FirmwareUpgradeStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DownloadMessage != nil {
		toSerialize["DownloadMessage"] = o.DownloadMessage
	}
	if o.DownloadPercentage != nil {
		toSerialize["DownloadPercentage"] = o.DownloadPercentage
	}
	if o.DownloadStage != nil {
		toSerialize["DownloadStage"] = o.DownloadStage
	}
	if o.EpPowerStatus != nil {
		toSerialize["EpPowerStatus"] = o.EpPowerStatus
	}
	if o.OverallError != nil {
		toSerialize["OverallError"] = o.OverallError
	}
	if o.OverallPercentage != nil {
		toSerialize["OverallPercentage"] = o.OverallPercentage
	}
	if o.Overallstatus != nil {
		toSerialize["Overallstatus"] = o.Overallstatus
	}
	if o.PendingType != nil {
		toSerialize["PendingType"] = o.PendingType
	}
	if o.Upgrade != nil {
		toSerialize["Upgrade"] = o.Upgrade
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareUpgradeStatusAllOf := _FirmwareUpgradeStatusAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareUpgradeStatusAllOf); err == nil {
		*o = FirmwareUpgradeStatusAllOf(varFirmwareUpgradeStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadMessage")
		delete(additionalProperties, "DownloadPercentage")
		delete(additionalProperties, "DownloadStage")
		delete(additionalProperties, "EpPowerStatus")
		delete(additionalProperties, "OverallError")
		delete(additionalProperties, "OverallPercentage")
		delete(additionalProperties, "Overallstatus")
		delete(additionalProperties, "PendingType")
		delete(additionalProperties, "Upgrade")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareUpgradeStatusAllOf struct {
	value *FirmwareUpgradeStatusAllOf
	isSet bool
}

func (v NullableFirmwareUpgradeStatusAllOf) Get() *FirmwareUpgradeStatusAllOf {
	return v.value
}

func (v *NullableFirmwareUpgradeStatusAllOf) Set(val *FirmwareUpgradeStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeStatusAllOf(val *FirmwareUpgradeStatusAllOf) *NullableFirmwareUpgradeStatusAllOf {
	return &NullableFirmwareUpgradeStatusAllOf{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
