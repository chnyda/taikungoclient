/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the PrometheusRuleListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusRuleListDto{}

// PrometheusRuleListDto struct for PrometheusRuleListDto
type PrometheusRuleListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Url *string `json:"url,omitempty"`
	MetricName *string `json:"metricName,omitempty"`
	Labels []PrometheusLabelUpdateDto `json:"labels,omitempty"`
	BoundOrganizations []PrometheusOrganizationDiscountDto `json:"boundOrganizations,omitempty"`
	Type *string `json:"type,omitempty"`
	Price *float64 `json:"price,omitempty"`
	BillingStartDate *string `json:"billingStartDate,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Partner *PartnerDetailsDto `json:"partner,omitempty"`
	OperationCredential *OperationCredentialsForOrganizationEntity `json:"operationCredential,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
}

// NewPrometheusRuleListDto instantiates a new PrometheusRuleListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusRuleListDto() *PrometheusRuleListDto {
	this := PrometheusRuleListDto{}
	return &this
}

// NewPrometheusRuleListDtoWithDefaults instantiates a new PrometheusRuleListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusRuleListDtoWithDefaults() *PrometheusRuleListDto {
	this := PrometheusRuleListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PrometheusRuleListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrometheusRuleListDto) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PrometheusRuleListDto) SetPassword(v string) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PrometheusRuleListDto) SetUserName(v string) {
	o.UserName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PrometheusRuleListDto) SetUrl(v string) {
	o.Url = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *PrometheusRuleListDto) SetMetricName(v string) {
	o.MetricName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetLabels() []PrometheusLabelUpdateDto {
	if o == nil || IsNil(o.Labels) {
		var ret []PrometheusLabelUpdateDto
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetLabelsOk() ([]PrometheusLabelUpdateDto, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []PrometheusLabelUpdateDto and assigns it to the Labels field.
func (o *PrometheusRuleListDto) SetLabels(v []PrometheusLabelUpdateDto) {
	o.Labels = v
}

// GetBoundOrganizations returns the BoundOrganizations field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetBoundOrganizations() []PrometheusOrganizationDiscountDto {
	if o == nil || IsNil(o.BoundOrganizations) {
		var ret []PrometheusOrganizationDiscountDto
		return ret
	}
	return o.BoundOrganizations
}

// GetBoundOrganizationsOk returns a tuple with the BoundOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetBoundOrganizationsOk() ([]PrometheusOrganizationDiscountDto, bool) {
	if o == nil || IsNil(o.BoundOrganizations) {
		return nil, false
	}
	return o.BoundOrganizations, true
}

// HasBoundOrganizations returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasBoundOrganizations() bool {
	if o != nil && !IsNil(o.BoundOrganizations) {
		return true
	}

	return false
}

// SetBoundOrganizations gets a reference to the given []PrometheusOrganizationDiscountDto and assigns it to the BoundOrganizations field.
func (o *PrometheusRuleListDto) SetBoundOrganizations(v []PrometheusOrganizationDiscountDto) {
	o.BoundOrganizations = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PrometheusRuleListDto) SetType(v string) {
	o.Type = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PrometheusRuleListDto) SetPrice(v float64) {
	o.Price = &v
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetBillingStartDate() string {
	if o == nil || IsNil(o.BillingStartDate) {
		var ret string
		return ret
	}
	return *o.BillingStartDate
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetBillingStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.BillingStartDate) {
		return nil, false
	}
	return o.BillingStartDate, true
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasBillingStartDate() bool {
	if o != nil && !IsNil(o.BillingStartDate) {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given string and assigns it to the BillingStartDate field.
func (o *PrometheusRuleListDto) SetBillingStartDate(v string) {
	o.BillingStartDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PrometheusRuleListDto) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetPartner() PartnerDetailsDto {
	if o == nil || IsNil(o.Partner) {
		var ret PartnerDetailsDto
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetPartnerOk() (*PartnerDetailsDto, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given PartnerDetailsDto and assigns it to the Partner field.
func (o *PrometheusRuleListDto) SetPartner(v PartnerDetailsDto) {
	o.Partner = &v
}

// GetOperationCredential returns the OperationCredential field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetOperationCredential() OperationCredentialsForOrganizationEntity {
	if o == nil || IsNil(o.OperationCredential) {
		var ret OperationCredentialsForOrganizationEntity
		return ret
	}
	return *o.OperationCredential
}

// GetOperationCredentialOk returns a tuple with the OperationCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetOperationCredentialOk() (*OperationCredentialsForOrganizationEntity, bool) {
	if o == nil || IsNil(o.OperationCredential) {
		return nil, false
	}
	return o.OperationCredential, true
}

// HasOperationCredential returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasOperationCredential() bool {
	if o != nil && !IsNil(o.OperationCredential) {
		return true
	}

	return false
}

// SetOperationCredential gets a reference to the given OperationCredentialsForOrganizationEntity and assigns it to the OperationCredential field.
func (o *PrometheusRuleListDto) SetOperationCredential(v OperationCredentialsForOrganizationEntity) {
	o.OperationCredential = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *PrometheusRuleListDto) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *PrometheusRuleListDto) SetLastModified(v string) {
	o.LastModified = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *PrometheusRuleListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

func (o PrometheusRuleListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusRuleListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.BoundOrganizations) {
		toSerialize["boundOrganizations"] = o.BoundOrganizations
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.BillingStartDate) {
		toSerialize["billingStartDate"] = o.BillingStartDate
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.OperationCredential) {
		toSerialize["operationCredential"] = o.OperationCredential
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	return toSerialize, nil
}

type NullablePrometheusRuleListDto struct {
	value *PrometheusRuleListDto
	isSet bool
}

func (v NullablePrometheusRuleListDto) Get() *PrometheusRuleListDto {
	return v.value
}

func (v *NullablePrometheusRuleListDto) Set(val *PrometheusRuleListDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusRuleListDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusRuleListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusRuleListDto(val *PrometheusRuleListDto) *NullablePrometheusRuleListDto {
	return &NullablePrometheusRuleListDto{value: val, isSet: true}
}

func (v NullablePrometheusRuleListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusRuleListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


