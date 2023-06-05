# ResetPasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**NewPassword** | [**AdminUserCreateCommandPassword**](AdminUserCreateCommandPassword.md) |  | 

## Methods

### NewResetPasswordCommand

`func NewResetPasswordCommand(token string, newPassword AdminUserCreateCommandPassword, ) *ResetPasswordCommand`

NewResetPasswordCommand instantiates a new ResetPasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordCommandWithDefaults

`func NewResetPasswordCommandWithDefaults() *ResetPasswordCommand`

NewResetPasswordCommandWithDefaults instantiates a new ResetPasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResetPasswordCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetPasswordCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetPasswordCommand) SetToken(v string)`

SetToken sets Token field to given value.


### GetEmail

`func (o *ResetPasswordCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResetPasswordCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNewPassword

`func (o *ResetPasswordCommand) GetNewPassword() AdminUserCreateCommandPassword`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordCommand) GetNewPasswordOk() (*AdminUserCreateCommandPassword, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordCommand) SetNewPassword(v AdminUserCreateCommandPassword)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


