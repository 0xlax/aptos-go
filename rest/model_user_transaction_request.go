/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// UserTransactionRequest struct for UserTransactionRequest
type UserTransactionRequest struct {
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
	ExpirationTimestampSecs string             `json:"expiration_timestamp_secs"`
	Payload                 TransactionPayload `json:"payload"`
}

// NewUserTransactionRequest instantiates a new UserTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransactionRequest(sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload) *UserTransactionRequest {
	this := UserTransactionRequest{}
	this.Sender = sender
	this.SequenceNumber = sequenceNumber
	this.MaxGasAmount = maxGasAmount
	this.GasUnitPrice = gasUnitPrice
	this.GasCurrencyCode = gasCurrencyCode
	this.ExpirationTimestampSecs = expirationTimestampSecs
	this.Payload = payload
	return &this
}

// NewUserTransactionRequestWithDefaults instantiates a new UserTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransactionRequestWithDefaults() *UserTransactionRequest {
	this := UserTransactionRequest{}
	return &this
}

// GetSender returns the Sender field value
func (o *UserTransactionRequest) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *UserTransactionRequest) SetSender(v string) {
	o.Sender = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *UserTransactionRequest) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *UserTransactionRequest) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetMaxGasAmount returns the MaxGasAmount field value
func (o *UserTransactionRequest) GetMaxGasAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxGasAmount
}

// GetMaxGasAmountOk returns a tuple with the MaxGasAmount field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetMaxGasAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGasAmount, true
}

// SetMaxGasAmount sets field value
func (o *UserTransactionRequest) SetMaxGasAmount(v string) {
	o.MaxGasAmount = v
}

// GetGasUnitPrice returns the GasUnitPrice field value
func (o *UserTransactionRequest) GetGasUnitPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUnitPrice
}

// GetGasUnitPriceOk returns a tuple with the GasUnitPrice field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetGasUnitPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUnitPrice, true
}

// SetGasUnitPrice sets field value
func (o *UserTransactionRequest) SetGasUnitPrice(v string) {
	o.GasUnitPrice = v
}

// GetGasCurrencyCode returns the GasCurrencyCode field value
func (o *UserTransactionRequest) GetGasCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasCurrencyCode
}

// GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetGasCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasCurrencyCode, true
}

// SetGasCurrencyCode sets field value
func (o *UserTransactionRequest) SetGasCurrencyCode(v string) {
	o.GasCurrencyCode = v
}

// GetExpirationTimestampSecs returns the ExpirationTimestampSecs field value
func (o *UserTransactionRequest) GetExpirationTimestampSecs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationTimestampSecs
}

// GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetExpirationTimestampSecsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestampSecs, true
}

// SetExpirationTimestampSecs sets field value
func (o *UserTransactionRequest) SetExpirationTimestampSecs(v string) {
	o.ExpirationTimestampSecs = v
}

// GetPayload returns the Payload field value
func (o *UserTransactionRequest) GetPayload() TransactionPayload {
	if o == nil {
		var ret TransactionPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *UserTransactionRequest) GetPayloadOk() (*TransactionPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *UserTransactionRequest) SetPayload(v TransactionPayload) {
	o.Payload = v
}

func (o UserTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableUserTransactionRequest struct {
	value *UserTransactionRequest
	isSet bool
}

func (v NullableUserTransactionRequest) Get() *UserTransactionRequest {
	return v.value
}

func (v *NullableUserTransactionRequest) Set(val *UserTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransactionRequest(val *UserTransactionRequest) *NullableUserTransactionRequest {
	return &NullableUserTransactionRequest{value: val, isSet: true}
}

func (v NullableUserTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
