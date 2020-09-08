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
	"time"
)

// LicenseAccountLicenseDataAllOf Definition of the list of properties defined in 'license.AccountLicenseData', excluding properties defined in parent classes.
type LicenseAccountLicenseDataAllOf struct {
	// Root user's ID of the account.
	AccountId *string `json:"AccountId,omitempty"`
	// Agent trusted store data.
	AgentData *string `json:"AgentData,omitempty"`
	// Authorization expiration time.
	AuthExpireTime *string `json:"AuthExpireTime,omitempty"`
	// Intial authorization time.
	AuthInitialTime *string `json:"AuthInitialTime,omitempty"`
	// Next time for the authorization.
	AuthNextTime *string `json:"AuthNextTime,omitempty"`
	// Account license data category name.
	Category *string `json:"Category,omitempty"`
	// Default license tier set by user. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type.
	DefaultLicenseType *string `json:"DefaultLicenseType,omitempty"`
	// The detailed error message when there is any error related to license sync of this account.
	ErrorDesc *string `json:"ErrorDesc,omitempty"`
	// Account license data group name.
	Group *string `json:"Group,omitempty"`
	// The highest license tier which is in compliant of this account. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type.
	HighestCompliantLicenseTier *string `json:"HighestCompliantLicenseTier,omitempty"`
	// Specifies last sync time with SA.
	LastSync *time.Time `json:"LastSync,omitempty"`
	// Record's last update datetime.
	LastUpdatedTime *time.Time `json:"LastUpdatedTime,omitempty"`
	// Aggregrated mode for the agent.
	LicenseState *string `json:"LicenseState,omitempty"`
	// Tech-support info of a smart-agent.
	LicenseTechSupportInfo *string `json:"LicenseTechSupportInfo,omitempty"`
	// Registration exipiration time.
	RegisterExpireTime *string `json:"RegisterExpireTime,omitempty"`
	// Initial time of registration.
	RegisterInitialTime *string `json:"RegisterInitialTime,omitempty"`
	// Next time for the license registration.
	RegisterNextTime *string `json:"RegisterNextTime,omitempty"`
	// Registration status of a smart-agent.
	RegistrationStatus *string `json:"RegistrationStatus,omitempty"`
	// License renewal failure message.
	RenewFailureString *string `json:"RenewFailureString,omitempty"`
	// Name of the smart account.
	SmartAccount *string `json:"SmartAccount,omitempty"`
	// Current sync status for the account.
	SyncStatus *string `json:"SyncStatus,omitempty"`
	// Name of the virtual account.
	VirtualAccount *string                        `json:"VirtualAccount,omitempty"`
	Account        *IamAccountRelationship        `json:"Account,omitempty"`
	CustomerOp     *LicenseCustomerOpRelationship `json:"CustomerOp,omitempty"`
	// An array of relationships to licenseLicenseInfo resources.
	Licenseinfos         []LicenseLicenseInfoRelationship      `json:"Licenseinfos,omitempty"`
	SmartlicenseToken    *LicenseSmartlicenseTokenRelationship `json:"SmartlicenseToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseAccountLicenseDataAllOf LicenseAccountLicenseDataAllOf

// NewLicenseAccountLicenseDataAllOf instantiates a new LicenseAccountLicenseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAccountLicenseDataAllOf() *LicenseAccountLicenseDataAllOf {
	this := LicenseAccountLicenseDataAllOf{}
	var defaultLicenseType string = "Base"
	this.DefaultLicenseType = &defaultLicenseType
	var highestCompliantLicenseTier string = "Base"
	this.HighestCompliantLicenseTier = &highestCompliantLicenseTier
	return &this
}

// NewLicenseAccountLicenseDataAllOfWithDefaults instantiates a new LicenseAccountLicenseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAccountLicenseDataAllOfWithDefaults() *LicenseAccountLicenseDataAllOf {
	this := LicenseAccountLicenseDataAllOf{}
	var defaultLicenseType string = "Base"
	this.DefaultLicenseType = &defaultLicenseType
	var highestCompliantLicenseTier string = "Base"
	this.HighestCompliantLicenseTier = &highestCompliantLicenseTier
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LicenseAccountLicenseDataAllOf) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgentData returns the AgentData field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAgentData() string {
	if o == nil || o.AgentData == nil {
		var ret string
		return ret
	}
	return *o.AgentData
}

// GetAgentDataOk returns a tuple with the AgentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAgentDataOk() (*string, bool) {
	if o == nil || o.AgentData == nil {
		return nil, false
	}
	return o.AgentData, true
}

// HasAgentData returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAgentData() bool {
	if o != nil && o.AgentData != nil {
		return true
	}

	return false
}

// SetAgentData gets a reference to the given string and assigns it to the AgentData field.
func (o *LicenseAccountLicenseDataAllOf) SetAgentData(v string) {
	o.AgentData = &v
}

// GetAuthExpireTime returns the AuthExpireTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAuthExpireTime() string {
	if o == nil || o.AuthExpireTime == nil {
		var ret string
		return ret
	}
	return *o.AuthExpireTime
}

// GetAuthExpireTimeOk returns a tuple with the AuthExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAuthExpireTimeOk() (*string, bool) {
	if o == nil || o.AuthExpireTime == nil {
		return nil, false
	}
	return o.AuthExpireTime, true
}

// HasAuthExpireTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAuthExpireTime() bool {
	if o != nil && o.AuthExpireTime != nil {
		return true
	}

	return false
}

// SetAuthExpireTime gets a reference to the given string and assigns it to the AuthExpireTime field.
func (o *LicenseAccountLicenseDataAllOf) SetAuthExpireTime(v string) {
	o.AuthExpireTime = &v
}

// GetAuthInitialTime returns the AuthInitialTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAuthInitialTime() string {
	if o == nil || o.AuthInitialTime == nil {
		var ret string
		return ret
	}
	return *o.AuthInitialTime
}

// GetAuthInitialTimeOk returns a tuple with the AuthInitialTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAuthInitialTimeOk() (*string, bool) {
	if o == nil || o.AuthInitialTime == nil {
		return nil, false
	}
	return o.AuthInitialTime, true
}

// HasAuthInitialTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAuthInitialTime() bool {
	if o != nil && o.AuthInitialTime != nil {
		return true
	}

	return false
}

// SetAuthInitialTime gets a reference to the given string and assigns it to the AuthInitialTime field.
func (o *LicenseAccountLicenseDataAllOf) SetAuthInitialTime(v string) {
	o.AuthInitialTime = &v
}

// GetAuthNextTime returns the AuthNextTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAuthNextTime() string {
	if o == nil || o.AuthNextTime == nil {
		var ret string
		return ret
	}
	return *o.AuthNextTime
}

// GetAuthNextTimeOk returns a tuple with the AuthNextTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAuthNextTimeOk() (*string, bool) {
	if o == nil || o.AuthNextTime == nil {
		return nil, false
	}
	return o.AuthNextTime, true
}

// HasAuthNextTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAuthNextTime() bool {
	if o != nil && o.AuthNextTime != nil {
		return true
	}

	return false
}

// SetAuthNextTime gets a reference to the given string and assigns it to the AuthNextTime field.
func (o *LicenseAccountLicenseDataAllOf) SetAuthNextTime(v string) {
	o.AuthNextTime = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *LicenseAccountLicenseDataAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetDefaultLicenseType returns the DefaultLicenseType field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetDefaultLicenseType() string {
	if o == nil || o.DefaultLicenseType == nil {
		var ret string
		return ret
	}
	return *o.DefaultLicenseType
}

// GetDefaultLicenseTypeOk returns a tuple with the DefaultLicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetDefaultLicenseTypeOk() (*string, bool) {
	if o == nil || o.DefaultLicenseType == nil {
		return nil, false
	}
	return o.DefaultLicenseType, true
}

// HasDefaultLicenseType returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasDefaultLicenseType() bool {
	if o != nil && o.DefaultLicenseType != nil {
		return true
	}

	return false
}

// SetDefaultLicenseType gets a reference to the given string and assigns it to the DefaultLicenseType field.
func (o *LicenseAccountLicenseDataAllOf) SetDefaultLicenseType(v string) {
	o.DefaultLicenseType = &v
}

// GetErrorDesc returns the ErrorDesc field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetErrorDesc() string {
	if o == nil || o.ErrorDesc == nil {
		var ret string
		return ret
	}
	return *o.ErrorDesc
}

// GetErrorDescOk returns a tuple with the ErrorDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetErrorDescOk() (*string, bool) {
	if o == nil || o.ErrorDesc == nil {
		return nil, false
	}
	return o.ErrorDesc, true
}

// HasErrorDesc returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasErrorDesc() bool {
	if o != nil && o.ErrorDesc != nil {
		return true
	}

	return false
}

// SetErrorDesc gets a reference to the given string and assigns it to the ErrorDesc field.
func (o *LicenseAccountLicenseDataAllOf) SetErrorDesc(v string) {
	o.ErrorDesc = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *LicenseAccountLicenseDataAllOf) SetGroup(v string) {
	o.Group = &v
}

// GetHighestCompliantLicenseTier returns the HighestCompliantLicenseTier field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetHighestCompliantLicenseTier() string {
	if o == nil || o.HighestCompliantLicenseTier == nil {
		var ret string
		return ret
	}
	return *o.HighestCompliantLicenseTier
}

// GetHighestCompliantLicenseTierOk returns a tuple with the HighestCompliantLicenseTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetHighestCompliantLicenseTierOk() (*string, bool) {
	if o == nil || o.HighestCompliantLicenseTier == nil {
		return nil, false
	}
	return o.HighestCompliantLicenseTier, true
}

// HasHighestCompliantLicenseTier returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasHighestCompliantLicenseTier() bool {
	if o != nil && o.HighestCompliantLicenseTier != nil {
		return true
	}

	return false
}

// SetHighestCompliantLicenseTier gets a reference to the given string and assigns it to the HighestCompliantLicenseTier field.
func (o *LicenseAccountLicenseDataAllOf) SetHighestCompliantLicenseTier(v string) {
	o.HighestCompliantLicenseTier = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetLastSync() time.Time {
	if o == nil || o.LastSync == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || o.LastSync == nil {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasLastSync() bool {
	if o != nil && o.LastSync != nil {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *LicenseAccountLicenseDataAllOf) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *LicenseAccountLicenseDataAllOf) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLicenseState returns the LicenseState field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetLicenseState() string {
	if o == nil || o.LicenseState == nil {
		var ret string
		return ret
	}
	return *o.LicenseState
}

// GetLicenseStateOk returns a tuple with the LicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetLicenseStateOk() (*string, bool) {
	if o == nil || o.LicenseState == nil {
		return nil, false
	}
	return o.LicenseState, true
}

// HasLicenseState returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasLicenseState() bool {
	if o != nil && o.LicenseState != nil {
		return true
	}

	return false
}

// SetLicenseState gets a reference to the given string and assigns it to the LicenseState field.
func (o *LicenseAccountLicenseDataAllOf) SetLicenseState(v string) {
	o.LicenseState = &v
}

// GetLicenseTechSupportInfo returns the LicenseTechSupportInfo field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetLicenseTechSupportInfo() string {
	if o == nil || o.LicenseTechSupportInfo == nil {
		var ret string
		return ret
	}
	return *o.LicenseTechSupportInfo
}

// GetLicenseTechSupportInfoOk returns a tuple with the LicenseTechSupportInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetLicenseTechSupportInfoOk() (*string, bool) {
	if o == nil || o.LicenseTechSupportInfo == nil {
		return nil, false
	}
	return o.LicenseTechSupportInfo, true
}

// HasLicenseTechSupportInfo returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasLicenseTechSupportInfo() bool {
	if o != nil && o.LicenseTechSupportInfo != nil {
		return true
	}

	return false
}

// SetLicenseTechSupportInfo gets a reference to the given string and assigns it to the LicenseTechSupportInfo field.
func (o *LicenseAccountLicenseDataAllOf) SetLicenseTechSupportInfo(v string) {
	o.LicenseTechSupportInfo = &v
}

// GetRegisterExpireTime returns the RegisterExpireTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterExpireTime() string {
	if o == nil || o.RegisterExpireTime == nil {
		var ret string
		return ret
	}
	return *o.RegisterExpireTime
}

// GetRegisterExpireTimeOk returns a tuple with the RegisterExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterExpireTimeOk() (*string, bool) {
	if o == nil || o.RegisterExpireTime == nil {
		return nil, false
	}
	return o.RegisterExpireTime, true
}

// HasRegisterExpireTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasRegisterExpireTime() bool {
	if o != nil && o.RegisterExpireTime != nil {
		return true
	}

	return false
}

// SetRegisterExpireTime gets a reference to the given string and assigns it to the RegisterExpireTime field.
func (o *LicenseAccountLicenseDataAllOf) SetRegisterExpireTime(v string) {
	o.RegisterExpireTime = &v
}

// GetRegisterInitialTime returns the RegisterInitialTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterInitialTime() string {
	if o == nil || o.RegisterInitialTime == nil {
		var ret string
		return ret
	}
	return *o.RegisterInitialTime
}

// GetRegisterInitialTimeOk returns a tuple with the RegisterInitialTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterInitialTimeOk() (*string, bool) {
	if o == nil || o.RegisterInitialTime == nil {
		return nil, false
	}
	return o.RegisterInitialTime, true
}

// HasRegisterInitialTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasRegisterInitialTime() bool {
	if o != nil && o.RegisterInitialTime != nil {
		return true
	}

	return false
}

// SetRegisterInitialTime gets a reference to the given string and assigns it to the RegisterInitialTime field.
func (o *LicenseAccountLicenseDataAllOf) SetRegisterInitialTime(v string) {
	o.RegisterInitialTime = &v
}

// GetRegisterNextTime returns the RegisterNextTime field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterNextTime() string {
	if o == nil || o.RegisterNextTime == nil {
		var ret string
		return ret
	}
	return *o.RegisterNextTime
}

// GetRegisterNextTimeOk returns a tuple with the RegisterNextTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetRegisterNextTimeOk() (*string, bool) {
	if o == nil || o.RegisterNextTime == nil {
		return nil, false
	}
	return o.RegisterNextTime, true
}

// HasRegisterNextTime returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasRegisterNextTime() bool {
	if o != nil && o.RegisterNextTime != nil {
		return true
	}

	return false
}

// SetRegisterNextTime gets a reference to the given string and assigns it to the RegisterNextTime field.
func (o *LicenseAccountLicenseDataAllOf) SetRegisterNextTime(v string) {
	o.RegisterNextTime = &v
}

// GetRegistrationStatus returns the RegistrationStatus field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetRegistrationStatus() string {
	if o == nil || o.RegistrationStatus == nil {
		var ret string
		return ret
	}
	return *o.RegistrationStatus
}

// GetRegistrationStatusOk returns a tuple with the RegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetRegistrationStatusOk() (*string, bool) {
	if o == nil || o.RegistrationStatus == nil {
		return nil, false
	}
	return o.RegistrationStatus, true
}

// HasRegistrationStatus returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasRegistrationStatus() bool {
	if o != nil && o.RegistrationStatus != nil {
		return true
	}

	return false
}

// SetRegistrationStatus gets a reference to the given string and assigns it to the RegistrationStatus field.
func (o *LicenseAccountLicenseDataAllOf) SetRegistrationStatus(v string) {
	o.RegistrationStatus = &v
}

// GetRenewFailureString returns the RenewFailureString field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetRenewFailureString() string {
	if o == nil || o.RenewFailureString == nil {
		var ret string
		return ret
	}
	return *o.RenewFailureString
}

// GetRenewFailureStringOk returns a tuple with the RenewFailureString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetRenewFailureStringOk() (*string, bool) {
	if o == nil || o.RenewFailureString == nil {
		return nil, false
	}
	return o.RenewFailureString, true
}

// HasRenewFailureString returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasRenewFailureString() bool {
	if o != nil && o.RenewFailureString != nil {
		return true
	}

	return false
}

// SetRenewFailureString gets a reference to the given string and assigns it to the RenewFailureString field.
func (o *LicenseAccountLicenseDataAllOf) SetRenewFailureString(v string) {
	o.RenewFailureString = &v
}

// GetSmartAccount returns the SmartAccount field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetSmartAccount() string {
	if o == nil || o.SmartAccount == nil {
		var ret string
		return ret
	}
	return *o.SmartAccount
}

// GetSmartAccountOk returns a tuple with the SmartAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetSmartAccountOk() (*string, bool) {
	if o == nil || o.SmartAccount == nil {
		return nil, false
	}
	return o.SmartAccount, true
}

// HasSmartAccount returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasSmartAccount() bool {
	if o != nil && o.SmartAccount != nil {
		return true
	}

	return false
}

// SetSmartAccount gets a reference to the given string and assigns it to the SmartAccount field.
func (o *LicenseAccountLicenseDataAllOf) SetSmartAccount(v string) {
	o.SmartAccount = &v
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetSyncStatus() string {
	if o == nil || o.SyncStatus == nil {
		var ret string
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetSyncStatusOk() (*string, bool) {
	if o == nil || o.SyncStatus == nil {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasSyncStatus() bool {
	if o != nil && o.SyncStatus != nil {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given string and assigns it to the SyncStatus field.
func (o *LicenseAccountLicenseDataAllOf) SetSyncStatus(v string) {
	o.SyncStatus = &v
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetVirtualAccount() string {
	if o == nil || o.VirtualAccount == nil {
		var ret string
		return ret
	}
	return *o.VirtualAccount
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetVirtualAccountOk() (*string, bool) {
	if o == nil || o.VirtualAccount == nil {
		return nil, false
	}
	return o.VirtualAccount, true
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasVirtualAccount() bool {
	if o != nil && o.VirtualAccount != nil {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given string and assigns it to the VirtualAccount field.
func (o *LicenseAccountLicenseDataAllOf) SetVirtualAccount(v string) {
	o.VirtualAccount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *LicenseAccountLicenseDataAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetCustomerOp returns the CustomerOp field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetCustomerOp() LicenseCustomerOpRelationship {
	if o == nil || o.CustomerOp == nil {
		var ret LicenseCustomerOpRelationship
		return ret
	}
	return *o.CustomerOp
}

// GetCustomerOpOk returns a tuple with the CustomerOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetCustomerOpOk() (*LicenseCustomerOpRelationship, bool) {
	if o == nil || o.CustomerOp == nil {
		return nil, false
	}
	return o.CustomerOp, true
}

// HasCustomerOp returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasCustomerOp() bool {
	if o != nil && o.CustomerOp != nil {
		return true
	}

	return false
}

// SetCustomerOp gets a reference to the given LicenseCustomerOpRelationship and assigns it to the CustomerOp field.
func (o *LicenseAccountLicenseDataAllOf) SetCustomerOp(v LicenseCustomerOpRelationship) {
	o.CustomerOp = &v
}

// GetLicenseinfos returns the Licenseinfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseAccountLicenseDataAllOf) GetLicenseinfos() []LicenseLicenseInfoRelationship {
	if o == nil {
		var ret []LicenseLicenseInfoRelationship
		return ret
	}
	return o.Licenseinfos
}

// GetLicenseinfosOk returns a tuple with the Licenseinfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseAccountLicenseDataAllOf) GetLicenseinfosOk() (*[]LicenseLicenseInfoRelationship, bool) {
	if o == nil || o.Licenseinfos == nil {
		return nil, false
	}
	return &o.Licenseinfos, true
}

// HasLicenseinfos returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasLicenseinfos() bool {
	if o != nil && o.Licenseinfos != nil {
		return true
	}

	return false
}

// SetLicenseinfos gets a reference to the given []LicenseLicenseInfoRelationship and assigns it to the Licenseinfos field.
func (o *LicenseAccountLicenseDataAllOf) SetLicenseinfos(v []LicenseLicenseInfoRelationship) {
	o.Licenseinfos = v
}

// GetSmartlicenseToken returns the SmartlicenseToken field value if set, zero value otherwise.
func (o *LicenseAccountLicenseDataAllOf) GetSmartlicenseToken() LicenseSmartlicenseTokenRelationship {
	if o == nil || o.SmartlicenseToken == nil {
		var ret LicenseSmartlicenseTokenRelationship
		return ret
	}
	return *o.SmartlicenseToken
}

// GetSmartlicenseTokenOk returns a tuple with the SmartlicenseToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAccountLicenseDataAllOf) GetSmartlicenseTokenOk() (*LicenseSmartlicenseTokenRelationship, bool) {
	if o == nil || o.SmartlicenseToken == nil {
		return nil, false
	}
	return o.SmartlicenseToken, true
}

// HasSmartlicenseToken returns a boolean if a field has been set.
func (o *LicenseAccountLicenseDataAllOf) HasSmartlicenseToken() bool {
	if o != nil && o.SmartlicenseToken != nil {
		return true
	}

	return false
}

// SetSmartlicenseToken gets a reference to the given LicenseSmartlicenseTokenRelationship and assigns it to the SmartlicenseToken field.
func (o *LicenseAccountLicenseDataAllOf) SetSmartlicenseToken(v LicenseSmartlicenseTokenRelationship) {
	o.SmartlicenseToken = &v
}

func (o LicenseAccountLicenseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.AgentData != nil {
		toSerialize["AgentData"] = o.AgentData
	}
	if o.AuthExpireTime != nil {
		toSerialize["AuthExpireTime"] = o.AuthExpireTime
	}
	if o.AuthInitialTime != nil {
		toSerialize["AuthInitialTime"] = o.AuthInitialTime
	}
	if o.AuthNextTime != nil {
		toSerialize["AuthNextTime"] = o.AuthNextTime
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.DefaultLicenseType != nil {
		toSerialize["DefaultLicenseType"] = o.DefaultLicenseType
	}
	if o.ErrorDesc != nil {
		toSerialize["ErrorDesc"] = o.ErrorDesc
	}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.HighestCompliantLicenseTier != nil {
		toSerialize["HighestCompliantLicenseTier"] = o.HighestCompliantLicenseTier
	}
	if o.LastSync != nil {
		toSerialize["LastSync"] = o.LastSync
	}
	if o.LastUpdatedTime != nil {
		toSerialize["LastUpdatedTime"] = o.LastUpdatedTime
	}
	if o.LicenseState != nil {
		toSerialize["LicenseState"] = o.LicenseState
	}
	if o.LicenseTechSupportInfo != nil {
		toSerialize["LicenseTechSupportInfo"] = o.LicenseTechSupportInfo
	}
	if o.RegisterExpireTime != nil {
		toSerialize["RegisterExpireTime"] = o.RegisterExpireTime
	}
	if o.RegisterInitialTime != nil {
		toSerialize["RegisterInitialTime"] = o.RegisterInitialTime
	}
	if o.RegisterNextTime != nil {
		toSerialize["RegisterNextTime"] = o.RegisterNextTime
	}
	if o.RegistrationStatus != nil {
		toSerialize["RegistrationStatus"] = o.RegistrationStatus
	}
	if o.RenewFailureString != nil {
		toSerialize["RenewFailureString"] = o.RenewFailureString
	}
	if o.SmartAccount != nil {
		toSerialize["SmartAccount"] = o.SmartAccount
	}
	if o.SyncStatus != nil {
		toSerialize["SyncStatus"] = o.SyncStatus
	}
	if o.VirtualAccount != nil {
		toSerialize["VirtualAccount"] = o.VirtualAccount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.CustomerOp != nil {
		toSerialize["CustomerOp"] = o.CustomerOp
	}
	if o.Licenseinfos != nil {
		toSerialize["Licenseinfos"] = o.Licenseinfos
	}
	if o.SmartlicenseToken != nil {
		toSerialize["SmartlicenseToken"] = o.SmartlicenseToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseAccountLicenseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseAccountLicenseDataAllOf := _LicenseAccountLicenseDataAllOf{}

	if err = json.Unmarshal(bytes, &varLicenseAccountLicenseDataAllOf); err == nil {
		*o = LicenseAccountLicenseDataAllOf(varLicenseAccountLicenseDataAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AccountId")
		delete(additionalProperties, "AgentData")
		delete(additionalProperties, "AuthExpireTime")
		delete(additionalProperties, "AuthInitialTime")
		delete(additionalProperties, "AuthNextTime")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DefaultLicenseType")
		delete(additionalProperties, "ErrorDesc")
		delete(additionalProperties, "Group")
		delete(additionalProperties, "HighestCompliantLicenseTier")
		delete(additionalProperties, "LastSync")
		delete(additionalProperties, "LastUpdatedTime")
		delete(additionalProperties, "LicenseState")
		delete(additionalProperties, "LicenseTechSupportInfo")
		delete(additionalProperties, "RegisterExpireTime")
		delete(additionalProperties, "RegisterInitialTime")
		delete(additionalProperties, "RegisterNextTime")
		delete(additionalProperties, "RegistrationStatus")
		delete(additionalProperties, "RenewFailureString")
		delete(additionalProperties, "SmartAccount")
		delete(additionalProperties, "SyncStatus")
		delete(additionalProperties, "VirtualAccount")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "CustomerOp")
		delete(additionalProperties, "Licenseinfos")
		delete(additionalProperties, "SmartlicenseToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseAccountLicenseDataAllOf struct {
	value *LicenseAccountLicenseDataAllOf
	isSet bool
}

func (v NullableLicenseAccountLicenseDataAllOf) Get() *LicenseAccountLicenseDataAllOf {
	return v.value
}

func (v *NullableLicenseAccountLicenseDataAllOf) Set(val *LicenseAccountLicenseDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAccountLicenseDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAccountLicenseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAccountLicenseDataAllOf(val *LicenseAccountLicenseDataAllOf) *NullableLicenseAccountLicenseDataAllOf {
	return &NullableLicenseAccountLicenseDataAllOf{value: val, isSet: true}
}

func (v NullableLicenseAccountLicenseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAccountLicenseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
