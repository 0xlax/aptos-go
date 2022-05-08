/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// Ed25519Signature Please refer to https://github.com/aptos-labs/aptos-core/tree/main/specifications/crypto#signature-and-verification for more details.
type Ed25519Signature struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	PublicKey string `json:"public_key"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Signature string `json:"signature"`
}

// NewEd25519Signature instantiates a new Ed25519Signature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEd25519Signature(type_ string, publicKey string, signature string) *Ed25519Signature {
	this := Ed25519Signature{}
	this.Type = type_
	this.PublicKey = publicKey
	this.Signature = signature
	return &this
}

// NewEd25519SignatureWithDefaults instantiates a new Ed25519Signature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEd25519SignatureWithDefaults() *Ed25519Signature {
	this := Ed25519Signature{}
	return &this
}

// GetType returns the Type field value
func (o *Ed25519Signature) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Ed25519Signature) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Ed25519Signature) SetType(v string) {
	o.Type = v
}

// GetPublicKey returns the PublicKey field value
func (o *Ed25519Signature) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *Ed25519Signature) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *Ed25519Signature) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetSignature returns the Signature field value
func (o *Ed25519Signature) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *Ed25519Signature) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *Ed25519Signature) SetSignature(v string) {
	o.Signature = v
}

func (o Ed25519Signature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableEd25519Signature struct {
	value *Ed25519Signature
	isSet bool
}

func (v NullableEd25519Signature) Get() *Ed25519Signature {
	return v.value
}

func (v *NullableEd25519Signature) Set(val *Ed25519Signature) {
	v.value = val
	v.isSet = true
}

func (v NullableEd25519Signature) IsSet() bool {
	return v.isSet
}

func (v *NullableEd25519Signature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEd25519Signature(val *Ed25519Signature) *NullableEd25519Signature {
	return &NullableEd25519Signature{value: val, isSet: true}
}

func (v NullableEd25519Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEd25519Signature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
