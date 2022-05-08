# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**MoveScript**](MoveScript.md) |  | 
**TypeArguments** | **[]string** |  | 
**Arguments** | **[]interface{}** |  | 

## Methods

### NewScript

`func NewScript(code MoveScript, typeArguments []string, arguments []interface{}, ) *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Script) GetCode() MoveScript`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Script) GetCodeOk() (*MoveScript, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Script) SetCode(v MoveScript)`

SetCode sets Code field to given value.


### GetTypeArguments

`func (o *Script) GetTypeArguments() []string`

GetTypeArguments returns the TypeArguments field if non-nil, zero value otherwise.

### GetTypeArgumentsOk

`func (o *Script) GetTypeArgumentsOk() (*[]string, bool)`

GetTypeArgumentsOk returns a tuple with the TypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeArguments

`func (o *Script) SetTypeArguments(v []string)`

SetTypeArguments sets TypeArguments field to given value.


### GetArguments

`func (o *Script) GetArguments() []interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Script) GetArgumentsOk() (*[]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Script) SetArguments(v []interface{})`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


