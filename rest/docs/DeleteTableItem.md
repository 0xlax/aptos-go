# DeleteTableItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StateKeyHash** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Data** | [**TableItemDeletion**](TableItemDeletion.md) |  | 

## Methods

### NewDeleteTableItem

`func NewDeleteTableItem(type_ string, stateKeyHash string, data TableItemDeletion, ) *DeleteTableItem`

NewDeleteTableItem instantiates a new DeleteTableItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTableItemWithDefaults

`func NewDeleteTableItemWithDefaults() *DeleteTableItem`

NewDeleteTableItemWithDefaults instantiates a new DeleteTableItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeleteTableItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeleteTableItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeleteTableItem) SetType(v string)`

SetType sets Type field to given value.


### GetStateKeyHash

`func (o *DeleteTableItem) GetStateKeyHash() string`

GetStateKeyHash returns the StateKeyHash field if non-nil, zero value otherwise.

### GetStateKeyHashOk

`func (o *DeleteTableItem) GetStateKeyHashOk() (*string, bool)`

GetStateKeyHashOk returns a tuple with the StateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateKeyHash

`func (o *DeleteTableItem) SetStateKeyHash(v string)`

SetStateKeyHash sets StateKeyHash field to given value.


### GetData

`func (o *DeleteTableItem) GetData() TableItemDeletion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteTableItem) GetDataOk() (*TableItemDeletion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteTableItem) SetData(v TableItemDeletion)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


