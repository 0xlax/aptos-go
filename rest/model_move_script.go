/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveScript struct for MoveScript
type MoveScript struct {
	// All bytes data are represented as hex-encoded string prefixed with `0x` and fulfilled with two hex digits per byte.  Different with `Address` type, hex-encoded bytes should not trim any zeros.
	Bytecode string        `json:"bytecode"`
	Abi      *MoveFunction `json:"abi,omitempty"`
}

// NewMoveScript instantiates a new MoveScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveScript(bytecode string) *MoveScript {
	this := MoveScript{}
	this.Bytecode = bytecode
	return &this
}

// NewMoveScriptWithDefaults instantiates a new MoveScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveScriptWithDefaults() *MoveScript {
	this := MoveScript{}
	return &this
}

// GetBytecode returns the Bytecode field value
func (o *MoveScript) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *MoveScript) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *MoveScript) SetBytecode(v string) {
	o.Bytecode = v
}

// GetAbi returns the Abi field value if set, zero value otherwise.
func (o *MoveScript) GetAbi() MoveFunction {
	if o == nil || o.Abi == nil {
		var ret MoveFunction
		return ret
	}
	return *o.Abi
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveScript) GetAbiOk() (*MoveFunction, bool) {
	if o == nil || o.Abi == nil {
		return nil, false
	}
	return o.Abi, true
}

// HasAbi returns a boolean if a field has been set.
func (o *MoveScript) HasAbi() bool {
	if o != nil && o.Abi != nil {
		return true
	}

	return false
}

// SetAbi gets a reference to the given MoveFunction and assigns it to the Abi field.
func (o *MoveScript) SetAbi(v MoveFunction) {
	o.Abi = &v
}

func (o MoveScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bytecode"] = o.Bytecode
	}
	if o.Abi != nil {
		toSerialize["abi"] = o.Abi
	}
	return json.Marshal(toSerialize)
}

type NullableMoveScript struct {
	value *MoveScript
	isSet bool
}

func (v NullableMoveScript) Get() *MoveScript {
	return v.value
}

func (v *NullableMoveScript) Set(val *MoveScript) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveScript) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveScript(val *MoveScript) *NullableMoveScript {
	return &NullableMoveScript{value: val, isSet: true}
}

func (v NullableMoveScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
