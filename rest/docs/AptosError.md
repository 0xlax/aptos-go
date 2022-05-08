# AptosError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | **string** |  | 
**AptosLedgerVersion** | Pointer to **string** | The version of the latest transaction in the ledger.  | [optional] 

## Methods

### NewAptosError

`func NewAptosError(code int32, message string, ) *AptosError`

NewAptosError instantiates a new AptosError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAptosErrorWithDefaults

`func NewAptosErrorWithDefaults() *AptosError`

NewAptosErrorWithDefaults instantiates a new AptosError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AptosError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AptosError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AptosError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *AptosError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AptosError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AptosError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAptosLedgerVersion

`func (o *AptosError) GetAptosLedgerVersion() string`

GetAptosLedgerVersion returns the AptosLedgerVersion field if non-nil, zero value otherwise.

### GetAptosLedgerVersionOk

`func (o *AptosError) GetAptosLedgerVersionOk() (*string, bool)`

GetAptosLedgerVersionOk returns a tuple with the AptosLedgerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptosLedgerVersion

`func (o *AptosError) SetAptosLedgerVersion(v string)`

SetAptosLedgerVersion sets AptosLedgerVersion field to given value.

### HasAptosLedgerVersion

`func (o *AptosError) HasAptosLedgerVersion() bool`

HasAptosLedgerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


