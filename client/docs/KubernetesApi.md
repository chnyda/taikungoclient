# \KubernetesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesClearCrds**](KubernetesApi.md#KubernetesClearCrds) | **Post** /api/v{v}/Kubernetes/remove-crds | Remove crds
[**KubernetesCreateKubernetesAlert**](KubernetesApi.md#KubernetesCreateKubernetesAlert) | **Post** /api/v{v}/Kubernetes/alert/{projectId} | Add k8s alert
[**KubernetesCreateKubernetesEvent**](KubernetesApi.md#KubernetesCreateKubernetesEvent) | **Post** /api/v{v}/Kubernetes/event/{projectId} | Add k8s event - opsgenie
[**KubernetesDeleteAlerts**](KubernetesApi.md#KubernetesDeleteAlerts) | **Post** /api/v{v}/Kubernetes/removealerts | Remove k8s alerts
[**KubernetesDeletePod**](KubernetesApi.md#KubernetesDeletePod) | **Post** /api/v{v}/Kubernetes/{projectId}/deletepod/{metadataName}/{podNamespace} | Remove k8s namespaced pod
[**KubernetesDescribeConfigMap**](KubernetesApi.md#KubernetesDescribeConfigMap) | **Post** /api/v{v}/Kubernetes/describe/configmap | Describe config map
[**KubernetesDescribeCrd**](KubernetesApi.md#KubernetesDescribeCrd) | **Post** /api/v{v}/Kubernetes/describe/crd | Describe crd
[**KubernetesDescribeCronJob**](KubernetesApi.md#KubernetesDescribeCronJob) | **Post** /api/v{v}/Kubernetes/describe/cronjob | Describe cronjob
[**KubernetesDescribeDaemonSet**](KubernetesApi.md#KubernetesDescribeDaemonSet) | **Post** /api/v{v}/Kubernetes/describe/daemonset | Describe daemonset
[**KubernetesDescribeDeployment**](KubernetesApi.md#KubernetesDescribeDeployment) | **Post** /api/v{v}/Kubernetes/describe/deployment | Describe deployment
[**KubernetesDescribeIngress**](KubernetesApi.md#KubernetesDescribeIngress) | **Post** /api/v{v}/Kubernetes/describe/ingress | Describe ingress
[**KubernetesDescribeJob**](KubernetesApi.md#KubernetesDescribeJob) | **Post** /api/v{v}/Kubernetes/describe/job | Describe job
[**KubernetesDescribeNetworkPolicy**](KubernetesApi.md#KubernetesDescribeNetworkPolicy) | **Post** /api/v{v}/Kubernetes/describe/network-policy | Describe network policy
[**KubernetesDescribeNode**](KubernetesApi.md#KubernetesDescribeNode) | **Post** /api/v{v}/Kubernetes/describe/node | Describe node
[**KubernetesDescribePdb**](KubernetesApi.md#KubernetesDescribePdb) | **Post** /api/v{v}/Kubernetes/describe/pdb | Describe pdb
[**KubernetesDescribePod**](KubernetesApi.md#KubernetesDescribePod) | **Post** /api/v{v}/Kubernetes/describe/pod | Describe pod
[**KubernetesDescribePvc**](KubernetesApi.md#KubernetesDescribePvc) | **Post** /api/v{v}/Kubernetes/describe/pvc | Describe pvc
[**KubernetesDescribeSecret**](KubernetesApi.md#KubernetesDescribeSecret) | **Post** /api/v{v}/Kubernetes/describe/secret | Describe secret
[**KubernetesDescribeService**](KubernetesApi.md#KubernetesDescribeService) | **Post** /api/v{v}/Kubernetes/describe/service | Describe service
[**KubernetesDescribeStorageClass**](KubernetesApi.md#KubernetesDescribeStorageClass) | **Post** /api/v{v}/Kubernetes/describe/storageclass | Describe storage class
[**KubernetesDescribeSts**](KubernetesApi.md#KubernetesDescribeSts) | **Post** /api/v{v}/Kubernetes/describe/sts | Describe stateful set
[**KubernetesDownload**](KubernetesApi.md#KubernetesDownload) | **Get** /api/v{v}/Kubernetes/{projectId}/download | Download kube config file bu project Id
[**KubernetesExecCli**](KubernetesApi.md#KubernetesExecCli) | **Post** /api/v{v}/Kubernetes/cli | Execute k8s web socket namespaced pod
[**KubernetesGetConfigMap**](KubernetesApi.md#KubernetesGetConfigMap) | **Get** /api/v{v}/Kubernetes/{projectId}/configmap | Retrieve a list of k8s config map for all namespaces
[**KubernetesGetCrdList**](KubernetesApi.md#KubernetesGetCrdList) | **Get** /api/v{v}/Kubernetes/{projectId}/crd | Retrieve a list of k8s custom resource definition
[**KubernetesGetCronJobsList**](KubernetesApi.md#KubernetesGetCronJobsList) | **Get** /api/v{v}/Kubernetes/{projectId}/cronjobs | Retrieve a list of k8s cron jobs for all namespaces
[**KubernetesGetDaemonSet**](KubernetesApi.md#KubernetesGetDaemonSet) | **Get** /api/v{v}/Kubernetes/{projectId}/daemonset | Retrieve a list of k8s daemon set for all namespaces
[**KubernetesGetDashboard**](KubernetesApi.md#KubernetesGetDashboard) | **Get** /api/v{v}/Kubernetes/{projectId}/dashboard | Retrieve all data of the k8s by project Id
[**KubernetesGetDeploymentList**](KubernetesApi.md#KubernetesGetDeploymentList) | **Get** /api/v{v}/Kubernetes/{projectId}/deployment | Retrieve a list of k8s deployment for all namespaces
[**KubernetesGetHelmReleaseList**](KubernetesApi.md#KubernetesGetHelmReleaseList) | **Get** /api/v{v}/Kubernetes/{projectId}/helmreleases | Retrieve a list of k8s helm releases for all namespaces
[**KubernetesGetIngresses**](KubernetesApi.md#KubernetesGetIngresses) | **Get** /api/v{v}/Kubernetes/{projectId}/ingress | Retrieve a list of k8s ingresses for all namespaces
[**KubernetesGetJobsList**](KubernetesApi.md#KubernetesGetJobsList) | **Get** /api/v{v}/Kubernetes/{projectId}/jobs | Retrieve a list of k8s jobs for all namespaces
[**KubernetesGetKubeConfigFile**](KubernetesApi.md#KubernetesGetKubeConfigFile) | **Get** /api/v{v}/Kubernetes/{projectId}/kubeconfig | Retrieve kube config file
[**KubernetesGetKubernetesAlertsList**](KubernetesApi.md#KubernetesGetKubernetesAlertsList) | **Get** /api/v{v}/Kubernetes/{projectId}/alerts | Retrieve all k8s alerts by project Id
[**KubernetesGetNamespacesList**](KubernetesApi.md#KubernetesGetNamespacesList) | **Get** /api/v{v}/Kubernetes/{projectId}/namespaces | Retrieve kube config file
[**KubernetesGetNetworkPolicies**](KubernetesApi.md#KubernetesGetNetworkPolicies) | **Get** /api/v{v}/Kubernetes/{projectId}/network-policies | Retrieve a list of k8s network-policies for all namespaces
[**KubernetesGetNodeList**](KubernetesApi.md#KubernetesGetNodeList) | **Get** /api/v{v}/Kubernetes/{projectId}/node | Retrieve a list of k8s node
[**KubernetesGetPdbList**](KubernetesApi.md#KubernetesGetPdbList) | **Get** /api/v{v}/Kubernetes/{projectId}/pdb | Retrieve a list of k8s pdb for all namespaces
[**KubernetesGetPersistentVolumeClaim**](KubernetesApi.md#KubernetesGetPersistentVolumeClaim) | **Get** /api/v{v}/Kubernetes/{projectId}/pvc | Retrieve a list of k8s persistent volume claim for all namespaces
[**KubernetesGetPodList**](KubernetesApi.md#KubernetesGetPodList) | **Get** /api/v{v}/Kubernetes/{projectId}/pod | Retrieve a list of k8s pod for all namespaces
[**KubernetesGetPodLogsList**](KubernetesApi.md#KubernetesGetPodLogsList) | **Post** /api/v{v}/Kubernetes/podlogs | Retrieve k8s pod logs
[**KubernetesGetQuota**](KubernetesApi.md#KubernetesGetQuota) | **Get** /api/v{v}/Kubernetes/quota/{projectId} | k8s quota usage
[**KubernetesGetSecret**](KubernetesApi.md#KubernetesGetSecret) | **Get** /api/v{v}/Kubernetes/{projectId}/secret | Retrieve a list of k8s secret for all namespaces
[**KubernetesGetService**](KubernetesApi.md#KubernetesGetService) | **Get** /api/v{v}/Kubernetes/{projectId}/service | Retrieve a list of k8s service for all namespaces
[**KubernetesGetStorageClasses**](KubernetesApi.md#KubernetesGetStorageClasses) | **Get** /api/v{v}/Kubernetes/{projectId}/storageclass | Retrieve a list of k8s storage classes
[**KubernetesGetSts**](KubernetesApi.md#KubernetesGetSts) | **Get** /api/v{v}/Kubernetes/{projectId}/sts | Retrieve a list of k8s stateful set for all namespaces
[**KubernetesGetSupportedList**](KubernetesApi.md#KubernetesGetSupportedList) | **Get** /api/v{v}/Kubernetes/supported/list | Retrieve Taikun supported kubernetes version
[**KubernetesInteractiveShell**](KubernetesApi.md#KubernetesInteractiveShell) | **Post** /api/v{v}/Kubernetes/interactive-shell | Produce interactive shell command
[**KubernetesOverview**](KubernetesApi.md#KubernetesOverview) | **Get** /api/v{v}/Kubernetes/overview | Overview kubernetes nodes and pods by organization id
[**KubernetesPatchConfigMap**](KubernetesApi.md#KubernetesPatchConfigMap) | **Post** /api/v{v}/Kubernetes/patch/configmap | Patch config map
[**KubernetesPatchCrd**](KubernetesApi.md#KubernetesPatchCrd) | **Post** /api/v{v}/Kubernetes/patch/crd | Patch crd
[**KubernetesPatchCronJob**](KubernetesApi.md#KubernetesPatchCronJob) | **Post** /api/v{v}/Kubernetes/patch/cronjob | Patch cronjob
[**KubernetesPatchDeployment**](KubernetesApi.md#KubernetesPatchDeployment) | **Post** /api/v{v}/Kubernetes/patch/deployment | Patch deployment
[**KubernetesPatchIngress**](KubernetesApi.md#KubernetesPatchIngress) | **Post** /api/v{v}/Kubernetes/patch/ingress | Patch ingress
[**KubernetesPatchJob**](KubernetesApi.md#KubernetesPatchJob) | **Post** /api/v{v}/Kubernetes/patch/job | Patch job
[**KubernetesPatchNode**](KubernetesApi.md#KubernetesPatchNode) | **Post** /api/v{v}/Kubernetes/patch/node | Patch node
[**KubernetesPatchPdb**](KubernetesApi.md#KubernetesPatchPdb) | **Post** /api/v{v}/Kubernetes/patch/pdb | Patch pdb
[**KubernetesPatchPod**](KubernetesApi.md#KubernetesPatchPod) | **Post** /api/v{v}/Kubernetes/patch/pod | Patch pod
[**KubernetesPatchPvc**](KubernetesApi.md#KubernetesPatchPvc) | **Post** /api/v{v}/Kubernetes/patch/pvc | Patch pvc
[**KubernetesPatchSecret**](KubernetesApi.md#KubernetesPatchSecret) | **Post** /api/v{v}/Kubernetes/patch/secret | Patch secret
[**KubernetesPatchSts**](KubernetesApi.md#KubernetesPatchSts) | **Post** /api/v{v}/Kubernetes/patch/sts | Patch sts
[**KubernetesRestartDaemonSet**](KubernetesApi.md#KubernetesRestartDaemonSet) | **Post** /api/v{v}/Kubernetes/restart/daemonset | Restart daemon set
[**KubernetesRestartDeployment**](KubernetesApi.md#KubernetesRestartDeployment) | **Post** /api/v{v}/Kubernetes/restart/deployment | Restart deployment
[**KubernetesRestartSts**](KubernetesApi.md#KubernetesRestartSts) | **Post** /api/v{v}/Kubernetes/restart/sts | Restart stateful set
[**KubernetesSilenceManager**](KubernetesApi.md#KubernetesSilenceManager) | **Post** /api/v{v}/Kubernetes/silencemanager | Silence management for k8s alerts
[**KubernetesUpdateKubernetesAlert**](KubernetesApi.md#KubernetesUpdateKubernetesAlert) | **Put** /api/v{v}/Kubernetes/updatealert/{alertId} | Update k8s alert



## KubernetesClearCrds

> KubernetesClearCrds(ctx, v).Body(body).Execute()

Remove crds

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewClearCrdsCommand() // ClearCrdsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesClearCrds(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesClearCrds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClearCrdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClearCrdsCommand**](ClearCrdsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCreateKubernetesAlert

> KubernetesCreateKubernetesAlert(ctx, projectId, v).Body(body).Execute()

Add k8s alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    body := *openapiclient.NewCreateAlertDto() // CreateAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesCreateKubernetesAlert(context.Background(), projectId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesCreateKubernetesAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCreateKubernetesAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateAlertDto**](CreateAlertDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCreateKubernetesEvent

> KubernetesCreateKubernetesEvent(ctx, projectId, v).Body(body).Execute()

Add k8s event - opsgenie

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    body := *openapiclient.NewKubernetesEventCreateDto() // KubernetesEventCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesCreateKubernetesEvent(context.Background(), projectId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesCreateKubernetesEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCreateKubernetesEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**KubernetesEventCreateDto**](KubernetesEventCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeleteAlerts

> KubernetesDeleteAlerts(ctx, v).Body(body).Execute()

Remove k8s alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDeleteAlertCommand(int32(123)) // DeleteAlertCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesDeleteAlerts(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDeleteAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeleteAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteAlertCommand**](DeleteAlertCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeletePod

> KubernetesDeletePod(ctx, projectId, metadataName, podNamespace, v).Execute()

Remove k8s namespaced pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    metadataName := "metadataName_example" // string | 
    podNamespace := "podNamespace_example" // string | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesDeletePod(context.Background(), projectId, metadataName, podNamespace, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDeletePod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**metadataName** | **string** |  | 
**podNamespace** | **string** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeletePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeConfigMap

> string KubernetesDescribeConfigMap(ctx, v).Body(body).Execute()

Describe config map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeConfigMapCommand(int32(123), "Name_example", "Namespace_example") // DescribeConfigMapCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeConfigMap(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeConfigMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeConfigMap`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeConfigMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeConfigMapCommand**](DescribeConfigMapCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCrd

> string KubernetesDescribeCrd(ctx, v).Body(body).Execute()

Describe crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeCrdCommand(int32(123), "Name_example") // DescribeCrdCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeCrd(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCrd`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeCrd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeCrdCommand**](DescribeCrdCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCronJob

> string KubernetesDescribeCronJob(ctx, v).Body(body).Execute()

Describe cronjob

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeCronJobCommand(int32(123), "Name_example", "Namespace_example") // DescribeCronJobCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeCronJob(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeCronJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCronJob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeCronJobCommand**](DescribeCronJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDaemonSet

> string KubernetesDescribeDaemonSet(ctx, v).Body(body).Execute()

Describe daemonset

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeDaemonSetCommand(int32(123), "Name_example", "Namespace_example") // DescribeDaemonSetCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeDaemonSet(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDaemonSet`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeDaemonSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeDaemonSetCommand**](DescribeDaemonSetCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDeployment

> string KubernetesDescribeDeployment(ctx, v).Body(body).Execute()

Describe deployment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeDeploymentCommand(int32(123), "Name_example", "Namespace_example") // DescribeDeploymentCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeDeployment(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDeployment`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeDeploymentCommand**](DescribeDeploymentCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeIngress

> string KubernetesDescribeIngress(ctx, v).Body(body).Execute()

Describe ingress

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeIngressCommand(int32(123), "Name_example", "Namespace_example") // DescribeIngressCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeIngress(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeIngress`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeIngress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeIngressCommand**](DescribeIngressCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeJob

> string KubernetesDescribeJob(ctx, v).Body(body).Execute()

Describe job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeJobCommand(int32(123), "Name_example", "Namespace_example") // DescribeJobCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeJob(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeJob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeJobCommand**](DescribeJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNetworkPolicy

> string KubernetesDescribeNetworkPolicy(ctx, v).Body(body).Execute()

Describe network policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeNetworkPolicyCommand(int32(123), "Name_example", "Namespace_example") // DescribeNetworkPolicyCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeNetworkPolicy(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNetworkPolicy`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeNetworkPolicyCommand**](DescribeNetworkPolicyCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNode

> string KubernetesDescribeNode(ctx, v).Body(body).Execute()

Describe node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeNodeCommand(int32(123), "Name_example") // DescribeNodeCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeNode(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNode`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeNodeCommand**](DescribeNodeCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePdb

> string KubernetesDescribePdb(ctx, v).Body(body).Execute()

Describe pdb

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribePodDisruptionCommand(int32(123), "Name_example", "Namespace_example") // DescribePodDisruptionCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePdb(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePdb`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribePodDisruptionCommand**](DescribePodDisruptionCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePod

> string KubernetesDescribePod(ctx, v).Body(body).Execute()

Describe pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribePodCommand(int32(123), "Name_example", "Namespace_example") // DescribePodCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePod(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePod`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribePodCommand**](DescribePodCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePvc

> string KubernetesDescribePvc(ctx, v).Body(body).Execute()

Describe pvc

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribePvcCommand(int32(123), "Name_example", "Namespace_example") // DescribePvcCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePvc(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePvc`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePvc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribePvcCommand**](DescribePvcCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSecret

> string KubernetesDescribeSecret(ctx, v).Body(body).Execute()

Describe secret

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeSecretCommand(int32(123), "Name_example", "Namespace_example") // DescribeSecretCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeSecret(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSecret`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeSecretCommand**](DescribeSecretCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeService

> string KubernetesDescribeService(ctx, v).Body(body).Execute()

Describe service

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeServiceCommand(int32(123), "Name_example", "Namespace_example") // DescribeServiceCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeService(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeService`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeServiceCommand**](DescribeServiceCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeStorageClass

> string KubernetesDescribeStorageClass(ctx, v).Body(body).Execute()

Describe storage class

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeStorageClassCommand(int32(123), "Name_example") // DescribeStorageClassCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeStorageClass(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeStorageClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeStorageClass`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeStorageClassCommand**](DescribeStorageClassCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSts

> string KubernetesDescribeSts(ctx, v).Body(body).Execute()

Describe stateful set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewDescribeStsCommand(int32(123), "Name_example", "Namespace_example") // DescribeStsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeSts(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSts`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeSts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DescribeStsCommand**](DescribeStsCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDownload

> map[string]interface{} KubernetesDownload(ctx, projectId, v).Execute()

Download kube config file bu project Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDownload(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDownload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesExecCli

> string KubernetesExecCli(ctx, v).Body(body).Execute()

Execute k8s web socket namespaced pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewKubernetesCliCommand() // KubernetesCliCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesExecCli(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesExecCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesExecCli`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesExecCli`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesExecCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KubernetesCliCommand**](KubernetesCliCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetConfigMap

> ConfigMaps KubernetesGetConfigMap(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s config map for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetConfigMap(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetConfigMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetConfigMap`: ConfigMaps
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetConfigMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**ConfigMaps**](ConfigMaps.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetCrdList

> map[string]interface{} KubernetesGetCrdList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s custom resource definition

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetCrdList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetCrdList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetCrdList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetCrdList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetCrdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetCronJobsList

> KubernetesCronJobsList KubernetesGetCronJobsList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s cron jobs for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetCronJobsList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetCronJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetCronJobsList`: KubernetesCronJobsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetCronJobsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetCronJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesCronJobsList**](KubernetesCronJobsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetDaemonSet

> DaemonSets KubernetesGetDaemonSet(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s daemon set for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetDaemonSet(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetDaemonSet`: DaemonSets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetDaemonSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**DaemonSets**](DaemonSets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetDashboard

> string KubernetesGetDashboard(ctx, projectId, v).Execute()

Retrieve all data of the k8s by project Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetDashboard(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetDashboard`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetDeploymentList

> Deployments KubernetesGetDeploymentList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s deployment for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetDeploymentList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetDeploymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetDeploymentList`: Deployments
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetDeploymentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetDeploymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Deployments**](Deployments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetHelmReleaseList

> HelmReleasesList KubernetesGetHelmReleaseList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s helm releases for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetHelmReleaseList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetHelmReleaseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetHelmReleaseList`: HelmReleasesList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetHelmReleaseList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetHelmReleaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**HelmReleasesList**](HelmReleasesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetIngresses

> Ingresses KubernetesGetIngresses(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s ingresses for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetIngresses(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetIngresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetIngresses`: Ingresses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetIngresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetIngressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Ingresses**](Ingresses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetJobsList

> KubernetesJobList KubernetesGetJobsList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s jobs for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetJobsList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetJobsList`: KubernetesJobList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetJobsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesJobList**](KubernetesJobList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetKubeConfigFile

> KubeConfigResponse KubernetesGetKubeConfigFile(ctx, projectId, v).Execute()

Retrieve kube config file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetKubeConfigFile(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetKubeConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetKubeConfigFile`: KubeConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetKubeConfigFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetKubeConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KubeConfigResponse**](KubeConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetKubernetesAlertsList

> KubernetesAlertList KubernetesGetKubernetesAlertsList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()

Retrieve all k8s alerts by project Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetKubernetesAlertsList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetKubernetesAlertsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetKubernetesAlertsList`: KubernetesAlertList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetKubernetesAlertsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetKubernetesAlertsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

### Return type

[**KubernetesAlertList**](KubernetesAlertList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetNamespacesList

> []string KubernetesGetNamespacesList(ctx, projectId, v).Execute()

Retrieve kube config file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetNamespacesList(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetNamespacesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetNamespacesList`: []string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetNamespacesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetNamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetNetworkPolicies

> NetworkPolicies KubernetesGetNetworkPolicies(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s network-policies for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetNetworkPolicies(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetNetworkPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetNetworkPolicies`: NetworkPolicies
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetNetworkPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetNetworkPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**NetworkPolicies**](NetworkPolicies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetNodeList

> map[string]interface{} KubernetesGetNodeList(ctx, projectId, v).SearchId(searchId).Execute()

Retrieve a list of k8s node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetNodeList(context.Background(), projectId, v).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetNodeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetNodeList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetNodeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetPdbList

> PodDisruptions KubernetesGetPdbList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pdb for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetPdbList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetPdbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetPdbList`: PodDisruptions
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetPdbList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetPdbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**PodDisruptions**](PodDisruptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetPersistentVolumeClaim

> Pvcs KubernetesGetPersistentVolumeClaim(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s persistent volume claim for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetPersistentVolumeClaim(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetPersistentVolumeClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetPersistentVolumeClaim`: Pvcs
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pvcs**](Pvcs.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetPodList

> Pods KubernetesGetPodList(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pod for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetPodList(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetPodList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetPodList`: Pods
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetPodList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetPodListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pods**](Pods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetPodLogsList

> map[string]interface{} KubernetesGetPodLogsList(ctx, v).Body(body).Execute()

Retrieve k8s pod logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewKubernetesPodLogsCommand() // KubernetesPodLogsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetPodLogsList(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetPodLogsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetPodLogsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetPodLogsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetPodLogsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KubernetesPodLogsCommand**](KubernetesPodLogsCommand.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetQuota

> KubernetesQuotaListDto KubernetesGetQuota(ctx, projectId, v).Execute()

k8s quota usage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetQuota(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetQuota`: KubernetesQuotaListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KubernetesQuotaListDto**](KubernetesQuotaListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSecret

> Secrets KubernetesGetSecret(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s secret for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetSecret(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetSecret`: Secrets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Secrets**](Secrets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetService

> Services KubernetesGetService(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s service for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetService(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetService`: Services
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Services**](Services.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetStorageClasses

> StorageClasses KubernetesGetStorageClasses(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s storage classes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetStorageClasses(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetStorageClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetStorageClasses`: StorageClasses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetStorageClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StorageClasses**](StorageClasses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSts

> StsList KubernetesGetSts(ctx, projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s stateful set for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetSts(context.Background(), projectId, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetSts`: StsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetSts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StsList**](StsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSupportedList

> []KubernetesVersionListDto KubernetesGetSupportedList(ctx, v).Execute()

Retrieve Taikun supported kubernetes version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesGetSupportedList(context.Background(), v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetSupportedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetSupportedList`: []KubernetesVersionListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetSupportedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetSupportedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]KubernetesVersionListDto**](KubernetesVersionListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesInteractiveShell

> string KubernetesInteractiveShell(ctx, v).Body(body).Execute()

Produce interactive shell command

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewInteractiveShellSendCommand(int32(123), "Token_example") // InteractiveShellSendCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesInteractiveShell(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesInteractiveShell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesInteractiveShell`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesInteractiveShell`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InteractiveShellSendCommand**](InteractiveShellSendCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesOverview

> string KubernetesOverview(ctx, v).OrganizationId(organizationId).Execute()

Overview kubernetes nodes and pods by organization id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesOverview(context.Background(), v).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesOverview`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchConfigMap

> KubernetesPatchConfigMap(ctx, v).Body(body).Execute()

Patch config map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchConfigMapCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchConfigMapCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchConfigMap(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchConfigMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchConfigMapCommand**](PatchConfigMapCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchCrd

> KubernetesPatchCrd(ctx, v).Body(body).Execute()

Patch crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchCrdCommand(int32(123), "Yaml_example", "Name_example") // PatchCrdCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchCrd(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchCrdCommand**](PatchCrdCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchCronJob

> KubernetesPatchCronJob(ctx, v).Body(body).Execute()

Patch cronjob

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchCronJobCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchCronJobCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchCronJob(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchCronJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchCronJobCommand**](PatchCronJobCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchDeployment

> KubernetesPatchDeployment(ctx, v).Body(body).Execute()

Patch deployment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchDeploymentCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchDeploymentCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchDeployment(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchDeploymentCommand**](PatchDeploymentCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchIngress

> KubernetesPatchIngress(ctx, v).Body(body).Execute()

Patch ingress

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchIngressCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchIngressCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchIngress(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchIngressCommand**](PatchIngressCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchJob

> KubernetesPatchJob(ctx, v).Body(body).Execute()

Patch job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchJobCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchJobCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchJob(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchJobCommand**](PatchJobCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchNode

> KubernetesPatchNode(ctx, v).Body(body).Execute()

Patch node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchNodeCommand(int32(123), "Name_example") // PatchNodeCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchNode(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchNodeCommand**](PatchNodeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPdb

> KubernetesPatchPdb(ctx, v).Body(body).Execute()

Patch pdb

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchPdbCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPdbCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPdb(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchPdbCommand**](PatchPdbCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPod

> KubernetesPatchPod(ctx, v).Body(body).Execute()

Patch pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchPodCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPodCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPod(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchPodCommand**](PatchPodCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPvc

> KubernetesPatchPvc(ctx, v).Body(body).Execute()

Patch pvc

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchPvcCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPvcCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPvc(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchPvcCommand**](PatchPvcCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchSecret

> KubernetesPatchSecret(ctx, v).Body(body).Execute()

Patch secret

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchSecretCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchSecretCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchSecret(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchSecretCommand**](PatchSecretCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchSts

> KubernetesPatchSts(ctx, v).Body(body).Execute()

Patch sts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewPatchStsCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchStsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchSts(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchStsCommand**](PatchStsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartDaemonSet

> KubernetesRestartDaemonSet(ctx, v).Body(body).Execute()

Restart daemon set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewRestartDaemonSetCommand(int32(123), "Name_example", "Namespace_example") // RestartDaemonSetCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartDaemonSet(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RestartDaemonSetCommand**](RestartDaemonSetCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartDeployment

> KubernetesRestartDeployment(ctx, v).Body(body).Execute()

Restart deployment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewRestartDeploymentCommand(int32(123), "Name_example", "Namespace_example") // RestartDeploymentCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartDeployment(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RestartDeploymentCommand**](RestartDeploymentCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartSts

> KubernetesRestartSts(ctx, v).Body(body).Execute()

Restart stateful set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewRestartStsCommand(int32(123), "Name_example", "Namespace_example") // RestartStsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartSts(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RestartStsCommand**](RestartStsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSilenceManager

> KubernetesSilenceManager(ctx, v).Body(body).Execute()

Silence management for k8s alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    body := *openapiclient.NewSilenceOperationsCommand(int32(123), "Mode_example") // SilenceOperationsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesSilenceManager(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesSilenceManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSilenceManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SilenceOperationsCommand**](SilenceOperationsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesUpdateKubernetesAlert

> KubernetesUpdateKubernetesAlert(ctx, alertId, v).Body(body).Execute()

Update k8s alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    alertId := int32(56) // int32 | 
    v := "v_example" // string | 
    body := *openapiclient.NewUpdateKubernetesAlertDto() // UpdateKubernetesAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesUpdateKubernetesAlert(context.Background(), alertId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesUpdateKubernetesAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesUpdateKubernetesAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateKubernetesAlertDto**](UpdateKubernetesAlertDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

