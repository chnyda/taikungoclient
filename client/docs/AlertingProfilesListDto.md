# AlertingProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**SlackConfigurationId** | Pointer to **int32** |  | [optional] 
**SlackConfigurationName** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Emails** | Pointer to [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | [optional] 
**Webhooks** | Pointer to [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**Reminder** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertingProfilesListDto

`func NewAlertingProfilesListDto() *AlertingProfilesListDto`

NewAlertingProfilesListDto instantiates a new AlertingProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfilesListDtoWithDefaults

`func NewAlertingProfilesListDtoWithDefaults() *AlertingProfilesListDto`

NewAlertingProfilesListDtoWithDefaults instantiates a new AlertingProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertingProfilesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertingProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertingProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertingProfilesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertingProfilesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AlertingProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertingProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertingProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AlertingProfilesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AlertingProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AlertingProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AlertingProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AlertingProfilesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetSlackConfigurationId

`func (o *AlertingProfilesListDto) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *AlertingProfilesListDto) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.

### HasSlackConfigurationId

`func (o *AlertingProfilesListDto) HasSlackConfigurationId() bool`

HasSlackConfigurationId returns a boolean if a field has been set.

### GetSlackConfigurationName

`func (o *AlertingProfilesListDto) GetSlackConfigurationName() string`

GetSlackConfigurationName returns the SlackConfigurationName field if non-nil, zero value otherwise.

### GetSlackConfigurationNameOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationNameOk() (*string, bool)`

GetSlackConfigurationNameOk returns a tuple with the SlackConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationName

`func (o *AlertingProfilesListDto) SetSlackConfigurationName(v string)`

SetSlackConfigurationName sets SlackConfigurationName field to given value.

### HasSlackConfigurationName

`func (o *AlertingProfilesListDto) HasSlackConfigurationName() bool`

HasSlackConfigurationName returns a boolean if a field has been set.

### GetIsLocked

`func (o *AlertingProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AlertingProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AlertingProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AlertingProfilesListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetEmails

`func (o *AlertingProfilesListDto) GetEmails() []AlertingEmailDto`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AlertingProfilesListDto) GetEmailsOk() (*[]AlertingEmailDto, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AlertingProfilesListDto) SetEmails(v []AlertingEmailDto)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AlertingProfilesListDto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetWebhooks

`func (o *AlertingProfilesListDto) GetWebhooks() []AlertingWebhookDto`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AlertingProfilesListDto) GetWebhooksOk() (*[]AlertingWebhookDto, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AlertingProfilesListDto) SetWebhooks(v []AlertingWebhookDto)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *AlertingProfilesListDto) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetProjects

`func (o *AlertingProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AlertingProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AlertingProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AlertingProfilesListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AlertingProfilesListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertingProfilesListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertingProfilesListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AlertingProfilesListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *AlertingProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AlertingProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AlertingProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AlertingProfilesListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *AlertingProfilesListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AlertingProfilesListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AlertingProfilesListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *AlertingProfilesListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetReminder

`func (o *AlertingProfilesListDto) GetReminder() string`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *AlertingProfilesListDto) GetReminderOk() (*string, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *AlertingProfilesListDto) SetReminder(v string)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *AlertingProfilesListDto) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertingProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertingProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertingProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertingProfilesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


