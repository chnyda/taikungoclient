# \StandaloneProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandaloneprofileCreate**](StandaloneProfileApi.md#StandaloneprofileCreate) | **Post** /api/v1/standaloneprofile/create | Create standalone profile.
[**StandaloneprofileDelete**](StandaloneProfileApi.md#StandaloneprofileDelete) | **Post** /api/v1/standaloneprofile/delete | Delete standalone profile.
[**StandaloneprofileDropdown**](StandaloneProfileApi.md#StandaloneprofileDropdown) | **Get** /api/v1/standaloneprofile/list | Retrieve all standalone profiles for organization
[**StandaloneprofileEdit**](StandaloneProfileApi.md#StandaloneprofileEdit) | **Post** /api/v1/standaloneprofile/edit | Edit standalone profile.
[**StandaloneprofileList**](StandaloneProfileApi.md#StandaloneprofileList) | **Get** /api/v1/standaloneprofile | Retrieve all standalone profiles
[**StandaloneprofileLockManagement**](StandaloneProfileApi.md#StandaloneprofileLockManagement) | **Post** /api/v1/standaloneprofile/lockmanager | Lock/unlock standalone profile.



## StandaloneprofileCreate

> ApiResponse StandaloneprofileCreate(ctx).StandAloneProfileCreateCommand(standAloneProfileCreateCommand).Execute()

Create standalone profile.

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
    standAloneProfileCreateCommand := *openapiclient.NewStandAloneProfileCreateCommand() // StandAloneProfileCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneProfileApi.StandaloneprofileCreate(context.Background()).StandAloneProfileCreateCommand(standAloneProfileCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneprofileCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `StandaloneProfileApi.StandaloneprofileCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneProfileCreateCommand** | [**StandAloneProfileCreateCommand**](StandAloneProfileCreateCommand.md) |  | 

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


## StandaloneprofileDelete

> StandaloneprofileDelete(ctx).DeleteStandAloneProfileCommand(deleteStandAloneProfileCommand).Execute()

Delete standalone profile.

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
    deleteStandAloneProfileCommand := *openapiclient.NewDeleteStandAloneProfileCommand() // DeleteStandAloneProfileCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneProfileApi.StandaloneprofileDelete(context.Background()).DeleteStandAloneProfileCommand(deleteStandAloneProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteStandAloneProfileCommand** | [**DeleteStandAloneProfileCommand**](DeleteStandAloneProfileCommand.md) |  | 

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


## StandaloneprofileDropdown

> []CommonDropdownDto StandaloneprofileDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all standalone profiles for organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneProfileApi.StandaloneprofileDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneprofileDropdown`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `StandaloneProfileApi.StandaloneprofileDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneprofileEdit

> StandaloneprofileEdit(ctx).StandAloneProfileUpdateCommand(standAloneProfileUpdateCommand).Execute()

Edit standalone profile.

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
    standAloneProfileUpdateCommand := *openapiclient.NewStandAloneProfileUpdateCommand() // StandAloneProfileUpdateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneProfileApi.StandaloneprofileEdit(context.Background()).StandAloneProfileUpdateCommand(standAloneProfileUpdateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneProfileUpdateCommand** | [**StandAloneProfileUpdateCommand**](StandAloneProfileUpdateCommand.md) |  | 

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


## StandaloneprofileList

> StandAloneProfiles StandaloneprofileList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all standalone profiles

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
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneProfileApi.StandaloneprofileList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneprofileList`: StandAloneProfiles
    fmt.Fprintf(os.Stdout, "Response from `StandaloneProfileApi.StandaloneprofileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**StandAloneProfiles**](StandAloneProfiles.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneprofileLockManagement

> StandaloneprofileLockManagement(ctx).StandAloneProfileLockManagementCommand(standAloneProfileLockManagementCommand).Execute()

Lock/unlock standalone profile.

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
    standAloneProfileLockManagementCommand := *openapiclient.NewStandAloneProfileLockManagementCommand() // StandAloneProfileLockManagementCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneProfileApi.StandaloneprofileLockManagement(context.Background()).StandAloneProfileLockManagementCommand(standAloneProfileLockManagementCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneProfileApi.StandaloneprofileLockManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneprofileLockManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneProfileLockManagementCommand** | [**StandAloneProfileLockManagementCommand**](StandAloneProfileLockManagementCommand.md) |  | 

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

