# \UserTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsertokenAvailableEndpoints**](UserTokenApi.md#UsertokenAvailableEndpoints) | **Get** /api/v1/usertoken/available-endpoints | Get available endpoint list
[**UsertokenBindUnbind**](UserTokenApi.md#UsertokenBindUnbind) | **Post** /api/v1/usertoken/bind-unbind | Bind and unbind endpoints
[**UsertokenCreate**](UserTokenApi.md#UsertokenCreate) | **Post** /api/v1/usertoken/create | Create a new user token
[**UsertokenDelete**](UserTokenApi.md#UsertokenDelete) | **Delete** /api/v1/usertoken/delete/{id} | 
[**UsertokenList**](UserTokenApi.md#UsertokenList) | **Get** /api/v1/usertoken/list | Get user token list



## UsertokenAvailableEndpoints

> AvailableEndpointsList UsertokenAvailableEndpoints(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsAdd(isAdd).IsReadonly(isReadonly).Execute()

Get available endpoint list

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := "id_example" // string |  (optional)
    isAdd := true // bool |  (optional)
    isReadonly := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTokenApi.UsertokenAvailableEndpoints(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsAdd(isAdd).IsReadonly(isReadonly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UsertokenAvailableEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsertokenAvailableEndpoints`: AvailableEndpointsList
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UsertokenAvailableEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsertokenAvailableEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **string** |  | 
 **isAdd** | **bool** |  | 
 **isReadonly** | **bool** |  | 

### Return type

[**AvailableEndpointsList**](AvailableEndpointsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsertokenBindUnbind

> UsertokenBindUnbind(ctx).BindUnbindEndpointToTokenCommand(bindUnbindEndpointToTokenCommand).Execute()

Bind and unbind endpoints

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
    bindUnbindEndpointToTokenCommand := *openapiclient.NewBindUnbindEndpointToTokenCommand() // BindUnbindEndpointToTokenCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTokenApi.UsertokenBindUnbind(context.Background()).BindUnbindEndpointToTokenCommand(bindUnbindEndpointToTokenCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UsertokenBindUnbind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsertokenBindUnbindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindUnbindEndpointToTokenCommand** | [**BindUnbindEndpointToTokenCommand**](BindUnbindEndpointToTokenCommand.md) |  | 

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


## UsertokenCreate

> UserTokenCreateDto UsertokenCreate(ctx).UserTokenCreateCommand(userTokenCreateCommand).Execute()

Create a new user token

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
    userTokenCreateCommand := *openapiclient.NewUserTokenCreateCommand() // UserTokenCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTokenApi.UsertokenCreate(context.Background()).UserTokenCreateCommand(userTokenCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UsertokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsertokenCreate`: UserTokenCreateDto
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UsertokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsertokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userTokenCreateCommand** | [**UserTokenCreateCommand**](UserTokenCreateCommand.md) |  | 

### Return type

[**UserTokenCreateDto**](UserTokenCreateDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsertokenDelete

> UsertokenDelete(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTokenApi.UsertokenDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UsertokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsertokenDeleteRequest struct via the builder pattern


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


## UsertokenList

> []UserTokensListDto UsertokenList(ctx).Execute()

Get user token list

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
    resp, r, err := apiClient.UserTokenApi.UsertokenList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UsertokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsertokenList`: []UserTokensListDto
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UsertokenList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsertokenListRequest struct via the builder pattern


### Return type

[**[]UserTokensListDto**](UserTokensListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

