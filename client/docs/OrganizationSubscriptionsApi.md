# \OrganizationSubscriptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationSubscriptionsList**](OrganizationSubscriptionsApi.md#OrganizationSubscriptionsList) | **Get** /api/v{v}/OrganizationSubscriptions | Retrieve all org subscriptions



## OrganizationSubscriptionsList

> OrganizationSubscriptionList OrganizationSubscriptionsList(ctx, v).Offset(offset).Limit(limit).OrganizationId(organizationId).Search(search).Execute()

Retrieve all org subscriptions

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationSubscriptionsApi.OrganizationSubscriptionsList(context.Background(), v).Offset(offset).Limit(limit).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSubscriptionsApi.OrganizationSubscriptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationSubscriptionsList`: OrganizationSubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationSubscriptionsApi.OrganizationSubscriptionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSubscriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**OrganizationSubscriptionList**](OrganizationSubscriptionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

