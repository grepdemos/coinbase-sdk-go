/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// WebhookEventType the model 'WebhookEventType'
type WebhookEventType string

// List of WebhookEventType
const (
	WEBHOOKEVENTTYPE_UNSPECIFIED WebhookEventType = "unspecified"
	WEBHOOKEVENTTYPE_ERC20_TRANSFER WebhookEventType = "erc20_transfer"
	WEBHOOKEVENTTYPE_ERC721_TRANSFER WebhookEventType = "erc721_transfer"
)

// All allowed values of WebhookEventType enum
var AllowedWebhookEventTypeEnumValues = []WebhookEventType{
	"unspecified",
	"erc20_transfer",
	"erc721_transfer",
}

func (v *WebhookEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookEventType(value)
	for _, existing := range AllowedWebhookEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookEventType", value)
}

// NewWebhookEventTypeFromValue returns a pointer to a valid WebhookEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookEventTypeFromValue(v string) (*WebhookEventType, error) {
	ev := WebhookEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookEventType: valid values are %v", v, AllowedWebhookEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookEventType) IsValid() bool {
	for _, existing := range AllowedWebhookEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookEventType value
func (v WebhookEventType) Ptr() *WebhookEventType {
	return &v
}

type NullableWebhookEventType struct {
	value *WebhookEventType
	isSet bool
}

func (v NullableWebhookEventType) Get() *WebhookEventType {
	return v.value
}

func (v *NullableWebhookEventType) Set(val *WebhookEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventType(val *WebhookEventType) *NullableWebhookEventType {
	return &NullableWebhookEventType{value: val, isSet: true}
}

func (v NullableWebhookEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

