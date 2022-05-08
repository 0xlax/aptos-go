# ScriptPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Code** | [**MoveScript**](MoveScript.md) |  | 
**TypeArguments** | **[]string** |  | 
**Arguments** | **[]interface{}** |  | 

## Methods

### NewScriptPayload

`func NewScriptPayload(type_ string, code MoveScript, typeArguments []string, arguments []interface{}, ) *ScriptPayload`

NewScriptPayload instantiates a new ScriptPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptPayloadWithDefaults

`func NewScriptPayloadWithDefaults() *ScriptPayload`

NewScriptPayloadWithDefaults instantiates a new ScriptPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScriptPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptPayload) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ScriptPayload) GetCode() MoveScript`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ScriptPayload) GetCodeOk() (*MoveScript, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ScriptPayload) SetCode(v MoveScript)`

SetCode sets Code field to given value.


### GetTypeArguments

`func (o *ScriptPayload) GetTypeArguments() []string`

GetTypeArguments returns the TypeArguments field if non-nil, zero value otherwise.

### GetTypeArgumentsOk

`func (o *ScriptPayload) GetTypeArgumentsOk() (*[]string, bool)`

GetTypeArgumentsOk returns a tuple with the TypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeArguments

`func (o *ScriptPayload) SetTypeArguments(v []string)`

SetTypeArguments sets TypeArguments field to given value.


### GetArguments

`func (o *ScriptPayload) GetArguments() []interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ScriptPayload) GetArgumentsOk() (*[]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ScriptPayload) SetArguments(v []interface{})`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


