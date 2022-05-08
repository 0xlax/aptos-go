# DeleteModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StateKeyHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Address** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Module** | **string** | Move module id is a string representation of Move module.  Format: \&quot;{address}::{module name}\&quot;  &#x60;address&#x60; should be hex-encoded 16 bytes account address that is prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  Module name is case-sensitive.  See [doc](https://diem.github.io/move/modules-and-scripts.html#modules) for more details.  | 

## Methods

### NewDeleteModule

`func NewDeleteModule(type_ string, stateKeyHash string, address string, module string, ) *DeleteModule`

NewDeleteModule instantiates a new DeleteModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteModuleWithDefaults

`func NewDeleteModuleWithDefaults() *DeleteModule`

NewDeleteModuleWithDefaults instantiates a new DeleteModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeleteModule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeleteModule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeleteModule) SetType(v string)`

SetType sets Type field to given value.


### GetStateKeyHash

`func (o *DeleteModule) GetStateKeyHash() string`

GetStateKeyHash returns the StateKeyHash field if non-nil, zero value otherwise.

### GetStateKeyHashOk

`func (o *DeleteModule) GetStateKeyHashOk() (*string, bool)`

GetStateKeyHashOk returns a tuple with the StateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateKeyHash

`func (o *DeleteModule) SetStateKeyHash(v string)`

SetStateKeyHash sets StateKeyHash field to given value.


### GetAddress

`func (o *DeleteModule) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeleteModule) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeleteModule) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetModule

`func (o *DeleteModule) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DeleteModule) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DeleteModule) SetModule(v string)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


