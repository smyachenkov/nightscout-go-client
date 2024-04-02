# \EntriesAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntries**](EntriesAPI.md#AddEntries) | **Post** /entries | Add new entries.
[**EchoStorageSpecGet**](EntriesAPI.md#EchoStorageSpecGet) | **Get** /echo/{storage}/{spec} | View generated Mongo Query object
[**EntriesGet**](EntriesAPI.md#EntriesGet) | **Get** /entries | All Entries matching query
[**EntriesSpecGet**](EntriesAPI.md#EntriesSpecGet) | **Get** /entries/{spec} | All Entries matching query
[**RemoveEntries**](EntriesAPI.md#RemoveEntries) | **Delete** /entries | Delete entries matching query.
[**SliceStorageFieldTypePrefixRegexGet**](EntriesAPI.md#SliceStorageFieldTypePrefixRegexGet) | **Get** /slice/{storage}/{field}/{type}/{prefix}/{regex} | All Entries matching query
[**TimesEchoPrefixRegexGet**](EntriesAPI.md#TimesEchoPrefixRegexGet) | **Get** /times/echo/{prefix}/{regex} | Echo the query object to be used.
[**TimesPrefixRegexGet**](EntriesAPI.md#TimesPrefixRegexGet) | **Get** /times/{prefix}/{regex} | All Entries matching query



## AddEntries

> AddEntries(ctx).Entry(entry).Execute()

Add new entries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	entry := []openapiclient.Entry{*openapiclient.NewEntry()} // []Entry | Entries to be uploaded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntriesAPI.AddEntries(context.Background()).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.AddEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entry** | [**[]Entry**](Entry.md) | Entries to be uploaded. | 

### Return type

 (empty response body)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EchoStorageSpecGet

> interface{} EchoStorageSpecGet(ctx, storage, spec).Find(find).Count(count).Execute()

View generated Mongo Query object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	storage := "storage_example" // string | `entries`, or `treatments` to select the storage layer.  (default to "sgv")
	spec := "spec_example" // string | entry id, such as `55cf81bc436037528ec75fa5` or a type filter such as `sgv`, `mbg`, etc. This parameter is optional.  (default to "sgv")
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings.  (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.EchoStorageSpecGet(context.Background(), storage, spec).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.EchoStorageSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EchoStorageSpecGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.EchoStorageSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storage** | **string** | &#x60;entries&#x60;, or &#x60;treatments&#x60; to select the storage layer.  | [default to &quot;sgv&quot;]
**spec** | **string** | entry id, such as &#x60;55cf81bc436037528ec75fa5&#x60; or a type filter such as &#x60;sgv&#x60;, &#x60;mbg&#x60;, etc. This parameter is optional.  | [default to &quot;sgv&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiEchoStorageSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings.  | 
 **count** | **float32** | Number of entries to return. | 

### Return type

**interface{}**

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntriesGet

> []Entry EntriesGet(ctx).Find(find).Count(count).Execute()

All Entries matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.EntriesGet(context.Background()).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.EntriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntriesGet`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.EntriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of entries to return. | 

### Return type

[**[]Entry**](Entry.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntriesSpecGet

> []Entry EntriesSpecGet(ctx, spec).Find(find).Count(count).Execute()

All Entries matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	spec := "spec_example" // string | entry id, such as `55cf81bc436037528ec75fa5` or a type filter such as `sgv`, `mbg`, etc.  (default to "sgv")
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings.  (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.EntriesSpecGet(context.Background(), spec).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.EntriesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntriesSpecGet`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.EntriesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spec** | **string** | entry id, such as &#x60;55cf81bc436037528ec75fa5&#x60; or a type filter such as &#x60;sgv&#x60;, &#x60;mbg&#x60;, etc.  | [default to &quot;sgv&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiEntriesSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings.  | 
 **count** | **float32** | Number of entries to return. | 

### Return type

[**[]Entry**](Entry.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEntries

> RemoveEntries(ctx).Find(find).Count(count).Execute()

Delete entries matching query.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntriesAPI.RemoveEntries(context.Background()).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.RemoveEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of entries to return. | 

### Return type

 (empty response body)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SliceStorageFieldTypePrefixRegexGet

> []Entry SliceStorageFieldTypePrefixRegexGet(ctx, storage, field, type_, prefix, regex).Find(find).Count(count).Execute()

All Entries matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	storage := "storage_example" // string | Prefix to use in constructing a prefix-based regex, default is `entries`. (default to "entries")
	field := "field_example" // string | Name of the field to use Regex against in query object, default is `dateString`. (default to "dateString")
	type_ := "type__example" // string | The type field to search against, default is sgv. (default to "sgv")
	prefix := "prefix_example" // string | Prefix to use in constructing a prefix-based regex. (default to "2015")
	regex := "regex_example" // string | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: `T{15..17}:.*`, this would search for all records from 3pm to 5pm.  (default to ".*")
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings.  (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.SliceStorageFieldTypePrefixRegexGet(context.Background(), storage, field, type_, prefix, regex).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.SliceStorageFieldTypePrefixRegexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SliceStorageFieldTypePrefixRegexGet`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.SliceStorageFieldTypePrefixRegexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storage** | **string** | Prefix to use in constructing a prefix-based regex, default is &#x60;entries&#x60;. | [default to &quot;entries&quot;]
**field** | **string** | Name of the field to use Regex against in query object, default is &#x60;dateString&#x60;. | [default to &quot;dateString&quot;]
**type_** | **string** | The type field to search against, default is sgv. | [default to &quot;sgv&quot;]
**prefix** | **string** | Prefix to use in constructing a prefix-based regex. | [default to &quot;2015&quot;]
**regex** | **string** | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: &#x60;T{15..17}:.*&#x60;, this would search for all records from 3pm to 5pm.  | [default to &quot;.*&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSliceStorageFieldTypePrefixRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings.  | 
 **count** | **float32** | Number of entries to return. | 

### Return type

[**[]Entry**](Entry.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesEchoPrefixRegexGet

> interface{} TimesEchoPrefixRegexGet(ctx, prefix, regex).Find(find).Count(count).Execute()

Echo the query object to be used.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	prefix := "prefix_example" // string | Prefix to use in constructing a prefix-based regex.
	regex := "regex_example" // string | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: `T{15..17}:.*`, this would search for all records from 3pm to 5pm. 
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.TimesEchoPrefixRegexGet(context.Background(), prefix, regex).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.TimesEchoPrefixRegexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesEchoPrefixRegexGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.TimesEchoPrefixRegexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | Prefix to use in constructing a prefix-based regex. | 
**regex** | **string** | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: &#x60;T{15..17}:.*&#x60;, this would search for all records from 3pm to 5pm.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimesEchoPrefixRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of entries to return. | 

### Return type

**interface{}**

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesPrefixRegexGet

> []Entry TimesPrefixRegexGet(ctx, prefix, regex).Find(find).Count(count).Execute()

All Entries matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client"
)

func main() {
	prefix := "prefix_example" // string | Prefix to use in constructing a prefix-based regex.
	regex := "regex_example" // string | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: `T{15..17}:.*`, this would search for all records from 3pm to 5pm. 
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.TimesPrefixRegexGet(context.Background(), prefix, regex).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.TimesPrefixRegexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesPrefixRegexGet`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.TimesPrefixRegexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | Prefix to use in constructing a prefix-based regex. | 
**regex** | **string** | Tail part of regexp to use in expanding/construccting a query object. Regexp also has bash-style brace and glob expansion applied to it, creating ways to search for modal times of day, perhaps using something like this syntax: &#x60;T{15..17}:.*&#x60;, this would search for all records from 3pm to 5pm.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimesPrefixRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of entries to return. | 

### Return type

[**[]Entry**](Entry.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

