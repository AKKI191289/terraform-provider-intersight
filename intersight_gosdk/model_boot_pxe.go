/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// BootPxe Device type used when booting from a PXE boot device.
type BootPxe struct {
	BootDeviceBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the underlying virtual ethernet interface used by the PXE boot device.
	InterfaceName *string `json:"InterfaceName,omitempty"`
	// Lists the supported Interface Source for PXE device. Supported values are \"name\" and \"mac\". * `name` - Use interface name to select virtual ethernet interface. * `mac` - Use MAC address to select virtual ethernet interface. * `port` - Use port to select virtual ethernet interface.
	InterfaceSource *string `json:"InterfaceSource,omitempty"`
	// The IP Address family type to use during the PXE Boot process. * `None` - Default value if IpType is not specified. * `IPv4` - The IPv4 address family type. * `IPv6` - The IPv6 address family type.
	IpType *string `json:"IpType,omitempty"`
	// The MAC Address of the underlying virtual ethernet interface used by the PXE boot device.
	MacAddress *string `json:"MacAddress,omitempty"`
	// The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are -1 to 255.
	Port *int64 `json:"Port,omitempty"`
	// The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \"MLOM\", \"L\", \"L1\", \"L2\", \"OCP\").
	Slot                 *string `json:"Slot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootPxe BootPxe

// NewBootPxe instantiates a new BootPxe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootPxe(classId string, objectType string) *BootPxe {
	this := BootPxe{}
	this.ClassId = classId
	this.ObjectType = objectType
	var interfaceSource string = "name"
	this.InterfaceSource = &interfaceSource
	var ipType string = "None"
	this.IpType = &ipType
	var port int64 = -1
	this.Port = &port
	return &this
}

// NewBootPxeWithDefaults instantiates a new BootPxe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootPxeWithDefaults() *BootPxe {
	this := BootPxe{}
	var classId string = "boot.Pxe"
	this.ClassId = classId
	var objectType string = "boot.Pxe"
	this.ObjectType = objectType
	var interfaceSource string = "name"
	this.InterfaceSource = &interfaceSource
	var ipType string = "None"
	this.IpType = &ipType
	var port int64 = -1
	this.Port = &port
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootPxe) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootPxe) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootPxe) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootPxe) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootPxe) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootPxe) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *BootPxe) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *BootPxe) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *BootPxe) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetInterfaceSource returns the InterfaceSource field value if set, zero value otherwise.
func (o *BootPxe) GetInterfaceSource() string {
	if o == nil || o.InterfaceSource == nil {
		var ret string
		return ret
	}
	return *o.InterfaceSource
}

// GetInterfaceSourceOk returns a tuple with the InterfaceSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetInterfaceSourceOk() (*string, bool) {
	if o == nil || o.InterfaceSource == nil {
		return nil, false
	}
	return o.InterfaceSource, true
}

// HasInterfaceSource returns a boolean if a field has been set.
func (o *BootPxe) HasInterfaceSource() bool {
	if o != nil && o.InterfaceSource != nil {
		return true
	}

	return false
}

// SetInterfaceSource gets a reference to the given string and assigns it to the InterfaceSource field.
func (o *BootPxe) SetInterfaceSource(v string) {
	o.InterfaceSource = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *BootPxe) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *BootPxe) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *BootPxe) SetIpType(v string) {
	o.IpType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *BootPxe) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *BootPxe) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *BootPxe) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *BootPxe) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *BootPxe) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *BootPxe) SetPort(v int64) {
	o.Port = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootPxe) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPxe) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootPxe) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootPxe) SetSlot(v string) {
	o.Slot = &v
}

func (o BootPxe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.InterfaceSource != nil {
		toSerialize["InterfaceSource"] = o.InterfaceSource
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootPxe) UnmarshalJSON(bytes []byte) (err error) {
	type BootPxeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the underlying virtual ethernet interface used by the PXE boot device.
		InterfaceName *string `json:"InterfaceName,omitempty"`
		// Lists the supported Interface Source for PXE device. Supported values are \"name\" and \"mac\". * `name` - Use interface name to select virtual ethernet interface. * `mac` - Use MAC address to select virtual ethernet interface. * `port` - Use port to select virtual ethernet interface.
		InterfaceSource *string `json:"InterfaceSource,omitempty"`
		// The IP Address family type to use during the PXE Boot process. * `None` - Default value if IpType is not specified. * `IPv4` - The IPv4 address family type. * `IPv6` - The IPv6 address family type.
		IpType *string `json:"IpType,omitempty"`
		// The MAC Address of the underlying virtual ethernet interface used by the PXE boot device.
		MacAddress *string `json:"MacAddress,omitempty"`
		// The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are -1 to 255.
		Port *int64 `json:"Port,omitempty"`
		// The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \"MLOM\", \"L\", \"L1\", \"L2\", \"OCP\").
		Slot *string `json:"Slot,omitempty"`
	}

	varBootPxeWithoutEmbeddedStruct := BootPxeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootPxeWithoutEmbeddedStruct)
	if err == nil {
		varBootPxe := _BootPxe{}
		varBootPxe.ClassId = varBootPxeWithoutEmbeddedStruct.ClassId
		varBootPxe.ObjectType = varBootPxeWithoutEmbeddedStruct.ObjectType
		varBootPxe.InterfaceName = varBootPxeWithoutEmbeddedStruct.InterfaceName
		varBootPxe.InterfaceSource = varBootPxeWithoutEmbeddedStruct.InterfaceSource
		varBootPxe.IpType = varBootPxeWithoutEmbeddedStruct.IpType
		varBootPxe.MacAddress = varBootPxeWithoutEmbeddedStruct.MacAddress
		varBootPxe.Port = varBootPxeWithoutEmbeddedStruct.Port
		varBootPxe.Slot = varBootPxeWithoutEmbeddedStruct.Slot
		*o = BootPxe(varBootPxe)
	} else {
		return err
	}

	varBootPxe := _BootPxe{}

	err = json.Unmarshal(bytes, &varBootPxe)
	if err == nil {
		o.BootDeviceBase = varBootPxe.BootDeviceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "InterfaceSource")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Slot")

		// remove fields from embedded structs
		reflectBootDeviceBase := reflect.ValueOf(o.BootDeviceBase)
		for i := 0; i < reflectBootDeviceBase.Type().NumField(); i++ {
			t := reflectBootDeviceBase.Type().Field(i)

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

type NullableBootPxe struct {
	value *BootPxe
	isSet bool
}

func (v NullableBootPxe) Get() *BootPxe {
	return v.value
}

func (v *NullableBootPxe) Set(val *BootPxe) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPxe) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPxe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPxe(val *BootPxe) *NullableBootPxe {
	return &NullableBootPxe{value: val, isSet: true}
}

func (v NullableBootPxe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPxe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
