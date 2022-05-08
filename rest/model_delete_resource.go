/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// DeleteResource Delete account resource change.
type DeleteResource struct {
	Type string `json:"type"`
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	StateKeyHash string `json:"state_key_hash"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Address string `json:"address"`
	// String representation of an on-chain Move struct type.  It is a combination of:   1. `Move module address`, `module name` and `struct name` joined by `::`.   2. `struct generic type parameters` joined by `, `.  Examples:   * `0x1::Aptos::Aptos<0x1::XDX::XDX>`   * `0x1::Abc::Abc<vector<u8>, vector<u64>>`   * `0x1::AptosAccount::AccountOperationsCapability`  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  See [doc](https://diem.github.io/move/structs-and-resources.html) for more details.
	Resource string `json:"resource"`
}

// NewDeleteResource instantiates a new DeleteResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResource(type_ string, stateKeyHash string, address string, resource string) *DeleteResource {
	this := DeleteResource{}
	this.Type = type_
	this.StateKeyHash = stateKeyHash
	this.Address = address
	this.Resource = resource
	return &this
}

// NewDeleteResourceWithDefaults instantiates a new DeleteResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResourceWithDefaults() *DeleteResource {
	this := DeleteResource{}
	return &this
}

// GetType returns the Type field value
func (o *DeleteResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeleteResource) SetType(v string) {
	o.Type = v
}

// GetStateKeyHash returns the StateKeyHash field value
func (o *DeleteResource) GetStateKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateKeyHash
}

// GetStateKeyHashOk returns a tuple with the StateKeyHash field value
// and a boolean to check if the value has been set.
func (o *DeleteResource) GetStateKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateKeyHash, true
}

// SetStateKeyHash sets field value
func (o *DeleteResource) SetStateKeyHash(v string) {
	o.StateKeyHash = v
}

// GetAddress returns the Address field value
func (o *DeleteResource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DeleteResource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DeleteResource) SetAddress(v string) {
	o.Address = v
}

// GetResource returns the Resource field value
func (o *DeleteResource) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *DeleteResource) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *DeleteResource) SetResource(v string) {
	o.Resource = v
}

func (o DeleteResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state_key_hash"] = o.StateKeyHash
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteResource struct {
	value *DeleteResource
	isSet bool
}

func (v NullableDeleteResource) Get() *DeleteResource {
	return v.value
}

func (v *NullableDeleteResource) Set(val *DeleteResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResource(val *DeleteResource) *NullableDeleteResource {
	return &NullableDeleteResource{value: val, isSet: true}
}

func (v NullableDeleteResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
