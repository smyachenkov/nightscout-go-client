# Xdripjs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **float32** | CGM Sensor Session State Code | [optional] 
**StateString** | Pointer to **string** | CGM Sensor Session State String | [optional] 
**StateStringShort** | Pointer to **string** | CGM Sensor Session State Short String | [optional] 
**TxId** | Pointer to **string** | CGM Transmitter ID | [optional] 
**TxStatus** | Pointer to **float32** | CGM Transmitter Status | [optional] 
**TxStatusString** | Pointer to **string** | CGM Transmitter Status String | [optional] 
**TxStatusStringShort** | Pointer to **string** | CGM Transmitter Status Short String | [optional] 
**TxActivation** | Pointer to **float32** | CGM Transmitter Activation Milliseconds After Epoch | [optional] 
**Mode** | Pointer to **string** | Mode xdrip-js Application Operationg: expired, not expired, etc. | [optional] 
**Timestamp** | Pointer to **float32** | Last Update Milliseconds After Epoch | [optional] 
**Rssi** | Pointer to **float32** | Receive Signal Strength of Transmitter | [optional] 
**Unfiltered** | Pointer to **float32** | Most Recent Raw Unfiltered Glucose | [optional] 
**Filtered** | Pointer to **float32** | Most Recent Raw Filtered Glucose | [optional] 
**Noise** | Pointer to **float32** | Calculated Noise Value - 1&#x3D;Clean, 2&#x3D;Light, 3&#x3D;Medium, 4&#x3D;Heavy | [optional] 
**NoiseString** | Pointer to **float32** | Noise Value String | [optional] 
**Slope** | Pointer to **float32** | Calibration Slope Value | [optional] 
**Intercept** | Pointer to **float32** | Calibration Intercept Value | [optional] 
**CalType** | Pointer to **string** | Algorithm Used to Calculate Calibration Values | [optional] 
**LastCalibrationDate** | Pointer to **float32** | Most Recent Calibration Milliseconds After Epoch | [optional] 
**SessionStart** | Pointer to **float32** | Sensor Session Start Milliseconds After Epoch | [optional] 
**BatteryTimestamp** | Pointer to **float32** | Most Recent Batter Status Read Milliseconds After Epoch | [optional] 
**Voltagea** | Pointer to **float32** | Voltage of Battery A | [optional] 
**Voltageb** | Pointer to **float32** | Voltage of Battery B | [optional] 
**Temperature** | Pointer to **float32** | Transmitter Temperature | [optional] 
**Resistance** | Pointer to **float32** | Sensor Resistance | [optional] 

## Methods

### NewXdripjs

`func NewXdripjs() *Xdripjs`

NewXdripjs instantiates a new Xdripjs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXdripjsWithDefaults

`func NewXdripjsWithDefaults() *Xdripjs`

NewXdripjsWithDefaults instantiates a new Xdripjs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Xdripjs) GetState() float32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Xdripjs) GetStateOk() (*float32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Xdripjs) SetState(v float32)`

SetState sets State field to given value.

### HasState

`func (o *Xdripjs) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *Xdripjs) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *Xdripjs) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *Xdripjs) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *Xdripjs) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### GetStateStringShort

`func (o *Xdripjs) GetStateStringShort() string`

GetStateStringShort returns the StateStringShort field if non-nil, zero value otherwise.

### GetStateStringShortOk

`func (o *Xdripjs) GetStateStringShortOk() (*string, bool)`

GetStateStringShortOk returns a tuple with the StateStringShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateStringShort

`func (o *Xdripjs) SetStateStringShort(v string)`

SetStateStringShort sets StateStringShort field to given value.

### HasStateStringShort

`func (o *Xdripjs) HasStateStringShort() bool`

HasStateStringShort returns a boolean if a field has been set.

### GetTxId

`func (o *Xdripjs) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *Xdripjs) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *Xdripjs) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *Xdripjs) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetTxStatus

`func (o *Xdripjs) GetTxStatus() float32`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *Xdripjs) GetTxStatusOk() (*float32, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *Xdripjs) SetTxStatus(v float32)`

SetTxStatus sets TxStatus field to given value.

### HasTxStatus

`func (o *Xdripjs) HasTxStatus() bool`

HasTxStatus returns a boolean if a field has been set.

### GetTxStatusString

`func (o *Xdripjs) GetTxStatusString() string`

GetTxStatusString returns the TxStatusString field if non-nil, zero value otherwise.

### GetTxStatusStringOk

`func (o *Xdripjs) GetTxStatusStringOk() (*string, bool)`

GetTxStatusStringOk returns a tuple with the TxStatusString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatusString

`func (o *Xdripjs) SetTxStatusString(v string)`

SetTxStatusString sets TxStatusString field to given value.

### HasTxStatusString

`func (o *Xdripjs) HasTxStatusString() bool`

HasTxStatusString returns a boolean if a field has been set.

### GetTxStatusStringShort

`func (o *Xdripjs) GetTxStatusStringShort() string`

GetTxStatusStringShort returns the TxStatusStringShort field if non-nil, zero value otherwise.

### GetTxStatusStringShortOk

`func (o *Xdripjs) GetTxStatusStringShortOk() (*string, bool)`

GetTxStatusStringShortOk returns a tuple with the TxStatusStringShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatusStringShort

`func (o *Xdripjs) SetTxStatusStringShort(v string)`

SetTxStatusStringShort sets TxStatusStringShort field to given value.

### HasTxStatusStringShort

`func (o *Xdripjs) HasTxStatusStringShort() bool`

HasTxStatusStringShort returns a boolean if a field has been set.

### GetTxActivation

`func (o *Xdripjs) GetTxActivation() float32`

GetTxActivation returns the TxActivation field if non-nil, zero value otherwise.

### GetTxActivationOk

`func (o *Xdripjs) GetTxActivationOk() (*float32, bool)`

GetTxActivationOk returns a tuple with the TxActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxActivation

`func (o *Xdripjs) SetTxActivation(v float32)`

SetTxActivation sets TxActivation field to given value.

### HasTxActivation

`func (o *Xdripjs) HasTxActivation() bool`

HasTxActivation returns a boolean if a field has been set.

### GetMode

`func (o *Xdripjs) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Xdripjs) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Xdripjs) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Xdripjs) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTimestamp

`func (o *Xdripjs) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Xdripjs) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Xdripjs) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Xdripjs) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRssi

`func (o *Xdripjs) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *Xdripjs) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *Xdripjs) SetRssi(v float32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *Xdripjs) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetUnfiltered

`func (o *Xdripjs) GetUnfiltered() float32`

GetUnfiltered returns the Unfiltered field if non-nil, zero value otherwise.

### GetUnfilteredOk

`func (o *Xdripjs) GetUnfilteredOk() (*float32, bool)`

GetUnfilteredOk returns a tuple with the Unfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfiltered

`func (o *Xdripjs) SetUnfiltered(v float32)`

SetUnfiltered sets Unfiltered field to given value.

### HasUnfiltered

`func (o *Xdripjs) HasUnfiltered() bool`

HasUnfiltered returns a boolean if a field has been set.

### GetFiltered

`func (o *Xdripjs) GetFiltered() float32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *Xdripjs) GetFilteredOk() (*float32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *Xdripjs) SetFiltered(v float32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *Xdripjs) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetNoise

`func (o *Xdripjs) GetNoise() float32`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *Xdripjs) GetNoiseOk() (*float32, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *Xdripjs) SetNoise(v float32)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *Xdripjs) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetNoiseString

`func (o *Xdripjs) GetNoiseString() float32`

GetNoiseString returns the NoiseString field if non-nil, zero value otherwise.

### GetNoiseStringOk

`func (o *Xdripjs) GetNoiseStringOk() (*float32, bool)`

GetNoiseStringOk returns a tuple with the NoiseString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseString

`func (o *Xdripjs) SetNoiseString(v float32)`

SetNoiseString sets NoiseString field to given value.

### HasNoiseString

`func (o *Xdripjs) HasNoiseString() bool`

HasNoiseString returns a boolean if a field has been set.

### GetSlope

`func (o *Xdripjs) GetSlope() float32`

GetSlope returns the Slope field if non-nil, zero value otherwise.

### GetSlopeOk

`func (o *Xdripjs) GetSlopeOk() (*float32, bool)`

GetSlopeOk returns a tuple with the Slope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlope

`func (o *Xdripjs) SetSlope(v float32)`

SetSlope sets Slope field to given value.

### HasSlope

`func (o *Xdripjs) HasSlope() bool`

HasSlope returns a boolean if a field has been set.

### GetIntercept

`func (o *Xdripjs) GetIntercept() float32`

GetIntercept returns the Intercept field if non-nil, zero value otherwise.

### GetInterceptOk

`func (o *Xdripjs) GetInterceptOk() (*float32, bool)`

GetInterceptOk returns a tuple with the Intercept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntercept

`func (o *Xdripjs) SetIntercept(v float32)`

SetIntercept sets Intercept field to given value.

### HasIntercept

`func (o *Xdripjs) HasIntercept() bool`

HasIntercept returns a boolean if a field has been set.

### GetCalType

`func (o *Xdripjs) GetCalType() string`

GetCalType returns the CalType field if non-nil, zero value otherwise.

### GetCalTypeOk

`func (o *Xdripjs) GetCalTypeOk() (*string, bool)`

GetCalTypeOk returns a tuple with the CalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalType

`func (o *Xdripjs) SetCalType(v string)`

SetCalType sets CalType field to given value.

### HasCalType

`func (o *Xdripjs) HasCalType() bool`

HasCalType returns a boolean if a field has been set.

### GetLastCalibrationDate

`func (o *Xdripjs) GetLastCalibrationDate() float32`

GetLastCalibrationDate returns the LastCalibrationDate field if non-nil, zero value otherwise.

### GetLastCalibrationDateOk

`func (o *Xdripjs) GetLastCalibrationDateOk() (*float32, bool)`

GetLastCalibrationDateOk returns a tuple with the LastCalibrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalibrationDate

`func (o *Xdripjs) SetLastCalibrationDate(v float32)`

SetLastCalibrationDate sets LastCalibrationDate field to given value.

### HasLastCalibrationDate

`func (o *Xdripjs) HasLastCalibrationDate() bool`

HasLastCalibrationDate returns a boolean if a field has been set.

### GetSessionStart

`func (o *Xdripjs) GetSessionStart() float32`

GetSessionStart returns the SessionStart field if non-nil, zero value otherwise.

### GetSessionStartOk

`func (o *Xdripjs) GetSessionStartOk() (*float32, bool)`

GetSessionStartOk returns a tuple with the SessionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStart

`func (o *Xdripjs) SetSessionStart(v float32)`

SetSessionStart sets SessionStart field to given value.

### HasSessionStart

`func (o *Xdripjs) HasSessionStart() bool`

HasSessionStart returns a boolean if a field has been set.

### GetBatteryTimestamp

`func (o *Xdripjs) GetBatteryTimestamp() float32`

GetBatteryTimestamp returns the BatteryTimestamp field if non-nil, zero value otherwise.

### GetBatteryTimestampOk

`func (o *Xdripjs) GetBatteryTimestampOk() (*float32, bool)`

GetBatteryTimestampOk returns a tuple with the BatteryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryTimestamp

`func (o *Xdripjs) SetBatteryTimestamp(v float32)`

SetBatteryTimestamp sets BatteryTimestamp field to given value.

### HasBatteryTimestamp

`func (o *Xdripjs) HasBatteryTimestamp() bool`

HasBatteryTimestamp returns a boolean if a field has been set.

### GetVoltagea

`func (o *Xdripjs) GetVoltagea() float32`

GetVoltagea returns the Voltagea field if non-nil, zero value otherwise.

### GetVoltageaOk

`func (o *Xdripjs) GetVoltageaOk() (*float32, bool)`

GetVoltageaOk returns a tuple with the Voltagea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltagea

`func (o *Xdripjs) SetVoltagea(v float32)`

SetVoltagea sets Voltagea field to given value.

### HasVoltagea

`func (o *Xdripjs) HasVoltagea() bool`

HasVoltagea returns a boolean if a field has been set.

### GetVoltageb

`func (o *Xdripjs) GetVoltageb() float32`

GetVoltageb returns the Voltageb field if non-nil, zero value otherwise.

### GetVoltagebOk

`func (o *Xdripjs) GetVoltagebOk() (*float32, bool)`

GetVoltagebOk returns a tuple with the Voltageb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltageb

`func (o *Xdripjs) SetVoltageb(v float32)`

SetVoltageb sets Voltageb field to given value.

### HasVoltageb

`func (o *Xdripjs) HasVoltageb() bool`

HasVoltageb returns a boolean if a field has been set.

### GetTemperature

`func (o *Xdripjs) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Xdripjs) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Xdripjs) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *Xdripjs) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetResistance

`func (o *Xdripjs) GetResistance() float32`

GetResistance returns the Resistance field if non-nil, zero value otherwise.

### GetResistanceOk

`func (o *Xdripjs) GetResistanceOk() (*float32, bool)`

GetResistanceOk returns a tuple with the Resistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResistance

`func (o *Xdripjs) SetResistance(v float32)`

SetResistance sets Resistance field to given value.

### HasResistance

`func (o *Xdripjs) HasResistance() bool`

HasResistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


