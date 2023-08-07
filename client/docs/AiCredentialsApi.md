# \AiCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiCredentialCreate**](AiCredentialsApi.md#AiCredentialCreate) | **Post** /api/v1/ai-credential/create | Create ai credential
[**AiCredentialDropdown**](AiCredentialsApi.md#AiCredentialDropdown) | **Get** /api/v1/ai-credential | Retrieve all AI credentials for organization
[**AiCredentialList**](AiCredentialsApi.md#AiCredentialList) | **Get** /api/v1/ai-credential/list | Retrieve all AI credentials



## AiCredentialCreate

> ApiResponse AiCredentialCreate(ctx).CreateAiCredentialCommand(createAiCredentialCommand).Execute()

Create ai credential

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
    createAiCredentialCommand := *openapiclient.NewCreateAiCredentialCommand() // CreateAiCredentialCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AiCredentialsApi.AiCredentialCreate(context.Background()).CreateAiCredentialCommand(createAiCredentialCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsApi.AiCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AiCredentialCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AiCredentialsApi.AiCredentialCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAiCredentialCommand** | [**CreateAiCredentialCommand**](CreateAiCredentialCommand.md) |  | 

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


## AiCredentialDropdown

> []AiCredentialsForOrganizationEntity AiCredentialDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all AI credentials for organization

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
    resp, r, err := apiClient.AiCredentialsApi.AiCredentialDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsApi.AiCredentialDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AiCredentialDropdown`: []AiCredentialsForOrganizationEntity
    fmt.Fprintf(os.Stdout, "Response from `AiCredentialsApi.AiCredentialDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]AiCredentialsForOrganizationEntity**](AiCredentialsForOrganizationEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCredentialList

> AiCredentials AiCredentialList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve all AI credentials

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
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AiCredentialsApi.AiCredentialList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsApi.AiCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AiCredentialList`: AiCredentials
    fmt.Fprintf(os.Stdout, "Response from `AiCredentialsApi.AiCredentialList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AiCredentials**](AiCredentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

