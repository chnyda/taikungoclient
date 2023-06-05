# Subresource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**HourlyCost** | Pointer to **string** |  | [optional] 
**MonthlyCost** | Pointer to **string** |  | [optional] 
**CostComponents** | Pointer to [**[]CostComponent**](CostComponent.md) |  | [optional] 

## Methods

### NewSubresource

`func NewSubresource() *Subresource`

NewSubresource instantiates a new Subresource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubresourceWithDefaults

`func NewSubresourceWithDefaults() *Subresource`

NewSubresourceWithDefaults instantiates a new Subresource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Subresource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subresource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subresource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subresource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *Subresource) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subresource) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subresource) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Subresource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetHourlyCost

`func (o *Subresource) GetHourlyCost() string`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Subresource) GetHourlyCostOk() (*string, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Subresource) SetHourlyCost(v string)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *Subresource) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetMonthlyCost

`func (o *Subresource) GetMonthlyCost() string`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *Subresource) GetMonthlyCostOk() (*string, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *Subresource) SetMonthlyCost(v string)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *Subresource) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### GetCostComponents

`func (o *Subresource) GetCostComponents() []CostComponent`

GetCostComponents returns the CostComponents field if non-nil, zero value otherwise.

### GetCostComponentsOk

`func (o *Subresource) GetCostComponentsOk() (*[]CostComponent, bool)`

GetCostComponentsOk returns a tuple with the CostComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostComponents

`func (o *Subresource) SetCostComponents(v []CostComponent)`

SetCostComponents sets CostComponents field to given value.

### HasCostComponents

`func (o *Subresource) HasCostComponents() bool`

HasCostComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


