/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SignatureCreationEventResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureCreationEventResult{}

// SignatureCreationEventResult The result to a SignatureCreationEvent.
type SignatureCreationEventResult struct {
	// The ID of the wallet that the event was created for.
	WalletId string `json:"wallet_id"`
	// The ID of the user that the wallet belongs to
	WalletUserId string `json:"wallet_user_id"`
	// The ID of the address the transfer belongs to
	AddressId string `json:"address_id"`
	TransactionType TransactionType `json:"transaction_type"`
	// The ID of the transaction that the Server-Signer has signed for
	TransactionId string `json:"transaction_id"`
	// The signature created by the server-signer.
	Signature string `json:"signature"`
}

type _SignatureCreationEventResult SignatureCreationEventResult

// NewSignatureCreationEventResult instantiates a new SignatureCreationEventResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureCreationEventResult(walletId string, walletUserId string, addressId string, transactionType TransactionType, transactionId string, signature string) *SignatureCreationEventResult {
	this := SignatureCreationEventResult{}
	this.WalletId = walletId
	this.WalletUserId = walletUserId
	this.AddressId = addressId
	this.TransactionType = transactionType
	this.TransactionId = transactionId
	this.Signature = signature
	return &this
}

// NewSignatureCreationEventResultWithDefaults instantiates a new SignatureCreationEventResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureCreationEventResultWithDefaults() *SignatureCreationEventResult {
	this := SignatureCreationEventResult{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *SignatureCreationEventResult) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SignatureCreationEventResult) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletUserId returns the WalletUserId field value
func (o *SignatureCreationEventResult) GetWalletUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletUserId
}

// GetWalletUserIdOk returns a tuple with the WalletUserId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetWalletUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletUserId, true
}

// SetWalletUserId sets field value
func (o *SignatureCreationEventResult) SetWalletUserId(v string) {
	o.WalletUserId = v
}

// GetAddressId returns the AddressId field value
func (o *SignatureCreationEventResult) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *SignatureCreationEventResult) SetAddressId(v string) {
	o.AddressId = v
}

// GetTransactionType returns the TransactionType field value
func (o *SignatureCreationEventResult) GetTransactionType() TransactionType {
	if o == nil {
		var ret TransactionType
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetTransactionTypeOk() (*TransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *SignatureCreationEventResult) SetTransactionType(v TransactionType) {
	o.TransactionType = v
}

// GetTransactionId returns the TransactionId field value
func (o *SignatureCreationEventResult) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *SignatureCreationEventResult) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetSignature returns the Signature field value
func (o *SignatureCreationEventResult) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEventResult) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *SignatureCreationEventResult) SetSignature(v string) {
	o.Signature = v
}

func (o SignatureCreationEventResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureCreationEventResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_user_id"] = o.WalletUserId
	toSerialize["address_id"] = o.AddressId
	toSerialize["transaction_type"] = o.TransactionType
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

func (o *SignatureCreationEventResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"wallet_user_id",
		"address_id",
		"transaction_type",
		"transaction_id",
		"signature",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSignatureCreationEventResult := _SignatureCreationEventResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureCreationEventResult)

	if err != nil {
		return err
	}

	*o = SignatureCreationEventResult(varSignatureCreationEventResult)

	return err
}

type NullableSignatureCreationEventResult struct {
	value *SignatureCreationEventResult
	isSet bool
}

func (v NullableSignatureCreationEventResult) Get() *SignatureCreationEventResult {
	return v.value
}

func (v *NullableSignatureCreationEventResult) Set(val *SignatureCreationEventResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureCreationEventResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureCreationEventResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureCreationEventResult(val *SignatureCreationEventResult) *NullableSignatureCreationEventResult {
	return &NullableSignatureCreationEventResult{value: val, isSet: true}
}

func (v NullableSignatureCreationEventResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureCreationEventResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


