/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveStructGenericTypeParams struct for MoveStructGenericTypeParams
type MoveStructGenericTypeParams struct {
	Constraints []MoveAbility `json:"constraints"`
	IsPhantom   bool          `json:"is_phantom"`
}

// NewMoveStructGenericTypeParams instantiates a new MoveStructGenericTypeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveStructGenericTypeParams(constraints []MoveAbility, isPhantom bool) *MoveStructGenericTypeParams {
	this := MoveStructGenericTypeParams{}
	this.Constraints = constraints
	this.IsPhantom = isPhantom
	return &this
}

// NewMoveStructGenericTypeParamsWithDefaults instantiates a new MoveStructGenericTypeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveStructGenericTypeParamsWithDefaults() *MoveStructGenericTypeParams {
	this := MoveStructGenericTypeParams{}
	return &this
}

// GetConstraints returns the Constraints field value
func (o *MoveStructGenericTypeParams) GetConstraints() []MoveAbility {
	if o == nil {
		var ret []MoveAbility
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *MoveStructGenericTypeParams) GetConstraintsOk() ([]MoveAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *MoveStructGenericTypeParams) SetConstraints(v []MoveAbility) {
	o.Constraints = v
}

// GetIsPhantom returns the IsPhantom field value
func (o *MoveStructGenericTypeParams) GetIsPhantom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPhantom
}

// GetIsPhantomOk returns a tuple with the IsPhantom field value
// and a boolean to check if the value has been set.
func (o *MoveStructGenericTypeParams) GetIsPhantomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPhantom, true
}

// SetIsPhantom sets field value
func (o *MoveStructGenericTypeParams) SetIsPhantom(v bool) {
	o.IsPhantom = v
}

func (o MoveStructGenericTypeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["constraints"] = o.Constraints
	}
	if true {
		toSerialize["is_phantom"] = o.IsPhantom
	}
	return json.Marshal(toSerialize)
}

type NullableMoveStructGenericTypeParams struct {
	value *MoveStructGenericTypeParams
	isSet bool
}

func (v NullableMoveStructGenericTypeParams) Get() *MoveStructGenericTypeParams {
	return v.value
}

func (v *NullableMoveStructGenericTypeParams) Set(val *MoveStructGenericTypeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveStructGenericTypeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveStructGenericTypeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveStructGenericTypeParams(val *MoveStructGenericTypeParams) *NullableMoveStructGenericTypeParams {
	return &NullableMoveStructGenericTypeParams{value: val, isSet: true}
}

func (v NullableMoveStructGenericTypeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveStructGenericTypeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
