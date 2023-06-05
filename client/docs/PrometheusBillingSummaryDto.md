# PrometheusBillingSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PrometheusRuleId** | Pointer to **int32** |  | [optional] 
**RuleName** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewPrometheusBillingSummaryDto

`func NewPrometheusBillingSummaryDto() *PrometheusBillingSummaryDto`

NewPrometheusBillingSummaryDto instantiates a new PrometheusBillingSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusBillingSummaryDtoWithDefaults

`func NewPrometheusBillingSummaryDtoWithDefaults() *PrometheusBillingSummaryDto`

NewPrometheusBillingSummaryDtoWithDefaults instantiates a new PrometheusBillingSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PrometheusBillingSummaryDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PrometheusBillingSummaryDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PrometheusBillingSummaryDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PrometheusBillingSummaryDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStartDate

`func (o *PrometheusBillingSummaryDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PrometheusBillingSummaryDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PrometheusBillingSummaryDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PrometheusBillingSummaryDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PrometheusBillingSummaryDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PrometheusBillingSummaryDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PrometheusBillingSummaryDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PrometheusBillingSummaryDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPrometheusRuleId

`func (o *PrometheusBillingSummaryDto) GetPrometheusRuleId() int32`

GetPrometheusRuleId returns the PrometheusRuleId field if non-nil, zero value otherwise.

### GetPrometheusRuleIdOk

`func (o *PrometheusBillingSummaryDto) GetPrometheusRuleIdOk() (*int32, bool)`

GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusRuleId

`func (o *PrometheusBillingSummaryDto) SetPrometheusRuleId(v int32)`

SetPrometheusRuleId sets PrometheusRuleId field to given value.

### HasPrometheusRuleId

`func (o *PrometheusBillingSummaryDto) HasPrometheusRuleId() bool`

HasPrometheusRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *PrometheusBillingSummaryDto) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *PrometheusBillingSummaryDto) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *PrometheusBillingSummaryDto) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *PrometheusBillingSummaryDto) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrometheusBillingSummaryDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrometheusBillingSummaryDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrometheusBillingSummaryDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrometheusBillingSummaryDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *PrometheusBillingSummaryDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PrometheusBillingSummaryDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PrometheusBillingSummaryDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *PrometheusBillingSummaryDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PrometheusBillingSummaryDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PrometheusBillingSummaryDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PrometheusBillingSummaryDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PrometheusBillingSummaryDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


