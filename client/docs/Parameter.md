# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParameter

`func NewParameter() *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Parameter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Parameter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Parameter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Parameter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Parameter) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Parameter) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *Parameter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Parameter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Parameter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Parameter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *Parameter) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *Parameter) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *Parameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Parameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Parameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Parameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Parameter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Parameter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


