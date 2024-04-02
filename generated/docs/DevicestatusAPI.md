# \DevicestatusAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDevicestatuses**](DevicestatusAPI.md#AddDevicestatuses) | **Post** /devicestatus/ | Add new devicestatus records.
[**DevicestatusDelete**](DevicestatusAPI.md#DevicestatusDelete) | **Delete** /devicestatus/ | Delete all Devicestatus records matching query
[**DevicestatusGet**](DevicestatusAPI.md#DevicestatusGet) | **Get** /devicestatus/ | All Devicestatuses matching query
[**DevicestatusSpecDelete**](DevicestatusAPI.md#DevicestatusSpecDelete) | **Delete** /devicestatus/{spec} | Delete devicestatus record with id provided in spec



## AddDevicestatuses

> AddDevicestatuses(ctx).Devicestatus(devicestatus).Execute()

Add new devicestatus records.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-openapi-client"
)

func main() {
	devicestatus := []openapiclient.Devicestatus{*openapiclient.NewDevicestatus("Device_example", "CreatedAt_example")} // []Devicestatus | Device statuses to be uploaded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicestatusAPI.AddDevicestatuses(context.Background()).Devicestatus(devicestatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicestatusAPI.AddDevicestatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicestatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devicestatus** | [**[]Devicestatus**](Devicestatus.md) | Device statuses to be uploaded. | 

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


## DevicestatusDelete

> DeleteStatus DevicestatusDelete(ctx).Find(find).Execute()

Delete all Devicestatus records matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-openapi-client"
)

func main() {
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[created_at][$gte]=2015-08-27`.  All find parameters are interpreted as strings.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicestatusAPI.DevicestatusDelete(context.Background()).Find(find).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicestatusAPI.DevicestatusDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicestatusDelete`: DeleteStatus
	fmt.Fprintf(os.Stdout, "Response from `DevicestatusAPI.DevicestatusDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicestatusDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[created_at][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings.  | 

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


## DevicestatusGet

> []Devicestatus DevicestatusGet(ctx).Find(find).Count(count).Execute()

All Devicestatuses matching query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-openapi-client"
)

func main() {
	find := "find_example" // string | The query used to find entries, support nested query syntax, for example `find[dateString][$gte]=2015-08-27`.  All find parameters are interpreted as strings. (optional)
	count := float32(8.14) // float32 | Number of devicestatus records to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicestatusAPI.DevicestatusGet(context.Background()).Find(find).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicestatusAPI.DevicestatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicestatusGet`: []Devicestatus
	fmt.Fprintf(os.Stdout, "Response from `DevicestatusAPI.DevicestatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicestatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **find** | **string** | The query used to find entries, support nested query syntax, for example &#x60;find[dateString][$gte]&#x3D;2015-08-27&#x60;.  All find parameters are interpreted as strings. | 
 **count** | **float32** | Number of devicestatus records to return. | 

### Return type

[**[]Devicestatus**](Devicestatus.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicestatusSpecDelete

> DeleteStatus DevicestatusSpecDelete(ctx, spec).Execute()

Delete devicestatus record with id provided in spec



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smyachenkov/nightscout-openapi-client"
)

func main() {
	spec := "spec_example" // string | entry id, such as `55cf81bc436037528ec75fa5` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicestatusAPI.DevicestatusSpecDelete(context.Background(), spec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicestatusAPI.DevicestatusSpecDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicestatusSpecDelete`: DeleteStatus
	fmt.Fprintf(os.Stdout, "Response from `DevicestatusAPI.DevicestatusSpecDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spec** | **string** | entry id, such as &#x60;55cf81bc436037528ec75fa5&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicestatusSpecDeleteRequest struct via the builder pattern


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

