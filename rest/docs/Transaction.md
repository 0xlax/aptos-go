# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Hash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Sender** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**SequenceNumber** | **string** | Unsigned int64 type value | 
**MaxGasAmount** | **string** | Unsigned int64 type value | 
**GasUnitPrice** | **string** | Unsigned int64 type value | 
**GasCurrencyCode** | **string** |  | 
**ExpirationTimestampSecs** | **string** | Timestamp in seconds, e.g. transaction expiration timestamp.  | 
**Payload** | [**TransactionPayload**](TransactionPayload.md) |  | 
**Signature** | [**TransactionSignature**](TransactionSignature.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 
**Version** | **string** | Unsigned int64 type value | 
**StateRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**EventRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**GasUsed** | **string** | Unsigned int64 type value | 
**Success** | **bool** | Transaction execution result (success: true, failure: false). See &#x60;vm_status&#x60; for human readable error message from Aptos VM.  | 
**VmStatus** | **string** | Human readable transaction execution result message from Aptos VM.  | 
**AccumulatorRootHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Changes** | [**[]WriteSetChange**](WriteSetChange.md) |  | 
**Timestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 
**Id** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Round** | **string** | Unsigned int64 type value | 
**PreviousBlockVotes** | **[]string** |  | 
**Proposer** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 

## Methods

### NewTransaction

`func NewTransaction(type_ string, hash string, sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature, events []Event, version string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange, timestamp string, id string, round string, previousBlockVotes []string, proposer string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### GetHash

`func (o *Transaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Transaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Transaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSender

`func (o *Transaction) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Transaction) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Transaction) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSequenceNumber

`func (o *Transaction) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Transaction) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Transaction) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetMaxGasAmount

`func (o *Transaction) GetMaxGasAmount() string`

GetMaxGasAmount returns the MaxGasAmount field if non-nil, zero value otherwise.

### GetMaxGasAmountOk

`func (o *Transaction) GetMaxGasAmountOk() (*string, bool)`

GetMaxGasAmountOk returns a tuple with the MaxGasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasAmount

`func (o *Transaction) SetMaxGasAmount(v string)`

SetMaxGasAmount sets MaxGasAmount field to given value.


### GetGasUnitPrice

`func (o *Transaction) GetGasUnitPrice() string`

GetGasUnitPrice returns the GasUnitPrice field if non-nil, zero value otherwise.

### GetGasUnitPriceOk

`func (o *Transaction) GetGasUnitPriceOk() (*string, bool)`

GetGasUnitPriceOk returns a tuple with the GasUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUnitPrice

`func (o *Transaction) SetGasUnitPrice(v string)`

SetGasUnitPrice sets GasUnitPrice field to given value.


### GetGasCurrencyCode

`func (o *Transaction) GetGasCurrencyCode() string`

GetGasCurrencyCode returns the GasCurrencyCode field if non-nil, zero value otherwise.

### GetGasCurrencyCodeOk

`func (o *Transaction) GetGasCurrencyCodeOk() (*string, bool)`

GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCurrencyCode

`func (o *Transaction) SetGasCurrencyCode(v string)`

SetGasCurrencyCode sets GasCurrencyCode field to given value.


### GetExpirationTimestampSecs

`func (o *Transaction) GetExpirationTimestampSecs() string`

GetExpirationTimestampSecs returns the ExpirationTimestampSecs field if non-nil, zero value otherwise.

### GetExpirationTimestampSecsOk

`func (o *Transaction) GetExpirationTimestampSecsOk() (*string, bool)`

GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestampSecs

`func (o *Transaction) SetExpirationTimestampSecs(v string)`

SetExpirationTimestampSecs sets ExpirationTimestampSecs field to given value.


### GetPayload

`func (o *Transaction) GetPayload() TransactionPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Transaction) GetPayloadOk() (*TransactionPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Transaction) SetPayload(v TransactionPayload)`

SetPayload sets Payload field to given value.


### GetSignature

`func (o *Transaction) GetSignature() TransactionSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Transaction) GetSignatureOk() (*TransactionSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Transaction) SetSignature(v TransactionSignature)`

SetSignature sets Signature field to given value.


### GetEvents

`func (o *Transaction) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Transaction) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Transaction) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetVersion

`func (o *Transaction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Transaction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Transaction) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStateRootHash

`func (o *Transaction) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *Transaction) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *Transaction) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetEventRootHash

`func (o *Transaction) GetEventRootHash() string`

GetEventRootHash returns the EventRootHash field if non-nil, zero value otherwise.

### GetEventRootHashOk

`func (o *Transaction) GetEventRootHashOk() (*string, bool)`

GetEventRootHashOk returns a tuple with the EventRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRootHash

`func (o *Transaction) SetEventRootHash(v string)`

SetEventRootHash sets EventRootHash field to given value.


### GetGasUsed

`func (o *Transaction) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *Transaction) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *Transaction) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetSuccess

`func (o *Transaction) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Transaction) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Transaction) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetVmStatus

`func (o *Transaction) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *Transaction) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *Transaction) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.


### GetAccumulatorRootHash

`func (o *Transaction) GetAccumulatorRootHash() string`

GetAccumulatorRootHash returns the AccumulatorRootHash field if non-nil, zero value otherwise.

### GetAccumulatorRootHashOk

`func (o *Transaction) GetAccumulatorRootHashOk() (*string, bool)`

GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatorRootHash

`func (o *Transaction) SetAccumulatorRootHash(v string)`

SetAccumulatorRootHash sets AccumulatorRootHash field to given value.


### GetChanges

`func (o *Transaction) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *Transaction) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *Transaction) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.


### GetTimestamp

`func (o *Transaction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Transaction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Transaction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetRound

`func (o *Transaction) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *Transaction) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *Transaction) SetRound(v string)`

SetRound sets Round field to given value.


### GetPreviousBlockVotes

`func (o *Transaction) GetPreviousBlockVotes() []string`

GetPreviousBlockVotes returns the PreviousBlockVotes field if non-nil, zero value otherwise.

### GetPreviousBlockVotesOk

`func (o *Transaction) GetPreviousBlockVotesOk() (*[]string, bool)`

GetPreviousBlockVotesOk returns a tuple with the PreviousBlockVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockVotes

`func (o *Transaction) SetPreviousBlockVotes(v []string)`

SetPreviousBlockVotes sets PreviousBlockVotes field to given value.


### GetProposer

`func (o *Transaction) GetProposer() string`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *Transaction) GetProposerOk() (*string, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *Transaction) SetProposer(v string)`

SetProposer sets Proposer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


