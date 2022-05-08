/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// MoveStruct struct for MoveStruct
type MoveStruct struct {
	Name              string                        `json:"name"`
	IsNative          bool                          `json:"is_native"`
	Abilities         []MoveAbility                 `json:"abilities"`
	GenericTypeParams []MoveStructGenericTypeParams `json:"generic_type_params"`
	Fields            []MoveStructField             `json:"fields"`
}

// NewMoveStruct instantiates a new MoveStruct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveStruct(name string, isNative bool, abilities []MoveAbility, genericTypeParams []MoveStructGenericTypeParams, fields []MoveStructField) *MoveStruct {
	this := MoveStruct{}
	this.Name = name
	this.IsNative = isNative
	this.Abilities = abilities
	this.GenericTypeParams = genericTypeParams
	this.Fields = fields
	return &this
}

// NewMoveStructWithDefaults instantiates a new MoveStruct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveStructWithDefaults() *MoveStruct {
	this := MoveStruct{}
	return &this
}

// GetName returns the Name field value
func (o *MoveStruct) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MoveStruct) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MoveStruct) SetName(v string) {
	o.Name = v
}

// GetIsNative returns the IsNative field value
func (o *MoveStruct) GetIsNative() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNative
}

// GetIsNativeOk returns a tuple with the IsNative field value
// and a boolean to check if the value has been set.
func (o *MoveStruct) GetIsNativeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNative, true
}

// SetIsNative sets field value
func (o *MoveStruct) SetIsNative(v bool) {
	o.IsNative = v
}

// GetAbilities returns the Abilities field value
func (o *MoveStruct) GetAbilities() []MoveAbility {
	if o == nil {
		var ret []MoveAbility
		return ret
	}

	return o.Abilities
}

// GetAbilitiesOk returns a tuple with the Abilities field value
// and a boolean to check if the value has been set.
func (o *MoveStruct) GetAbilitiesOk() ([]MoveAbility, bool) {
	if o == nil {
		return nil, false
	}
	return o.Abilities, true
}

// SetAbilities sets field value
func (o *MoveStruct) SetAbilities(v []MoveAbility) {
	o.Abilities = v
}

// GetGenericTypeParams returns the GenericTypeParams field value
func (o *MoveStruct) GetGenericTypeParams() []MoveStructGenericTypeParams {
	if o == nil {
		var ret []MoveStructGenericTypeParams
		return ret
	}

	return o.GenericTypeParams
}

// GetGenericTypeParamsOk returns a tuple with the GenericTypeParams field value
// and a boolean to check if the value has been set.
func (o *MoveStruct) GetGenericTypeParamsOk() ([]MoveStructGenericTypeParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenericTypeParams, true
}

// SetGenericTypeParams sets field value
func (o *MoveStruct) SetGenericTypeParams(v []MoveStructGenericTypeParams) {
	o.GenericTypeParams = v
}

// GetFields returns the Fields field value
func (o *MoveStruct) GetFields() []MoveStructField {
	if o == nil {
		var ret []MoveStructField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *MoveStruct) GetFieldsOk() ([]MoveStructField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *MoveStruct) SetFields(v []MoveStructField) {
	o.Fields = v
}

func (o MoveStruct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["is_native"] = o.IsNative
	}
	if true {
		toSerialize["abilities"] = o.Abilities
	}
	if true {
		toSerialize["generic_type_params"] = o.GenericTypeParams
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableMoveStruct struct {
	value *MoveStruct
	isSet bool
}

func (v NullableMoveStruct) Get() *MoveStruct {
	return v.value
}

func (v *NullableMoveStruct) Set(val *MoveStruct) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveStruct) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveStruct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveStruct(val *MoveStruct) *NullableMoveStruct {
	return &NullableMoveStruct{value: val, isSet: true}
}

func (v NullableMoveStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveStruct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
