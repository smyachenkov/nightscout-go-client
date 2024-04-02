# \ProfileAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfileGet**](ProfileAPI.md#ProfileGet) | **Get** /profile | Profile



## ProfileGet

> Profile ProfileGet(ctx).Execute()

Profile



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.ProfileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProfileGet`: Profile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.ProfileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProfileGetRequest struct via the builder pattern


### Return type

[**Profile**](Profile.md)

### Authorization

[jwtoken](../README.md#jwtoken), [token_in_url](../README.md#token_in_url), [api_secret](../README.md#api_secret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

