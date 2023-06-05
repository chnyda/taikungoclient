# CreateSubscriptionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ProjectLimit** | Pointer to **int32** |  | [optional] 
**ServerLimit** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**CloudCredentialLimit** | Pointer to **int32** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**MonthlyPrice** | Pointer to **float64** |  | [optional] 
**TcuPrice** | Pointer to **float64** |  | [optional] 
**YearlyPrice** | Pointer to **float64** |  | [optional] 

## Methods

### NewCreateSubscriptionCommand

`func NewCreateSubscriptionCommand(name string, ) *CreateSubscriptionCommand`

NewCreateSubscriptionCommand instantiates a new CreateSubscriptionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionCommandWithDefaults

`func NewCreateSubscriptionCommandWithDefaults() *CreateSubscriptionCommand`

NewCreateSubscriptionCommandWithDefaults instantiates a new CreateSubscriptionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSubscriptionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubscriptionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubscriptionCommand) SetName(v string)`

SetName sets Name field to given value.


### GetProjectLimit

`func (o *CreateSubscriptionCommand) GetProjectLimit() int32`

GetProjectLimit returns the ProjectLimit field if non-nil, zero value otherwise.

### GetProjectLimitOk

`func (o *CreateSubscriptionCommand) GetProjectLimitOk() (*int32, bool)`

GetProjectLimitOk returns a tuple with the ProjectLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLimit

`func (o *CreateSubscriptionCommand) SetProjectLimit(v int32)`

SetProjectLimit sets ProjectLimit field to given value.

### HasProjectLimit

`func (o *CreateSubscriptionCommand) HasProjectLimit() bool`

HasProjectLimit returns a boolean if a field has been set.

### GetServerLimit

`func (o *CreateSubscriptionCommand) GetServerLimit() int32`

GetServerLimit returns the ServerLimit field if non-nil, zero value otherwise.

### GetServerLimitOk

`func (o *CreateSubscriptionCommand) GetServerLimitOk() (*int32, bool)`

GetServerLimitOk returns a tuple with the ServerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLimit

`func (o *CreateSubscriptionCommand) SetServerLimit(v int32)`

SetServerLimit sets ServerLimit field to given value.

### HasServerLimit

`func (o *CreateSubscriptionCommand) HasServerLimit() bool`

HasServerLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *CreateSubscriptionCommand) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *CreateSubscriptionCommand) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *CreateSubscriptionCommand) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *CreateSubscriptionCommand) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCloudCredentialLimit

`func (o *CreateSubscriptionCommand) GetCloudCredentialLimit() int32`

GetCloudCredentialLimit returns the CloudCredentialLimit field if non-nil, zero value otherwise.

### GetCloudCredentialLimitOk

`func (o *CreateSubscriptionCommand) GetCloudCredentialLimitOk() (*int32, bool)`

GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialLimit

`func (o *CreateSubscriptionCommand) SetCloudCredentialLimit(v int32)`

SetCloudCredentialLimit sets CloudCredentialLimit field to given value.

### HasCloudCredentialLimit

`func (o *CreateSubscriptionCommand) HasCloudCredentialLimit() bool`

HasCloudCredentialLimit returns a boolean if a field has been set.

### GetTrialDays

`func (o *CreateSubscriptionCommand) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *CreateSubscriptionCommand) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *CreateSubscriptionCommand) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *CreateSubscriptionCommand) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetMonthlyPrice

`func (o *CreateSubscriptionCommand) GetMonthlyPrice() float64`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *CreateSubscriptionCommand) GetMonthlyPriceOk() (*float64, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *CreateSubscriptionCommand) SetMonthlyPrice(v float64)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *CreateSubscriptionCommand) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetTcuPrice

`func (o *CreateSubscriptionCommand) GetTcuPrice() float64`

GetTcuPrice returns the TcuPrice field if non-nil, zero value otherwise.

### GetTcuPriceOk

`func (o *CreateSubscriptionCommand) GetTcuPriceOk() (*float64, bool)`

GetTcuPriceOk returns a tuple with the TcuPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcuPrice

`func (o *CreateSubscriptionCommand) SetTcuPrice(v float64)`

SetTcuPrice sets TcuPrice field to given value.

### HasTcuPrice

`func (o *CreateSubscriptionCommand) HasTcuPrice() bool`

HasTcuPrice returns a boolean if a field has been set.

### GetYearlyPrice

`func (o *CreateSubscriptionCommand) GetYearlyPrice() float64`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *CreateSubscriptionCommand) GetYearlyPriceOk() (*float64, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *CreateSubscriptionCommand) SetYearlyPrice(v float64)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *CreateSubscriptionCommand) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


