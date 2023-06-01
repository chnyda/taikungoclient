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

// checks if the EditArticleCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditArticleCommand{}

// EditArticleCommand struct for EditArticleCommand
type EditArticleCommand struct {
	MessageId string `json:"messageId"`
	Body string `json:"body"`
}

// NewEditArticleCommand instantiates a new EditArticleCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditArticleCommand(messageId string, body string) *EditArticleCommand {
	this := EditArticleCommand{}
	this.MessageId = messageId
	this.Body = body
	return &this
}

// NewEditArticleCommandWithDefaults instantiates a new EditArticleCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditArticleCommandWithDefaults() *EditArticleCommand {
	this := EditArticleCommand{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *EditArticleCommand) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *EditArticleCommand) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *EditArticleCommand) SetMessageId(v string) {
	o.MessageId = v
}

// GetBody returns the Body field value
func (o *EditArticleCommand) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EditArticleCommand) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EditArticleCommand) SetBody(v string) {
	o.Body = v
}

func (o EditArticleCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditArticleCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageId"] = o.MessageId
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullableEditArticleCommand struct {
	value *EditArticleCommand
	isSet bool
}

func (v NullableEditArticleCommand) Get() *EditArticleCommand {
	return v.value
}

func (v *NullableEditArticleCommand) Set(val *EditArticleCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditArticleCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditArticleCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditArticleCommand(val *EditArticleCommand) *NullableEditArticleCommand {
	return &NullableEditArticleCommand{value: val, isSet: true}
}

func (v NullableEditArticleCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditArticleCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

