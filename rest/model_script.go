/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// Script struct for Script
type Script struct {
	Code          MoveScript    `json:"code"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
}

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(code MoveScript, typeArguments []string, arguments []interface{}) *Script {
	this := Script{}
	this.Code = code
	this.TypeArguments = typeArguments
	this.Arguments = arguments
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetCode returns the Code field value
func (o *Script) GetCode() MoveScript {
	if o == nil {
		var ret MoveScript
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Script) GetCodeOk() (*MoveScript, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Script) SetCode(v MoveScript) {
	o.Code = v
}

// GetTypeArguments returns the TypeArguments field value
func (o *Script) GetTypeArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TypeArguments
}

// GetTypeArgumentsOk returns a tuple with the TypeArguments field value
// and a boolean to check if the value has been set.
func (o *Script) GetTypeArgumentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeArguments, true
}

// SetTypeArguments sets field value
func (o *Script) SetTypeArguments(v []string) {
	o.TypeArguments = v
}

// GetArguments returns the Arguments field value
func (o *Script) GetArguments() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *Script) GetArgumentsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *Script) SetArguments(v []interface{}) {
	o.Arguments = v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["type_arguments"] = o.TypeArguments
	}
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	return json.Marshal(toSerialize)
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
