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

// Transaction - struct for Transaction
type Transaction struct {
	BlockMetadataTransaction *BlockMetadataTransaction
	GenesisTransaction       *GenesisTransaction
	PendingTransaction       *PendingTransaction
	UserTransaction          *UserTransaction
}

// BlockMetadataTransactionAsTransaction is a convenience function that returns BlockMetadataTransaction wrapped in Transaction
func BlockMetadataTransactionAsTransaction(v *BlockMetadataTransaction) Transaction {
	return Transaction{
		BlockMetadataTransaction: v,
	}
}

// GenesisTransactionAsTransaction is a convenience function that returns GenesisTransaction wrapped in Transaction
func GenesisTransactionAsTransaction(v *GenesisTransaction) Transaction {
	return Transaction{
		GenesisTransaction: v,
	}
}

// PendingTransactionAsTransaction is a convenience function that returns PendingTransaction wrapped in Transaction
func PendingTransactionAsTransaction(v *PendingTransaction) Transaction {
	return Transaction{
		PendingTransaction: v,
	}
}

// UserTransactionAsTransaction is a convenience function that returns UserTransaction wrapped in Transaction
func UserTransactionAsTransaction(v *UserTransaction) Transaction {
	return Transaction{
		UserTransaction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Transaction) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into PendingTransaction
	err = newStrictDecoder(data).Decode(&dst.PendingTransaction)
	if err == nil {
		jsonPendingTransaction, _ := json.Marshal(dst.PendingTransaction)
		if string(jsonPendingTransaction) == "{}" { // empty struct
			dst.PendingTransaction = nil
		} else {
			match++
		}
	} else {
		dst.PendingTransaction = nil
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
		dst.PendingTransaction = nil
		dst.UserTransaction = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Transaction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Transaction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Transaction) MarshalJSON() ([]byte, error) {
	if src.BlockMetadataTransaction != nil {
		return json.Marshal(&src.BlockMetadataTransaction)
	}

	if src.GenesisTransaction != nil {
		return json.Marshal(&src.GenesisTransaction)
	}

	if src.PendingTransaction != nil {
		return json.Marshal(&src.PendingTransaction)
	}

	if src.UserTransaction != nil {
		return json.Marshal(&src.UserTransaction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Transaction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BlockMetadataTransaction != nil {
		return obj.BlockMetadataTransaction
	}

	if obj.GenesisTransaction != nil {
		return obj.GenesisTransaction
	}

	if obj.PendingTransaction != nil {
		return obj.PendingTransaction
	}

	if obj.UserTransaction != nil {
		return obj.UserTransaction
	}

	// all schemas are nil
	return nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
