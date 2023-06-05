# \AlertingIntegrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingIntegrationsCreate**](AlertingIntegrationsApi.md#AlertingIntegrationsCreate) | **Post** /api/v{v}/AlertingIntegrations/create | Create alerting profile alerting integration
[**AlertingIntegrationsDelete**](AlertingIntegrationsApi.md#AlertingIntegrationsDelete) | **Delete** /api/v{v}/AlertingIntegrations/{id} | Delete alerting profile alerting integration
[**AlertingIntegrationsEdit**](AlertingIntegrationsApi.md#AlertingIntegrationsEdit) | **Put** /api/v{v}/AlertingIntegrations/edit | Edit alerting profile alerting integration
[**AlertingIntegrationsList**](AlertingIntegrationsApi.md#AlertingIntegrationsList) | **Get** /api/v{v}/AlertingIntegrations/{alertingProfileId} | List alerting integrations by profile id



## AlertingIntegrationsCreate

> ApiResponse AlertingIntegrationsCreate(ctx, v).Body(body).Execute()

Create alerting profile alerting integration

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
    body := *openapiclient.NewCreateAlertingIntegrationCommand("Url_example", openapiclient.AlertingIntegrationType(100), int32(123)) // CreateAlertingIntegrationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingIntegrationsApi.AlertingIntegrationsCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsApi.AlertingIntegrationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingIntegrationsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertingIntegrationsApi.AlertingIntegrationsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingIntegrationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateAlertingIntegrationCommand**](CreateAlertingIntegrationCommand.md) |  | 

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


## AlertingIntegrationsDelete

> AlertingIntegrationsDelete(ctx, id, v).Execute()

Delete alerting profile alerting integration

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
    r, err := apiClient.AlertingIntegrationsApi.AlertingIntegrationsDelete(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsApi.AlertingIntegrationsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingIntegrationsDeleteRequest struct via the builder pattern


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


## AlertingIntegrationsEdit

> AlertingIntegrationsEdit(ctx, v).Body(body).Execute()

Edit alerting profile alerting integration

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
    body := *openapiclient.NewEditAlertingIntegrationCommand(int32(123), "Url_example", openapiclient.AlertingIntegrationType(100), int32(123)) // EditAlertingIntegrationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingIntegrationsApi.AlertingIntegrationsEdit(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsApi.AlertingIntegrationsEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingIntegrationsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EditAlertingIntegrationCommand**](EditAlertingIntegrationCommand.md) |  | 

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


## AlertingIntegrationsList

> []AlertingIntegrationsListDto AlertingIntegrationsList(ctx, alertingProfileId, v).Search(search).Execute()

List alerting integrations by profile id

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
    alertingProfileId := int32(56) // int32 | 
    v := "v_example" // string | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingIntegrationsApi.AlertingIntegrationsList(context.Background(), alertingProfileId, v).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsApi.AlertingIntegrationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingIntegrationsList`: []AlertingIntegrationsListDto
    fmt.Fprintf(os.Stdout, "Response from `AlertingIntegrationsApi.AlertingIntegrationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertingProfileId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingIntegrationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **string** |  | 

### Return type

[**[]AlertingIntegrationsListDto**](AlertingIntegrationsListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

