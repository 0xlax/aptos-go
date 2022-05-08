# Ed25519Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PublicKey** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Signature** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 

## Methods

### NewEd25519Signature

`func NewEd25519Signature(type_ string, publicKey string, signature string, ) *Ed25519Signature`

NewEd25519Signature instantiates a new Ed25519Signature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEd25519SignatureWithDefaults

`func NewEd25519SignatureWithDefaults() *Ed25519Signature`

NewEd25519SignatureWithDefaults instantiates a new Ed25519Signature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Ed25519Signature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ed25519Signature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ed25519Signature) SetType(v string)`

SetType sets Type field to given value.


### GetPublicKey

`func (o *Ed25519Signature) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Ed25519Signature) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Ed25519Signature) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSignature

`func (o *Ed25519Signature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Ed25519Signature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Ed25519Signature) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


