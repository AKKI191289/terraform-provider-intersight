# VirtualizationVirtualDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualDisk"]
**Capacity** | Pointer to **string** | Disk capacity to be provided with units example - 10Gi. | [optional] 
**Discovered** | Pointer to **bool** | Flag to indicate whether the configuration is created from inventory object. | [optional] [readonly] 
**Mode** | Pointer to **string** | File mode of the disk  example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. | [optional] [default to "Block"]
**Name** | Pointer to **string** | Name of the storage disk. Name must be unique within a Datastore. | [optional] 
**SourceCerts** | Pointer to **string** | Base64 encoded CA certificates of the https source to check against. | [optional] 
**SourceDiskToClone** | Pointer to **string** | Source disk from which the content is copied. | [optional] 
**SourceFilePath** | Pointer to **string** | Image path used to import on the created disk. | [optional] 
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](virtualization.BaseCluster.Relationship.md) |  | [optional] 
**Inventory** | Pointer to [**VirtualizationBaseVirtualDiskRelationship**](virtualization.BaseVirtualDisk.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVirtualDiskAllOf

`func NewVirtualizationVirtualDiskAllOf(classId string, objectType string, ) *VirtualizationVirtualDiskAllOf`

NewVirtualizationVirtualDiskAllOf instantiates a new VirtualizationVirtualDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualDiskAllOfWithDefaults

`func NewVirtualizationVirtualDiskAllOfWithDefaults() *VirtualizationVirtualDiskAllOf`

NewVirtualizationVirtualDiskAllOfWithDefaults instantiates a new VirtualizationVirtualDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationVirtualDiskAllOf) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVirtualDiskAllOf) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVirtualDiskAllOf) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVirtualDiskAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualDiskAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualDiskAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualDiskAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualDiskAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVirtualDiskAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVirtualDiskAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVirtualDiskAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVirtualDiskAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) GetSourceCerts() string`

GetSourceCerts returns the SourceCerts field if non-nil, zero value otherwise.

### GetSourceCertsOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceCertsOk() (*string, bool)`

GetSourceCertsOk returns a tuple with the SourceCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) SetSourceCerts(v string)`

SetSourceCerts sets SourceCerts field to given value.

### HasSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) HasSourceCerts() bool`

HasSourceCerts returns a boolean if a field has been set.

### GetSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToClone() string`

GetSourceDiskToClone returns the SourceDiskToClone field if non-nil, zero value otherwise.

### GetSourceDiskToCloneOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToCloneOk() (*string, bool)`

GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) SetSourceDiskToClone(v string)`

SetSourceDiskToClone sets SourceDiskToClone field to given value.

### HasSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) HasSourceDiskToClone() bool`

HasSourceDiskToClone returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVirtualDiskAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualDiskAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualDiskAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualDiskAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualDiskAllOf) GetInventory() VirtualizationBaseVirtualDiskRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualDiskAllOf) GetInventoryOk() (*VirtualizationBaseVirtualDiskRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualDiskAllOf) SetInventory(v VirtualizationBaseVirtualDiskRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualDiskAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


