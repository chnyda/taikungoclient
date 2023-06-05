# PrivateSubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ListForPartnersDto**](ListForPartnersDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**IsEligibleToSwitch** | Pointer to **bool** |  | [optional] 
**ActiveSubscriptionStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivateSubscriptionList

`func NewPrivateSubscriptionList() *PrivateSubscriptionList`

NewPrivateSubscriptionList instantiates a new PrivateSubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSubscriptionListWithDefaults

`func NewPrivateSubscriptionListWithDefaults() *PrivateSubscriptionList`

NewPrivateSubscriptionListWithDefaults instantiates a new PrivateSubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PrivateSubscriptionList) GetData() []ListForPartnersDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrivateSubscriptionList) GetDataOk() (*[]ListForPartnersDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrivateSubscriptionList) SetData(v []ListForPartnersDto)`

SetData sets Data field to given value.

### HasData

`func (o *PrivateSubscriptionList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalCount

`func (o *PrivateSubscriptionList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PrivateSubscriptionList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PrivateSubscriptionList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PrivateSubscriptionList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetIsEligibleToSwitch

`func (o *PrivateSubscriptionList) GetIsEligibleToSwitch() bool`

GetIsEligibleToSwitch returns the IsEligibleToSwitch field if non-nil, zero value otherwise.

### GetIsEligibleToSwitchOk

`func (o *PrivateSubscriptionList) GetIsEligibleToSwitchOk() (*bool, bool)`

GetIsEligibleToSwitchOk returns a tuple with the IsEligibleToSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleToSwitch

`func (o *PrivateSubscriptionList) SetIsEligibleToSwitch(v bool)`

SetIsEligibleToSwitch sets IsEligibleToSwitch field to given value.

### HasIsEligibleToSwitch

`func (o *PrivateSubscriptionList) HasIsEligibleToSwitch() bool`

HasIsEligibleToSwitch returns a boolean if a field has been set.

### GetActiveSubscriptionStatus

`func (o *PrivateSubscriptionList) GetActiveSubscriptionStatus() string`

GetActiveSubscriptionStatus returns the ActiveSubscriptionStatus field if non-nil, zero value otherwise.

### GetActiveSubscriptionStatusOk

`func (o *PrivateSubscriptionList) GetActiveSubscriptionStatusOk() (*string, bool)`

GetActiveSubscriptionStatusOk returns a tuple with the ActiveSubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptionStatus

`func (o *PrivateSubscriptionList) SetActiveSubscriptionStatus(v string)`

SetActiveSubscriptionStatus sets ActiveSubscriptionStatus field to given value.

### HasActiveSubscriptionStatus

`func (o *PrivateSubscriptionList) HasActiveSubscriptionStatus() bool`

HasActiveSubscriptionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


