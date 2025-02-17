package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCondHclStatusDetail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCondHclStatusDetailRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hardware_status": {
				Description: "The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_cimc_version": {
				Description: "The current CIMC version for the server normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_driver_name": {
				Description: "The current driver name of the component we are validating normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_driver_version": {
				Description: "The current driver version of the component we are validating normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_firmware_version": {
				Description: "The current firmware version of the component model normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_model": {
				Description: "The component model we are trying to validate normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_cimc_version": {
				Description: "The current CIMC version for the server as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_driver_name": {
				Description: "The current driver name of the component we are validating as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_driver_version": {
				Description: "The current driver version of the component we are validating as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_firmware_version": {
				Description: "The current firmware version of the component model as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_model": {
				Description: "The component model we are trying to validate as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"reason": {
				Description: "The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_status": {
				Description: "The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"component": {
						Description: "A reference to a inventoryBase resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"hardware_status": {
						Description: "The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_cimc_version": {
						Description: "The current CIMC version for the server normalized for querying HCL data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_driver_name": {
						Description: "The current driver name of the component we are validating normalized for querying HCL data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_driver_version": {
						Description: "The current driver version of the component we are validating normalized for querying HCL data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_firmware_version": {
						Description: "The current firmware version of the component model normalized for querying HCL data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_model": {
						Description: "The component model we are trying to validate normalized for querying HCL data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_status": {
						Description: "A reference to a condHclStatus resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"inv_cimc_version": {
						Description: "The current CIMC version for the server as received from inventory.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_driver_name": {
						Description: "The current driver name of the component we are validating as received from inventory.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_driver_version": {
						Description: "The current driver version of the component we are validating as received from inventory.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_firmware_version": {
						Description: "The current firmware version of the component model as received from inventory.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_model": {
						Description: "The component model we are trying to validate as received from inventory.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"reason": {
						Description: "The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"software_status": {
						Description: "The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.\n* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.\n* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.\n* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.\n* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.\n* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.\n* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.\n* `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceCondHclStatusDetailRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CondHclStatusDetail{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("hardware_status"); ok {
		x := (v.(string))
		o.SetHardwareStatus(x)
	}
	if v, ok := d.GetOk("hcl_cimc_version"); ok {
		x := (v.(string))
		o.SetHclCimcVersion(x)
	}
	if v, ok := d.GetOk("hcl_driver_name"); ok {
		x := (v.(string))
		o.SetHclDriverName(x)
	}
	if v, ok := d.GetOk("hcl_driver_version"); ok {
		x := (v.(string))
		o.SetHclDriverVersion(x)
	}
	if v, ok := d.GetOk("hcl_firmware_version"); ok {
		x := (v.(string))
		o.SetHclFirmwareVersion(x)
	}
	if v, ok := d.GetOk("hcl_model"); ok {
		x := (v.(string))
		o.SetHclModel(x)
	}
	if v, ok := d.GetOk("inv_cimc_version"); ok {
		x := (v.(string))
		o.SetInvCimcVersion(x)
	}
	if v, ok := d.GetOk("inv_driver_name"); ok {
		x := (v.(string))
		o.SetInvDriverName(x)
	}
	if v, ok := d.GetOk("inv_driver_version"); ok {
		x := (v.(string))
		o.SetInvDriverVersion(x)
	}
	if v, ok := d.GetOk("inv_firmware_version"); ok {
		x := (v.(string))
		o.SetInvFirmwareVersion(x)
	}
	if v, ok := d.GetOk("inv_model"); ok {
		x := (v.(string))
		o.SetInvModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("reason"); ok {
		x := (v.(string))
		o.SetReason(x)
	}
	if v, ok := d.GetOk("software_status"); ok {
		x := (v.(string))
		o.SetSoftwareStatus(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CondHclStatusDetail object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CondApi.GetCondHclStatusDetailList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CondHclStatusDetail: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CondHclStatusDetailList.GetCount()
	var i int32
	var condHclStatusDetailResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CondApi.GetCondHclStatusDetailList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CondHclStatusDetail: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CondHclStatusDetailList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CondHclStatusDetail data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["component"] = flattenMapInventoryBaseRelationship(s.GetComponent(), d)
				temp["hardware_status"] = (s.GetHardwareStatus())
				temp["hcl_cimc_version"] = (s.GetHclCimcVersion())
				temp["hcl_driver_name"] = (s.GetHclDriverName())
				temp["hcl_driver_version"] = (s.GetHclDriverVersion())
				temp["hcl_firmware_version"] = (s.GetHclFirmwareVersion())
				temp["hcl_model"] = (s.GetHclModel())

				temp["hcl_status"] = flattenMapCondHclStatusRelationship(s.GetHclStatus(), d)
				temp["inv_cimc_version"] = (s.GetInvCimcVersion())
				temp["inv_driver_name"] = (s.GetInvDriverName())
				temp["inv_driver_version"] = (s.GetInvDriverVersion())
				temp["inv_firmware_version"] = (s.GetInvFirmwareVersion())
				temp["inv_model"] = (s.GetInvModel())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["reason"] = (s.GetReason())
				temp["software_status"] = (s.GetSoftwareStatus())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				condHclStatusDetailResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(condHclStatusDetailResults))
	if err := d.Set("results", condHclStatusDetailResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(condHclStatusDetailResults[0]["moid"].(string))
	return de
}
