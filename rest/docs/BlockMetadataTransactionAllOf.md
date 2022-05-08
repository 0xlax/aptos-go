# BlockMetadataTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** | All bytes data are represented as hex-encoded string prefixed with &#x60;0x&#x60; and fulfilled with two hex digits per byte.  Different with &#x60;Address&#x60; type, hex-encoded bytes should not trim any zeros.  | 
**Round** | **string** | Unsigned int64 type value | 
**PreviousBlockVotes** | **[]string** |  | 
**Proposer** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Timestamp** | **string** | Timestamp in microseconds, e.g. ledger / block creation timestamp.  | 

## Methods

### NewBlockMetadataTransactionAllOf

`func NewBlockMetadataTransactionAllOf(type_ string, id string, round string, previousBlockVotes []string, proposer string, timestamp string, ) *BlockMetadataTransactionAllOf`

NewBlockMetadataTransactionAllOf instantiates a new BlockMetadataTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockMetadataTransactionAllOfWithDefaults

`func NewBlockMetadataTransactionAllOfWithDefaults() *BlockMetadataTransactionAllOf`

NewBlockMetadataTransactionAllOfWithDefaults instantiates a new BlockMetadataTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlockMetadataTransactionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlockMetadataTransactionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlockMetadataTransactionAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BlockMetadataTransactionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockMetadataTransactionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockMetadataTransactionAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetRound

`func (o *BlockMetadataTransactionAllOf) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BlockMetadataTransactionAllOf) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BlockMetadataTransactionAllOf) SetRound(v string)`

SetRound sets Round field to given value.


### GetPreviousBlockVotes

`func (o *BlockMetadataTransactionAllOf) GetPreviousBlockVotes() []string`

GetPreviousBlockVotes returns the PreviousBlockVotes field if non-nil, zero value otherwise.

### GetPreviousBlockVotesOk

`func (o *BlockMetadataTransactionAllOf) GetPreviousBlockVotesOk() (*[]string, bool)`

GetPreviousBlockVotesOk returns a tuple with the PreviousBlockVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockVotes

`func (o *BlockMetadataTransactionAllOf) SetPreviousBlockVotes(v []string)`

SetPreviousBlockVotes sets PreviousBlockVotes field to given value.


### GetProposer

`func (o *BlockMetadataTransactionAllOf) GetProposer() string`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *BlockMetadataTransactionAllOf) GetProposerOk() (*string, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *BlockMetadataTransactionAllOf) SetProposer(v string)`

SetProposer sets Proposer field to given value.


### GetTimestamp

`func (o *BlockMetadataTransactionAllOf) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockMetadataTransactionAllOf) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockMetadataTransactionAllOf) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


