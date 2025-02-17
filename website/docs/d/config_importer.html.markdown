---
subcategory: "config"
layout: "intersight"
page_title: "Intersight: intersight_config_importer"
description: |-
  All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
---

# Data Source: intersight_config_importer
All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_config_importer.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `import_path`:(string) The path to the archive in Intersight storage that has all the MOsto be imported. 
* `import_source`:(string) The source of the archive in Intersight storage that has all the MOsto be imported.* `ImageRepo` - The 'ImageRepo' source if the source of exporter archive is image repository.* `URL` - The 'URL' source if the source of exported archive is a URL. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An identifier for the importer instance. 
* `skip_integrity_checks`:(bool) Specifies whether integrity checks must be skipped. 
* `status`:(string) Status of the import operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been inititiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out. 
* `status_message`:(string) Status message associated with failures or progress indication. 
 