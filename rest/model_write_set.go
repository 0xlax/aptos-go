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

// WriteSet - struct for WriteSet
type WriteSet struct {
	DirectWriteSet *DirectWriteSet
	ScriptWriteSet *ScriptWriteSet
}

// DirectWriteSetAsWriteSet is a convenience function that returns DirectWriteSet wrapped in WriteSet
func DirectWriteSetAsWriteSet(v *DirectWriteSet) WriteSet {
	return WriteSet{
		DirectWriteSet: v,
	}
}

// ScriptWriteSetAsWriteSet is a convenience function that returns ScriptWriteSet wrapped in WriteSet
func ScriptWriteSetAsWriteSet(v *ScriptWriteSet) WriteSet {
	return WriteSet{
		ScriptWriteSet: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WriteSet) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DirectWriteSet
	err = newStrictDecoder(data).Decode(&dst.DirectWriteSet)
	if err == nil {
		jsonDirectWriteSet, _ := json.Marshal(dst.DirectWriteSet)
		if string(jsonDirectWriteSet) == "{}" { // empty struct
			dst.DirectWriteSet = nil
		} else {
			match++
		}
	} else {
		dst.DirectWriteSet = nil
	}

	// try to unmarshal data into ScriptWriteSet
	err = newStrictDecoder(data).Decode(&dst.ScriptWriteSet)
	if err == nil {
		jsonScriptWriteSet, _ := json.Marshal(dst.ScriptWriteSet)
		if string(jsonScriptWriteSet) == "{}" { // empty struct
			dst.ScriptWriteSet = nil
		} else {
			match++
		}
	} else {
		dst.ScriptWriteSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DirectWriteSet = nil
		dst.ScriptWriteSet = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WriteSet)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WriteSet)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WriteSet) MarshalJSON() ([]byte, error) {
	if src.DirectWriteSet != nil {
		return json.Marshal(&src.DirectWriteSet)
	}

	if src.ScriptWriteSet != nil {
		return json.Marshal(&src.ScriptWriteSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WriteSet) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DirectWriteSet != nil {
		return obj.DirectWriteSet
	}

	if obj.ScriptWriteSet != nil {
		return obj.ScriptWriteSet
	}

	// all schemas are nil
	return nil
}

type NullableWriteSet struct {
	value *WriteSet
	isSet bool
}

func (v NullableWriteSet) Get() *WriteSet {
	return v.value
}

func (v *NullableWriteSet) Set(val *WriteSet) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteSet) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteSet(val *WriteSet) *NullableWriteSet {
	return &NullableWriteSet{value: val, isSet: true}
}

func (v NullableWriteSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
