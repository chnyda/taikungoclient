# UpdateUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**ForceToResetPassword** | Pointer to **bool** |  | [optional] 
**Disable** | Pointer to **bool** |  | [optional] 
**IsApprovedByPartner** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateUserCommand

`func NewUpdateUserCommand() *UpdateUserCommand`

NewUpdateUserCommand instantiates a new UpdateUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserCommandWithDefaults

`func NewUpdateUserCommandWithDefaults() *UpdateUserCommand`

NewUpdateUserCommandWithDefaults instantiates a new UpdateUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateUserCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserCommand) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUserCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateUserCommand) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateUserCommand) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDisplayName

`func (o *UpdateUserCommand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateUserCommand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateUserCommand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateUserCommand) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateUserCommand) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateUserCommand) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUsername

`func (o *UpdateUserCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserCommand) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserCommand) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateUserCommand) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateUserCommand) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *UpdateUserCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateUserCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateUserCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRole

`func (o *UpdateUserCommand) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateUserCommand) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateUserCommand) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateUserCommand) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetForceToResetPassword

`func (o *UpdateUserCommand) GetForceToResetPassword() bool`

GetForceToResetPassword returns the ForceToResetPassword field if non-nil, zero value otherwise.

### GetForceToResetPasswordOk

`func (o *UpdateUserCommand) GetForceToResetPasswordOk() (*bool, bool)`

GetForceToResetPasswordOk returns a tuple with the ForceToResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceToResetPassword

`func (o *UpdateUserCommand) SetForceToResetPassword(v bool)`

SetForceToResetPassword sets ForceToResetPassword field to given value.

### HasForceToResetPassword

`func (o *UpdateUserCommand) HasForceToResetPassword() bool`

HasForceToResetPassword returns a boolean if a field has been set.

### GetDisable

`func (o *UpdateUserCommand) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *UpdateUserCommand) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *UpdateUserCommand) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *UpdateUserCommand) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetIsApprovedByPartner

`func (o *UpdateUserCommand) GetIsApprovedByPartner() bool`

GetIsApprovedByPartner returns the IsApprovedByPartner field if non-nil, zero value otherwise.

### GetIsApprovedByPartnerOk

`func (o *UpdateUserCommand) GetIsApprovedByPartnerOk() (*bool, bool)`

GetIsApprovedByPartnerOk returns a tuple with the IsApprovedByPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovedByPartner

`func (o *UpdateUserCommand) SetIsApprovedByPartner(v bool)`

SetIsApprovedByPartner sets IsApprovedByPartner field to given value.

### HasIsApprovedByPartner

`func (o *UpdateUserCommand) HasIsApprovedByPartner() bool`

HasIsApprovedByPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


