# Pumpstatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Pump Status String | [optional] 
**Bolusing** | Pointer to **bool** | Is Pump Bolusing | [optional] 
**Suspended** | Pointer to **bool** | Is Pump Suspended | [optional] 
**Timestamp** | Pointer to **string** | dateString, MUST be ISO &#x60;8601&#x60; format date parseable by Javascript Date() | [optional] 

## Methods

### NewPumpstatus

`func NewPumpstatus() *Pumpstatus`

NewPumpstatus instantiates a new Pumpstatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPumpstatusWithDefaults

`func NewPumpstatusWithDefaults() *Pumpstatus`

NewPumpstatusWithDefaults instantiates a new Pumpstatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Pumpstatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pumpstatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pumpstatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pumpstatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBolusing

`func (o *Pumpstatus) GetBolusing() bool`

GetBolusing returns the Bolusing field if non-nil, zero value otherwise.

### GetBolusingOk

`func (o *Pumpstatus) GetBolusingOk() (*bool, bool)`

GetBolusingOk returns a tuple with the Bolusing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBolusing

`func (o *Pumpstatus) SetBolusing(v bool)`

SetBolusing sets Bolusing field to given value.

### HasBolusing

`func (o *Pumpstatus) HasBolusing() bool`

HasBolusing returns a boolean if a field has been set.

### GetSuspended

`func (o *Pumpstatus) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Pumpstatus) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Pumpstatus) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Pumpstatus) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTimestamp

`func (o *Pumpstatus) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Pumpstatus) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Pumpstatus) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Pumpstatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


