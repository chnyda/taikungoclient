# OrganizationCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**BillingEmail** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**DiscountRate** | Pointer to **float64** |  | [optional] 
**IsEligibleUpdateSubscription** | Pointer to **bool** |  | [optional] 
**AdminCloudCredentialId** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationCreateCommand

`func NewOrganizationCreateCommand(name string, fullName string, ) *OrganizationCreateCommand`

NewOrganizationCreateCommand instantiates a new OrganizationCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateCommandWithDefaults

`func NewOrganizationCreateCommandWithDefaults() *OrganizationCreateCommand`

NewOrganizationCreateCommandWithDefaults instantiates a new OrganizationCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *OrganizationCreateCommand) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationCreateCommand) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationCreateCommand) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetPhone

`func (o *OrganizationCreateCommand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationCreateCommand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationCreateCommand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationCreateCommand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationCreateCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationCreateCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationCreateCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationCreateCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBillingEmail

`func (o *OrganizationCreateCommand) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrganizationCreateCommand) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrganizationCreateCommand) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrganizationCreateCommand) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetAddress

`func (o *OrganizationCreateCommand) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationCreateCommand) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationCreateCommand) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganizationCreateCommand) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountry

`func (o *OrganizationCreateCommand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationCreateCommand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationCreateCommand) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationCreateCommand) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *OrganizationCreateCommand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationCreateCommand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationCreateCommand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationCreateCommand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetVatNumber

`func (o *OrganizationCreateCommand) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *OrganizationCreateCommand) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *OrganizationCreateCommand) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *OrganizationCreateCommand) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetDiscountRate

`func (o *OrganizationCreateCommand) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *OrganizationCreateCommand) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *OrganizationCreateCommand) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *OrganizationCreateCommand) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) GetIsEligibleUpdateSubscription() bool`

GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field if non-nil, zero value otherwise.

### GetIsEligibleUpdateSubscriptionOk

`func (o *OrganizationCreateCommand) GetIsEligibleUpdateSubscriptionOk() (*bool, bool)`

GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) SetIsEligibleUpdateSubscription(v bool)`

SetIsEligibleUpdateSubscription sets IsEligibleUpdateSubscription field to given value.

### HasIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) HasIsEligibleUpdateSubscription() bool`

HasIsEligibleUpdateSubscription returns a boolean if a field has been set.

### GetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialId() int32`

GetAdminCloudCredentialId returns the AdminCloudCredentialId field if non-nil, zero value otherwise.

### GetAdminCloudCredentialIdOk

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialIdOk() (*int32, bool)`

GetAdminCloudCredentialIdOk returns a tuple with the AdminCloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) SetAdminCloudCredentialId(v int32)`

SetAdminCloudCredentialId sets AdminCloudCredentialId field to given value.

### HasAdminCloudCredentialId

`func (o *OrganizationCreateCommand) HasAdminCloudCredentialId() bool`

HasAdminCloudCredentialId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


