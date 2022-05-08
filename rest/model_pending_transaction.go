/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// PendingTransaction struct for PendingTransaction
type PendingTransaction struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Hash string `json:"hash"`
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
}

// NewPendingTransaction instantiates a new PendingTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransaction(type_ string, hash string, sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature) *PendingTransaction {
	this := PendingTransaction{}
	this.Type = type_
	this.Hash = hash
	this.Sender = sender
	this.SequenceNumber = sequenceNumber
	this.MaxGasAmount = maxGasAmount
	this.GasUnitPrice = gasUnitPrice
	this.GasCurrencyCode = gasCurrencyCode
	this.ExpirationTimestampSecs = expirationTimestampSecs
	this.Payload = payload
	this.Signature = signature
	return &this
}

// NewPendingTransactionWithDefaults instantiates a new PendingTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionWithDefaults() *PendingTransaction {
	this := PendingTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *PendingTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PendingTransaction) SetType(v string) {
	o.Type = v
}

// GetHash returns the Hash field value
func (o *PendingTransaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *PendingTransaction) SetHash(v string) {
	o.Hash = v
}

// GetSender returns the Sender field value
func (o *PendingTransaction) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *PendingTransaction) SetSender(v string) {
	o.Sender = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *PendingTransaction) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *PendingTransaction) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetMaxGasAmount returns the MaxGasAmount field value
func (o *PendingTransaction) GetMaxGasAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxGasAmount
}

// GetMaxGasAmountOk returns a tuple with the MaxGasAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetMaxGasAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGasAmount, true
}

// SetMaxGasAmount sets field value
func (o *PendingTransaction) SetMaxGasAmount(v string) {
	o.MaxGasAmount = v
}

// GetGasUnitPrice returns the GasUnitPrice field value
func (o *PendingTransaction) GetGasUnitPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUnitPrice
}

// GetGasUnitPriceOk returns a tuple with the GasUnitPrice field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetGasUnitPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUnitPrice, true
}

// SetGasUnitPrice sets field value
func (o *PendingTransaction) SetGasUnitPrice(v string) {
	o.GasUnitPrice = v
}

// GetGasCurrencyCode returns the GasCurrencyCode field value
func (o *PendingTransaction) GetGasCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasCurrencyCode
}

// GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetGasCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasCurrencyCode, true
}

// SetGasCurrencyCode sets field value
func (o *PendingTransaction) SetGasCurrencyCode(v string) {
	o.GasCurrencyCode = v
}

// GetExpirationTimestampSecs returns the ExpirationTimestampSecs field value
func (o *PendingTransaction) GetExpirationTimestampSecs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationTimestampSecs
}

// GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetExpirationTimestampSecsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestampSecs, true
}

// SetExpirationTimestampSecs sets field value
func (o *PendingTransaction) SetExpirationTimestampSecs(v string) {
	o.ExpirationTimestampSecs = v
}

// GetPayload returns the Payload field value
func (o *PendingTransaction) GetPayload() TransactionPayload {
	if o == nil {
		var ret TransactionPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetPayloadOk() (*TransactionPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *PendingTransaction) SetPayload(v TransactionPayload) {
	o.Payload = v
}

// GetSignature returns the Signature field value
func (o *PendingTransaction) GetSignature() TransactionSignature {
	if o == nil {
		var ret TransactionSignature
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetSignatureOk() (*TransactionSignature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *PendingTransaction) SetSignature(v TransactionSignature) {
	o.Signature = v
}

func (o PendingTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["hash"] = o.Hash
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
	return json.Marshal(toSerialize)
}

type NullablePendingTransaction struct {
	value *PendingTransaction
	isSet bool
}

func (v NullablePendingTransaction) Get() *PendingTransaction {
	return v.value
}

func (v *NullablePendingTransaction) Set(val *PendingTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransaction(val *PendingTransaction) *NullablePendingTransaction {
	return &NullablePendingTransaction{value: val, isSet: true}
}

func (v NullablePendingTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
