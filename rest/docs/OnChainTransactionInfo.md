# OnChainTransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewOnChainTransactionInfo

`func NewOnChainTransactionInfo(version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange, ) *OnChainTransactionInfo`

NewOnChainTransactionInfo instantiates a new OnChainTransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnChainTransactionInfoWithDefaults

`func NewOnChainTransactionInfoWithDefaults() *OnChainTransactionInfo`

NewOnChainTransactionInfoWithDefaults instantiates a new OnChainTransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *OnChainTransactionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OnChainTransactionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OnChainTransactionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHash

`func (o *OnChainTransactionInfo) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *OnChainTransactionInfo) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *OnChainTransactionInfo) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStateRootHash

`func (o *OnChainTransactionInfo) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *OnChainTransactionInfo) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *OnChainTransactionInfo) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetEventRootHash

`func (o *OnChainTransactionInfo) GetEventRootHash() string`

GetEventRootHash returns the EventRootHash field if non-nil, zero value otherwise.

### GetEventRootHashOk

`func (o *OnChainTransactionInfo) GetEventRootHashOk() (*string, bool)`

GetEventRootHashOk returns a tuple with the EventRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRootHash

`func (o *OnChainTransactionInfo) SetEventRootHash(v string)`

SetEventRootHash sets EventRootHash field to given value.


### GetGasUsed

`func (o *OnChainTransactionInfo) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *OnChainTransactionInfo) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *OnChainTransactionInfo) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetSuccess

`func (o *OnChainTransactionInfo) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OnChainTransactionInfo) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OnChainTransactionInfo) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetVmStatus

`func (o *OnChainTransactionInfo) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *OnChainTransactionInfo) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *OnChainTransactionInfo) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.


### GetAccumulatorRootHash

`func (o *OnChainTransactionInfo) GetAccumulatorRootHash() string`

GetAccumulatorRootHash returns the AccumulatorRootHash field if non-nil, zero value otherwise.

### GetAccumulatorRootHashOk

`func (o *OnChainTransactionInfo) GetAccumulatorRootHashOk() (*string, bool)`

GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatorRootHash

`func (o *OnChainTransactionInfo) SetAccumulatorRootHash(v string)`

SetAccumulatorRootHash sets AccumulatorRootHash field to given value.


### GetChanges

`func (o *OnChainTransactionInfo) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *OnChainTransactionInfo) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *OnChainTransactionInfo) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


