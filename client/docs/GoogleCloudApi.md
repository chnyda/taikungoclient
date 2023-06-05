# \GoogleCloudApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GoogleCloudBillingAccountList**](GoogleCloudApi.md#GoogleCloudBillingAccountList) | **Post** /api/v{v}/GoogleCloud/billing-accounts | Retrieve google billing accounts list
[**GoogleCloudCreate**](GoogleCloudApi.md#GoogleCloudCreate) | **Post** /api/v{v}/GoogleCloud/create | Create google cloud credential
[**GoogleCloudList**](GoogleCloudApi.md#GoogleCloudList) | **Get** /api/v{v}/GoogleCloud/list | Retrieve list of google cloud credentials
[**GoogleCloudRegionList**](GoogleCloudApi.md#GoogleCloudRegionList) | **Post** /api/v{v}/GoogleCloud/regions | Retrieve google region list
[**GoogleCloudZoneList**](GoogleCloudApi.md#GoogleCloudZoneList) | **Post** /api/v{v}/GoogleCloud/zones | Retrieve google zone list



## GoogleCloudBillingAccountList

> []CommonStringBasedDropdownDto GoogleCloudBillingAccountList(ctx, v).Config(config).Execute()

Retrieve google billing accounts list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleCloudApi.GoogleCloudBillingAccountList(context.Background(), v).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleCloudApi.GoogleCloudBillingAccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleCloudBillingAccountList`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `GoogleCloudApi.GoogleCloudBillingAccountList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleCloudBillingAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | ***os.File** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoogleCloudCreate

> ApiResponse GoogleCloudCreate(ctx, v).Name(name).Config(config).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).Execute()

Create google cloud credential

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
    name := "name_example" // string |  (optional)
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)
    importProject := true // bool |  (optional)
    folderId := "folderId_example" // string |  (optional)
    billingAccountId := "billingAccountId_example" // string |  (optional)
    azCount := int32(56) // int32 |  (optional)
    region := "region_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleCloudApi.GoogleCloudCreate(context.Background(), v).Name(name).Config(config).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleCloudApi.GoogleCloudCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleCloudCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `GoogleCloudApi.GoogleCloudCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleCloudCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **config** | ***os.File** |  | 
 **importProject** | **bool** |  | 
 **folderId** | **string** |  | 
 **billingAccountId** | **string** |  | 
 **azCount** | **int32** |  | 
 **region** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoogleCloudList

> GoogleCredentialList GoogleCloudList(ctx, v).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of google cloud credentials

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
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleCloudApi.GoogleCloudList(context.Background(), v).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleCloudApi.GoogleCloudList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleCloudList`: GoogleCredentialList
    fmt.Fprintf(os.Stdout, "Response from `GoogleCloudApi.GoogleCloudList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleCloudListRequest struct via the builder pattern


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

[**GoogleCredentialList**](GoogleCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoogleCloudRegionList

> []string GoogleCloudRegionList(ctx, v).Config(config).Execute()

Retrieve google region list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleCloudApi.GoogleCloudRegionList(context.Background(), v).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleCloudApi.GoogleCloudRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleCloudRegionList`: []string
    fmt.Fprintf(os.Stdout, "Response from `GoogleCloudApi.GoogleCloudRegionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleCloudRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | ***os.File** |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoogleCloudZoneList

> AzResult GoogleCloudZoneList(ctx, v).Config(config).Region(region).CloudId(cloudId).Execute()

Retrieve google zone list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)
    region := "region_example" // string |  (optional)
    cloudId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleCloudApi.GoogleCloudZoneList(context.Background(), v).Config(config).Region(region).CloudId(cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleCloudApi.GoogleCloudZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleCloudZoneList`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `GoogleCloudApi.GoogleCloudZoneList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleCloudZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | ***os.File** |  | 
 **region** | **string** |  | 
 **cloudId** | **int32** |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

