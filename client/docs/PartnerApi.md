# \PartnerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnerAddWhiteListDomain**](PartnerApi.md#PartnerAddWhiteListDomain) | **Post** /api/v{v}/Partner/add/whitelist/domain | 
[**PartnerBecomePartner**](PartnerApi.md#PartnerBecomePartner) | **Post** /api/v{v}/Partner/become-a-partner | 
[**PartnerBindOrganizations**](PartnerApi.md#PartnerBindOrganizations) | **Post** /api/v{v}/Partner/bindorganizations | 
[**PartnerContactUs**](PartnerApi.md#PartnerContactUs) | **Post** /api/v{v}/Partner/contact-us | 
[**PartnerCreate**](PartnerApi.md#PartnerCreate) | **Post** /api/v{v}/Partner/create | Add partner
[**PartnerDeleteWhiteListDomain**](PartnerApi.md#PartnerDeleteWhiteListDomain) | **Post** /api/v{v}/Partner/delete/whitelist/domain | 
[**PartnerDetails**](PartnerApi.md#PartnerDetails) | **Get** /api/v{v}/Partner/details | Details of partners
[**PartnerList**](PartnerApi.md#PartnerList) | **Get** /api/v{v}/Partner | Get partners
[**PartnerPartnerInfoRegistration**](PartnerApi.md#PartnerPartnerInfoRegistration) | **Get** /api/v{v}/Partner/info | 
[**PartnerPartnerList**](PartnerApi.md#PartnerPartnerList) | **Get** /api/v{v}/Partner/list | Get partners dropdown
[**PartnerUpdate**](PartnerApi.md#PartnerUpdate) | **Put** /api/v{v}/Partner/update/{id} | Edit partner data by Id



## PartnerAddWhiteListDomain

> PartnerAddWhiteListDomain(ctx, v).Body(body).Execute()



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
    body := *openapiclient.NewWhiteListDomainCreateCommand() // WhiteListDomainCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerAddWhiteListDomain(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerAddWhiteListDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerAddWhiteListDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WhiteListDomainCreateCommand**](WhiteListDomainCreateCommand.md) |  | 

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


## PartnerBecomePartner

> PartnerBecomePartner(ctx, v).Body(body).Execute()



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
    body := *openapiclient.NewBecomePartnerCommand("FullName_example", "Email_example") // BecomePartnerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerBecomePartner(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerBecomePartner``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerBecomePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BecomePartnerCommand**](BecomePartnerCommand.md) |  | 

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


## PartnerBindOrganizations

> PartnerBindOrganizations(ctx, v).Body(body).Execute()



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
    body := *openapiclient.NewBindOrganizationsCommand(int32(123)) // BindOrganizationsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerBindOrganizations(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerBindOrganizations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerBindOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BindOrganizationsCommand**](BindOrganizationsCommand.md) |  | 

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


## PartnerContactUs

> PartnerContactUs(ctx, v).Body(body).Execute()



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
    body := *openapiclient.NewContactUsCommand("Name_example", "BusinessEmail_example", "CompanyName_example") // ContactUsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerContactUs(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerContactUs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerContactUsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ContactUsCommand**](ContactUsCommand.md) |  | 

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


## PartnerCreate

> PartnerCreate(ctx, v).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Execute()

Add partner

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
    domain := "domain_example" // string |  (optional)
    link := "link_example" // string |  (optional)
    phone := "phone_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    vatNumber := "vatNumber_example" // string |  (optional)
    address := "address_example" // string |  (optional)
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)
    backgroundImage := os.NewFile(1234, "some_file") // *os.File |  (optional)
    allowRegistration := true // bool |  (optional)
    requiredUserApproval := true // bool |  (optional)
    paymentEnabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerCreate(context.Background(), v).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **domain** | **string** |  | 
 **link** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **country** | **string** |  | 
 **city** | **string** |  | 
 **vatNumber** | **string** |  | 
 **address** | **string** |  | 
 **logo** | ***os.File** |  | 
 **backgroundImage** | ***os.File** |  | 
 **allowRegistration** | **bool** |  | 
 **requiredUserApproval** | **bool** |  | 
 **paymentEnabled** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerDeleteWhiteListDomain

> PartnerDeleteWhiteListDomain(ctx, v).Body(body).Execute()



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
    body := *openapiclient.NewWhiteListDomainDeleteCommand() // WhiteListDomainDeleteCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerDeleteWhiteListDomain(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerDeleteWhiteListDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerDeleteWhiteListDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WhiteListDomainDeleteCommand**](WhiteListDomainDeleteCommand.md) |  | 

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


## PartnerDetails

> PartnerDetailsDto PartnerDetails(ctx, v).Execute()

Details of partners

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
    resp, r, err := apiClient.PartnerApi.PartnerDetails(context.Background(), v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerDetails`: PartnerDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `PartnerApi.PartnerDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartnerDetailsDto**](PartnerDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerList

> PartnersList PartnerList(ctx, v).Offset(offset).Limit(limit).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()

Get partners

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerApi.PartnerList(context.Background(), v).Offset(offset).Limit(limit).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerList`: PartnersList
    fmt.Fprintf(os.Stdout, "Response from `PartnerApi.PartnerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 

### Return type

[**PartnersList**](PartnersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerPartnerInfoRegistration

> PartnerRecordDto PartnerPartnerInfoRegistration(ctx, v).Domain(domain).Execute()



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
    domain := "domain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerApi.PartnerPartnerInfoRegistration(context.Background(), v).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerPartnerInfoRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerPartnerInfoRegistration`: PartnerRecordDto
    fmt.Fprintf(os.Stdout, "Response from `PartnerApi.PartnerPartnerInfoRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerPartnerInfoRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | **string** |  | 

### Return type

[**PartnerRecordDto**](PartnerRecordDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerPartnerList

> []PartnerEntity PartnerPartnerList(ctx, v).Search(search).Execute()

Get partners dropdown

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
    resp, r, err := apiClient.PartnerApi.PartnerPartnerList(context.Background(), v).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerPartnerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerPartnerList`: []PartnerEntity
    fmt.Fprintf(os.Stdout, "Response from `PartnerApi.PartnerPartnerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerPartnerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]PartnerEntity**](PartnerEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerUpdate

> PartnerUpdate(ctx, id, v).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Execute()

Edit partner data by Id

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
    name := "name_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    link := "link_example" // string |  (optional)
    phone := "phone_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    vatNumber := "vatNumber_example" // string |  (optional)
    address := "address_example" // string |  (optional)
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)
    backgroundImage := os.NewFile(1234, "some_file") // *os.File |  (optional)
    allowRegistration := true // bool |  (optional)
    requiredUserApproval := true // bool |  (optional)
    paymentEnabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnerApi.PartnerUpdate(context.Background(), id, v).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApi.PartnerUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** |  | 
 **domain** | **string** |  | 
 **link** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **country** | **string** |  | 
 **city** | **string** |  | 
 **vatNumber** | **string** |  | 
 **address** | **string** |  | 
 **logo** | ***os.File** |  | 
 **backgroundImage** | ***os.File** |  | 
 **allowRegistration** | **bool** |  | 
 **requiredUserApproval** | **bool** |  | 
 **paymentEnabled** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

