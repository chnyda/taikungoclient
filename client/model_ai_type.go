/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
	"fmt"
)

// AiType the model 'AiType'
type AiType string

// List of AiType
const (
	AITYPE_TAIKUN AiType = "Taikun"
	AITYPE_OPEN_AI AiType = "OpenAi"
)

// All allowed values of AiType enum
var AllowedAiTypeEnumValues = []AiType{
	"Taikun",
	"OpenAi",
}

func (v *AiType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiType(value)
	for _, existing := range AllowedAiTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiType", value)
}

// NewAiTypeFromValue returns a pointer to a valid AiType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiTypeFromValue(v string) (*AiType, error) {
	ev := AiType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiType: valid values are %v", v, AllowedAiTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiType) IsValid() bool {
	for _, existing := range AllowedAiTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiType value
func (v AiType) Ptr() *AiType {
	return &v
}

type NullableAiType struct {
	value *AiType
	isSet bool
}

func (v NullableAiType) Get() *AiType {
	return v.value
}

func (v *NullableAiType) Set(val *AiType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiType(val *AiType) *NullableAiType {
	return &NullableAiType{value: val, isSet: true}
}

func (v NullableAiType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
