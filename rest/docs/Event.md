# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Event key is a global index for an event stream.  It is hex-encoded BCS bytes of &#x60;EventHandle&#x60; &#x60;guid&#x60; field value, which is a combination of a &#x60;uint64&#x60; creation number and account address (without trimming leading zeros).  For example, event key &#x60;0x00000000000000000000000000000000000000000a550c18&#x60; is combined by the following 2 parts:   1. &#x60;0000000000000000&#x60;: &#x60;uint64&#x60; representation of &#x60;0&#x60;.   2. &#x60;0000000000000000000000000a550c18&#x60;: 16 bytes of account address.  | 
**SequenceNumber** | **string** | Event &#x60;sequence_number&#x60; is unique id of an event in an event stream. Event &#x60;sequence_number&#x60; starts from 0 for each event key.  | 
**Type** | **string** | String representation of an on-chain Move type tag that is exposed in transaction payload.  Values:   - bool   - u8   - u64   - u128   - address   - signer   - vector: &#x60;vector&lt;{non-reference MoveTypeId}&gt;&#x60;   - struct: &#x60;{address}::{module_name}::{struct_name}::&lt;{generic types}&gt;&#x60;  Vector type value examples:   * &#x60;vector&lt;u8&gt;&#x60;   * &#x60;vector&lt;vector&lt;u64&gt;&gt;&#x60;   * &#x60;vector&lt;0x1::AptosAccount::Balance&lt;0x1::XDX::XDX&gt;&gt;&#x60;  Struct type value examples:   * &#x60;0x1::Aptos::Aptos&lt;0x1::XDX::XDX&gt;&#x60;   * &#x60;0x1::Abc::Abc&lt;vector&lt;u8&gt;, vector&lt;u64&gt;&gt;&#x60;   * &#x60;0x1::AptosAccount::AccountOperationsCapability&#x60;  Note:   1. Empty chars should be ignored when comparing 2 struct tag ids.   2. When used in an URL path, should be encoded by url-encoding (AKA percent-encoding).  | 
**Data** | **interface{}** | Move &#x60;bool&#x60; type value is serialized into &#x60;boolean&#x60;.  Move &#x60;u8&#x60; type value is serialized into &#x60;integer&#x60;.  Move &#x60;u64&#x60; and &#x60;u128&#x60; type value is serialized into &#x60;string&#x60;.  Move &#x60;address&#x60; type value(16 bytes Aptos account address) is serialized into hex-encoded string, which is prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  For example:   * &#x60;0x1&#x60;   * &#x60;0x1668f6be25668c1a17cd8caf6b8d2f25&#x60;  Move &#x60;vector&#x60; type value is serialized into &#x60;array&#x60;, except &#x60;vector&lt;u8&gt;&#x60; which is serialized into hex-encoded string with &#x60;0x&#x60; prefix.  For example:   * &#x60;vector&lt;u64&gt;{255, 255}&#x60; &#x3D;&gt; &#x60;[\&quot;255\&quot;, \&quot;255\&quot;]&#x60;   * &#x60;vector&lt;u8&gt;{255, 255}&#x60; &#x3D;&gt; &#x60;0xffff&#x60;  Move &#x60;struct&#x60; type value is serialized into &#x60;object&#x60; that looks like this (except some Move stdlib types, see the following section):    &#x60;&#x60;&#x60;json   {     field1_name: field1_value,     field2_name: field2_value,     ......   }   &#x60;&#x60;&#x60;  For example:   &#x60;{ \&quot;created\&quot;: \&quot;0xa550c18\&quot;, \&quot;role_id\&quot;: \&quot;0\&quot; }&#x60;  **Special serialization for Move stdlib types:**  * [0x1::ASCII::String](https://github.com/aptos-labs/aptos-core/blob/main/language/move-stdlib/docs/ASCII.md) is serialized into &#x60;string&#x60;. For example, struct value &#x60;0x1::ASCII::String{bytes: b\&quot;hello world\&quot;}&#x60; is serialized as &#x60;\&quot;hello world\&quot;&#x60; in JSON.  | 

## Methods

### NewEvent

`func NewEvent(key string, sequenceNumber string, type_ string, data interface{}, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Event) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Event) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Event) SetKey(v string)`

SetKey sets Key field to given value.


### GetSequenceNumber

`func (o *Event) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Event) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Event) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *Event) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Event) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Event) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *Event) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Event) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


