# \StandaloneVMDisksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandalonevmdisksCreate**](StandaloneVMDisksApi.md#StandalonevmdisksCreate) | **Post** /api/v1/standalonevmdisks/create | Add disk for standalone vm
[**StandalonevmdisksDelete**](StandaloneVMDisksApi.md#StandalonevmdisksDelete) | **Post** /api/v1/standalonevmdisks/delete | Remove disk from standalone vm
[**StandalonevmdisksPurge**](StandaloneVMDisksApi.md#StandalonevmdisksPurge) | **Post** /api/v1/standalonevmdisks/purge | Purge vm disks
[**StandalonevmdisksReset**](StandaloneVMDisksApi.md#StandalonevmdisksReset) | **Post** /api/v1/standalonevmdisks/reset | Update status of disk
[**StandalonevmdisksUpdate**](StandaloneVMDisksApi.md#StandalonevmdisksUpdate) | **Post** /api/v1/standalonevmdisks/update | Update disk
[**StandalonevmdisksUpdateSize**](StandaloneVMDisksApi.md#StandalonevmdisksUpdateSize) | **Post** /api/v1/standalonevmdisks/update-size | Update disk size



## StandalonevmdisksCreate

> ApiResponse StandalonevmdisksCreate(ctx).CreateStandAloneDiskCommand(createStandAloneDiskCommand).Execute()

Add disk for standalone vm

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
    createStandAloneDiskCommand := *openapiclient.NewCreateStandAloneDiskCommand() // CreateStandAloneDiskCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksCreate(context.Background()).CreateStandAloneDiskCommand(createStandAloneDiskCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandalonevmdisksCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `StandaloneVMDisksApi.StandalonevmdisksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandAloneDiskCommand** | [**CreateStandAloneDiskCommand**](CreateStandAloneDiskCommand.md) |  | 

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


## StandalonevmdisksDelete

> StandalonevmdisksDelete(ctx).DeleteStandAloneVmDiskCommand(deleteStandAloneVmDiskCommand).Execute()

Remove disk from standalone vm

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
    deleteStandAloneVmDiskCommand := *openapiclient.NewDeleteStandAloneVmDiskCommand() // DeleteStandAloneVmDiskCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksDelete(context.Background()).DeleteStandAloneVmDiskCommand(deleteStandAloneVmDiskCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteStandAloneVmDiskCommand** | [**DeleteStandAloneVmDiskCommand**](DeleteStandAloneVmDiskCommand.md) |  | 

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


## StandalonevmdisksPurge

> StandalonevmdisksPurge(ctx).PurgeStandAloneVmDiskCommand(purgeStandAloneVmDiskCommand).Execute()

Purge vm disks

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
    purgeStandAloneVmDiskCommand := *openapiclient.NewPurgeStandAloneVmDiskCommand() // PurgeStandAloneVmDiskCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksPurge(context.Background()).PurgeStandAloneVmDiskCommand(purgeStandAloneVmDiskCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purgeStandAloneVmDiskCommand** | [**PurgeStandAloneVmDiskCommand**](PurgeStandAloneVmDiskCommand.md) |  | 

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


## StandalonevmdisksReset

> StandalonevmdisksReset(ctx).ResetStandAloneVmDiskStatusCommand(resetStandAloneVmDiskStatusCommand).Execute()

Update status of disk

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
    resetStandAloneVmDiskStatusCommand := *openapiclient.NewResetStandAloneVmDiskStatusCommand() // ResetStandAloneVmDiskStatusCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksReset(context.Background()).ResetStandAloneVmDiskStatusCommand(resetStandAloneVmDiskStatusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetStandAloneVmDiskStatusCommand** | [**ResetStandAloneVmDiskStatusCommand**](ResetStandAloneVmDiskStatusCommand.md) |  | 

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


## StandalonevmdisksUpdate

> StandalonevmdisksUpdate(ctx).UpdateStandaloneVmDiskCommand(updateStandaloneVmDiskCommand).Execute()

Update disk

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
    updateStandaloneVmDiskCommand := *openapiclient.NewUpdateStandaloneVmDiskCommand() // UpdateStandaloneVmDiskCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksUpdate(context.Background()).UpdateStandaloneVmDiskCommand(updateStandaloneVmDiskCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandaloneVmDiskCommand** | [**UpdateStandaloneVmDiskCommand**](UpdateStandaloneVmDiskCommand.md) |  | 

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


## StandalonevmdisksUpdateSize

> StandalonevmdisksUpdateSize(ctx).UpdateStandaloneVmDiskSizeCommand(updateStandaloneVmDiskSizeCommand).Execute()

Update disk size

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
    updateStandaloneVmDiskSizeCommand := *openapiclient.NewUpdateStandaloneVmDiskSizeCommand() // UpdateStandaloneVmDiskSizeCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksUpdateSize(context.Background()).UpdateStandaloneVmDiskSizeCommand(updateStandaloneVmDiskSizeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksApi.StandalonevmdisksUpdateSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksUpdateSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandaloneVmDiskSizeCommand** | [**UpdateStandaloneVmDiskSizeCommand**](UpdateStandaloneVmDiskSizeCommand.md) |  | 

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

