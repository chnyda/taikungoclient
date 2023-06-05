# OrganizationDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**BillingEmail** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**IsEligibleUpdateSubscription** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 
**Projects** | Pointer to **int32** |  | [optional] 
**Servers** | Pointer to **int32** |  | [optional] 
**CloudCredentials** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**PartnerName** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**PartnerDetailsForOrganizationsDto**](PartnerDetailsForOrganizationsDto.md) |  | [optional] 
**DiscountRate** | Pointer to **float64** |  | [optional] 
**BoundRules** | Pointer to [**[]PrometheusEntity**](PrometheusEntity.md) |  | [optional] 

## Methods

### NewOrganizationDetailsDto

`func NewOrganizationDetailsDto() *OrganizationDetailsDto`

NewOrganizationDetailsDto instantiates a new OrganizationDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsDtoWithDefaults

`func NewOrganizationDetailsDtoWithDefaults() *OrganizationDetailsDto`

NewOrganizationDetailsDtoWithDefaults instantiates a new OrganizationDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullName

`func (o *OrganizationDetailsDto) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationDetailsDto) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationDetailsDto) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OrganizationDetailsDto) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationDetailsDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBillingEmail

`func (o *OrganizationDetailsDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrganizationDetailsDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrganizationDetailsDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrganizationDetailsDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetPhone

`func (o *OrganizationDetailsDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationDetailsDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationDetailsDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationDetailsDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCountry

`func (o *OrganizationDetailsDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationDetailsDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationDetailsDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationDetailsDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *OrganizationDetailsDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationDetailsDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationDetailsDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationDetailsDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetVatNumber

`func (o *OrganizationDetailsDto) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *OrganizationDetailsDto) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *OrganizationDetailsDto) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *OrganizationDetailsDto) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetAddress

`func (o *OrganizationDetailsDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationDetailsDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationDetailsDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganizationDetailsDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsEligibleUpdateSubscription

`func (o *OrganizationDetailsDto) GetIsEligibleUpdateSubscription() bool`

GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field if non-nil, zero value otherwise.

### GetIsEligibleUpdateSubscriptionOk

`func (o *OrganizationDetailsDto) GetIsEligibleUpdateSubscriptionOk() (*bool, bool)`

GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleUpdateSubscription

`func (o *OrganizationDetailsDto) SetIsEligibleUpdateSubscription(v bool)`

SetIsEligibleUpdateSubscription sets IsEligibleUpdateSubscription field to given value.

### HasIsEligibleUpdateSubscription

`func (o *OrganizationDetailsDto) HasIsEligibleUpdateSubscription() bool`

HasIsEligibleUpdateSubscription returns a boolean if a field has been set.

### GetIsLocked

`func (o *OrganizationDetailsDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OrganizationDetailsDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OrganizationDetailsDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *OrganizationDetailsDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *OrganizationDetailsDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *OrganizationDetailsDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *OrganizationDetailsDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *OrganizationDetailsDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetUsers

`func (o *OrganizationDetailsDto) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OrganizationDetailsDto) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OrganizationDetailsDto) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OrganizationDetailsDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetProjects

`func (o *OrganizationDetailsDto) GetProjects() int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationDetailsDto) GetProjectsOk() (*int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationDetailsDto) SetProjects(v int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationDetailsDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetServers

`func (o *OrganizationDetailsDto) GetServers() int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OrganizationDetailsDto) GetServersOk() (*int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OrganizationDetailsDto) SetServers(v int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OrganizationDetailsDto) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetCloudCredentials

`func (o *OrganizationDetailsDto) GetCloudCredentials() int32`

GetCloudCredentials returns the CloudCredentials field if non-nil, zero value otherwise.

### GetCloudCredentialsOk

`func (o *OrganizationDetailsDto) GetCloudCredentialsOk() (*int32, bool)`

GetCloudCredentialsOk returns a tuple with the CloudCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentials

`func (o *OrganizationDetailsDto) SetCloudCredentials(v int32)`

SetCloudCredentials sets CloudCredentials field to given value.

### HasCloudCredentials

`func (o *OrganizationDetailsDto) HasCloudCredentials() bool`

HasCloudCredentials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationDetailsDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPartnerId

`func (o *OrganizationDetailsDto) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *OrganizationDetailsDto) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *OrganizationDetailsDto) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *OrganizationDetailsDto) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerName

`func (o *OrganizationDetailsDto) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *OrganizationDetailsDto) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *OrganizationDetailsDto) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *OrganizationDetailsDto) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetPartner

`func (o *OrganizationDetailsDto) GetPartner() PartnerDetailsForOrganizationsDto`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *OrganizationDetailsDto) GetPartnerOk() (*PartnerDetailsForOrganizationsDto, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *OrganizationDetailsDto) SetPartner(v PartnerDetailsForOrganizationsDto)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *OrganizationDetailsDto) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetDiscountRate

`func (o *OrganizationDetailsDto) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *OrganizationDetailsDto) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *OrganizationDetailsDto) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *OrganizationDetailsDto) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetBoundRules

`func (o *OrganizationDetailsDto) GetBoundRules() []PrometheusEntity`

GetBoundRules returns the BoundRules field if non-nil, zero value otherwise.

### GetBoundRulesOk

`func (o *OrganizationDetailsDto) GetBoundRulesOk() (*[]PrometheusEntity, bool)`

GetBoundRulesOk returns a tuple with the BoundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRules

`func (o *OrganizationDetailsDto) SetBoundRules(v []PrometheusEntity)`

SetBoundRules sets BoundRules field to given value.

### HasBoundRules

`func (o *OrganizationDetailsDto) HasBoundRules() bool`

HasBoundRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


