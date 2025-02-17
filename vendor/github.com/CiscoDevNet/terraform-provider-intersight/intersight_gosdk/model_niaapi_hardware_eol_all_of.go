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
	"time"
)

// NiaapiHardwareEolAllOf Definition of the list of properties defined in 'niaapi.HardwareEol', excluding properties defined in parent classes.
type NiaapiHardwareEolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// String contains the PID of hardwares affected by this notice, seperated by comma.
	AffectedPids *string `json:"AffectedPids,omitempty"`
	// When this notice is announced.
	AnnouncementDate *time.Time `json:"AnnouncementDate,omitempty"`
	// Epoch time of Announcement Date.
	AnnouncementDateEpoch *int64 `json:"AnnouncementDateEpoch,omitempty"`
	// The bulletinno of this hardware end of life notice.
	BulletinNo *string `json:"BulletinNo,omitempty"`
	// The description of this hardware end of life notice.
	Description *string `json:"Description,omitempty"`
	// Date time of end of new services attachment  .
	EndofNewServiceAttachmentDate *time.Time `json:"EndofNewServiceAttachmentDate,omitempty"`
	// Epoch time of New service attachment Date .
	EndofNewServiceAttachmentDateEpoch *int64 `json:"EndofNewServiceAttachmentDateEpoch,omitempty"`
	// Date time of end of routinefailure analysis.
	EndofRoutineFailureAnalysisDate *time.Time `json:"EndofRoutineFailureAnalysisDate,omitempty"`
	// Epoch time of End of Routine Failure Analysis Date.
	EndofRoutineFailureAnalysisDateEpoch *int64 `json:"EndofRoutineFailureAnalysisDateEpoch,omitempty"`
	// When this hardware will end sale.
	EndofSaleDate *time.Time `json:"EndofSaleDate,omitempty"`
	// Epoch time of End of Sale Date.
	EndofSaleDateEpoch *int64 `json:"EndofSaleDateEpoch,omitempty"`
	// Date time of end of security support .
	EndofSecuritySupport *time.Time `json:"EndofSecuritySupport,omitempty"`
	// Epoch time of End of Security Support Date .
	EndofSecuritySupportEpoch *int64 `json:"EndofSecuritySupportEpoch,omitempty"`
	// Date time of end of service contract renew .
	EndofServiceContractRenewalDate *time.Time `json:"EndofServiceContractRenewalDate,omitempty"`
	// Epoch time of End of Renewal service contract.
	EndofServiceContractRenewalDateEpoch *int64 `json:"EndofServiceContractRenewalDateEpoch,omitempty"`
	// The date of end of maintainance.
	EndofSwMaintenanceDate *time.Time `json:"EndofSwMaintenanceDate,omitempty"`
	// Epoch time of End of maintenance Date.
	EndofSwMaintenanceDateEpoch *int64 `json:"EndofSwMaintenanceDateEpoch,omitempty"`
	// Hardware end of sale URL link to the notice webpage.
	HardwareEolUrl *string `json:"HardwareEolUrl,omitempty"`
	// The title of this hardware end of life notice.
	Headline *string `json:"Headline,omitempty"`
	// Date time of end of support .
	LastDateofSupport *time.Time `json:"LastDateofSupport,omitempty"`
	// Epoch time of last date of support .
	LastDateofSupportEpoch *int64 `json:"LastDateofSupportEpoch,omitempty"`
	// Date time of Lastship Date.
	LastShipDate *time.Time `json:"LastShipDate,omitempty"`
	// Epoch time of last ship Date.
	LastShipDateEpoch *int64 `json:"LastShipDateEpoch,omitempty"`
	// The URL of this migration notice.
	MigrationUrl         *string `json:"MigrationUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiHardwareEolAllOf NiaapiHardwareEolAllOf

// NewNiaapiHardwareEolAllOf instantiates a new NiaapiHardwareEolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiHardwareEolAllOf(classId string, objectType string) *NiaapiHardwareEolAllOf {
	this := NiaapiHardwareEolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiHardwareEolAllOfWithDefaults instantiates a new NiaapiHardwareEolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiHardwareEolAllOfWithDefaults() *NiaapiHardwareEolAllOf {
	this := NiaapiHardwareEolAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiHardwareEolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiHardwareEolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiHardwareEolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiHardwareEolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAffectedPids returns the AffectedPids field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetAffectedPids() string {
	if o == nil || o.AffectedPids == nil {
		var ret string
		return ret
	}
	return *o.AffectedPids
}

// GetAffectedPidsOk returns a tuple with the AffectedPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetAffectedPidsOk() (*string, bool) {
	if o == nil || o.AffectedPids == nil {
		return nil, false
	}
	return o.AffectedPids, true
}

// HasAffectedPids returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasAffectedPids() bool {
	if o != nil && o.AffectedPids != nil {
		return true
	}

	return false
}

// SetAffectedPids gets a reference to the given string and assigns it to the AffectedPids field.
func (o *NiaapiHardwareEolAllOf) SetAffectedPids(v string) {
	o.AffectedPids = &v
}

// GetAnnouncementDate returns the AnnouncementDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetAnnouncementDate() time.Time {
	if o == nil || o.AnnouncementDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AnnouncementDate
}

// GetAnnouncementDateOk returns a tuple with the AnnouncementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateOk() (*time.Time, bool) {
	if o == nil || o.AnnouncementDate == nil {
		return nil, false
	}
	return o.AnnouncementDate, true
}

// HasAnnouncementDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasAnnouncementDate() bool {
	if o != nil && o.AnnouncementDate != nil {
		return true
	}

	return false
}

// SetAnnouncementDate gets a reference to the given time.Time and assigns it to the AnnouncementDate field.
func (o *NiaapiHardwareEolAllOf) SetAnnouncementDate(v time.Time) {
	o.AnnouncementDate = &v
}

// GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateEpoch() int64 {
	if o == nil || o.AnnouncementDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.AnnouncementDateEpoch
}

// GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateEpochOk() (*int64, bool) {
	if o == nil || o.AnnouncementDateEpoch == nil {
		return nil, false
	}
	return o.AnnouncementDateEpoch, true
}

// HasAnnouncementDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasAnnouncementDateEpoch() bool {
	if o != nil && o.AnnouncementDateEpoch != nil {
		return true
	}

	return false
}

// SetAnnouncementDateEpoch gets a reference to the given int64 and assigns it to the AnnouncementDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetAnnouncementDateEpoch(v int64) {
	o.AnnouncementDateEpoch = &v
}

// GetBulletinNo returns the BulletinNo field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetBulletinNo() string {
	if o == nil || o.BulletinNo == nil {
		var ret string
		return ret
	}
	return *o.BulletinNo
}

// GetBulletinNoOk returns a tuple with the BulletinNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetBulletinNoOk() (*string, bool) {
	if o == nil || o.BulletinNo == nil {
		return nil, false
	}
	return o.BulletinNo, true
}

// HasBulletinNo returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasBulletinNo() bool {
	if o != nil && o.BulletinNo != nil {
		return true
	}

	return false
}

// SetBulletinNo gets a reference to the given string and assigns it to the BulletinNo field.
func (o *NiaapiHardwareEolAllOf) SetBulletinNo(v string) {
	o.BulletinNo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiaapiHardwareEolAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDate() time.Time {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofNewServiceAttachmentDate
}

// GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool) {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDate, true
}

// HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofNewServiceAttachmentDate() bool {
	if o != nil && o.EndofNewServiceAttachmentDate != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDate gets a reference to the given time.Time and assigns it to the EndofNewServiceAttachmentDate field.
func (o *NiaapiHardwareEolAllOf) SetEndofNewServiceAttachmentDate(v time.Time) {
	o.EndofNewServiceAttachmentDate = &v
}

// GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateEpoch() int64 {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofNewServiceAttachmentDateEpoch
}

// GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDateEpoch, true
}

// HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofNewServiceAttachmentDateEpoch() bool {
	if o != nil && o.EndofNewServiceAttachmentDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDateEpoch gets a reference to the given int64 and assigns it to the EndofNewServiceAttachmentDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofNewServiceAttachmentDateEpoch(v int64) {
	o.EndofNewServiceAttachmentDateEpoch = &v
}

// GetEndofRoutineFailureAnalysisDate returns the EndofRoutineFailureAnalysisDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDate() time.Time {
	if o == nil || o.EndofRoutineFailureAnalysisDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofRoutineFailureAnalysisDate
}

// GetEndofRoutineFailureAnalysisDateOk returns a tuple with the EndofRoutineFailureAnalysisDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateOk() (*time.Time, bool) {
	if o == nil || o.EndofRoutineFailureAnalysisDate == nil {
		return nil, false
	}
	return o.EndofRoutineFailureAnalysisDate, true
}

// HasEndofRoutineFailureAnalysisDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofRoutineFailureAnalysisDate() bool {
	if o != nil && o.EndofRoutineFailureAnalysisDate != nil {
		return true
	}

	return false
}

// SetEndofRoutineFailureAnalysisDate gets a reference to the given time.Time and assigns it to the EndofRoutineFailureAnalysisDate field.
func (o *NiaapiHardwareEolAllOf) SetEndofRoutineFailureAnalysisDate(v time.Time) {
	o.EndofRoutineFailureAnalysisDate = &v
}

// GetEndofRoutineFailureAnalysisDateEpoch returns the EndofRoutineFailureAnalysisDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateEpoch() int64 {
	if o == nil || o.EndofRoutineFailureAnalysisDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofRoutineFailureAnalysisDateEpoch
}

// GetEndofRoutineFailureAnalysisDateEpochOk returns a tuple with the EndofRoutineFailureAnalysisDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofRoutineFailureAnalysisDateEpoch == nil {
		return nil, false
	}
	return o.EndofRoutineFailureAnalysisDateEpoch, true
}

// HasEndofRoutineFailureAnalysisDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofRoutineFailureAnalysisDateEpoch() bool {
	if o != nil && o.EndofRoutineFailureAnalysisDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofRoutineFailureAnalysisDateEpoch gets a reference to the given int64 and assigns it to the EndofRoutineFailureAnalysisDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofRoutineFailureAnalysisDateEpoch(v int64) {
	o.EndofRoutineFailureAnalysisDateEpoch = &v
}

// GetEndofSaleDate returns the EndofSaleDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSaleDate() time.Time {
	if o == nil || o.EndofSaleDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSaleDate
}

// GetEndofSaleDateOk returns a tuple with the EndofSaleDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateOk() (*time.Time, bool) {
	if o == nil || o.EndofSaleDate == nil {
		return nil, false
	}
	return o.EndofSaleDate, true
}

// HasEndofSaleDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSaleDate() bool {
	if o != nil && o.EndofSaleDate != nil {
		return true
	}

	return false
}

// SetEndofSaleDate gets a reference to the given time.Time and assigns it to the EndofSaleDate field.
func (o *NiaapiHardwareEolAllOf) SetEndofSaleDate(v time.Time) {
	o.EndofSaleDate = &v
}

// GetEndofSaleDateEpoch returns the EndofSaleDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateEpoch() int64 {
	if o == nil || o.EndofSaleDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSaleDateEpoch
}

// GetEndofSaleDateEpochOk returns a tuple with the EndofSaleDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofSaleDateEpoch == nil {
		return nil, false
	}
	return o.EndofSaleDateEpoch, true
}

// HasEndofSaleDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSaleDateEpoch() bool {
	if o != nil && o.EndofSaleDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofSaleDateEpoch gets a reference to the given int64 and assigns it to the EndofSaleDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofSaleDateEpoch(v int64) {
	o.EndofSaleDateEpoch = &v
}

// GetEndofSecuritySupport returns the EndofSecuritySupport field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupport() time.Time {
	if o == nil || o.EndofSecuritySupport == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSecuritySupport
}

// GetEndofSecuritySupportOk returns a tuple with the EndofSecuritySupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportOk() (*time.Time, bool) {
	if o == nil || o.EndofSecuritySupport == nil {
		return nil, false
	}
	return o.EndofSecuritySupport, true
}

// HasEndofSecuritySupport returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSecuritySupport() bool {
	if o != nil && o.EndofSecuritySupport != nil {
		return true
	}

	return false
}

// SetEndofSecuritySupport gets a reference to the given time.Time and assigns it to the EndofSecuritySupport field.
func (o *NiaapiHardwareEolAllOf) SetEndofSecuritySupport(v time.Time) {
	o.EndofSecuritySupport = &v
}

// GetEndofSecuritySupportEpoch returns the EndofSecuritySupportEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportEpoch() int64 {
	if o == nil || o.EndofSecuritySupportEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSecuritySupportEpoch
}

// GetEndofSecuritySupportEpochOk returns a tuple with the EndofSecuritySupportEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportEpochOk() (*int64, bool) {
	if o == nil || o.EndofSecuritySupportEpoch == nil {
		return nil, false
	}
	return o.EndofSecuritySupportEpoch, true
}

// HasEndofSecuritySupportEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSecuritySupportEpoch() bool {
	if o != nil && o.EndofSecuritySupportEpoch != nil {
		return true
	}

	return false
}

// SetEndofSecuritySupportEpoch gets a reference to the given int64 and assigns it to the EndofSecuritySupportEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofSecuritySupportEpoch(v int64) {
	o.EndofSecuritySupportEpoch = &v
}

// GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDate() time.Time {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofServiceContractRenewalDate
}

// GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateOk() (*time.Time, bool) {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDate, true
}

// HasEndofServiceContractRenewalDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofServiceContractRenewalDate() bool {
	if o != nil && o.EndofServiceContractRenewalDate != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDate gets a reference to the given time.Time and assigns it to the EndofServiceContractRenewalDate field.
func (o *NiaapiHardwareEolAllOf) SetEndofServiceContractRenewalDate(v time.Time) {
	o.EndofServiceContractRenewalDate = &v
}

// GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateEpoch() int64 {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofServiceContractRenewalDateEpoch
}

// GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDateEpoch, true
}

// HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofServiceContractRenewalDateEpoch() bool {
	if o != nil && o.EndofServiceContractRenewalDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDateEpoch gets a reference to the given int64 and assigns it to the EndofServiceContractRenewalDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofServiceContractRenewalDateEpoch(v int64) {
	o.EndofServiceContractRenewalDateEpoch = &v
}

// GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDate() time.Time {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSwMaintenanceDate
}

// GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateOk() (*time.Time, bool) {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDate, true
}

// HasEndofSwMaintenanceDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSwMaintenanceDate() bool {
	if o != nil && o.EndofSwMaintenanceDate != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDate gets a reference to the given time.Time and assigns it to the EndofSwMaintenanceDate field.
func (o *NiaapiHardwareEolAllOf) SetEndofSwMaintenanceDate(v time.Time) {
	o.EndofSwMaintenanceDate = &v
}

// GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateEpoch() int64 {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSwMaintenanceDateEpoch
}

// GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDateEpoch, true
}

// HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasEndofSwMaintenanceDateEpoch() bool {
	if o != nil && o.EndofSwMaintenanceDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDateEpoch gets a reference to the given int64 and assigns it to the EndofSwMaintenanceDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetEndofSwMaintenanceDateEpoch(v int64) {
	o.EndofSwMaintenanceDateEpoch = &v
}

// GetHardwareEolUrl returns the HardwareEolUrl field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetHardwareEolUrl() string {
	if o == nil || o.HardwareEolUrl == nil {
		var ret string
		return ret
	}
	return *o.HardwareEolUrl
}

// GetHardwareEolUrlOk returns a tuple with the HardwareEolUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetHardwareEolUrlOk() (*string, bool) {
	if o == nil || o.HardwareEolUrl == nil {
		return nil, false
	}
	return o.HardwareEolUrl, true
}

// HasHardwareEolUrl returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasHardwareEolUrl() bool {
	if o != nil && o.HardwareEolUrl != nil {
		return true
	}

	return false
}

// SetHardwareEolUrl gets a reference to the given string and assigns it to the HardwareEolUrl field.
func (o *NiaapiHardwareEolAllOf) SetHardwareEolUrl(v string) {
	o.HardwareEolUrl = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *NiaapiHardwareEolAllOf) SetHeadline(v string) {
	o.Headline = &v
}

// GetLastDateofSupport returns the LastDateofSupport field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetLastDateofSupport() time.Time {
	if o == nil || o.LastDateofSupport == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDateofSupport
}

// GetLastDateofSupportOk returns a tuple with the LastDateofSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportOk() (*time.Time, bool) {
	if o == nil || o.LastDateofSupport == nil {
		return nil, false
	}
	return o.LastDateofSupport, true
}

// HasLastDateofSupport returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasLastDateofSupport() bool {
	if o != nil && o.LastDateofSupport != nil {
		return true
	}

	return false
}

// SetLastDateofSupport gets a reference to the given time.Time and assigns it to the LastDateofSupport field.
func (o *NiaapiHardwareEolAllOf) SetLastDateofSupport(v time.Time) {
	o.LastDateofSupport = &v
}

// GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportEpoch() int64 {
	if o == nil || o.LastDateofSupportEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastDateofSupportEpoch
}

// GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportEpochOk() (*int64, bool) {
	if o == nil || o.LastDateofSupportEpoch == nil {
		return nil, false
	}
	return o.LastDateofSupportEpoch, true
}

// HasLastDateofSupportEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasLastDateofSupportEpoch() bool {
	if o != nil && o.LastDateofSupportEpoch != nil {
		return true
	}

	return false
}

// SetLastDateofSupportEpoch gets a reference to the given int64 and assigns it to the LastDateofSupportEpoch field.
func (o *NiaapiHardwareEolAllOf) SetLastDateofSupportEpoch(v int64) {
	o.LastDateofSupportEpoch = &v
}

// GetLastShipDate returns the LastShipDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetLastShipDate() time.Time {
	if o == nil || o.LastShipDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastShipDate
}

// GetLastShipDateOk returns a tuple with the LastShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetLastShipDateOk() (*time.Time, bool) {
	if o == nil || o.LastShipDate == nil {
		return nil, false
	}
	return o.LastShipDate, true
}

// HasLastShipDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasLastShipDate() bool {
	if o != nil && o.LastShipDate != nil {
		return true
	}

	return false
}

// SetLastShipDate gets a reference to the given time.Time and assigns it to the LastShipDate field.
func (o *NiaapiHardwareEolAllOf) SetLastShipDate(v time.Time) {
	o.LastShipDate = &v
}

// GetLastShipDateEpoch returns the LastShipDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetLastShipDateEpoch() int64 {
	if o == nil || o.LastShipDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastShipDateEpoch
}

// GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetLastShipDateEpochOk() (*int64, bool) {
	if o == nil || o.LastShipDateEpoch == nil {
		return nil, false
	}
	return o.LastShipDateEpoch, true
}

// HasLastShipDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasLastShipDateEpoch() bool {
	if o != nil && o.LastShipDateEpoch != nil {
		return true
	}

	return false
}

// SetLastShipDateEpoch gets a reference to the given int64 and assigns it to the LastShipDateEpoch field.
func (o *NiaapiHardwareEolAllOf) SetLastShipDateEpoch(v int64) {
	o.LastShipDateEpoch = &v
}

// GetMigrationUrl returns the MigrationUrl field value if set, zero value otherwise.
func (o *NiaapiHardwareEolAllOf) GetMigrationUrl() string {
	if o == nil || o.MigrationUrl == nil {
		var ret string
		return ret
	}
	return *o.MigrationUrl
}

// GetMigrationUrlOk returns a tuple with the MigrationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEolAllOf) GetMigrationUrlOk() (*string, bool) {
	if o == nil || o.MigrationUrl == nil {
		return nil, false
	}
	return o.MigrationUrl, true
}

// HasMigrationUrl returns a boolean if a field has been set.
func (o *NiaapiHardwareEolAllOf) HasMigrationUrl() bool {
	if o != nil && o.MigrationUrl != nil {
		return true
	}

	return false
}

// SetMigrationUrl gets a reference to the given string and assigns it to the MigrationUrl field.
func (o *NiaapiHardwareEolAllOf) SetMigrationUrl(v string) {
	o.MigrationUrl = &v
}

func (o NiaapiHardwareEolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AffectedPids != nil {
		toSerialize["AffectedPids"] = o.AffectedPids
	}
	if o.AnnouncementDate != nil {
		toSerialize["AnnouncementDate"] = o.AnnouncementDate
	}
	if o.AnnouncementDateEpoch != nil {
		toSerialize["AnnouncementDateEpoch"] = o.AnnouncementDateEpoch
	}
	if o.BulletinNo != nil {
		toSerialize["BulletinNo"] = o.BulletinNo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndofNewServiceAttachmentDate != nil {
		toSerialize["EndofNewServiceAttachmentDate"] = o.EndofNewServiceAttachmentDate
	}
	if o.EndofNewServiceAttachmentDateEpoch != nil {
		toSerialize["EndofNewServiceAttachmentDateEpoch"] = o.EndofNewServiceAttachmentDateEpoch
	}
	if o.EndofRoutineFailureAnalysisDate != nil {
		toSerialize["EndofRoutineFailureAnalysisDate"] = o.EndofRoutineFailureAnalysisDate
	}
	if o.EndofRoutineFailureAnalysisDateEpoch != nil {
		toSerialize["EndofRoutineFailureAnalysisDateEpoch"] = o.EndofRoutineFailureAnalysisDateEpoch
	}
	if o.EndofSaleDate != nil {
		toSerialize["EndofSaleDate"] = o.EndofSaleDate
	}
	if o.EndofSaleDateEpoch != nil {
		toSerialize["EndofSaleDateEpoch"] = o.EndofSaleDateEpoch
	}
	if o.EndofSecuritySupport != nil {
		toSerialize["EndofSecuritySupport"] = o.EndofSecuritySupport
	}
	if o.EndofSecuritySupportEpoch != nil {
		toSerialize["EndofSecuritySupportEpoch"] = o.EndofSecuritySupportEpoch
	}
	if o.EndofServiceContractRenewalDate != nil {
		toSerialize["EndofServiceContractRenewalDate"] = o.EndofServiceContractRenewalDate
	}
	if o.EndofServiceContractRenewalDateEpoch != nil {
		toSerialize["EndofServiceContractRenewalDateEpoch"] = o.EndofServiceContractRenewalDateEpoch
	}
	if o.EndofSwMaintenanceDate != nil {
		toSerialize["EndofSwMaintenanceDate"] = o.EndofSwMaintenanceDate
	}
	if o.EndofSwMaintenanceDateEpoch != nil {
		toSerialize["EndofSwMaintenanceDateEpoch"] = o.EndofSwMaintenanceDateEpoch
	}
	if o.HardwareEolUrl != nil {
		toSerialize["HardwareEolUrl"] = o.HardwareEolUrl
	}
	if o.Headline != nil {
		toSerialize["Headline"] = o.Headline
	}
	if o.LastDateofSupport != nil {
		toSerialize["LastDateofSupport"] = o.LastDateofSupport
	}
	if o.LastDateofSupportEpoch != nil {
		toSerialize["LastDateofSupportEpoch"] = o.LastDateofSupportEpoch
	}
	if o.LastShipDate != nil {
		toSerialize["LastShipDate"] = o.LastShipDate
	}
	if o.LastShipDateEpoch != nil {
		toSerialize["LastShipDateEpoch"] = o.LastShipDateEpoch
	}
	if o.MigrationUrl != nil {
		toSerialize["MigrationUrl"] = o.MigrationUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiHardwareEolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiHardwareEolAllOf := _NiaapiHardwareEolAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiHardwareEolAllOf); err == nil {
		*o = NiaapiHardwareEolAllOf(varNiaapiHardwareEolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffectedPids")
		delete(additionalProperties, "AnnouncementDate")
		delete(additionalProperties, "AnnouncementDateEpoch")
		delete(additionalProperties, "BulletinNo")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndofNewServiceAttachmentDate")
		delete(additionalProperties, "EndofNewServiceAttachmentDateEpoch")
		delete(additionalProperties, "EndofRoutineFailureAnalysisDate")
		delete(additionalProperties, "EndofRoutineFailureAnalysisDateEpoch")
		delete(additionalProperties, "EndofSaleDate")
		delete(additionalProperties, "EndofSaleDateEpoch")
		delete(additionalProperties, "EndofSecuritySupport")
		delete(additionalProperties, "EndofSecuritySupportEpoch")
		delete(additionalProperties, "EndofServiceContractRenewalDate")
		delete(additionalProperties, "EndofServiceContractRenewalDateEpoch")
		delete(additionalProperties, "EndofSwMaintenanceDate")
		delete(additionalProperties, "EndofSwMaintenanceDateEpoch")
		delete(additionalProperties, "HardwareEolUrl")
		delete(additionalProperties, "Headline")
		delete(additionalProperties, "LastDateofSupport")
		delete(additionalProperties, "LastDateofSupportEpoch")
		delete(additionalProperties, "LastShipDate")
		delete(additionalProperties, "LastShipDateEpoch")
		delete(additionalProperties, "MigrationUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiHardwareEolAllOf struct {
	value *NiaapiHardwareEolAllOf
	isSet bool
}

func (v NullableNiaapiHardwareEolAllOf) Get() *NiaapiHardwareEolAllOf {
	return v.value
}

func (v *NullableNiaapiHardwareEolAllOf) Set(val *NiaapiHardwareEolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiHardwareEolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiHardwareEolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiHardwareEolAllOf(val *NiaapiHardwareEolAllOf) *NullableNiaapiHardwareEolAllOf {
	return &NullableNiaapiHardwareEolAllOf{value: val, isSet: true}
}

func (v NullableNiaapiHardwareEolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiHardwareEolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
