# PrometheusLabelListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewPrometheusLabelListDto

`func NewPrometheusLabelListDto(value string, ) *PrometheusLabelListDto`

NewPrometheusLabelListDto instantiates a new PrometheusLabelListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusLabelListDtoWithDefaults

`func NewPrometheusLabelListDtoWithDefaults() *PrometheusLabelListDto`

NewPrometheusLabelListDtoWithDefaults instantiates a new PrometheusLabelListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PrometheusLabelListDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PrometheusLabelListDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PrometheusLabelListDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PrometheusLabelListDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *PrometheusLabelListDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrometheusLabelListDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrometheusLabelListDto) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


