# \ServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServersConsole**](ServersApi.md#ServersConsole) | **Post** /api/v1/servers/console | Console screenshot or terminal access for server
[**ServersCreate**](ServersApi.md#ServersCreate) | **Post** /api/v1/servers/create | Create a new server in the given project.
[**ServersDelete**](ServersApi.md#ServersDelete) | **Post** /api/v1/servers/delete | Delete server by project id
[**ServersDetails**](ServersApi.md#ServersDetails) | **Get** /api/v1/servers/{projectId} | Retrieve all servers by given project
[**ServersList**](ServersApi.md#ServersList) | **Get** /api/v1/servers | Retrieve all servers
[**ServersReboot**](ServersApi.md#ServersReboot) | **Post** /api/v1/servers/reboot | Reboot server
[**ServersReset**](ServersApi.md#ServersReset) | **Post** /api/v1/servers/reset | Update server(s) status(es)
[**ServersStatus**](ServersApi.md#ServersStatus) | **Get** /api/v1/servers/status/{serverId} | Show server status
[**ServersUpdate**](ServersApi.md#ServersUpdate) | **Post** /api/v1/servers/update | Update server
[**ServersUpdateByProjectId**](ServersApi.md#ServersUpdateByProjectId) | **Put** /api/v1/servers/update/{projectId} | Update server by projectId



## ServersConsole

> string ServersConsole(ctx).ConsoleScreenshotCommand(consoleScreenshotCommand).Execute()

Console screenshot or terminal access for server

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
    consoleScreenshotCommand := *openapiclient.NewConsoleScreenshotCommand() // ConsoleScreenshotCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersConsole(context.Background()).ConsoleScreenshotCommand(consoleScreenshotCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersConsole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersConsole`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersConsole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleScreenshotCommand** | [**ConsoleScreenshotCommand**](ConsoleScreenshotCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersCreate

> ApiResponse ServersCreate(ctx).ServerForCreateDto(serverForCreateDto).Execute()

Create a new server in the given project.

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
    serverForCreateDto := *openapiclient.NewServerForCreateDto() // ServerForCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersCreate(context.Background()).ServerForCreateDto(serverForCreateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverForCreateDto** | [**ServerForCreateDto**](ServerForCreateDto.md) |  | 

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


## ServersDelete

> ServersDelete(ctx).DeleteServerCommand(deleteServerCommand).Execute()

Delete server by project id

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
    deleteServerCommand := *openapiclient.NewDeleteServerCommand() // DeleteServerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersDelete(context.Background()).DeleteServerCommand(deleteServerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServerCommand** | [**DeleteServerCommand**](DeleteServerCommand.md) |  | 

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


## ServersDetails

> ServersListForDetails ServersDetails(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()

Retrieve all servers by given project

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    withAutoscalingGroup := "withAutoscalingGroup_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersDetails(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersDetails`: ServersListForDetails
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **withAutoscalingGroup** | **string** |  | 

### Return type

[**ServersListForDetails**](ServersListForDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersList

> ServersList ServersList(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).Execute()

Retrieve all servers

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
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    startRam := "startRam_example" // string |  (optional)
    endRam := "endRam_example" // string |  (optional)
    startDiskSize := int64(789) // int64 |  (optional)
    endDiskSize := int64(789) // int64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    id := int32(56) // int32 |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    autoscalingGroup := "autoscalingGroup_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersList(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersList`: ServersList
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **startRam** | **string** |  | 
 **endRam** | **string** |  | 
 **startDiskSize** | **int64** |  | 
 **endDiskSize** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **id** | **int32** |  | 
 **filterBy** | **string** |  | 
 **autoscalingGroup** | **string** |  | 

### Return type

[**ServersList**](ServersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersReboot

> ServersReboot(ctx).RebootServerCommand(rebootServerCommand).Execute()

Reboot server

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
    rebootServerCommand := *openapiclient.NewRebootServerCommand() // RebootServerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersReboot(context.Background()).RebootServerCommand(rebootServerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersReboot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootServerCommand** | [**RebootServerCommand**](RebootServerCommand.md) |  | 

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


## ServersReset

> ServersReset(ctx).ResetServerStatusCommand(resetServerStatusCommand).Execute()

Update server(s) status(es)

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
    resetServerStatusCommand := *openapiclient.NewResetServerStatusCommand() // ResetServerStatusCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersReset(context.Background()).ResetServerStatusCommand(resetServerStatusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetServerStatusCommand** | [**ResetServerStatusCommand**](ResetServerStatusCommand.md) |  | 

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


## ServersStatus

> string ServersStatus(ctx, serverId).Execute()

Show server status

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
    serverId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersStatus(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersStatus`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersUpdate

> ServersUpdate(ctx).UpdateServerCommand(updateServerCommand).Execute()

Update server

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
    updateServerCommand := *openapiclient.NewUpdateServerCommand() // UpdateServerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersUpdate(context.Background()).UpdateServerCommand(updateServerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateServerCommand** | [**UpdateServerCommand**](UpdateServerCommand.md) |  | 

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


## ServersUpdateByProjectId

> ServersUpdateByProjectId(ctx, projectId).UpdateServerHealthDto(updateServerHealthDto).Execute()

Update server by projectId

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
    updateServerHealthDto := []openapiclient.UpdateServerHealthDto{*openapiclient.NewUpdateServerHealthDto()} // []UpdateServerHealthDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersUpdateByProjectId(context.Background(), projectId).UpdateServerHealthDto(updateServerHealthDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersUpdateByProjectId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersUpdateByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerHealthDto** | [**[]UpdateServerHealthDto**](UpdateServerHealthDto.md) |  | 

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

