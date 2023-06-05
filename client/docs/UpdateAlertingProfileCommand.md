# UpdateAlertingProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**SlackConfigurationId** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Reminder** | Pointer to [**AlertingReminder**](AlertingReminder.md) |  | [optional] 

## Methods

### NewUpdateAlertingProfileCommand

`func NewUpdateAlertingProfileCommand(id int32, ) *UpdateAlertingProfileCommand`

NewUpdateAlertingProfileCommand instantiates a new UpdateAlertingProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertingProfileCommandWithDefaults

`func NewUpdateAlertingProfileCommandWithDefaults() *UpdateAlertingProfileCommand`

NewUpdateAlertingProfileCommandWithDefaults instantiates a new UpdateAlertingProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAlertingProfileCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAlertingProfileCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAlertingProfileCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAlertingProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAlertingProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAlertingProfileCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAlertingProfileCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *UpdateAlertingProfileCommand) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.

### HasSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) HasSlackConfigurationId() bool`

HasSlackConfigurationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UpdateAlertingProfileCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateAlertingProfileCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateAlertingProfileCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UpdateAlertingProfileCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetReminder

`func (o *UpdateAlertingProfileCommand) GetReminder() AlertingReminder`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *UpdateAlertingProfileCommand) GetReminderOk() (*AlertingReminder, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *UpdateAlertingProfileCommand) SetReminder(v AlertingReminder)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *UpdateAlertingProfileCommand) HasReminder() bool`

HasReminder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


