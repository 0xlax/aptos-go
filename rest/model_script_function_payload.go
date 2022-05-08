/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// ScriptFunctionPayload struct for ScriptFunctionPayload
type ScriptFunctionPayload struct {
	Type string `json:"type"`
	// Script function id is string representation of a script function defined on-chain.  Format: `{address}::{module name}::{function name}`  Both `module name` and `function name` are case-sensitive.
	Function string `json:"function"`
	// Generic type arguments required by the script function.
	TypeArguments []string `json:"type_arguments"`
	// The script function arguments.
	Arguments []interface{} `json:"arguments"`
}

// NewScriptFunctionPayload instantiates a new ScriptFunctionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptFunctionPayload(type_ string, function string, typeArguments []string, arguments []interface{}) *ScriptFunctionPayload {
	this := ScriptFunctionPayload{}
	this.Type = type_
	this.Function = function
	this.TypeArguments = typeArguments
	this.Arguments = arguments
	return &this
}

// NewScriptFunctionPayloadWithDefaults instantiates a new ScriptFunctionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptFunctionPayloadWithDefaults() *ScriptFunctionPayload {
	this := ScriptFunctionPayload{}
	return &this
}

// GetType returns the Type field value
func (o *ScriptFunctionPayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScriptFunctionPayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScriptFunctionPayload) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value
func (o *ScriptFunctionPayload) GetFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *ScriptFunctionPayload) GetFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *ScriptFunctionPayload) SetFunction(v string) {
	o.Function = v
}

// GetTypeArguments returns the TypeArguments field value
func (o *ScriptFunctionPayload) GetTypeArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TypeArguments
}

// GetTypeArgumentsOk returns a tuple with the TypeArguments field value
// and a boolean to check if the value has been set.
func (o *ScriptFunctionPayload) GetTypeArgumentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeArguments, true
}

// SetTypeArguments sets field value
func (o *ScriptFunctionPayload) SetTypeArguments(v []string) {
	o.TypeArguments = v
}

// GetArguments returns the Arguments field value
func (o *ScriptFunctionPayload) GetArguments() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *ScriptFunctionPayload) GetArgumentsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *ScriptFunctionPayload) SetArguments(v []interface{}) {
	o.Arguments = v
}

func (o ScriptFunctionPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["function"] = o.Function
	}
	if true {
		toSerialize["type_arguments"] = o.TypeArguments
	}
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	return json.Marshal(toSerialize)
}

type NullableScriptFunctionPayload struct {
	value *ScriptFunctionPayload
	isSet bool
}

func (v NullableScriptFunctionPayload) Get() *ScriptFunctionPayload {
	return v.value
}

func (v *NullableScriptFunctionPayload) Set(val *ScriptFunctionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptFunctionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptFunctionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptFunctionPayload(val *ScriptFunctionPayload) *NullableScriptFunctionPayload {
	return &NullableScriptFunctionPayload{value: val, isSet: true}
}

func (v NullableScriptFunctionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptFunctionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
