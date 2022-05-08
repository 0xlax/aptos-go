# AccountResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | String representation of an on-chain Move struct type.  It is a combination of:   1. &#x60;Move module address&#x60;, &#x60;module name&#x60; and &#x60;struct name&#x60; joined by &#x60;::&#x60;.   2. &#x60;struct generic type parameters&#x60; joined by &#x60;, &#x60;.  Examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  See [doc](https://diem.github.io/move/structs-and-resources.html) for more details.  | 
**Data** | **map[string]interface{}** | Account resource data is JSON representation of the Move struct &#x60;type&#x60;.  Move struct field name and value are serialized as object property name and value.  | 

## Methods

### NewAccountResource

`func NewAccountResource(type_ string, data map[string]interface{}, ) *AccountResource`

NewAccountResource instantiates a new AccountResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourceWithDefaults

`func NewAccountResourceWithDefaults() *AccountResource`

NewAccountResourceWithDefaults instantiates a new AccountResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResource) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *AccountResource) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountResource) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountResource) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


