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
)

// AssetTargetAllOf Definition of the list of properties defined in 'asset.Target', excluding properties defined in parent classes.
type AssetTargetAllOf struct {
	Connections *[]AssetConnection `json:"Connections,omitempty"`
	Services    *[]AssetService    `json:"Services,omitempty"`
	// Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. * `` - The device registered with Intersight but subsequently did not establish a persistent websocket connection. * `Connected` - The device's connection to Intersight has been established and is active. * `NotConnected` - The device's connection to Intersight has been disconnected. * `ClaimInProgress` - Claim of the device is in progress. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device.
	Status *string `json:"Status,omitempty"`
	// StatusErrorReason provides additional context for the Status.
	StatusErrorReason *string `json:"StatusErrorReason,omitempty"`
	// The type of the managed target. For example a UCS Server or Vmware Vcenter target. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - Intersight on-premise appliance. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NetAppOntap` - A NetApp ONTAP storage system. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A Dynatrace controller that monitors applications. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `MicrosoftAzure` - A Microsoft Azure target. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `IMCBlade` - An Intersight managed UCS Blade Server.
	TargetType           *string                              `json:"TargetType,omitempty"`
	Account              *IamAccountRelationship              `json:"Account,omitempty"`
	Assist               *AssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetAllOf AssetTargetAllOf

// NewAssetTargetAllOf instantiates a new AssetTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetAllOf() *AssetTargetAllOf {
	this := AssetTargetAllOf{}
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewAssetTargetAllOfWithDefaults instantiates a new AssetTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetAllOfWithDefaults() *AssetTargetAllOf {
	this := AssetTargetAllOf{}
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetConnections() []AssetConnection {
	if o == nil || o.Connections == nil {
		var ret []AssetConnection
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetConnectionsOk() (*[]AssetConnection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []AssetConnection and assigns it to the Connections field.
func (o *AssetTargetAllOf) SetConnections(v []AssetConnection) {
	o.Connections = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetServices() []AssetService {
	if o == nil || o.Services == nil {
		var ret []AssetService
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetServicesOk() (*[]AssetService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []AssetService and assigns it to the Services field.
func (o *AssetTargetAllOf) SetServices(v []AssetService) {
	o.Services = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetTargetAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusErrorReason returns the StatusErrorReason field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetStatusErrorReason() string {
	if o == nil || o.StatusErrorReason == nil {
		var ret string
		return ret
	}
	return *o.StatusErrorReason
}

// GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetStatusErrorReasonOk() (*string, bool) {
	if o == nil || o.StatusErrorReason == nil {
		return nil, false
	}
	return o.StatusErrorReason, true
}

// HasStatusErrorReason returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasStatusErrorReason() bool {
	if o != nil && o.StatusErrorReason != nil {
		return true
	}

	return false
}

// SetStatusErrorReason gets a reference to the given string and assigns it to the StatusErrorReason field.
func (o *AssetTargetAllOf) SetStatusErrorReason(v string) {
	o.StatusErrorReason = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AssetTargetAllOf) SetTargetType(v string) {
	o.TargetType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetTargetAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || o.Assist == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Assist == nil {
		return nil, false
	}
	return o.Assist, true
}

// HasAssist returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasAssist() bool {
	if o != nil && o.Assist != nil {
		return true
	}

	return false
}

// SetAssist gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *AssetTargetAllOf) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist = &v
}

func (o AssetTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connections != nil {
		toSerialize["Connections"] = o.Connections
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusErrorReason != nil {
		toSerialize["StatusErrorReason"] = o.StatusErrorReason
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Assist != nil {
		toSerialize["Assist"] = o.Assist
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTargetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTargetAllOf := _AssetTargetAllOf{}

	if err = json.Unmarshal(bytes, &varAssetTargetAllOf); err == nil {
		*o = AssetTargetAllOf(varAssetTargetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Connections")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusErrorReason")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Assist")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTargetAllOf struct {
	value *AssetTargetAllOf
	isSet bool
}

func (v NullableAssetTargetAllOf) Get() *AssetTargetAllOf {
	return v.value
}

func (v *NullableAssetTargetAllOf) Set(val *AssetTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetAllOf(val *AssetTargetAllOf) *NullableAssetTargetAllOf {
	return &NullableAssetTargetAllOf{value: val, isSet: true}
}

func (v NullableAssetTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
