# ScriptFunctionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Function** | **string** | Script function id is string representation of a script function defined on-chain.  Format: &#x60;{address}::{module name}::{function name}&#x60;  Both &#x60;module name&#x60; and &#x60;function name&#x60; are case-sensitive.  | 
**TypeArguments** | **[]string** | Generic type arguments required by the script function. | 
**Arguments** | **[]interface{}** | The script function arguments. | 

## Methods

### NewScriptFunctionPayload

`func NewScriptFunctionPayload(type_ string, function string, typeArguments []string, arguments []interface{}, ) *ScriptFunctionPayload`

NewScriptFunctionPayload instantiates a new ScriptFunctionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptFunctionPayloadWithDefaults

`func NewScriptFunctionPayloadWithDefaults() *ScriptFunctionPayload`

NewScriptFunctionPayloadWithDefaults instantiates a new ScriptFunctionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScriptFunctionPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptFunctionPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptFunctionPayload) SetType(v string)`

SetType sets Type field to given value.


### GetFunction

`func (o *ScriptFunctionPayload) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ScriptFunctionPayload) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ScriptFunctionPayload) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetTypeArguments

`func (o *ScriptFunctionPayload) GetTypeArguments() []string`

GetTypeArguments returns the TypeArguments field if non-nil, zero value otherwise.

### GetTypeArgumentsOk

`func (o *ScriptFunctionPayload) GetTypeArgumentsOk() (*[]string, bool)`

GetTypeArgumentsOk returns a tuple with the TypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeArguments

`func (o *ScriptFunctionPayload) SetTypeArguments(v []string)`

SetTypeArguments sets TypeArguments field to given value.


### GetArguments

`func (o *ScriptFunctionPayload) GetArguments() []interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ScriptFunctionPayload) GetArgumentsOk() (*[]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ScriptFunctionPayload) SetArguments(v []interface{})`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


