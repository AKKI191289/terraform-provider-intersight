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
	"fmt"
)

// HyperflexHealthResponse - The response body of a HTTP GET request for the 'hyperflex.Health' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'hyperflex.Health' resources.
type HyperflexHealthResponse struct {
	HyperflexHealthList  *HyperflexHealthList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// HyperflexHealthListAsHyperflexHealthResponse is a convenience function that returns HyperflexHealthList wrapped in HyperflexHealthResponse
func HyperflexHealthListAsHyperflexHealthResponse(v *HyperflexHealthList) HyperflexHealthResponse {
	return HyperflexHealthResponse{HyperflexHealthList: v}
}

// MoAggregateTransformAsHyperflexHealthResponse is a convenience function that returns MoAggregateTransform wrapped in HyperflexHealthResponse
func MoAggregateTransformAsHyperflexHealthResponse(v *MoAggregateTransform) HyperflexHealthResponse {
	return HyperflexHealthResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsHyperflexHealthResponse is a convenience function that returns MoDocumentCount wrapped in HyperflexHealthResponse
func MoDocumentCountAsHyperflexHealthResponse(v *MoDocumentCount) HyperflexHealthResponse {
	return HyperflexHealthResponse{MoDocumentCount: v}
}

// MoTagSummaryAsHyperflexHealthResponse is a convenience function that returns MoTagSummary wrapped in HyperflexHealthResponse
func MoTagSummaryAsHyperflexHealthResponse(v *MoTagSummary) HyperflexHealthResponse {
	return HyperflexHealthResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexHealthResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'hyperflex.Health.List'
	if jsonDict["ObjectType"] == "hyperflex.Health.List" {
		// try to unmarshal JSON data into HyperflexHealthList
		err = json.Unmarshal(data, &dst.HyperflexHealthList)
		if err == nil {
			return nil // data stored in dst.HyperflexHealthList, return on the first match
		} else {
			dst.HyperflexHealthList = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHealthResponse as HyperflexHealthList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHealthResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHealthResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHealthResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexHealthResponse) MarshalJSON() ([]byte, error) {
	if src.HyperflexHealthList != nil {
		return json.Marshal(&src.HyperflexHealthList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexHealthResponse) GetActualInstance() interface{} {
	if obj.HyperflexHealthList != nil {
		return obj.HyperflexHealthList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexHealthResponse struct {
	value *HyperflexHealthResponse
	isSet bool
}

func (v NullableHyperflexHealthResponse) Get() *HyperflexHealthResponse {
	return v.value
}

func (v *NullableHyperflexHealthResponse) Set(val *HyperflexHealthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthResponse(val *HyperflexHealthResponse) *NullableHyperflexHealthResponse {
	return &NullableHyperflexHealthResponse{value: val, isSet: true}
}

func (v NullableHyperflexHealthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
