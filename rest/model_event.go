/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// Event Event `key` and `sequence_number` are global identifier of the event.  Event `sequence_number` starts from 0 for each event key.  Event `type` is the type information of the event `data`, you can use the `type` to decode the `data` JSON.
type Event struct {
	// Event key is a global index for an event stream.  It is hex-encoded BCS bytes of `EventHandle` `guid` field value, which is a combination of a `uint64` creation number and account address (without trimming leading zeros).  For example, event key `0x00000000000000000000000000000000000000000a550c18` is combined by the following 2 parts:   1. `0000000000000000`: `uint64` representation of `0`.   2. `0000000000000000000000000a550c18`: 16 bytes of account address.
	Key string `json:"key"`
	// Event `sequence_number` is unique id of an event in an event stream. Event `sequence_number` starts from 0 for each event key.
	SequenceNumber string `json:"sequence_number"`
	// String representation of an on-chain Move type tag that is exposed in transaction payload.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: `vector<{non-reference MoveTypeId}>`   - struct: `{address}::{module_name}::{struct_name}::<{generic types}>`  Vector type value examples:   * `vector<u8>`   * `vector<vector<u64>>`   * `vector<0x1::AptosAccount::Balance<0x1::XDX::XDX>>`  Struct type value examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).
	Type string `json:"type"`
	// Move `bool` type value is serialized into `boolean`.  Move `u8` type value is serialized into `integer`.  Move `u64` and `u128` type value is serialized into `string`.  Move `address` type value(16 bytes Aptos account address) is serialized into hex-encoded string, which is prefixed with `0x` and leading zeros are trimmed.  For example:   * `0x1`   * `0x1668f6be25668c1a17cd8caf6b8d2f25`  Move `vector` type value is serialized into `array`, except `vector<u8>` which is serialized into hex-encoded string with `0x` prefix.  For example:   * `vector<u64>{255, 255}` => `[\"255\", \"255\"]`   * `vector<u8>{255, 255}` => `0xffff`  Move `struct` type value is serialized into `object` that looks like this (except some Move stdlib types, see the following section):    ```json   {     field1_name: field1_value,     field2_name: field2_value,     ......   }   ```  For example:   `{ \"created\": \"0xa550c18\", \"role_id\": \"0\" }`  **Special serialization for Move stdlib types:**  * [0x1::ASCII::String](https://github.com/aptos-labs/aptos-core/blob/main/language/move-stdlib/docs/ASCII.md) is serialized into `string`. For example, struct value `0x1::ASCII::String{bytes: b\"hello world\"}` is serialized as `\"hello world\"` in JSON.
	Data interface{} `json:"data"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(key string, sequenceNumber string, type_ string, data interface{}) *Event {
	this := Event{}
	this.Key = key
	this.SequenceNumber = sequenceNumber
	this.Type = type_
	this.Data = data
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetKey returns the Key field value
func (o *Event) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Event) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Event) SetKey(v string) {
	o.Key = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *Event) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *Event) GetSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *Event) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetType returns the Type field value
func (o *Event) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Event) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Event) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Event) SetData(v interface{}) {
	o.Data = v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
