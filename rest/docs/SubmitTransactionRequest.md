# SubmitTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**SequenceNumber** | **string** | Unsigned int64 type value | 
**MaxGasAmount** | **string** | Unsigned int64 type value | 
**GasUnitPrice** | **string** | Unsigned int64 type value | 
**GasCurrencyCode** | **string** |  | 
**ExpirationTimestampSecs** | **string** | Timestamp in seconds, e.g. transaction expiration timestamp.  | 
**Payload** | [**TransactionPayload**](TransactionPayload.md) |  | 
**Signature** | [**TransactionSignature**](TransactionSignature.md) |  | 

## Methods

### NewSubmitTransactionRequest

`func NewSubmitTransactionRequest(sender string, sequenceNumber string, maxGasAmount string, gasUnitPrice string, gasCurrencyCode string, expirationTimestampSecs string, payload TransactionPayload, signature TransactionSignature, ) *SubmitTransactionRequest`

NewSubmitTransactionRequest instantiates a new SubmitTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTransactionRequestWithDefaults

`func NewSubmitTransactionRequestWithDefaults() *SubmitTransactionRequest`

NewSubmitTransactionRequestWithDefaults instantiates a new SubmitTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *SubmitTransactionRequest) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SubmitTransactionRequest) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SubmitTransactionRequest) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSequenceNumber

`func (o *SubmitTransactionRequest) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *SubmitTransactionRequest) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *SubmitTransactionRequest) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetMaxGasAmount

`func (o *SubmitTransactionRequest) GetMaxGasAmount() string`

GetMaxGasAmount returns the MaxGasAmount field if non-nil, zero value otherwise.

### GetMaxGasAmountOk

`func (o *SubmitTransactionRequest) GetMaxGasAmountOk() (*string, bool)`

GetMaxGasAmountOk returns a tuple with the MaxGasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasAmount

`func (o *SubmitTransactionRequest) SetMaxGasAmount(v string)`

SetMaxGasAmount sets MaxGasAmount field to given value.


### GetGasUnitPrice

`func (o *SubmitTransactionRequest) GetGasUnitPrice() string`

GetGasUnitPrice returns the GasUnitPrice field if non-nil, zero value otherwise.

### GetGasUnitPriceOk

`func (o *SubmitTransactionRequest) GetGasUnitPriceOk() (*string, bool)`

GetGasUnitPriceOk returns a tuple with the GasUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUnitPrice

`func (o *SubmitTransactionRequest) SetGasUnitPrice(v string)`

SetGasUnitPrice sets GasUnitPrice field to given value.


### GetGasCurrencyCode

`func (o *SubmitTransactionRequest) GetGasCurrencyCode() string`

GetGasCurrencyCode returns the GasCurrencyCode field if non-nil, zero value otherwise.

### GetGasCurrencyCodeOk

`func (o *SubmitTransactionRequest) GetGasCurrencyCodeOk() (*string, bool)`

GetGasCurrencyCodeOk returns a tuple with the GasCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCurrencyCode

`func (o *SubmitTransactionRequest) SetGasCurrencyCode(v string)`

SetGasCurrencyCode sets GasCurrencyCode field to given value.


### GetExpirationTimestampSecs

`func (o *SubmitTransactionRequest) GetExpirationTimestampSecs() string`

GetExpirationTimestampSecs returns the ExpirationTimestampSecs field if non-nil, zero value otherwise.

### GetExpirationTimestampSecsOk

`func (o *SubmitTransactionRequest) GetExpirationTimestampSecsOk() (*string, bool)`

GetExpirationTimestampSecsOk returns a tuple with the ExpirationTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestampSecs

`func (o *SubmitTransactionRequest) SetExpirationTimestampSecs(v string)`

SetExpirationTimestampSecs sets ExpirationTimestampSecs field to given value.


### GetPayload

`func (o *SubmitTransactionRequest) GetPayload() TransactionPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SubmitTransactionRequest) GetPayloadOk() (*TransactionPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SubmitTransactionRequest) SetPayload(v TransactionPayload)`

SetPayload sets Payload field to given value.


### GetSignature

`func (o *SubmitTransactionRequest) GetSignature() TransactionSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SubmitTransactionRequest) GetSignatureOk() (*TransactionSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SubmitTransactionRequest) SetSignature(v TransactionSignature)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


