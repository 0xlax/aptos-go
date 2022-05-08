/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// DirectWriteSet struct for DirectWriteSet
type DirectWriteSet struct {
	Type    string           `json:"type"`
	Changes []WriteSetChange `json:"changes"`
	Events  []Event          `json:"events"`
}

// NewDirectWriteSet instantiates a new DirectWriteSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectWriteSet(type_ string, changes []WriteSetChange, events []Event) *DirectWriteSet {
	this := DirectWriteSet{}
	this.Type = type_
	this.Changes = changes
	this.Events = events
	return &this
}

// NewDirectWriteSetWithDefaults instantiates a new DirectWriteSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectWriteSetWithDefaults() *DirectWriteSet {
	this := DirectWriteSet{}
	return &this
}

// GetType returns the Type field value
func (o *DirectWriteSet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DirectWriteSet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DirectWriteSet) SetType(v string) {
	o.Type = v
}

// GetChanges returns the Changes field value
func (o *DirectWriteSet) GetChanges() []WriteSetChange {
	if o == nil {
		var ret []WriteSetChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *DirectWriteSet) GetChangesOk() ([]WriteSetChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *DirectWriteSet) SetChanges(v []WriteSetChange) {
	o.Changes = v
}

// GetEvents returns the Events field value
func (o *DirectWriteSet) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *DirectWriteSet) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *DirectWriteSet) SetEvents(v []Event) {
	o.Events = v
}

func (o DirectWriteSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["changes"] = o.Changes
	}
	if true {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableDirectWriteSet struct {
	value *DirectWriteSet
	isSet bool
}

func (v NullableDirectWriteSet) Get() *DirectWriteSet {
	return v.value
}

func (v *NullableDirectWriteSet) Set(val *DirectWriteSet) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectWriteSet) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectWriteSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectWriteSet(val *DirectWriteSet) *NullableDirectWriteSet {
	return &NullableDirectWriteSet{value: val, isSet: true}
}

func (v NullableDirectWriteSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectWriteSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
