# \EventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsByEventHandle**](EventsApi.md#GetEventsByEventHandle) | **Get** /accounts/{address}/events/{event_handle_struct}/{field_name} | Get events by event handle
[**GetEventsByEventKey**](EventsApi.md#GetEventsByEventKey) | **Get** /events/{event_key} | Get events by event key



## GetEventsByEventHandle

> []Event GetEventsByEventHandle(ctx, address, eventHandleStruct, fieldName).Start(start).Limit(limit).Execute()

Get events by event handle



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    address := "address_example" // string | 
    eventHandleStruct := "0x1::AptosAccount::AptosAccount" // string | 
    fieldName := "sent_events" // string | The field name of the `EventHandle` in the struct. 
    start := int32(56) // int32 | The start sequence number in the EVENT STREAM, defaulting to the latest event. The events are returned in the reverse order of sequence numbers.  (optional)
    limit := int32(25) // int32 | The number of events to be returned for the page default is 5 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsByEventHandle(context.Background(), address, eventHandleStruct, fieldName).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsByEventHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsByEventHandle`: []Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsByEventHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 
**eventHandleStruct** | **string** |  | 
**fieldName** | **string** | The field name of the &#x60;EventHandle&#x60; in the struct.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsByEventHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int32** | The start sequence number in the EVENT STREAM, defaulting to the latest event. The events are returned in the reverse order of sequence numbers.  | 
 **limit** | **int32** | The number of events to be returned for the page default is 5 | 

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByEventKey

> []Event GetEventsByEventKey(ctx, eventKey).Execute()

Get events by event key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventKey := "eventKey_example" // string | Event key for an event stream. It is BCS serialized bytes of `guid` field in the Move struct `EventHandle`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsByEventKey(context.Background(), eventKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsByEventKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsByEventKey`: []Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsByEventKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string** | Event key for an event stream. It is BCS serialized bytes of &#x60;guid&#x60; field in the Move struct &#x60;EventHandle&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsByEventKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

