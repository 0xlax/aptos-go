/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// BlockMetadataTransaction struct for BlockMetadataTransaction
type BlockMetadataTransaction struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Id string `json:"id"`
	// Unsigned int64 type value
	Round              string   `json:"round"`
	PreviousBlockVotes []string `json:"previous_block_votes"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Proposer string `json:"proposer"`
	// Timestamp in microseconds, e.g. ledger / block creation timestamp.
	Timestamp string `json:"timestamp"`
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

// NewBlockMetadataTransaction instantiates a new BlockMetadataTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockMetadataTransaction(type_ string, id string, round string, previousBlockVotes []string, proposer string, timestamp string, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange) *BlockMetadataTransaction {
	this := BlockMetadataTransaction{}
	this.Type = type_
	this.Id = id
	this.Round = round
	this.PreviousBlockVotes = previousBlockVotes
	this.Proposer = proposer
	this.Timestamp = timestamp
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

// NewBlockMetadataTransactionWithDefaults instantiates a new BlockMetadataTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockMetadataTransactionWithDefaults() *BlockMetadataTransaction {
	this := BlockMetadataTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *BlockMetadataTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlockMetadataTransaction) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BlockMetadataTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlockMetadataTransaction) SetId(v string) {
	o.Id = v
}

// GetRound returns the Round field value
func (o *BlockMetadataTransaction) GetRound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetRoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *BlockMetadataTransaction) SetRound(v string) {
	o.Round = v
}

// GetPreviousBlockVotes returns the PreviousBlockVotes field value
func (o *BlockMetadataTransaction) GetPreviousBlockVotes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreviousBlockVotes
}

// GetPreviousBlockVotesOk returns a tuple with the PreviousBlockVotes field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetPreviousBlockVotesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousBlockVotes, true
}

// SetPreviousBlockVotes sets field value
func (o *BlockMetadataTransaction) SetPreviousBlockVotes(v []string) {
	o.PreviousBlockVotes = v
}

// GetProposer returns the Proposer field value
func (o *BlockMetadataTransaction) GetProposer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proposer
}

// GetProposerOk returns a tuple with the Proposer field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetProposerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proposer, true
}

// SetProposer sets field value
func (o *BlockMetadataTransaction) SetProposer(v string) {
	o.Proposer = v
}

// GetTimestamp returns the Timestamp field value
func (o *BlockMetadataTransaction) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BlockMetadataTransaction) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetVersion returns the Version field value
func (o *BlockMetadataTransaction) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BlockMetadataTransaction) SetVersion(v string) {
	o.Version = v
}

// GetHash returns the Hash field value
func (o *BlockMetadataTransaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockMetadataTransaction) SetHash(v string) {
	o.Hash = v
}

// GetStateRootHash returns the StateRootHash field value
func (o *BlockMetadataTransaction) GetStateRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRootHash
}

// GetStateRootHashOk returns a tuple with the StateRootHash field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetStateRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRootHash, true
}

// SetStateRootHash sets field value
func (o *BlockMetadataTransaction) SetStateRootHash(v string) {
	o.StateRootHash = v
}

// GetEventRootHash returns the EventRootHash field value
func (o *BlockMetadataTransaction) GetEventRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventRootHash
}

// GetEventRootHashOk returns a tuple with the EventRootHash field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetEventRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventRootHash, true
}

// SetEventRootHash sets field value
func (o *BlockMetadataTransaction) SetEventRootHash(v string) {
	o.EventRootHash = v
}

// GetGasUsed returns the GasUsed field value
func (o *BlockMetadataTransaction) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *BlockMetadataTransaction) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetSuccess returns the Success field value
func (o *BlockMetadataTransaction) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *BlockMetadataTransaction) SetSuccess(v bool) {
	o.Success = v
}

// GetVmStatus returns the VmStatus field value
func (o *BlockMetadataTransaction) GetVmStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmStatus
}

// GetVmStatusOk returns a tuple with the VmStatus field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetVmStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmStatus, true
}

// SetVmStatus sets field value
func (o *BlockMetadataTransaction) SetVmStatus(v string) {
	o.VmStatus = v
}

// GetAccumulatorRootHash returns the AccumulatorRootHash field value
func (o *BlockMetadataTransaction) GetAccumulatorRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccumulatorRootHash
}

// GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetAccumulatorRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccumulatorRootHash, true
}

// SetAccumulatorRootHash sets field value
func (o *BlockMetadataTransaction) SetAccumulatorRootHash(v string) {
	o.AccumulatorRootHash = v
}

// GetChanges returns the Changes field value
func (o *BlockMetadataTransaction) GetChanges() []WriteSetChange {
	if o == nil {
		var ret []WriteSetChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransaction) GetChangesOk() ([]WriteSetChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *BlockMetadataTransaction) SetChanges(v []WriteSetChange) {
	o.Changes = v
}

func (o BlockMetadataTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["round"] = o.Round
	}
	if true {
		toSerialize["previous_block_votes"] = o.PreviousBlockVotes
	}
	if true {
		toSerialize["proposer"] = o.Proposer
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
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

type NullableBlockMetadataTransaction struct {
	value *BlockMetadataTransaction
	isSet bool
}

func (v NullableBlockMetadataTransaction) Get() *BlockMetadataTransaction {
	return v.value
}

func (v *NullableBlockMetadataTransaction) Set(val *BlockMetadataTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockMetadataTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockMetadataTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockMetadataTransaction(val *BlockMetadataTransaction) *NullableBlockMetadataTransaction {
	return &NullableBlockMetadataTransaction{value: val, isSet: true}
}

func (v NullableBlockMetadataTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockMetadataTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
