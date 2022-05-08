# \StateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](StateApi.md#GetAccount) | **Get** /accounts/{address} | Get account
[**GetAccountModule**](StateApi.md#GetAccountModule) | **Get** /accounts/{address}/module/{module_name} | Get module by module id.
[**GetAccountModules**](StateApi.md#GetAccountModules) | **Get** /accounts/{address}/modules | Get account modules
[**GetAccountResource**](StateApi.md#GetAccountResource) | **Get** /accounts/{address}/resource/{resource_type} | Get resource by account address and resource type.
[**GetAccountResources**](StateApi.md#GetAccountResources) | **Get** /accounts/{address}/resources | Get account resources
[**GetTableItem**](StateApi.md#GetTableItem) | **Post** /tables/{table_handle}/item | Get table item by handle and key.



## GetAccount

> Account GetAccount(ctx, address).Execute()

Get account

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.GetAccount(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountModule

> MoveModule GetAccountModule(ctx, address, moduleName).Version(version).Execute()

Get module by module id.



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
    moduleName := "GUID" // string | The name of the module.
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.GetAccountModule(context.Background(), address, moduleName).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetAccountModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountModule`: MoveModule
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetAccountModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 
**moduleName** | **string** | The name of the module. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** |  | 

### Return type

[**MoveModule**](MoveModule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountModules

> []MoveModule GetAccountModules(ctx, address).Version(version).Execute()

Get account modules

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
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.GetAccountModules(context.Background(), address).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetAccountModules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountModules`: []MoveModule
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetAccountModules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountModulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** |  | 

### Return type

[**[]MoveModule**](MoveModule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountResource

> AccountResource GetAccountResource(ctx, address, resourceType).Version(version).Execute()

Get resource by account address and resource type.



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
    resourceType := "0x1::AptosAccount::AptosAccount" // string | 
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.GetAccountResource(context.Background(), address, resourceType).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetAccountResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountResource`: AccountResource
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetAccountResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** |  | 

### Return type

[**AccountResource**](AccountResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountResources

> []AccountResource GetAccountResources(ctx, address).Version(version).Execute()

Get account resources

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
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.GetAccountResources(context.Background(), address).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetAccountResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountResources`: []AccountResource
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetAccountResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** |  | 

### Return type

[**[]AccountResource**](AccountResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.StateApi.GetTableItem(context.Background(), tableHandle).TableItemRequest(tableItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.GetTableItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StateApi.GetTableItem`: %v\n", resp)
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

