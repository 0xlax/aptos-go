# BlockMetadataTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Round** | **string** | Unsigned int64 type value | 
**PreviousBlockVotes** | **[]string** |  | 
**Proposer** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Timestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 
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

### NewBlockMetadataTransaction

`func NewBlockMetadataTransaction(type_ string, id string, round string, previousBlockVotes []string, proposer string, timestamp string, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange, ) *BlockMetadataTransaction`

NewBlockMetadataTransaction instantiates a new BlockMetadataTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockMetadataTransactionWithDefaults

`func NewBlockMetadataTransactionWithDefaults() *BlockMetadataTransaction`

NewBlockMetadataTransactionWithDefaults instantiates a new BlockMetadataTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlockMetadataTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlockMetadataTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlockMetadataTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BlockMetadataTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockMetadataTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockMetadataTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetRound

`func (o *BlockMetadataTransaction) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BlockMetadataTransaction) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BlockMetadataTransaction) SetRound(v string)`

SetRound sets Round field to given value.


### GetPreviousBlockVotes

`func (o *BlockMetadataTransaction) GetPreviousBlockVotes() []string`

GetPreviousBlockVotes returns the PreviousBlockVotes field if non-nil, zero value otherwise.

### GetPreviousBlockVotesOk

`func (o *BlockMetadataTransaction) GetPreviousBlockVotesOk() (*[]string, bool)`

GetPreviousBlockVotesOk returns a tuple with the PreviousBlockVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockVotes

`func (o *BlockMetadataTransaction) SetPreviousBlockVotes(v []string)`

SetPreviousBlockVotes sets PreviousBlockVotes field to given value.


### GetProposer

`func (o *BlockMetadataTransaction) GetProposer() string`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *BlockMetadataTransaction) GetProposerOk() (*string, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *BlockMetadataTransaction) SetProposer(v string)`

SetProposer sets Proposer field to given value.


### GetTimestamp

`func (o *BlockMetadataTransaction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockMetadataTransaction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockMetadataTransaction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *BlockMetadataTransaction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockMetadataTransaction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockMetadataTransaction) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHash

`func (o *BlockMetadataTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockMetadataTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockMetadataTransaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStateRootHash

`func (o *BlockMetadataTransaction) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *BlockMetadataTransaction) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *BlockMetadataTransaction) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetEventRootHash

`func (o *BlockMetadataTransaction) GetEventRootHash() string`

GetEventRootHash returns the EventRootHash field if non-nil, zero value otherwise.

### GetEventRootHashOk

`func (o *BlockMetadataTransaction) GetEventRootHashOk() (*string, bool)`

GetEventRootHashOk returns a tuple with the EventRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRootHash

`func (o *BlockMetadataTransaction) SetEventRootHash(v string)`

SetEventRootHash sets EventRootHash field to given value.


### GetGasUsed

`func (o *BlockMetadataTransaction) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *BlockMetadataTransaction) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *BlockMetadataTransaction) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetSuccess

`func (o *BlockMetadataTransaction) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BlockMetadataTransaction) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BlockMetadataTransaction) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetVmStatus

`func (o *BlockMetadataTransaction) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *BlockMetadataTransaction) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *BlockMetadataTransaction) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.


### GetAccumulatorRootHash

`func (o *BlockMetadataTransaction) GetAccumulatorRootHash() string`

GetAccumulatorRootHash returns the AccumulatorRootHash field if non-nil, zero value otherwise.

### GetAccumulatorRootHashOk

`func (o *BlockMetadataTransaction) GetAccumulatorRootHashOk() (*string, bool)`

GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatorRootHash

`func (o *BlockMetadataTransaction) SetAccumulatorRootHash(v string)`

SetAccumulatorRootHash sets AccumulatorRootHash field to given value.


### GetChanges

`func (o *BlockMetadataTransaction) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *BlockMetadataTransaction) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *BlockMetadataTransaction) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


