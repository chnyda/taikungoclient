# PrometheusDashboardListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryName** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]PrometheusDashboardDto**](PrometheusDashboardDto.md) |  | [optional] 

## Methods

### NewPrometheusDashboardListDto

`func NewPrometheusDashboardListDto() *PrometheusDashboardListDto`

NewPrometheusDashboardListDto instantiates a new PrometheusDashboardListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusDashboardListDtoWithDefaults

`func NewPrometheusDashboardListDtoWithDefaults() *PrometheusDashboardListDto`

NewPrometheusDashboardListDtoWithDefaults instantiates a new PrometheusDashboardListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryName

`func (o *PrometheusDashboardListDto) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *PrometheusDashboardListDto) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *PrometheusDashboardListDto) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *PrometheusDashboardListDto) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetData

`func (o *PrometheusDashboardListDto) GetData() []PrometheusDashboardDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrometheusDashboardListDto) GetDataOk() (*[]PrometheusDashboardDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrometheusDashboardListDto) SetData(v []PrometheusDashboardDto)`

SetData sets Data field to given value.

### HasData

`func (o *PrometheusDashboardListDto) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


