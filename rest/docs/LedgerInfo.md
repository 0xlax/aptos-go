# LedgerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **int32** | The blockchain chain id.  | 
**LedgerVersion** | **string** | The version of the latest transaction in the ledger.  | 
**LedgerTimestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 

## Methods

### NewLedgerInfo

`func NewLedgerInfo(chainId int32, ledgerVersion string, ledgerTimestamp string, ) *LedgerInfo`

NewLedgerInfo instantiates a new LedgerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerInfoWithDefaults

`func NewLedgerInfoWithDefaults() *LedgerInfo`

NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *LedgerInfo) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *LedgerInfo) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *LedgerInfo) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetLedgerVersion

`func (o *LedgerInfo) GetLedgerVersion() string`

GetLedgerVersion returns the LedgerVersion field if non-nil, zero value otherwise.

### GetLedgerVersionOk

`func (o *LedgerInfo) GetLedgerVersionOk() (*string, bool)`

GetLedgerVersionOk returns a tuple with the LedgerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerVersion

`func (o *LedgerInfo) SetLedgerVersion(v string)`

SetLedgerVersion sets LedgerVersion field to given value.


### GetLedgerTimestamp

`func (o *LedgerInfo) GetLedgerTimestamp() string`

GetLedgerTimestamp returns the LedgerTimestamp field if non-nil, zero value otherwise.

### GetLedgerTimestampOk

`func (o *LedgerInfo) GetLedgerTimestampOk() (*string, bool)`

GetLedgerTimestampOk returns a tuple with the LedgerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerTimestamp

`func (o *LedgerInfo) SetLedgerTimestamp(v string)`

SetLedgerTimestamp sets LedgerTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


