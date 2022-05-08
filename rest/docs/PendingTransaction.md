# PendingTransaction

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

## Methods

### NewPendingTransaction

`func NewPendingTransaction(type_ string, hash string, sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature, ) *PendingTransaction`

NewPendingTransaction instantiates a new PendingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingTransactionWithDefaults

`func NewPendingTransactionWithDefaults() *PendingTransaction`

NewPendingTransactionWithDefaults instantiates a new PendingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PendingTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetHash

`func (o *PendingTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PendingTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PendingTransaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSender

`func (o *PendingTransaction) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PendingTransaction) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PendingTransaction) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSequenceNumber

`func (o *PendingTransaction) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *PendingTransaction) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *PendingTransaction) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetMaxGasAmount

`func (o *PendingTransaction) GetMaxGasAmount() string`

GetMaxGasAmount returns the MaxGasAmount field if non-nil, zero value otherwise.

### GetMaxGasAmountOk

`func (o *PendingTransaction) GetMaxGasAmountOk() (*string, bool)`

GetMaxGasAmountOk returns a tuple with the MaxGasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasAmount

`func (o *PendingTransaction) SetMaxGasAmount(v string)`

SetMaxGasAmount sets MaxGasAmount field to given value.


### GetGasUnitPrice

`func (o *PendingTransaction) GetGasUnitPrice() string`

GetGasUnitPrice returns the GasUnitPrice field if non-nil, zero value otherwise.

### GetGasUnitPriceOk

`func (o *PendingTransaction) GetGasUnitPriceOk() (*string, bool)`

GetGasUnitPriceOk returns a tuple with the GasUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUnitPrice

`func (o *PendingTransaction) SetGasUnitPrice(v string)`

SetGasUnitPrice sets GasUnitPrice field to given value.


### GetGasCurrencyCode

`func (o *PendingTransaction) GetGasCurrencyCode() string`

GetGasCurrencyCode returns the GasCurrencyCode field if non-nil, zero value otherwise.

### GetGasCurrencyCodeOk

`func (o *PendingTransaction) GetGasCurrencyCodeOk() (*string, bool)`

GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCurrencyCode

`func (o *PendingTransaction) SetGasCurrencyCode(v string)`

SetGasCurrencyCode sets GasCurrencyCode field to given value.


### GetExpirationTimestampSecs

`func (o *PendingTransaction) GetExpirationTimestampSecs() string`

GetExpirationTimestampSecs returns the ExpirationTimestampSecs field if non-nil, zero value otherwise.

### GetExpirationTimestampSecsOk

`func (o *PendingTransaction) GetExpirationTimestampSecsOk() (*string, bool)`

GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestampSecs

`func (o *PendingTransaction) SetExpirationTimestampSecs(v string)`

SetExpirationTimestampSecs sets ExpirationTimestampSecs field to given value.


### GetPayload

`func (o *PendingTransaction) GetPayload() TransactionPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PendingTransaction) GetPayloadOk() (*TransactionPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PendingTransaction) SetPayload(v TransactionPayload)`

SetPayload sets Payload field to given value.


### GetSignature

`func (o *PendingTransaction) GetSignature() TransactionSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PendingTransaction) GetSignatureOk() (*TransactionSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PendingTransaction) SetSignature(v TransactionSignature)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


