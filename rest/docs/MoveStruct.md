# MoveStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsNative** | **bool** |  | 
**Abilities** | [**[]MoveAbility**](MoveAbility.md) |  | 
**GenericTypeParams** | [**[]MoveStructGenericTypeParams**](MoveStructGenericTypeParams.md) |  | 
**Fields** | [**[]MoveStructField**](MoveStructField.md) |  | 

## Methods

### NewMoveStruct

`func NewMoveStruct(name string, isNative bool, abilities []MoveAbility, genericTypeParams []MoveStructGenericTypeParams, fields []MoveStructField, ) *MoveStruct`

NewMoveStruct instantiates a new MoveStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveStructWithDefaults

`func NewMoveStructWithDefaults() *MoveStruct`

NewMoveStructWithDefaults instantiates a new MoveStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MoveStruct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveStruct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveStruct) SetName(v string)`

SetName sets Name field to given value.


### GetIsNative

`func (o *MoveStruct) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *MoveStruct) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *MoveStruct) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.


### GetAbilities

`func (o *MoveStruct) GetAbilities() []MoveAbility`

GetAbilities returns the Abilities field if non-nil, zero value otherwise.

### GetAbilitiesOk

`func (o *MoveStruct) GetAbilitiesOk() (*[]MoveAbility, bool)`

GetAbilitiesOk returns a tuple with the Abilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilities

`func (o *MoveStruct) SetAbilities(v []MoveAbility)`

SetAbilities sets Abilities field to given value.


### GetGenericTypeParams

`func (o *MoveStruct) GetGenericTypeParams() []MoveStructGenericTypeParams`

GetGenericTypeParams returns the GenericTypeParams field if non-nil, zero value otherwise.

### GetGenericTypeParamsOk

`func (o *MoveStruct) GetGenericTypeParamsOk() (*[]MoveStructGenericTypeParams, bool)`

GetGenericTypeParamsOk returns a tuple with the GenericTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypeParams

`func (o *MoveStruct) SetGenericTypeParams(v []MoveStructGenericTypeParams)`

SetGenericTypeParams sets GenericTypeParams field to given value.


### GetFields

`func (o *MoveStruct) GetFields() []MoveStructField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MoveStruct) GetFieldsOk() (*[]MoveStructField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MoveStruct) SetFields(v []MoveStructField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


