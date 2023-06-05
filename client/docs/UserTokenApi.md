# \UserTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserTokenAvailableEndpointList**](UserTokenApi.md#UserTokenAvailableEndpointList) | **Get** /api/v{v}/UserToken/available-endpoints | Get available endpoints list
[**UserTokenBindUnbind**](UserTokenApi.md#UserTokenBindUnbind) | **Post** /api/v{v}/UserToken/bind-unbind | Bind and unbind endpoints
[**UserTokenCreate**](UserTokenApi.md#UserTokenCreate) | **Post** /api/v{v}/UserToken/create | Create a new user token
[**UserTokenDelete**](UserTokenApi.md#UserTokenDelete) | **Delete** /api/v{v}/UserToken/delete/{id} | Delete token
[**UserTokenList**](UserTokenApi.md#UserTokenList) | **Get** /api/v{v}/UserToken/list | Get user token list



## UserTokenAvailableEndpointList

> AvailableEndpointsList UserTokenAvailableEndpointList(ctx, v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsAdd(isAdd).IsReadonly(isReadonly).Execute()

Get available endpoints list

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
    resp, r, err := apiClient.UserTokenApi.UserTokenAvailableEndpointList(context.Background(), v).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsAdd(isAdd).IsReadonly(isReadonly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UserTokenAvailableEndpointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTokenAvailableEndpointList`: AvailableEndpointsList
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UserTokenAvailableEndpointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenAvailableEndpointListRequest struct via the builder pattern


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
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokenBindUnbind

> UserTokenBindUnbind(ctx, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := *openapiclient.NewBindUnbindEndpointToTokenCommand() // BindUnbindEndpointToTokenCommand | Create command (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTokenApi.UserTokenBindUnbind(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UserTokenBindUnbind``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUserTokenBindUnbindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BindUnbindEndpointToTokenCommand**](BindUnbindEndpointToTokenCommand.md) | Create command | 

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


## UserTokenCreate

> UserTokenCreateDto UserTokenCreate(ctx, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := *openapiclient.NewUserTokenCreateCommand() // UserTokenCreateCommand | Create command (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTokenApi.UserTokenCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UserTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTokenCreate`: UserTokenCreateDto
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UserTokenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserTokenCreateCommand**](UserTokenCreateCommand.md) | Create command | 

### Return type

[**UserTokenCreateDto**](UserTokenCreateDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokenDelete

> UserTokenDelete(ctx, id, v).Execute()

Delete token

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
    id := "id_example" // string | Delete command
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTokenApi.UserTokenDelete(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UserTokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Delete command | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenDeleteRequest struct via the builder pattern


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


## UserTokenList

> []UserTokensListDto UserTokenList(ctx, v).Execute()

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
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTokenApi.UserTokenList(context.Background(), v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokenApi.UserTokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTokenList`: []UserTokensListDto
    fmt.Fprintf(os.Stdout, "Response from `UserTokenApi.UserTokenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserTokensListDto**](UserTokensListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

