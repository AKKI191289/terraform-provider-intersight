---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_aci_cni_apic"
description: |-
  Internally generated object of claimed APIC device known to Razor.
---

# Data Source: intersight_kubernetes_aci_cni_apic
Internally generated object of claimed APIC device known to Razor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_aci_cni_apic.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `asset_apic_moid`:(string) The Moid of the apic device under the asset service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_aci_cni_profiles`:(int) The number of ACI CNI profiles configured for this APIC. 
 