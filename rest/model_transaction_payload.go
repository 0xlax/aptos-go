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

// TransactionPayload - struct for TransactionPayload
type TransactionPayload struct {
	ModuleBundlePayload   *ModuleBundlePayload
	ScriptFunctionPayload *ScriptFunctionPayload
	ScriptPayload         *ScriptPayload
	WriteSetPayload       *WriteSetPayload
}

// ModuleBundlePayloadAsTransactionPayload is a convenience function that returns ModuleBundlePayload wrapped in TransactionPayload
func ModuleBundlePayloadAsTransactionPayload(v *ModuleBundlePayload) TransactionPayload {
	return TransactionPayload{
		ModuleBundlePayload: v,
	}
}

// ScriptFunctionPayloadAsTransactionPayload is a convenience function that returns ScriptFunctionPayload wrapped in TransactionPayload
func ScriptFunctionPayloadAsTransactionPayload(v *ScriptFunctionPayload) TransactionPayload {
	return TransactionPayload{
		ScriptFunctionPayload: v,
	}
}

// ScriptPayloadAsTransactionPayload is a convenience function that returns ScriptPayload wrapped in TransactionPayload
func ScriptPayloadAsTransactionPayload(v *ScriptPayload) TransactionPayload {
	return TransactionPayload{
		ScriptPayload: v,
	}
}

// WriteSetPayloadAsTransactionPayload is a convenience function that returns WriteSetPayload wrapped in TransactionPayload
func WriteSetPayloadAsTransactionPayload(v *WriteSetPayload) TransactionPayload {
	return TransactionPayload{
		WriteSetPayload: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionPayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ModuleBundlePayload
	err = newStrictDecoder(data).Decode(&dst.ModuleBundlePayload)
	if err == nil {
		jsonModuleBundlePayload, _ := json.Marshal(dst.ModuleBundlePayload)
		if string(jsonModuleBundlePayload) == "{}" { // empty struct
			dst.ModuleBundlePayload = nil
		} else {
			match++
		}
	} else {
		dst.ModuleBundlePayload = nil
	}

	// try to unmarshal data into ScriptFunctionPayload
	err = newStrictDecoder(data).Decode(&dst.ScriptFunctionPayload)
	if err == nil {
		jsonScriptFunctionPayload, _ := json.Marshal(dst.ScriptFunctionPayload)
		if string(jsonScriptFunctionPayload) == "{}" { // empty struct
			dst.ScriptFunctionPayload = nil
		} else {
			match++
		}
	} else {
		dst.ScriptFunctionPayload = nil
	}

	// try to unmarshal data into ScriptPayload
	err = newStrictDecoder(data).Decode(&dst.ScriptPayload)
	if err == nil {
		jsonScriptPayload, _ := json.Marshal(dst.ScriptPayload)
		if string(jsonScriptPayload) == "{}" { // empty struct
			dst.ScriptPayload = nil
		} else {
			match++
		}
	} else {
		dst.ScriptPayload = nil
	}

	// try to unmarshal data into WriteSetPayload
	err = newStrictDecoder(data).Decode(&dst.WriteSetPayload)
	if err == nil {
		jsonWriteSetPayload, _ := json.Marshal(dst.WriteSetPayload)
		if string(jsonWriteSetPayload) == "{}" { // empty struct
			dst.WriteSetPayload = nil
		} else {
			match++
		}
	} else {
		dst.WriteSetPayload = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ModuleBundlePayload = nil
		dst.ScriptFunctionPayload = nil
		dst.ScriptPayload = nil
		dst.WriteSetPayload = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TransactionPayload)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TransactionPayload)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionPayload) MarshalJSON() ([]byte, error) {
	if src.ModuleBundlePayload != nil {
		return json.Marshal(&src.ModuleBundlePayload)
	}

	if src.ScriptFunctionPayload != nil {
		return json.Marshal(&src.ScriptFunctionPayload)
	}

	if src.ScriptPayload != nil {
		return json.Marshal(&src.ScriptPayload)
	}

	if src.WriteSetPayload != nil {
		return json.Marshal(&src.WriteSetPayload)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionPayload) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ModuleBundlePayload != nil {
		return obj.ModuleBundlePayload
	}

	if obj.ScriptFunctionPayload != nil {
		return obj.ScriptFunctionPayload
	}

	if obj.ScriptPayload != nil {
		return obj.ScriptPayload
	}

	if obj.WriteSetPayload != nil {
		return obj.WriteSetPayload
	}

	// all schemas are nil
	return nil
}

type NullableTransactionPayload struct {
	value *TransactionPayload
	isSet bool
}

func (v NullableTransactionPayload) Get() *TransactionPayload {
	return v.value
}

func (v *NullableTransactionPayload) Set(val *TransactionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPayload(val *TransactionPayload) *NullableTransactionPayload {
	return &NullableTransactionPayload{value: val, isSet: true}
}

func (v NullableTransactionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
