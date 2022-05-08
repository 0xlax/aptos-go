# UserTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Events** | [**[]Event**](Event.md) |  | 
**Timestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 

## Methods

### NewUserTransactionAllOf

`func NewUserTransactionAllOf(type_ string, events []Event, timestamp string, ) *UserTransactionAllOf`

NewUserTransactionAllOf instantiates a new UserTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTransactionAllOfWithDefaults

`func NewUserTransactionAllOfWithDefaults() *UserTransactionAllOf`

NewUserTransactionAllOfWithDefaults instantiates a new UserTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserTransactionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserTransactionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserTransactionAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *UserTransactionAllOf) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UserTransactionAllOf) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UserTransactionAllOf) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetTimestamp

`func (o *UserTransactionAllOf) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserTransactionAllOf) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserTransactionAllOf) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


