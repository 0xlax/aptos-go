/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// Account Core account resource, used for identifying account and transaction execution.
type Account struct {
	// Unsigned int64 type value
	SequenceNumber string `json:"sequence_number"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	AuthenticationKey string `json:"authentication_key"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(sequenceNumber string, authenticationKey string) *Account {
	this := Account{}
	this.SequenceNumber = sequenceNumber
	this.AuthenticationKey = authenticationKey
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *Account) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *Account) GetSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *Account) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetAuthenticationKey returns the AuthenticationKey field value
func (o *Account) GetAuthenticationKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationKey
}

// GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field value
// and a boolean to check if the value has been set.
func (o *Account) GetAuthenticationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationKey, true
}

// SetAuthenticationKey sets field value
func (o *Account) SetAuthenticationKey(v string) {
	o.AuthenticationKey = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if true {
		toSerialize["authentication_key"] = o.AuthenticationKey
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
