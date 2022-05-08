# MultiEd25519Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PublicKeys** | **[]string** | all public keys of the sender account | 
**Signatures** | **[]string** | signatures created based on the &#x60;threshold&#x60; | 
**Threshold** | **int32** | The threshold of the multi ed25519 account key. | 
**Bitmap** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 

## Methods

### NewMultiEd25519Signature

`func NewMultiEd25519Signature(type_ string, publicKeys []string, signatures []string, threshold int32, bitmap string, ) *MultiEd25519Signature`

NewMultiEd25519Signature instantiates a new MultiEd25519Signature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEd25519SignatureWithDefaults

`func NewMultiEd25519SignatureWithDefaults() *MultiEd25519Signature`

NewMultiEd25519SignatureWithDefaults instantiates a new MultiEd25519Signature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultiEd25519Signature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiEd25519Signature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiEd25519Signature) SetType(v string)`

SetType sets Type field to given value.


### GetPublicKeys

`func (o *MultiEd25519Signature) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *MultiEd25519Signature) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *MultiEd25519Signature) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.


### GetSignatures

`func (o *MultiEd25519Signature) GetSignatures() []string`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *MultiEd25519Signature) GetSignaturesOk() (*[]string, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *MultiEd25519Signature) SetSignatures(v []string)`

SetSignatures sets Signatures field to given value.


### GetThreshold

`func (o *MultiEd25519Signature) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MultiEd25519Signature) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MultiEd25519Signature) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetBitmap

`func (o *MultiEd25519Signature) GetBitmap() string`

GetBitmap returns the Bitmap field if non-nil, zero value otherwise.

### GetBitmapOk

`func (o *MultiEd25519Signature) GetBitmapOk() (*string, bool)`

GetBitmapOk returns a tuple with the Bitmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitmap

`func (o *MultiEd25519Signature) SetBitmap(v string)`

SetBitmap sets Bitmap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


