# Optime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | Time the operation started | [optional] 
**T** | Pointer to **int32** | Time the operation took to complete | [optional] 

## Methods

### NewOptime

`func NewOptime() *Optime`

NewOptime instantiates a new Optime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimeWithDefaults

`func NewOptimeWithDefaults() *Optime`

NewOptimeWithDefaults instantiates a new Optime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *Optime) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *Optime) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *Optime) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *Optime) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetT

`func (o *Optime) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *Optime) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *Optime) SetT(v int32)`

SetT sets T field to given value.

### HasT

`func (o *Optime) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


