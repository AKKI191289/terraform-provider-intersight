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

// IaasServiceRequest Gets last six months Service Requests from UCSD.
type IaasServiceRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Service request duration.
	Duration *string `json:"Duration,omitempty"`
	// Service Request initiating user.
	InitiatingUser *string `json:"InitiatingUser,omitempty"`
	// Service request end time.
	RequestEndTime *string `json:"RequestEndTime,omitempty"`
	// Service request id of an SR.
	RequestId *string `json:"RequestId,omitempty"`
	// Service request start time.
	RequestStartTime *string `json:"RequestStartTime,omitempty"`
	// Service request type of an SR.
	RequestType *string `json:"RequestType,omitempty"`
	// UCSD service request status.
	Status *string `json:"Status,omitempty"`
	// Executed workflow name for an SR.
	WorkflowName         *string                              `json:"WorkflowName,omitempty"`
	WorkflowSteps        []IaasWorkflowSteps                  `json:"WorkflowSteps,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasServiceRequest IaasServiceRequest

// NewIaasServiceRequest instantiates a new IaasServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasServiceRequest(classId string, objectType string) *IaasServiceRequest {
	this := IaasServiceRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasServiceRequestWithDefaults instantiates a new IaasServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasServiceRequestWithDefaults() *IaasServiceRequest {
	this := IaasServiceRequest{}
	var classId string = "iaas.ServiceRequest"
	this.ClassId = classId
	var objectType string = "iaas.ServiceRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasServiceRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasServiceRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasServiceRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasServiceRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *IaasServiceRequest) SetDuration(v string) {
	o.Duration = &v
}

// GetInitiatingUser returns the InitiatingUser field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetInitiatingUser() string {
	if o == nil || o.InitiatingUser == nil {
		var ret string
		return ret
	}
	return *o.InitiatingUser
}

// GetInitiatingUserOk returns a tuple with the InitiatingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetInitiatingUserOk() (*string, bool) {
	if o == nil || o.InitiatingUser == nil {
		return nil, false
	}
	return o.InitiatingUser, true
}

// HasInitiatingUser returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasInitiatingUser() bool {
	if o != nil && o.InitiatingUser != nil {
		return true
	}

	return false
}

// SetInitiatingUser gets a reference to the given string and assigns it to the InitiatingUser field.
func (o *IaasServiceRequest) SetInitiatingUser(v string) {
	o.InitiatingUser = &v
}

// GetRequestEndTime returns the RequestEndTime field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestEndTime() string {
	if o == nil || o.RequestEndTime == nil {
		var ret string
		return ret
	}
	return *o.RequestEndTime
}

// GetRequestEndTimeOk returns a tuple with the RequestEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestEndTimeOk() (*string, bool) {
	if o == nil || o.RequestEndTime == nil {
		return nil, false
	}
	return o.RequestEndTime, true
}

// HasRequestEndTime returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestEndTime() bool {
	if o != nil && o.RequestEndTime != nil {
		return true
	}

	return false
}

// SetRequestEndTime gets a reference to the given string and assigns it to the RequestEndTime field.
func (o *IaasServiceRequest) SetRequestEndTime(v string) {
	o.RequestEndTime = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *IaasServiceRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestStartTime returns the RequestStartTime field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestStartTime() string {
	if o == nil || o.RequestStartTime == nil {
		var ret string
		return ret
	}
	return *o.RequestStartTime
}

// GetRequestStartTimeOk returns a tuple with the RequestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestStartTimeOk() (*string, bool) {
	if o == nil || o.RequestStartTime == nil {
		return nil, false
	}
	return o.RequestStartTime, true
}

// HasRequestStartTime returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestStartTime() bool {
	if o != nil && o.RequestStartTime != nil {
		return true
	}

	return false
}

// SetRequestStartTime gets a reference to the given string and assigns it to the RequestStartTime field.
func (o *IaasServiceRequest) SetRequestStartTime(v string) {
	o.RequestStartTime = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *IaasServiceRequest) SetRequestType(v string) {
	o.RequestType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasServiceRequest) SetStatus(v string) {
	o.Status = &v
}

// GetWorkflowName returns the WorkflowName field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetWorkflowName() string {
	if o == nil || o.WorkflowName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetWorkflowNameOk() (*string, bool) {
	if o == nil || o.WorkflowName == nil {
		return nil, false
	}
	return o.WorkflowName, true
}

// HasWorkflowName returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasWorkflowName() bool {
	if o != nil && o.WorkflowName != nil {
		return true
	}

	return false
}

// SetWorkflowName gets a reference to the given string and assigns it to the WorkflowName field.
func (o *IaasServiceRequest) SetWorkflowName(v string) {
	o.WorkflowName = &v
}

// GetWorkflowSteps returns the WorkflowSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasServiceRequest) GetWorkflowSteps() []IaasWorkflowSteps {
	if o == nil {
		var ret []IaasWorkflowSteps
		return ret
	}
	return o.WorkflowSteps
}

// GetWorkflowStepsOk returns a tuple with the WorkflowSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasServiceRequest) GetWorkflowStepsOk() (*[]IaasWorkflowSteps, bool) {
	if o == nil || o.WorkflowSteps == nil {
		return nil, false
	}
	return &o.WorkflowSteps, true
}

// HasWorkflowSteps returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasWorkflowSteps() bool {
	if o != nil && o.WorkflowSteps != nil {
		return true
	}

	return false
}

// SetWorkflowSteps gets a reference to the given []IaasWorkflowSteps and assigns it to the WorkflowSteps field.
func (o *IaasServiceRequest) SetWorkflowSteps(v []IaasWorkflowSteps) {
	o.WorkflowSteps = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasServiceRequest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o IaasServiceRequest) MarshalJSON() ([]byte, error) {
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
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.InitiatingUser != nil {
		toSerialize["InitiatingUser"] = o.InitiatingUser
	}
	if o.RequestEndTime != nil {
		toSerialize["RequestEndTime"] = o.RequestEndTime
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.RequestStartTime != nil {
		toSerialize["RequestStartTime"] = o.RequestStartTime
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.WorkflowName != nil {
		toSerialize["WorkflowName"] = o.WorkflowName
	}
	if o.WorkflowSteps != nil {
		toSerialize["WorkflowSteps"] = o.WorkflowSteps
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasServiceRequest) UnmarshalJSON(bytes []byte) (err error) {
	type IaasServiceRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Service request duration.
		Duration *string `json:"Duration,omitempty"`
		// Service Request initiating user.
		InitiatingUser *string `json:"InitiatingUser,omitempty"`
		// Service request end time.
		RequestEndTime *string `json:"RequestEndTime,omitempty"`
		// Service request id of an SR.
		RequestId *string `json:"RequestId,omitempty"`
		// Service request start time.
		RequestStartTime *string `json:"RequestStartTime,omitempty"`
		// Service request type of an SR.
		RequestType *string `json:"RequestType,omitempty"`
		// UCSD service request status.
		Status *string `json:"Status,omitempty"`
		// Executed workflow name for an SR.
		WorkflowName     *string                              `json:"WorkflowName,omitempty"`
		WorkflowSteps    []IaasWorkflowSteps                  `json:"WorkflowSteps,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varIaasServiceRequestWithoutEmbeddedStruct := IaasServiceRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasServiceRequestWithoutEmbeddedStruct)
	if err == nil {
		varIaasServiceRequest := _IaasServiceRequest{}
		varIaasServiceRequest.ClassId = varIaasServiceRequestWithoutEmbeddedStruct.ClassId
		varIaasServiceRequest.ObjectType = varIaasServiceRequestWithoutEmbeddedStruct.ObjectType
		varIaasServiceRequest.Duration = varIaasServiceRequestWithoutEmbeddedStruct.Duration
		varIaasServiceRequest.InitiatingUser = varIaasServiceRequestWithoutEmbeddedStruct.InitiatingUser
		varIaasServiceRequest.RequestEndTime = varIaasServiceRequestWithoutEmbeddedStruct.RequestEndTime
		varIaasServiceRequest.RequestId = varIaasServiceRequestWithoutEmbeddedStruct.RequestId
		varIaasServiceRequest.RequestStartTime = varIaasServiceRequestWithoutEmbeddedStruct.RequestStartTime
		varIaasServiceRequest.RequestType = varIaasServiceRequestWithoutEmbeddedStruct.RequestType
		varIaasServiceRequest.Status = varIaasServiceRequestWithoutEmbeddedStruct.Status
		varIaasServiceRequest.WorkflowName = varIaasServiceRequestWithoutEmbeddedStruct.WorkflowName
		varIaasServiceRequest.WorkflowSteps = varIaasServiceRequestWithoutEmbeddedStruct.WorkflowSteps
		varIaasServiceRequest.RegisteredDevice = varIaasServiceRequestWithoutEmbeddedStruct.RegisteredDevice
		*o = IaasServiceRequest(varIaasServiceRequest)
	} else {
		return err
	}

	varIaasServiceRequest := _IaasServiceRequest{}

	err = json.Unmarshal(bytes, &varIaasServiceRequest)
	if err == nil {
		o.MoBaseMo = varIaasServiceRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "InitiatingUser")
		delete(additionalProperties, "RequestEndTime")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "RequestStartTime")
		delete(additionalProperties, "RequestType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "WorkflowName")
		delete(additionalProperties, "WorkflowSteps")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableIaasServiceRequest struct {
	value *IaasServiceRequest
	isSet bool
}

func (v NullableIaasServiceRequest) Get() *IaasServiceRequest {
	return v.value
}

func (v *NullableIaasServiceRequest) Set(val *IaasServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasServiceRequest(val *IaasServiceRequest) *NullableIaasServiceRequest {
	return &NullableIaasServiceRequest{value: val, isSet: true}
}

func (v NullableIaasServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
