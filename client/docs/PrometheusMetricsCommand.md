# PrometheusMetricsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**IsGraphEnabled** | Pointer to **bool** |  | [optional] 
**IsAutoComplete** | Pointer to **bool** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 

## Methods

### NewPrometheusMetricsCommand

`func NewPrometheusMetricsCommand() *PrometheusMetricsCommand`

NewPrometheusMetricsCommand instantiates a new PrometheusMetricsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMetricsCommandWithDefaults

`func NewPrometheusMetricsCommandWithDefaults() *PrometheusMetricsCommand`

NewPrometheusMetricsCommandWithDefaults instantiates a new PrometheusMetricsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PrometheusMetricsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PrometheusMetricsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PrometheusMetricsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PrometheusMetricsCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParameters

`func (o *PrometheusMetricsCommand) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PrometheusMetricsCommand) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PrometheusMetricsCommand) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PrometheusMetricsCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTime

`func (o *PrometheusMetricsCommand) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PrometheusMetricsCommand) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PrometheusMetricsCommand) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *PrometheusMetricsCommand) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStart

`func (o *PrometheusMetricsCommand) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PrometheusMetricsCommand) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PrometheusMetricsCommand) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *PrometheusMetricsCommand) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *PrometheusMetricsCommand) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PrometheusMetricsCommand) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PrometheusMetricsCommand) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PrometheusMetricsCommand) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIsGraphEnabled

`func (o *PrometheusMetricsCommand) GetIsGraphEnabled() bool`

GetIsGraphEnabled returns the IsGraphEnabled field if non-nil, zero value otherwise.

### GetIsGraphEnabledOk

`func (o *PrometheusMetricsCommand) GetIsGraphEnabledOk() (*bool, bool)`

GetIsGraphEnabledOk returns a tuple with the IsGraphEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGraphEnabled

`func (o *PrometheusMetricsCommand) SetIsGraphEnabled(v bool)`

SetIsGraphEnabled sets IsGraphEnabled field to given value.

### HasIsGraphEnabled

`func (o *PrometheusMetricsCommand) HasIsGraphEnabled() bool`

HasIsGraphEnabled returns a boolean if a field has been set.

### GetIsAutoComplete

`func (o *PrometheusMetricsCommand) GetIsAutoComplete() bool`

GetIsAutoComplete returns the IsAutoComplete field if non-nil, zero value otherwise.

### GetIsAutoCompleteOk

`func (o *PrometheusMetricsCommand) GetIsAutoCompleteOk() (*bool, bool)`

GetIsAutoCompleteOk returns a tuple with the IsAutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoComplete

`func (o *PrometheusMetricsCommand) SetIsAutoComplete(v bool)`

SetIsAutoComplete sets IsAutoComplete field to given value.

### HasIsAutoComplete

`func (o *PrometheusMetricsCommand) HasIsAutoComplete() bool`

HasIsAutoComplete returns a boolean if a field has been set.

### GetStep

`func (o *PrometheusMetricsCommand) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PrometheusMetricsCommand) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PrometheusMetricsCommand) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *PrometheusMetricsCommand) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


