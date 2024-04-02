# Pumpbattery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Pump Battery Status String | [optional] 
**Voltage** | Pointer to **float32** | Pump Battery Voltage Level | [optional] 

## Methods

### NewPumpbattery

`func NewPumpbattery() *Pumpbattery`

NewPumpbattery instantiates a new Pumpbattery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPumpbatteryWithDefaults

`func NewPumpbatteryWithDefaults() *Pumpbattery`

NewPumpbatteryWithDefaults instantiates a new Pumpbattery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Pumpbattery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pumpbattery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pumpbattery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pumpbattery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVoltage

`func (o *Pumpbattery) GetVoltage() float32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *Pumpbattery) GetVoltageOk() (*float32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *Pumpbattery) SetVoltage(v float32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *Pumpbattery) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


