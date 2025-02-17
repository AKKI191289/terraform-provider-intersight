---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_deployment"
description: |-
  Deployment inventory represents a Kubernetes Deployment. A Kubernetes Deployment provides declarative for Pods and ReplicaSets. For the Deployment in Intersight, it is an read only object that represent the Deployment. In Kubernetes world, you can describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.
---

# Data Source: intersight_kubernetes_deployment
Deployment inventory represents a Kubernetes Deployment. A Kubernetes Deployment provides declarative for Pods and ReplicaSets. For the Deployment in Intersight, it is an read only object that represent the Deployment. In Kubernetes world, you can describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_deployment.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the referenced kubernetes resource. 
* `uuid`:(string) UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself. 
 