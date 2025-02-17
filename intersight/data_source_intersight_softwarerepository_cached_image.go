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

func dataSourceSoftwarerepositoryCachedImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryCachedImageRead,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in endpoint. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.\n* `None` - No action should be taken on the imported file.\n* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.\n* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.\n* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.\n* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.\n* `PreCache` - Cache the file into the Intersight Appliance.\n* `Cancel` - The cancel import process for the file into the repository.\n* `Extract` - The action to extract the file in the external repository.\n* `Evict` - Evict the cached file from the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cache_state": {
				Description: "The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process.\n* `ReadyForImport` - The image is ready to be imported into the repository.\n* `Importing` - The image is being imported into the repository.\n* `Imported` - The image has been extracted and imported into the repository.\n* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.\n* `Extracting` - Indicates that the image is being extracted into the repository.\n* `Extracted` - Indicates that the image has been extracted into the repository.\n* `Failed` - The image import from an external source to the repository has failed.\n* `MetaOnly` - The image is present in an external repository.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.\n* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.\n* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.\n* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.\n* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cached_time": {
				Description: "The time when the image or file was cached into the FI storage.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_error": {
				Description: "Any error encountered. Set to empty when download is in progress or completed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_progress": {
				Description: "The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_retries": {
				Description: "The number of retries the plugin attempted before succeeding or failing the download.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"md5sum": {
				Description: "The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image.",
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
			"original_sha512sum": {
				Description: "The actual sha512sum of the cached image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"path": {
				Description: "The absolute path of the imported file in the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"used_count": {
				Description: "The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"action": {
					Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in endpoint. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.\n* `None` - No action should be taken on the imported file.\n* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.\n* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.\n* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.\n* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.\n* `PreCache` - Cache the file into the Intersight Appliance.\n* `Cancel` - The cancel import process for the file into the repository.\n* `Extract` - The action to extract the file in the external repository.\n* `Evict` - Evict the cached file from the Intersight Appliance.",
					Type:        schema.TypeString,
					Optional:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"cache_state": {
						Description: "The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process.\n* `ReadyForImport` - The image is ready to be imported into the repository.\n* `Importing` - The image is being imported into the repository.\n* `Imported` - The image has been extracted and imported into the repository.\n* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.\n* `Extracting` - Indicates that the image is being extracted into the repository.\n* `Extracted` - Indicates that the image has been extracted into the repository.\n* `Failed` - The image import from an external source to the repository has failed.\n* `MetaOnly` - The image is present in an external repository.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.\n* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.\n* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.\n* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.\n* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"cached_time": {
						Description: "The time when the image or file was cached into the FI storage.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"checksum": {
						Description: "The checksum of the downloaded file as calculated by the download plugin after successfully downloading a file.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
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
								"hash_algorithm": {
									Description: "The hash algorithm used to calculate the checksum.\n* `crc` - A CRC hash as definded by RFC 3385. Generated with the IEEE polynomial.\n* `sha256` - A SHA256 hash as defined by RFC 4634.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"download_error": {
						Description: "Any error encountered. Set to empty when download is in progress or completed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"download_progress": {
						Description: "The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"download_retries": {
						Description: "The number of retries the plugin attempted before succeeding or failing the download.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file": {
						Description: "A reference to a softwarerepositoryFile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"md5sum": {
						Description: "The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image.",
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
					"network_element": {
						Description: "A reference to a networkElement resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"original_sha512sum": {
						Description: "The actual sha512sum of the cached image.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"path": {
						Description: "The absolute path of the imported file in the endpoint.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"registered_workflows": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
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
					"used_count": {
						Description: "The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceSoftwarerepositoryCachedImageRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryCachedImage{}
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}
	if v, ok := d.GetOk("cache_state"); ok {
		x := (v.(string))
		o.SetCacheState(x)
	}
	if v, ok := d.GetOk("cached_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCachedTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("download_error"); ok {
		x := (v.(string))
		o.SetDownloadError(x)
	}
	if v, ok := d.GetOk("download_progress"); ok {
		x := int64(v.(int))
		o.SetDownloadProgress(x)
	}
	if v, ok := d.GetOk("download_retries"); ok {
		x := int64(v.(int))
		o.SetDownloadRetries(x)
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("original_sha512sum"); ok {
		x := (v.(string))
		o.SetOriginalSha512sum(x)
	}
	if v, ok := d.GetOk("path"); ok {
		x := (v.(string))
		o.SetPath(x)
	}
	if v, ok := d.GetOk("used_count"); ok {
		x := int64(v.(int))
		o.SetUsedCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryCachedImage object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCachedImageList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SoftwarerepositoryCachedImage: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SoftwarerepositoryCachedImageList.GetCount()
	var i int32
	var softwarerepositoryCachedImageResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCachedImageList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SoftwarerepositoryCachedImage: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SoftwarerepositoryCachedImageList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SoftwarerepositoryCachedImage data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["cache_state"] = (s.GetCacheState())

				temp["cached_time"] = (s.GetCachedTime()).String()

				temp["checksum"] = flattenMapConnectorFileChecksum(s.GetChecksum(), d)
				temp["class_id"] = (s.GetClassId())
				temp["download_error"] = (s.GetDownloadError())
				temp["download_progress"] = (s.GetDownloadProgress())
				temp["download_retries"] = (s.GetDownloadRetries())

				temp["file"] = flattenMapSoftwarerepositoryFileRelationship(s.GetFile(), d)
				temp["md5sum"] = (s.GetMd5sum())
				temp["moid"] = (s.GetMoid())

				temp["network_element"] = flattenMapNetworkElementRelationship(s.GetNetworkElement(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["original_sha512sum"] = (s.GetOriginalSha512sum())
				temp["path"] = (s.GetPath())
				temp["registered_workflows"] = (s.GetRegisteredWorkflows())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["used_count"] = (s.GetUsedCount())
				softwarerepositoryCachedImageResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(softwarerepositoryCachedImageResults))
	if err := d.Set("results", softwarerepositoryCachedImageResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(softwarerepositoryCachedImageResults[0]["moid"].(string))
	return de
}
