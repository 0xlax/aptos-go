/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveStructField struct for MoveStructField
type MoveStructField struct {
	Name string `json:"name"`
	// String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: `vector<{non-reference MoveTypeId}>`   - struct: `{address}::{module_name}::{struct_name}::<{generic types}>`   - reference: immutable `&` and mutable `&mut` references.   - generic_type_parameter: it is always start with `T` and following an index number,     which is the position of the generic type parameter in the `struct` or     `function` generic type parameters definition.  Vector type value examples:   * `vector<u8>`   * `vector<vector<u64>>`   * `vector<0x1::AptosAccount::Balance<0x1::XDX::XDX>>`  Struct type value examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Reference type value examples:   * `&signer`   * `&mut address`   * `&mut vector<u8>`  Generic type parameter value example, the following is `0x1::TransactionFee::TransactionFee` JSON representation:      {         \"name\": \"TransactionFee\",         \"is_native\": false,         \"abilities\": [\"key\"],         \"generic_type_params\": [             {\"constraints\": [], \"is_phantom\": true}         ],         \"fields\": [             { \"name\": \"balance\", \"type\": \"0x1::Aptos::Aptos<T0>\" },             { \"name\": \"preburn\", \"type\": \"0x1::Aptos::Preburn<T0>\" }         ]     }  It's Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee<phantom CoinType> has key {             balance: Aptos<CoinType>,             preburn: Preburn<CoinType>,         }     }  The `T0` in the above JSON representation is the generic type place holder for the `CoinType` in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).
	Type string `json:"type"`
}

// NewMoveStructField instantiates a new MoveStructField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveStructField(name string, type_ string) *MoveStructField {
	this := MoveStructField{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewMoveStructFieldWithDefaults instantiates a new MoveStructField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveStructFieldWithDefaults() *MoveStructField {
	this := MoveStructField{}
	return &this
}

// GetName returns the Name field value
func (o *MoveStructField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MoveStructField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MoveStructField) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MoveStructField) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MoveStructField) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MoveStructField) SetType(v string) {
	o.Type = v
}

func (o MoveStructField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMoveStructField struct {
	value *MoveStructField
	isSet bool
}

func (v NullableMoveStructField) Get() *MoveStructField {
	return v.value
}

func (v *NullableMoveStructField) Set(val *MoveStructField) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveStructField) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveStructField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveStructField(val *MoveStructField) *NullableMoveStructField {
	return &NullableMoveStructField{value: val, isSet: true}
}

func (v NullableMoveStructField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveStructField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
