# \KubeConfigRoleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubeConfigRoleList**](KubeConfigRoleApi.md#KubeConfigRoleList) | **Get** /api/v{v}/KubeConfigRole | Retrieve list of kubeconfig role



## KubeConfigRoleList

> KubeConfigRoleResponse KubeConfigRoleList(ctx, v).Search(search).Execute()

Retrieve list of kubeconfig role

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubeConfigRoleApi.KubeConfigRoleList(context.Background(), v).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigRoleApi.KubeConfigRoleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubeConfigRoleList`: KubeConfigRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `KubeConfigRoleApi.KubeConfigRoleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubeConfigRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**KubeConfigRoleResponse**](KubeConfigRoleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

