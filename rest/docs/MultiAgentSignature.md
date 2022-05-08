# MultiAgentSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Sender** | [**AccountSignature**](AccountSignature.md) |  | 
**SecondarySignerAddresses** | **[]string** |  | 
**SecondarySigners** | [**[]AccountSignature**](AccountSignature.md) |  | 

## Methods

### NewMultiAgentSignature

`func NewMultiAgentSignature(type_ string, sender AccountSignature, secondarySignerAddresses []string, secondarySigners []AccountSignature, ) *MultiAgentSignature`

NewMultiAgentSignature instantiates a new MultiAgentSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiAgentSignatureWithDefaults

`func NewMultiAgentSignatureWithDefaults() *MultiAgentSignature`

NewMultiAgentSignatureWithDefaults instantiates a new MultiAgentSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultiAgentSignature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiAgentSignature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiAgentSignature) SetType(v string)`

SetType sets Type field to given value.


### GetSender

`func (o *MultiAgentSignature) GetSender() AccountSignature`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MultiAgentSignature) GetSenderOk() (*AccountSignature, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MultiAgentSignature) SetSender(v AccountSignature)`

SetSender sets Sender field to given value.


### GetSecondarySignerAddresses

`func (o *MultiAgentSignature) GetSecondarySignerAddresses() []string`

GetSecondarySignerAddresses returns the SecondarySignerAddresses field if non-nil, zero value otherwise.

### GetSecondarySignerAddressesOk

`func (o *MultiAgentSignature) GetSecondarySignerAddressesOk() (*[]string, bool)`

GetSecondarySignerAddressesOk returns a tuple with the SecondarySignerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySignerAddresses

`func (o *MultiAgentSignature) SetSecondarySignerAddresses(v []string)`

SetSecondarySignerAddresses sets SecondarySignerAddresses field to given value.


### GetSecondarySigners

`func (o *MultiAgentSignature) GetSecondarySigners() []AccountSignature`

GetSecondarySigners returns the SecondarySigners field if non-nil, zero value otherwise.

### GetSecondarySignersOk

`func (o *MultiAgentSignature) GetSecondarySignersOk() (*[]AccountSignature, bool)`

GetSecondarySignersOk returns a tuple with the SecondarySigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySigners

`func (o *MultiAgentSignature) SetSecondarySigners(v []AccountSignature)`

SetSecondarySigners sets SecondarySigners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


