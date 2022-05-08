# \TableApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTableItem**](TableApi.md#GetTableItem) | **Post** /tables/{table_handle}/item | Get table item by handle and key.



## GetTableItem

> map[string]interface{} GetTableItem(ctx, tableHandle).TableItemRequest(tableItemRequest).Execute()

Get table item by handle and key.



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
    tableHandle := "1283023094380" // string | 
    tableItemRequest := *openapiclient.NewTableItemRequest("0x1::AptosAccount::Balance<0x1::XUS::XUS>", "0x1::AptosAccount::Balance<0x1::XUS::XUS>", interface{}(3344000000)) // TableItemRequest | Table item request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetTableItem(context.Background(), tableHandle).TableItemRequest(tableItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetTableItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableHandle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableItemRequest** | [**TableItemRequest**](TableItemRequest.md) | Table item request | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

