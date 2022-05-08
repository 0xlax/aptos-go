# WriteTableItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StateKeyHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Data** | [**TableItemWrite**](TableItemWrite.md) |  | 

## Methods

### NewWriteTableItem

`func NewWriteTableItem(type_ string, stateKeyHash string, data TableItemWrite, ) *WriteTableItem`

NewWriteTableItem instantiates a new WriteTableItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteTableItemWithDefaults

`func NewWriteTableItemWithDefaults() *WriteTableItem`

NewWriteTableItemWithDefaults instantiates a new WriteTableItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WriteTableItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteTableItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteTableItem) SetType(v string)`

SetType sets Type field to given value.


### GetStateKeyHash

`func (o *WriteTableItem) GetStateKeyHash() string`

GetStateKeyHash returns the StateKeyHash field if non-nil, zero value otherwise.

### GetStateKeyHashOk

`func (o *WriteTableItem) GetStateKeyHashOk() (*string, bool)`

GetStateKeyHashOk returns a tuple with the StateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateKeyHash

`func (o *WriteTableItem) SetStateKeyHash(v string)`

SetStateKeyHash sets StateKeyHash field to given value.


### GetData

`func (o *WriteTableItem) GetData() TableItemWrite`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WriteTableItem) GetDataOk() (*TableItemWrite, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WriteTableItem) SetData(v TableItemWrite)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


