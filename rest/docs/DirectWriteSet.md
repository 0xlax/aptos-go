# DirectWriteSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Changes** | [**[]WriteSetChange**](WriteSetChange.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 

## Methods

### NewDirectWriteSet

`func NewDirectWriteSet(type_ string, changes []WriteSetChange, events []Event, ) *DirectWriteSet`

NewDirectWriteSet instantiates a new DirectWriteSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectWriteSetWithDefaults

`func NewDirectWriteSetWithDefaults() *DirectWriteSet`

NewDirectWriteSetWithDefaults instantiates a new DirectWriteSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DirectWriteSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectWriteSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectWriteSet) SetType(v string)`

SetType sets Type field to given value.


### GetChanges

`func (o *DirectWriteSet) GetChanges() []WriteSetChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *DirectWriteSet) GetChangesOk() (*[]WriteSetChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *DirectWriteSet) SetChanges(v []WriteSetChange)`

SetChanges sets Changes field to given value.


### GetEvents

`func (o *DirectWriteSet) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DirectWriteSet) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DirectWriteSet) SetEvents(v []Event)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


