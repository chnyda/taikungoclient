# \ProjectInfracostsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectInfracostsDelete**](ProjectInfracostsApi.md#ProjectInfracostsDelete) | **Post** /api/v{v}/ProjectInfracosts/delete | Delete the project infracost.
[**ProjectInfracostsDetails**](ProjectInfracostsApi.md#ProjectInfracostsDetails) | **Get** /api/v{v}/ProjectInfracosts/{projectId} | Project Infracost details
[**ProjectInfracostsEdit**](ProjectInfracostsApi.md#ProjectInfracostsEdit) | **Post** /api/v{v}/ProjectInfracosts/upsert/{projectId} | Upsert project infracost by ProjectId



## ProjectInfracostsDelete

> ProjectInfracostsDelete(ctx, v).Body(body).Execute()

Delete the project infracost.

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
    body := *openapiclient.NewDeleteProjectInfracostCommand() // DeleteProjectInfracostCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectInfracostsApi.ProjectInfracostsDelete(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsApi.ProjectInfracostsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectInfracostsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteProjectInfracostCommand**](DeleteProjectInfracostCommand.md) |  | 

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


## ProjectInfracostsDetails

> EstimatedInfracost ProjectInfracostsDetails(ctx, projectId, v).Execute()

Project Infracost details

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
    projectId := int32(56) // int32 | Project Id
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectInfracostsApi.ProjectInfracostsDetails(context.Background(), projectId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsApi.ProjectInfracostsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectInfracostsDetails`: EstimatedInfracost
    fmt.Fprintf(os.Stdout, "Response from `ProjectInfracostsApi.ProjectInfracostsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project Id | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectInfracostsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EstimatedInfracost**](EstimatedInfracost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectInfracostsEdit

> ProjectInfracostsEdit(ctx, projectId, v).Body(body).Execute()

Upsert project infracost by ProjectId

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
    body := *openapiclient.NewProjectInfracostUpsertDto() // ProjectInfracostUpsertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectInfracostsApi.ProjectInfracostsEdit(context.Background(), projectId, v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsApi.ProjectInfracostsEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectInfracostsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProjectInfracostUpsertDto**](ProjectInfracostUpsertDto.md) |  | 

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

