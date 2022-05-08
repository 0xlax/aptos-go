# DeleteResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StateKeyHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Address** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Resource** | **string** | String representation of an on-chain Move struct type.  It is a combination of:   1. &#x60;Move module address&#x60;, &#x60;module name&#x60; and &#x60;struct name&#x60; joined by &#x60;::&#x60;.   2. &#x60;struct generic type parameters&#x60; joined by &#x60;, &#x60;.  Examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  See [doc](https://diem.github.io/move/structs-and-resources.html) for more details.  | 

## Methods

### NewDeleteResource

`func NewDeleteResource(type_ string, stateKeyHash string, address string, resource string, ) *DeleteResource`

NewDeleteResource instantiates a new DeleteResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResourceWithDefaults

`func NewDeleteResourceWithDefaults() *DeleteResource`

NewDeleteResourceWithDefaults instantiates a new DeleteResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeleteResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeleteResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeleteResource) SetType(v string)`

SetType sets Type field to given value.


### GetStateKeyHash

`func (o *DeleteResource) GetStateKeyHash() string`

GetStateKeyHash returns the StateKeyHash field if non-nil, zero value otherwise.

### GetStateKeyHashOk

`func (o *DeleteResource) GetStateKeyHashOk() (*string, bool)`

GetStateKeyHashOk returns a tuple with the StateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateKeyHash

`func (o *DeleteResource) SetStateKeyHash(v string)`

SetStateKeyHash sets StateKeyHash field to given value.


### GetAddress

`func (o *DeleteResource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeleteResource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeleteResource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetResource

`func (o *DeleteResource) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DeleteResource) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DeleteResource) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


