# MoveModuleABI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Hex-encoded 16 bytes Aptos account address.  Prefixed with &#x60;0x&#x60; and leading zeros are trimmed.  See [doc](https://diem.github.io/move/address.html) for more details.  | 
**Name** | **string** |  | 
**Friends** | **[]string** |  | 
**ExposedFunctions** | [**[]MoveFunction**](MoveFunction.md) |  | 
**Structs** | [**[]MoveStruct**](MoveStruct.md) |  | 

## Methods

### NewMoveModuleABI

`func NewMoveModuleABI(address string, name string, friends []string, exposedFunctions []MoveFunction, structs []MoveStruct, ) *MoveModuleABI`

NewMoveModuleABI instantiates a new MoveModuleABI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveModuleABIWithDefaults

`func NewMoveModuleABIWithDefaults() *MoveModuleABI`

NewMoveModuleABIWithDefaults instantiates a new MoveModuleABI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MoveModuleABI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MoveModuleABI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MoveModuleABI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *MoveModuleABI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveModuleABI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveModuleABI) SetName(v string)`

SetName sets Name field to given value.


### GetFriends

`func (o *MoveModuleABI) GetFriends() []string`

GetFriends returns the Friends field if non-nil, zero value otherwise.

### GetFriendsOk

`func (o *MoveModuleABI) GetFriendsOk() (*[]string, bool)`

GetFriendsOk returns a tuple with the Friends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriends

`func (o *MoveModuleABI) SetFriends(v []string)`

SetFriends sets Friends field to given value.


### GetExposedFunctions

`func (o *MoveModuleABI) GetExposedFunctions() []MoveFunction`

GetExposedFunctions returns the ExposedFunctions field if non-nil, zero value otherwise.

### GetExposedFunctionsOk

`func (o *MoveModuleABI) GetExposedFunctionsOk() (*[]MoveFunction, bool)`

GetExposedFunctionsOk returns a tuple with the ExposedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedFunctions

`func (o *MoveModuleABI) SetExposedFunctions(v []MoveFunction)`

SetExposedFunctions sets ExposedFunctions field to given value.


### GetStructs

`func (o *MoveModuleABI) GetStructs() []MoveStruct`

GetStructs returns the Structs field if non-nil, zero value otherwise.

### GetStructsOk

`func (o *MoveModuleABI) GetStructsOk() (*[]MoveStruct, bool)`

GetStructsOk returns a tuple with the Structs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructs

`func (o *MoveModuleABI) SetStructs(v []MoveStruct)`

SetStructs sets Structs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


