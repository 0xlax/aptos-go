/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// GenesisTransactionAllOf struct for GenesisTransactionAllOf
type GenesisTransactionAllOf struct {
	Type    string          `json:"type"`
	Events  []Event         `json:"events"`
	Payload WriteSetPayload `json:"payload"`
}

// NewGenesisTransactionAllOf instantiates a new GenesisTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenesisTransactionAllOf(type_ string, events []Event, payload WriteSetPayload) *GenesisTransactionAllOf {
	this := GenesisTransactionAllOf{}
	this.Type = type_
	this.Events = events
	this.Payload = payload
	return &this
}

// NewGenesisTransactionAllOfWithDefaults instantiates a new GenesisTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenesisTransactionAllOfWithDefaults() *GenesisTransactionAllOf {
	this := GenesisTransactionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *GenesisTransactionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GenesisTransactionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GenesisTransactionAllOf) SetType(v string) {
	o.Type = v
}

// GetEvents returns the Events field value
func (o *GenesisTransactionAllOf) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *GenesisTransactionAllOf) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *GenesisTransactionAllOf) SetEvents(v []Event) {
	o.Events = v
}

// GetPayload returns the Payload field value
func (o *GenesisTransactionAllOf) GetPayload() WriteSetPayload {
	if o == nil {
		var ret WriteSetPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *GenesisTransactionAllOf) GetPayloadOk() (*WriteSetPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *GenesisTransactionAllOf) SetPayload(v WriteSetPayload) {
	o.Payload = v
}

func (o GenesisTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableGenesisTransactionAllOf struct {
	value *GenesisTransactionAllOf
	isSet bool
}

func (v NullableGenesisTransactionAllOf) Get() *GenesisTransactionAllOf {
	return v.value
}

func (v *NullableGenesisTransactionAllOf) Set(val *GenesisTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGenesisTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGenesisTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenesisTransactionAllOf(val *GenesisTransactionAllOf) *NullableGenesisTransactionAllOf {
	return &NullableGenesisTransactionAllOf{value: val, isSet: true}
}

func (v NullableGenesisTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenesisTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
