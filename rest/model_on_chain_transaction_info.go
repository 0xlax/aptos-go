/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// OnChainTransactionInfo struct for OnChainTransactionInfo
type OnChainTransactionInfo struct {
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

// NewOnChainTransactionInfo instantiates a new OnChainTransactionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainTransactionInfo(version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange) *OnChainTransactionInfo {
	this := OnChainTransactionInfo{}
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

// NewOnChainTransactionInfoWithDefaults instantiates a new OnChainTransactionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainTransactionInfoWithDefaults() *OnChainTransactionInfo {
	this := OnChainTransactionInfo{}
	return &this
}

// GetVersion returns the Version field value
func (o *OnChainTransactionInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OnChainTransactionInfo) SetVersion(v string) {
	o.Version = v
}

// GetHash returns the Hash field value
func (o *OnChainTransactionInfo) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *OnChainTransactionInfo) SetHash(v string) {
	o.Hash = v
}

// GetStateRootHash returns the StateRootHash field value
func (o *OnChainTransactionInfo) GetStateRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRootHash
}

// GetStateRootHashOk returns a tuple with the StateRootHash field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetStateRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRootHash, true
}

// SetStateRootHash sets field value
func (o *OnChainTransactionInfo) SetStateRootHash(v string) {
	o.StateRootHash = v
}

// GetEventRootHash returns the EventRootHash field value
func (o *OnChainTransactionInfo) GetEventRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventRootHash
}

// GetEventRootHashOk returns a tuple with the EventRootHash field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetEventRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventRootHash, true
}

// SetEventRootHash sets field value
func (o *OnChainTransactionInfo) SetEventRootHash(v string) {
	o.EventRootHash = v
}

// GetGasUsed returns the GasUsed field value
func (o *OnChainTransactionInfo) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *OnChainTransactionInfo) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetSuccess returns the Success field value
func (o *OnChainTransactionInfo) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *OnChainTransactionInfo) SetSuccess(v bool) {
	o.Success = v
}

// GetVmStatus returns the VmStatus field value
func (o *OnChainTransactionInfo) GetVmStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmStatus
}

// GetVmStatusOk returns a tuple with the VmStatus field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetVmStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmStatus, true
}

// SetVmStatus sets field value
func (o *OnChainTransactionInfo) SetVmStatus(v string) {
	o.VmStatus = v
}

// GetAccumulatorRootHash returns the AccumulatorRootHash field value
func (o *OnChainTransactionInfo) GetAccumulatorRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccumulatorRootHash
}

// GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetAccumulatorRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccumulatorRootHash, true
}

// SetAccumulatorRootHash sets field value
func (o *OnChainTransactionInfo) SetAccumulatorRootHash(v string) {
	o.AccumulatorRootHash = v
}

// GetChanges returns the Changes field value
func (o *OnChainTransactionInfo) GetChanges() []WriteSetChange {
	if o == nil {
		var ret []WriteSetChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *OnChainTransactionInfo) GetChangesOk() ([]WriteSetChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *OnChainTransactionInfo) SetChanges(v []WriteSetChange) {
	o.Changes = v
}

func (o OnChainTransactionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableOnChainTransactionInfo struct {
	value *OnChainTransactionInfo
	isSet bool
}

func (v NullableOnChainTransactionInfo) Get() *OnChainTransactionInfo {
	return v.value
}

func (v *NullableOnChainTransactionInfo) Set(val *OnChainTransactionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainTransactionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainTransactionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainTransactionInfo(val *OnChainTransactionInfo) *NullableOnChainTransactionInfo {
	return &NullableOnChainTransactionInfo{value: val, isSet: true}
}

func (v NullableOnChainTransactionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainTransactionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
