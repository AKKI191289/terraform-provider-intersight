# FirmwareUpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.UpgradeStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.UpgradeStatus"]
**DownloadMessage** | Pointer to **string** | The message from the endpoint during the download. | [optional] 
**DownloadPercentage** | Pointer to **int64** | The percentage of the image downloaded in the endpoint. | [optional] 
**DownloadStage** | Pointer to **string** | The image download stages. Example:downloading, flashing. | [optional] 
**EpPowerStatus** | Pointer to **string** | The server power status after the upgrade request is submitted in the endpoint. * &#x60;none&#x60; - Server power status is none. * &#x60;powered on&#x60; - Server power status is powered on. * &#x60;powered off&#x60; - Server power status is powered off. | [optional] [default to "none"]
**OverallError** | Pointer to **string** | The reason for the operation failure. | [optional] 
**OverallPercentage** | Pointer to **int64** | The overall percentage of the operation. | [optional] 
**Overallstatus** | Pointer to **string** | The overall status of the operation. * &#x60;none&#x60; - Upgrade stage is no upgrade stage. * &#x60;started&#x60; - Upgrade stage is started. * &#x60;prepare initiating&#x60; - Upgrade configuration is being prepared. * &#x60;prepare initiated&#x60; - Upgrade configuration is initiated. * &#x60;prepared&#x60; - Upgrade configuration is prepared. * &#x60;download initiating&#x60; - Upgrade stage is download initiating. * &#x60;download initiated&#x60; - Upgrade stage is download initiated. * &#x60;downloading&#x60; - Upgrade stage is downloading. * &#x60;downloaded&#x60; - Upgrade stage is downloaded. * &#x60;upgrade initiating on fabric A&#x60; - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * &#x60;upgrade initiated on fabric A&#x60; - Upgrade stage is in upgrade initiated when the upgrade has started in endpoint. * &#x60;upgrading fabric A&#x60; - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * &#x60;rebooting fabric A&#x60; - Upgrade is in rebooting when the endpoint is being rebooted. * &#x60;upgraded fabric A&#x60; - Upgrade stage is in upgraded when the corresponding endpoint has completed. * &#x60;upgrade initiating on fabric B&#x60; - Upgrade stage is in upgrade initiating when upgrade is being started in endopint. * &#x60;upgrade initiated on fabric B&#x60; - Upgrade stage is in upgrade initiated when upgrade has started in endpoint. * &#x60;upgrading fabric B&#x60; - Upgrade stage is in upgrading when the upgrade requires reboot to complete. * &#x60;rebooting fabric B&#x60; - Upgrade is in rebooting when the endpoint is being rebooted. * &#x60;upgraded fabric B&#x60; - Upgrade stage is in upgraded when the corresponding endpoint has completed. * &#x60;upgrade initiating&#x60; - Upgrade stage is upgrade initiating. * &#x60;upgrade initiated&#x60; - Upgrade stage is upgrade initiated. * &#x60;upgrading&#x60; - Upgrade stage is upgrading. * &#x60;oob images staging&#x60; - Out-of-band component images staging. * &#x60;oob images staged&#x60; - Out-of-band component images staged. * &#x60;rebooting&#x60; - Upgrade is rebooting the endpoint. * &#x60;upgraded&#x60; - Upgrade stage is upgraded. * &#x60;success&#x60; - Upgrade stage is success. * &#x60;failed&#x60; - Upgrade stage is upgrade failed. * &#x60;terminated&#x60; - Upgrade stage is terminated. * &#x60;pending&#x60; - Upgrade stage is pending. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - The image will be cached into Intersight Appliance or an endpoint cache. * &#x60;Cached&#x60; - The image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - The image caching into the Intersight Appliance failed or endpoint cache. | [optional] [default to "none"]
**PendingType** | Pointer to **string** | Pending reason for the upgrade waiting. * &#x60;none&#x60; - Upgrade pending reason is none. * &#x60;pending for next reboot&#x60; - Upgrade pending reason is pending for next reboot. | [optional] [default to "none"]
**Upgrade** | Pointer to [**FirmwareUpgradeBaseRelationship**](firmware.UpgradeBase.Relationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeStatus

`func NewFirmwareUpgradeStatus(classId string, objectType string, ) *FirmwareUpgradeStatus`

NewFirmwareUpgradeStatus instantiates a new FirmwareUpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeStatusWithDefaults

`func NewFirmwareUpgradeStatusWithDefaults() *FirmwareUpgradeStatus`

NewFirmwareUpgradeStatusWithDefaults instantiates a new FirmwareUpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDownloadMessage

`func (o *FirmwareUpgradeStatus) GetDownloadMessage() string`

GetDownloadMessage returns the DownloadMessage field if non-nil, zero value otherwise.

### GetDownloadMessageOk

`func (o *FirmwareUpgradeStatus) GetDownloadMessageOk() (*string, bool)`

GetDownloadMessageOk returns a tuple with the DownloadMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMessage

`func (o *FirmwareUpgradeStatus) SetDownloadMessage(v string)`

SetDownloadMessage sets DownloadMessage field to given value.

### HasDownloadMessage

`func (o *FirmwareUpgradeStatus) HasDownloadMessage() bool`

HasDownloadMessage returns a boolean if a field has been set.

### GetDownloadPercentage

`func (o *FirmwareUpgradeStatus) GetDownloadPercentage() int64`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *FirmwareUpgradeStatus) GetDownloadPercentageOk() (*int64, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *FirmwareUpgradeStatus) SetDownloadPercentage(v int64)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *FirmwareUpgradeStatus) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetDownloadStage

`func (o *FirmwareUpgradeStatus) GetDownloadStage() string`

GetDownloadStage returns the DownloadStage field if non-nil, zero value otherwise.

### GetDownloadStageOk

`func (o *FirmwareUpgradeStatus) GetDownloadStageOk() (*string, bool)`

GetDownloadStageOk returns a tuple with the DownloadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStage

`func (o *FirmwareUpgradeStatus) SetDownloadStage(v string)`

SetDownloadStage sets DownloadStage field to given value.

### HasDownloadStage

`func (o *FirmwareUpgradeStatus) HasDownloadStage() bool`

HasDownloadStage returns a boolean if a field has been set.

### GetEpPowerStatus

`func (o *FirmwareUpgradeStatus) GetEpPowerStatus() string`

GetEpPowerStatus returns the EpPowerStatus field if non-nil, zero value otherwise.

### GetEpPowerStatusOk

`func (o *FirmwareUpgradeStatus) GetEpPowerStatusOk() (*string, bool)`

GetEpPowerStatusOk returns a tuple with the EpPowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpPowerStatus

`func (o *FirmwareUpgradeStatus) SetEpPowerStatus(v string)`

SetEpPowerStatus sets EpPowerStatus field to given value.

### HasEpPowerStatus

`func (o *FirmwareUpgradeStatus) HasEpPowerStatus() bool`

HasEpPowerStatus returns a boolean if a field has been set.

### GetOverallError

`func (o *FirmwareUpgradeStatus) GetOverallError() string`

GetOverallError returns the OverallError field if non-nil, zero value otherwise.

### GetOverallErrorOk

`func (o *FirmwareUpgradeStatus) GetOverallErrorOk() (*string, bool)`

GetOverallErrorOk returns a tuple with the OverallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallError

`func (o *FirmwareUpgradeStatus) SetOverallError(v string)`

SetOverallError sets OverallError field to given value.

### HasOverallError

`func (o *FirmwareUpgradeStatus) HasOverallError() bool`

HasOverallError returns a boolean if a field has been set.

### GetOverallPercentage

`func (o *FirmwareUpgradeStatus) GetOverallPercentage() int64`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *FirmwareUpgradeStatus) GetOverallPercentageOk() (*int64, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *FirmwareUpgradeStatus) SetOverallPercentage(v int64)`

SetOverallPercentage sets OverallPercentage field to given value.

### HasOverallPercentage

`func (o *FirmwareUpgradeStatus) HasOverallPercentage() bool`

HasOverallPercentage returns a boolean if a field has been set.

### GetOverallstatus

`func (o *FirmwareUpgradeStatus) GetOverallstatus() string`

GetOverallstatus returns the Overallstatus field if non-nil, zero value otherwise.

### GetOverallstatusOk

`func (o *FirmwareUpgradeStatus) GetOverallstatusOk() (*string, bool)`

GetOverallstatusOk returns a tuple with the Overallstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallstatus

`func (o *FirmwareUpgradeStatus) SetOverallstatus(v string)`

SetOverallstatus sets Overallstatus field to given value.

### HasOverallstatus

`func (o *FirmwareUpgradeStatus) HasOverallstatus() bool`

HasOverallstatus returns a boolean if a field has been set.

### GetPendingType

`func (o *FirmwareUpgradeStatus) GetPendingType() string`

GetPendingType returns the PendingType field if non-nil, zero value otherwise.

### GetPendingTypeOk

`func (o *FirmwareUpgradeStatus) GetPendingTypeOk() (*string, bool)`

GetPendingTypeOk returns a tuple with the PendingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingType

`func (o *FirmwareUpgradeStatus) SetPendingType(v string)`

SetPendingType sets PendingType field to given value.

### HasPendingType

`func (o *FirmwareUpgradeStatus) HasPendingType() bool`

HasPendingType returns a boolean if a field has been set.

### GetUpgrade

`func (o *FirmwareUpgradeStatus) GetUpgrade() FirmwareUpgradeBaseRelationship`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *FirmwareUpgradeStatus) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *FirmwareUpgradeStatus) SetUpgrade(v FirmwareUpgradeBaseRelationship)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *FirmwareUpgradeStatus) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetWorkflow

`func (o *FirmwareUpgradeStatus) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *FirmwareUpgradeStatus) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *FirmwareUpgradeStatus) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *FirmwareUpgradeStatus) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


