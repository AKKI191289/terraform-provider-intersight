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

// SoftwarerepositoryCifsServerAllOf Definition of the list of properties defined in 'softwarerepository.CifsServer', excluding properties defined in parent classes.
type SoftwarerepositoryCifsServerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
	FileLocation *string `json:"FileLocation,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// For CIFS, leave the field blank or enter one or more comma seperated options from the following. For Example, \" \" , \" soft \" , \" soft , nounix \" . * soft. * nounix. * noserviceino. * guest. * USERNAME=VALUE. * PASSWORD=VALUE. * sec=VALUE (VALUE could be None, Ntlm, Ntlmi, Ntlmssp, Ntlmsspi, Ntlmv2, Ntlmv2i).
	MountOption *string `json:"MountOption,omitempty"`
	// Password configured on the CIFS server.
	Password *string `json:"Password,omitempty"`
	// Filename of the image in the CIFS server. For example:ucs-c220m5-huu-3.1.2c.iso.
	RemoteFile *string `json:"RemoteFile,omitempty"`
	// Hostname or IP Address of the CIFS server.
	RemoteIp *string `json:"RemoteIp,omitempty"`
	// Remote directory where the image is present. For example:/share/subfolder.
	RemoteShare *string `json:"RemoteShare,omitempty"`
	// Username configured on the CIFS server.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCifsServerAllOf SoftwarerepositoryCifsServerAllOf

// NewSoftwarerepositoryCifsServerAllOf instantiates a new SoftwarerepositoryCifsServerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCifsServerAllOf(classId string, objectType string) *SoftwarerepositoryCifsServerAllOf {
	this := SoftwarerepositoryCifsServerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewSoftwarerepositoryCifsServerAllOfWithDefaults instantiates a new SoftwarerepositoryCifsServerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCifsServerAllOfWithDefaults() *SoftwarerepositoryCifsServerAllOf {
	this := SoftwarerepositoryCifsServerAllOf{}
	var classId string = "softwarerepository.CifsServer"
	this.ClassId = classId
	var objectType string = "softwarerepository.CifsServer"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCifsServerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCifsServerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCifsServerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCifsServerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *SoftwarerepositoryCifsServerAllOf) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryCifsServerAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMountOption returns the MountOption field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetMountOption() string {
	if o == nil || o.MountOption == nil {
		var ret string
		return ret
	}
	return *o.MountOption
}

// GetMountOptionOk returns a tuple with the MountOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetMountOptionOk() (*string, bool) {
	if o == nil || o.MountOption == nil {
		return nil, false
	}
	return o.MountOption, true
}

// HasMountOption returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasMountOption() bool {
	if o != nil && o.MountOption != nil {
		return true
	}

	return false
}

// SetMountOption gets a reference to the given string and assigns it to the MountOption field.
func (o *SoftwarerepositoryCifsServerAllOf) SetMountOption(v string) {
	o.MountOption = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryCifsServerAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetRemoteFile returns the RemoteFile field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteFile() string {
	if o == nil || o.RemoteFile == nil {
		var ret string
		return ret
	}
	return *o.RemoteFile
}

// GetRemoteFileOk returns a tuple with the RemoteFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteFileOk() (*string, bool) {
	if o == nil || o.RemoteFile == nil {
		return nil, false
	}
	return o.RemoteFile, true
}

// HasRemoteFile returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteFile() bool {
	if o != nil && o.RemoteFile != nil {
		return true
	}

	return false
}

// SetRemoteFile gets a reference to the given string and assigns it to the RemoteFile field.
func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteFile(v string) {
	o.RemoteFile = &v
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteIp() string {
	if o == nil || o.RemoteIp == nil {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteIpOk() (*string, bool) {
	if o == nil || o.RemoteIp == nil {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteIp() bool {
	if o != nil && o.RemoteIp != nil {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetRemoteShare returns the RemoteShare field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteShare() string {
	if o == nil || o.RemoteShare == nil {
		var ret string
		return ret
	}
	return *o.RemoteShare
}

// GetRemoteShareOk returns a tuple with the RemoteShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetRemoteShareOk() (*string, bool) {
	if o == nil || o.RemoteShare == nil {
		return nil, false
	}
	return o.RemoteShare, true
}

// HasRemoteShare returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasRemoteShare() bool {
	if o != nil && o.RemoteShare != nil {
		return true
	}

	return false
}

// SetRemoteShare gets a reference to the given string and assigns it to the RemoteShare field.
func (o *SoftwarerepositoryCifsServerAllOf) SetRemoteShare(v string) {
	o.RemoteShare = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServerAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServerAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServerAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SoftwarerepositoryCifsServerAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o SoftwarerepositoryCifsServerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.MountOption != nil {
		toSerialize["MountOption"] = o.MountOption
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RemoteFile != nil {
		toSerialize["RemoteFile"] = o.RemoteFile
	}
	if o.RemoteIp != nil {
		toSerialize["RemoteIp"] = o.RemoteIp
	}
	if o.RemoteShare != nil {
		toSerialize["RemoteShare"] = o.RemoteShare
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCifsServerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryCifsServerAllOf := _SoftwarerepositoryCifsServerAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryCifsServerAllOf); err == nil {
		*o = SoftwarerepositoryCifsServerAllOf(varSoftwarerepositoryCifsServerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "MountOption")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "RemoteFile")
		delete(additionalProperties, "RemoteIp")
		delete(additionalProperties, "RemoteShare")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryCifsServerAllOf struct {
	value *SoftwarerepositoryCifsServerAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCifsServerAllOf) Get() *SoftwarerepositoryCifsServerAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCifsServerAllOf) Set(val *SoftwarerepositoryCifsServerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCifsServerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCifsServerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCifsServerAllOf(val *SoftwarerepositoryCifsServerAllOf) *NullableSoftwarerepositoryCifsServerAllOf {
	return &NullableSoftwarerepositoryCifsServerAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCifsServerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCifsServerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
