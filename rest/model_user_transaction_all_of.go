/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// UserTransactionAllOf struct for UserTransactionAllOf
type UserTransactionAllOf struct {
	Type   string  `json:"type"`
	Events []Event `json:"events"`
	// Timestamp in microseconds, e.g. ledger / block creation timestamp.
	Timestamp string `json:"timestamp"`
}

// NewUserTransactionAllOf instantiates a new UserTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransactionAllOf(type_ string, events []Event, timestamp string) *UserTransactionAllOf {
	this := UserTransactionAllOf{}
	this.Type = type_
	this.Events = events
	this.Timestamp = timestamp
	return &this
}

// NewUserTransactionAllOfWithDefaults instantiates a new UserTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransactionAllOfWithDefaults() *UserTransactionAllOf {
	this := UserTransactionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *UserTransactionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserTransactionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserTransactionAllOf) SetType(v string) {
	o.Type = v
}

// GetEvents returns the Events field value
func (o *UserTransactionAllOf) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *UserTransactionAllOf) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *UserTransactionAllOf) SetEvents(v []Event) {
	o.Events = v
}

// GetTimestamp returns the Timestamp field value
func (o *UserTransactionAllOf) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UserTransactionAllOf) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UserTransactionAllOf) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o UserTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableUserTransactionAllOf struct {
	value *UserTransactionAllOf
	isSet bool
}

func (v NullableUserTransactionAllOf) Get() *UserTransactionAllOf {
	return v.value
}

func (v *NullableUserTransactionAllOf) Set(val *UserTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransactionAllOf(val *UserTransactionAllOf) *NullableUserTransactionAllOf {
	return &NullableUserTransactionAllOf{value: val, isSet: true}
}

func (v NullableUserTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
