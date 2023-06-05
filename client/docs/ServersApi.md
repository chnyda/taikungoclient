# \ServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServersConsole**](ServersApi.md#ServersConsole) | **Post** /api/v{v}/Servers/console | Console servers
[**ServersCreate**](ServersApi.md#ServersCreate) | **Post** /api/v{v}/Servers | Create a new server in the given project.
[**ServersDelete**](ServersApi.md#ServersDelete) | **Post** /api/v{v}/Servers/delete | Delete servers by project Id
[**ServersDetails**](ServersApi.md#ServersDetails) | **Get** /api/v{v}/Servers/{projectId} | Retrieve a list of servers with detailed info
[**ServersList**](ServersApi.md#ServersList) | **Get** /api/v{v}/Servers | List all servers according to roles
[**ServersReboot**](ServersApi.md#ServersReboot) | **Post** /api/v{v}/Servers/reboot | Reboot servers
[**ServersReset**](ServersApi.md#ServersReset) | **Post** /api/v{v}/Servers/reset | Update statuses of servers by project Id
[**ServersShowServerStatus**](ServersApi.md#ServersShowServerStatus) | **Get** /api/v{v}/Servers/status/{serverId} | Show server status
[**ServersUpdate**](ServersApi.md#ServersUpdate) | **Post** /api/v{v}/Servers/update | Update server
[**ServersUpdateByProjectId**](ServersApi.md#ServersUpdateByProjectId) | **Put** /api/v{v}/Servers/update/{projectId} | Update server by projectId



## ServersConsole

> string ServersConsole(ctx, v).Body(body).Execute()

Console servers

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
    body := *openapiclient.NewConsoleScreenshotCommand() // ConsoleScreenshotCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersConsole(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersConsole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersConsole`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersConsole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConsoleScreenshotCommand**](ConsoleScreenshotCommand.md) |  | 

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


## ServersCreate

> ApiResponse ServersCreate(ctx, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := *openapiclient.NewServerForCreateDto() // ServerForCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ServerForCreateDto**](ServerForCreateDto.md) |  | 

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


## ServersDelete

> ServersDelete(ctx, v).Body(body).Execute()

Delete servers by project Id

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
    body := *openapiclient.NewDeleteServerCommand() // DeleteServerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersDelete(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteServerCommand**](DeleteServerCommand.md) |  | 

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


## ServersDetails

> ServersListForDetails ServersDetails(ctx, projectId, v).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()

Retrieve a list of servers with detailed info

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    withAutoscalingGroup := "withAutoscalingGroup_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersDetails(context.Background(), projectId, v).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()
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
**v** | **string** |  | 

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
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersList

> ServersList ServersList(ctx, v).Offset(offset).Limit(limit).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).Execute()

List all servers according to roles

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
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    startRam := float64(1.2) // float64 |  (optional)
    endRam := float64(1.2) // float64 |  (optional)
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
    resp, r, err := apiClient.ServersApi.ServersList(context.Background(), v).Offset(offset).Limit(limit).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersList`: ServersList
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits user size (by default 50) | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
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
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersReboot

> ServersReboot(ctx, v).Body(body).Execute()

Reboot servers

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
    body := *openapiclient.NewRebootServerCommand() // RebootServerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersReboot(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersReboot``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RebootServerCommand**](RebootServerCommand.md) |  | 

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


## ServersReset

> ServersReset(ctx, v).Body(body).Execute()

Update statuses of servers by project Id

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
    body := *openapiclient.NewResetServerStatusCommand(int32(123)) // ResetServerStatusCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersReset(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersReset``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ResetServerStatusCommand**](ResetServerStatusCommand.md) |  | 

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


## ServersShowServerStatus

> string ServersShowServerStatus(ctx, serverId, v).Execute()

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
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersShowServerStatus(context.Background(), serverId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersShowServerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersShowServerStatus`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersShowServerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersShowServerStatusRequest struct via the builder pattern


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


## ServersUpdate

> ServersUpdate(ctx, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := *openapiclient.NewUpdateServerCommand(int32(123), "Ip_example") // UpdateServerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersUpdate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateServerCommand**](UpdateServerCommand.md) |  | 

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


## ServersUpdateByProjectId

> ServersUpdateByProjectId(ctx, projectId, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := []openapiclient.UpdateServerHealthDto{*openapiclient.NewUpdateServerHealthDto()} // []UpdateServerHealthDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersApi.ServersUpdateByProjectId(context.Background(), projectId, v).Body(body).Execute()
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
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersUpdateByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**[]UpdateServerHealthDto**](UpdateServerHealthDto.md) |  | 

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

