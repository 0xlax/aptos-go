/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// UserTransaction struct for UserTransaction
type UserTransaction struct {
	Type   string  `json:"type"`
	Events []Event `json:"events"`
	// Timestamp in microseconds, e.g. ledger / block creation timestamp.
	Timestamp string `json:"timestamp"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Sender string `json:"sender"`
	// Unsigned int64 type value
	SequenceNumber string `json:"sequence_number"`
	// Unsigned int64 type value
	MaxGasAmount string `json:"max_gas_amount"`
	// Unsigned int64 type value
	GasUnitPrice    string `json:"gas_unit_price"`
	GasCurrencyCode string `json:"gas_currency_code"`
	// Timestamp in seconds, e.g. transaction expiration timestamp.
	ExpirationTimestampSecs string               `json:"expiration_timestamp_secs"`
	Payload                 TransactionPayload   `json:"payload"`
	Signature               TransactionSignature `json:"signature"`
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

// NewUserTransaction instantiates a new UserTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransaction(type_ string, events []Event, timestamp string, sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange) *UserTransaction {
	this := UserTransaction{}
	this.Type = type_
	this.Events = events
	this.Timestamp = timestamp
	this.Sender = sender
	this.SequenceNumber = sequenceNumber
	this.MaxGasAmount = maxGasAmount
	this.GasUnitPrice = gasUnitPrice
	this.GasCurrencyCode = gasCurrencyCode
	this.ExpirationTimestampSecs = expirationTimestampSecs
	this.Payload = payload
	this.Signature = signature
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

// NewUserTransactionWithDefaults instantiates a new UserTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransactionWithDefaults() *UserTransaction {
	this := UserTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *UserTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserTransaction) SetType(v string) {
	o.Type = v
}

// GetEvents returns the Events field value
func (o *UserTransaction) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *UserTransaction) SetEvents(v []Event) {
	o.Events = v
}

// GetTimestamp returns the Timestamp field value
func (o *UserTransaction) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UserTransaction) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetSender returns the Sender field value
func (o *UserTransaction) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *UserTransaction) SetSender(v string) {
	o.Sender = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *UserTransaction) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *UserTransaction) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetMaxGasAmount returns the MaxGasAmount field value
func (o *UserTransaction) GetMaxGasAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxGasAmount
}

// GetMaxGasAmountOk returns a tuple with the MaxGasAmount field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetMaxGasAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGasAmount, true
}

// SetMaxGasAmount sets field value
func (o *UserTransaction) SetMaxGasAmount(v string) {
	o.MaxGasAmount = v
}

// GetGasUnitPrice returns the GasUnitPrice field value
func (o *UserTransaction) GetGasUnitPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUnitPrice
}

// GetGasUnitPriceOk returns a tuple with the GasUnitPrice field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetGasUnitPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUnitPrice, true
}

// SetGasUnitPrice sets field value
func (o *UserTransaction) SetGasUnitPrice(v string) {
	o.GasUnitPrice = v
}

// GetGasCurrencyCode returns the GasCurrencyCode field value
func (o *UserTransaction) GetGasCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasCurrencyCode
}

// GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetGasCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasCurrencyCode, true
}

// SetGasCurrencyCode sets field value
func (o *UserTransaction) SetGasCurrencyCode(v string) {
	o.GasCurrencyCode = v
}

// GetExpirationTimestampSecs returns the ExpirationTimestampSecs field value
func (o *UserTransaction) GetExpirationTimestampSecs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationTimestampSecs
}

// GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetExpirationTimestampSecsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestampSecs, true
}

// SetExpirationTimestampSecs sets field value
func (o *UserTransaction) SetExpirationTimestampSecs(v string) {
	o.ExpirationTimestampSecs = v
}

// GetPayload returns the Payload field value
func (o *UserTransaction) GetPayload() TransactionPayload {
	if o == nil {
		var ret TransactionPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetPayloadOk() (*TransactionPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *UserTransaction) SetPayload(v TransactionPayload) {
	o.Payload = v
}

// GetSignature returns the Signature field value
func (o *UserTransaction) GetSignature() TransactionSignature {
	if o == nil {
		var ret TransactionSignature
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetSignatureOk() (*TransactionSignature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *UserTransaction) SetSignature(v TransactionSignature) {
	o.Signature = v
}

// GetVersion returns the Version field value
func (o *UserTransaction) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UserTransaction) SetVersion(v string) {
	o.Version = v
}

// GetHash returns the Hash field value
func (o *UserTransaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *UserTransaction) SetHash(v string) {
	o.Hash = v
}

// GetStateRootHash returns the StateRootHash field value
func (o *UserTransaction) GetStateRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRootHash
}

// GetStateRootHashOk returns a tuple with the StateRootHash field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetStateRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRootHash, true
}

// SetStateRootHash sets field value
func (o *UserTransaction) SetStateRootHash(v string) {
	o.StateRootHash = v
}

// GetEventRootHash returns the EventRootHash field value
func (o *UserTransaction) GetEventRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventRootHash
}

// GetEventRootHashOk returns a tuple with the EventRootHash field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetEventRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventRootHash, true
}

// SetEventRootHash sets field value
func (o *UserTransaction) SetEventRootHash(v string) {
	o.EventRootHash = v
}

// GetGasUsed returns the GasUsed field value
func (o *UserTransaction) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *UserTransaction) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetSuccess returns the Success field value
func (o *UserTransaction) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *UserTransaction) SetSuccess(v bool) {
	o.Success = v
}

// GetVmStatus returns the VmStatus field value
func (o *UserTransaction) GetVmStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmStatus
}

// GetVmStatusOk returns a tuple with the VmStatus field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetVmStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmStatus, true
}

// SetVmStatus sets field value
func (o *UserTransaction) SetVmStatus(v string) {
	o.VmStatus = v
}

// GetAccumulatorRootHash returns the AccumulatorRootHash field value
func (o *UserTransaction) GetAccumulatorRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccumulatorRootHash
}

// GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetAccumulatorRootHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccumulatorRootHash, true
}

// SetAccumulatorRootHash sets field value
func (o *UserTransaction) SetAccumulatorRootHash(v string) {
	o.AccumulatorRootHash = v
}

// GetChanges returns the Changes field value
func (o *UserTransaction) GetChanges() []WriteSetChange {
	if o == nil {
		var ret []WriteSetChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *UserTransaction) GetChangesOk() ([]WriteSetChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *UserTransaction) SetChanges(v []WriteSetChange) {
	o.Changes = v
}

func (o UserTransaction) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if true {
		toSerialize["max_gas_amount"] = o.MaxGasAmount
	}
	if true {
		toSerialize["gas_unit_price"] = o.GasUnitPrice
	}
	if true {
		toSerialize["gas_currency_code"] = o.GasCurrencyCode
	}
	if true {
		toSerialize["expiration_timestamp_secs"] = o.ExpirationTimestampSecs
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if true {
		toSerialize["signature"] = o.Signature
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

type NullableUserTransaction struct {
	value *UserTransaction
	isSet bool
}

func (v NullableUserTransaction) Get() *UserTransaction {
	return v.value
}

func (v *NullableUserTransaction) Set(val *UserTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransaction(val *UserTransaction) *NullableUserTransaction {
	return &NullableUserTransaction{value: val, isSet: true}
}

func (v NullableUserTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
