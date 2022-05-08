/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// GenesisTransaction struct for GenesisTransaction
type GenesisTransaction struct {
	Type    string          `json:"type"`
	Events  []Event         `json:"events"`
	Payload WriteSetPayload `json:"payload"`
	// Unsigned int64 type value
	Version string `json:"version"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Hash string `json:"hash"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateRootHash string `json:"state_root_hash"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	EventRootHash string `json:"event_root_hash"`
	// Unsigned int64 type value
	GasUsed string `json:"gas_used"`
	// Transaction execution result (success: true, failure: false). See `vm_status` for human readable error message from Aptos VM.
	Success bool `json:"success"`
	// Human readable transaction execution result message from Aptos VM.
	VmStatus string `json:"vm_status"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	AccumulatorRootHash string           `json:"accumulator_root_hash"`
	Changes             []WriteSetChange `json:"changes"`
}

// NewGenesisTransaction instantiates a new GenesisTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenesisTransaction(type_ string, events []Event, payload WriteSetPayload, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange) *GenesisTransaction {
	this := GenesisTransaction{}
	this.Type = type_
	this.Events = events
	this.Payload = payload
	this.Version = version
	this.Hash = hash
	this.StateRootHash = stateRootHash
	this.EventRootHash = eventRootHash
	this.GasUsed = gasUsed
	this.Success = success
	this.VmStatus = vmStatus
	this.AccumulatorRootHash = accumulatorRootHash
	this.Changes = changes
	return &this
}

// NewGenesisTransactionWithDefaults instantiates a new GenesisTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenesisTransactionWithDefaults() *GenesisTransaction {
	this := GenesisTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *GenesisTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GenesisTransaction) SetType(v string) {
	o.Type = v
}

// GetEvents returns the Events field value
func (o *GenesisTransaction) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *GenesisTransaction) SetEvents(v []Event) {
	o.Events = v
}

// GetPayload returns the Payload field value
func (o *GenesisTransaction) GetPayload() WriteSetPayload {
	if o == nil {
		var ret WriteSetPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetPayloadOk() (*WriteSetPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *GenesisTransaction) SetPayload(v WriteSetPayload) {
	o.Payload = v
}

// GetVersion returns the Version field value
func (o *GenesisTransaction) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GenesisTransaction) SetVersion(v string) {
	o.Version = v
}

// GetHash returns the Hash field value
func (o *GenesisTransaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *GenesisTransaction) SetHash(v string) {
	o.Hash = v
}

// GetStateRootHash returns the StateRootHash field value
func (o *GenesisTransaction) GetStateRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRootHash
}

// GetStateRootHashOk returns a tuple with the StateRootHash field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetStateRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRootHash, true
}

// SetStateRootHash sets field value
func (o *GenesisTransaction) SetStateRootHash(v string) {
	o.StateRootHash = v
}

// GetEventRootHash returns the EventRootHash field value
func (o *GenesisTransaction) GetEventRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventRootHash
}

// GetEventRootHashOk returns a tuple with the EventRootHash field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetEventRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventRootHash, true
}

// SetEventRootHash sets field value
func (o *GenesisTransaction) SetEventRootHash(v string) {
	o.EventRootHash = v
}

// GetGasUsed returns the GasUsed field value
func (o *GenesisTransaction) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GenesisTransaction) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetSuccess returns the Success field value
func (o *GenesisTransaction) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *GenesisTransaction) SetSuccess(v bool) {
	o.Success = v
}

// GetVmStatus returns the VmStatus field value
func (o *GenesisTransaction) GetVmStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmStatus
}

// GetVmStatusOk returns a tuple with the VmStatus field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetVmStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmStatus, true
}

// SetVmStatus sets field value
func (o *GenesisTransaction) SetVmStatus(v string) {
	o.VmStatus = v
}

// GetAccumulatorRootHash returns the AccumulatorRootHash field value
func (o *GenesisTransaction) GetAccumulatorRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccumulatorRootHash
}

// GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetAccumulatorRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccumulatorRootHash, true
}

// SetAccumulatorRootHash sets field value
func (o *GenesisTransaction) SetAccumulatorRootHash(v string) {
	o.AccumulatorRootHash = v
}

// GetChanges returns the Changes field value
func (o *GenesisTransaction) GetChanges() []WriteSetChange {
	if o == nil {
		var ret []WriteSetChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *GenesisTransaction) GetChangesOk() ([]WriteSetChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *GenesisTransaction) SetChanges(v []WriteSetChange) {
	o.Changes = v
}

func (o GenesisTransaction) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["state_root_hash"] = o.StateRootHash
	}
	if true {
		toSerialize["event_root_hash"] = o.EventRootHash
	}
	if true {
		toSerialize["gas_used"] = o.GasUsed
	}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["vm_status"] = o.VmStatus
	}
	if true {
		toSerialize["accumulator_root_hash"] = o.AccumulatorRootHash
	}
	if true {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullableGenesisTransaction struct {
	value *GenesisTransaction
	isSet bool
}

func (v NullableGenesisTransaction) Get() *GenesisTransaction {
	return v.value
}

func (v *NullableGenesisTransaction) Set(val *GenesisTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableGenesisTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableGenesisTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenesisTransaction(val *GenesisTransaction) *NullableGenesisTransaction {
	return &NullableGenesisTransaction{value: val, isSet: true}
}

func (v NullableGenesisTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenesisTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
