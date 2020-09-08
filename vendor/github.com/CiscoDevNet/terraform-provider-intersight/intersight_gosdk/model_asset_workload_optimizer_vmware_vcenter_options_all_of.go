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

// AssetWorkloadOptimizerVmwareVcenterOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerVmwareVcenterOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerVmwareVcenterOptionsAllOf struct {
	// DatastoreBrowsingEnabled controls whether Workload Optimizer scans vCenter datastores to identify files which are not used and can be deleted to reclaim space and improve actual disk utilization. For example orphaned VMDK files.
	DatastoreBrowsingEnabled *bool `json:"DatastoreBrowsingEnabled,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerVmwareVcenterOptionsAllOf AssetWorkloadOptimizerVmwareVcenterOptionsAllOf

// NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOf instantiates a new AssetWorkloadOptimizerVmwareVcenterOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOf() *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf {
	this := AssetWorkloadOptimizerVmwareVcenterOptionsAllOf{}
	return &this
}

// NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerVmwareVcenterOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOfWithDefaults() *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf {
	this := AssetWorkloadOptimizerVmwareVcenterOptionsAllOf{}
	return &this
}

// GetDatastoreBrowsingEnabled returns the DatastoreBrowsingEnabled field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetDatastoreBrowsingEnabled() bool {
	if o == nil || o.DatastoreBrowsingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DatastoreBrowsingEnabled
}

// GetDatastoreBrowsingEnabledOk returns a tuple with the DatastoreBrowsingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetDatastoreBrowsingEnabledOk() (*bool, bool) {
	if o == nil || o.DatastoreBrowsingEnabled == nil {
		return nil, false
	}
	return o.DatastoreBrowsingEnabled, true
}

// HasDatastoreBrowsingEnabled returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) HasDatastoreBrowsingEnabled() bool {
	if o != nil && o.DatastoreBrowsingEnabled != nil {
		return true
	}

	return false
}

// SetDatastoreBrowsingEnabled gets a reference to the given bool and assigns it to the DatastoreBrowsingEnabled field.
func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) SetDatastoreBrowsingEnabled(v bool) {
	o.DatastoreBrowsingEnabled = &v
}

func (o AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatastoreBrowsingEnabled != nil {
		toSerialize["DatastoreBrowsingEnabled"] = o.DatastoreBrowsingEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerVmwareVcenterOptionsAllOf := _AssetWorkloadOptimizerVmwareVcenterOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerVmwareVcenterOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerVmwareVcenterOptionsAllOf(varAssetWorkloadOptimizerVmwareVcenterOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DatastoreBrowsingEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf struct {
	value *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) Get() *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) Set(val *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf(val *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) *NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf {
	return &NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
