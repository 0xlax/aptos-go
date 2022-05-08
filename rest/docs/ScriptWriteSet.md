# ScriptWriteSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExecuteAs** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Script** | [**Script**](Script.md) |  | 

## Methods

### NewScriptWriteSet

`func NewScriptWriteSet(type_ string, executeAs string, script Script, ) *ScriptWriteSet`

NewScriptWriteSet instantiates a new ScriptWriteSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWriteSetWithDefaults

`func NewScriptWriteSetWithDefaults() *ScriptWriteSet`

NewScriptWriteSetWithDefaults instantiates a new ScriptWriteSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScriptWriteSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptWriteSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptWriteSet) SetType(v string)`

SetType sets Type field to given value.


### GetExecuteAs

`func (o *ScriptWriteSet) GetExecuteAs() string`

GetExecuteAs returns the ExecuteAs field if non-nil, zero value otherwise.

### GetExecuteAsOk

`func (o *ScriptWriteSet) GetExecuteAsOk() (*string, bool)`

GetExecuteAsOk returns a tuple with the ExecuteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteAs

`func (o *ScriptWriteSet) SetExecuteAs(v string)`

SetExecuteAs sets ExecuteAs field to given value.


### GetScript

`func (o *ScriptWriteSet) GetScript() Script`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ScriptWriteSet) GetScriptOk() (*Script, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ScriptWriteSet) SetScript(v Script)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


