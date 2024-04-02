# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **string** | Default units for glucose measurements across the server. | [optional] 
**TimeFormat** | Pointer to **string** | Default time format | [optional] 
**CustomTitle** | Pointer to **string** | Default custom title to be displayed system wide. | [optional] 
**NightMode** | Pointer to **bool** | Should Night mode be enabled by default? | [optional] 
**Theme** | Pointer to **string** | Default theme to be displayed system wide, &#x60;default&#x60;, &#x60;colors&#x60;, &#x60;colorblindfriendly&#x60;. | [optional] 
**Language** | Pointer to **string** | Default language code to be used system wide | [optional] 
**ShowPlugins** | Pointer to **string** | Plugins that should be shown by default | [optional] 
**ShowRawbg** | Pointer to **string** | If Raw BG is enabled when should it be shown? &#x60;never&#x60;, &#x60;always&#x60;, &#x60;noise&#x60; | [optional] 
**AlarmTypes** | Pointer to **[]string** | Enabled alarm types, can be multiple | [optional] 
**AlarmUrgentHigh** | Pointer to **bool** | Enable/Disable client-side Urgent High alarms by default, for use with &#x60;simple&#x60; alarms. | [optional] 
**AlarmHigh** | Pointer to **bool** | Enable/Disable client-side High alarms by default, for use with &#x60;simple&#x60; alarms. | [optional] 
**AlarmLow** | Pointer to **bool** | Enable/Disable client-side Low alarms by default, for use with &#x60;simple&#x60; alarms. | [optional] 
**AlarmUrgentLow** | Pointer to **bool** | Enable/Disable client-side Urgent Low alarms by default, for use with &#x60;simple&#x60; alarms. | [optional] 
**AlarmTimeagoWarn** | Pointer to **bool** | Enable/Disable client-side stale data alarms by default. | [optional] 
**AlarmTimeagoWarnMins** | Pointer to **float32** | Number of minutes before a stale data warning is generated. | [optional] 
**AlarmTimeagoUrgent** | Pointer to **bool** | Enable/Disable client-side urgent stale data alarms by default. | [optional] 
**AlarmTimeagoUrgentMins** | Pointer to **float32** | Number of minutes before a stale data warning is generated. | [optional] 
**Enable** | Pointer to **[]string** | Enabled features | [optional] 
**Thresholds** | Pointer to [**Threshold**](Threshold.md) |  | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *Settings) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Settings) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Settings) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Settings) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetTimeFormat

`func (o *Settings) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *Settings) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *Settings) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *Settings) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetCustomTitle

`func (o *Settings) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *Settings) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *Settings) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *Settings) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.

### GetNightMode

`func (o *Settings) GetNightMode() bool`

GetNightMode returns the NightMode field if non-nil, zero value otherwise.

### GetNightModeOk

`func (o *Settings) GetNightModeOk() (*bool, bool)`

GetNightModeOk returns a tuple with the NightMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightMode

`func (o *Settings) SetNightMode(v bool)`

SetNightMode sets NightMode field to given value.

### HasNightMode

`func (o *Settings) HasNightMode() bool`

HasNightMode returns a boolean if a field has been set.

### GetTheme

`func (o *Settings) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Settings) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Settings) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Settings) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetLanguage

`func (o *Settings) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Settings) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Settings) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Settings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetShowPlugins

`func (o *Settings) GetShowPlugins() string`

GetShowPlugins returns the ShowPlugins field if non-nil, zero value otherwise.

### GetShowPluginsOk

`func (o *Settings) GetShowPluginsOk() (*string, bool)`

GetShowPluginsOk returns a tuple with the ShowPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPlugins

`func (o *Settings) SetShowPlugins(v string)`

SetShowPlugins sets ShowPlugins field to given value.

### HasShowPlugins

`func (o *Settings) HasShowPlugins() bool`

HasShowPlugins returns a boolean if a field has been set.

### GetShowRawbg

`func (o *Settings) GetShowRawbg() string`

GetShowRawbg returns the ShowRawbg field if non-nil, zero value otherwise.

### GetShowRawbgOk

`func (o *Settings) GetShowRawbgOk() (*string, bool)`

GetShowRawbgOk returns a tuple with the ShowRawbg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRawbg

`func (o *Settings) SetShowRawbg(v string)`

SetShowRawbg sets ShowRawbg field to given value.

### HasShowRawbg

`func (o *Settings) HasShowRawbg() bool`

HasShowRawbg returns a boolean if a field has been set.

### GetAlarmTypes

`func (o *Settings) GetAlarmTypes() []string`

GetAlarmTypes returns the AlarmTypes field if non-nil, zero value otherwise.

### GetAlarmTypesOk

`func (o *Settings) GetAlarmTypesOk() (*[]string, bool)`

GetAlarmTypesOk returns a tuple with the AlarmTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTypes

`func (o *Settings) SetAlarmTypes(v []string)`

SetAlarmTypes sets AlarmTypes field to given value.

### HasAlarmTypes

`func (o *Settings) HasAlarmTypes() bool`

HasAlarmTypes returns a boolean if a field has been set.

### GetAlarmUrgentHigh

`func (o *Settings) GetAlarmUrgentHigh() bool`

GetAlarmUrgentHigh returns the AlarmUrgentHigh field if non-nil, zero value otherwise.

### GetAlarmUrgentHighOk

`func (o *Settings) GetAlarmUrgentHighOk() (*bool, bool)`

GetAlarmUrgentHighOk returns a tuple with the AlarmUrgentHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmUrgentHigh

`func (o *Settings) SetAlarmUrgentHigh(v bool)`

SetAlarmUrgentHigh sets AlarmUrgentHigh field to given value.

### HasAlarmUrgentHigh

`func (o *Settings) HasAlarmUrgentHigh() bool`

HasAlarmUrgentHigh returns a boolean if a field has been set.

### GetAlarmHigh

`func (o *Settings) GetAlarmHigh() bool`

GetAlarmHigh returns the AlarmHigh field if non-nil, zero value otherwise.

### GetAlarmHighOk

`func (o *Settings) GetAlarmHighOk() (*bool, bool)`

GetAlarmHighOk returns a tuple with the AlarmHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmHigh

`func (o *Settings) SetAlarmHigh(v bool)`

SetAlarmHigh sets AlarmHigh field to given value.

### HasAlarmHigh

`func (o *Settings) HasAlarmHigh() bool`

HasAlarmHigh returns a boolean if a field has been set.

### GetAlarmLow

`func (o *Settings) GetAlarmLow() bool`

GetAlarmLow returns the AlarmLow field if non-nil, zero value otherwise.

### GetAlarmLowOk

`func (o *Settings) GetAlarmLowOk() (*bool, bool)`

GetAlarmLowOk returns a tuple with the AlarmLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmLow

`func (o *Settings) SetAlarmLow(v bool)`

SetAlarmLow sets AlarmLow field to given value.

### HasAlarmLow

`func (o *Settings) HasAlarmLow() bool`

HasAlarmLow returns a boolean if a field has been set.

### GetAlarmUrgentLow

`func (o *Settings) GetAlarmUrgentLow() bool`

GetAlarmUrgentLow returns the AlarmUrgentLow field if non-nil, zero value otherwise.

### GetAlarmUrgentLowOk

`func (o *Settings) GetAlarmUrgentLowOk() (*bool, bool)`

GetAlarmUrgentLowOk returns a tuple with the AlarmUrgentLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmUrgentLow

`func (o *Settings) SetAlarmUrgentLow(v bool)`

SetAlarmUrgentLow sets AlarmUrgentLow field to given value.

### HasAlarmUrgentLow

`func (o *Settings) HasAlarmUrgentLow() bool`

HasAlarmUrgentLow returns a boolean if a field has been set.

### GetAlarmTimeagoWarn

`func (o *Settings) GetAlarmTimeagoWarn() bool`

GetAlarmTimeagoWarn returns the AlarmTimeagoWarn field if non-nil, zero value otherwise.

### GetAlarmTimeagoWarnOk

`func (o *Settings) GetAlarmTimeagoWarnOk() (*bool, bool)`

GetAlarmTimeagoWarnOk returns a tuple with the AlarmTimeagoWarn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTimeagoWarn

`func (o *Settings) SetAlarmTimeagoWarn(v bool)`

SetAlarmTimeagoWarn sets AlarmTimeagoWarn field to given value.

### HasAlarmTimeagoWarn

`func (o *Settings) HasAlarmTimeagoWarn() bool`

HasAlarmTimeagoWarn returns a boolean if a field has been set.

### GetAlarmTimeagoWarnMins

`func (o *Settings) GetAlarmTimeagoWarnMins() float32`

GetAlarmTimeagoWarnMins returns the AlarmTimeagoWarnMins field if non-nil, zero value otherwise.

### GetAlarmTimeagoWarnMinsOk

`func (o *Settings) GetAlarmTimeagoWarnMinsOk() (*float32, bool)`

GetAlarmTimeagoWarnMinsOk returns a tuple with the AlarmTimeagoWarnMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTimeagoWarnMins

`func (o *Settings) SetAlarmTimeagoWarnMins(v float32)`

SetAlarmTimeagoWarnMins sets AlarmTimeagoWarnMins field to given value.

### HasAlarmTimeagoWarnMins

`func (o *Settings) HasAlarmTimeagoWarnMins() bool`

HasAlarmTimeagoWarnMins returns a boolean if a field has been set.

### GetAlarmTimeagoUrgent

`func (o *Settings) GetAlarmTimeagoUrgent() bool`

GetAlarmTimeagoUrgent returns the AlarmTimeagoUrgent field if non-nil, zero value otherwise.

### GetAlarmTimeagoUrgentOk

`func (o *Settings) GetAlarmTimeagoUrgentOk() (*bool, bool)`

GetAlarmTimeagoUrgentOk returns a tuple with the AlarmTimeagoUrgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTimeagoUrgent

`func (o *Settings) SetAlarmTimeagoUrgent(v bool)`

SetAlarmTimeagoUrgent sets AlarmTimeagoUrgent field to given value.

### HasAlarmTimeagoUrgent

`func (o *Settings) HasAlarmTimeagoUrgent() bool`

HasAlarmTimeagoUrgent returns a boolean if a field has been set.

### GetAlarmTimeagoUrgentMins

`func (o *Settings) GetAlarmTimeagoUrgentMins() float32`

GetAlarmTimeagoUrgentMins returns the AlarmTimeagoUrgentMins field if non-nil, zero value otherwise.

### GetAlarmTimeagoUrgentMinsOk

`func (o *Settings) GetAlarmTimeagoUrgentMinsOk() (*float32, bool)`

GetAlarmTimeagoUrgentMinsOk returns a tuple with the AlarmTimeagoUrgentMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTimeagoUrgentMins

`func (o *Settings) SetAlarmTimeagoUrgentMins(v float32)`

SetAlarmTimeagoUrgentMins sets AlarmTimeagoUrgentMins field to given value.

### HasAlarmTimeagoUrgentMins

`func (o *Settings) HasAlarmTimeagoUrgentMins() bool`

HasAlarmTimeagoUrgentMins returns a boolean if a field has been set.

### GetEnable

`func (o *Settings) GetEnable() []string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Settings) GetEnableOk() (*[]string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Settings) SetEnable(v []string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Settings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetThresholds

`func (o *Settings) GetThresholds() Threshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Settings) GetThresholdsOk() (*Threshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Settings) SetThresholds(v Threshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Settings) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


