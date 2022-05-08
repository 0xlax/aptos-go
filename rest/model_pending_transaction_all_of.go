/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// PendingTransactionAllOf struct for PendingTransactionAllOf
type PendingTransactionAllOf struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Hash string `json:"hash"`
}

// NewPendingTransactionAllOf instantiates a new PendingTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactionAllOf(type_ string, hash string) *PendingTransactionAllOf {
	this := PendingTransactionAllOf{}
	this.Type = type_
	this.Hash = hash
	return &this
}

// NewPendingTransactionAllOfWithDefaults instantiates a new PendingTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionAllOfWithDefaults() *PendingTransactionAllOf {
	this := PendingTransactionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *PendingTransactionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PendingTransactionAllOf) SetType(v string) {
	o.Type = v
}

// GetHash returns the Hash field value
func (o *PendingTransactionAllOf) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionAllOf) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *PendingTransactionAllOf) SetHash(v string) {
	o.Hash = v
}

func (o PendingTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullablePendingTransactionAllOf struct {
	value *PendingTransactionAllOf
	isSet bool
}

func (v NullablePendingTransactionAllOf) Get() *PendingTransactionAllOf {
	return v.value
}

func (v *NullablePendingTransactionAllOf) Set(val *PendingTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactionAllOf(val *PendingTransactionAllOf) *NullablePendingTransactionAllOf {
	return &NullablePendingTransactionAllOf{value: val, isSet: true}
}

func (v NullablePendingTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
