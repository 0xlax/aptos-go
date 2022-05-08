/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// UserTransactionSignature This schema is used for appending `signature` field to another schema.
type UserTransactionSignature struct {
	Signature TransactionSignature `json:"signature"`
}

// NewUserTransactionSignature instantiates a new UserTransactionSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransactionSignature(signature TransactionSignature) *UserTransactionSignature {
	this := UserTransactionSignature{}
	this.Signature = signature
	return &this
}

// NewUserTransactionSignatureWithDefaults instantiates a new UserTransactionSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransactionSignatureWithDefaults() *UserTransactionSignature {
	this := UserTransactionSignature{}
	return &this
}

// GetSignature returns the Signature field value
func (o *UserTransactionSignature) GetSignature() TransactionSignature {
	if o == nil {
		var ret TransactionSignature
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *UserTransactionSignature) GetSignatureOk() (*TransactionSignature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *UserTransactionSignature) SetSignature(v TransactionSignature) {
	o.Signature = v
}

func (o UserTransactionSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableUserTransactionSignature struct {
	value *UserTransactionSignature
	isSet bool
}

func (v NullableUserTransactionSignature) Get() *UserTransactionSignature {
	return v.value
}

func (v *NullableUserTransactionSignature) Set(val *UserTransactionSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransactionSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransactionSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransactionSignature(val *UserTransactionSignature) *NullableUserTransactionSignature {
	return &NullableUserTransactionSignature{value: val, isSet: true}
}

func (v NullableUserTransactionSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransactionSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
