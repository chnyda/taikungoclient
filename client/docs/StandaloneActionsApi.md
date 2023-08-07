# \StandaloneActionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandaloneactionsConsole**](StandaloneActionsApi.md#StandaloneactionsConsole) | **Post** /api/v1/standaloneactions/console | Console screenshot or terminal for vm
[**StandaloneactionsDownloadRdp**](StandaloneActionsApi.md#StandaloneactionsDownloadRdp) | **Get** /api/v1/standaloneactions/download/rdp/{id} | Download RDP file
[**StandaloneactionsReboot**](StandaloneActionsApi.md#StandaloneactionsReboot) | **Post** /api/v1/standaloneactions/reboot | Reboot standalone vm
[**StandaloneactionsShelve**](StandaloneActionsApi.md#StandaloneactionsShelve) | **Post** /api/v1/standaloneactions/shelve | Shelve standalone vm
[**StandaloneactionsStart**](StandaloneActionsApi.md#StandaloneactionsStart) | **Post** /api/v1/standaloneactions/start | Start standalone vm
[**StandaloneactionsStatus**](StandaloneActionsApi.md#StandaloneactionsStatus) | **Get** /api/v1/standaloneactions/status/{id} | Show standalone vm status
[**StandaloneactionsStop**](StandaloneActionsApi.md#StandaloneactionsStop) | **Post** /api/v1/standaloneactions/stop | Stop standalone vm
[**StandaloneactionsUnshelve**](StandaloneActionsApi.md#StandaloneactionsUnshelve) | **Post** /api/v1/standaloneactions/unshelve | Unshelve standalone vm
[**StandaloneactionsWindowsInstancePassword**](StandaloneActionsApi.md#StandaloneactionsWindowsInstancePassword) | **Post** /api/v1/standaloneactions/password | Retrieve aws windows admin instance password



## StandaloneactionsConsole

> string StandaloneactionsConsole(ctx).VmConsoleScreenshotCommand(vmConsoleScreenshotCommand).Execute()

Console screenshot or terminal for vm

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
    vmConsoleScreenshotCommand := *openapiclient.NewVmConsoleScreenshotCommand() // VmConsoleScreenshotCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneActionsApi.StandaloneactionsConsole(context.Background()).VmConsoleScreenshotCommand(vmConsoleScreenshotCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsConsole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneactionsConsole`: string
    fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsApi.StandaloneactionsConsole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmConsoleScreenshotCommand** | [**VmConsoleScreenshotCommand**](VmConsoleScreenshotCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneactionsDownloadRdp

> CsvExporter StandaloneactionsDownloadRdp(ctx, id).Execute()

Download RDP file

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneActionsApi.StandaloneactionsDownloadRdp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsDownloadRdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneactionsDownloadRdp`: CsvExporter
    fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsApi.StandaloneactionsDownloadRdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsDownloadRdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneactionsReboot

> StandaloneactionsReboot(ctx).RebootStandAloneVmCommand(rebootStandAloneVmCommand).Execute()

Reboot standalone vm

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
    rebootStandAloneVmCommand := *openapiclient.NewRebootStandAloneVmCommand() // RebootStandAloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneActionsApi.StandaloneactionsReboot(context.Background()).RebootStandAloneVmCommand(rebootStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsReboot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootStandAloneVmCommand** | [**RebootStandAloneVmCommand**](RebootStandAloneVmCommand.md) |  | 

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


## StandaloneactionsShelve

> StandaloneactionsShelve(ctx).ShelveStandAloneVmCommand(shelveStandAloneVmCommand).Execute()

Shelve standalone vm

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
    shelveStandAloneVmCommand := *openapiclient.NewShelveStandAloneVmCommand() // ShelveStandAloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneActionsApi.StandaloneactionsShelve(context.Background()).ShelveStandAloneVmCommand(shelveStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsShelve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsShelveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shelveStandAloneVmCommand** | [**ShelveStandAloneVmCommand**](ShelveStandAloneVmCommand.md) |  | 

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


## StandaloneactionsStart

> StandaloneactionsStart(ctx).StartStandaloneVmCommand(startStandaloneVmCommand).Execute()

Start standalone vm

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
    startStandaloneVmCommand := *openapiclient.NewStartStandaloneVmCommand() // StartStandaloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneActionsApi.StandaloneactionsStart(context.Background()).StartStandaloneVmCommand(startStandaloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startStandaloneVmCommand** | [**StartStandaloneVmCommand**](StartStandaloneVmCommand.md) |  | 

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


## StandaloneactionsStatus

> string StandaloneactionsStatus(ctx, id).Execute()

Show standalone vm status

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneActionsApi.StandaloneactionsStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneactionsStatus`: string
    fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsApi.StandaloneactionsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneactionsStop

> StandaloneactionsStop(ctx).StopStandaloneVmCommand(stopStandaloneVmCommand).Execute()

Stop standalone vm

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
    stopStandaloneVmCommand := *openapiclient.NewStopStandaloneVmCommand() // StopStandaloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneActionsApi.StandaloneactionsStop(context.Background()).StopStandaloneVmCommand(stopStandaloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopStandaloneVmCommand** | [**StopStandaloneVmCommand**](StopStandaloneVmCommand.md) |  | 

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


## StandaloneactionsUnshelve

> StandaloneactionsUnshelve(ctx).UnshelveStandaloneVmCommand(unshelveStandaloneVmCommand).Execute()

Unshelve standalone vm

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
    unshelveStandaloneVmCommand := *openapiclient.NewUnshelveStandaloneVmCommand() // UnshelveStandaloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneActionsApi.StandaloneactionsUnshelve(context.Background()).UnshelveStandaloneVmCommand(unshelveStandaloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsUnshelve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsUnshelveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unshelveStandaloneVmCommand** | [**UnshelveStandaloneVmCommand**](UnshelveStandaloneVmCommand.md) |  | 

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


## StandaloneactionsWindowsInstancePassword

> string StandaloneactionsWindowsInstancePassword(ctx).Id(id).Key(key).Execute()

Retrieve aws windows admin instance password

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
    id := int32(56) // int32 | 
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneActionsApi.StandaloneactionsWindowsInstancePassword(context.Background()).Id(id).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsApi.StandaloneactionsWindowsInstancePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneactionsWindowsInstancePassword`: string
    fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsApi.StandaloneactionsWindowsInstancePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsWindowsInstancePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **key** | **string** |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

