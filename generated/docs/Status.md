# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiEnabled** | Pointer to **bool** | Whether or not the REST API is enabled. | [optional] 
**CareportalEnabled** | Pointer to **bool** | Whether or not the careportal is enabled in the API. | [optional] 
**Head** | Pointer to **string** | The git identifier for the running instance of the app. | [optional] 
**Name** | Pointer to **string** | Nightscout (static) | [optional] 
**Version** | Pointer to **string** | The version label of the app. | [optional] 
**Settings** | Pointer to [**Settings**](Settings.md) |  | [optional] 
**ExtendedSettings** | Pointer to **interface{}** | Extended settings of client side plugins | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiEnabled

`func (o *Status) GetApiEnabled() bool`

GetApiEnabled returns the ApiEnabled field if non-nil, zero value otherwise.

### GetApiEnabledOk

`func (o *Status) GetApiEnabledOk() (*bool, bool)`

GetApiEnabledOk returns a tuple with the ApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEnabled

`func (o *Status) SetApiEnabled(v bool)`

SetApiEnabled sets ApiEnabled field to given value.

### HasApiEnabled

`func (o *Status) HasApiEnabled() bool`

HasApiEnabled returns a boolean if a field has been set.

### GetCareportalEnabled

`func (o *Status) GetCareportalEnabled() bool`

GetCareportalEnabled returns the CareportalEnabled field if non-nil, zero value otherwise.

### GetCareportalEnabledOk

`func (o *Status) GetCareportalEnabledOk() (*bool, bool)`

GetCareportalEnabledOk returns a tuple with the CareportalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareportalEnabled

`func (o *Status) SetCareportalEnabled(v bool)`

SetCareportalEnabled sets CareportalEnabled field to given value.

### HasCareportalEnabled

`func (o *Status) HasCareportalEnabled() bool`

HasCareportalEnabled returns a boolean if a field has been set.

### GetHead

`func (o *Status) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *Status) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *Status) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *Status) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetName

`func (o *Status) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Status) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Status) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Status) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Status) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Status) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Status) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Status) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSettings

`func (o *Status) GetSettings() Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Status) GetSettingsOk() (*Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Status) SetSettings(v Settings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Status) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetExtendedSettings

`func (o *Status) GetExtendedSettings() interface{}`

GetExtendedSettings returns the ExtendedSettings field if non-nil, zero value otherwise.

### GetExtendedSettingsOk

`func (o *Status) GetExtendedSettingsOk() (*interface{}, bool)`

GetExtendedSettingsOk returns a tuple with the ExtendedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSettings

`func (o *Status) SetExtendedSettings(v interface{})`

SetExtendedSettings sets ExtendedSettings field to given value.

### HasExtendedSettings

`func (o *Status) HasExtendedSettings() bool`

HasExtendedSettings returns a boolean if a field has been set.

### SetExtendedSettingsNil

`func (o *Status) SetExtendedSettingsNil(b bool)`

 SetExtendedSettingsNil sets the value for ExtendedSettings to be an explicit nil

### UnsetExtendedSettings
`func (o *Status) UnsetExtendedSettings()`

UnsetExtendedSettings ensures that no value is present for ExtendedSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


