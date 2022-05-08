# ModuleBundlePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Modules** | [**[]MoveModule**](MoveModule.md) |  | 

## Methods

### NewModuleBundlePayload

`func NewModuleBundlePayload(type_ string, modules []MoveModule, ) *ModuleBundlePayload`

NewModuleBundlePayload instantiates a new ModuleBundlePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleBundlePayloadWithDefaults

`func NewModuleBundlePayloadWithDefaults() *ModuleBundlePayload`

NewModuleBundlePayloadWithDefaults instantiates a new ModuleBundlePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModuleBundlePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModuleBundlePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModuleBundlePayload) SetType(v string)`

SetType sets Type field to given value.


### GetModules

`func (o *ModuleBundlePayload) GetModules() []MoveModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ModuleBundlePayload) GetModulesOk() (*[]MoveModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ModuleBundlePayload) SetModules(v []MoveModule)`

SetModules sets Modules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


