# AdminUserCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Username** | **string** |  | 
**Password** | [**AdminUserCreateCommandPassword**](AdminUserCreateCommandPassword.md) |  | 
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminUserCreateCommand

`func NewAdminUserCreateCommand(email string, username string, password AdminUserCreateCommandPassword, ) *AdminUserCreateCommand`

NewAdminUserCreateCommand instantiates a new AdminUserCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserCreateCommandWithDefaults

`func NewAdminUserCreateCommandWithDefaults() *AdminUserCreateCommand`

NewAdminUserCreateCommandWithDefaults instantiates a new AdminUserCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminUserCreateCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserCreateCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserCreateCommand) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsername

`func (o *AdminUserCreateCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AdminUserCreateCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AdminUserCreateCommand) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AdminUserCreateCommand) GetPassword() AdminUserCreateCommandPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminUserCreateCommand) GetPasswordOk() (*AdminUserCreateCommandPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminUserCreateCommand) SetPassword(v AdminUserCreateCommandPassword)`

SetPassword sets Password field to given value.


### GetRole

`func (o *AdminUserCreateCommand) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminUserCreateCommand) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminUserCreateCommand) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminUserCreateCommand) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AdminUserCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AdminUserCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AdminUserCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AdminUserCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


