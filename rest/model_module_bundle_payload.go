/*
Aptos Dev API Specification

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.

API version: 0.1.1
*/

package openapi

import (
	"encoding/json"
)

// ModuleBundlePayload struct for ModuleBundlePayload
type ModuleBundlePayload struct {
	Type    string       `json:"type"`
	Modules []MoveModule `json:"modules"`
}

// NewModuleBundlePayload instantiates a new ModuleBundlePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBundlePayload(type_ string, modules []MoveModule) *ModuleBundlePayload {
	this := ModuleBundlePayload{}
	this.Type = type_
	this.Modules = modules
	return &this
}

// NewModuleBundlePayloadWithDefaults instantiates a new ModuleBundlePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBundlePayloadWithDefaults() *ModuleBundlePayload {
	this := ModuleBundlePayload{}
	return &this
}

// GetType returns the Type field value
func (o *ModuleBundlePayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModuleBundlePayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModuleBundlePayload) SetType(v string) {
	o.Type = v
}

// GetModules returns the Modules field value
func (o *ModuleBundlePayload) GetModules() []MoveModule {
	if o == nil {
		var ret []MoveModule
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *ModuleBundlePayload) GetModulesOk() ([]MoveModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *ModuleBundlePayload) SetModules(v []MoveModule) {
	o.Modules = v
}

func (o ModuleBundlePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["modules"] = o.Modules
	}
	return json.Marshal(toSerialize)
}

type NullableModuleBundlePayload struct {
	value *ModuleBundlePayload
	isSet bool
}

func (v NullableModuleBundlePayload) Get() *ModuleBundlePayload {
	return v.value
}

func (v *NullableModuleBundlePayload) Set(val *ModuleBundlePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBundlePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBundlePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBundlePayload(val *ModuleBundlePayload) *NullableModuleBundlePayload {
	return &NullableModuleBundlePayload{value: val, isSet: true}
}

func (v NullableModuleBundlePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBundlePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
