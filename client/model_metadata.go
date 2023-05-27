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
	"time"
)

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata struct for Metadata
type Metadata struct {
	InfracostCommand NullableString `json:"infracostCommand,omitempty"`
	VcsBranch NullableString `json:"vcsBranch,omitempty"`
	VcsCommitSha NullableString `json:"vcsCommitSha,omitempty"`
	VcsCommitAuthorName NullableString `json:"vcsCommitAuthorName,omitempty"`
	VcsCommitAuthorEmail NullableString `json:"vcsCommitAuthorEmail,omitempty"`
	VcsCommitTimestamp *time.Time `json:"vcsCommitTimestamp,omitempty"`
	VcsCommitMessage NullableString `json:"vcsCommitMessage,omitempty"`
	VcsRepositoryUrl NullableString `json:"vcsRepositoryUrl,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Type NullableString `json:"type,omitempty"`
	VcsSubPath NullableString `json:"vcsSubPath,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetInfracostCommand returns the InfracostCommand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetInfracostCommand() string {
	if o == nil || IsNil(o.InfracostCommand.Get()) {
		var ret string
		return ret
	}
	return *o.InfracostCommand.Get()
}

// GetInfracostCommandOk returns a tuple with the InfracostCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetInfracostCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfracostCommand.Get(), o.InfracostCommand.IsSet()
}

// HasInfracostCommand returns a boolean if a field has been set.
func (o *Metadata) HasInfracostCommand() bool {
	if o != nil && o.InfracostCommand.IsSet() {
		return true
	}

	return false
}

// SetInfracostCommand gets a reference to the given NullableString and assigns it to the InfracostCommand field.
func (o *Metadata) SetInfracostCommand(v string) {
	o.InfracostCommand.Set(&v)
}
// SetInfracostCommandNil sets the value for InfracostCommand to be an explicit nil
func (o *Metadata) SetInfracostCommandNil() {
	o.InfracostCommand.Set(nil)
}

// UnsetInfracostCommand ensures that no value is present for InfracostCommand, not even an explicit nil
func (o *Metadata) UnsetInfracostCommand() {
	o.InfracostCommand.Unset()
}

// GetVcsBranch returns the VcsBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsBranch() string {
	if o == nil || IsNil(o.VcsBranch.Get()) {
		var ret string
		return ret
	}
	return *o.VcsBranch.Get()
}

// GetVcsBranchOk returns a tuple with the VcsBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsBranch.Get(), o.VcsBranch.IsSet()
}

// HasVcsBranch returns a boolean if a field has been set.
func (o *Metadata) HasVcsBranch() bool {
	if o != nil && o.VcsBranch.IsSet() {
		return true
	}

	return false
}

// SetVcsBranch gets a reference to the given NullableString and assigns it to the VcsBranch field.
func (o *Metadata) SetVcsBranch(v string) {
	o.VcsBranch.Set(&v)
}
// SetVcsBranchNil sets the value for VcsBranch to be an explicit nil
func (o *Metadata) SetVcsBranchNil() {
	o.VcsBranch.Set(nil)
}

// UnsetVcsBranch ensures that no value is present for VcsBranch, not even an explicit nil
func (o *Metadata) UnsetVcsBranch() {
	o.VcsBranch.Unset()
}

// GetVcsCommitSha returns the VcsCommitSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsCommitSha() string {
	if o == nil || IsNil(o.VcsCommitSha.Get()) {
		var ret string
		return ret
	}
	return *o.VcsCommitSha.Get()
}

// GetVcsCommitShaOk returns a tuple with the VcsCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsCommitSha.Get(), o.VcsCommitSha.IsSet()
}

// HasVcsCommitSha returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitSha() bool {
	if o != nil && o.VcsCommitSha.IsSet() {
		return true
	}

	return false
}

// SetVcsCommitSha gets a reference to the given NullableString and assigns it to the VcsCommitSha field.
func (o *Metadata) SetVcsCommitSha(v string) {
	o.VcsCommitSha.Set(&v)
}
// SetVcsCommitShaNil sets the value for VcsCommitSha to be an explicit nil
func (o *Metadata) SetVcsCommitShaNil() {
	o.VcsCommitSha.Set(nil)
}

// UnsetVcsCommitSha ensures that no value is present for VcsCommitSha, not even an explicit nil
func (o *Metadata) UnsetVcsCommitSha() {
	o.VcsCommitSha.Unset()
}

// GetVcsCommitAuthorName returns the VcsCommitAuthorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsCommitAuthorName() string {
	if o == nil || IsNil(o.VcsCommitAuthorName.Get()) {
		var ret string
		return ret
	}
	return *o.VcsCommitAuthorName.Get()
}

// GetVcsCommitAuthorNameOk returns a tuple with the VcsCommitAuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsCommitAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsCommitAuthorName.Get(), o.VcsCommitAuthorName.IsSet()
}

// HasVcsCommitAuthorName returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitAuthorName() bool {
	if o != nil && o.VcsCommitAuthorName.IsSet() {
		return true
	}

	return false
}

// SetVcsCommitAuthorName gets a reference to the given NullableString and assigns it to the VcsCommitAuthorName field.
func (o *Metadata) SetVcsCommitAuthorName(v string) {
	o.VcsCommitAuthorName.Set(&v)
}
// SetVcsCommitAuthorNameNil sets the value for VcsCommitAuthorName to be an explicit nil
func (o *Metadata) SetVcsCommitAuthorNameNil() {
	o.VcsCommitAuthorName.Set(nil)
}

// UnsetVcsCommitAuthorName ensures that no value is present for VcsCommitAuthorName, not even an explicit nil
func (o *Metadata) UnsetVcsCommitAuthorName() {
	o.VcsCommitAuthorName.Unset()
}

// GetVcsCommitAuthorEmail returns the VcsCommitAuthorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsCommitAuthorEmail() string {
	if o == nil || IsNil(o.VcsCommitAuthorEmail.Get()) {
		var ret string
		return ret
	}
	return *o.VcsCommitAuthorEmail.Get()
}

// GetVcsCommitAuthorEmailOk returns a tuple with the VcsCommitAuthorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsCommitAuthorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsCommitAuthorEmail.Get(), o.VcsCommitAuthorEmail.IsSet()
}

// HasVcsCommitAuthorEmail returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitAuthorEmail() bool {
	if o != nil && o.VcsCommitAuthorEmail.IsSet() {
		return true
	}

	return false
}

// SetVcsCommitAuthorEmail gets a reference to the given NullableString and assigns it to the VcsCommitAuthorEmail field.
func (o *Metadata) SetVcsCommitAuthorEmail(v string) {
	o.VcsCommitAuthorEmail.Set(&v)
}
// SetVcsCommitAuthorEmailNil sets the value for VcsCommitAuthorEmail to be an explicit nil
func (o *Metadata) SetVcsCommitAuthorEmailNil() {
	o.VcsCommitAuthorEmail.Set(nil)
}

// UnsetVcsCommitAuthorEmail ensures that no value is present for VcsCommitAuthorEmail, not even an explicit nil
func (o *Metadata) UnsetVcsCommitAuthorEmail() {
	o.VcsCommitAuthorEmail.Unset()
}

// GetVcsCommitTimestamp returns the VcsCommitTimestamp field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitTimestamp() time.Time {
	if o == nil || IsNil(o.VcsCommitTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.VcsCommitTimestamp
}

// GetVcsCommitTimestampOk returns a tuple with the VcsCommitTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VcsCommitTimestamp) {
		return nil, false
	}
	return o.VcsCommitTimestamp, true
}

// HasVcsCommitTimestamp returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitTimestamp() bool {
	if o != nil && !IsNil(o.VcsCommitTimestamp) {
		return true
	}

	return false
}

// SetVcsCommitTimestamp gets a reference to the given time.Time and assigns it to the VcsCommitTimestamp field.
func (o *Metadata) SetVcsCommitTimestamp(v time.Time) {
	o.VcsCommitTimestamp = &v
}

// GetVcsCommitMessage returns the VcsCommitMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsCommitMessage() string {
	if o == nil || IsNil(o.VcsCommitMessage.Get()) {
		var ret string
		return ret
	}
	return *o.VcsCommitMessage.Get()
}

// GetVcsCommitMessageOk returns a tuple with the VcsCommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsCommitMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsCommitMessage.Get(), o.VcsCommitMessage.IsSet()
}

// HasVcsCommitMessage returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitMessage() bool {
	if o != nil && o.VcsCommitMessage.IsSet() {
		return true
	}

	return false
}

// SetVcsCommitMessage gets a reference to the given NullableString and assigns it to the VcsCommitMessage field.
func (o *Metadata) SetVcsCommitMessage(v string) {
	o.VcsCommitMessage.Set(&v)
}
// SetVcsCommitMessageNil sets the value for VcsCommitMessage to be an explicit nil
func (o *Metadata) SetVcsCommitMessageNil() {
	o.VcsCommitMessage.Set(nil)
}

// UnsetVcsCommitMessage ensures that no value is present for VcsCommitMessage, not even an explicit nil
func (o *Metadata) UnsetVcsCommitMessage() {
	o.VcsCommitMessage.Unset()
}

// GetVcsRepositoryUrl returns the VcsRepositoryUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsRepositoryUrl() string {
	if o == nil || IsNil(o.VcsRepositoryUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VcsRepositoryUrl.Get()
}

// GetVcsRepositoryUrlOk returns a tuple with the VcsRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsRepositoryUrl.Get(), o.VcsRepositoryUrl.IsSet()
}

// HasVcsRepositoryUrl returns a boolean if a field has been set.
func (o *Metadata) HasVcsRepositoryUrl() bool {
	if o != nil && o.VcsRepositoryUrl.IsSet() {
		return true
	}

	return false
}

// SetVcsRepositoryUrl gets a reference to the given NullableString and assigns it to the VcsRepositoryUrl field.
func (o *Metadata) SetVcsRepositoryUrl(v string) {
	o.VcsRepositoryUrl.Set(&v)
}
// SetVcsRepositoryUrlNil sets the value for VcsRepositoryUrl to be an explicit nil
func (o *Metadata) SetVcsRepositoryUrlNil() {
	o.VcsRepositoryUrl.Set(nil)
}

// UnsetVcsRepositoryUrl ensures that no value is present for VcsRepositoryUrl, not even an explicit nil
func (o *Metadata) UnsetVcsRepositoryUrl() {
	o.VcsRepositoryUrl.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *Metadata) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *Metadata) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *Metadata) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *Metadata) UnsetPath() {
	o.Path.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Metadata) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Metadata) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Metadata) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Metadata) UnsetType() {
	o.Type.Unset()
}

// GetVcsSubPath returns the VcsSubPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Metadata) GetVcsSubPath() string {
	if o == nil || IsNil(o.VcsSubPath.Get()) {
		var ret string
		return ret
	}
	return *o.VcsSubPath.Get()
}

// GetVcsSubPathOk returns a tuple with the VcsSubPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetVcsSubPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcsSubPath.Get(), o.VcsSubPath.IsSet()
}

// HasVcsSubPath returns a boolean if a field has been set.
func (o *Metadata) HasVcsSubPath() bool {
	if o != nil && o.VcsSubPath.IsSet() {
		return true
	}

	return false
}

// SetVcsSubPath gets a reference to the given NullableString and assigns it to the VcsSubPath field.
func (o *Metadata) SetVcsSubPath(v string) {
	o.VcsSubPath.Set(&v)
}
// SetVcsSubPathNil sets the value for VcsSubPath to be an explicit nil
func (o *Metadata) SetVcsSubPathNil() {
	o.VcsSubPath.Set(nil)
}

// UnsetVcsSubPath ensures that no value is present for VcsSubPath, not even an explicit nil
func (o *Metadata) UnsetVcsSubPath() {
	o.VcsSubPath.Unset()
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InfracostCommand.IsSet() {
		toSerialize["infracostCommand"] = o.InfracostCommand.Get()
	}
	if o.VcsBranch.IsSet() {
		toSerialize["vcsBranch"] = o.VcsBranch.Get()
	}
	if o.VcsCommitSha.IsSet() {
		toSerialize["vcsCommitSha"] = o.VcsCommitSha.Get()
	}
	if o.VcsCommitAuthorName.IsSet() {
		toSerialize["vcsCommitAuthorName"] = o.VcsCommitAuthorName.Get()
	}
	if o.VcsCommitAuthorEmail.IsSet() {
		toSerialize["vcsCommitAuthorEmail"] = o.VcsCommitAuthorEmail.Get()
	}
	if !IsNil(o.VcsCommitTimestamp) {
		toSerialize["vcsCommitTimestamp"] = o.VcsCommitTimestamp
	}
	if o.VcsCommitMessage.IsSet() {
		toSerialize["vcsCommitMessage"] = o.VcsCommitMessage.Get()
	}
	if o.VcsRepositoryUrl.IsSet() {
		toSerialize["vcsRepositoryUrl"] = o.VcsRepositoryUrl.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.VcsSubPath.IsSet() {
		toSerialize["vcsSubPath"] = o.VcsSubPath.Get()
	}
	return toSerialize, nil
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


