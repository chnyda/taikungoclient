# ListForPartnersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectLimit** | Pointer to **int32** |  | [optional] 
**ServerLimit** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**CloudCredentialLimit** | Pointer to **int32** |  | [optional] 
**MonthlyPrice** | Pointer to **float64** |  | [optional] 
**YearlyPrice** | Pointer to **float64** |  | [optional] 
**TcuPrice** | Pointer to **float64** |  | [optional] 
**IsDeprecated** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**IsEnterprise** | Pointer to **bool** |  | [optional] 
**Partner** | Pointer to [**PartnerDetailsForSubscription**](PartnerDetailsForSubscription.md) |  | [optional] 
**ExceededUser** | Pointer to **bool** |  | [optional] 
**ExceededProject** | Pointer to **bool** |  | [optional] 
**ExceededCloudCredential** | Pointer to **bool** |  | [optional] 
**ExceededServers** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsYearly** | Pointer to **bool** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsDemo** | Pointer to **bool** |  | [optional] 

## Methods

### NewListForPartnersDto

`func NewListForPartnersDto() *ListForPartnersDto`

NewListForPartnersDto instantiates a new ListForPartnersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListForPartnersDtoWithDefaults

`func NewListForPartnersDtoWithDefaults() *ListForPartnersDto`

NewListForPartnersDtoWithDefaults instantiates a new ListForPartnersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListForPartnersDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListForPartnersDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListForPartnersDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListForPartnersDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListForPartnersDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListForPartnersDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListForPartnersDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListForPartnersDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectLimit

`func (o *ListForPartnersDto) GetProjectLimit() int32`

GetProjectLimit returns the ProjectLimit field if non-nil, zero value otherwise.

### GetProjectLimitOk

`func (o *ListForPartnersDto) GetProjectLimitOk() (*int32, bool)`

GetProjectLimitOk returns a tuple with the ProjectLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLimit

`func (o *ListForPartnersDto) SetProjectLimit(v int32)`

SetProjectLimit sets ProjectLimit field to given value.

### HasProjectLimit

`func (o *ListForPartnersDto) HasProjectLimit() bool`

HasProjectLimit returns a boolean if a field has been set.

### GetServerLimit

`func (o *ListForPartnersDto) GetServerLimit() int32`

GetServerLimit returns the ServerLimit field if non-nil, zero value otherwise.

### GetServerLimitOk

`func (o *ListForPartnersDto) GetServerLimitOk() (*int32, bool)`

GetServerLimitOk returns a tuple with the ServerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLimit

`func (o *ListForPartnersDto) SetServerLimit(v int32)`

SetServerLimit sets ServerLimit field to given value.

### HasServerLimit

`func (o *ListForPartnersDto) HasServerLimit() bool`

HasServerLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *ListForPartnersDto) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *ListForPartnersDto) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *ListForPartnersDto) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *ListForPartnersDto) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCloudCredentialLimit

`func (o *ListForPartnersDto) GetCloudCredentialLimit() int32`

GetCloudCredentialLimit returns the CloudCredentialLimit field if non-nil, zero value otherwise.

### GetCloudCredentialLimitOk

`func (o *ListForPartnersDto) GetCloudCredentialLimitOk() (*int32, bool)`

GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialLimit

`func (o *ListForPartnersDto) SetCloudCredentialLimit(v int32)`

SetCloudCredentialLimit sets CloudCredentialLimit field to given value.

### HasCloudCredentialLimit

`func (o *ListForPartnersDto) HasCloudCredentialLimit() bool`

HasCloudCredentialLimit returns a boolean if a field has been set.

### GetMonthlyPrice

`func (o *ListForPartnersDto) GetMonthlyPrice() float64`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *ListForPartnersDto) GetMonthlyPriceOk() (*float64, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *ListForPartnersDto) SetMonthlyPrice(v float64)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *ListForPartnersDto) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetYearlyPrice

`func (o *ListForPartnersDto) GetYearlyPrice() float64`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *ListForPartnersDto) GetYearlyPriceOk() (*float64, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *ListForPartnersDto) SetYearlyPrice(v float64)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *ListForPartnersDto) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.

### GetTcuPrice

`func (o *ListForPartnersDto) GetTcuPrice() float64`

GetTcuPrice returns the TcuPrice field if non-nil, zero value otherwise.

### GetTcuPriceOk

`func (o *ListForPartnersDto) GetTcuPriceOk() (*float64, bool)`

GetTcuPriceOk returns a tuple with the TcuPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcuPrice

`func (o *ListForPartnersDto) SetTcuPrice(v float64)`

SetTcuPrice sets TcuPrice field to given value.

### HasTcuPrice

`func (o *ListForPartnersDto) HasTcuPrice() bool`

HasTcuPrice returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *ListForPartnersDto) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ListForPartnersDto) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ListForPartnersDto) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *ListForPartnersDto) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetCurrency

`func (o *ListForPartnersDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListForPartnersDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListForPartnersDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListForPartnersDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIsEnterprise

`func (o *ListForPartnersDto) GetIsEnterprise() bool`

GetIsEnterprise returns the IsEnterprise field if non-nil, zero value otherwise.

### GetIsEnterpriseOk

`func (o *ListForPartnersDto) GetIsEnterpriseOk() (*bool, bool)`

GetIsEnterpriseOk returns a tuple with the IsEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterprise

`func (o *ListForPartnersDto) SetIsEnterprise(v bool)`

SetIsEnterprise sets IsEnterprise field to given value.

### HasIsEnterprise

`func (o *ListForPartnersDto) HasIsEnterprise() bool`

HasIsEnterprise returns a boolean if a field has been set.

### GetPartner

`func (o *ListForPartnersDto) GetPartner() PartnerDetailsForSubscription`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *ListForPartnersDto) GetPartnerOk() (*PartnerDetailsForSubscription, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *ListForPartnersDto) SetPartner(v PartnerDetailsForSubscription)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *ListForPartnersDto) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetExceededUser

`func (o *ListForPartnersDto) GetExceededUser() bool`

GetExceededUser returns the ExceededUser field if non-nil, zero value otherwise.

### GetExceededUserOk

`func (o *ListForPartnersDto) GetExceededUserOk() (*bool, bool)`

GetExceededUserOk returns a tuple with the ExceededUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededUser

`func (o *ListForPartnersDto) SetExceededUser(v bool)`

SetExceededUser sets ExceededUser field to given value.

### HasExceededUser

`func (o *ListForPartnersDto) HasExceededUser() bool`

HasExceededUser returns a boolean if a field has been set.

### GetExceededProject

`func (o *ListForPartnersDto) GetExceededProject() bool`

GetExceededProject returns the ExceededProject field if non-nil, zero value otherwise.

### GetExceededProjectOk

`func (o *ListForPartnersDto) GetExceededProjectOk() (*bool, bool)`

GetExceededProjectOk returns a tuple with the ExceededProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededProject

`func (o *ListForPartnersDto) SetExceededProject(v bool)`

SetExceededProject sets ExceededProject field to given value.

### HasExceededProject

`func (o *ListForPartnersDto) HasExceededProject() bool`

HasExceededProject returns a boolean if a field has been set.

### GetExceededCloudCredential

`func (o *ListForPartnersDto) GetExceededCloudCredential() bool`

GetExceededCloudCredential returns the ExceededCloudCredential field if non-nil, zero value otherwise.

### GetExceededCloudCredentialOk

`func (o *ListForPartnersDto) GetExceededCloudCredentialOk() (*bool, bool)`

GetExceededCloudCredentialOk returns a tuple with the ExceededCloudCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededCloudCredential

`func (o *ListForPartnersDto) SetExceededCloudCredential(v bool)`

SetExceededCloudCredential sets ExceededCloudCredential field to given value.

### HasExceededCloudCredential

`func (o *ListForPartnersDto) HasExceededCloudCredential() bool`

HasExceededCloudCredential returns a boolean if a field has been set.

### GetExceededServers

`func (o *ListForPartnersDto) GetExceededServers() bool`

GetExceededServers returns the ExceededServers field if non-nil, zero value otherwise.

### GetExceededServersOk

`func (o *ListForPartnersDto) GetExceededServersOk() (*bool, bool)`

GetExceededServersOk returns a tuple with the ExceededServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededServers

`func (o *ListForPartnersDto) SetExceededServers(v bool)`

SetExceededServers sets ExceededServers field to given value.

### HasExceededServers

`func (o *ListForPartnersDto) HasExceededServers() bool`

HasExceededServers returns a boolean if a field has been set.

### GetIsActive

`func (o *ListForPartnersDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListForPartnersDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListForPartnersDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ListForPartnersDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsYearly

`func (o *ListForPartnersDto) GetIsYearly() bool`

GetIsYearly returns the IsYearly field if non-nil, zero value otherwise.

### GetIsYearlyOk

`func (o *ListForPartnersDto) GetIsYearlyOk() (*bool, bool)`

GetIsYearlyOk returns a tuple with the IsYearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYearly

`func (o *ListForPartnersDto) SetIsYearly(v bool)`

SetIsYearly sets IsYearly field to given value.

### HasIsYearly

`func (o *ListForPartnersDto) HasIsYearly() bool`

HasIsYearly returns a boolean if a field has been set.

### GetTrialDays

`func (o *ListForPartnersDto) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *ListForPartnersDto) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *ListForPartnersDto) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *ListForPartnersDto) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetDescription

`func (o *ListForPartnersDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListForPartnersDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListForPartnersDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListForPartnersDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDemo

`func (o *ListForPartnersDto) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *ListForPartnersDto) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *ListForPartnersDto) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *ListForPartnersDto) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


