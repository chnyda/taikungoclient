# \DnsServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsServersCreate**](DnsServersApi.md#DnsServersCreate) | **Post** /api/v{v}/DnsServers/create | Create access profile dns server
[**DnsServersDelete**](DnsServersApi.md#DnsServersDelete) | **Delete** /api/v{v}/DnsServers/{id} | Delete access profile dns server
[**DnsServersEdit**](DnsServersApi.md#DnsServersEdit) | **Put** /api/v{v}/DnsServers/edit/{id} | Edit access profile dns server
[**DnsServersList**](DnsServersApi.md#DnsServersList) | **Get** /api/v{v}/DnsServers/list/{accessProfileId} | List dns servers by profile id



## DnsServersCreate

> ApiResponse DnsServersCreate(ctx, v).Body(body).Execute()

Create access profile dns server

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
    body := *openapiclient.NewCreateDnsServerCommand("Address_example", int32(123)) // CreateDnsServerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersApi.DnsServersCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsServersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsServersCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsServersApi.DnsServersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateDnsServerCommand**](CreateDnsServerCommand.md) |  | 

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


## DnsServersDelete

> DnsServersDelete(ctx, id, v).Execute()

Delete access profile dns server

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
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsServersApi.DnsServersDelete(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsServersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsServersDeleteRequest struct via the builder pattern


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


## DnsServersEdit

> DnsServersEdit(ctx, id, v).Body(body).Execute()

Edit access profile dns server

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
    v := "v_example" // string | 
    body := *openapiclient.NewDnsNtpAddressEditDto() // DnsNtpAddressEditDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsServersApi.DnsServersEdit(context.Background(), id, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsServersEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsServersEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DnsNtpAddressEditDto**](DnsNtpAddressEditDto.md) |  | 

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


## DnsServersList

> []DnsServersListDto DnsServersList(ctx, accessProfileId, v).Search(search).Execute()

List dns servers by profile id

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
    accessProfileId := int32(56) // int32 | 
    v := "v_example" // string | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersApi.DnsServersList(context.Background(), accessProfileId, v).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsServersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsServersList`: []DnsServersListDto
    fmt.Fprintf(os.Stdout, "Response from `DnsServersApi.DnsServersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **string** |  | 

### Return type

[**[]DnsServersListDto**](DnsServersListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

