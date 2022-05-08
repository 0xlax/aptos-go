/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// ScriptWriteSet struct for ScriptWriteSet
type ScriptWriteSet struct {
	Type string `json:"type"`
	// Hex-encoded 16 bytes Aptos account address.  Prefixed with `0x` and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.
	ExecuteAs string `json:"execute_as"`
	Script    Script `json:"script"`
}

// NewScriptWriteSet instantiates a new ScriptWriteSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptWriteSet(type_ string, executeAs string, script Script) *ScriptWriteSet {
	this := ScriptWriteSet{}
	this.Type = type_
	this.ExecuteAs = executeAs
	this.Script = script
	return &this
}

// NewScriptWriteSetWithDefaults instantiates a new ScriptWriteSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWriteSetWithDefaults() *ScriptWriteSet {
	this := ScriptWriteSet{}
	return &this
}

// GetType returns the Type field value
func (o *ScriptWriteSet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScriptWriteSet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScriptWriteSet) SetType(v string) {
	o.Type = v
}

// GetExecuteAs returns the ExecuteAs field value
func (o *ScriptWriteSet) GetExecuteAs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecuteAs
}

// GetExecuteAsOk returns a tuple with the ExecuteAs field value
// and a boolean to check if the value has been set.
func (o *ScriptWriteSet) GetExecuteAsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecuteAs, true
}

// SetExecuteAs sets field value
func (o *ScriptWriteSet) SetExecuteAs(v string) {
	o.ExecuteAs = v
}

// GetScript returns the Script field value
func (o *ScriptWriteSet) GetScript() Script {
	if o == nil {
		var ret Script
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *ScriptWriteSet) GetScriptOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *ScriptWriteSet) SetScript(v Script) {
	o.Script = v
}

func (o ScriptWriteSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["execute_as"] = o.ExecuteAs
	}
	if true {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullableScriptWriteSet struct {
	value *ScriptWriteSet
	isSet bool
}

func (v NullableScriptWriteSet) Get() *ScriptWriteSet {
	return v.value
}

func (v *NullableScriptWriteSet) Set(val *ScriptWriteSet) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptWriteSet) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptWriteSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptWriteSet(val *ScriptWriteSet) *NullableScriptWriteSet {
	return &NullableScriptWriteSet{value: val, isSet: true}
}

func (v NullableScriptWriteSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptWriteSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
