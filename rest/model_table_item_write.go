/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// TableItemWrite struct for TableItemWrite
type TableItemWrite struct {
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Handle string `json:"handle"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Key string `json:"key"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Value string `json:"value"`
}

// NewTableItemWrite instantiates a new TableItemWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableItemWrite(handle string, key string, value string) *TableItemWrite {
	this := TableItemWrite{}
	this.Handle = handle
	this.Key = key
	this.Value = value
	return &this
}

// NewTableItemWriteWithDefaults instantiates a new TableItemWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableItemWriteWithDefaults() *TableItemWrite {
	this := TableItemWrite{}
	return &this
}

// GetHandle returns the Handle field value
func (o *TableItemWrite) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TableItemWrite) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *TableItemWrite) SetHandle(v string) {
	o.Handle = v
}

// GetKey returns the Key field value
func (o *TableItemWrite) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TableItemWrite) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TableItemWrite) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *TableItemWrite) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TableItemWrite) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TableItemWrite) SetValue(v string) {
	o.Value = v
}

func (o TableItemWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTableItemWrite struct {
	value *TableItemWrite
	isSet bool
}

func (v NullableTableItemWrite) Get() *TableItemWrite {
	return v.value
}

func (v *NullableTableItemWrite) Set(val *TableItemWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableTableItemWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableTableItemWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableItemWrite(val *TableItemWrite) *NullableTableItemWrite {
	return &NullableTableItemWrite{value: val, isSet: true}
}

func (v NullableTableItemWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableItemWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
