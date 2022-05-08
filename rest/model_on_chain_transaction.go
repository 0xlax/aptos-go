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

// OnChainTransaction - struct for OnChainTransaction
type OnChainTransaction struct {
	BlockMetadataTransaction *BlockMetadataTransaction
	GenesisTransaction       *GenesisTransaction
	UserTransaction          *UserTransaction
}

// BlockMetadataTransactionAsOnChainTransaction is a convenience function that returns BlockMetadataTransaction wrapped in OnChainTransaction
func BlockMetadataTransactionAsOnChainTransaction(v *BlockMetadataTransaction) OnChainTransaction {
	return OnChainTransaction{
		BlockMetadataTransaction: v,
	}
}

// GenesisTransactionAsOnChainTransaction is a convenience function that returns GenesisTransaction wrapped in OnChainTransaction
func GenesisTransactionAsOnChainTransaction(v *GenesisTransaction) OnChainTransaction {
	return OnChainTransaction{
		GenesisTransaction: v,
	}
}

// UserTransactionAsOnChainTransaction is a convenience function that returns UserTransaction wrapped in OnChainTransaction
func UserTransactionAsOnChainTransaction(v *UserTransaction) OnChainTransaction {
	return OnChainTransaction{
		UserTransaction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OnChainTransaction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlockMetadataTransaction
	err = newStrictDecoder(data).Decode(&dst.BlockMetadataTransaction)
	if err == nil {
		jsonBlockMetadataTransaction, _ := json.Marshal(dst.BlockMetadataTransaction)
		if string(jsonBlockMetadataTransaction) == "{}" { // empty struct
			dst.BlockMetadataTransaction = nil
		} else {
			match++
		}
	} else {
		dst.BlockMetadataTransaction = nil
	}

	// try to unmarshal data into GenesisTransaction
	err = newStrictDecoder(data).Decode(&dst.GenesisTransaction)
	if err == nil {
		jsonGenesisTransaction, _ := json.Marshal(dst.GenesisTransaction)
		if string(jsonGenesisTransaction) == "{}" { // empty struct
			dst.GenesisTransaction = nil
		} else {
			match++
		}
	} else {
		dst.GenesisTransaction = nil
	}

	// try to unmarshal data into UserTransaction
	err = newStrictDecoder(data).Decode(&dst.UserTransaction)
	if err == nil {
		jsonUserTransaction, _ := json.Marshal(dst.UserTransaction)
		if string(jsonUserTransaction) == "{}" { // empty struct
			dst.UserTransaction = nil
		} else {
			match++
		}
	} else {
		dst.UserTransaction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlockMetadataTransaction = nil
		dst.GenesisTransaction = nil
		dst.UserTransaction = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(OnChainTransaction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(OnChainTransaction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OnChainTransaction) MarshalJSON() ([]byte, error) {
	if src.BlockMetadataTransaction != nil {
		return json.Marshal(&src.BlockMetadataTransaction)
	}

	if src.GenesisTransaction != nil {
		return json.Marshal(&src.GenesisTransaction)
	}

	if src.UserTransaction != nil {
		return json.Marshal(&src.UserTransaction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OnChainTransaction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BlockMetadataTransaction != nil {
		return obj.BlockMetadataTransaction
	}

	if obj.GenesisTransaction != nil {
		return obj.GenesisTransaction
	}

	if obj.UserTransaction != nil {
		return obj.UserTransaction
	}

	// all schemas are nil
	return nil
}

type NullableOnChainTransaction struct {
	value *OnChainTransaction
	isSet bool
}

func (v NullableOnChainTransaction) Get() *OnChainTransaction {
	return v.value
}

func (v *NullableOnChainTransaction) Set(val *OnChainTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainTransaction(val *OnChainTransaction) *NullableOnChainTransaction {
	return &NullableOnChainTransaction{value: val, isSet: true}
}

func (v NullableOnChainTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
