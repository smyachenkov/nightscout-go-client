# Devicestatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** | Device type and hostname for example openaps://hostname | 
**CreatedAt** | **string** | dateString, MUST be ISO &#x60;8601&#x60; format date parseable by Javascript Date() | 
**Openaps** | Pointer to **string** | OpenAPS devicestatus record - TODO: Fill Out Details | [optional] 
**Loop** | Pointer to **string** | Loop devicestatus record - TODO: Fill Out Details | [optional] 
**Pump** | Pointer to [**Pump**](Pump.md) |  | [optional] 
**Uploader** | Pointer to [**Uploader**](Uploader.md) |  | [optional] 
**Xdripjs** | Pointer to [**Xdripjs**](Xdripjs.md) |  | [optional] 

## Methods

### NewDevicestatus

`func NewDevicestatus(device string, createdAt string, ) *Devicestatus`

NewDevicestatus instantiates a new Devicestatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicestatusWithDefaults

`func NewDevicestatusWithDefaults() *Devicestatus`

NewDevicestatusWithDefaults instantiates a new Devicestatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *Devicestatus) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Devicestatus) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Devicestatus) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetCreatedAt

`func (o *Devicestatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Devicestatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Devicestatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetOpenaps

`func (o *Devicestatus) GetOpenaps() string`

GetOpenaps returns the Openaps field if non-nil, zero value otherwise.

### GetOpenapsOk

`func (o *Devicestatus) GetOpenapsOk() (*string, bool)`

GetOpenapsOk returns a tuple with the Openaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaps

`func (o *Devicestatus) SetOpenaps(v string)`

SetOpenaps sets Openaps field to given value.

### HasOpenaps

`func (o *Devicestatus) HasOpenaps() bool`

HasOpenaps returns a boolean if a field has been set.

### GetLoop

`func (o *Devicestatus) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *Devicestatus) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *Devicestatus) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *Devicestatus) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetPump

`func (o *Devicestatus) GetPump() Pump`

GetPump returns the Pump field if non-nil, zero value otherwise.

### GetPumpOk

`func (o *Devicestatus) GetPumpOk() (*Pump, bool)`

GetPumpOk returns a tuple with the Pump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPump

`func (o *Devicestatus) SetPump(v Pump)`

SetPump sets Pump field to given value.

### HasPump

`func (o *Devicestatus) HasPump() bool`

HasPump returns a boolean if a field has been set.

### GetUploader

`func (o *Devicestatus) GetUploader() Uploader`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *Devicestatus) GetUploaderOk() (*Uploader, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *Devicestatus) SetUploader(v Uploader)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *Devicestatus) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetXdripjs

`func (o *Devicestatus) GetXdripjs() Xdripjs`

GetXdripjs returns the Xdripjs field if non-nil, zero value otherwise.

### GetXdripjsOk

`func (o *Devicestatus) GetXdripjsOk() (*Xdripjs, bool)`

GetXdripjsOk returns a tuple with the Xdripjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXdripjs

`func (o *Devicestatus) SetXdripjs(v Xdripjs)`

SetXdripjs sets Xdripjs field to given value.

### HasXdripjs

`func (o *Devicestatus) HasXdripjs() bool`

HasXdripjs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


