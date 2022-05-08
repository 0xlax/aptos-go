/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// AccountResource Account resource is a Move struct value belongs to an account.
type AccountResource struct {
	// String representation of an on-chain Move struct type.  It is a combination of:   1. `Move module address`, `module name` and `struct name` joined by `::`.   2. `struct generic type parameters` joined by `, `.  Examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  See [doc](https://diem.github.io/move/structs-and-resources.html) for more details.
	Type string `json:"type"`
	// Account resource data is JSON representation of the Move struct `type`.  Move struct field name and value are serialized as object property name and value.
	Data map[string]interface{} `json:"data"`
}

// NewAccountResource instantiates a new AccountResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResource(type_ string, data map[string]interface{}) *AccountResource {
	this := AccountResource{}
	this.Type = type_
	this.Data = data
	return &this
}

// NewAccountResourceWithDefaults instantiates a new AccountResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResourceWithDefaults() *AccountResource {
	this := AccountResource{}
	return &this
}

// GetType returns the Type field value
func (o *AccountResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountResource) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *AccountResource) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountResource) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AccountResource) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o AccountResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAccountResource struct {
	value *AccountResource
	isSet bool
}

func (v NullableAccountResource) Get() *AccountResource {
	return v.value
}

func (v *NullableAccountResource) Set(val *AccountResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResource(val *AccountResource) *NullableAccountResource {
	return &NullableAccountResource{value: val, isSet: true}
}

func (v NullableAccountResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
