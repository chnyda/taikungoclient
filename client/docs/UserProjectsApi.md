# \UserProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserprojectsBindProjects**](UserProjectsApi.md#UserprojectsBindProjects) | **Post** /api/v1/userprojects/bindprojects | Bind projects to user
[**UserprojectsBindUsers**](UserProjectsApi.md#UserprojectsBindUsers) | **Post** /api/v1/userprojects/bindusers | Bind users to project
[**UserprojectsProjectListByUser**](UserProjectsApi.md#UserprojectsProjectListByUser) | **Get** /api/v1/userprojects/projects/list | Projects list for user
[**UserprojectsUserListByProject**](UserProjectsApi.md#UserprojectsUserListByProject) | **Get** /api/v1/userprojects/users/list/{projectId} | Users list by project id



## UserprojectsBindProjects

> UserprojectsBindProjects(ctx).BindProjectsCommand(bindProjectsCommand).Execute()

Bind projects to user

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
    bindProjectsCommand := *openapiclient.NewBindProjectsCommand() // BindProjectsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserProjectsApi.UserprojectsBindProjects(context.Background()).BindProjectsCommand(bindProjectsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProjectsApi.UserprojectsBindProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprojectsBindProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindProjectsCommand** | [**BindProjectsCommand**](BindProjectsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprojectsBindUsers

> UserprojectsBindUsers(ctx).BindUsersCommand(bindUsersCommand).Execute()

Bind users to project

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
    bindUsersCommand := *openapiclient.NewBindUsersCommand() // BindUsersCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserProjectsApi.UserprojectsBindUsers(context.Background()).BindUsersCommand(bindUsersCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProjectsApi.UserprojectsBindUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprojectsBindUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindUsersCommand** | [**BindUsersCommand**](BindUsersCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprojectsProjectListByUser

> []CommonDropdownDto UserprojectsProjectListByUser(ctx).Execute()

Projects list for user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProjectsApi.UserprojectsProjectListByUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProjectsApi.UserprojectsProjectListByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserprojectsProjectListByUser`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `UserProjectsApi.UserprojectsProjectListByUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserprojectsProjectListByUserRequest struct via the builder pattern


### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprojectsUserListByProject

> []CommonStringBasedDropdownDto UserprojectsUserListByProject(ctx, projectId).Execute()

Users list by project id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProjectsApi.UserprojectsUserListByProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProjectsApi.UserprojectsUserListByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserprojectsUserListByProject`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `UserProjectsApi.UserprojectsUserListByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserprojectsUserListByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

