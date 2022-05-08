# TransactionSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PublicKey** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Signature** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**PublicKeys** | **[]string** | all public keys of the sender account | 
**Signatures** | **[]string** | signatures created based on the &#x60;threshold&#x60; | 
**Threshold** | **int32** | The threshold of the multi ed25519 account key. | 
**Bitmap** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Sender** | [**AccountSignature**](AccountSignature.md) |  | 
**SecondarySignerAddresses** | **[]string** |  | 
**SecondarySigners** | [**[]AccountSignature**](AccountSignature.md) |  | 

## Methods

### NewTransactionSignature

`func NewTransactionSignature(type_ string, publicKey string, signature string, publicKeys []string, signatures []string, threshold int32, bitmap string, sender AccountSignature, secondarySignerAddresses []string, secondarySigners []AccountSignature, ) *TransactionSignature`

NewTransactionSignature instantiates a new TransactionSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSignatureWithDefaults

`func NewTransactionSignatureWithDefaults() *TransactionSignature`

NewTransactionSignatureWithDefaults instantiates a new TransactionSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionSignature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSignature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSignature) SetType(v string)`

SetType sets Type field to given value.


### GetPublicKey

`func (o *TransactionSignature) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *TransactionSignature) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *TransactionSignature) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSignature

`func (o *TransactionSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TransactionSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TransactionSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetPublicKeys

`func (o *TransactionSignature) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *TransactionSignature) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *TransactionSignature) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.


### GetSignatures

`func (o *TransactionSignature) GetSignatures() []string`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *TransactionSignature) GetSignaturesOk() (*[]string, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *TransactionSignature) SetSignatures(v []string)`

SetSignatures sets Signatures field to given value.


### GetThreshold

`func (o *TransactionSignature) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TransactionSignature) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TransactionSignature) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetBitmap

`func (o *TransactionSignature) GetBitmap() string`

GetBitmap returns the Bitmap field if non-nil, zero value otherwise.

### GetBitmapOk

`func (o *TransactionSignature) GetBitmapOk() (*string, bool)`

GetBitmapOk returns a tuple with the Bitmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitmap

`func (o *TransactionSignature) SetBitmap(v string)`

SetBitmap sets Bitmap field to given value.


### GetSender

`func (o *TransactionSignature) GetSender() AccountSignature`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TransactionSignature) GetSenderOk() (*AccountSignature, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TransactionSignature) SetSender(v AccountSignature)`

SetSender sets Sender field to given value.


### GetSecondarySignerAddresses

`func (o *TransactionSignature) GetSecondarySignerAddresses() []string`

GetSecondarySignerAddresses returns the SecondarySignerAddresses field if non-nil, zero value otherwise.

### GetSecondarySignerAddressesOk

`func (o *TransactionSignature) GetSecondarySignerAddressesOk() (*[]string, bool)`

GetSecondarySignerAddressesOk returns a tuple with the SecondarySignerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySignerAddresses

`func (o *TransactionSignature) SetSecondarySignerAddresses(v []string)`

SetSecondarySignerAddresses sets SecondarySignerAddresses field to given value.


### GetSecondarySigners

`func (o *TransactionSignature) GetSecondarySigners() []AccountSignature`

GetSecondarySigners returns the SecondarySigners field if non-nil, zero value otherwise.

### GetSecondarySignersOk

`func (o *TransactionSignature) GetSecondarySignersOk() (*[]AccountSignature, bool)`

GetSecondarySignersOk returns a tuple with the SecondarySigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySigners

`func (o *TransactionSignature) SetSecondarySigners(v []AccountSignature)`

SetSecondarySigners sets SecondarySigners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


