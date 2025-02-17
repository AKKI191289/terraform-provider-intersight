package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssetDeviceContractInformation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAssetDeviceContractInformationRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"contract_status": {
				Description: "Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.\n* `Unknown` - The device's contract status cannot be determined.\n* `Not Covered` - The Cisco device does not have a valid support contract.\n* `Active` - The Cisco device is covered under a active support contract.\n* `Expiring Soon` - The contract for this Cisco device is going to expire in the next 30 days.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"contract_status_reason": {
				Description: "Reason for contract status. In case of Not Covered, reason is either Terminated or Expired.\n* `` - There is no reason for the specified contract status.\n* `Line Item Expired` - The Cisco device does not have a valid support contract, it has expired.\n* `Line Item Terminated` - The Cisco device does not have a valid support contract, it has been terminated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"contract_updated_time": {
				Description: "Date and time indicating when the contract data is last fetched from Cisco's Contract API successfully.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"covered_product_line_end_date": {
				Description: "End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_type": {
				Description: "Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.\n* `None` - A default value to catch cases where device type is not correctly detected.\n* `CiscoUcsServer` - A device of type server. It includes Cisco IMC and UCS Managed servers.\n* `CiscoUcsFI` - A device of type Fabric Interconnect. It includes the various types of Cisco Fabric Interconnects supported by Cisco Intersight.\n* `CiscoUcsChassis` - A device of type Chassis. It includes various UCS chassis supported by Cisco Intersight.\n* `CiscoNexusSwitch` - A device of type Nexus switch. It includes various Nexus switches supported by Cisco Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_valid": {
				Description: "Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"item_type": {
				Description: "Item type of this specific Cisco device. example \"Chassis\".",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_purchase_order_number": {
				Description: "Maintenance purchase order number for the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_sales_order_number": {
				Description: "Maintenance sales order number for the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"platform_type": {
				Description: "The platform type of the Cisco device.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.\n* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.\n* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"purchase_order_number": {
				Description: "Purchase order number for the Cisco device. It is a unique number assigned for every purchase.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sales_order_number": {
				Description: "Sales order number for the Cisco device. It is a unique number assigned for every sale.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_description": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_end_date": {
				Description: "End date for the Cisco service contract that covers this Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_level": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_sku": {
				Description: "The SKU of the service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_start_date": {
				Description: "Start date for the Cisco service contract that covers this Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state_contract": {
				Description: "Internal property used for triggering and tracking actions for contract information.\n* `Update` - Sn2Info/Contract information needs to be updated.\n* `OK` - Sn2Info/Contract information was fetched succcessfuly and updated.\n* `Failed` - Sn2Info/Contract information was not available  or failed while fetching.\n* `Retry` - Sn2Info/Contract information update failed and will be retried later.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"warranty_end_date": {
				Description: "End date for the warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"warranty_type": {
				Description: "Type of warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
					"contract": {
						Description: "Contract information for the Cisco support contract purchased for the Cisco device.",
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
								"bill_to": {
									Description: "BillTo address of listed for the contract.",
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
											"address1": {
												Description: "Address Line one of the address information. example \"PO BOX 641570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address2": {
												Description: "Address Line two of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address3": {
												Description: "Address Line three of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"city": {
												Description: "City in which the address resides. example \"San Jose\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"country": {
												Description: "Country in which the address resides. example \"US\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"county": {
												Description: "County in which the address resides. example \"Washington County\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"location": {
												Description: "Location in which the address resides. example \"14852\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the user whose address is being populated.",
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
											"postal_code": {
												Description: "Postal Code in which the address resides. example \"95164-1570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"province": {
												Description: "Province in which the address resides. example \"AB\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"state": {
												Description: "State in which the address resides. example \"CA\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"bill_to_global_ultimate": {
									Description: "BillToGlobalUltimate information listed in the contract.",
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
											"id": {
												Description: "ID of the user in BillToGlobal.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the user in BillToGlobal.",
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
										},
									},
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"contract_number": {
									Description: "Contract number for the Cisco support contract purchased for the Cisco device.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"line_status": {
									Description: "Contract status as per the Cisco Contract APIx.",
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
							},
						},
					},
					"contract_status": {
						Description: "Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.\n* `Unknown` - The device's contract status cannot be determined.\n* `Not Covered` - The Cisco device does not have a valid support contract.\n* `Active` - The Cisco device is covered under a active support contract.\n* `Expiring Soon` - The contract for this Cisco device is going to expire in the next 30 days.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"contract_status_reason": {
						Description: "Reason for contract status. In case of Not Covered, reason is either Terminated or Expired.\n* `` - There is no reason for the specified contract status.\n* `Line Item Expired` - The Cisco device does not have a valid support contract, it has expired.\n* `Line Item Terminated` - The Cisco device does not have a valid support contract, it has been terminated.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"contract_updated_time": {
						Description: "Date and time indicating when the contract data is last fetched from Cisco's Contract API successfully.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"covered_product_line_end_date": {
						Description: "End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_id": {
						Description: "Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_type": {
						Description: "Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.\n* `None` - A default value to catch cases where device type is not correctly detected.\n* `CiscoUcsServer` - A device of type server. It includes Cisco IMC and UCS Managed servers.\n* `CiscoUcsFI` - A device of type Fabric Interconnect. It includes the various types of Cisco Fabric Interconnects supported by Cisco Intersight.\n* `CiscoUcsChassis` - A device of type Chassis. It includes various UCS chassis supported by Cisco Intersight.\n* `CiscoNexusSwitch` - A device of type Nexus switch. It includes various Nexus switches supported by Cisco Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"end_customer": {
						Description: "End customer information for the Cisco support contract purchased for the Cisco device.",
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
								"address": {
									Description: "Address as per the information provided by the user.",
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
											"address1": {
												Description: "Address Line one of the address information. example \"PO BOX 641570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address2": {
												Description: "Address Line two of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address3": {
												Description: "Address Line three of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"city": {
												Description: "City in which the address resides. example \"San Jose\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"country": {
												Description: "Country in which the address resides. example \"US\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"county": {
												Description: "County in which the address resides. example \"Washington County\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"location": {
												Description: "Location in which the address resides. example \"14852\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the user whose address is being populated.",
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
											"postal_code": {
												Description: "Postal Code in which the address resides. example \"95164-1570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"province": {
												Description: "Province in which the address resides. example \"AB\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"state": {
												Description: "State in which the address resides. example \"CA\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"id": {
									Description: "Unique identifier for an end customer. This identifier is allocated by Cisco.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name as per the information provided by the user.",
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
							},
						},
					},
					"end_user_global_ultimate": {
						Description: "EndUserGlobalUltimate information listed in the contract.",
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
								"id": {
									Description: "ID of the user in BillToGlobal.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the user in BillToGlobal.",
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
							},
						},
					},
					"is_valid": {
						Description: "Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"item_type": {
						Description: "Item type of this specific Cisco device. example \"Chassis\".",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"maintenance_purchase_order_number": {
						Description: "Maintenance purchase order number for the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"maintenance_sales_order_number": {
						Description: "Maintenance sales order number for the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"platform_type": {
						Description: "The platform type of the Cisco device.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.\n* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.\n* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"product": {
						Description: "Product information of the offering under a contract.",
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
								"bill_to": {
									Description: "Billing address provided by customer while buying this Cisco product.",
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
											"address1": {
												Description: "Address Line one of the address information. example \"PO BOX 641570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address2": {
												Description: "Address Line two of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address3": {
												Description: "Address Line three of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"city": {
												Description: "City in which the address resides. example \"San Jose\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"country": {
												Description: "Country in which the address resides. example \"US\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"county": {
												Description: "County in which the address resides. example \"Washington County\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"location": {
												Description: "Location in which the address resides. example \"14852\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the user whose address is being populated.",
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
											"postal_code": {
												Description: "Postal Code in which the address resides. example \"95164-1570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"province": {
												Description: "Province in which the address resides. example \"AB\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"state": {
												Description: "State in which the address resides. example \"CA\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"description": {
									Description: "Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"family": {
									Description: "Family that the product belongs to. Example \"UCSB\".",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"group": {
									Description: "Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"number": {
									Description: "Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".",
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
								"ship_to": {
									Description: "Shipping address provided by customer while buying this Cisco product.",
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
											"address1": {
												Description: "Address Line one of the address information. example \"PO BOX 641570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address2": {
												Description: "Address Line two of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"address3": {
												Description: "Address Line three of the address information. example \"Cisco Systems\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"city": {
												Description: "City in which the address resides. example \"San Jose\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"country": {
												Description: "Country in which the address resides. example \"US\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"county": {
												Description: "County in which the address resides. example \"Washington County\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"location": {
												Description: "Location in which the address resides. example \"14852\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the user whose address is being populated.",
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
											"postal_code": {
												Description: "Postal Code in which the address resides. example \"95164-1570\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"province": {
												Description: "Province in which the address resides. example \"AB\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"state": {
												Description: "State in which the address resides. example \"CA\".",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"sub_type": {
									Description: "Sub type of the product being specified. example \"UCS 6200 SER\".",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"purchase_order_number": {
						Description: "Purchase order number for the Cisco device. It is a unique number assigned for every purchase.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"reseller_global_ultimate": {
						Description: "ResellerGlobalUltimate information listed in the contract.",
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
								"id": {
									Description: "ID of the user in BillToGlobal.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the user in BillToGlobal.",
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
							},
						},
					},
					"sales_order_number": {
						Description: "Sales order number for the Cisco device. It is a unique number assigned for every sale.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"service_description": {
						Description: "The type of service contract that covers the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"service_end_date": {
						Description: "End date for the Cisco service contract that covers this Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"service_level": {
						Description: "The type of service contract that covers the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"service_sku": {
						Description: "The SKU of the service contract that covers the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"service_start_date": {
						Description: "Start date for the Cisco service contract that covers this Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"state_contract": {
						Description: "Internal property used for triggering and tracking actions for contract information.\n* `Update` - Sn2Info/Contract information needs to be updated.\n* `OK` - Sn2Info/Contract information was fetched succcessfuly and updated.\n* `Failed` - Sn2Info/Contract information was not available  or failed while fetching.\n* `Retry` - Sn2Info/Contract information update failed and will be retried later.",
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
					"warranty_end_date": {
						Description: "End date for the warranty that covers the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"warranty_type": {
						Description: "Type of warranty that covers the Cisco device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceAssetDeviceContractInformationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.AssetDeviceContractInformation{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("contract_status"); ok {
		x := (v.(string))
		o.SetContractStatus(x)
	}
	if v, ok := d.GetOk("contract_status_reason"); ok {
		x := (v.(string))
		o.SetContractStatusReason(x)
	}
	if v, ok := d.GetOk("contract_updated_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetContractUpdatedTime(x)
	}
	if v, ok := d.GetOk("covered_product_line_end_date"); ok {
		x := (v.(string))
		o.SetCoveredProductLineEndDate(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
	}
	if v, ok := d.GetOk("is_valid"); ok {
		x := (v.(bool))
		o.SetIsValid(x)
	}
	if v, ok := d.GetOk("item_type"); ok {
		x := (v.(string))
		o.SetItemType(x)
	}
	if v, ok := d.GetOk("maintenance_purchase_order_number"); ok {
		x := (v.(string))
		o.SetMaintenancePurchaseOrderNumber(x)
	}
	if v, ok := d.GetOk("maintenance_sales_order_number"); ok {
		x := (v.(string))
		o.SetMaintenanceSalesOrderNumber(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("purchase_order_number"); ok {
		x := (v.(string))
		o.SetPurchaseOrderNumber(x)
	}
	if v, ok := d.GetOk("sales_order_number"); ok {
		x := (v.(string))
		o.SetSalesOrderNumber(x)
	}
	if v, ok := d.GetOk("service_description"); ok {
		x := (v.(string))
		o.SetServiceDescription(x)
	}
	if v, ok := d.GetOk("service_end_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetServiceEndDate(x)
	}
	if v, ok := d.GetOk("service_level"); ok {
		x := (v.(string))
		o.SetServiceLevel(x)
	}
	if v, ok := d.GetOk("service_sku"); ok {
		x := (v.(string))
		o.SetServiceSku(x)
	}
	if v, ok := d.GetOk("service_start_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetServiceStartDate(x)
	}
	if v, ok := d.GetOk("state_contract"); ok {
		x := (v.(string))
		o.SetStateContract(x)
	}
	if v, ok := d.GetOk("warranty_end_date"); ok {
		x := (v.(string))
		o.SetWarrantyEndDate(x)
	}
	if v, ok := d.GetOk("warranty_type"); ok {
		x := (v.(string))
		o.SetWarrantyType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of AssetDeviceContractInformation object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.AssetApi.GetAssetDeviceContractInformationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of AssetDeviceContractInformation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.AssetDeviceContractInformationList.GetCount()
	var i int32
	var assetDeviceContractInformationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.AssetApi.GetAssetDeviceContractInformationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching AssetDeviceContractInformation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.AssetDeviceContractInformationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for AssetDeviceContractInformation data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["contract"] = flattenMapAssetContractInformation(s.GetContract(), d)
				temp["contract_status"] = (s.GetContractStatus())
				temp["contract_status_reason"] = (s.GetContractStatusReason())

				temp["contract_updated_time"] = (s.GetContractUpdatedTime()).String()
				temp["covered_product_line_end_date"] = (s.GetCoveredProductLineEndDate())
				temp["device_id"] = (s.GetDeviceId())
				temp["device_type"] = (s.GetDeviceType())

				temp["end_customer"] = flattenMapAssetCustomerInformation(s.GetEndCustomer(), d)

				temp["end_user_global_ultimate"] = flattenMapAssetGlobalUltimate(s.GetEndUserGlobalUltimate(), d)
				temp["is_valid"] = (s.GetIsValid())
				temp["item_type"] = (s.GetItemType())
				temp["maintenance_purchase_order_number"] = (s.GetMaintenancePurchaseOrderNumber())
				temp["maintenance_sales_order_number"] = (s.GetMaintenanceSalesOrderNumber())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["platform_type"] = (s.GetPlatformType())

				temp["product"] = flattenMapAssetProductInformation(s.GetProduct(), d)
				temp["purchase_order_number"] = (s.GetPurchaseOrderNumber())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["reseller_global_ultimate"] = flattenMapAssetGlobalUltimate(s.GetResellerGlobalUltimate(), d)
				temp["sales_order_number"] = (s.GetSalesOrderNumber())
				temp["service_description"] = (s.GetServiceDescription())

				temp["service_end_date"] = (s.GetServiceEndDate()).String()
				temp["service_level"] = (s.GetServiceLevel())
				temp["service_sku"] = (s.GetServiceSku())

				temp["service_start_date"] = (s.GetServiceStartDate()).String()
				temp["state_contract"] = (s.GetStateContract())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["warranty_end_date"] = (s.GetWarrantyEndDate())
				temp["warranty_type"] = (s.GetWarrantyType())
				assetDeviceContractInformationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(assetDeviceContractInformationResults))
	if err := d.Set("results", assetDeviceContractInformationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(assetDeviceContractInformationResults[0]["moid"].(string))
	return de
}
