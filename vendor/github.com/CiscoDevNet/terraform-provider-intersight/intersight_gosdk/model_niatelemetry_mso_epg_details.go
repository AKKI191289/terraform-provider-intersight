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

// NiatelemetryMsoEpgDetails Details of Epgs configured from the Multi-Site Orchestrator.
type NiatelemetryMsoEpgDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Site Ids to which this EPG is deployed to.
	DeployedSites *string `json:"DeployedSites,omitempty"`
	// Name of EPG in Multi-Site Orchestrator.
	EpgName *string `json:"EpgName,omitempty"`
	// Is the EPG local to the Multi-Site Orchestrator.
	IsLocal *string `json:"IsLocal,omitempty"`
	// Unique reference for the EPG in the Multi-Site Orchestrator.
	Reference *string `json:"Reference,omitempty"`
	// Schema ID in Multi-Site Orchestrator.
	SchemaId *string `json:"SchemaId,omitempty"`
	// Schema name in Multi-Site Orchestrator.
	SchemaName *string `json:"SchemaName,omitempty"`
	// Template name in Multi-Site Orchestrator.
	TemplateName         *string                              `json:"TemplateName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoEpgDetails NiatelemetryMsoEpgDetails

// NewNiatelemetryMsoEpgDetails instantiates a new NiatelemetryMsoEpgDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoEpgDetails(classId string, objectType string) *NiatelemetryMsoEpgDetails {
	this := NiatelemetryMsoEpgDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoEpgDetailsWithDefaults instantiates a new NiatelemetryMsoEpgDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoEpgDetailsWithDefaults() *NiatelemetryMsoEpgDetails {
	this := NiatelemetryMsoEpgDetails{}
	var classId string = "niatelemetry.MsoEpgDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoEpgDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoEpgDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoEpgDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoEpgDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoEpgDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeployedSites returns the DeployedSites field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetDeployedSites() string {
	if o == nil || o.DeployedSites == nil {
		var ret string
		return ret
	}
	return *o.DeployedSites
}

// GetDeployedSitesOk returns a tuple with the DeployedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetDeployedSitesOk() (*string, bool) {
	if o == nil || o.DeployedSites == nil {
		return nil, false
	}
	return o.DeployedSites, true
}

// HasDeployedSites returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasDeployedSites() bool {
	if o != nil && o.DeployedSites != nil {
		return true
	}

	return false
}

// SetDeployedSites gets a reference to the given string and assigns it to the DeployedSites field.
func (o *NiatelemetryMsoEpgDetails) SetDeployedSites(v string) {
	o.DeployedSites = &v
}

// GetEpgName returns the EpgName field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetEpgName() string {
	if o == nil || o.EpgName == nil {
		var ret string
		return ret
	}
	return *o.EpgName
}

// GetEpgNameOk returns a tuple with the EpgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetEpgNameOk() (*string, bool) {
	if o == nil || o.EpgName == nil {
		return nil, false
	}
	return o.EpgName, true
}

// HasEpgName returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasEpgName() bool {
	if o != nil && o.EpgName != nil {
		return true
	}

	return false
}

// SetEpgName gets a reference to the given string and assigns it to the EpgName field.
func (o *NiatelemetryMsoEpgDetails) SetEpgName(v string) {
	o.EpgName = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetIsLocal() string {
	if o == nil || o.IsLocal == nil {
		var ret string
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetIsLocalOk() (*string, bool) {
	if o == nil || o.IsLocal == nil {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasIsLocal() bool {
	if o != nil && o.IsLocal != nil {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given string and assigns it to the IsLocal field.
func (o *NiatelemetryMsoEpgDetails) SetIsLocal(v string) {
	o.IsLocal = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *NiatelemetryMsoEpgDetails) SetReference(v string) {
	o.Reference = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *NiatelemetryMsoEpgDetails) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasSchemaName() bool {
	if o != nil && o.SchemaName != nil {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *NiatelemetryMsoEpgDetails) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *NiatelemetryMsoEpgDetails) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoEpgDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoEpgDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoEpgDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoEpgDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoEpgDetails) MarshalJSON() ([]byte, error) {
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
	if o.DeployedSites != nil {
		toSerialize["DeployedSites"] = o.DeployedSites
	}
	if o.EpgName != nil {
		toSerialize["EpgName"] = o.EpgName
	}
	if o.IsLocal != nil {
		toSerialize["IsLocal"] = o.IsLocal
	}
	if o.Reference != nil {
		toSerialize["Reference"] = o.Reference
	}
	if o.SchemaId != nil {
		toSerialize["SchemaId"] = o.SchemaId
	}
	if o.SchemaName != nil {
		toSerialize["SchemaName"] = o.SchemaName
	}
	if o.TemplateName != nil {
		toSerialize["TemplateName"] = o.TemplateName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMsoEpgDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryMsoEpgDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Site Ids to which this EPG is deployed to.
		DeployedSites *string `json:"DeployedSites,omitempty"`
		// Name of EPG in Multi-Site Orchestrator.
		EpgName *string `json:"EpgName,omitempty"`
		// Is the EPG local to the Multi-Site Orchestrator.
		IsLocal *string `json:"IsLocal,omitempty"`
		// Unique reference for the EPG in the Multi-Site Orchestrator.
		Reference *string `json:"Reference,omitempty"`
		// Schema ID in Multi-Site Orchestrator.
		SchemaId *string `json:"SchemaId,omitempty"`
		// Schema name in Multi-Site Orchestrator.
		SchemaName *string `json:"SchemaName,omitempty"`
		// Template name in Multi-Site Orchestrator.
		TemplateName     *string                              `json:"TemplateName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct := NiatelemetryMsoEpgDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMsoEpgDetails := _NiatelemetryMsoEpgDetails{}
		varNiatelemetryMsoEpgDetails.ClassId = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMsoEpgDetails.ObjectType = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMsoEpgDetails.DeployedSites = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.DeployedSites
		varNiatelemetryMsoEpgDetails.EpgName = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.EpgName
		varNiatelemetryMsoEpgDetails.IsLocal = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.IsLocal
		varNiatelemetryMsoEpgDetails.Reference = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.Reference
		varNiatelemetryMsoEpgDetails.SchemaId = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.SchemaId
		varNiatelemetryMsoEpgDetails.SchemaName = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.SchemaName
		varNiatelemetryMsoEpgDetails.TemplateName = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.TemplateName
		varNiatelemetryMsoEpgDetails.RegisteredDevice = varNiatelemetryMsoEpgDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMsoEpgDetails(varNiatelemetryMsoEpgDetails)
	} else {
		return err
	}

	varNiatelemetryMsoEpgDetails := _NiatelemetryMsoEpgDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoEpgDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMsoEpgDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeployedSites")
		delete(additionalProperties, "EpgName")
		delete(additionalProperties, "IsLocal")
		delete(additionalProperties, "Reference")
		delete(additionalProperties, "SchemaId")
		delete(additionalProperties, "SchemaName")
		delete(additionalProperties, "TemplateName")
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

type NullableNiatelemetryMsoEpgDetails struct {
	value *NiatelemetryMsoEpgDetails
	isSet bool
}

func (v NullableNiatelemetryMsoEpgDetails) Get() *NiatelemetryMsoEpgDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoEpgDetails) Set(val *NiatelemetryMsoEpgDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoEpgDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoEpgDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoEpgDetails(val *NiatelemetryMsoEpgDetails) *NullableNiatelemetryMsoEpgDetails {
	return &NullableNiatelemetryMsoEpgDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoEpgDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoEpgDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
