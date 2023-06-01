# StandAloneProfileSecurityGroupForDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**PortMinRange** | Pointer to **int32** |  | [optional] 
**PortMaxRange** | Pointer to **int32** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**IsRdpPortEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewStandAloneProfileSecurityGroupForDetailsDto

`func NewStandAloneProfileSecurityGroupForDetailsDto() *StandAloneProfileSecurityGroupForDetailsDto`

NewStandAloneProfileSecurityGroupForDetailsDto instantiates a new StandAloneProfileSecurityGroupForDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfileSecurityGroupForDetailsDtoWithDefaults

`func NewStandAloneProfileSecurityGroupForDetailsDtoWithDefaults() *StandAloneProfileSecurityGroupForDetailsDto`

NewStandAloneProfileSecurityGroupForDetailsDtoWithDefaults instantiates a new StandAloneProfileSecurityGroupForDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPortMinRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.

### HasPortMinRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasPortMinRange() bool`

HasPortMinRange returns a boolean if a field has been set.

### GetPortMaxRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.

### HasPortMaxRange

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasPortMaxRange() bool`

HasPortMaxRange returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetIsRdpPortEnabled

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIsRdpPortEnabled() bool`

GetIsRdpPortEnabled returns the IsRdpPortEnabled field if non-nil, zero value otherwise.

### GetIsRdpPortEnabledOk

`func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIsRdpPortEnabledOk() (*bool, bool)`

GetIsRdpPortEnabledOk returns a tuple with the IsRdpPortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRdpPortEnabled

`func (o *StandAloneProfileSecurityGroupForDetailsDto) SetIsRdpPortEnabled(v bool)`

SetIsRdpPortEnabled sets IsRdpPortEnabled field to given value.

### HasIsRdpPortEnabled

`func (o *StandAloneProfileSecurityGroupForDetailsDto) HasIsRdpPortEnabled() bool`

HasIsRdpPortEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


