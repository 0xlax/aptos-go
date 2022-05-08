/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
	"fmt"
)

// MoveAbility Abilities are a typing feature in Move that control what actions are permissible for values of a given type.  See [doc](https://diem.github.io/move/abilities.html) for more details.
type MoveAbility string

// List of MoveAbility
const (
	COPY  MoveAbility = "copy"
	DROP  MoveAbility = "drop"
	STORE MoveAbility = "store"
	KEY   MoveAbility = "key"
)

// All allowed values of MoveAbility enum
var AllowedMoveAbilityEnumValues = []MoveAbility{
	"copy",
	"drop",
	"store",
	"key",
}

func (v *MoveAbility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MoveAbility(value)
	for _, existing := range AllowedMoveAbilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MoveAbility", value)
}

// NewMoveAbilityFromValue returns a pointer to a valid MoveAbility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMoveAbilityFromValue(v string) (*MoveAbility, error) {
	ev := MoveAbility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MoveAbility: valid values are %v", v, AllowedMoveAbilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MoveAbility) IsValid() bool {
	for _, existing := range AllowedMoveAbilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MoveAbility value
func (v MoveAbility) Ptr() *MoveAbility {
	return &v
}

type NullableMoveAbility struct {
	value *MoveAbility
	isSet bool
}

func (v NullableMoveAbility) Get() *MoveAbility {
	return v.value
}

func (v *NullableMoveAbility) Set(val *MoveAbility) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveAbility) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveAbility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveAbility(val *MoveAbility) *NullableMoveAbility {
	return &NullableMoveAbility{value: val, isSet: true}
}

func (v NullableMoveAbility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveAbility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
