# TableItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | **string** | String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: &#x60;vector&lt;{non-reference MoveTypeId}&gt;&#x60;   - struct: &#x60;{address}::{module_name}::{struct_name}::&lt;{generic types}&gt;&#x60;   - reference: immutable &#x60;&amp;&#x60; and mutable &#x60;&amp;mut&#x60; references.   - generic_type_parameter: it is always start with &#x60;T&#x60; and following an index number,     which is the position of the generic type parameter in the &#x60;struct&#x60; or     &#x60;function&#x60; generic type parameters definition.  Vector type value examples:   * &#x60;vector&lt;u8&gt;&#x60;   * &#x60;vector&lt;vector&lt;u64&gt;&gt;&#x60;   * &#x60;vector&lt;0x1::AptosAccount::Balance&lt;0x1::XDX::XDX&gt;&gt;&#x60;  Struct type value examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Reference type value examples:   * &#x60;&amp;signer&#x60;   * &#x60;&amp;mut address&#x60;   * &#x60;&amp;mut vector&lt;u8&gt;&#x60;  Generic type parameter value example, the following is &#x60;0x1::TransactionFee::TransactionFee&#x60; JSON representation:      {         \&quot;name\&quot;: \&quot;TransactionFee\&quot;,         \&quot;is_native\&quot;: false,         \&quot;abilities\&quot;: [\&quot;key\&quot;],         \&quot;generic_type_params\&quot;: [             {\&quot;constraints\&quot;: [], \&quot;is_phantom\&quot;: true}         ],         \&quot;fields\&quot;: [             { \&quot;name\&quot;: \&quot;balance\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Aptos&lt;T0&gt;\&quot; },             { \&quot;name\&quot;: \&quot;preburn\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Preburn&lt;T0&gt;\&quot; }         ]     }  It&#39;s Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee&lt;phantom CoinType&gt; has key {             balance: Aptos&lt;CoinType&gt;,             preburn: Preburn&lt;CoinType&gt;,         }     }  The &#x60;T0&#x60; in the above JSON representation is the generic type place holder for the &#x60;CoinType&#x60; in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  | 
**ValueType** | **string** | String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: &#x60;vector&lt;{non-reference MoveTypeId}&gt;&#x60;   - struct: &#x60;{address}::{module_name}::{struct_name}::&lt;{generic types}&gt;&#x60;   - reference: immutable &#x60;&amp;&#x60; and mutable &#x60;&amp;mut&#x60; references.   - generic_type_parameter: it is always start with &#x60;T&#x60; and following an index number,     which is the position of the generic type parameter in the &#x60;struct&#x60; or     &#x60;function&#x60; generic type parameters definition.  Vector type value examples:   * &#x60;vector&lt;u8&gt;&#x60;   * &#x60;vector&lt;vector&lt;u64&gt;&gt;&#x60;   * &#x60;vector&lt;0x1::AptosAccount::Balance&lt;0x1::XDX::XDX&gt;&gt;&#x60;  Struct type value examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Reference type value examples:   * &#x60;&amp;signer&#x60;   * &#x60;&amp;mut address&#x60;   * &#x60;&amp;mut vector&lt;u8&gt;&#x60;  Generic type parameter value example, the following is &#x60;0x1::TransactionFee::TransactionFee&#x60; JSON representation:      {         \&quot;name\&quot;: \&quot;TransactionFee\&quot;,         \&quot;is_native\&quot;: false,         \&quot;abilities\&quot;: [\&quot;key\&quot;],         \&quot;generic_type_params\&quot;: [             {\&quot;constraints\&quot;: [], \&quot;is_phantom\&quot;: true}         ],         \&quot;fields\&quot;: [             { \&quot;name\&quot;: \&quot;balance\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Aptos&lt;T0&gt;\&quot; },             { \&quot;name\&quot;: \&quot;preburn\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Preburn&lt;T0&gt;\&quot; }         ]     }  It&#39;s Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee&lt;phantom CoinType&gt; has key {             balance: Aptos&lt;CoinType&gt;,             preburn: Preburn&lt;CoinType&gt;,         }     }  The &#x60;T0&#x60; in the above JSON representation is the generic type place holder for the &#x60;CoinType&#x60; in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  | 
**Key** | **interface{}** | Move &#x60;bool&#x60; type value is serialized into &#x60;boolean&#x60;.  Move &#x60;u8&#x60; type value is serialized into &#x60;integer&#x60;.  Move &#x60;u64&#x60; and &#x60;u128&#x60; type value is serialized into &#x60;string&#x60;.  Move &#x60;address&#x60; type value(16 bytes Aptos account address) is serialized into hex-encoded string, which is prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  For example:   * &#x60;0x1&#x60;   * &#x60;0x1668f6be25668c1a17cd8caf6b8d2f25&#x60;  Move &#x60;vector&#x60; type value is serialized into &#x60;array&#x60;, except &#x60;vector&lt;u8&gt;&#x60; which is serialized into hex-encoded string with &#x60;0x&#x60; prefix.  For example:   * &#x60;vector&lt;u64&gt;{255, 255}&#x60; &#x3D;&gt; &#x60;[\&quot;255\&quot;, \&quot;255\&quot;]&#x60;   * &#x60;vector&lt;u8&gt;{255, 255}&#x60; &#x3D;&gt; &#x60;0xffff&#x60;  Move &#x60;struct&#x60; type value is serialized into &#x60;object&#x60; that looks like this (except some Move stdlib types, see the following section):    &#x60;&#x60;&#x60;json   {     field1_name: field1_value,     field2_name: field2_value,     ......   }   &#x60;&#x60;&#x60;  For example:   &#x60;{ \&quot;created\&quot;: \&quot;0xa550c18\&quot;, \&quot;role_id\&quot;: \&quot;0\&quot; }&#x60;  **Special serialization for Move stdlib types:**  * [0x1::ASCII::String](https://github.com/aptos-labs/aptos-core/blob/main/language/move-stdlib/docs/ASCII.md) is serialized into &#x60;string&#x60;. For example, struct value &#x60;0x1::ASCII::String{bytes: b\&quot;hello world\&quot;}&#x60; is serialized as &#x60;\&quot;hello world\&quot;&#x60; in JSON.  | 

## Methods

### NewTableItemRequest

`func NewTableItemRequest(keyType string, valueType string, key interface{}, ) *TableItemRequest`

NewTableItemRequest instantiates a new TableItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableItemRequestWithDefaults

`func NewTableItemRequestWithDefaults() *TableItemRequest`

NewTableItemRequestWithDefaults instantiates a new TableItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *TableItemRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TableItemRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TableItemRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetValueType

`func (o *TableItemRequest) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *TableItemRequest) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *TableItemRequest) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetKey

`func (o *TableItemRequest) GetKey() interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TableItemRequest) GetKeyOk() (*interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TableItemRequest) SetKey(v interface{})`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *TableItemRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *TableItemRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


