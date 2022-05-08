# GenesisTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Events** | [**[]Event**](Event.md) |  | 
**Payload** | [**WriteSetPayload**](WriteSetPayload.md) |  | 
**Version** | **string** | Unsigned int64 type value | 
**Hash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**StateRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**EventRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**GasUsed** | **string** | Unsigned int64 type value | 
**Success** | **bool** | Transaction execution result (success: true, failure: false). See &#x60;vm_status&#x60; for human readable error message from Aptos VM.  | 
**VmStatus** | **string** | Human readable transaction execution result message from Aptos VM.  | 
**AccumulatorRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Changes** | [**[]WriteSetChange**](WriteSetChange.md) |  | 

## Methods

### NewGenesisTransaction

`func NewGenesisTransaction(type_ string, events []Event, payload WriteSetPayload, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange, ) *GenesisTransaction`

NewGenesisTransaction instantiates a new GenesisTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenesisTransactionWithDefaults

`func NewGenesisTransactionWithDefaults() *GenesisTransaction`

NewGenesisTransactionWithDefaults instantiates a new GenesisTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenesisTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenesisTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenesisTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *GenesisTransaction) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GenesisTransaction) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GenesisTransaction) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetPayload

`func (o *GenesisTransaction) GetPayload() WriteSetPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GenesisTransaction) GetPayloadOk() (*WriteSetPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GenesisTransaction) SetPayload(v WriteSetPayload)`

SetPayload sets Payload field to given value.


### GetVersion

`func (o *GenesisTransaction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GenesisTransaction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GenesisTransaction) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHash

`func (o *GenesisTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GenesisTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GenesisTransaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStateRootHash

`func (o *GenesisTransaction) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *GenesisTransaction) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *GenesisTransaction) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetEventRootHash

`func (o *GenesisTransaction) GetEventRootHash() string`

GetEventRootHash returns the EventRootHash field if non-nil, zero value otherwise.

### GetEventRootHashOk

`func (o *GenesisTransaction) GetEventRootHashOk() (*string, bool)`

GetEventRootHashOk returns a tuple with the EventRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRootHash

`func (o *GenesisTransaction) SetEventRootHash(v string)`

SetEventRootHash sets EventRootHash field to given value.


### GetGasUsed

`func (o *GenesisTransaction) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GenesisTransaction) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GenesisTransaction) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetSuccess

`func (o *GenesisTransaction) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GenesisTransaction) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GenesisTransaction) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetVmStatus

`func (o *GenesisTransaction) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *GenesisTransaction) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *GenesisTransaction) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.


### GetAccumulatorRootHash

`func (o *GenesisTransaction) GetAccumulatorRootHash() string`

GetAccumulatorRootHash returns the AccumulatorRootHash field if non-nil, zero value otherwise.

### GetAccumulatorRootHashOk

`func (o *GenesisTransaction) GetAccumulatorRootHashOk() (*string, bool)`

GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatorRootHash

`func (o *GenesisTransaction) SetAccumulatorRootHash(v string)`

SetAccumulatorRootHash sets AccumulatorRootHash field to given value.


### GetChanges

`func (o *GenesisTransaction) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *GenesisTransaction) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *GenesisTransaction) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


