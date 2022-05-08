# MoveFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Move function name | 
**Visibility** | **string** |  | 
**GenericTypeParams** | [**[]MoveFunctionGenericTypeParams**](MoveFunctionGenericTypeParams.md) |  | 
**Params** | **[]string** |  | 
**Return** | **[]string** |  | 

## Methods

### NewMoveFunction

`func NewMoveFunction(name string, visibility string, genericTypeParams []MoveFunctionGenericTypeParams, params []string, return_ []string, ) *MoveFunction`

NewMoveFunction instantiates a new MoveFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveFunctionWithDefaults

`func NewMoveFunctionWithDefaults() *MoveFunction`

NewMoveFunctionWithDefaults instantiates a new MoveFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MoveFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveFunction) SetName(v string)`

SetName sets Name field to given value.


### GetVisibility

`func (o *MoveFunction) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MoveFunction) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MoveFunction) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetGenericTypeParams

`func (o *MoveFunction) GetGenericTypeParams() []MoveFunctionGenericTypeParams`

GetGenericTypeParams returns the GenericTypeParams field if non-nil, zero value otherwise.

### GetGenericTypeParamsOk

`func (o *MoveFunction) GetGenericTypeParamsOk() (*[]MoveFunctionGenericTypeParams, bool)`

GetGenericTypeParamsOk returns a tuple with the GenericTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypeParams

`func (o *MoveFunction) SetGenericTypeParams(v []MoveFunctionGenericTypeParams)`

SetGenericTypeParams sets GenericTypeParams field to given value.


### GetParams

`func (o *MoveFunction) GetParams() []string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MoveFunction) GetParamsOk() (*[]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MoveFunction) SetParams(v []string)`

SetParams sets Params field to given value.


### GetReturn

`func (o *MoveFunction) GetReturn() []string`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *MoveFunction) GetReturnOk() (*[]string, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *MoveFunction) SetReturn(v []string)`

SetReturn sets Return field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


