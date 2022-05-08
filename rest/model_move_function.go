/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveFunction struct for MoveFunction
type MoveFunction struct {
	// Move function name
	Name              string                          `json:"name"`
	Visibility        string                          `json:"visibility"`
	GenericTypeParams []MoveFunctionGenericTypeParams `json:"generic_type_params"`
	Params            []string                        `json:"params"`
	Return            []string                        `json:"return"`
}

// NewMoveFunction instantiates a new MoveFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveFunction(name string, visibility string, genericTypeParams []MoveFunctionGenericTypeParams, params []string, return_ []string) *MoveFunction {
	this := MoveFunction{}
	this.Name = name
	this.Visibility = visibility
	this.GenericTypeParams = genericTypeParams
	this.Params = params
	this.Return = return_
	return &this
}

// NewMoveFunctionWithDefaults instantiates a new MoveFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveFunctionWithDefaults() *MoveFunction {
	this := MoveFunction{}
	return &this
}

// GetName returns the Name field value
func (o *MoveFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MoveFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MoveFunction) SetName(v string) {
	o.Name = v
}

// GetVisibility returns the Visibility field value
func (o *MoveFunction) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *MoveFunction) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *MoveFunction) SetVisibility(v string) {
	o.Visibility = v
}

// GetGenericTypeParams returns the GenericTypeParams field value
func (o *MoveFunction) GetGenericTypeParams() []MoveFunctionGenericTypeParams {
	if o == nil {
		var ret []MoveFunctionGenericTypeParams
		return ret
	}

	return o.GenericTypeParams
}

// GetGenericTypeParamsOk returns a tuple with the GenericTypeParams field value
// and a boolean to check if the value has been set.
func (o *MoveFunction) GetGenericTypeParamsOk() ([]MoveFunctionGenericTypeParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenericTypeParams, true
}

// SetGenericTypeParams sets field value
func (o *MoveFunction) SetGenericTypeParams(v []MoveFunctionGenericTypeParams) {
	o.GenericTypeParams = v
}

// GetParams returns the Params field value
func (o *MoveFunction) GetParams() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *MoveFunction) GetParamsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *MoveFunction) SetParams(v []string) {
	o.Params = v
}

// GetReturn returns the Return field value
func (o *MoveFunction) GetReturn() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Return
}

// GetReturnOk returns a tuple with the Return field value
// and a boolean to check if the value has been set.
func (o *MoveFunction) GetReturnOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Return, true
}

// SetReturn sets field value
func (o *MoveFunction) SetReturn(v []string) {
	o.Return = v
}

func (o MoveFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if true {
		toSerialize["generic_type_params"] = o.GenericTypeParams
	}
	if true {
		toSerialize["params"] = o.Params
	}
	if true {
		toSerialize["return"] = o.Return
	}
	return json.Marshal(toSerialize)
}

type NullableMoveFunction struct {
	value *MoveFunction
	isSet bool
}

func (v NullableMoveFunction) Get() *MoveFunction {
	return v.value
}

func (v *NullableMoveFunction) Set(val *MoveFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveFunction(val *MoveFunction) *NullableMoveFunction {
	return &NullableMoveFunction{value: val, isSet: true}
}

func (v NullableMoveFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
