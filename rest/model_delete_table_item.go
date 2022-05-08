/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// DeleteTableItem Delete table item change.
type DeleteTableItem struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateKeyHash string            `json:"state_key_hash"`
	Data         TableItemDeletion `json:"data"`
}

// NewDeleteTableItem instantiates a new DeleteTableItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTableItem(type_ string, stateKeyHash string, data TableItemDeletion) *DeleteTableItem {
	this := DeleteTableItem{}
	this.Type = type_
	this.StateKeyHash = stateKeyHash
	this.Data = data
	return &this
}

// NewDeleteTableItemWithDefaults instantiates a new DeleteTableItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTableItemWithDefaults() *DeleteTableItem {
	this := DeleteTableItem{}
	return &this
}

// GetType returns the Type field value
func (o *DeleteTableItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteTableItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeleteTableItem) SetType(v string) {
	o.Type = v
}

// GetStateKeyHash returns the StateKeyHash field value
func (o *DeleteTableItem) GetStateKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateKeyHash
}

// GetStateKeyHashOk returns a tuple with the StateKeyHash field value
// and a boolean to check if the value has been set.
func (o *DeleteTableItem) GetStateKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateKeyHash, true
}

// SetStateKeyHash sets field value
func (o *DeleteTableItem) SetStateKeyHash(v string) {
	o.StateKeyHash = v
}

// GetData returns the Data field value
func (o *DeleteTableItem) GetData() TableItemDeletion {
	if o == nil {
		var ret TableItemDeletion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteTableItem) GetDataOk() (*TableItemDeletion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteTableItem) SetData(v TableItemDeletion) {
	o.Data = v
}

func (o DeleteTableItem) MarshalJSON() ([]byte, error) {
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

type NullableDeleteTableItem struct {
	value *DeleteTableItem
	isSet bool
}

func (v NullableDeleteTableItem) Get() *DeleteTableItem {
	return v.value
}

func (v *NullableDeleteTableItem) Set(val *DeleteTableItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTableItem(val *DeleteTableItem) *NullableDeleteTableItem {
	return &NullableDeleteTableItem{value: val, isSet: true}
}

func (v NullableDeleteTableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
