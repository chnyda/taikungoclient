# UpdatePaymentIdCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**PaymentIntentId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePaymentIdCommand

`func NewUpdatePaymentIdCommand() *UpdatePaymentIdCommand`

NewUpdatePaymentIdCommand instantiates a new UpdatePaymentIdCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentIdCommandWithDefaults

`func NewUpdatePaymentIdCommandWithDefaults() *UpdatePaymentIdCommand`

NewUpdatePaymentIdCommandWithDefaults instantiates a new UpdatePaymentIdCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *UpdatePaymentIdCommand) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UpdatePaymentIdCommand) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UpdatePaymentIdCommand) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UpdatePaymentIdCommand) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *UpdatePaymentIdCommand) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *UpdatePaymentIdCommand) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *UpdatePaymentIdCommand) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *UpdatePaymentIdCommand) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


