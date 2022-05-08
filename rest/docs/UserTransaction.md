# UserTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Events** | [**[]Event**](Event.md) |  | 
**Timestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 
**Sender** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**SequenceNumber** | **string** | Unsigned int64 type value | 
**MaxGasAmount** | **string** | Unsigned int64 type value | 
**GasUnitPrice** | **string** | Unsigned int64 type value | 
**GasCurrencyCode** | **string** |  | 
**ExpirationTimestampSecs** | **string** | Timestamp in seconds, e.g. transaction expiration timestamp.  | 
**Payload** | [**TransactionPayload**](TransactionPayload.md) |  | 
**Signature** | [**TransactionSignature**](TransactionSignature.md) |  | 
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

### NewUserTransaction

`func NewUserTransaction(type_ string, events []Event, timestamp string, sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature, version string, hash string, stateRootHash string, eventRootHash string, gasUsed string, success bool, vmStatus string, accumulatorRootHash string, changes []WriteSetChange, ) *UserTransaction`

NewUserTransaction instantiates a new UserTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTransactionWithDefaults

`func NewUserTransactionWithDefaults() *UserTransaction`

NewUserTransactionWithDefaults instantiates a new UserTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *UserTransaction) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UserTransaction) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UserTransaction) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetTimestamp

`func (o *UserTransaction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserTransaction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserTransaction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetSender

`func (o *UserTransaction) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *UserTransaction) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *UserTransaction) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSequenceNumber

`func (o *UserTransaction) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *UserTransaction) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *UserTransaction) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetMaxGasAmount

`func (o *UserTransaction) GetMaxGasAmount() string`

GetMaxGasAmount returns the MaxGasAmount field if non-nil, zero value otherwise.

### GetMaxGasAmountOk

`func (o *UserTransaction) GetMaxGasAmountOk() (*string, bool)`

GetMaxGasAmountOk returns a tuple with the MaxGasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasAmount

`func (o *UserTransaction) SetMaxGasAmount(v string)`

SetMaxGasAmount sets MaxGasAmount field to given value.


### GetGasUnitPrice

`func (o *UserTransaction) GetGasUnitPrice() string`

GetGasUnitPrice returns the GasUnitPrice field if non-nil, zero value otherwise.

### GetGasUnitPriceOk

`func (o *UserTransaction) GetGasUnitPriceOk() (*string, bool)`

GetGasUnitPriceOk returns a tuple with the GasUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUnitPrice

`func (o *UserTransaction) SetGasUnitPrice(v string)`

SetGasUnitPrice sets GasUnitPrice field to given value.


### GetGasCurrencyCode

`func (o *UserTransaction) GetGasCurrencyCode() string`

GetGasCurrencyCode returns the GasCurrencyCode field if non-nil, zero value otherwise.

### GetGasCurrencyCodeOk

`func (o *UserTransaction) GetGasCurrencyCodeOk() (*string, bool)`

GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCurrencyCode

`func (o *UserTransaction) SetGasCurrencyCode(v string)`

SetGasCurrencyCode sets GasCurrencyCode field to given value.


### GetExpirationTimestampSecs

`func (o *UserTransaction) GetExpirationTimestampSecs() string`

GetExpirationTimestampSecs returns the ExpirationTimestampSecs field if non-nil, zero value otherwise.

### GetExpirationTimestampSecsOk

`func (o *UserTransaction) GetExpirationTimestampSecsOk() (*string, bool)`

GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestampSecs

`func (o *UserTransaction) SetExpirationTimestampSecs(v string)`

SetExpirationTimestampSecs sets ExpirationTimestampSecs field to given value.


### GetPayload

`func (o *UserTransaction) GetPayload() TransactionPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UserTransaction) GetPayloadOk() (*TransactionPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UserTransaction) SetPayload(v TransactionPayload)`

SetPayload sets Payload field to given value.


### GetSignature

`func (o *UserTransaction) GetSignature() TransactionSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserTransaction) GetSignatureOk() (*TransactionSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserTransaction) SetSignature(v TransactionSignature)`

SetSignature sets Signature field to given value.


### GetVersion

`func (o *UserTransaction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserTransaction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserTransaction) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHash

`func (o *UserTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UserTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UserTransaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStateRootHash

`func (o *UserTransaction) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *UserTransaction) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *UserTransaction) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetEventRootHash

`func (o *UserTransaction) GetEventRootHash() string`

GetEventRootHash returns the EventRootHash field if non-nil, zero value otherwise.

### GetEventRootHashOk

`func (o *UserTransaction) GetEventRootHashOk() (*string, bool)`

GetEventRootHashOk returns a tuple with the EventRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRootHash

`func (o *UserTransaction) SetEventRootHash(v string)`

SetEventRootHash sets EventRootHash field to given value.


### GetGasUsed

`func (o *UserTransaction) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *UserTransaction) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *UserTransaction) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetSuccess

`func (o *UserTransaction) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UserTransaction) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UserTransaction) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetVmStatus

`func (o *UserTransaction) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *UserTransaction) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *UserTransaction) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.


### GetAccumulatorRootHash

`func (o *UserTransaction) GetAccumulatorRootHash() string`

GetAccumulatorRootHash returns the AccumulatorRootHash field if non-nil, zero value otherwise.

### GetAccumulatorRootHashOk

`func (o *UserTransaction) GetAccumulatorRootHashOk() (*string, bool)`

GetAccumulatorRootHashOk returns a tuple with the AccumulatorRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatorRootHash

`func (o *UserTransaction) SetAccumulatorRootHash(v string)`

SetAccumulatorRootHash sets AccumulatorRootHash field to given value.


### GetChanges

`func (o *UserTransaction) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *UserTransaction) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *UserTransaction) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


