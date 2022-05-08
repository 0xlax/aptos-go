/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveFunctionGenericTypeParams struct for MoveFunctionGenericTypeParams
type MoveFunctionGenericTypeParams struct {
	Constraints []MoveAbility `json:"constraints"`
}

// NewMoveFunctionGenericTypeParams instantiates a new MoveFunctionGenericTypeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveFunctionGenericTypeParams(constraints []MoveAbility) *MoveFunctionGenericTypeParams {
	this := MoveFunctionGenericTypeParams{}
	this.Constraints = constraints
	return &this
}

// NewMoveFunctionGenericTypeParamsWithDefaults instantiates a new MoveFunctionGenericTypeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveFunctionGenericTypeParamsWithDefaults() *MoveFunctionGenericTypeParams {
	this := MoveFunctionGenericTypeParams{}
	return &this
}

// GetConstraints returns the Constraints field value
func (o *MoveFunctionGenericTypeParams) GetConstraints() []MoveAbility {
	if o == nil {
		var ret []MoveAbility
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *MoveFunctionGenericTypeParams) GetConstraintsOk() ([]MoveAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *MoveFunctionGenericTypeParams) SetConstraints(v []MoveAbility) {
	o.Constraints = v
}

func (o MoveFunctionGenericTypeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["constraints"] = o.Constraints
	}
	return json.Marshal(toSerialize)
}

type NullableMoveFunctionGenericTypeParams struct {
	value *MoveFunctionGenericTypeParams
	isSet bool
}

func (v NullableMoveFunctionGenericTypeParams) Get() *MoveFunctionGenericTypeParams {
	return v.value
}

func (v *NullableMoveFunctionGenericTypeParams) Set(val *MoveFunctionGenericTypeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveFunctionGenericTypeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveFunctionGenericTypeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveFunctionGenericTypeParams(val *MoveFunctionGenericTypeParams) *NullableMoveFunctionGenericTypeParams {
	return &NullableMoveFunctionGenericTypeParams{value: val, isSet: true}
}

func (v NullableMoveFunctionGenericTypeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveFunctionGenericTypeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
