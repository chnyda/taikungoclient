# \PreDefinedQueriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreDefinedQueriesCreatePrometheusDashboard**](PreDefinedQueriesApi.md#PreDefinedQueriesCreatePrometheusDashboard) | **Post** /api/v{v}/PreDefinedQueries/prometheus/dashboard/create | Create prometheus dashboard pre defined query
[**PreDefinedQueriesDeletePrometheusDashboard**](PreDefinedQueriesApi.md#PreDefinedQueriesDeletePrometheusDashboard) | **Delete** /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id} | Delete prometheus dashboard pre defined query
[**PreDefinedQueriesGetPrometheusCommonDashboardList**](PreDefinedQueriesApi.md#PreDefinedQueriesGetPrometheusCommonDashboardList) | **Get** /api/v{v}/PreDefinedQueries/prometheus/dashboard/common/{projectId} | Get list of pre defined common prometheus dashboard elements
[**PreDefinedQueriesGetPrometheusDashboardList**](PreDefinedQueriesApi.md#PreDefinedQueriesGetPrometheusDashboardList) | **Get** /api/v{v}/PreDefinedQueries/prometheus/dashboard/list/{projectId} | Get list of pre defined organization prometheus dashboard elements
[**PreDefinedQueriesUpdatePrometheusDashboard**](PreDefinedQueriesApi.md#PreDefinedQueriesUpdatePrometheusDashboard) | **Post** /api/v{v}/PreDefinedQueries/prometheus/dashboard/update | Update prometheus dashboard pre defined query



## PreDefinedQueriesCreatePrometheusDashboard

> PreDefinedQueriesCreatePrometheusDashboard(ctx, v).Body(body).Execute()

Create prometheus dashboard pre defined query

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
    body := *openapiclient.NewPrometheusDashboardCreateCommand("Name_example", "Expression_example", "Description_example", "CategoryName_example") // PrometheusDashboardCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PreDefinedQueriesApi.PreDefinedQueriesCreatePrometheusDashboard(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesApi.PreDefinedQueriesCreatePrometheusDashboard``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPreDefinedQueriesCreatePrometheusDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PrometheusDashboardCreateCommand**](PrometheusDashboardCreateCommand.md) |  | 

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


## PreDefinedQueriesDeletePrometheusDashboard

> PreDefinedQueriesDeletePrometheusDashboard(ctx, id, v).Execute()

Delete prometheus dashboard pre defined query

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
    r, err := apiClient.PreDefinedQueriesApi.PreDefinedQueriesDeletePrometheusDashboard(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesApi.PreDefinedQueriesDeletePrometheusDashboard``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPreDefinedQueriesDeletePrometheusDashboardRequest struct via the builder pattern


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


## PreDefinedQueriesGetPrometheusCommonDashboardList

> []PrometheusDashboardListDto PreDefinedQueriesGetPrometheusCommonDashboardList(ctx, projectId, v).Execute()

Get list of pre defined common prometheus dashboard elements

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
    resp, r, err := apiClient.PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusCommonDashboardList(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusCommonDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreDefinedQueriesGetPrometheusCommonDashboardList`: []PrometheusDashboardListDto
    fmt.Fprintf(os.Stdout, "Response from `PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusCommonDashboardList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreDefinedQueriesGetPrometheusCommonDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PrometheusDashboardListDto**](PrometheusDashboardListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreDefinedQueriesGetPrometheusDashboardList

> []PrometheusDashboardListDto PreDefinedQueriesGetPrometheusDashboardList(ctx, projectId, v).Execute()

Get list of pre defined organization prometheus dashboard elements

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
    resp, r, err := apiClient.PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusDashboardList(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreDefinedQueriesGetPrometheusDashboardList`: []PrometheusDashboardListDto
    fmt.Fprintf(os.Stdout, "Response from `PreDefinedQueriesApi.PreDefinedQueriesGetPrometheusDashboardList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreDefinedQueriesGetPrometheusDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PrometheusDashboardListDto**](PrometheusDashboardListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreDefinedQueriesUpdatePrometheusDashboard

> PreDefinedQueriesUpdatePrometheusDashboard(ctx, v).Body(body).Execute()

Update prometheus dashboard pre defined query

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
    body := *openapiclient.NewPrometheusDashboardUpdateCommand(int32(123), "Name_example", "Expression_example", "Description_example", "CategoryName_example") // PrometheusDashboardUpdateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PreDefinedQueriesApi.PreDefinedQueriesUpdatePrometheusDashboard(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesApi.PreDefinedQueriesUpdatePrometheusDashboard``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPreDefinedQueriesUpdatePrometheusDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PrometheusDashboardUpdateCommand**](PrometheusDashboardUpdateCommand.md) |  | 

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

