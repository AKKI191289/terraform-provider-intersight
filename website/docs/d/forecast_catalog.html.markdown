---
subcategory: "forecast"
layout: "intersight"
page_title: "Intersight: intersight_forecast_catalog"
description: |-
  A catalog for managing forecast settings.
---

# Data Source: intersight_forecast_catalog
A catalog for managing forecast settings.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_forecast_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `sched_time`:(string) The time at which the regression model needs to run for all the metrics specified in catalog. 
* `nr_version`:(string) The catalog version used in forecast configuration service. 
 