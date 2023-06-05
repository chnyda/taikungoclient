# PartnerDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PaymentEnabled** | Pointer to **bool** |  | [optional] 
**AllowRegistration** | Pointer to **bool** |  | [optional] 
**RequiredUserApproval** | Pointer to **bool** |  | [optional] 
**Organizations** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**WhiteListDomains** | Pointer to [**[]WhiteListDomainDto**](WhiteListDomainDto.md) |  | [optional] 

## Methods

### NewPartnerDetailsDto

`func NewPartnerDetailsDto() *PartnerDetailsDto`

NewPartnerDetailsDto instantiates a new PartnerDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerDetailsDtoWithDefaults

`func NewPartnerDetailsDtoWithDefaults() *PartnerDetailsDto`

NewPartnerDetailsDtoWithDefaults instantiates a new PartnerDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnerDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PartnerDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PartnerDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLink

`func (o *PartnerDetailsDto) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PartnerDetailsDto) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PartnerDetailsDto) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PartnerDetailsDto) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDomain

`func (o *PartnerDetailsDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PartnerDetailsDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PartnerDetailsDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PartnerDetailsDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCountry

`func (o *PartnerDetailsDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PartnerDetailsDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PartnerDetailsDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PartnerDetailsDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *PartnerDetailsDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PartnerDetailsDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PartnerDetailsDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PartnerDetailsDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetVatNumber

`func (o *PartnerDetailsDto) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *PartnerDetailsDto) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *PartnerDetailsDto) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *PartnerDetailsDto) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetAddress

`func (o *PartnerDetailsDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartnerDetailsDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartnerDetailsDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartnerDetailsDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLogo

`func (o *PartnerDetailsDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PartnerDetailsDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PartnerDetailsDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PartnerDetailsDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPhone

`func (o *PartnerDetailsDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartnerDetailsDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartnerDetailsDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartnerDetailsDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *PartnerDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartnerDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartnerDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartnerDetailsDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPaymentEnabled

`func (o *PartnerDetailsDto) GetPaymentEnabled() bool`

GetPaymentEnabled returns the PaymentEnabled field if non-nil, zero value otherwise.

### GetPaymentEnabledOk

`func (o *PartnerDetailsDto) GetPaymentEnabledOk() (*bool, bool)`

GetPaymentEnabledOk returns a tuple with the PaymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentEnabled

`func (o *PartnerDetailsDto) SetPaymentEnabled(v bool)`

SetPaymentEnabled sets PaymentEnabled field to given value.

### HasPaymentEnabled

`func (o *PartnerDetailsDto) HasPaymentEnabled() bool`

HasPaymentEnabled returns a boolean if a field has been set.

### GetAllowRegistration

`func (o *PartnerDetailsDto) GetAllowRegistration() bool`

GetAllowRegistration returns the AllowRegistration field if non-nil, zero value otherwise.

### GetAllowRegistrationOk

`func (o *PartnerDetailsDto) GetAllowRegistrationOk() (*bool, bool)`

GetAllowRegistrationOk returns a tuple with the AllowRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRegistration

`func (o *PartnerDetailsDto) SetAllowRegistration(v bool)`

SetAllowRegistration sets AllowRegistration field to given value.

### HasAllowRegistration

`func (o *PartnerDetailsDto) HasAllowRegistration() bool`

HasAllowRegistration returns a boolean if a field has been set.

### GetRequiredUserApproval

`func (o *PartnerDetailsDto) GetRequiredUserApproval() bool`

GetRequiredUserApproval returns the RequiredUserApproval field if non-nil, zero value otherwise.

### GetRequiredUserApprovalOk

`func (o *PartnerDetailsDto) GetRequiredUserApprovalOk() (*bool, bool)`

GetRequiredUserApprovalOk returns a tuple with the RequiredUserApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredUserApproval

`func (o *PartnerDetailsDto) SetRequiredUserApproval(v bool)`

SetRequiredUserApproval sets RequiredUserApproval field to given value.

### HasRequiredUserApproval

`func (o *PartnerDetailsDto) HasRequiredUserApproval() bool`

HasRequiredUserApproval returns a boolean if a field has been set.

### GetOrganizations

`func (o *PartnerDetailsDto) GetOrganizations() []CommonDropdownDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *PartnerDetailsDto) GetOrganizationsOk() (*[]CommonDropdownDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *PartnerDetailsDto) SetOrganizations(v []CommonDropdownDto)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *PartnerDetailsDto) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetWhiteListDomains

`func (o *PartnerDetailsDto) GetWhiteListDomains() []WhiteListDomainDto`

GetWhiteListDomains returns the WhiteListDomains field if non-nil, zero value otherwise.

### GetWhiteListDomainsOk

`func (o *PartnerDetailsDto) GetWhiteListDomainsOk() (*[]WhiteListDomainDto, bool)`

GetWhiteListDomainsOk returns a tuple with the WhiteListDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListDomains

`func (o *PartnerDetailsDto) SetWhiteListDomains(v []WhiteListDomainDto)`

SetWhiteListDomains sets WhiteListDomains field to given value.

### HasWhiteListDomains

`func (o *PartnerDetailsDto) HasWhiteListDomains() bool`

HasWhiteListDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


