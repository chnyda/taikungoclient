# \UserGroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupsBindProjectsGroup**](UserGroupApi.md#UsergroupsBindProjectsGroup) | **Post** /api/v1/usergroups/bind-project-groups | Bind project groups
[**UsergroupsCreate**](UserGroupApi.md#UsergroupsCreate) | **Post** /api/v1/usergroups/create | Add user groups
[**UsergroupsDelete**](UserGroupApi.md#UsergroupsDelete) | **Delete** /api/v1/usergroups | Remove user group(s)
[**UsergroupsList**](UserGroupApi.md#UsergroupsList) | **Get** /api/v1/usergroups/list | Retrieve list of user groups
[**UsergroupsListByProjectGroupId**](UserGroupApi.md#UsergroupsListByProjectGroupId) | **Get** /api/v1/usergroups/list-by-project-group-id/{projectGroupId} | Dropdown list
[**UsergroupsUpdate**](UserGroupApi.md#UsergroupsUpdate) | **Put** /api/v1/usergroups/update/{userGroupId} | Update user groups
[**UsergroupsUserGroupUsers**](UserGroupApi.md#UsergroupsUserGroupUsers) | **Get** /api/v1/usergroups/users/{userGroupId} | Dropdown list



## UsergroupsBindProjectsGroup

> UsergroupsBindProjectsGroup(ctx).BindProjectGroupsToUserGroupCommand(bindProjectGroupsToUserGroupCommand).Execute()

Bind project groups

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
    bindProjectGroupsToUserGroupCommand := *openapiclient.NewBindProjectGroupsToUserGroupCommand() // BindProjectGroupsToUserGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserGroupApi.UsergroupsBindProjectsGroup(context.Background()).BindProjectGroupsToUserGroupCommand(bindProjectGroupsToUserGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsBindProjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsBindProjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindProjectGroupsToUserGroupCommand** | [**BindProjectGroupsToUserGroupCommand**](BindProjectGroupsToUserGroupCommand.md) |  | 

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


## UsergroupsCreate

> ApiResponse UsergroupsCreate(ctx).CreateUserGroupCommand(createUserGroupCommand).Execute()

Add user groups

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
    createUserGroupCommand := *openapiclient.NewCreateUserGroupCommand("Name_example") // CreateUserGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupApi.UsergroupsCreate(context.Background()).CreateUserGroupCommand(createUserGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.UsergroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserGroupCommand** | [**CreateUserGroupCommand**](CreateUserGroupCommand.md) |  | 

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


## UsergroupsDelete

> UsergroupsDelete(ctx).DeleteUserGroupCommand(deleteUserGroupCommand).Execute()

Remove user group(s)

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
    deleteUserGroupCommand := *openapiclient.NewDeleteUserGroupCommand() // DeleteUserGroupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserGroupApi.UsergroupsDelete(context.Background()).DeleteUserGroupCommand(deleteUserGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserGroupCommand** | [**DeleteUserGroupCommand**](DeleteUserGroupCommand.md) |  | 

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


## UsergroupsList

> UserGroupList UsergroupsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of user groups

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupApi.UsergroupsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupsList`: UserGroupList
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.UsergroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsListRequest struct via the builder pattern


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

[**UserGroupList**](UserGroupList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupsListByProjectGroupId

> []CommonDropdownDto UsergroupsListByProjectGroupId(ctx, projectGroupId).Execute()

Dropdown list

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
    projectGroupId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupApi.UsergroupsListByProjectGroupId(context.Background(), projectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsListByProjectGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupsListByProjectGroupId`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.UsergroupsListByProjectGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsListByProjectGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UsergroupsUpdate

> UsergroupsUpdate(ctx, userGroupId).UpdateUserGroupDto(updateUserGroupDto).Execute()

Update user groups

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
    userGroupId := int32(56) // int32 | 
    updateUserGroupDto := *openapiclient.NewUpdateUserGroupDto() // UpdateUserGroupDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserGroupApi.UsergroupsUpdate(context.Background(), userGroupId).UpdateUserGroupDto(updateUserGroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserGroupDto** | [**UpdateUserGroupDto**](UpdateUserGroupDto.md) |  | 

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


## UsergroupsUserGroupUsers

> []UserListForUserGroupDto UsergroupsUserGroupUsers(ctx, userGroupId).OrganizationId(organizationId).Execute()

Dropdown list

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
    userGroupId := int32(56) // int32 | 
    organizationId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupApi.UsergroupsUserGroupUsers(context.Background(), userGroupId).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UsergroupsUserGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsergroupsUserGroupUsers`: []UserListForUserGroupDto
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.UsergroupsUserGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupsUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 

### Return type

[**[]UserListForUserGroupDto**](UserListForUserGroupDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

