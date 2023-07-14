# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsAiAnalyzer**](ProjectsApi.md#ProjectsAiAnalyzer) | **Get** /api/v1/projects/ai-analyze/{projectId} | Analyze cluster by AI model
[**ProjectsChatCompletions**](ProjectsApi.md#ProjectsChatCompletions) | **Post** /api/v1/projects/chat/completions | AI Chat completions
[**ProjectsCommit**](ProjectsApi.md#ProjectsCommit) | **Post** /api/v1/projects/commit/{projectId} | Commit changes for the given project. The changes will then be applied and the project will be updated. The project must be in the READY state.
[**ProjectsCreate**](ProjectsApi.md#ProjectsCreate) | **Post** /api/v1/projects | Create a new project
[**ProjectsDelete**](ProjectsApi.md#ProjectsDelete) | **Post** /api/v1/projects/delete | Delete the project. The project must be empty (no server) and in READY state
[**ProjectsDeleteWholeProject**](ProjectsApi.md#ProjectsDeleteWholeProject) | **Post** /api/v1/projects/deletewholeproject | Delete whole project by project Id
[**ProjectsDescribe**](ProjectsApi.md#ProjectsDescribe) | **Get** /api/v1/projects/describe/{projectId} | Describe project by Id
[**ProjectsDetails**](ProjectsApi.md#ProjectsDetails) | **Get** /api/v1/projects/{projectId} | Retrieve details of the project by Id
[**ProjectsDropdown**](ProjectsApi.md#ProjectsDropdown) | **Get** /api/v1/projects/list | Retrieve list of projects for dropdown
[**ProjectsEdit**](ProjectsApi.md#ProjectsEdit) | **Put** /api/v1/projects/edit/{projectId} | Update project by Id for poller
[**ProjectsEditHealth**](ProjectsApi.md#ProjectsEditHealth) | **Put** /api/v1/projects/edit/health | Update health status of the project by Id
[**ProjectsEditStatus**](ProjectsApi.md#ProjectsEditStatus) | **Put** /api/v1/projects/edit/status | Change the project status for the given project. Only available for admin.
[**ProjectsExtendLifetime**](ProjectsApi.md#ProjectsExtendLifetime) | **Post** /api/v1/projects/extend/lifetime | Extend life time of project
[**ProjectsForAlerting**](ProjectsApi.md#ProjectsForAlerting) | **Get** /api/v1/projects/foralerting | Retrieve a list of projects for alert poller. Only available for admins.
[**ProjectsForBilling**](ProjectsApi.md#ProjectsForBilling) | **Get** /api/v1/projects/forbilling | Retrieve a list of projects for billing
[**ProjectsForPoller**](ProjectsApi.md#ProjectsForPoller) | **Get** /api/v1/projects/forpoller | Retrieve a list of projects for poller. Only available for admins.
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /api/v1/projects | Retrieve all projects
[**ProjectsLockManager**](ProjectsApi.md#ProjectsLockManager) | **Post** /api/v1/projects/lockmanager | Lock/Unlock project
[**ProjectsLokiLogs**](ProjectsApi.md#ProjectsLokiLogs) | **Post** /api/v1/projects/lokilogs | Retrieve loki logs
[**ProjectsMonitoring**](ProjectsApi.md#ProjectsMonitoring) | **Post** /api/v1/projects/monitoring | Monitoring operations enable/disable
[**ProjectsMonitoringAlerts**](ProjectsApi.md#ProjectsMonitoringAlerts) | **Post** /api/v1/projects/monitoringalerts | Monitoring alerts for project
[**ProjectsPrometheusMetrics**](ProjectsApi.md#ProjectsPrometheusMetrics) | **Post** /api/v1/projects/prometheusmetrics | Prometheus metrics data project
[**ProjectsPurge**](ProjectsApi.md#ProjectsPurge) | **Post** /api/v1/projects/purge | Purge a list of servers from project by project Id
[**ProjectsPurgeWholeProject**](ProjectsApi.md#ProjectsPurgeWholeProject) | **Post** /api/v1/projects/purgewholeproject | Purge a whole project by project Id
[**ProjectsRepair**](ProjectsApi.md#ProjectsRepair) | **Post** /api/v1/projects/repair/{projectId} | Repair project by Id
[**ProjectsToggleFullSpot**](ProjectsApi.md#ProjectsToggleFullSpot) | **Post** /api/v1/projects/toggle-full-spot | Full spot operations enable/disable
[**ProjectsToggleSpotVms**](ProjectsApi.md#ProjectsToggleSpotVms) | **Post** /api/v1/projects/toggle-spot-vms | Spot vm(s) operations enable/disable
[**ProjectsToggleSpotWorkers**](ProjectsApi.md#ProjectsToggleSpotWorkers) | **Post** /api/v1/projects/toggle-spot-workers | Spot worker(s) operations enable/disable
[**ProjectsUpgrade**](ProjectsApi.md#ProjectsUpgrade) | **Post** /api/v1/projects/upgrade/{projectId} | Upgrade the project&#39;s Kubernetes to the next available version. Project must be READY.
[**ProjectsVisibility**](ProjectsApi.md#ProjectsVisibility) | **Get** /api/v1/projects/visibility/{projectId} | Visibility of project actions



## ProjectsAiAnalyzer

> interface{} ProjectsAiAnalyzer(ctx, projectId).Execute()

Analyze cluster by AI model

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsAiAnalyzer(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsAiAnalyzer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsAiAnalyzer`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsAiAnalyzer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAiAnalyzerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsChatCompletions

> interface{} ProjectsChatCompletions(ctx).ChatCompletionsCommand(chatCompletionsCommand).Execute()

AI Chat completions

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
    chatCompletionsCommand := *openapiclient.NewChatCompletionsCommand() // ChatCompletionsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsChatCompletions(context.Background()).ChatCompletionsCommand(chatCompletionsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsChatCompletions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsChatCompletions`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionsCommand** | [**ChatCompletionsCommand**](ChatCompletionsCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCommit

> ProjectsCommit(ctx, projectId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsCommit(context.Background(), projectId).Execute()
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreate

> ApiResponse ProjectsCreate(ctx).CreateProjectCommand(createProjectCommand).Execute()

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
    createProjectCommand := *openapiclient.NewCreateProjectCommand() // CreateProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreate(context.Background()).CreateProjectCommand(createProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectCommand** | [**CreateProjectCommand**](CreateProjectCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> ProjectsDelete(ctx).DeleteProjectCommand(deleteProjectCommand).Execute()

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
    deleteProjectCommand := *openapiclient.NewDeleteProjectCommand() // DeleteProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsDelete(context.Background()).DeleteProjectCommand(deleteProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteProjectCommand** | [**DeleteProjectCommand**](DeleteProjectCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDeleteWholeProject

> ProjectsDeleteWholeProject(ctx).DeleteWholeProjectCommand(deleteWholeProjectCommand).Execute()

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
    deleteWholeProjectCommand := *openapiclient.NewDeleteWholeProjectCommand() // DeleteWholeProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsDeleteWholeProject(context.Background()).DeleteWholeProjectCommand(deleteWholeProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDeleteWholeProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteWholeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteWholeProjectCommand** | [**DeleteWholeProjectCommand**](DeleteWholeProjectCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDescribe

> interface{} ProjectsDescribe(ctx, projectId).IsYaml(isYaml).Execute()

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
    isYaml := true // bool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDescribe(context.Background(), projectId).IsYaml(isYaml).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsDescribe`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isYaml** | **bool** |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDetails

> ProjectForListDto ProjectsDetails(ctx, projectId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDetails(context.Background(), projectId).Execute()
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDropdown

> []CommonDropdownIsBoundDtoForProject ProjectsDropdown(ctx).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).Execute()

Retrieve list of projects for dropdown

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
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    catalogId := int32(56) // int32 |  (optional)
    healthy := true // bool |  (optional)
    userId := "userId_example" // string |  (optional)
    ready := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDropdown(context.Background()).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsDropdown`: []CommonDropdownIsBoundDtoForProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDropdownRequest struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsEdit

> ProjectsEdit(ctx, projectId).ProjectForUpdateDto(projectForUpdateDto).Execute()

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
    projectForUpdateDto := *openapiclient.NewProjectForUpdateDto() // ProjectForUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsEdit(context.Background(), projectId).ProjectForUpdateDto(projectForUpdateDto).Execute()
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

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectForUpdateDto** | [**ProjectForUpdateDto**](ProjectForUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsEditHealth

> ProjectsEditHealth(ctx).UpdateHealthStatusCommand(updateHealthStatusCommand).Execute()

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
    updateHealthStatusCommand := *openapiclient.NewUpdateHealthStatusCommand() // UpdateHealthStatusCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsEditHealth(context.Background()).UpdateHealthStatusCommand(updateHealthStatusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsEditHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsEditHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateHealthStatusCommand** | [**UpdateHealthStatusCommand**](UpdateHealthStatusCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsEditStatus

> ProjectsEditStatus(ctx).ResetProjectStatusCommand(resetProjectStatusCommand).Execute()

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
    resetProjectStatusCommand := *openapiclient.NewResetProjectStatusCommand() // ResetProjectStatusCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsEditStatus(context.Background()).ResetProjectStatusCommand(resetProjectStatusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsEditStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsEditStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetProjectStatusCommand** | [**ResetProjectStatusCommand**](ResetProjectStatusCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsExtendLifetime

> ProjectsExtendLifetime(ctx).ProjectExtendLifeTimeCommand(projectExtendLifeTimeCommand).Execute()

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
    projectExtendLifeTimeCommand := *openapiclient.NewProjectExtendLifeTimeCommand() // ProjectExtendLifeTimeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsExtendLifetime(context.Background()).ProjectExtendLifeTimeCommand(projectExtendLifeTimeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsExtendLifetime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsExtendLifetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectExtendLifeTimeCommand** | [**ProjectExtendLifeTimeCommand**](ProjectExtendLifeTimeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsForAlerting

> []ProjectListForAlert ProjectsForAlerting(ctx).Limit(limit).Offset(offset).Status(status).ProjectId(projectId).Execute()

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    status := "status_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsForAlerting(context.Background()).Limit(limit).Offset(offset).Status(status).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsForAlerting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsForAlerting`: []ProjectListForAlert
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsForAlerting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsForAlertingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **status** | **string** |  | 
 **projectId** | **int32** |  | 

### Return type

[**[]ProjectListForAlert**](ProjectListForAlert.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsForBilling

> []ProjectsForBillingDto ProjectsForBilling(ctx).Execute()

Retrieve a list of projects for billing

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsForBilling(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsForBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsForBilling`: []ProjectsForBillingDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsForBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsForBillingRequest struct via the builder pattern


### Return type

[**[]ProjectsForBillingDto**](ProjectsForBillingDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsForPoller

> ProjectListForPoller ProjectsForPoller(ctx).Limit(limit).Offset(offset).Search(search).UpdatedAt(updatedAt).Execute()

Retrieve a list of projects for poller. Only available for admins.

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsForPoller(context.Background()).Limit(limit).Offset(offset).Search(search).UpdatedAt(updatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsForPoller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsForPoller`: ProjectListForPoller
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsForPoller`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsForPollerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **search** | **string** |  | 
 **updatedAt** | **string** |  | 

### Return type

[**ProjectListForPoller**](ProjectListForPoller.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ProjectsList ProjectsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).Execute()

Retrieve all projects

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    updatedAt := "updatedAt_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    backupCredentialId := int32(56) // int32 |  (optional)
    healthy := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsList`: ProjectsList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **updatedAt** | **string** |  | 
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLockManager

> ProjectsLockManager(ctx).ProjectLockManagerCommand(projectLockManagerCommand).Execute()

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
    projectLockManagerCommand := *openapiclient.NewProjectLockManagerCommand() // ProjectLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsLockManager(context.Background()).ProjectLockManagerCommand(projectLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectLockManagerCommand** | [**ProjectLockManagerCommand**](ProjectLockManagerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLokiLogs

> ProjectsLokiLogs(ctx).LokiResponseDto(lokiResponseDto).Execute()

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
    lokiResponseDto := *openapiclient.NewLokiResponseDto() // LokiResponseDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsLokiLogs(context.Background()).LokiResponseDto(lokiResponseDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsLokiLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLokiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lokiResponseDto** | [**LokiResponseDto**](LokiResponseDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsMonitoring

> ProjectsMonitoring(ctx).MonitoringOperationsCommand(monitoringOperationsCommand).Execute()

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
    monitoringOperationsCommand := *openapiclient.NewMonitoringOperationsCommand() // MonitoringOperationsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsMonitoring(context.Background()).MonitoringOperationsCommand(monitoringOperationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitoringOperationsCommand** | [**MonitoringOperationsCommand**](MonitoringOperationsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsMonitoringAlerts

> interface{} ProjectsMonitoringAlerts(ctx).ProjectsMonitoringAlertsCommand(projectsMonitoringAlertsCommand).Execute()

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
    projectsMonitoringAlertsCommand := *openapiclient.NewProjectsMonitoringAlertsCommand() // ProjectsMonitoringAlertsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsMonitoringAlerts(context.Background()).ProjectsMonitoringAlertsCommand(projectsMonitoringAlertsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsMonitoringAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsMonitoringAlerts`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsMonitoringAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMonitoringAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsMonitoringAlertsCommand** | [**ProjectsMonitoringAlertsCommand**](ProjectsMonitoringAlertsCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPrometheusMetrics

> interface{} ProjectsPrometheusMetrics(ctx).PrometheusMetricsCommand(prometheusMetricsCommand).Execute()

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
    prometheusMetricsCommand := *openapiclient.NewPrometheusMetricsCommand() // PrometheusMetricsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsPrometheusMetrics(context.Background()).PrometheusMetricsCommand(prometheusMetricsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPrometheusMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsPrometheusMetrics`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsPrometheusMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPrometheusMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusMetricsCommand** | [**PrometheusMetricsCommand**](PrometheusMetricsCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPurge

> ProjectsPurge(ctx).PurgeCommand(purgeCommand).Execute()

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
    purgeCommand := *openapiclient.NewPurgeCommand() // PurgeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsPurge(context.Background()).PurgeCommand(purgeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purgeCommand** | [**PurgeCommand**](PurgeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPurgeWholeProject

> ProjectsPurgeWholeProject(ctx).PurgeWholeProjectCommand(purgeWholeProjectCommand).Execute()

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
    purgeWholeProjectCommand := *openapiclient.NewPurgeWholeProjectCommand() // PurgeWholeProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsPurgeWholeProject(context.Background()).PurgeWholeProjectCommand(purgeWholeProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPurgeWholeProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPurgeWholeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purgeWholeProjectCommand** | [**PurgeWholeProjectCommand**](PurgeWholeProjectCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsRepair

> ProjectsRepair(ctx, projectId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsRepair(context.Background(), projectId).Execute()
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsToggleFullSpot

> ProjectsToggleFullSpot(ctx).FullSpotOperationCommand(fullSpotOperationCommand).Execute()

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
    fullSpotOperationCommand := *openapiclient.NewFullSpotOperationCommand() // FullSpotOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsToggleFullSpot(context.Background()).FullSpotOperationCommand(fullSpotOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsToggleFullSpot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleFullSpotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullSpotOperationCommand** | [**FullSpotOperationCommand**](FullSpotOperationCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsToggleSpotVms

> ProjectsToggleSpotVms(ctx).SpotVmOperationCommand(spotVmOperationCommand).Execute()

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
    spotVmOperationCommand := *openapiclient.NewSpotVmOperationCommand() // SpotVmOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsToggleSpotVms(context.Background()).SpotVmOperationCommand(spotVmOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsToggleSpotVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleSpotVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spotVmOperationCommand** | [**SpotVmOperationCommand**](SpotVmOperationCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsToggleSpotWorkers

> ProjectsToggleSpotWorkers(ctx).SpotWorkerOperationCommand(spotWorkerOperationCommand).Execute()

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
    spotWorkerOperationCommand := *openapiclient.NewSpotWorkerOperationCommand() // SpotWorkerOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsToggleSpotWorkers(context.Background()).SpotWorkerOperationCommand(spotWorkerOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsToggleSpotWorkers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleSpotWorkersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spotWorkerOperationCommand** | [**SpotWorkerOperationCommand**](SpotWorkerOperationCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpgrade

> ProjectsUpgrade(ctx, projectId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ProjectsUpgrade(context.Background(), projectId).Execute()
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsVisibility

> ProjectActionVisibilityDto ProjectsVisibility(ctx, projectId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsVisibility(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsVisibility`: ProjectActionVisibilityDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectActionVisibilityDto**](ProjectActionVisibilityDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

