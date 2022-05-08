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

// WriteSetChange - struct for WriteSetChange
type WriteSetChange struct {
	DeleteModule    *DeleteModule
	DeleteResource  *DeleteResource
	DeleteTableItem *DeleteTableItem
	WriteModule     *WriteModule
	WriteResource   *WriteResource
	WriteTableItem  *WriteTableItem
}

// DeleteModuleAsWriteSetChange is a convenience function that returns DeleteModule wrapped in WriteSetChange
func DeleteModuleAsWriteSetChange(v *DeleteModule) WriteSetChange {
	return WriteSetChange{
		DeleteModule: v,
	}
}

// DeleteResourceAsWriteSetChange is a convenience function that returns DeleteResource wrapped in WriteSetChange
func DeleteResourceAsWriteSetChange(v *DeleteResource) WriteSetChange {
	return WriteSetChange{
		DeleteResource: v,
	}
}

// DeleteTableItemAsWriteSetChange is a convenience function that returns DeleteTableItem wrapped in WriteSetChange
func DeleteTableItemAsWriteSetChange(v *DeleteTableItem) WriteSetChange {
	return WriteSetChange{
		DeleteTableItem: v,
	}
}

// WriteModuleAsWriteSetChange is a convenience function that returns WriteModule wrapped in WriteSetChange
func WriteModuleAsWriteSetChange(v *WriteModule) WriteSetChange {
	return WriteSetChange{
		WriteModule: v,
	}
}

// WriteResourceAsWriteSetChange is a convenience function that returns WriteResource wrapped in WriteSetChange
func WriteResourceAsWriteSetChange(v *WriteResource) WriteSetChange {
	return WriteSetChange{
		WriteResource: v,
	}
}

// WriteTableItemAsWriteSetChange is a convenience function that returns WriteTableItem wrapped in WriteSetChange
func WriteTableItemAsWriteSetChange(v *WriteTableItem) WriteSetChange {
	return WriteSetChange{
		WriteTableItem: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WriteSetChange) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeleteModule
	err = newStrictDecoder(data).Decode(&dst.DeleteModule)
	if err == nil {
		jsonDeleteModule, _ := json.Marshal(dst.DeleteModule)
		if string(jsonDeleteModule) == "{}" { // empty struct
			dst.DeleteModule = nil
		} else {
			match++
		}
	} else {
		dst.DeleteModule = nil
	}

	// try to unmarshal data into DeleteResource
	err = newStrictDecoder(data).Decode(&dst.DeleteResource)
	if err == nil {
		jsonDeleteResource, _ := json.Marshal(dst.DeleteResource)
		if string(jsonDeleteResource) == "{}" { // empty struct
			dst.DeleteResource = nil
		} else {
			match++
		}
	} else {
		dst.DeleteResource = nil
	}

	// try to unmarshal data into DeleteTableItem
	err = newStrictDecoder(data).Decode(&dst.DeleteTableItem)
	if err == nil {
		jsonDeleteTableItem, _ := json.Marshal(dst.DeleteTableItem)
		if string(jsonDeleteTableItem) == "{}" { // empty struct
			dst.DeleteTableItem = nil
		} else {
			match++
		}
	} else {
		dst.DeleteTableItem = nil
	}

	// try to unmarshal data into WriteModule
	err = newStrictDecoder(data).Decode(&dst.WriteModule)
	if err == nil {
		jsonWriteModule, _ := json.Marshal(dst.WriteModule)
		if string(jsonWriteModule) == "{}" { // empty struct
			dst.WriteModule = nil
		} else {
			match++
		}
	} else {
		dst.WriteModule = nil
	}

	// try to unmarshal data into WriteResource
	err = newStrictDecoder(data).Decode(&dst.WriteResource)
	if err == nil {
		jsonWriteResource, _ := json.Marshal(dst.WriteResource)
		if string(jsonWriteResource) == "{}" { // empty struct
			dst.WriteResource = nil
		} else {
			match++
		}
	} else {
		dst.WriteResource = nil
	}

	// try to unmarshal data into WriteTableItem
	err = newStrictDecoder(data).Decode(&dst.WriteTableItem)
	if err == nil {
		jsonWriteTableItem, _ := json.Marshal(dst.WriteTableItem)
		if string(jsonWriteTableItem) == "{}" { // empty struct
			dst.WriteTableItem = nil
		} else {
			match++
		}
	} else {
		dst.WriteTableItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeleteModule = nil
		dst.DeleteResource = nil
		dst.DeleteTableItem = nil
		dst.WriteModule = nil
		dst.WriteResource = nil
		dst.WriteTableItem = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WriteSetChange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WriteSetChange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WriteSetChange) MarshalJSON() ([]byte, error) {
	if src.DeleteModule != nil {
		return json.Marshal(&src.DeleteModule)
	}

	if src.DeleteResource != nil {
		return json.Marshal(&src.DeleteResource)
	}

	if src.DeleteTableItem != nil {
		return json.Marshal(&src.DeleteTableItem)
	}

	if src.WriteModule != nil {
		return json.Marshal(&src.WriteModule)
	}

	if src.WriteResource != nil {
		return json.Marshal(&src.WriteResource)
	}

	if src.WriteTableItem != nil {
		return json.Marshal(&src.WriteTableItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WriteSetChange) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DeleteModule != nil {
		return obj.DeleteModule
	}

	if obj.DeleteResource != nil {
		return obj.DeleteResource
	}

	if obj.DeleteTableItem != nil {
		return obj.DeleteTableItem
	}

	if obj.WriteModule != nil {
		return obj.WriteModule
	}

	if obj.WriteResource != nil {
		return obj.WriteResource
	}

	if obj.WriteTableItem != nil {
		return obj.WriteTableItem
	}

	// all schemas are nil
	return nil
}

type NullableWriteSetChange struct {
	value *WriteSetChange
	isSet bool
}

func (v NullableWriteSetChange) Get() *WriteSetChange {
	return v.value
}

func (v *NullableWriteSetChange) Set(val *WriteSetChange) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteSetChange) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteSetChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteSetChange(val *WriteSetChange) *NullableWriteSetChange {
	return &NullableWriteSetChange{value: val, isSet: true}
}

func (v NullableWriteSetChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteSetChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
