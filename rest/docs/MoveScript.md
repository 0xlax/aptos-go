# MoveScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytecode** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Abi** | Pointer to [**MoveFunction**](MoveFunction.md) |  | [optional] 

## Methods

### NewMoveScript

`func NewMoveScript(bytecode string, ) *MoveScript`

NewMoveScript instantiates a new MoveScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveScriptWithDefaults

`func NewMoveScriptWithDefaults() *MoveScript`

NewMoveScriptWithDefaults instantiates a new MoveScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytecode

`func (o *MoveScript) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *MoveScript) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *MoveScript) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetAbi

`func (o *MoveScript) GetAbi() MoveFunction`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *MoveScript) GetAbiOk() (*MoveFunction, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *MoveScript) SetAbi(v MoveFunction)`

SetAbi sets Abi field to given value.

### HasAbi

`func (o *MoveScript) HasAbi() bool`

HasAbi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


