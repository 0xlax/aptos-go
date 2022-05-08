/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// DeleteModule struct for DeleteModule
type DeleteModule struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateKeyHash string `json:"state_key_hash"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Address string `json:"address"`
	// Move module id is a string representation of Move module.  Format: \"{address}::{module name}\"  `address` should be hex-encoded 16 bytes account address that is prefixed with `0x` and leading zeros are trimmed.  Module name is case-sensitive.  See [doc](https://diem.github.io/move/modules-and-scripts.html#modules) for more details.
	Module string `json:"module"`
}

// NewDeleteModule instantiates a new DeleteModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteModule(type_ string, stateKeyHash string, address string, module string) *DeleteModule {
	this := DeleteModule{}
	this.Type = type_
	this.StateKeyHash = stateKeyHash
	this.Address = address
	this.Module = module
	return &this
}

// NewDeleteModuleWithDefaults instantiates a new DeleteModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteModuleWithDefaults() *DeleteModule {
	this := DeleteModule{}
	return &this
}

// GetType returns the Type field value
func (o *DeleteModule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteModule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeleteModule) SetType(v string) {
	o.Type = v
}

// GetStateKeyHash returns the StateKeyHash field value
func (o *DeleteModule) GetStateKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateKeyHash
}

// GetStateKeyHashOk returns a tuple with the StateKeyHash field value
// and a boolean to check if the value has been set.
func (o *DeleteModule) GetStateKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateKeyHash, true
}

// SetStateKeyHash sets field value
func (o *DeleteModule) SetStateKeyHash(v string) {
	o.StateKeyHash = v
}

// GetAddress returns the Address field value
func (o *DeleteModule) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DeleteModule) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DeleteModule) SetAddress(v string) {
	o.Address = v
}

// GetModule returns the Module field value
func (o *DeleteModule) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *DeleteModule) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *DeleteModule) SetModule(v string) {
	o.Module = v
}

func (o DeleteModule) MarshalJSON() ([]byte, error) {
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
		toSerialize["module"] = o.Module
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteModule struct {
	value *DeleteModule
	isSet bool
}

func (v NullableDeleteModule) Get() *DeleteModule {
	return v.value
}

func (v *NullableDeleteModule) Set(val *DeleteModule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteModule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteModule(val *DeleteModule) *NullableDeleteModule {
	return &NullableDeleteModule{value: val, isSet: true}
}

func (v NullableDeleteModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
