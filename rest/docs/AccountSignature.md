# AccountSignature

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

## Methods

### NewAccountSignature

`func NewAccountSignature(type_ string, publicKey string, signature string, publicKeys []string, signatures []string, threshold int32, bitmap string, ) *AccountSignature`

NewAccountSignature instantiates a new AccountSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSignatureWithDefaults

`func NewAccountSignatureWithDefaults() *AccountSignature`

NewAccountSignatureWithDefaults instantiates a new AccountSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountSignature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountSignature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountSignature) SetType(v string)`

SetType sets Type field to given value.


### GetPublicKey

`func (o *AccountSignature) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AccountSignature) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AccountSignature) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSignature

`func (o *AccountSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AccountSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AccountSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetPublicKeys

`func (o *AccountSignature) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *AccountSignature) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *AccountSignature) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.


### GetSignatures

`func (o *AccountSignature) GetSignatures() []string`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *AccountSignature) GetSignaturesOk() (*[]string, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *AccountSignature) SetSignatures(v []string)`

SetSignatures sets Signatures field to given value.


### GetThreshold

`func (o *AccountSignature) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AccountSignature) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AccountSignature) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetBitmap

`func (o *AccountSignature) GetBitmap() string`

GetBitmap returns the Bitmap field if non-nil, zero value otherwise.

### GetBitmapOk

`func (o *AccountSignature) GetBitmapOk() (*string, bool)`

GetBitmapOk returns a tuple with the Bitmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitmap

`func (o *AccountSignature) SetBitmap(v string)`

SetBitmap sets Bitmap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


