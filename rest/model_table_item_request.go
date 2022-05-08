/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// TableItemRequest struct for TableItemRequest
type TableItemRequest struct {
	// String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: `vector<{non-reference MoveTypeId}>`   - struct: `{address}::{module_name}::{struct_name}::<{generic types}>`   - reference: immutable `&` and mutable `&mut` references.   - generic_type_parameter: it is always start with `T` and following an index number,     which is the position of the generic type parameter in the `struct` or     `function` generic type parameters definition.  Vector type value examples:   * `vector<u8>`   * `vector<vector<u64>>`   * `vector<0x1::AptosAccount::Balance<0x1::XDX::XDX>>`  Struct type value examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Reference type value examples:   * `&signer`   * `&mut address`   * `&mut vector<u8>`  Generic type parameter value example, the following is `0x1::TransactionFee::TransactionFee` JSON representation:      {         \"name\": \"TransactionFee\",         \"is_native\": false,         \"abilities\": [\"key\"],         \"generic_type_params\": [             {\"constraints\": [], \"is_phantom\": true}         ],         \"fields\": [             { \"name\": \"balance\", \"type\": \"0x1::Aptos::Aptos<T0>\" },             { \"name\": \"preburn\", \"type\": \"0x1::Aptos::Preburn<T0>\" }         ]     }  It's Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee<phantom CoinType> has key {             balance: Aptos<CoinType>,             preburn: Preburn<CoinType>,         }     }  The `T0` in the above JSON representation is the generic type place holder for the `CoinType` in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).
	KeyType string `json:"key_type"`
	// String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: `vector<{non-reference MoveTypeId}>`   - struct: `{address}::{module_name}::{struct_name}::<{generic types}>`   - reference: immutable `&` and mutable `&mut` references.   - generic_type_parameter: it is always start with `T` and following an index number,     which is the position of the generic type parameter in the `struct` or     `function` generic type parameters definition.  Vector type value examples:   * `vector<u8>`   * `vector<vector<u64>>`   * `vector<0x1::AptosAccount::Balance<0x1::XDX::XDX>>`  Struct type value examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Reference type value examples:   * `&signer`   * `&mut address`   * `&mut vector<u8>`  Generic type parameter value example, the following is `0x1::TransactionFee::TransactionFee` JSON representation:      {         \"name\": \"TransactionFee\",         \"is_native\": false,         \"abilities\": [\"key\"],         \"generic_type_params\": [             {\"constraints\": [], \"is_phantom\": true}         ],         \"fields\": [             { \"name\": \"balance\", \"type\": \"0x1::Aptos::Aptos<T0>\" },             { \"name\": \"preburn\", \"type\": \"0x1::Aptos::Preburn<T0>\" }         ]     }  It's Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee<phantom CoinType> has key {             balance: Aptos<CoinType>,             preburn: Preburn<CoinType>,         }     }  The `T0` in the above JSON representation is the generic type place holder for the `CoinType` in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).
	ValueType string `json:"value_type"`
	// Move `bool` type value is serialized into `boolean`.  Move `u8` type value is serialized into `integer`.  Move `u64` and `u128` type value is serialized into `string`.  Move `address` type value(16 bytes Aptos account address) is serialized into hex-encoded string, which is prefixed with `0x` and leading zeros are trimmed.  For example:   * `0x1`   * `0x1668f6be25668c1a17cd8caf6b8d2f25`  Move `vector` type value is serialized into `array`, except `vector<u8>` which is serialized into hex-encoded string with `0x` prefix.  For example:   * `vector<u64>{255, 255}` => `[\"255\", \"255\"]`   * `vector<u8>{255, 255}` => `0xffff`  Move `struct` type value is serialized into `object` that looks like this (except some Move stdlib types, see the following section):    ```json   {     field1_name: field1_value,     field2_name: field2_value,     ......   }   ```  For example:   `{ \"created\": \"0xa550c18\", \"role_id\": \"0\" }`  **Special serialization for Move stdlib types:**  * [0x1::ASCII::String](https://github.com/aptos-labs/aptos-core/blob/main/language/move-stdlib/docs/ASCII.md) is serialized into `string`. For example, struct value `0x1::ASCII::String{bytes: b\"hello world\"}` is serialized as `\"hello world\"` in JSON.
	Key interface{} `json:"key"`
}

// NewTableItemRequest instantiates a new TableItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableItemRequest(keyType string, valueType string, key interface{}) *TableItemRequest {
	this := TableItemRequest{}
	this.KeyType = keyType
	this.ValueType = valueType
	this.Key = key
	return &this
}

// NewTableItemRequestWithDefaults instantiates a new TableItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableItemRequestWithDefaults() *TableItemRequest {
	this := TableItemRequest{}
	return &this
}

// GetKeyType returns the KeyType field value
func (o *TableItemRequest) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *TableItemRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *TableItemRequest) SetKeyType(v string) {
	o.KeyType = v
}

// GetValueType returns the ValueType field value
func (o *TableItemRequest) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *TableItemRequest) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *TableItemRequest) SetValueType(v string) {
	o.ValueType = v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TableItemRequest) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TableItemRequest) GetKeyOk() (*interface{}, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TableItemRequest) SetKey(v interface{}) {
	o.Key = v
}

func (o TableItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_type"] = o.KeyType
	}
	if true {
		toSerialize["value_type"] = o.ValueType
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableTableItemRequest struct {
	value *TableItemRequest
	isSet bool
}

func (v NullableTableItemRequest) Get() *TableItemRequest {
	return v.value
}

func (v *NullableTableItemRequest) Set(val *TableItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTableItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTableItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableItemRequest(val *TableItemRequest) *NullableTableItemRequest {
	return &NullableTableItemRequest{value: val, isSet: true}
}

func (v NullableTableItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
