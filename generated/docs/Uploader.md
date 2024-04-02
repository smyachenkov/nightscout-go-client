# Uploader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryVoltage** | Pointer to **float32** | Uploader Device Battery Voltage | [optional] 
**Battery** | Pointer to **float32** | Uploader Device Battery Percentage Charge Remaining | [optional] 

## Methods

### NewUploader

`func NewUploader() *Uploader`

NewUploader instantiates a new Uploader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploaderWithDefaults

`func NewUploaderWithDefaults() *Uploader`

NewUploaderWithDefaults instantiates a new Uploader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryVoltage

`func (o *Uploader) GetBatteryVoltage() float32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *Uploader) GetBatteryVoltageOk() (*float32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *Uploader) SetBatteryVoltage(v float32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *Uploader) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### GetBattery

`func (o *Uploader) GetBattery() float32`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *Uploader) GetBatteryOk() (*float32, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *Uploader) SetBattery(v float32)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *Uploader) HasBattery() bool`

HasBattery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


