# Pump

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clock** | Pointer to **string** | dateString, MUST be ISO &#x60;8601&#x60; format date parseable by Javascript Date() | [optional] 
**Battery** | Pointer to [**Pumpbattery**](Pumpbattery.md) |  | [optional] 
**Reservoir** | Pointer to **float32** | Amount of insulin remaining in pump reservoir | [optional] 
**Status** | Pointer to [**Pumpstatus**](Pumpstatus.md) |  | [optional] 

## Methods

### NewPump

`func NewPump() *Pump`

NewPump instantiates a new Pump object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPumpWithDefaults

`func NewPumpWithDefaults() *Pump`

NewPumpWithDefaults instantiates a new Pump object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClock

`func (o *Pump) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *Pump) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *Pump) SetClock(v string)`

SetClock sets Clock field to given value.

### HasClock

`func (o *Pump) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetBattery

`func (o *Pump) GetBattery() Pumpbattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *Pump) GetBatteryOk() (*Pumpbattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *Pump) SetBattery(v Pumpbattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *Pump) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetReservoir

`func (o *Pump) GetReservoir() float32`

GetReservoir returns the Reservoir field if non-nil, zero value otherwise.

### GetReservoirOk

`func (o *Pump) GetReservoirOk() (*float32, bool)`

GetReservoirOk returns a tuple with the Reservoir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservoir

`func (o *Pump) SetReservoir(v float32)`

SetReservoir sets Reservoir field to given value.

### HasReservoir

`func (o *Pump) HasReservoir() bool`

HasReservoir returns a boolean if a field has been set.

### GetStatus

`func (o *Pump) GetStatus() Pumpstatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pump) GetStatusOk() (*Pumpstatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pump) SetStatus(v Pumpstatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pump) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


