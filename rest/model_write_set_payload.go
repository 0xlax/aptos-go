/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// WriteSetPayload struct for WriteSetPayload
type WriteSetPayload struct {
	Type     string   `json:"type"`
	WriteSet WriteSet `json:"write_set"`
}

// NewWriteSetPayload instantiates a new WriteSetPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteSetPayload(type_ string, writeSet WriteSet) *WriteSetPayload {
	this := WriteSetPayload{}
	this.Type = type_
	this.WriteSet = writeSet
	return &this
}

// NewWriteSetPayloadWithDefaults instantiates a new WriteSetPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteSetPayloadWithDefaults() *WriteSetPayload {
	this := WriteSetPayload{}
	return &this
}

// GetType returns the Type field value
func (o *WriteSetPayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WriteSetPayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WriteSetPayload) SetType(v string) {
	o.Type = v
}

// GetWriteSet returns the WriteSet field value
func (o *WriteSetPayload) GetWriteSet() WriteSet {
	if o == nil {
		var ret WriteSet
		return ret
	}

	return o.WriteSet
}

// GetWriteSetOk returns a tuple with the WriteSet field value
// and a boolean to check if the value has been set.
func (o *WriteSetPayload) GetWriteSetOk() (*WriteSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteSet, true
}

// SetWriteSet sets field value
func (o *WriteSetPayload) SetWriteSet(v WriteSet) {
	o.WriteSet = v
}

func (o WriteSetPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["write_set"] = o.WriteSet
	}
	return json.Marshal(toSerialize)
}

type NullableWriteSetPayload struct {
	value *WriteSetPayload
	isSet bool
}

func (v NullableWriteSetPayload) Get() *WriteSetPayload {
	return v.value
}

func (v *NullableWriteSetPayload) Set(val *WriteSetPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteSetPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteSetPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteSetPayload(val *WriteSetPayload) *NullableWriteSetPayload {
	return &NullableWriteSetPayload{value: val, isSet: true}
}

func (v NullableWriteSetPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteSetPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
