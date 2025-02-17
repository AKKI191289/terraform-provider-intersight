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

// FirmwareDirectDownloadAllOf Definition of the list of properties defined in 'firmware.DirectDownload', excluding properties defined in parent classes.
type FirmwareDirectDownloadAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                     `json:"ObjectType"`
	HttpServer NullableFirmwareHttpServer `json:"HttpServer,omitempty"`
	// Source type referring the image to be downloaded from CCO or from a local HTTPS server. * `cisco` - Image to be downloaded from Cisco software repository. * `localHttp` - Image to be downloaded from a https server.
	ImageSource *string `json:"ImageSource,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Password as configured on the local https server.
	Password *string `json:"Password,omitempty"`
	// Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot. * `sd_upgrade_mount_only` - Direct upgrade SD upgrade mount only. * `sd_download_only` - Direct upgrade SD download only. * `sd_upgrade_only` - Direct upgrade SD upgrade only. * `sd_upgrade_full` - Direct upgrade SD upgrade full. * `upgrade_full` - The upgrade downloads or mounts the image, and reboots immediately for an upgrade. * `upgrade_mount_only` - The upgrade downloads or mounts the image. The upgrade happens in next reboot. * `chassis_upgrade_full` - Direct upgrade chassis upgrade full.
	Upgradeoption *string `json:"Upgradeoption,omitempty"`
	// Username as configured on the local https server.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDirectDownloadAllOf FirmwareDirectDownloadAllOf

// NewFirmwareDirectDownloadAllOf instantiates a new FirmwareDirectDownloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDirectDownloadAllOf(classId string, objectType string) *FirmwareDirectDownloadAllOf {
	this := FirmwareDirectDownloadAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var imageSource string = "cisco"
	this.ImageSource = &imageSource
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var upgradeoption string = "sd_upgrade_mount_only"
	this.Upgradeoption = &upgradeoption
	return &this
}

// NewFirmwareDirectDownloadAllOfWithDefaults instantiates a new FirmwareDirectDownloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDirectDownloadAllOfWithDefaults() *FirmwareDirectDownloadAllOf {
	this := FirmwareDirectDownloadAllOf{}
	var classId string = "firmware.DirectDownload"
	this.ClassId = classId
	var objectType string = "firmware.DirectDownload"
	this.ObjectType = objectType
	var imageSource string = "cisco"
	this.ImageSource = &imageSource
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var upgradeoption string = "sd_upgrade_mount_only"
	this.Upgradeoption = &upgradeoption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareDirectDownloadAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareDirectDownloadAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareDirectDownloadAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareDirectDownloadAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareDirectDownloadAllOf) GetHttpServer() FirmwareHttpServer {
	if o == nil || o.HttpServer.Get() == nil {
		var ret FirmwareHttpServer
		return ret
	}
	return *o.HttpServer.Get()
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareDirectDownloadAllOf) GetHttpServerOk() (*FirmwareHttpServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpServer.Get(), o.HttpServer.IsSet()
}

// HasHttpServer returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasHttpServer() bool {
	if o != nil && o.HttpServer.IsSet() {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given NullableFirmwareHttpServer and assigns it to the HttpServer field.
func (o *FirmwareDirectDownloadAllOf) SetHttpServer(v FirmwareHttpServer) {
	o.HttpServer.Set(&v)
}

// SetHttpServerNil sets the value for HttpServer to be an explicit nil
func (o *FirmwareDirectDownloadAllOf) SetHttpServerNil() {
	o.HttpServer.Set(nil)
}

// UnsetHttpServer ensures that no value is present for HttpServer, not even an explicit nil
func (o *FirmwareDirectDownloadAllOf) UnsetHttpServer() {
	o.HttpServer.Unset()
}

// GetImageSource returns the ImageSource field value if set, zero value otherwise.
func (o *FirmwareDirectDownloadAllOf) GetImageSource() string {
	if o == nil || o.ImageSource == nil {
		var ret string
		return ret
	}
	return *o.ImageSource
}

// GetImageSourceOk returns a tuple with the ImageSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetImageSourceOk() (*string, bool) {
	if o == nil || o.ImageSource == nil {
		return nil, false
	}
	return o.ImageSource, true
}

// HasImageSource returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasImageSource() bool {
	if o != nil && o.ImageSource != nil {
		return true
	}

	return false
}

// SetImageSource gets a reference to the given string and assigns it to the ImageSource field.
func (o *FirmwareDirectDownloadAllOf) SetImageSource(v string) {
	o.ImageSource = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *FirmwareDirectDownloadAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *FirmwareDirectDownloadAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FirmwareDirectDownloadAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FirmwareDirectDownloadAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUpgradeoption returns the Upgradeoption field value if set, zero value otherwise.
func (o *FirmwareDirectDownloadAllOf) GetUpgradeoption() string {
	if o == nil || o.Upgradeoption == nil {
		var ret string
		return ret
	}
	return *o.Upgradeoption
}

// GetUpgradeoptionOk returns a tuple with the Upgradeoption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetUpgradeoptionOk() (*string, bool) {
	if o == nil || o.Upgradeoption == nil {
		return nil, false
	}
	return o.Upgradeoption, true
}

// HasUpgradeoption returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasUpgradeoption() bool {
	if o != nil && o.Upgradeoption != nil {
		return true
	}

	return false
}

// SetUpgradeoption gets a reference to the given string and assigns it to the Upgradeoption field.
func (o *FirmwareDirectDownloadAllOf) SetUpgradeoption(v string) {
	o.Upgradeoption = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *FirmwareDirectDownloadAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDirectDownloadAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *FirmwareDirectDownloadAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *FirmwareDirectDownloadAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o FirmwareDirectDownloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HttpServer.IsSet() {
		toSerialize["HttpServer"] = o.HttpServer.Get()
	}
	if o.ImageSource != nil {
		toSerialize["ImageSource"] = o.ImageSource
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Upgradeoption != nil {
		toSerialize["Upgradeoption"] = o.Upgradeoption
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDirectDownloadAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareDirectDownloadAllOf := _FirmwareDirectDownloadAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareDirectDownloadAllOf); err == nil {
		*o = FirmwareDirectDownloadAllOf(varFirmwareDirectDownloadAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HttpServer")
		delete(additionalProperties, "ImageSource")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Upgradeoption")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareDirectDownloadAllOf struct {
	value *FirmwareDirectDownloadAllOf
	isSet bool
}

func (v NullableFirmwareDirectDownloadAllOf) Get() *FirmwareDirectDownloadAllOf {
	return v.value
}

func (v *NullableFirmwareDirectDownloadAllOf) Set(val *FirmwareDirectDownloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDirectDownloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDirectDownloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDirectDownloadAllOf(val *FirmwareDirectDownloadAllOf) *NullableFirmwareDirectDownloadAllOf {
	return &NullableFirmwareDirectDownloadAllOf{value: val, isSet: true}
}

func (v NullableFirmwareDirectDownloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDirectDownloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
