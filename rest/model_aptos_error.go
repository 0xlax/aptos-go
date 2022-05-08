/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// AptosError struct for AptosError
type AptosError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	// The version of the latest transaction in the ledger.
	AptosLedgerVersion *string `json:"aptos_ledger_version,omitempty"`
}

// NewAptosError instantiates a new AptosError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptosError(code int32, message string) *AptosError {
	this := AptosError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAptosErrorWithDefaults instantiates a new AptosError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptosErrorWithDefaults() *AptosError {
	this := AptosError{}
	return &this
}

// GetCode returns the Code field value
func (o *AptosError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AptosError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AptosError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *AptosError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AptosError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AptosError) SetMessage(v string) {
	o.Message = v
}

// GetAptosLedgerVersion returns the AptosLedgerVersion field value if set, zero value otherwise.
func (o *AptosError) GetAptosLedgerVersion() string {
	if o == nil || o.AptosLedgerVersion == nil {
		var ret string
		return ret
	}
	return *o.AptosLedgerVersion
}

// GetAptosLedgerVersionOk returns a tuple with the AptosLedgerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptosError) GetAptosLedgerVersionOk() (*string, bool) {
	if o == nil || o.AptosLedgerVersion == nil {
		return nil, false
	}
	return o.AptosLedgerVersion, true
}

// HasAptosLedgerVersion returns a boolean if a field has been set.
func (o *AptosError) HasAptosLedgerVersion() bool {
	if o != nil && o.AptosLedgerVersion != nil {
		return true
	}

	return false
}

// SetAptosLedgerVersion gets a reference to the given string and assigns it to the AptosLedgerVersion field.
func (o *AptosError) SetAptosLedgerVersion(v string) {
	o.AptosLedgerVersion = &v
}

func (o AptosError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.AptosLedgerVersion != nil {
		toSerialize["aptos_ledger_version"] = o.AptosLedgerVersion
	}
	return json.Marshal(toSerialize)
}

type NullableAptosError struct {
	value *AptosError
	isSet bool
}

func (v NullableAptosError) Get() *AptosError {
	return v.value
}

func (v *NullableAptosError) Set(val *AptosError) {
	v.value = val
	v.isSet = true
}

func (v NullableAptosError) IsSet() bool {
	return v.isSet
}

func (v *NullableAptosError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptosError(val *AptosError) *NullableAptosError {
	return &NullableAptosError{value: val, isSet: true}
}

func (v NullableAptosError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptosError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
