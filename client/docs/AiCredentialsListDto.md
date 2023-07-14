# AiCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AiType**](AiType.md) |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewAiCredentialsListDto

`func NewAiCredentialsListDto() *AiCredentialsListDto`

NewAiCredentialsListDto instantiates a new AiCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiCredentialsListDtoWithDefaults

`func NewAiCredentialsListDtoWithDefaults() *AiCredentialsListDto`

NewAiCredentialsListDtoWithDefaults instantiates a new AiCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AiCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *AiCredentialsListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiCredentialsListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiCredentialsListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiCredentialsListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AiCredentialsListDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AiCredentialsListDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetName

`func (o *AiCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AiCredentialsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AiCredentialsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetApiKey

`func (o *AiCredentialsListDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AiCredentialsListDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AiCredentialsListDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AiCredentialsListDto) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *AiCredentialsListDto) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *AiCredentialsListDto) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetType

`func (o *AiCredentialsListDto) GetType() AiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiCredentialsListDto) GetTypeOk() (*AiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiCredentialsListDto) SetType(v AiType)`

SetType sets Type field to given value.

### HasType

`func (o *AiCredentialsListDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AiCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AiCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AiCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AiCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *AiCredentialsListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AiCredentialsListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *AiCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AiCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AiCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AiCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AiCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AiCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetProjects

`func (o *AiCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AiCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AiCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AiCredentialsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *AiCredentialsListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *AiCredentialsListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetIsDefault

`func (o *AiCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AiCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AiCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AiCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


