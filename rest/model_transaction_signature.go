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

// TransactionSignature - struct for TransactionSignature
type TransactionSignature struct {
	Ed25519Signature      *Ed25519Signature
	MultiAgentSignature   *MultiAgentSignature
	MultiEd25519Signature *MultiEd25519Signature
}

// Ed25519SignatureAsTransactionSignature is a convenience function that returns Ed25519Signature wrapped in TransactionSignature
func Ed25519SignatureAsTransactionSignature(v *Ed25519Signature) TransactionSignature {
	return TransactionSignature{
		Ed25519Signature: v,
	}
}

// MultiAgentSignatureAsTransactionSignature is a convenience function that returns MultiAgentSignature wrapped in TransactionSignature
func MultiAgentSignatureAsTransactionSignature(v *MultiAgentSignature) TransactionSignature {
	return TransactionSignature{
		MultiAgentSignature: v,
	}
}

// MultiEd25519SignatureAsTransactionSignature is a convenience function that returns MultiEd25519Signature wrapped in TransactionSignature
func MultiEd25519SignatureAsTransactionSignature(v *MultiEd25519Signature) TransactionSignature {
	return TransactionSignature{
		MultiEd25519Signature: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionSignature) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into MultiAgentSignature
	err = newStrictDecoder(data).Decode(&dst.MultiAgentSignature)
	if err == nil {
		jsonMultiAgentSignature, _ := json.Marshal(dst.MultiAgentSignature)
		if string(jsonMultiAgentSignature) == "{}" { // empty struct
			dst.MultiAgentSignature = nil
		} else {
			match++
		}
	} else {
		dst.MultiAgentSignature = nil
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
		dst.MultiAgentSignature = nil
		dst.MultiEd25519Signature = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TransactionSignature)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TransactionSignature)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionSignature) MarshalJSON() ([]byte, error) {
	if src.Ed25519Signature != nil {
		return json.Marshal(&src.Ed25519Signature)
	}

	if src.MultiAgentSignature != nil {
		return json.Marshal(&src.MultiAgentSignature)
	}

	if src.MultiEd25519Signature != nil {
		return json.Marshal(&src.MultiEd25519Signature)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionSignature) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ed25519Signature != nil {
		return obj.Ed25519Signature
	}

	if obj.MultiAgentSignature != nil {
		return obj.MultiAgentSignature
	}

	if obj.MultiEd25519Signature != nil {
		return obj.MultiEd25519Signature
	}

	// all schemas are nil
	return nil
}

type NullableTransactionSignature struct {
	value *TransactionSignature
	isSet bool
}

func (v NullableTransactionSignature) Get() *TransactionSignature {
	return v.value
}

func (v *NullableTransactionSignature) Set(val *TransactionSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSignature(val *TransactionSignature) *NullableTransactionSignature {
	return &NullableTransactionSignature{value: val, isSet: true}
}

func (v NullableTransactionSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
