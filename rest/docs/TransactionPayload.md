# TransactionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Function** | **string** | Script function id is string representation of a script function defined on-chain.  Format: &#x60;{address}::{module name}::{function name}&#x60;  Both &#x60;module name&#x60; and &#x60;function name&#x60; are case-sensitive.  | 
**TypeArguments** | **[]string** |  | 
**Arguments** | **[]interface{}** |  | 
**Code** | [**MoveScript**](MoveScript.md) |  | 
**Modules** | [**[]MoveModule**](MoveModule.md) |  | 
**WriteSet** | [**WriteSet**](WriteSet.md) |  | 

## Methods

### NewTransactionPayload

`func NewTransactionPayload(type_ string, function string, typeArguments []string, arguments []interface{}, code MoveScript, modules []MoveModule, writeSet WriteSet, ) *TransactionPayload`

NewTransactionPayload instantiates a new TransactionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPayloadWithDefaults

`func NewTransactionPayloadWithDefaults() *TransactionPayload`

NewTransactionPayloadWithDefaults instantiates a new TransactionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPayload) SetType(v string)`

SetType sets Type field to given value.


### GetFunction

`func (o *TransactionPayload) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *TransactionPayload) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *TransactionPayload) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetTypeArguments

`func (o *TransactionPayload) GetTypeArguments() []string`

GetTypeArguments returns the TypeArguments field if non-nil, zero value otherwise.

### GetTypeArgumentsOk

`func (o *TransactionPayload) GetTypeArgumentsOk() (*[]string, bool)`

GetTypeArgumentsOk returns a tuple with the TypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeArguments

`func (o *TransactionPayload) SetTypeArguments(v []string)`

SetTypeArguments sets TypeArguments field to given value.


### GetArguments

`func (o *TransactionPayload) GetArguments() []interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *TransactionPayload) GetArgumentsOk() (*[]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *TransactionPayload) SetArguments(v []interface{})`

SetArguments sets Arguments field to given value.


### GetCode

`func (o *TransactionPayload) GetCode() MoveScript`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransactionPayload) GetCodeOk() (*MoveScript, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransactionPayload) SetCode(v MoveScript)`

SetCode sets Code field to given value.


### GetModules

`func (o *TransactionPayload) GetModules() []MoveModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *TransactionPayload) GetModulesOk() (*[]MoveModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *TransactionPayload) SetModules(v []MoveModule)`

SetModules sets Modules field to given value.


### GetWriteSet

`func (o *TransactionPayload) GetWriteSet() WriteSet`

GetWriteSet returns the WriteSet field if non-nil, zero value otherwise.

### GetWriteSetOk

`func (o *TransactionPayload) GetWriteSetOk() (*WriteSet, bool)`

GetWriteSetOk returns a tuple with the WriteSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteSet

`func (o *TransactionPayload) SetWriteSet(v WriteSet)`

SetWriteSet sets WriteSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


