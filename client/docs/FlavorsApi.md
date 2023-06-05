# \FlavorsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlavorsAwsFlavors**](FlavorsApi.md#FlavorsAwsFlavors) | **Get** /api/v{v}/Flavors/aws/{cloudId} | Retrieve aws flavors
[**FlavorsAzureFlavors**](FlavorsApi.md#FlavorsAzureFlavors) | **Get** /api/v{v}/Flavors/azure/{cloudId} | Retrieve azure flavors
[**FlavorsBindToProject**](FlavorsApi.md#FlavorsBindToProject) | **Post** /api/v{v}/Flavors/bind | Bind flavors to project
[**FlavorsDropdownRecordDtos**](FlavorsApi.md#FlavorsDropdownRecordDtos) | **Get** /api/v{v}/Flavors/credentials/dropdown/list | Retrieve cloud credentials dropdown list
[**FlavorsGetSelectedFlavorsForProject**](FlavorsApi.md#FlavorsGetSelectedFlavorsForProject) | **Get** /api/v{v}/Flavors/projects/list | Retrieve selected flavors for projects
[**FlavorsGoogleFlavors**](FlavorsApi.md#FlavorsGoogleFlavors) | **Get** /api/v{v}/Flavors/google/{cloudId} | Retrieve google flavors
[**FlavorsOpenstackFlavors**](FlavorsApi.md#FlavorsOpenstackFlavors) | **Get** /api/v{v}/Flavors/openstack/{cloudId} | Retrieve openstack flavors
[**FlavorsProxmoxFlavors**](FlavorsApi.md#FlavorsProxmoxFlavors) | **Get** /api/v{v}/Flavors/proxmox/{cloudId} | Retrieve proxmox flavors
[**FlavorsTanzuFlavors**](FlavorsApi.md#FlavorsTanzuFlavors) | **Get** /api/v{v}/Flavors/tanzu/{cloudId} | Retrieve tanzu flavors
[**FlavorsUnbindFromProject**](FlavorsApi.md#FlavorsUnbindFromProject) | **Post** /api/v{v}/Flavors/unbind | Unbind flavors from project



## FlavorsAwsFlavors

> AwsFlavorList FlavorsAwsFlavors(ctx, cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve aws flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    startRam := float64(1.2) // float64 |  (optional)
    endRam := float64(1.2) // float64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsAwsFlavors(context.Background(), cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsAwsFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsAwsFlavors`: AwsFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsAwsFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAwsFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits size (by default 50) | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AwsFlavorList**](AwsFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsAzureFlavors

> AzureFlavorList FlavorsAzureFlavors(ctx, cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve azure flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    startRam := int32(56) // int32 |  (optional)
    endRam := int32(56) // int32 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsAzureFlavors(context.Background(), cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsAzureFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsAzureFlavors`: AzureFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsAzureFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAzureFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits size (by default 50) | 
 **startRam** | **int32** |  | 
 **endRam** | **int32** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AzureFlavorList**](AzureFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsBindToProject

> FlavorsBindToProject(ctx, v).Body(body).Execute()

Bind flavors to project

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
    body := *openapiclient.NewBindFlavorToProjectCommand() // BindFlavorToProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FlavorsApi.FlavorsBindToProject(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsBindToProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlavorsBindToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BindFlavorToProjectCommand**](BindFlavorToProjectCommand.md) |  | 

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


## FlavorsDropdownRecordDtos

> []CloudCredentialsDropdownRecordDto FlavorsDropdownRecordDtos(ctx, v).OrganizationId(organizationId).FilterBy(filterBy).Search(search).Execute()

Retrieve cloud credentials dropdown list

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
    organizationId := int32(56) // int32 |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsDropdownRecordDtos(context.Background(), v).OrganizationId(organizationId).FilterBy(filterBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsDropdownRecordDtos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsDropdownRecordDtos`: []CloudCredentialsDropdownRecordDto
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsDropdownRecordDtos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsDropdownRecordDtosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]CloudCredentialsDropdownRecordDto**](CloudCredentialsDropdownRecordDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsGetSelectedFlavorsForProject

> BoundFlavorsForProjectsList FlavorsGetSelectedFlavorsForProject(ctx, v).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).WithPrice(withPrice).Execute()

Retrieve selected flavors for projects

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
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    flavorName := "flavorName_example" // string |  (optional)
    withPrice := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsGetSelectedFlavorsForProject(context.Background(), v).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).WithPrice(withPrice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsGetSelectedFlavorsForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsGetSelectedFlavorsForProject`: BoundFlavorsForProjectsList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsGetSelectedFlavorsForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsGetSelectedFlavorsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 
 **flavorName** | **string** |  | 
 **withPrice** | **bool** |  | 

### Return type

[**BoundFlavorsForProjectsList**](BoundFlavorsForProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsGoogleFlavors

> GoogleFlavorList FlavorsGoogleFlavors(ctx, cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve google flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    startRam := float64(1.2) // float64 |  (optional)
    endRam := float64(1.2) // float64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsGoogleFlavors(context.Background(), cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsGoogleFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsGoogleFlavors`: GoogleFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsGoogleFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsGoogleFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits size (by default 50) | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**GoogleFlavorList**](GoogleFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsOpenstackFlavors

> OpenstackFlavorList FlavorsOpenstackFlavors(ctx, cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve openstack flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    startRam := float64(1.2) // float64 |  (optional)
    endRam := float64(1.2) // float64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsOpenstackFlavors(context.Background(), cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsOpenstackFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsOpenstackFlavors`: OpenstackFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsOpenstackFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsOpenstackFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits size (by default 50) | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**OpenstackFlavorList**](OpenstackFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsProxmoxFlavors

> ProxmoxFlavorList FlavorsProxmoxFlavors(ctx, cloudId, v).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve proxmox flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    startRam := int64(789) // int64 |  (optional)
    endRam := int64(789) // int64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsProxmoxFlavors(context.Background(), cloudId, v).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsProxmoxFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsProxmoxFlavors`: ProxmoxFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsProxmoxFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsProxmoxFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **startRam** | **int64** |  | 
 **endRam** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**ProxmoxFlavorList**](ProxmoxFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsTanzuFlavors

> TanzuFlavorList FlavorsTanzuFlavors(ctx, cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve tanzu flavors

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
    cloudId := int32(56) // int32 | 
    v := "v_example" // string | 
    offset := int32(56) // int32 | Skip elements (optional)
    limit := int32(56) // int32 | Limits size (by default 50) (optional)
    startRam := int32(56) // int32 |  (optional)
    endRam := int32(56) // int32 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsTanzuFlavors(context.Background(), cloudId, v).Offset(offset).Limit(limit).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsTanzuFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsTanzuFlavors`: TanzuFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsTanzuFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsTanzuFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Skip elements | 
 **limit** | **int32** | Limits size (by default 50) | 
 **startRam** | **int32** |  | 
 **endRam** | **int32** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**TanzuFlavorList**](TanzuFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsUnbindFromProject

> FlavorsUnbindFromProject(ctx, v).Body(body).Execute()

Unbind flavors from project

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
    body := *openapiclient.NewUnbindFlavorFromProjectCommand() // UnbindFlavorFromProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FlavorsApi.FlavorsUnbindFromProject(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsUnbindFromProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlavorsUnbindFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnbindFlavorFromProjectCommand**](UnbindFlavorFromProjectCommand.md) |  | 

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

