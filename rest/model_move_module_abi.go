/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveModuleABI Move Module ABI is JSON representation of Move module binary interface.
type MoveModuleABI struct {
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	Address          string         `json:"address"`
	Name             string         `json:"name"`
	Friends          []string       `json:"friends"`
	ExposedFunctions []MoveFunction `json:"exposed_functions"`
	Structs          []MoveStruct   `json:"structs"`
}

// NewMoveModuleABI instantiates a new MoveModuleABI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveModuleABI(address string, name string, friends []string, exposedFunctions []MoveFunction, structs []MoveStruct) *MoveModuleABI {
	this := MoveModuleABI{}
	this.Address = address
	this.Name = name
	this.Friends = friends
	this.ExposedFunctions = exposedFunctions
	this.Structs = structs
	return &this
}

// NewMoveModuleABIWithDefaults instantiates a new MoveModuleABI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveModuleABIWithDefaults() *MoveModuleABI {
	this := MoveModuleABI{}
	return &this
}

// GetAddress returns the Address field value
func (o *MoveModuleABI) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *MoveModuleABI) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *MoveModuleABI) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value
func (o *MoveModuleABI) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MoveModuleABI) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MoveModuleABI) SetName(v string) {
	o.Name = v
}

// GetFriends returns the Friends field value
func (o *MoveModuleABI) GetFriends() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Friends
}

// GetFriendsOk returns a tuple with the Friends field value
// and a boolean to check if the value has been set.
func (o *MoveModuleABI) GetFriendsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Friends, true
}

// SetFriends sets field value
func (o *MoveModuleABI) SetFriends(v []string) {
	o.Friends = v
}

// GetExposedFunctions returns the ExposedFunctions field value
func (o *MoveModuleABI) GetExposedFunctions() []MoveFunction {
	if o == nil {
		var ret []MoveFunction
		return ret
	}

	return o.ExposedFunctions
}

// GetExposedFunctionsOk returns a tuple with the ExposedFunctions field value
// and a boolean to check if the value has been set.
func (o *MoveModuleABI) GetExposedFunctionsOk() ([]MoveFunction, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposedFunctions, true
}

// SetExposedFunctions sets field value
func (o *MoveModuleABI) SetExposedFunctions(v []MoveFunction) {
	o.ExposedFunctions = v
}

// GetStructs returns the Structs field value
func (o *MoveModuleABI) GetStructs() []MoveStruct {
	if o == nil {
		var ret []MoveStruct
		return ret
	}

	return o.Structs
}

// GetStructsOk returns a tuple with the Structs field value
// and a boolean to check if the value has been set.
func (o *MoveModuleABI) GetStructsOk() ([]MoveStruct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Structs, true
}

// SetStructs sets field value
func (o *MoveModuleABI) SetStructs(v []MoveStruct) {
	o.Structs = v
}

func (o MoveModuleABI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["friends"] = o.Friends
	}
	if true {
		toSerialize["exposed_functions"] = o.ExposedFunctions
	}
	if true {
		toSerialize["structs"] = o.Structs
	}
	return json.Marshal(toSerialize)
}

type NullableMoveModuleABI struct {
	value *MoveModuleABI
	isSet bool
}

func (v NullableMoveModuleABI) Get() *MoveModuleABI {
	return v.value
}

func (v *NullableMoveModuleABI) Set(val *MoveModuleABI) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveModuleABI) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveModuleABI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveModuleABI(val *MoveModuleABI) *NullableMoveModuleABI {
	return &NullableMoveModuleABI{value: val, isSet: true}
}

func (v NullableMoveModuleABI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveModuleABI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
