# \TreatmentsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTreatments**](TreatmentsAPI.md#AddTreatments) | **Post** /treatments | Add new treatments.
[**RemoveTreatments**](TreatmentsAPI.md#RemoveTreatments) | **Delete** /treatments | Delete treatments matching query.
[**TreatmentsGet**](TreatmentsAPI.md#TreatmentsGet) | **Get** /treatments | Treatments
[**TreatmentsSpecDelete**](TreatmentsAPI.md#TreatmentsSpecDelete) | **Delete** /treatments/{spec} | Delete treatments record with id provided in spec



## AddTreatments

> AddTreatments(ctx).Treatment(treatment).Execute()

Add new treatments.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client/generated"
)

func main() {
	treatment := []openapiclient.Treatment{*openapiclient.NewTreatment()} // []Treatment | Treatments to be uploaded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreatmentsAPI.AddTreatments(context.Background()).Treatment(treatment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreatmentsAPI.AddTreatments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTreatmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **treatment** | [**[]Treatment**](Treatment.md) | Treatments to be uploaded. | 

### Return type

 (empty response body)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTreatments

> RemoveTreatments(ctx).Find(find).Count(count).Execute()

Delete treatments matching query.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client/generated"
)

func main() {
	find := "find_example" // string | The query used to find treatments to delete, support nested query syntax, for example `find[insulin][$gte]=3`. All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreatmentsAPI.RemoveTreatments(context.Background()).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreatmentsAPI.RemoveTreatments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTreatmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find treatments to delete, support nested query syntax, for example &#x60;find[insulin][$gte]&#x3D;3&#x60;. All find parameters are interpreted as strings. | 
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


## TreatmentsGet

> []Treatment TreatmentsGet(ctx).Find(find).Count(count).Execute()

Treatments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client/generated"
)

func main() {
	find := "find_example" // string | The query used to find entries, supports nested query syntax.  Examples `find[insulin][$gte]=3` `find[carb][$gte]=100` `find[eventType]=Correction+Bolus` All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of entries to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreatmentsAPI.TreatmentsGet(context.Background()).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreatmentsAPI.TreatmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TreatmentsGet`: []Treatment
	fmt.Fprintf(os.Stdout, "Response from `TreatmentsAPI.TreatmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTreatmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find entries, supports nested query syntax.  Examples &#x60;find[insulin][$gte]&#x3D;3&#x60; &#x60;find[carb][$gte]&#x3D;100&#x60; &#x60;find[eventType]&#x3D;Correction+Bolus&#x60; All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of entries to return. | 

### Return type

[**[]Treatment**](Treatment.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TreatmentsSpecDelete

> DeleteStatus TreatmentsSpecDelete(ctx, spec).Execute()

Delete treatments record with id provided in spec



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-go-client/generated"
)

func main() {
	spec := "spec_example" // string | treatment id, such as `55cf81bc436037528ec75fa5` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreatmentsAPI.TreatmentsSpecDelete(context.Background(), spec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreatmentsAPI.TreatmentsSpecDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TreatmentsSpecDelete`: DeleteStatus
	fmt.Fprintf(os.Stdout, "Response from `TreatmentsAPI.TreatmentsSpecDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spec** | **string** | treatment id, such as &#x60;55cf81bc436037528ec75fa5&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTreatmentsSpecDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteStatus**](DeleteStatus.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

