# PrometheusRuleListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]PrometheusLabelUpdateDto**](PrometheusLabelUpdateDto.md) |  | [optional] 
**BoundOrganizations** | Pointer to [**[]PrometheusOrganizationDiscountDto**](PrometheusOrganizationDiscountDto.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**BillingStartDate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**PartnerDetailsDto**](PartnerDetailsDto.md) |  | [optional] 
**OperationCredential** | Pointer to [**OperationCredentialsForOrganizationEntity**](OperationCredentialsForOrganizationEntity.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewPrometheusRuleListDto

`func NewPrometheusRuleListDto() *PrometheusRuleListDto`

NewPrometheusRuleListDto instantiates a new PrometheusRuleListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusRuleListDtoWithDefaults

`func NewPrometheusRuleListDtoWithDefaults() *PrometheusRuleListDto`

NewPrometheusRuleListDtoWithDefaults instantiates a new PrometheusRuleListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusRuleListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusRuleListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusRuleListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PrometheusRuleListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrometheusRuleListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrometheusRuleListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrometheusRuleListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrometheusRuleListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *PrometheusRuleListDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PrometheusRuleListDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PrometheusRuleListDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PrometheusRuleListDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *PrometheusRuleListDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PrometheusRuleListDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PrometheusRuleListDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PrometheusRuleListDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUrl

`func (o *PrometheusRuleListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrometheusRuleListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrometheusRuleListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PrometheusRuleListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMetricName

`func (o *PrometheusRuleListDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *PrometheusRuleListDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *PrometheusRuleListDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *PrometheusRuleListDto) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetLabels

`func (o *PrometheusRuleListDto) GetLabels() []PrometheusLabelUpdateDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PrometheusRuleListDto) GetLabelsOk() (*[]PrometheusLabelUpdateDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PrometheusRuleListDto) SetLabels(v []PrometheusLabelUpdateDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PrometheusRuleListDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetBoundOrganizations

`func (o *PrometheusRuleListDto) GetBoundOrganizations() []PrometheusOrganizationDiscountDto`

GetBoundOrganizations returns the BoundOrganizations field if non-nil, zero value otherwise.

### GetBoundOrganizationsOk

`func (o *PrometheusRuleListDto) GetBoundOrganizationsOk() (*[]PrometheusOrganizationDiscountDto, bool)`

GetBoundOrganizationsOk returns a tuple with the BoundOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizations

`func (o *PrometheusRuleListDto) SetBoundOrganizations(v []PrometheusOrganizationDiscountDto)`

SetBoundOrganizations sets BoundOrganizations field to given value.

### HasBoundOrganizations

`func (o *PrometheusRuleListDto) HasBoundOrganizations() bool`

HasBoundOrganizations returns a boolean if a field has been set.

### GetType

`func (o *PrometheusRuleListDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrometheusRuleListDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrometheusRuleListDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrometheusRuleListDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *PrometheusRuleListDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PrometheusRuleListDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PrometheusRuleListDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PrometheusRuleListDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *PrometheusRuleListDto) GetBillingStartDate() string`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *PrometheusRuleListDto) GetBillingStartDateOk() (*string, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *PrometheusRuleListDto) SetBillingStartDate(v string)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *PrometheusRuleListDto) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrometheusRuleListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrometheusRuleListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrometheusRuleListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrometheusRuleListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPartner

`func (o *PrometheusRuleListDto) GetPartner() PartnerDetailsDto`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *PrometheusRuleListDto) GetPartnerOk() (*PartnerDetailsDto, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *PrometheusRuleListDto) SetPartner(v PartnerDetailsDto)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *PrometheusRuleListDto) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetOperationCredential

`func (o *PrometheusRuleListDto) GetOperationCredential() OperationCredentialsForOrganizationEntity`

GetOperationCredential returns the OperationCredential field if non-nil, zero value otherwise.

### GetOperationCredentialOk

`func (o *PrometheusRuleListDto) GetOperationCredentialOk() (*OperationCredentialsForOrganizationEntity, bool)`

GetOperationCredentialOk returns a tuple with the OperationCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCredential

`func (o *PrometheusRuleListDto) SetOperationCredential(v OperationCredentialsForOrganizationEntity)`

SetOperationCredential sets OperationCredential field to given value.

### HasOperationCredential

`func (o *PrometheusRuleListDto) HasOperationCredential() bool`

HasOperationCredential returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrometheusRuleListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrometheusRuleListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrometheusRuleListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrometheusRuleListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *PrometheusRuleListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PrometheusRuleListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PrometheusRuleListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *PrometheusRuleListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PrometheusRuleListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PrometheusRuleListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PrometheusRuleListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PrometheusRuleListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


