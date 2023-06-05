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

// checks if the BoundImagesForProjectsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundImagesForProjectsListDto{}

// BoundImagesForProjectsListDto struct for BoundImagesForProjectsListDto
type BoundImagesForProjectsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	IsAzure *bool `json:"isAzure,omitempty"`
	IsAws *bool `json:"isAws,omitempty"`
	IsOpenstack *bool `json:"isOpenstack,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
	Size *float64 `json:"size,omitempty"`
	ImageId *string `json:"imageId,omitempty"`
	CloudId *int32 `json:"cloudId,omitempty"`
	IsWindows *bool `json:"isWindows,omitempty"`
}

// NewBoundImagesForProjectsListDto instantiates a new BoundImagesForProjectsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundImagesForProjectsListDto() *BoundImagesForProjectsListDto {
	this := BoundImagesForProjectsListDto{}
	return &this
}

// NewBoundImagesForProjectsListDtoWithDefaults instantiates a new BoundImagesForProjectsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundImagesForProjectsListDtoWithDefaults() *BoundImagesForProjectsListDto {
	this := BoundImagesForProjectsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BoundImagesForProjectsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BoundImagesForProjectsListDto) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *BoundImagesForProjectsListDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetIsAzure returns the IsAzure field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetIsAzure() bool {
	if o == nil || IsNil(o.IsAzure) {
		var ret bool
		return ret
	}
	return *o.IsAzure
}

// GetIsAzureOk returns a tuple with the IsAzure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetIsAzureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAzure) {
		return nil, false
	}
	return o.IsAzure, true
}

// HasIsAzure returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasIsAzure() bool {
	if o != nil && !IsNil(o.IsAzure) {
		return true
	}

	return false
}

// SetIsAzure gets a reference to the given bool and assigns it to the IsAzure field.
func (o *BoundImagesForProjectsListDto) SetIsAzure(v bool) {
	o.IsAzure = &v
}

// GetIsAws returns the IsAws field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetIsAws() bool {
	if o == nil || IsNil(o.IsAws) {
		var ret bool
		return ret
	}
	return *o.IsAws
}

// GetIsAwsOk returns a tuple with the IsAws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetIsAwsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAws) {
		return nil, false
	}
	return o.IsAws, true
}

// HasIsAws returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasIsAws() bool {
	if o != nil && !IsNil(o.IsAws) {
		return true
	}

	return false
}

// SetIsAws gets a reference to the given bool and assigns it to the IsAws field.
func (o *BoundImagesForProjectsListDto) SetIsAws(v bool) {
	o.IsAws = &v
}

// GetIsOpenstack returns the IsOpenstack field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetIsOpenstack() bool {
	if o == nil || IsNil(o.IsOpenstack) {
		var ret bool
		return ret
	}
	return *o.IsOpenstack
}

// GetIsOpenstackOk returns a tuple with the IsOpenstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetIsOpenstackOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpenstack) {
		return nil, false
	}
	return o.IsOpenstack, true
}

// HasIsOpenstack returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasIsOpenstack() bool {
	if o != nil && !IsNil(o.IsOpenstack) {
		return true
	}

	return false
}

// SetIsOpenstack gets a reference to the given bool and assigns it to the IsOpenstack field.
func (o *BoundImagesForProjectsListDto) SetIsOpenstack(v bool) {
	o.IsOpenstack = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *BoundImagesForProjectsListDto) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetSize() float64 {
	if o == nil || IsNil(o.Size) {
		var ret float64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float64 and assigns it to the Size field.
func (o *BoundImagesForProjectsListDto) SetSize(v float64) {
	o.Size = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *BoundImagesForProjectsListDto) SetImageId(v string) {
	o.ImageId = &v
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *BoundImagesForProjectsListDto) SetCloudId(v int32) {
	o.CloudId = &v
}

// GetIsWindows returns the IsWindows field value if set, zero value otherwise.
func (o *BoundImagesForProjectsListDto) GetIsWindows() bool {
	if o == nil || IsNil(o.IsWindows) {
		var ret bool
		return ret
	}
	return *o.IsWindows
}

// GetIsWindowsOk returns a tuple with the IsWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsListDto) GetIsWindowsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWindows) {
		return nil, false
	}
	return o.IsWindows, true
}

// HasIsWindows returns a boolean if a field has been set.
func (o *BoundImagesForProjectsListDto) HasIsWindows() bool {
	if o != nil && !IsNil(o.IsWindows) {
		return true
	}

	return false
}

// SetIsWindows gets a reference to the given bool and assigns it to the IsWindows field.
func (o *BoundImagesForProjectsListDto) SetIsWindows(v bool) {
	o.IsWindows = &v
}

func (o BoundImagesForProjectsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoundImagesForProjectsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsAzure) {
		toSerialize["isAzure"] = o.IsAzure
	}
	if !IsNil(o.IsAws) {
		toSerialize["isAws"] = o.IsAws
	}
	if !IsNil(o.IsOpenstack) {
		toSerialize["isOpenstack"] = o.IsOpenstack
	}
	if !IsNil(o.ProjectName) {
		toSerialize["projectName"] = o.ProjectName
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	if !IsNil(o.IsWindows) {
		toSerialize["isWindows"] = o.IsWindows
	}
	return toSerialize, nil
}

type NullableBoundImagesForProjectsListDto struct {
	value *BoundImagesForProjectsListDto
	isSet bool
}

func (v NullableBoundImagesForProjectsListDto) Get() *BoundImagesForProjectsListDto {
	return v.value
}

func (v *NullableBoundImagesForProjectsListDto) Set(val *BoundImagesForProjectsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBoundImagesForProjectsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBoundImagesForProjectsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoundImagesForProjectsListDto(val *BoundImagesForProjectsListDto) *NullableBoundImagesForProjectsListDto {
	return &NullableBoundImagesForProjectsListDto{value: val, isSet: true}
}

func (v NullableBoundImagesForProjectsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoundImagesForProjectsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


