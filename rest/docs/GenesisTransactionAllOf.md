# GenesisTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Events** | [**[]Event**](Event.md) |  | 
**Payload** | [**WriteSetPayload**](WriteSetPayload.md) |  | 

## Methods

### NewGenesisTransactionAllOf

`func NewGenesisTransactionAllOf(type_ string, events []Event, payload WriteSetPayload, ) *GenesisTransactionAllOf`

NewGenesisTransactionAllOf instantiates a new GenesisTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenesisTransactionAllOfWithDefaults

`func NewGenesisTransactionAllOfWithDefaults() *GenesisTransactionAllOf`

NewGenesisTransactionAllOfWithDefaults instantiates a new GenesisTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenesisTransactionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenesisTransactionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenesisTransactionAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *GenesisTransactionAllOf) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GenesisTransactionAllOf) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GenesisTransactionAllOf) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetPayload

`func (o *GenesisTransactionAllOf) GetPayload() WriteSetPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GenesisTransactionAllOf) GetPayloadOk() (*WriteSetPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GenesisTransactionAllOf) SetPayload(v WriteSetPayload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


