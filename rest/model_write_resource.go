/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// WriteResource Write account resource
type WriteResource struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateKeyHash string `json:"state_key_hash"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Address string          `json:"address"`
	Data    AccountResource `json:"data"`
}

// NewWriteResource instantiates a new WriteResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteResource(type_ string, stateKeyHash string, address string, data AccountResource) *WriteResource {
	this := WriteResource{}
	this.Type = type_
	this.StateKeyHash = stateKeyHash
	this.Address = address
	this.Data = data
	return &this
}

// NewWriteResourceWithDefaults instantiates a new WriteResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteResourceWithDefaults() *WriteResource {
	this := WriteResource{}
	return &this
}

// GetType returns the Type field value
func (o *WriteResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WriteResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WriteResource) SetType(v string) {
	o.Type = v
}

// GetStateKeyHash returns the StateKeyHash field value
func (o *WriteResource) GetStateKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateKeyHash
}

// GetStateKeyHashOk returns a tuple with the StateKeyHash field value
// and a boolean to check if the value has been set.
func (o *WriteResource) GetStateKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateKeyHash, true
}

// SetStateKeyHash sets field value
func (o *WriteResource) SetStateKeyHash(v string) {
	o.StateKeyHash = v
}

// GetAddress returns the Address field value
func (o *WriteResource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *WriteResource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *WriteResource) SetAddress(v string) {
	o.Address = v
}

// GetData returns the Data field value
func (o *WriteResource) GetData() AccountResource {
	if o == nil {
		var ret AccountResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WriteResource) GetDataOk() (*AccountResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WriteResource) SetData(v AccountResource) {
	o.Data = v
}

func (o WriteResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state_key_hash"] = o.StateKeyHash
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWriteResource struct {
	value *WriteResource
	isSet bool
}

func (v NullableWriteResource) Get() *WriteResource {
	return v.value
}

func (v *NullableWriteResource) Set(val *WriteResource) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteResource) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteResource(val *WriteResource) *NullableWriteResource {
	return &NullableWriteResource{value: val, isSet: true}
}

func (v NullableWriteResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
