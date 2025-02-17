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

// BulkRequest The bulk.Request API allows users to perform API actions (Create, Update or Delete) in bulk, on a given URI. It is possible to operate on multiple subpaths relative to the provided URI (For example, it would be possible to perform a PATCH action on multiple objects of a given REST resource type).
type BulkRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string           `json:"ObjectType"`
	Requests   []BulkSubRequest `json:"Requests,omitempty"`
	Results    []BulkApiResult  `json:"Results,omitempty"`
	// The URI on which this bulk action is to be performed.
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	Verb                 *string                               `json:"Verb,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRequest BulkRequest

// NewBulkRequest instantiates a new BulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRequest(classId string, objectType string) *BulkRequest {
	this := BulkRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkRequestWithDefaults instantiates a new BulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRequestWithDefaults() *BulkRequest {
	this := BulkRequest{}
	var classId string = "bulk.Request"
	this.ClassId = classId
	var objectType string = "bulk.Request"
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRequests returns the Requests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetRequests() []BulkSubRequest {
	if o == nil {
		var ret []BulkSubRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetRequestsOk() (*[]BulkSubRequest, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return &o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BulkRequest) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []BulkSubRequest and assigns it to the Requests field.
func (o *BulkRequest) SetRequests(v []BulkSubRequest) {
	o.Requests = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetResults() []BulkApiResult {
	if o == nil {
		var ret []BulkApiResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetResultsOk() (*[]BulkApiResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkRequest) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkApiResult and assigns it to the Results field.
func (o *BulkRequest) SetResults(v []BulkApiResult) {
	o.Results = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkRequest) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkRequest) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkRequest) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *BulkRequest) GetVerb() string {
	if o == nil || o.Verb == nil {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetVerbOk() (*string, bool) {
	if o == nil || o.Verb == nil {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkRequest) HasVerb() bool {
	if o != nil && o.Verb != nil {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *BulkRequest) SetVerb(v string) {
	o.Verb = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkRequest) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkRequest) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkRequest) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkRequest) MarshalJSON() ([]byte, error) {
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
	if o.Requests != nil {
		toSerialize["Requests"] = o.Requests
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.Verb != nil {
		toSerialize["Verb"] = o.Verb
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkRequest) UnmarshalJSON(bytes []byte) (err error) {
	type BulkRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string           `json:"ObjectType"`
		Requests   []BulkSubRequest `json:"Requests,omitempty"`
		Results    []BulkApiResult  `json:"Results,omitempty"`
		// The URI on which this bulk action is to be performed.
		Uri *string `json:"Uri,omitempty"`
		// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
		Verb         *string                               `json:"Verb,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkRequestWithoutEmbeddedStruct := BulkRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkRequestWithoutEmbeddedStruct)
	if err == nil {
		varBulkRequest := _BulkRequest{}
		varBulkRequest.ClassId = varBulkRequestWithoutEmbeddedStruct.ClassId
		varBulkRequest.ObjectType = varBulkRequestWithoutEmbeddedStruct.ObjectType
		varBulkRequest.Requests = varBulkRequestWithoutEmbeddedStruct.Requests
		varBulkRequest.Results = varBulkRequestWithoutEmbeddedStruct.Results
		varBulkRequest.Uri = varBulkRequestWithoutEmbeddedStruct.Uri
		varBulkRequest.Verb = varBulkRequestWithoutEmbeddedStruct.Verb
		varBulkRequest.Organization = varBulkRequestWithoutEmbeddedStruct.Organization
		*o = BulkRequest(varBulkRequest)
	} else {
		return err
	}

	varBulkRequest := _BulkRequest{}

	err = json.Unmarshal(bytes, &varBulkRequest)
	if err == nil {
		o.MoBaseMo = varBulkRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Requests")
		delete(additionalProperties, "Results")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		delete(additionalProperties, "Organization")

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

type NullableBulkRequest struct {
	value *BulkRequest
	isSet bool
}

func (v NullableBulkRequest) Get() *BulkRequest {
	return v.value
}

func (v *NullableBulkRequest) Set(val *BulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRequest(val *BulkRequest) *NullableBulkRequest {
	return &NullableBulkRequest{value: val, isSet: true}
}

func (v NullableBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
