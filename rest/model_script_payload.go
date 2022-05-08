/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// ScriptPayload struct for ScriptPayload
type ScriptPayload struct {
	Type          string        `json:"type"`
	Code          MoveScript    `json:"code"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
}

// NewScriptPayload instantiates a new ScriptPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptPayload(type_ string, code MoveScript, typeArguments []string, arguments []interface{}) *ScriptPayload {
	this := ScriptPayload{}
	this.Type = type_
	this.Code = code
	this.TypeArguments = typeArguments
	this.Arguments = arguments
	return &this
}

// NewScriptPayloadWithDefaults instantiates a new ScriptPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptPayloadWithDefaults() *ScriptPayload {
	this := ScriptPayload{}
	return &this
}

// GetType returns the Type field value
func (o *ScriptPayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScriptPayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScriptPayload) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *ScriptPayload) GetCode() MoveScript {
	if o == nil {
		var ret MoveScript
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ScriptPayload) GetCodeOk() (*MoveScript, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ScriptPayload) SetCode(v MoveScript) {
	o.Code = v
}

// GetTypeArguments returns the TypeArguments field value
func (o *ScriptPayload) GetTypeArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TypeArguments
}

// GetTypeArgumentsOk returns a tuple with the TypeArguments field value
// and a boolean to check if the value has been set.
func (o *ScriptPayload) GetTypeArgumentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeArguments, true
}

// SetTypeArguments sets field value
func (o *ScriptPayload) SetTypeArguments(v []string) {
	o.TypeArguments = v
}

// GetArguments returns the Arguments field value
func (o *ScriptPayload) GetArguments() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *ScriptPayload) GetArgumentsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *ScriptPayload) SetArguments(v []interface{}) {
	o.Arguments = v
}

func (o ScriptPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
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

type NullableScriptPayload struct {
	value *ScriptPayload
	isSet bool
}

func (v NullableScriptPayload) Get() *ScriptPayload {
	return v.value
}

func (v *NullableScriptPayload) Set(val *ScriptPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptPayload(val *ScriptPayload) *NullableScriptPayload {
	return &NullableScriptPayload{value: val, isSet: true}
}

func (v NullableScriptPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
