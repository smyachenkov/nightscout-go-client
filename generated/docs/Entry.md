# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | sgv, mbg, cal, etc | [optional] 
**DateString** | Pointer to **string** | dateString, MUST be ISO &#x60;8601&#x60; format date parseable by Javascript Date() | [optional] 
**Date** | Pointer to **float32** | Epoch | [optional] 
**Sgv** | Pointer to **float32** | The glucose reading. (only available for sgv types) | [optional] 
**Direction** | Pointer to **string** | Direction of glucose trend reported by CGM. (only available for sgv types) | [optional] 
**Noise** | Pointer to **float32** | Noise level at time of reading. (only available for sgv types) | [optional] 
**Filtered** | Pointer to **float32** | The raw filtered value directly from CGM transmitter. (only available for sgv types) | [optional] 
**Unfiltered** | Pointer to **float32** | The raw unfiltered value directly from CGM transmitter. (only available for sgv types) | [optional] 
**Rssi** | Pointer to **float32** | The signal strength from CGM transmitter. (only available for sgv types) | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Entry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Entry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateString

`func (o *Entry) GetDateString() string`

GetDateString returns the DateString field if non-nil, zero value otherwise.

### GetDateStringOk

`func (o *Entry) GetDateStringOk() (*string, bool)`

GetDateStringOk returns a tuple with the DateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateString

`func (o *Entry) SetDateString(v string)`

SetDateString sets DateString field to given value.

### HasDateString

`func (o *Entry) HasDateString() bool`

HasDateString returns a boolean if a field has been set.

### GetDate

`func (o *Entry) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Entry) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Entry) SetDate(v float32)`

SetDate sets Date field to given value.

### HasDate

`func (o *Entry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSgv

`func (o *Entry) GetSgv() float32`

GetSgv returns the Sgv field if non-nil, zero value otherwise.

### GetSgvOk

`func (o *Entry) GetSgvOk() (*float32, bool)`

GetSgvOk returns a tuple with the Sgv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgv

`func (o *Entry) SetSgv(v float32)`

SetSgv sets Sgv field to given value.

### HasSgv

`func (o *Entry) HasSgv() bool`

HasSgv returns a boolean if a field has been set.

### GetDirection

`func (o *Entry) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Entry) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Entry) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Entry) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetNoise

`func (o *Entry) GetNoise() float32`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *Entry) GetNoiseOk() (*float32, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *Entry) SetNoise(v float32)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *Entry) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetFiltered

`func (o *Entry) GetFiltered() float32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *Entry) GetFilteredOk() (*float32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *Entry) SetFiltered(v float32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *Entry) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetUnfiltered

`func (o *Entry) GetUnfiltered() float32`

GetUnfiltered returns the Unfiltered field if non-nil, zero value otherwise.

### GetUnfilteredOk

`func (o *Entry) GetUnfilteredOk() (*float32, bool)`

GetUnfilteredOk returns a tuple with the Unfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfiltered

`func (o *Entry) SetUnfiltered(v float32)`

SetUnfiltered sets Unfiltered field to given value.

### HasUnfiltered

`func (o *Entry) HasUnfiltered() bool`

HasUnfiltered returns a boolean if a field has been set.

### GetRssi

`func (o *Entry) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *Entry) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *Entry) SetRssi(v float32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *Entry) HasRssi() bool`

HasRssi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


