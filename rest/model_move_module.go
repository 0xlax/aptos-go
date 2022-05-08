/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveModule struct for MoveModule
type MoveModule struct {
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Bytecode string         `json:"bytecode"`
	Abi      *MoveModuleABI `json:"abi,omitempty"`
}

// NewMoveModule instantiates a new MoveModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveModule(bytecode string) *MoveModule {
	this := MoveModule{}
	this.Bytecode = bytecode
	return &this
}

// NewMoveModuleWithDefaults instantiates a new MoveModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveModuleWithDefaults() *MoveModule {
	this := MoveModule{}
	return &this
}

// GetBytecode returns the Bytecode field value
func (o *MoveModule) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *MoveModule) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *MoveModule) SetBytecode(v string) {
	o.Bytecode = v
}

// GetAbi returns the Abi field value if set, zero value otherwise.
func (o *MoveModule) GetAbi() MoveModuleABI {
	if o == nil || o.Abi == nil {
		var ret MoveModuleABI
		return ret
	}
	return *o.Abi
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveModule) GetAbiOk() (*MoveModuleABI, bool) {
	if o == nil || o.Abi == nil {
		return nil, false
	}
	return o.Abi, true
}

// HasAbi returns a boolean if a field has been set.
func (o *MoveModule) HasAbi() bool {
	if o != nil && o.Abi != nil {
		return true
	}

	return false
}

// SetAbi gets a reference to the given MoveModuleABI and assigns it to the Abi field.
func (o *MoveModule) SetAbi(v MoveModuleABI) {
	o.Abi = &v
}

func (o MoveModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bytecode"] = o.Bytecode
	}
	if o.Abi != nil {
		toSerialize["abi"] = o.Abi
	}
	return json.Marshal(toSerialize)
}

type NullableMoveModule struct {
	value *MoveModule
	isSet bool
}

func (v NullableMoveModule) Get() *MoveModule {
	return v.value
}

func (v *NullableMoveModule) Set(val *MoveModule) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveModule) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveModule(val *MoveModule) *NullableMoveModule {
	return &NullableMoveModule{value: val, isSet: true}
}

func (v NullableMoveModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
