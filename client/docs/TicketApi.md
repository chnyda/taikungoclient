# \TicketApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketArchive**](TicketApi.md#TicketArchive) | **Post** /api/v{v}/Ticket/archive | Archive ticket
[**TicketArticleList**](TicketApi.md#TicketArticleList) | **Get** /api/v{v}/Ticket/{ticketId}/messages | Retrieve articles of ticket.
[**TicketClose**](TicketApi.md#TicketClose) | **Post** /api/v{v}/Ticket/close | Close ticket
[**TicketCreate**](TicketApi.md#TicketCreate) | **Post** /api/v{v}/Ticket/create | Create ticket
[**TicketDelete**](TicketApi.md#TicketDelete) | **Delete** /api/v{v}/Ticket/delete/{ticketId} | Delete ticket
[**TicketDeleteMessage**](TicketApi.md#TicketDeleteMessage) | **Delete** /api/v{v}/Ticket/delete/message/{messageId} | Delete ticket message
[**TicketEdit**](TicketApi.md#TicketEdit) | **Post** /api/v{v}/Ticket/edit | Edit ticket
[**TicketOpen**](TicketApi.md#TicketOpen) | **Post** /api/v{v}/Ticket/open | Open ticket
[**TicketReply**](TicketApi.md#TicketReply) | **Post** /api/v{v}/Ticket/reply | Reply message
[**TicketSetPriority**](TicketApi.md#TicketSetPriority) | **Post** /api/v{v}/Ticket/set-priority | Set priority
[**TicketTicketList**](TicketApi.md#TicketTicketList) | **Get** /api/v{v}/Ticket/list | Retrieve a list of tickets.
[**TicketTransfer**](TicketApi.md#TicketTransfer) | **Post** /api/v{v}/Ticket/transfer | Transfer ticket
[**TicketTransferList**](TicketApi.md#TicketTransferList) | **Get** /api/v{v}/Ticket/transfer/list | Retrieve organization managers.
[**TicketUpdateMessage**](TicketApi.md#TicketUpdateMessage) | **Post** /api/v{v}/Ticket/edit/message | Edit ticket message



## TicketArchive

> TicketArchive(ctx, v).Body(body).Execute()

Archive ticket

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
    body := *openapiclient.NewArchiveTicketCommand("TicketId_example") // ArchiveTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketArchive(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ArchiveTicketCommand**](ArchiveTicketCommand.md) |  | 

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


## TicketArticleList

> ArticleList TicketArticleList(ctx, ticketId, v).Limit(limit).Offset(offset).Execute()

Retrieve articles of ticket.

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
    ticketId := "ticketId_example" // string | 
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.TicketArticleList(context.Background(), ticketId, v).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketArticleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketArticleList`: ArticleList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketArticleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketArticleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 

### Return type

[**ArticleList**](ArticleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketClose

> TicketClose(ctx, v).Body(body).Execute()

Close ticket

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
    body := *openapiclient.NewCloseTicketCommand("TicketId_example") // CloseTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketClose(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketClose``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CloseTicketCommand**](CloseTicketCommand.md) |  | 

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


## TicketCreate

> TicketCreate(ctx, v).Body(body).Execute()

Create ticket

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
    body := *openapiclient.NewCreateTicketCommand("Name_example", "Description_example") // CreateTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateTicketCommand**](CreateTicketCommand.md) |  | 

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


## TicketDelete

> TicketDelete(ctx, ticketId, v).Execute()

Delete ticket

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
    ticketId := "ticketId_example" // string | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketDelete(context.Background(), ticketId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteRequest struct via the builder pattern


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


## TicketDeleteMessage

> TicketDeleteMessage(ctx, messageId, v).Execute()

Delete ticket message

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
    messageId := "messageId_example" // string | 
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketDeleteMessage(context.Background(), messageId, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketDeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteMessageRequest struct via the builder pattern


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


## TicketEdit

> TicketEdit(ctx, v).Body(body).Execute()

Edit ticket

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
    body := *openapiclient.NewEditTicketCommand("TicketId_example", "Name_example", "Description_example") // EditTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketEdit(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EditTicketCommand**](EditTicketCommand.md) |  | 

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


## TicketOpen

> TicketOpen(ctx, v).Body(body).Execute()

Open ticket

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
    body := *openapiclient.NewOpenTicketCommand("TicketId_example") // OpenTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketOpen(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketOpen``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenTicketCommand**](OpenTicketCommand.md) |  | 

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


## TicketReply

> TicketReply(ctx, v).Body(body).Execute()

Reply message

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
    body := *openapiclient.NewReplyTicketCommand("TicketId_example", "Body_example") // ReplyTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketReply(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketReply``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplyTicketCommand**](ReplyTicketCommand.md) |  | 

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


## TicketSetPriority

> TicketSetPriority(ctx, v).Body(body).Execute()

Set priority

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
    body := *openapiclient.NewSetTicketPriorityCommand("Id_example") // SetTicketPriorityCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketSetPriority(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketSetPriority``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketSetPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetTicketPriorityCommand**](SetTicketPriorityCommand.md) |  | 

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


## TicketTicketList

> AllTicketsList TicketTicketList(ctx, v).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()

Retrieve a list of tickets.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    limit := int32(56) // int32 | Limits user size (by default 50) (optional)
    offset := int32(56) // int32 | Skip elements (optional)
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string | Keyword for searching (optional)
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)
    ticketId := "ticketId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.TicketTicketList(context.Background(), v).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketTicketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketTicketList`: AllTicketsList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketTicketList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketTicketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limits user size (by default 50) | 
 **offset** | **int32** | Skip elements | 
 **organizationId** | **int32** |  | 
 **search** | **string** | Keyword for searching | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **ticketId** | **string** |  | 

### Return type

[**AllTicketsList**](AllTicketsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketTransfer

> TicketTransfer(ctx, v).Body(body).Execute()

Transfer ticket

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
    body := *openapiclient.NewTransferTicketCommand("TicketId_example", "UserId_example") // TransferTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketTransfer(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketTransfer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TransferTicketCommand**](TransferTicketCommand.md) |  | 

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


## TicketTransferList

> []TransferList TicketTransferList(ctx, v).OrganizationId(organizationId).Execute()

Retrieve organization managers.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.TicketTransferList(context.Background(), v).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketTransferList`: []TransferList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketTransferList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 

### Return type

[**[]TransferList**](TransferList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketUpdateMessage

> TicketUpdateMessage(ctx, v).Body(body).Execute()

Edit ticket message

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
    body := *openapiclient.NewEditArticleCommand("MessageId_example", "Body_example") // EditArticleCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketUpdateMessage(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketUpdateMessage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTicketUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EditArticleCommand**](EditArticleCommand.md) |  | 

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

