# \LogsIndexesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogsIndex**](LogsIndexesApi.md#GetLogsIndex) | **Get** /api/v1/logs/config/indexes/{name} | Get an index
[**GetLogsIndexOrder**](LogsIndexesApi.md#GetLogsIndexOrder) | **Get** /api/v1/logs/config/index-order | Get indexes order
[**ListLogIndexes**](LogsIndexesApi.md#ListLogIndexes) | **Get** /api/v1/logs/config/indexes | Get all indexes
[**UpdateLogsIndex**](LogsIndexesApi.md#UpdateLogsIndex) | **Put** /api/v1/logs/config/indexes/{name} | Update an index
[**UpdateLogsIndexOrder**](LogsIndexesApi.md#UpdateLogsIndexOrder) | **Put** /api/v1/logs/config/index-order | Update indexes order



## GetLogsIndex

> LogsIndex GetLogsIndex(ctx, name).Execute()

Get an index



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    name := "name_example" // string | Name of the log index.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetLogsIndex", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsIndexesApi.GetLogsIndex(ctx, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.GetLogsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsIndex`: LogsIndex
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.GetLogsIndex:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the log index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsIndexOrder

> LogsIndexesOrder GetLogsIndexOrder(ctx).Execute()

Get indexes order



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )


    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetLogsIndexOrder", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsIndexesApi.GetLogsIndexOrder(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.GetLogsIndexOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsIndexOrder`: LogsIndexesOrder
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.GetLogsIndexOrder:\n%s\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsIndexOrderRequest struct via the builder pattern


### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogIndexes

> LogsIndexListResponse ListLogIndexes(ctx).Execute()

Get all indexes



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )


    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListLogIndexes", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsIndexesApi.ListLogIndexes(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.ListLogIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogIndexes`: LogsIndexListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.ListLogIndexes:\n%s\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogIndexesRequest struct via the builder pattern


### Return type

[**LogsIndexListResponse**](LogsIndexListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsIndex

> LogsIndex UpdateLogsIndex(ctx, name).Body(body).Execute()

Update an index



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    name := "name_example" // string | Name of the log index.
    body := *datadog.NewLogsIndexUpdateRequest(*datadog.NewLogsFilter()) // LogsIndexUpdateRequest | Object containing the new `LogsIndexUpdateRequest`.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateLogsIndex", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsIndexesApi.UpdateLogsIndex(ctx, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.UpdateLogsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsIndex`: LogsIndex
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.UpdateLogsIndex:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the log index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsIndexUpdateRequest**](LogsIndexUpdateRequest.md) | Object containing the new &#x60;LogsIndexUpdateRequest&#x60;. | 

### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsIndexOrder

> LogsIndexesOrder UpdateLogsIndexOrder(ctx).Body(body).Execute()

Update indexes order



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := *datadog.NewLogsIndexesOrder([]string{"IndexNames_example"}) // LogsIndexesOrder | Object containing the new ordered list of index names

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateLogsIndexOrder", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsIndexesApi.UpdateLogsIndexOrder(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.UpdateLogsIndexOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsIndexOrder`: LogsIndexesOrder
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.UpdateLogsIndexOrder:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsIndexOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsIndexesOrder**](LogsIndexesOrder.md) | Object containing the new ordered list of index names | 

### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

