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

// EquipmentPhysicalIdentityRelationship - A relationship to the 'equipment.PhysicalIdentity' resource, or the expanded 'equipment.PhysicalIdentity' resource, or the 'null' value.
type EquipmentPhysicalIdentityRelationship struct {
	EquipmentPhysicalIdentity *EquipmentPhysicalIdentity
	MoMoRef                   *MoMoRef
}

// EquipmentPhysicalIdentityAsEquipmentPhysicalIdentityRelationship is a convenience function that returns EquipmentPhysicalIdentity wrapped in EquipmentPhysicalIdentityRelationship
func EquipmentPhysicalIdentityAsEquipmentPhysicalIdentityRelationship(v *EquipmentPhysicalIdentity) EquipmentPhysicalIdentityRelationship {
	return EquipmentPhysicalIdentityRelationship{EquipmentPhysicalIdentity: v}
}

// MoMoRefAsEquipmentPhysicalIdentityRelationship is a convenience function that returns MoMoRef wrapped in EquipmentPhysicalIdentityRelationship
func MoMoRefAsEquipmentPhysicalIdentityRelationship(v *MoMoRef) EquipmentPhysicalIdentityRelationship {
	return EquipmentPhysicalIdentityRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentPhysicalIdentityRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'equipment.PhysicalIdentity'
	if jsonDict["ClassId"] == "equipment.PhysicalIdentity" {
		// try to unmarshal JSON data into EquipmentPhysicalIdentity
		err = json.Unmarshal(data, &dst.EquipmentPhysicalIdentity)
		if err == nil {
			return nil // data stored in dst.EquipmentPhysicalIdentity, return on the first match
		} else {
			dst.EquipmentPhysicalIdentity = nil
			return fmt.Errorf("Failed to unmarshal EquipmentPhysicalIdentityRelationship as EquipmentPhysicalIdentity: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal EquipmentPhysicalIdentityRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentPhysicalIdentityRelationship) MarshalJSON() ([]byte, error) {
	if src.EquipmentPhysicalIdentity != nil {
		return json.Marshal(&src.EquipmentPhysicalIdentity)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EquipmentPhysicalIdentityRelationship) GetActualInstance() interface{} {
	if obj.EquipmentPhysicalIdentity != nil {
		return obj.EquipmentPhysicalIdentity
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEquipmentPhysicalIdentityRelationship struct {
	value *EquipmentPhysicalIdentityRelationship
	isSet bool
}

func (v NullableEquipmentPhysicalIdentityRelationship) Get() *EquipmentPhysicalIdentityRelationship {
	return v.value
}

func (v *NullableEquipmentPhysicalIdentityRelationship) Set(val *EquipmentPhysicalIdentityRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentPhysicalIdentityRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentPhysicalIdentityRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentPhysicalIdentityRelationship(val *EquipmentPhysicalIdentityRelationship) *NullableEquipmentPhysicalIdentityRelationship {
	return &NullableEquipmentPhysicalIdentityRelationship{value: val, isSet: true}
}

func (v NullableEquipmentPhysicalIdentityRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentPhysicalIdentityRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
