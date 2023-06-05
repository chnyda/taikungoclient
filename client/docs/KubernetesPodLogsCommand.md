# KubernetesPodLogsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesPodLogsCommand

`func NewKubernetesPodLogsCommand() *KubernetesPodLogsCommand`

NewKubernetesPodLogsCommand instantiates a new KubernetesPodLogsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodLogsCommandWithDefaults

`func NewKubernetesPodLogsCommandWithDefaults() *KubernetesPodLogsCommand`

NewKubernetesPodLogsCommandWithDefaults instantiates a new KubernetesPodLogsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *KubernetesPodLogsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubernetesPodLogsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubernetesPodLogsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KubernetesPodLogsCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesPodLogsCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesPodLogsCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesPodLogsCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesPodLogsCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KubernetesPodLogsCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesPodLogsCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesPodLogsCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesPodLogsCommand) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetContainer

`func (o *KubernetesPodLogsCommand) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *KubernetesPodLogsCommand) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *KubernetesPodLogsCommand) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *KubernetesPodLogsCommand) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


