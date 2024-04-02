# DeleteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N** | Pointer to **int32** | Number of records deleted | [optional] 
**Optime** | Pointer to [**Optime**](Optime.md) |  | [optional] 
**ElectionId** | Pointer to **string** | Election id of operation | [optional] 
**Ok** | Pointer to **int32** | Status of whether delete was successful | [optional] 
**OperationTime** | Pointer to **string** | Time delete operation was executed | [optional] 
**ClusterTime** | Pointer to **string** | Information about execution time in cluster environment | [optional] 

## Methods

### NewDeleteStatus

`func NewDeleteStatus() *DeleteStatus`

NewDeleteStatus instantiates a new DeleteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteStatusWithDefaults

`func NewDeleteStatusWithDefaults() *DeleteStatus`

NewDeleteStatusWithDefaults instantiates a new DeleteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN

`func (o *DeleteStatus) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *DeleteStatus) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *DeleteStatus) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *DeleteStatus) HasN() bool`

HasN returns a boolean if a field has been set.

### GetOptime

`func (o *DeleteStatus) GetOptime() Optime`

GetOptime returns the Optime field if non-nil, zero value otherwise.

### GetOptimeOk

`func (o *DeleteStatus) GetOptimeOk() (*Optime, bool)`

GetOptimeOk returns a tuple with the Optime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptime

`func (o *DeleteStatus) SetOptime(v Optime)`

SetOptime sets Optime field to given value.

### HasOptime

`func (o *DeleteStatus) HasOptime() bool`

HasOptime returns a boolean if a field has been set.

### GetElectionId

`func (o *DeleteStatus) GetElectionId() string`

GetElectionId returns the ElectionId field if non-nil, zero value otherwise.

### GetElectionIdOk

`func (o *DeleteStatus) GetElectionIdOk() (*string, bool)`

GetElectionIdOk returns a tuple with the ElectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectionId

`func (o *DeleteStatus) SetElectionId(v string)`

SetElectionId sets ElectionId field to given value.

### HasElectionId

`func (o *DeleteStatus) HasElectionId() bool`

HasElectionId returns a boolean if a field has been set.

### GetOk

`func (o *DeleteStatus) GetOk() int32`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DeleteStatus) GetOkOk() (*int32, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DeleteStatus) SetOk(v int32)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DeleteStatus) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetOperationTime

`func (o *DeleteStatus) GetOperationTime() string`

GetOperationTime returns the OperationTime field if non-nil, zero value otherwise.

### GetOperationTimeOk

`func (o *DeleteStatus) GetOperationTimeOk() (*string, bool)`

GetOperationTimeOk returns a tuple with the OperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTime

`func (o *DeleteStatus) SetOperationTime(v string)`

SetOperationTime sets OperationTime field to given value.

### HasOperationTime

`func (o *DeleteStatus) HasOperationTime() bool`

HasOperationTime returns a boolean if a field has been set.

### GetClusterTime

`func (o *DeleteStatus) GetClusterTime() string`

GetClusterTime returns the ClusterTime field if non-nil, zero value otherwise.

### GetClusterTimeOk

`func (o *DeleteStatus) GetClusterTimeOk() (*string, bool)`

GetClusterTimeOk returns a tuple with the ClusterTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTime

`func (o *DeleteStatus) SetClusterTime(v string)`

SetClusterTime sets ClusterTime field to given value.

### HasClusterTime

`func (o *DeleteStatus) HasClusterTime() bool`

HasClusterTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


