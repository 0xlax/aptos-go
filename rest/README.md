# Go API client for openapi

The Aptos Node API is a RESTful API for client applications to interact with the Aptos blockchain.


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/0xlax/aptos-go/rest"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /accounts/{address} | Get account
*AccountsApi* | [**GetAccountModule**](docs/AccountsApi.md#getaccountmodule) | **Get** /accounts/{address}/module/{module_name} | Get module by module id.
*AccountsApi* | [**GetAccountModules**](docs/AccountsApi.md#getaccountmodules) | **Get** /accounts/{address}/modules | Get account modules
*AccountsApi* | [**GetAccountResource**](docs/AccountsApi.md#getaccountresource) | **Get** /accounts/{address}/resource/{resource_type} | Get resource by account address and resource type.
*AccountsApi* | [**GetAccountResources**](docs/AccountsApi.md#getaccountresources) | **Get** /accounts/{address}/resources | Get account resources
*EventsApi* | [**GetEventsByEventHandle**](docs/EventsApi.md#geteventsbyeventhandle) | **Get** /accounts/{address}/events/{event_handle_struct}/{field_name} | Get events by event handle
*EventsApi* | [**GetEventsByEventKey**](docs/EventsApi.md#geteventsbyeventkey) | **Get** /events/{event_key} | Get events by event key
*GeneralApi* | [**GetLedgerInfo**](docs/GeneralApi.md#getledgerinfo) | **Get** / | Ledger information
*GeneralApi* | [**GetSpecHtml**](docs/GeneralApi.md#getspechtml) | **Get** /spec.html | API document
*GeneralApi* | [**GetSpecYaml**](docs/GeneralApi.md#getspecyaml) | **Get** /openapi.yaml | OpenAPI specification
*StateApi* | [**GetAccount**](docs/StateApi.md#getaccount) | **Get** /accounts/{address} | Get account
*StateApi* | [**GetAccountModule**](docs/StateApi.md#getaccountmodule) | **Get** /accounts/{address}/module/{module_name} | Get module by module id.
*StateApi* | [**GetAccountModules**](docs/StateApi.md#getaccountmodules) | **Get** /accounts/{address}/modules | Get account modules
*StateApi* | [**GetAccountResource**](docs/StateApi.md#getaccountresource) | **Get** /accounts/{address}/resource/{resource_type} | Get resource by account address and resource type.
*StateApi* | [**GetAccountResources**](docs/StateApi.md#getaccountresources) | **Get** /accounts/{address}/resources | Get account resources
*StateApi* | [**GetTableItem**](docs/StateApi.md#gettableitem) | **Post** /tables/{table_handle}/item | Get table item by handle and key.
*TableApi* | [**GetTableItem**](docs/TableApi.md#gettableitem) | **Post** /tables/{table_handle}/item | Get table item by handle and key.
*TransactionsApi* | [**CreateSigningMessage**](docs/TransactionsApi.md#createsigningmessage) | **Post** /transactions/signing_message | Create transaction signing message
*TransactionsApi* | [**GetAccountTransactions**](docs/TransactionsApi.md#getaccounttransactions) | **Get** /accounts/{address}/transactions | Get account transactions
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /transactions/{txn_hash_or_version} | Get transaction
*TransactionsApi* | [**GetTransactions**](docs/TransactionsApi.md#gettransactions) | **Get** /transactions | Get transactions
*TransactionsApi* | [**SubmitTransaction**](docs/TransactionsApi.md#submittransaction) | **Post** /transactions | Submit transaction


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountResource](docs/AccountResource.md)
 - [AccountSignature](docs/AccountSignature.md)
 - [AptosError](docs/AptosError.md)
 - [BlockMetadataTransaction](docs/BlockMetadataTransaction.md)
 - [BlockMetadataTransactionAllOf](docs/BlockMetadataTransactionAllOf.md)
 - [DeleteModule](docs/DeleteModule.md)
 - [DeleteResource](docs/DeleteResource.md)
 - [DeleteTableItem](docs/DeleteTableItem.md)
 - [DirectWriteSet](docs/DirectWriteSet.md)
 - [Ed25519Signature](docs/Ed25519Signature.md)
 - [Event](docs/Event.md)
 - [GenesisTransaction](docs/GenesisTransaction.md)
 - [GenesisTransactionAllOf](docs/GenesisTransactionAllOf.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [LedgerInfo](docs/LedgerInfo.md)
 - [ModuleBundlePayload](docs/ModuleBundlePayload.md)
 - [MoveAbility](docs/MoveAbility.md)
 - [MoveFunction](docs/MoveFunction.md)
 - [MoveFunctionGenericTypeParams](docs/MoveFunctionGenericTypeParams.md)
 - [MoveModule](docs/MoveModule.md)
 - [MoveModuleABI](docs/MoveModuleABI.md)
 - [MoveScript](docs/MoveScript.md)
 - [MoveStruct](docs/MoveStruct.md)
 - [MoveStructField](docs/MoveStructField.md)
 - [MoveStructGenericTypeParams](docs/MoveStructGenericTypeParams.md)
 - [MultiAgentSignature](docs/MultiAgentSignature.md)
 - [MultiEd25519Signature](docs/MultiEd25519Signature.md)
 - [OnChainTransaction](docs/OnChainTransaction.md)
 - [OnChainTransactionInfo](docs/OnChainTransactionInfo.md)
 - [PendingTransaction](docs/PendingTransaction.md)
 - [PendingTransactionAllOf](docs/PendingTransactionAllOf.md)
 - [Script](docs/Script.md)
 - [ScriptFunctionPayload](docs/ScriptFunctionPayload.md)
 - [ScriptPayload](docs/ScriptPayload.md)
 - [ScriptWriteSet](docs/ScriptWriteSet.md)
 - [SubmitTransactionRequest](docs/SubmitTransactionRequest.md)
 - [TableItemDeletion](docs/TableItemDeletion.md)
 - [TableItemRequest](docs/TableItemRequest.md)
 - [TableItemWrite](docs/TableItemWrite.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionPayload](docs/TransactionPayload.md)
 - [TransactionSignature](docs/TransactionSignature.md)
 - [UserTransaction](docs/UserTransaction.md)
 - [UserTransactionAllOf](docs/UserTransactionAllOf.md)
 - [UserTransactionRequest](docs/UserTransactionRequest.md)
 - [UserTransactionSignature](docs/UserTransactionSignature.md)
 - [WriteModule](docs/WriteModule.md)
 - [WriteResource](docs/WriteResource.md)
 - [WriteSet](docs/WriteSet.md)
 - [WriteSetChange](docs/WriteSetChange.md)
 - [WriteSetPayload](docs/WriteSetPayload.md)
 - [WriteTableItem](docs/WriteTableItem.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



