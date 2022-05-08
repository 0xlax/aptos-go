# \TransactionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigningMessage**](TransactionsApi.md#CreateSigningMessage) | **Post** /transactions/signing_message | Create transaction signing message
[**GetAccountTransactions**](TransactionsApi.md#GetAccountTransactions) | **Get** /accounts/{address}/transactions | Get account transactions
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /transactions/{txn_hash_or_version} | Get transaction
[**GetTransactions**](TransactionsApi.md#GetTransactions) | **Get** /transactions | Get transactions
[**SubmitTransaction**](TransactionsApi.md#SubmitTransaction) | **Post** /transactions | Submit transaction



## CreateSigningMessage

> InlineResponse200 CreateSigningMessage(ctx).UserTransactionRequest(userTransactionRequest).Execute()

Create transaction signing message



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
    userTransactionRequest := *openapiclient.NewUserTransactionRequest("0xdd", "32425224034", "32425224034", "32425224034", "XDX", "1635447454", openapiclient.TransactionPayload{ModuleBundlePayload: openapiclient.NewModuleBundlePayload("module_bundle_payload", []openapiclient.MoveModule{*openapiclient.NewMoveModule("0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1")})}) // UserTransactionRequest | User transaction request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.CreateSigningMessage(context.Background()).UserTransactionRequest(userTransactionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateSigningMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSigningMessage`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateSigningMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSigningMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userTransactionRequest** | [**UserTransactionRequest**](UserTransactionRequest.md) | User transaction request | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> []OnChainTransaction GetAccountTransactions(ctx, address).Start(start).Limit(limit).Execute()

Get account transactions

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
    start := int32(1) // int32 | The start transaction version of the page. Default is the latest ledger version. (optional)
    limit := int32(25) // int32 | The max number of transactions should be returned for the page. Default is 25. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetAccountTransactions(context.Background(), address).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactions`: []OnChainTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | The start transaction version of the page. Default is the latest ledger version. | 
 **limit** | **int32** | The max number of transactions should be returned for the page. Default is 25. | 

### Return type

[**[]OnChainTransaction**](OnChainTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> Transaction GetTransaction(ctx, txnHashOrVersion).Execute()

Get transaction



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
    txnHashOrVersion := "txnHashOrVersion_example" // string | * Transaction hash should be hex-encoded bytes string with `0x` prefix. * Transaction version is an `uint64` number. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransaction(context.Background(), txnHashOrVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txnHashOrVersion** | **string** | * Transaction hash should be hex-encoded bytes string with &#x60;0x&#x60; prefix. * Transaction version is an &#x60;uint64&#x60; number.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> []OnChainTransaction GetTransactions(ctx).Start(start).Limit(limit).Execute()

Get transactions

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
    start := int32(1) // int32 | The start transaction version of the page. Default is the latest ledger version. (optional)
    limit := int32(25) // int32 | The max number of transactions should be returned for the page. Default is 25. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactions(context.Background()).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactions`: []OnChainTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | The start transaction version of the page. Default is the latest ledger version. | 
 **limit** | **int32** | The max number of transactions should be returned for the page. Default is 25. | 

### Return type

[**[]OnChainTransaction**](OnChainTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTransaction

> PendingTransaction SubmitTransaction(ctx).SubmitTransactionRequest(submitTransactionRequest).Execute()

Submit transaction



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
    submitTransactionRequest := *openapiclient.NewSubmitTransactionRequest("0xdd", "32425224034", "32425224034", "32425224034", "XDX", "1635447454", openapiclient.TransactionPayload{ModuleBundlePayload: openapiclient.NewModuleBundlePayload("module_bundle_payload", []openapiclient.MoveModule{*openapiclient.NewMoveModule("0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1")})}, openapiclient.TransactionSignature{Ed25519Signature: openapiclient.NewEd25519Signature("ed25519_signature", "0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1", "0x88fbd33f54e1126269769780feb24480428179f552e2313fbe571b72e62a1ca1")}) // SubmitTransactionRequest | User transaction request with transaction sender's signature. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.SubmitTransaction(context.Background()).SubmitTransactionRequest(submitTransactionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.SubmitTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitTransaction`: PendingTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.SubmitTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTransactionRequest** | [**SubmitTransactionRequest**](SubmitTransactionRequest.md) | User transaction request with transaction sender&#39;s signature.  | 

### Return type

[**PendingTransaction**](PendingTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

