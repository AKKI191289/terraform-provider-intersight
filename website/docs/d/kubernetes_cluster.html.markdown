---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_cluster"
description: |-
  Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.
---

# Data Source: intersight_kubernetes_cluster
Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_cluster.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `connection_status`:(string) Status of the endpoint connection of this Kubernetes cluster.* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.* `NotConnected` - Intersight is unable to establish a connection to the target.* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. 
* `kube_config`:(string) Kubeconfig for the cluster to collect inventory for. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Kubernetes cluster. 
 