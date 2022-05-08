/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// LedgerInfo struct for LedgerInfo
type LedgerInfo struct {
	// The blockchain chain id.
	ChainId int32 `json:"chain_id"`
	// The version of the latest transaction in the ledger.
	LedgerVersion string `json:"ledger_version"`
	// Timestamp in microseconds, e.g. ledger / block creation timestamp.
	LedgerTimestamp string `json:"ledger_timestamp"`
}

// NewLedgerInfo instantiates a new LedgerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerInfo(chainId int32, ledgerVersion string, ledgerTimestamp string) *LedgerInfo {
	this := LedgerInfo{}
	this.ChainId = chainId
	this.LedgerVersion = ledgerVersion
	this.LedgerTimestamp = ledgerTimestamp
	return &this
}

// NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerInfoWithDefaults() *LedgerInfo {
	this := LedgerInfo{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *LedgerInfo) GetChainId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetChainIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *LedgerInfo) SetChainId(v int32) {
	o.ChainId = v
}

// GetLedgerVersion returns the LedgerVersion field value
func (o *LedgerInfo) GetLedgerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerVersion
}

// GetLedgerVersionOk returns a tuple with the LedgerVersion field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetLedgerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerVersion, true
}

// SetLedgerVersion sets field value
func (o *LedgerInfo) SetLedgerVersion(v string) {
	o.LedgerVersion = v
}

// GetLedgerTimestamp returns the LedgerTimestamp field value
func (o *LedgerInfo) GetLedgerTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerTimestamp
}

// GetLedgerTimestampOk returns a tuple with the LedgerTimestamp field value
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetLedgerTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerTimestamp, true
}

// SetLedgerTimestamp sets field value
func (o *LedgerInfo) SetLedgerTimestamp(v string) {
	o.LedgerTimestamp = v
}

func (o LedgerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain_id"] = o.ChainId
	}
	if true {
		toSerialize["ledger_version"] = o.LedgerVersion
	}
	if true {
		toSerialize["ledger_timestamp"] = o.LedgerTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableLedgerInfo struct {
	value *LedgerInfo
	isSet bool
}

func (v NullableLedgerInfo) Get() *LedgerInfo {
	return v.value
}

func (v *NullableLedgerInfo) Set(val *LedgerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerInfo(val *LedgerInfo) *NullableLedgerInfo {
	return &NullableLedgerInfo{value: val, isSet: true}
}

func (v NullableLedgerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
