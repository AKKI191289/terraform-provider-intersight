/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// OsInstall This MO models the target server, answers and other properties needed for OS installation. The OS installation can be started in the target server by doing a POST on this MO. The requests to this MO starts a OS installation workflow that can be tracked using workflow engine MO workflow.WorkflowInfo.
type OsInstall struct {
	OsBaseInstallConfig
	// The name of the OS install configuration.
	Name                 *string                                                      `json:"Name,omitempty"`
	ConfigurationFile    *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
	Image                *SoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
	Organization         *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
	OsduImage            *FirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
	Server               *ComputePhysicalRelationship                                 `json:"Server,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsInstall OsInstall

// NewOsInstall instantiates a new OsInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsInstall() *OsInstall {
	this := OsInstall{}
	return &this
}

// NewOsInstallWithDefaults instantiates a new OsInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsInstallWithDefaults() *OsInstall {
	this := OsInstall{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsInstall) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsInstall) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsInstall) SetName(v string) {
	o.Name = &v
}

// GetConfigurationFile returns the ConfigurationFile field value if set, zero value otherwise.
func (o *OsInstall) GetConfigurationFile() OsConfigurationFileRelationship {
	if o == nil || o.ConfigurationFile == nil {
		var ret OsConfigurationFileRelationship
		return ret
	}
	return *o.ConfigurationFile
}

// GetConfigurationFileOk returns a tuple with the ConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool) {
	if o == nil || o.ConfigurationFile == nil {
		return nil, false
	}
	return o.ConfigurationFile, true
}

// HasConfigurationFile returns a boolean if a field has been set.
func (o *OsInstall) HasConfigurationFile() bool {
	if o != nil && o.ConfigurationFile != nil {
		return true
	}

	return false
}

// SetConfigurationFile gets a reference to the given OsConfigurationFileRelationship and assigns it to the ConfigurationFile field.
func (o *OsInstall) SetConfigurationFile(v OsConfigurationFileRelationship) {
	o.ConfigurationFile = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OsInstall) GetImage() SoftwarerepositoryOperatingSystemFileRelationship {
	if o == nil || o.Image == nil {
		var ret SoftwarerepositoryOperatingSystemFileRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OsInstall) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given SoftwarerepositoryOperatingSystemFileRelationship and assigns it to the Image field.
func (o *OsInstall) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship) {
	o.Image = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OsInstall) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsInstall) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsInstall) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetOsduImage returns the OsduImage field value if set, zero value otherwise.
func (o *OsInstall) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship {
	if o == nil || o.OsduImage == nil {
		var ret FirmwareServerConfigurationUtilityDistributableRelationship
		return ret
	}
	return *o.OsduImage
}

// GetOsduImageOk returns a tuple with the OsduImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool) {
	if o == nil || o.OsduImage == nil {
		return nil, false
	}
	return o.OsduImage, true
}

// HasOsduImage returns a boolean if a field has been set.
func (o *OsInstall) HasOsduImage() bool {
	if o != nil && o.OsduImage != nil {
		return true
	}

	return false
}

// SetOsduImage gets a reference to the given FirmwareServerConfigurationUtilityDistributableRelationship and assigns it to the OsduImage field.
func (o *OsInstall) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship) {
	o.OsduImage = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *OsInstall) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *OsInstall) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *OsInstall) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *OsInstall) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *OsInstall) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *OsInstall) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o OsInstall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsBaseInstallConfig, errOsBaseInstallConfig := json.Marshal(o.OsBaseInstallConfig)
	if errOsBaseInstallConfig != nil {
		return []byte{}, errOsBaseInstallConfig
	}
	errOsBaseInstallConfig = json.Unmarshal([]byte(serializedOsBaseInstallConfig), &toSerialize)
	if errOsBaseInstallConfig != nil {
		return []byte{}, errOsBaseInstallConfig
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ConfigurationFile != nil {
		toSerialize["ConfigurationFile"] = o.ConfigurationFile
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OsduImage != nil {
		toSerialize["OsduImage"] = o.OsduImage
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsInstall) UnmarshalJSON(bytes []byte) (err error) {
	type OsInstallWithoutEmbeddedStruct struct {
		// The name of the OS install configuration.
		Name              *string                                                      `json:"Name,omitempty"`
		ConfigurationFile *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
		Image             *SoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
		Organization      *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
		OsduImage         *FirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
		Server            *ComputePhysicalRelationship                                 `json:"Server,omitempty"`
		WorkflowInfo      *WorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
	}

	varOsInstallWithoutEmbeddedStruct := OsInstallWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsInstallWithoutEmbeddedStruct)
	if err == nil {
		varOsInstall := _OsInstall{}
		varOsInstall.Name = varOsInstallWithoutEmbeddedStruct.Name
		varOsInstall.ConfigurationFile = varOsInstallWithoutEmbeddedStruct.ConfigurationFile
		varOsInstall.Image = varOsInstallWithoutEmbeddedStruct.Image
		varOsInstall.Organization = varOsInstallWithoutEmbeddedStruct.Organization
		varOsInstall.OsduImage = varOsInstallWithoutEmbeddedStruct.OsduImage
		varOsInstall.Server = varOsInstallWithoutEmbeddedStruct.Server
		varOsInstall.WorkflowInfo = varOsInstallWithoutEmbeddedStruct.WorkflowInfo
		*o = OsInstall(varOsInstall)
	} else {
		return err
	}

	varOsInstall := _OsInstall{}

	err = json.Unmarshal(bytes, &varOsInstall)
	if err == nil {
		o.OsBaseInstallConfig = varOsInstall.OsBaseInstallConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ConfigurationFile")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OsduImage")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "WorkflowInfo")

		// remove fields from embedded structs
		reflectOsBaseInstallConfig := reflect.ValueOf(o.OsBaseInstallConfig)
		for i := 0; i < reflectOsBaseInstallConfig.Type().NumField(); i++ {
			t := reflectOsBaseInstallConfig.Type().Field(i)

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

type NullableOsInstall struct {
	value *OsInstall
	isSet bool
}

func (v NullableOsInstall) Get() *OsInstall {
	return v.value
}

func (v *NullableOsInstall) Set(val *OsInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableOsInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableOsInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsInstall(val *OsInstall) *NullableOsInstall {
	return &NullableOsInstall{value: val, isSet: true}
}

func (v NullableOsInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
