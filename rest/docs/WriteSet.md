# WriteSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExecuteAs** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Script** | [**Script**](Script.md) |  | 
**Changes** | [**[]WriteSetChange**](WriteSetChange.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 

## Methods

### NewWriteSet

`func NewWriteSet(type_ string, executeAs string, script Script, changes []WriteSetChange, events []Event, ) *WriteSet`

NewWriteSet instantiates a new WriteSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteSetWithDefaults

`func NewWriteSetWithDefaults() *WriteSet`

NewWriteSetWithDefaults instantiates a new WriteSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WriteSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteSet) SetType(v string)`

SetType sets Type field to given value.


### GetExecuteAs

`func (o *WriteSet) GetExecuteAs() string`

GetExecuteAs returns the ExecuteAs field if non-nil, zero value otherwise.

### GetExecuteAsOk

`func (o *WriteSet) GetExecuteAsOk() (*string, bool)`

GetExecuteAsOk returns a tuple with the ExecuteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteAs

`func (o *WriteSet) SetExecuteAs(v string)`

SetExecuteAs sets ExecuteAs field to given value.


### GetScript

`func (o *WriteSet) GetScript() Script`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *WriteSet) GetScriptOk() (*Script, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *WriteSet) SetScript(v Script)`

SetScript sets Script field to given value.


### GetChanges

`func (o *WriteSet) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *WriteSet) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *WriteSet) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.


### GetEvents

`func (o *WriteSet) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WriteSet) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WriteSet) SetEvents(v []Event)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


