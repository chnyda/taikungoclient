# SshUsersListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SshPublicKey** | Pointer to **string** |  | [optional] 
**AccessProfileName** | Pointer to **string** |  | [optional] 

## Methods

### NewSshUsersListDto

`func NewSshUsersListDto() *SshUsersListDto`

NewSshUsersListDto instantiates a new SshUsersListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUsersListDtoWithDefaults

`func NewSshUsersListDtoWithDefaults() *SshUsersListDto`

NewSshUsersListDtoWithDefaults instantiates a new SshUsersListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshUsersListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshUsersListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshUsersListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SshUsersListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SshUsersListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshUsersListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshUsersListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshUsersListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *SshUsersListDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *SshUsersListDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *SshUsersListDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *SshUsersListDto) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetAccessProfileName

`func (o *SshUsersListDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *SshUsersListDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *SshUsersListDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.

### HasAccessProfileName

`func (o *SshUsersListDto) HasAccessProfileName() bool`

HasAccessProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


