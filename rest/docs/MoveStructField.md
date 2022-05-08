# MoveStructField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** | String representation of an on-chain Move type identifier defined by the Move language.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: &#x60;vector&lt;{non-reference MoveTypeId}&gt;&#x60;   - struct: &#x60;{address}::{module_name}::{struct_name}::&lt;{generic types}&gt;&#x60;   - reference: immutable &#x60;&amp;&#x60; and mutable &#x60;&amp;mut&#x60; references.   - generic_type_parameter: it is always start with &#x60;T&#x60; and following an index number,     which is the position of the generic type parameter in the &#x60;struct&#x60; or     &#x60;function&#x60; generic type parameters definition.  Vector type value examples:   * &#x60;vector&lt;u8&gt;&#x60;   * &#x60;vector&lt;vector&lt;u64&gt;&gt;&#x60;   * &#x60;vector&lt;0x1::AptosAccount::Balance&lt;0x1::XDX::XDX&gt;&gt;&#x60;  Struct type value examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Reference type value examples:   * &#x60;&amp;signer&#x60;   * &#x60;&amp;mut address&#x60;   * &#x60;&amp;mut vector&lt;u8&gt;&#x60;  Generic type parameter value example, the following is &#x60;0x1::TransactionFee::TransactionFee&#x60; JSON representation:      {         \&quot;name\&quot;: \&quot;TransactionFee\&quot;,         \&quot;is_native\&quot;: false,         \&quot;abilities\&quot;: [\&quot;key\&quot;],         \&quot;generic_type_params\&quot;: [             {\&quot;constraints\&quot;: [], \&quot;is_phantom\&quot;: true}         ],         \&quot;fields\&quot;: [             { \&quot;name\&quot;: \&quot;balance\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Aptos&lt;T0&gt;\&quot; },             { \&quot;name\&quot;: \&quot;preburn\&quot;, \&quot;type\&quot;: \&quot;0x1::Aptos::Preburn&lt;T0&gt;\&quot; }         ]     }  It&#39;s Move source code:      module AptosFramework::TransactionFee {         struct TransactionFee&lt;phantom CoinType&gt; has key {             balance: Aptos&lt;CoinType&gt;,             preburn: Preburn&lt;CoinType&gt;,         }     }  The &#x60;T0&#x60; in the above JSON representation is the generic type place holder for the &#x60;CoinType&#x60; in the Move source code.  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  | 

## Methods

### NewMoveStructField

`func NewMoveStructField(name string, type_ string, ) *MoveStructField`

NewMoveStructField instantiates a new MoveStructField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveStructFieldWithDefaults

`func NewMoveStructFieldWithDefaults() *MoveStructField`

NewMoveStructFieldWithDefaults instantiates a new MoveStructField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MoveStructField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveStructField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveStructField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MoveStructField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MoveStructField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MoveStructField) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


