package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCondHclStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCondHclStatusRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"component_status": {
				Description: "The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \"Validated\" - all the components hardware/software profiles are listed in the HCL. \"Not-Listed\" - one or more components hardware/software profiles are not listed in the HCL \"Incomplete\" - the components are not evaluated as the server's software/hardware profiles are not listed in the HCL. \"Not-Evaluated\" - The components are not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hardware_status": {
				Description: "The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \"Validated\" - The server model, processor and firmware combination is listed in the HCL \"Not-Listed\" - The server model, processor and firmware combination is not listed in the HCL. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_firmware_version": {
				Description: "The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_model": {
				Description: "The managed object's model to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_os_vendor": {
				Description: "The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_os_version": {
				Description: "The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_processor": {
				Description: "The managed object's processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_firmware_version": {
				Description: "The current CIMC version for the server as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_model": {
				Description: "The managed object's model to validate as received from the inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_os_vendor": {
				Description: "The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_os_version": {
				Description: "The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_processor": {
				Description: "The managed object's processor to validate if applicable as received from inventory. It is empty if we are missing this information.",
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
				Description: "The reason for the HCL status. It will be one of the following \"Missing-Os-Info\" - we are missing os information in the inventory from the device connector \"Incompatible-Components\" - we have 1 or more components with \"Not-Validated\" status \"Compatible\" - all the components have \"Validated\" status. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Missing-Os-Info` - This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Components` - This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails.\n* `Compatible` - This means the HclStatus for the sever has passed HCL validation for all of its related components.\n* `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"server_reason": {
				Description: "The reason generated by the server's HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server's hardware status 2. Evaluate the server's software status 3. Evaluate the server's components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \"Incompatible-Server\" - the server model is not listed in the HCL. \"Incompatible-Processor\" - the server model and processor combination is not listed in HCL. \"Incompatible-Firmware\" - the server model, processor and server firmware is not listed in HCL. \"Missing-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Components\" - there is one or more components with \"Not-Validated\" status \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). \"Compatible\" - the server and all its components are validated. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information.\n* `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination.\n* `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result.\n* `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.\n* `Incompatible-Components` - The validation has failed for this server because one or more components have \"Not-Listed\" status.\n* `Compatible` - The validation has passed for this server's model, processor, OS vendor and version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_status": {
				Description: "The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \"Validated\" - The os vendor/version is listed in the HCL for the server model, processor and firmware \"Not-Listed\" - The os vendor/version is not listed in the HCL for the server model, processor and firmware \"Incomplete\" - The inventory is missing os vendor/version and HCL validation was not performed. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The HCL compatibility status of the managed object. The status can be one of the following \"Incomplete\" - there is no enough information to evaluate against the HCL data \"Validated\" - all components have been evaluated against the HCL and they all have \"Validated\" status \"Not-Listed\" - all components have been evaluated against the HCL and one or more have \"Not-Listed\" status. \"Not-Evaluated\" - server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
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
					"component_status": {
						Description: "The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \"Validated\" - all the components hardware/software profiles are listed in the HCL. \"Not-Listed\" - one or more components hardware/software profiles are not listed in the HCL \"Incomplete\" - the components are not evaluated as the server's software/hardware profiles are not listed in the HCL. \"Not-Evaluated\" - The components are not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"details": {
						Description: "An array of relationships to condHclStatusDetail resources.",
						Type:        schema.TypeList,
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
						Description: "The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \"Validated\" - The server model, processor and firmware combination is listed in the HCL \"Not-Listed\" - The server model, processor and firmware combination is not listed in the HCL. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_firmware_version": {
						Description: "The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_model": {
						Description: "The managed object's model to validate normalized for querying HCL data. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_os_vendor": {
						Description: "The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_os_version": {
						Description: "The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hcl_processor": {
						Description: "The managed object's processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_firmware_version": {
						Description: "The current CIMC version for the server as received from inventory. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_model": {
						Description: "The managed object's model to validate as received from the inventory. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_os_vendor": {
						Description: "The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_os_version": {
						Description: "The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inv_processor": {
						Description: "The managed object's processor to validate if applicable as received from inventory. It is empty if we are missing this information.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"managed_object": {
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
						Description: "The reason for the HCL status. It will be one of the following \"Missing-Os-Info\" - we are missing os information in the inventory from the device connector \"Incompatible-Components\" - we have 1 or more components with \"Not-Validated\" status \"Compatible\" - all the components have \"Validated\" status. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Missing-Os-Info` - This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Incompatible-Components` - This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails.\n* `Compatible` - This means the HclStatus for the sever has passed HCL validation for all of its related components.\n* `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"server_reason": {
						Description: "The reason generated by the server's HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server's hardware status 2. Evaluate the server's software status 3. Evaluate the server's components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \"Incompatible-Server\" - the server model is not listed in the HCL. \"Incompatible-Processor\" - the server model and processor combination is not listed in HCL. \"Incompatible-Firmware\" - the server model, processor and server firmware is not listed in HCL. \"Missing-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Components\" - there is one or more components with \"Not-Validated\" status \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). \"Compatible\" - the server and all its components are validated. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information.\n* `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL.\n* `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model.\n* `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination.\n* `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version.\n* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.\n* `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result.\n* `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.\n* `Incompatible-Components` - The validation has failed for this server because one or more components have \"Not-Listed\" status.\n* `Compatible` - The validation has passed for this server's model, processor, OS vendor and version.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"software_status": {
						Description: "The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \"Validated\" - The os vendor/version is listed in the HCL for the server model, processor and firmware \"Not-Listed\" - The os vendor/version is not listed in the HCL for the server model, processor and firmware \"Incomplete\" - The inventory is missing os vendor/version and HCL validation was not performed. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "The HCL compatibility status of the managed object. The status can be one of the following \"Incomplete\" - there is no enough information to evaluate against the HCL data \"Validated\" - all components have been evaluated against the HCL and they all have \"Validated\" status \"Not-Listed\" - all components have been evaluated against the HCL and one or more have \"Not-Listed\" status. \"Not-Evaluated\" - server is not evaluated against the HCL because it is exempted.\n* `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.\n* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.\n* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.\n* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.\n* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.",
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

func dataSourceCondHclStatusRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CondHclStatus{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("component_status"); ok {
		x := (v.(string))
		o.SetComponentStatus(x)
	}
	if v, ok := d.GetOk("hardware_status"); ok {
		x := (v.(string))
		o.SetHardwareStatus(x)
	}
	if v, ok := d.GetOk("hcl_firmware_version"); ok {
		x := (v.(string))
		o.SetHclFirmwareVersion(x)
	}
	if v, ok := d.GetOk("hcl_model"); ok {
		x := (v.(string))
		o.SetHclModel(x)
	}
	if v, ok := d.GetOk("hcl_os_vendor"); ok {
		x := (v.(string))
		o.SetHclOsVendor(x)
	}
	if v, ok := d.GetOk("hcl_os_version"); ok {
		x := (v.(string))
		o.SetHclOsVersion(x)
	}
	if v, ok := d.GetOk("hcl_processor"); ok {
		x := (v.(string))
		o.SetHclProcessor(x)
	}
	if v, ok := d.GetOk("inv_firmware_version"); ok {
		x := (v.(string))
		o.SetInvFirmwareVersion(x)
	}
	if v, ok := d.GetOk("inv_model"); ok {
		x := (v.(string))
		o.SetInvModel(x)
	}
	if v, ok := d.GetOk("inv_os_vendor"); ok {
		x := (v.(string))
		o.SetInvOsVendor(x)
	}
	if v, ok := d.GetOk("inv_os_version"); ok {
		x := (v.(string))
		o.SetInvOsVersion(x)
	}
	if v, ok := d.GetOk("inv_processor"); ok {
		x := (v.(string))
		o.SetInvProcessor(x)
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
	if v, ok := d.GetOk("server_reason"); ok {
		x := (v.(string))
		o.SetServerReason(x)
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
		return diag.Errorf("json marshal of CondHclStatus object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CondApi.GetCondHclStatusList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CondHclStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CondHclStatusList.GetCount()
	var i int32
	var condHclStatusResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CondApi.GetCondHclStatusList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CondHclStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CondHclStatusList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CondHclStatus data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["component_status"] = (s.GetComponentStatus())

				temp["details"] = flattenListCondHclStatusDetailRelationship(s.GetDetails(), d)
				temp["hardware_status"] = (s.GetHardwareStatus())
				temp["hcl_firmware_version"] = (s.GetHclFirmwareVersion())
				temp["hcl_model"] = (s.GetHclModel())
				temp["hcl_os_vendor"] = (s.GetHclOsVendor())
				temp["hcl_os_version"] = (s.GetHclOsVersion())
				temp["hcl_processor"] = (s.GetHclProcessor())
				temp["inv_firmware_version"] = (s.GetInvFirmwareVersion())
				temp["inv_model"] = (s.GetInvModel())
				temp["inv_os_vendor"] = (s.GetInvOsVendor())
				temp["inv_os_version"] = (s.GetInvOsVersion())
				temp["inv_processor"] = (s.GetInvProcessor())

				temp["managed_object"] = flattenMapInventoryBaseRelationship(s.GetManagedObject(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["reason"] = (s.GetReason())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["server_reason"] = (s.GetServerReason())
				temp["software_status"] = (s.GetSoftwareStatus())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				condHclStatusResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(condHclStatusResults))
	if err := d.Set("results", condHclStatusResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(condHclStatusResults[0]["moid"].(string))
	return de
}
