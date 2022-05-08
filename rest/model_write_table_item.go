/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// WriteTableItem Write table item
type WriteTableItem struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateKeyHash string         `json:"state_key_hash"`
	Data         TableItemWrite `json:"data"`
}

// NewWriteTableItem instantiates a new WriteTableItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteTableItem(type_ string, stateKeyHash string, data TableItemWrite) *WriteTableItem {
	this := WriteTableItem{}
	this.Type = type_
	this.StateKeyHash = stateKeyHash
	this.Data = data
	return &this
}

// NewWriteTableItemWithDefaults instantiates a new WriteTableItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteTableItemWithDefaults() *WriteTableItem {
	this := WriteTableItem{}
	return &this
}

// GetType returns the Type field value
func (o *WriteTableItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WriteTableItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WriteTableItem) SetType(v string) {
	o.Type = v
}

// GetStateKeyHash returns the StateKeyHash field value
func (o *WriteTableItem) GetStateKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateKeyHash
}

// GetStateKeyHashOk returns a tuple with the StateKeyHash field value
// and a boolean to check if the value has been set.
func (o *WriteTableItem) GetStateKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateKeyHash, true
}

// SetStateKeyHash sets field value
func (o *WriteTableItem) SetStateKeyHash(v string) {
	o.StateKeyHash = v
}

// GetData returns the Data field value
func (o *WriteTableItem) GetData() TableItemWrite {
	if o == nil {
		var ret TableItemWrite
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WriteTableItem) GetDataOk() (*TableItemWrite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WriteTableItem) SetData(v TableItemWrite) {
	o.Data = v
}

func (o WriteTableItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state_key_hash"] = o.StateKeyHash
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWriteTableItem struct {
	value *WriteTableItem
	isSet bool
}

func (v NullableWriteTableItem) Get() *WriteTableItem {
	return v.value
}

func (v *NullableWriteTableItem) Set(val *WriteTableItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteTableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteTableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteTableItem(val *WriteTableItem) *NullableWriteTableItem {
	return &NullableWriteTableItem{value: val, isSet: true}
}

func (v NullableWriteTableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteTableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
