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

// VmediaMappingAllOf Definition of the list of properties defined in 'vmedia.Mapping', excluding properties defined in parent classes.
type VmediaMappingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of Authentication protocol when CIFS is used for communication with the remote server. * `none` - No authentication is used. * `ntlm` - NT LAN Manager (NTLM) security protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * `ntlmi` - NTLMi security protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. * `ntlmv2` - NTLMv2 security protocol. Use this option only with Samba Linux. * `ntlmv2i` - NTLMv2i security protocol. Use this option only with Samba Linux. * `ntlmssp` - NT LAN Manager Security Support Provider (NTLMSSP) protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2. * `ntlmsspi` - NTLMSSPi protocol. Use this option only when you enable Digital Signing in the CIFS Windows server.
	AuthenticationProtocol *string `json:"AuthenticationProtocol,omitempty"`
	// Type of remote Virtual Media device. * `cdd` - Uses compact disc drive as the virtual media mount device. * `hdd` - Uses hard disk drive as the virtual media mount device.
	DeviceType *string `json:"DeviceType,omitempty"`
	// Remote location of image. Preferred format is 'hostname/filePath/fileName'.
	FileLocation *string `json:"FileLocation,omitempty"`
	// IP address or hostname of the remote server.
	HostName *string `json:"HostName,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\\n For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\\n For CIFS, supported options are soft, nounix, noserverino, guest.\\n For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\\n For HTTP/HTTPS, the only supported option is noauto.
	MountOptions *string `json:"MountOptions,omitempty"`
	// Protocol to use to communicate with the remote server. * `nfs` - NFS protocol for vmedia mount. * `cifs` - CIFS protocol for vmedia mount. * `http` - HTTP protocol for vmedia mount. * `https` - HTTPS protocol for vmedia mount.
	MountProtocol *string `json:"MountProtocol,omitempty"`
	// Password associated with the username.
	Password *string `json:"Password,omitempty"`
	// The remote file location path for the virtual media mapping. Accepted formats are: HDD for CIFS/NFS: hostname-or-IP/filePath/fileName.img. CDD for CIFS/NFS: hostname-or-IP/filePath/fileName.iso. HDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.img. CDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.iso.
	RemoteFile *string `json:"RemoteFile,omitempty"`
	// URL path to the location of the image on the remote server. The preferred format is '/path'.
	RemotePath *string `json:"RemotePath,omitempty"`
	// File Location in standard format 'hostname/filePath/fileName'. This field should be used to calculate config drift. User input format may vary while inventory will return data in format in compliance with mount option for the mount. Both will be converged to this standard format for comparison.
	SanitizedFileLocation *string `json:"SanitizedFileLocation,omitempty"`
	// Username to log in to the remote server.
	Username *string `json:"Username,omitempty"`
	// Identity of the image for Virtual Media mapping.
	VolumeName           *string `json:"VolumeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmediaMappingAllOf VmediaMappingAllOf

// NewVmediaMappingAllOf instantiates a new VmediaMappingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmediaMappingAllOf(classId string, objectType string) *VmediaMappingAllOf {
	this := VmediaMappingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var authenticationProtocol string = "none"
	this.AuthenticationProtocol = &authenticationProtocol
	var deviceType string = "cdd"
	this.DeviceType = &deviceType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var mountProtocol string = "nfs"
	this.MountProtocol = &mountProtocol
	return &this
}

// NewVmediaMappingAllOfWithDefaults instantiates a new VmediaMappingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmediaMappingAllOfWithDefaults() *VmediaMappingAllOf {
	this := VmediaMappingAllOf{}
	var classId string = "vmedia.Mapping"
	this.ClassId = classId
	var objectType string = "vmedia.Mapping"
	this.ObjectType = objectType
	var authenticationProtocol string = "none"
	this.AuthenticationProtocol = &authenticationProtocol
	var deviceType string = "cdd"
	this.DeviceType = &deviceType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var mountProtocol string = "nfs"
	this.MountProtocol = &mountProtocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *VmediaMappingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VmediaMappingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VmediaMappingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VmediaMappingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetAuthenticationProtocol() string {
	if o == nil || o.AuthenticationProtocol == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || o.AuthenticationProtocol == nil {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasAuthenticationProtocol() bool {
	if o != nil && o.AuthenticationProtocol != nil {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *VmediaMappingAllOf) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *VmediaMappingAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *VmediaMappingAllOf) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *VmediaMappingAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *VmediaMappingAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetMountOptions() string {
	if o == nil || o.MountOptions == nil {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetMountOptionsOk() (*string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasMountOptions() bool {
	if o != nil && o.MountOptions != nil {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *VmediaMappingAllOf) SetMountOptions(v string) {
	o.MountOptions = &v
}

// GetMountProtocol returns the MountProtocol field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetMountProtocol() string {
	if o == nil || o.MountProtocol == nil {
		var ret string
		return ret
	}
	return *o.MountProtocol
}

// GetMountProtocolOk returns a tuple with the MountProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetMountProtocolOk() (*string, bool) {
	if o == nil || o.MountProtocol == nil {
		return nil, false
	}
	return o.MountProtocol, true
}

// HasMountProtocol returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasMountProtocol() bool {
	if o != nil && o.MountProtocol != nil {
		return true
	}

	return false
}

// SetMountProtocol gets a reference to the given string and assigns it to the MountProtocol field.
func (o *VmediaMappingAllOf) SetMountProtocol(v string) {
	o.MountProtocol = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VmediaMappingAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetRemoteFile returns the RemoteFile field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetRemoteFile() string {
	if o == nil || o.RemoteFile == nil {
		var ret string
		return ret
	}
	return *o.RemoteFile
}

// GetRemoteFileOk returns a tuple with the RemoteFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetRemoteFileOk() (*string, bool) {
	if o == nil || o.RemoteFile == nil {
		return nil, false
	}
	return o.RemoteFile, true
}

// HasRemoteFile returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasRemoteFile() bool {
	if o != nil && o.RemoteFile != nil {
		return true
	}

	return false
}

// SetRemoteFile gets a reference to the given string and assigns it to the RemoteFile field.
func (o *VmediaMappingAllOf) SetRemoteFile(v string) {
	o.RemoteFile = &v
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetRemotePath() string {
	if o == nil || o.RemotePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePath
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetRemotePathOk() (*string, bool) {
	if o == nil || o.RemotePath == nil {
		return nil, false
	}
	return o.RemotePath, true
}

// HasRemotePath returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasRemotePath() bool {
	if o != nil && o.RemotePath != nil {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given string and assigns it to the RemotePath field.
func (o *VmediaMappingAllOf) SetRemotePath(v string) {
	o.RemotePath = &v
}

// GetSanitizedFileLocation returns the SanitizedFileLocation field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetSanitizedFileLocation() string {
	if o == nil || o.SanitizedFileLocation == nil {
		var ret string
		return ret
	}
	return *o.SanitizedFileLocation
}

// GetSanitizedFileLocationOk returns a tuple with the SanitizedFileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetSanitizedFileLocationOk() (*string, bool) {
	if o == nil || o.SanitizedFileLocation == nil {
		return nil, false
	}
	return o.SanitizedFileLocation, true
}

// HasSanitizedFileLocation returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasSanitizedFileLocation() bool {
	if o != nil && o.SanitizedFileLocation != nil {
		return true
	}

	return false
}

// SetSanitizedFileLocation gets a reference to the given string and assigns it to the SanitizedFileLocation field.
func (o *VmediaMappingAllOf) SetSanitizedFileLocation(v string) {
	o.SanitizedFileLocation = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *VmediaMappingAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *VmediaMappingAllOf) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaMappingAllOf) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *VmediaMappingAllOf) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *VmediaMappingAllOf) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o VmediaMappingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthenticationProtocol != nil {
		toSerialize["AuthenticationProtocol"] = o.AuthenticationProtocol
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.MountOptions != nil {
		toSerialize["MountOptions"] = o.MountOptions
	}
	if o.MountProtocol != nil {
		toSerialize["MountProtocol"] = o.MountProtocol
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RemoteFile != nil {
		toSerialize["RemoteFile"] = o.RemoteFile
	}
	if o.RemotePath != nil {
		toSerialize["RemotePath"] = o.RemotePath
	}
	if o.SanitizedFileLocation != nil {
		toSerialize["SanitizedFileLocation"] = o.SanitizedFileLocation
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.VolumeName != nil {
		toSerialize["VolumeName"] = o.VolumeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VmediaMappingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVmediaMappingAllOf := _VmediaMappingAllOf{}

	if err = json.Unmarshal(bytes, &varVmediaMappingAllOf); err == nil {
		*o = VmediaMappingAllOf(varVmediaMappingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthenticationProtocol")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "MountOptions")
		delete(additionalProperties, "MountProtocol")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "RemoteFile")
		delete(additionalProperties, "RemotePath")
		delete(additionalProperties, "SanitizedFileLocation")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "VolumeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVmediaMappingAllOf struct {
	value *VmediaMappingAllOf
	isSet bool
}

func (v NullableVmediaMappingAllOf) Get() *VmediaMappingAllOf {
	return v.value
}

func (v *NullableVmediaMappingAllOf) Set(val *VmediaMappingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVmediaMappingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVmediaMappingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmediaMappingAllOf(val *VmediaMappingAllOf) *NullableVmediaMappingAllOf {
	return &NullableVmediaMappingAllOf{value: val, isSet: true}
}

func (v NullableVmediaMappingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmediaMappingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
