/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// BlockMetadataTransactionAllOf struct for BlockMetadataTransactionAllOf
type BlockMetadataTransactionAllOf struct {
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
}

// NewBlockMetadataTransactionAllOf instantiates a new BlockMetadataTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockMetadataTransactionAllOf(type_ string, id string, round string, previousBlockVotes []string, proposer string, timestamp string) *BlockMetadataTransactionAllOf {
	this := BlockMetadataTransactionAllOf{}
	this.Type = type_
	this.Id = id
	this.Round = round
	this.PreviousBlockVotes = previousBlockVotes
	this.Proposer = proposer
	this.Timestamp = timestamp
	return &this
}

// NewBlockMetadataTransactionAllOfWithDefaults instantiates a new BlockMetadataTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockMetadataTransactionAllOfWithDefaults() *BlockMetadataTransactionAllOf {
	this := BlockMetadataTransactionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *BlockMetadataTransactionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlockMetadataTransactionAllOf) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BlockMetadataTransactionAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlockMetadataTransactionAllOf) SetId(v string) {
	o.Id = v
}

// GetRound returns the Round field value
func (o *BlockMetadataTransactionAllOf) GetRound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetRoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *BlockMetadataTransactionAllOf) SetRound(v string) {
	o.Round = v
}

// GetPreviousBlockVotes returns the PreviousBlockVotes field value
func (o *BlockMetadataTransactionAllOf) GetPreviousBlockVotes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreviousBlockVotes
}

// GetPreviousBlockVotesOk returns a tuple with the PreviousBlockVotes field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetPreviousBlockVotesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousBlockVotes, true
}

// SetPreviousBlockVotes sets field value
func (o *BlockMetadataTransactionAllOf) SetPreviousBlockVotes(v []string) {
	o.PreviousBlockVotes = v
}

// GetProposer returns the Proposer field value
func (o *BlockMetadataTransactionAllOf) GetProposer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proposer
}

// GetProposerOk returns a tuple with the Proposer field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetProposerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proposer, true
}

// SetProposer sets field value
func (o *BlockMetadataTransactionAllOf) SetProposer(v string) {
	o.Proposer = v
}

// GetTimestamp returns the Timestamp field value
func (o *BlockMetadataTransactionAllOf) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BlockMetadataTransactionAllOf) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BlockMetadataTransactionAllOf) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o BlockMetadataTransactionAllOf) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBlockMetadataTransactionAllOf struct {
	value *BlockMetadataTransactionAllOf
	isSet bool
}

func (v NullableBlockMetadataTransactionAllOf) Get() *BlockMetadataTransactionAllOf {
	return v.value
}

func (v *NullableBlockMetadataTransactionAllOf) Set(val *BlockMetadataTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockMetadataTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockMetadataTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockMetadataTransactionAllOf(val *BlockMetadataTransactionAllOf) *NullableBlockMetadataTransactionAllOf {
	return &NullableBlockMetadataTransactionAllOf{value: val, isSet: true}
}

func (v NullableBlockMetadataTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockMetadataTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
