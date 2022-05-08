/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
	"fmt"
)

// AccountSignature - struct for AccountSignature
type AccountSignature struct {
	Ed25519Signature      *Ed25519Signature
	MultiEd25519Signature *MultiEd25519Signature
}

// Ed25519SignatureAsAccountSignature is a convenience function that returns Ed25519Signature wrapped in AccountSignature
func Ed25519SignatureAsAccountSignature(v *Ed25519Signature) AccountSignature {
	return AccountSignature{
		Ed25519Signature: v,
	}
}

// MultiEd25519SignatureAsAccountSignature is a convenience function that returns MultiEd25519Signature wrapped in AccountSignature
func MultiEd25519SignatureAsAccountSignature(v *MultiEd25519Signature) AccountSignature {
	return AccountSignature{
		MultiEd25519Signature: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccountSignature) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ed25519Signature
	err = newStrictDecoder(data).Decode(&dst.Ed25519Signature)
	if err == nil {
		jsonEd25519Signature, _ := json.Marshal(dst.Ed25519Signature)
		if string(jsonEd25519Signature) == "{}" { // empty struct
			dst.Ed25519Signature = nil
		} else {
			match++
		}
	} else {
		dst.Ed25519Signature = nil
	}

	// try to unmarshal data into MultiEd25519Signature
	err = newStrictDecoder(data).Decode(&dst.MultiEd25519Signature)
	if err == nil {
		jsonMultiEd25519Signature, _ := json.Marshal(dst.MultiEd25519Signature)
		if string(jsonMultiEd25519Signature) == "{}" { // empty struct
			dst.MultiEd25519Signature = nil
		} else {
			match++
		}
	} else {
		dst.MultiEd25519Signature = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ed25519Signature = nil
		dst.MultiEd25519Signature = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AccountSignature)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AccountSignature)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccountSignature) MarshalJSON() ([]byte, error) {
	if src.Ed25519Signature != nil {
		return json.Marshal(&src.Ed25519Signature)
	}

	if src.MultiEd25519Signature != nil {
		return json.Marshal(&src.MultiEd25519Signature)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccountSignature) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ed25519Signature != nil {
		return obj.Ed25519Signature
	}

	if obj.MultiEd25519Signature != nil {
		return obj.MultiEd25519Signature
	}

	// all schemas are nil
	return nil
}

type NullableAccountSignature struct {
	value *AccountSignature
	isSet bool
}

func (v NullableAccountSignature) Get() *AccountSignature {
	return v.value
}

func (v *NullableAccountSignature) Set(val *AccountSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSignature(val *AccountSignature) *NullableAccountSignature {
	return &NullableAccountSignature{value: val, isSet: true}
}

func (v NullableAccountSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
