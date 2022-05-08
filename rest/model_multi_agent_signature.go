/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MultiAgentSignature Multi agent signature, please refer to TBD.
type MultiAgentSignature struct {
	Type                     string             `json:"type"`
	Sender                   AccountSignature   `json:"sender"`
	SecondarySignerAddresses []string           `json:"secondary_signer_addresses"`
	SecondarySigners         []AccountSignature `json:"secondary_signers"`
}

// NewMultiAgentSignature instantiates a new MultiAgentSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAgentSignature(type_ string, sender AccountSignature, secondarySignerAddresses []string, secondarySigners []AccountSignature) *MultiAgentSignature {
	this := MultiAgentSignature{}
	this.Type = type_
	this.Sender = sender
	this.SecondarySignerAddresses = secondarySignerAddresses
	this.SecondarySigners = secondarySigners
	return &this
}

// NewMultiAgentSignatureWithDefaults instantiates a new MultiAgentSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAgentSignatureWithDefaults() *MultiAgentSignature {
	this := MultiAgentSignature{}
	return &this
}

// GetType returns the Type field value
func (o *MultiAgentSignature) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MultiAgentSignature) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MultiAgentSignature) SetType(v string) {
	o.Type = v
}

// GetSender returns the Sender field value
func (o *MultiAgentSignature) GetSender() AccountSignature {
	if o == nil {
		var ret AccountSignature
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *MultiAgentSignature) GetSenderOk() (*AccountSignature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *MultiAgentSignature) SetSender(v AccountSignature) {
	o.Sender = v
}

// GetSecondarySignerAddresses returns the SecondarySignerAddresses field value
func (o *MultiAgentSignature) GetSecondarySignerAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecondarySignerAddresses
}

// GetSecondarySignerAddressesOk returns a tuple with the SecondarySignerAddresses field value
// and a boolean to check if the value has been set.
func (o *MultiAgentSignature) GetSecondarySignerAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondarySignerAddresses, true
}

// SetSecondarySignerAddresses sets field value
func (o *MultiAgentSignature) SetSecondarySignerAddresses(v []string) {
	o.SecondarySignerAddresses = v
}

// GetSecondarySigners returns the SecondarySigners field value
func (o *MultiAgentSignature) GetSecondarySigners() []AccountSignature {
	if o == nil {
		var ret []AccountSignature
		return ret
	}

	return o.SecondarySigners
}

// GetSecondarySignersOk returns a tuple with the SecondarySigners field value
// and a boolean to check if the value has been set.
func (o *MultiAgentSignature) GetSecondarySignersOk() ([]AccountSignature, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondarySigners, true
}

// SetSecondarySigners sets field value
func (o *MultiAgentSignature) SetSecondarySigners(v []AccountSignature) {
	o.SecondarySigners = v
}

func (o MultiAgentSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["secondary_signer_addresses"] = o.SecondarySignerAddresses
	}
	if true {
		toSerialize["secondary_signers"] = o.SecondarySigners
	}
	return json.Marshal(toSerialize)
}

type NullableMultiAgentSignature struct {
	value *MultiAgentSignature
	isSet bool
}

func (v NullableMultiAgentSignature) Get() *MultiAgentSignature {
	return v.value
}

func (v *NullableMultiAgentSignature) Set(val *MultiAgentSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAgentSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAgentSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiAgentSignature(val *MultiAgentSignature) *NullableMultiAgentSignature {
	return &NullableMultiAgentSignature{value: val, isSet: true}
}

func (v NullableMultiAgentSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAgentSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
