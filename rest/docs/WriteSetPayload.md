# WriteSetPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WriteSet** | [**WriteSet**](WriteSet.md) |  | 

## Methods

### NewWriteSetPayload

`func NewWriteSetPayload(type_ string, writeSet WriteSet, ) *WriteSetPayload`

NewWriteSetPayload instantiates a new WriteSetPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteSetPayloadWithDefaults

`func NewWriteSetPayloadWithDefaults() *WriteSetPayload`

NewWriteSetPayloadWithDefaults instantiates a new WriteSetPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WriteSetPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteSetPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteSetPayload) SetType(v string)`

SetType sets Type field to given value.


### GetWriteSet

`func (o *WriteSetPayload) GetWriteSet() WriteSet`

GetWriteSet returns the WriteSet field if non-nil, zero value otherwise.

### GetWriteSetOk

`func (o *WriteSetPayload) GetWriteSetOk() (*WriteSet, bool)`

GetWriteSetOk returns a tuple with the WriteSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteSet

`func (o *WriteSetPayload) SetWriteSet(v WriteSet)`

SetWriteSet sets WriteSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


