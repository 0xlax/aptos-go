/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// TableItemDeletion struct for TableItemDeletion
type TableItemDeletion struct {
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Handle string `json:"handle"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Key string `json:"key"`
}

// NewTableItemDeletion instantiates a new TableItemDeletion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableItemDeletion(handle string, key string) *TableItemDeletion {
	this := TableItemDeletion{}
	this.Handle = handle
	this.Key = key
	return &this
}

// NewTableItemDeletionWithDefaults instantiates a new TableItemDeletion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableItemDeletionWithDefaults() *TableItemDeletion {
	this := TableItemDeletion{}
	return &this
}

// GetHandle returns the Handle field value
func (o *TableItemDeletion) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TableItemDeletion) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *TableItemDeletion) SetHandle(v string) {
	o.Handle = v
}

// GetKey returns the Key field value
func (o *TableItemDeletion) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TableItemDeletion) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TableItemDeletion) SetKey(v string) {
	o.Key = v
}

func (o TableItemDeletion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableTableItemDeletion struct {
	value *TableItemDeletion
	isSet bool
}

func (v NullableTableItemDeletion) Get() *TableItemDeletion {
	return v.value
}

func (v *NullableTableItemDeletion) Set(val *TableItemDeletion) {
	v.value = val
	v.isSet = true
}

func (v NullableTableItemDeletion) IsSet() bool {
	return v.isSet
}

func (v *NullableTableItemDeletion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableItemDeletion(val *TableItemDeletion) *NullableTableItemDeletion {
	return &NullableTableItemDeletion{value: val, isSet: true}
}

func (v NullableTableItemDeletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableItemDeletion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
