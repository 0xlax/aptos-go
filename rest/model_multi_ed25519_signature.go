/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MultiEd25519Signature Multi ed25519 signature, please refer to https://github.com/aptos-labs/aptos-core/tree/main/specifications/crypto#multi-signatures for more details.
type MultiEd25519Signature struct {
	Type string `json:"type"`
	// all public keys of the sender account
	PublicKeys []string `json:"public_keys"`
	// signatures created based on the `threshold`
	Signatures []string `json:"signatures"`
	// The threshold of the multi ed25519 account key.
	Threshold int32 `json:"threshold"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Bitmap string `json:"bitmap"`
}

// NewMultiEd25519Signature instantiates a new MultiEd25519Signature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEd25519Signature(type_ string, publicKeys []string, signatures []string, threshold int32, bitmap string) *MultiEd25519Signature {
	this := MultiEd25519Signature{}
	this.Type = type_
	this.PublicKeys = publicKeys
	this.Signatures = signatures
	this.Threshold = threshold
	this.Bitmap = bitmap
	return &this
}

// NewMultiEd25519SignatureWithDefaults instantiates a new MultiEd25519Signature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEd25519SignatureWithDefaults() *MultiEd25519Signature {
	this := MultiEd25519Signature{}
	return &this
}

// GetType returns the Type field value
func (o *MultiEd25519Signature) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MultiEd25519Signature) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MultiEd25519Signature) SetType(v string) {
	o.Type = v
}

// GetPublicKeys returns the PublicKeys field value
func (o *MultiEd25519Signature) GetPublicKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value
// and a boolean to check if the value has been set.
func (o *MultiEd25519Signature) GetPublicKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKeys, true
}

// SetPublicKeys sets field value
func (o *MultiEd25519Signature) SetPublicKeys(v []string) {
	o.PublicKeys = v
}

// GetSignatures returns the Signatures field value
func (o *MultiEd25519Signature) GetSignatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value
// and a boolean to check if the value has been set.
func (o *MultiEd25519Signature) GetSignaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures, true
}

// SetSignatures sets field value
func (o *MultiEd25519Signature) SetSignatures(v []string) {
	o.Signatures = v
}

// GetThreshold returns the Threshold field value
func (o *MultiEd25519Signature) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *MultiEd25519Signature) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *MultiEd25519Signature) SetThreshold(v int32) {
	o.Threshold = v
}

// GetBitmap returns the Bitmap field value
func (o *MultiEd25519Signature) GetBitmap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bitmap
}

// GetBitmapOk returns a tuple with the Bitmap field value
// and a boolean to check if the value has been set.
func (o *MultiEd25519Signature) GetBitmapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bitmap, true
}

// SetBitmap sets field value
func (o *MultiEd25519Signature) SetBitmap(v string) {
	o.Bitmap = v
}

func (o MultiEd25519Signature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["public_keys"] = o.PublicKeys
	}
	if true {
		toSerialize["signatures"] = o.Signatures
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["bitmap"] = o.Bitmap
	}
	return json.Marshal(toSerialize)
}

type NullableMultiEd25519Signature struct {
	value *MultiEd25519Signature
	isSet bool
}

func (v NullableMultiEd25519Signature) Get() *MultiEd25519Signature {
	return v.value
}

func (v *NullableMultiEd25519Signature) Set(val *MultiEd25519Signature) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEd25519Signature) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEd25519Signature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEd25519Signature(val *MultiEd25519Signature) *NullableMultiEd25519Signature {
	return &NullableMultiEd25519Signature{value: val, isSet: true}
}

func (v NullableMultiEd25519Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEd25519Signature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
