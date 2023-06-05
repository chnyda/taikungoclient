# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsCommit**](ProjectsApi.md#ProjectsCommit) | **Post** /api/v{v}/Projects/commit/{projectId} | Commit changes for the given project. The changes will then be applied and the project will be updated. The project must be in the READY state.
[**ProjectsCreate**](ProjectsApi.md#ProjectsCreate) | **Post** /api/v{v}/Projects | Create a new project
[**ProjectsDelete**](ProjectsApi.md#ProjectsDelete) | **Post** /api/v{v}/Projects/delete | Delete the project. The project must be empty (no server) and in READY state
[**ProjectsDeleteWholeProject**](ProjectsApi.md#ProjectsDeleteWholeProject) | **Post** /api/v{v}/Projects/deletewholeproject | Delete whole project by project Id
[**ProjectsDescribe**](ProjectsApi.md#ProjectsDescribe) | **Get** /api/v{v}/Projects/describe/{projectId} | Describe project by Id
[**ProjectsDetails**](ProjectsApi.md#ProjectsDetails) | **Get** /api/v{v}/Projects/{projectId} | Retrieve details of the project by Id
[**ProjectsEdit**](ProjectsApi.md#ProjectsEdit) | **Post** /api/v{v}/Projects/update/{projectId} | Update project by Id for poller
[**ProjectsEditHealthStatus**](ProjectsApi.md#ProjectsEditHealthStatus) | **Post** /api/v{v}/Projects/updatehealth/{projectId} | Update health status of the project by Id
[**ProjectsExtendLifeTime**](ProjectsApi.md#ProjectsExtendLifeTime) | **Post** /api/v{v}/Projects/extend/lifetime | Extend life time of project
[**ProjectsFullSpotWorkersOperations**](ProjectsApi.md#ProjectsFullSpotWorkersOperations) | **Post** /api/v{v}/Projects/toggle-full-spot | Full spot operations enable/disable
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /api/v{v}/Projects | Retrieve a list of projects.
[**ProjectsListForBilling**](ProjectsApi.md#ProjectsListForBilling) | **Get** /api/v{v}/Projects/forbilling | Retrieve projects for billing
[**ProjectsListForPoller**](ProjectsApi.md#ProjectsListForPoller) | **Get** /api/v{v}/Projects/forpoller | Retrieve a list of projects for poller. Only available for admins.
[**ProjectsListSelector**](ProjectsApi.md#ProjectsListSelector) | **Get** /api/v{v}/Projects/list | List of projects with dropdown
[**ProjectsListWithAlerts**](ProjectsApi.md#ProjectsListWithAlerts) | **Get** /api/v{v}/Projects/foralerting | Retrieve a list of projects for alert poller. Only available for admins.
[**ProjectsLockManager**](ProjectsApi.md#ProjectsLockManager) | **Post** /api/v{v}/Projects/lockmanager | Lock/Unlock project
[**ProjectsLokiLogs**](ProjectsApi.md#ProjectsLokiLogs) | **Post** /api/v{v}/Projects/lokilogs | Retrieve loki logs
[**ProjectsMonitoringOperations**](ProjectsApi.md#ProjectsMonitoringOperations) | **Post** /api/v{v}/Projects/monitoring | Monitoring operations enable/disable
[**ProjectsProjectActionButtonVisibility**](ProjectsApi.md#ProjectsProjectActionButtonVisibility) | **Get** /api/v{v}/Projects/visibility/{projectId} | Visibility of project actions
[**ProjectsProjectMonitoringAlerts**](ProjectsApi.md#ProjectsProjectMonitoringAlerts) | **Post** /api/v{v}/Projects/monitoringalerts | Monitoring alerts for project
[**ProjectsPrometheusMetrics**](ProjectsApi.md#ProjectsPrometheusMetrics) | **Post** /api/v{v}/Projects/prometheusmetrics | Prometheus metrics data project
[**ProjectsPurge**](ProjectsApi.md#ProjectsPurge) | **Post** /api/v{v}/Projects/purge | Purge a list of servers from project by project Id
[**ProjectsPurgeWholeProject**](ProjectsApi.md#ProjectsPurgeWholeProject) | **Post** /api/v{v}/Projects/purgewholeproject | Purge a whole project by project Id
[**ProjectsRepair**](ProjectsApi.md#ProjectsRepair) | **Post** /api/v{v}/Projects/repair/{projectId} | Repair project by Id
[**ProjectsResetProjectStatus**](ProjectsApi.md#ProjectsResetProjectStatus) | **Post** /api/v{v}/Projects/reset | Change the project status for the given project. Only available for admin.
[**ProjectsSpotVmsOperations**](ProjectsApi.md#ProjectsSpotVmsOperations) | **Post** /api/v{v}/Projects/toggle-spot-vms | Spot vm(s) operations enable/disable
[**ProjectsSpotWorkersOperations**](ProjectsApi.md#ProjectsSpotWorkersOperations) | **Post** /api/v{v}/Projects/toggle-spot-workers | Spot worker(s) operations enable/disable
[**ProjectsUpgrade**](ProjectsApi.md#ProjectsUpgrade) | **Post** /api/v{v}/Projects/upgrade/{projectId} | Upgrade the project&#39;s Kubernetes to the next available version. Project must be READY.



## ProjectsCommit

> ProjectsCommit(ctx, projectId, v).Execute()

Commit changes for the given project. The changes will then be applied and the project will be updated. The project must be in the READY state.

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
    r, err := apiClient.ProjectsApi.ProjectsCommit(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCommit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsCommitRequest struct via the builder pattern


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


## ProjectsCreate

> ApiResponse ProjectsCreate(ctx, v).Body(body).Execute()

Create a new project

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
    body := *openapiclient.NewCreateProjectCommand("Name_example", int32(123)) // CreateProjectCommand | Create command (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateProjectCommand**](CreateProjectCommand.md) | Create command | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> ProjectsDelete(ctx, v).Body(body).Execute()

Delete the project. The project must be empty (no server) and in READY state

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
    body := *openapiclient.NewDeleteProjectCommand() // DeleteProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsDelete(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteProjectCommand**](DeleteProjectCommand.md) |  | 

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


## ProjectsDeleteWholeProject

> ProjectsDeleteWholeProject(ctx, v).Body(body).Execute()

Delete whole project by project Id

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
    body := *openapiclient.NewDeleteWholeProjectCommand() // DeleteWholeProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsDeleteWholeProject(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDeleteWholeProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsDeleteWholeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteWholeProjectCommand**](DeleteWholeProjectCommand.md) |  | 

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


## ProjectsDescribe

> map[string]interface{} ProjectsDescribe(ctx, projectId, v).IsYaml(isYaml).Execute()

Describe project by Id

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
    isYaml := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDescribe(context.Background(), projectId, v).IsYaml(isYaml).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsDescribe`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isYaml** | **bool** |  | 

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


## ProjectsDetails

> ProjectForListDto ProjectsDetails(ctx, projectId, v).Execute()

Retrieve details of the project by Id

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
    resp, r, err := apiClient.ProjectsApi.ProjectsDetails(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsDetails`: ProjectForListDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectForListDto**](ProjectForListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsEdit

> ProjectsEdit(ctx, projectId, v).Body(body).Execute()

Update project by Id for poller

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
    body := *openapiclient.NewProjectForUpdateDto() // ProjectForUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsEdit(context.Background(), projectId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProjectForUpdateDto**](ProjectForUpdateDto.md) |  | 

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


## ProjectsEditHealthStatus

> ProjectsEditHealthStatus(ctx, projectId, v).Body(body).Execute()

Update health status of the project by Id

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
    body := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsEditHealthStatus(context.Background(), projectId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsEditHealthStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsEditHealthStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** |  | 

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


## ProjectsExtendLifeTime

> ProjectsExtendLifeTime(ctx, v).Body(body).Execute()

Extend life time of project

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
    body := *openapiclient.NewProjectExtendLifeTimeCommand(int32(123)) // ProjectExtendLifeTimeCommand | Command (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsExtendLifeTime(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsExtendLifeTime``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsExtendLifeTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProjectExtendLifeTimeCommand**](ProjectExtendLifeTimeCommand.md) | Command | 

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


## ProjectsFullSpotWorkersOperations

> ProjectsFullSpotWorkersOperations(ctx, v).Body(body).Execute()

Full spot operations enable/disable

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
    body := *openapiclient.NewFullSpotOperationCommand() // FullSpotOperationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsFullSpotWorkersOperations(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsFullSpotWorkersOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsFullSpotWorkersOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FullSpotOperationCommand**](FullSpotOperationCommand.md) |  | 

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


## ProjectsList

> ProjectsList ProjectsList(ctx, v).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).Execute()

Retrieve a list of projects.

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
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string | Keyword for searching (optional)
    updatedAt := time.Now() // time.Time |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    backupCredentialId := int32(56) // int32 |  (optional)
    healthy := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsList(context.Background(), v).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsList`: ProjectsList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** | Keyword for searching | 
 **updatedAt** | **time.Time** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **backupCredentialId** | **int32** |  | 
 **healthy** | **bool** |  | 

### Return type

[**ProjectsList**](ProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListForBilling

> []ProjectsForBillingDto ProjectsListForBilling(ctx, v).Execute()

Retrieve projects for billing

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
    resp, r, err := apiClient.ProjectsApi.ProjectsListForBilling(context.Background(), v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListForBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListForBilling`: []ProjectsForBillingDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListForBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListForBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectsForBillingDto**](ProjectsForBillingDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListForPoller

> ProjectListForPoller ProjectsListForPoller(ctx, v).Limit(limit).Offset(offset).Search(search).UpdatedAt(updatedAt).Execute()

Retrieve a list of projects for poller. Only available for admins.

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
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    search := "search_example" // string | Keyword for searching (optional)
    updatedAt := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListForPoller(context.Background(), v).Limit(limit).Offset(offset).Search(search).UpdatedAt(updatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListForPoller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListForPoller`: ProjectListForPoller
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListForPoller`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListForPollerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **search** | **string** | Keyword for searching | 
 **updatedAt** | **time.Time** |  | 

### Return type

[**ProjectListForPoller**](ProjectListForPoller.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListSelector

> []CommonDropdownIsBoundDtoForProject ProjectsListSelector(ctx, v).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).Execute()

List of projects with dropdown

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
    search := "search_example" // string |  (optional)
    catalogId := int32(56) // int32 |  (optional)
    healthy := true // bool |  (optional)
    userId := "userId_example" // string |  (optional)
    ready := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListSelector(context.Background(), v).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListSelector`: []CommonDropdownIsBoundDtoForProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **catalogId** | **int32** |  | 
 **healthy** | **bool** |  | 
 **userId** | **string** |  | 
 **ready** | **bool** |  | 

### Return type

[**[]CommonDropdownIsBoundDtoForProject**](CommonDropdownIsBoundDtoForProject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListWithAlerts

> []ProjectListForAlert ProjectsListWithAlerts(ctx, v).Limit(limit).Offset(offset).Status(status).ProjectId(projectId).Execute()

Retrieve a list of projects for alert poller. Only available for admins.

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
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    status := "status_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListWithAlerts(context.Background(), v).Limit(limit).Offset(offset).Status(status).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListWithAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListWithAlerts`: []ProjectListForAlert
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListWithAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListWithAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **status** | **string** |  | 
 **projectId** | **int32** |  | 

### Return type

[**[]ProjectListForAlert**](ProjectListForAlert.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLockManager

> ProjectsLockManager(ctx, v).Id(id).Mode(mode).Execute()

Lock/Unlock project

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
    id := int32(56) // int32 |  (optional)
    mode := "mode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsLockManager(context.Background(), v).Id(id).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsLockManager``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **int32** |  | 
 **mode** | **string** |  | 

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


## ProjectsLokiLogs

> ProjectsLokiLogs(ctx, v).Body(body).Execute()

Retrieve loki logs

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
    body := *openapiclient.NewLokiResponseDto() // LokiResponseDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsLokiLogs(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsLokiLogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsLokiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LokiResponseDto**](LokiResponseDto.md) |  | 

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


## ProjectsMonitoringOperations

> ProjectsMonitoringOperations(ctx, v).Body(body).Execute()

Monitoring operations enable/disable

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
    body := *openapiclient.NewMonitoringOperationsCommand() // MonitoringOperationsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsMonitoringOperations(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsMonitoringOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsMonitoringOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MonitoringOperationsCommand**](MonitoringOperationsCommand.md) |  | 

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


## ProjectsProjectActionButtonVisibility

> ProjectActionVisibilityDto ProjectsProjectActionButtonVisibility(ctx, projectId, v).Execute()

Visibility of project actions

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
    resp, r, err := apiClient.ProjectsApi.ProjectsProjectActionButtonVisibility(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsProjectActionButtonVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsProjectActionButtonVisibility`: ProjectActionVisibilityDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsProjectActionButtonVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectActionButtonVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectActionVisibilityDto**](ProjectActionVisibilityDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectMonitoringAlerts

> map[string]interface{} ProjectsProjectMonitoringAlerts(ctx, v).Body(body).Execute()

Monitoring alerts for project

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
    body := *openapiclient.NewProjectsMonitoringAlertsCommand() // ProjectsMonitoringAlertsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsProjectMonitoringAlerts(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsProjectMonitoringAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsProjectMonitoringAlerts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsProjectMonitoringAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectMonitoringAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProjectsMonitoringAlertsCommand**](ProjectsMonitoringAlertsCommand.md) |  | 

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


## ProjectsPrometheusMetrics

> map[string]interface{} ProjectsPrometheusMetrics(ctx, v).Body(body).Execute()

Prometheus metrics data project

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
    body := *openapiclient.NewPrometheusMetricsCommand() // PrometheusMetricsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsPrometheusMetrics(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPrometheusMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsPrometheusMetrics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsPrometheusMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPrometheusMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PrometheusMetricsCommand**](PrometheusMetricsCommand.md) |  | 

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


## ProjectsPurge

> ProjectsPurge(ctx, v).Body(body).Execute()

Purge a list of servers from project by project Id

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
    body := *openapiclient.NewPurgeCommand() // PurgeCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsPurge(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPurge``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PurgeCommand**](PurgeCommand.md) |  | 

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


## ProjectsPurgeWholeProject

> ProjectsPurgeWholeProject(ctx, v).Body(body).Execute()

Purge a whole project by project Id

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
    body := *openapiclient.NewPurgeWholeProjectCommand() // PurgeWholeProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsPurgeWholeProject(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPurgeWholeProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsPurgeWholeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PurgeWholeProjectCommand**](PurgeWholeProjectCommand.md) |  | 

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


## ProjectsRepair

> ProjectsRepair(ctx, projectId, v).Execute()

Repair project by Id

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
    r, err := apiClient.ProjectsApi.ProjectsRepair(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsRepair``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsRepairRequest struct via the builder pattern


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


## ProjectsResetProjectStatus

> ProjectsResetProjectStatus(ctx, v).Body(body).Execute()

Change the project status for the given project. Only available for admin.

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
    body := *openapiclient.NewResetProjectStatusCommand() // ResetProjectStatusCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsResetProjectStatus(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsResetProjectStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsResetProjectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ResetProjectStatusCommand**](ResetProjectStatusCommand.md) |  | 

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


## ProjectsSpotVmsOperations

> ProjectsSpotVmsOperations(ctx, v).Body(body).Execute()

Spot vm(s) operations enable/disable

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
    body := *openapiclient.NewSpotVmOperationCommand() // SpotVmOperationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsSpotVmsOperations(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsSpotVmsOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsSpotVmsOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SpotVmOperationCommand**](SpotVmOperationCommand.md) |  | 

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


## ProjectsSpotWorkersOperations

> ProjectsSpotWorkersOperations(ctx, v).Body(body).Execute()

Spot worker(s) operations enable/disable

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
    body := *openapiclient.NewSpotWorkerOperationCommand() // SpotWorkerOperationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsSpotWorkersOperations(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsSpotWorkersOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsSpotWorkersOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SpotWorkerOperationCommand**](SpotWorkerOperationCommand.md) |  | 

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


## ProjectsUpgrade

> ProjectsUpgrade(ctx, projectId, v).Execute()

Upgrade the project's Kubernetes to the next available version. Project must be READY.

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
    r, err := apiClient.ProjectsApi.ProjectsUpgrade(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpgrade``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsUpgradeRequest struct via the builder pattern


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

