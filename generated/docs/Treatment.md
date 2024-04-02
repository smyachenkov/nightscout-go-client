# Treatment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internally assigned id. | [optional] 
**EventType** | Pointer to **string** | The type of treatment event. | [optional] 
**CreatedAt** | Pointer to **string** | The date of the event, might be set retroactively . | [optional] 
**Glucose** | Pointer to **string** | Current glucose. | [optional] 
**GlucoseType** | Pointer to **string** | Method used to obtain glucose, Finger or Sensor. | [optional] 
**Carbs** | Pointer to **float32** | Amount of carbs consumed in grams. | [optional] 
**Protein** | Pointer to **float32** | Amount of protein consumed in grams. | [optional] 
**Fat** | Pointer to **float32** | Amount of fat consumed in grams. | [optional] 
**Insulin** | Pointer to **float32** | Amount of insulin, if any. | [optional] 
**Units** | Pointer to **string** | The units for the glucose value, mg/dl or mmol. | [optional] 
**TransmitterId** | Pointer to **string** | The transmitter ID of the transmitter being started. | [optional] 
**SensorCode** | Pointer to **string** | The code used to start a Dexcom G6 sensor. | [optional] 
**Notes** | Pointer to **string** | Description/notes of treatment. | [optional] 
**EnteredBy** | Pointer to **string** | Who entered the treatment. | [optional] 

## Methods

### NewTreatment

`func NewTreatment() *Treatment`

NewTreatment instantiates a new Treatment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreatmentWithDefaults

`func NewTreatmentWithDefaults() *Treatment`

NewTreatmentWithDefaults instantiates a new Treatment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Treatment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Treatment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Treatment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Treatment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *Treatment) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Treatment) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Treatment) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Treatment) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Treatment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Treatment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Treatment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Treatment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGlucose

`func (o *Treatment) GetGlucose() string`

GetGlucose returns the Glucose field if non-nil, zero value otherwise.

### GetGlucoseOk

`func (o *Treatment) GetGlucoseOk() (*string, bool)`

GetGlucoseOk returns a tuple with the Glucose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlucose

`func (o *Treatment) SetGlucose(v string)`

SetGlucose sets Glucose field to given value.

### HasGlucose

`func (o *Treatment) HasGlucose() bool`

HasGlucose returns a boolean if a field has been set.

### GetGlucoseType

`func (o *Treatment) GetGlucoseType() string`

GetGlucoseType returns the GlucoseType field if non-nil, zero value otherwise.

### GetGlucoseTypeOk

`func (o *Treatment) GetGlucoseTypeOk() (*string, bool)`

GetGlucoseTypeOk returns a tuple with the GlucoseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlucoseType

`func (o *Treatment) SetGlucoseType(v string)`

SetGlucoseType sets GlucoseType field to given value.

### HasGlucoseType

`func (o *Treatment) HasGlucoseType() bool`

HasGlucoseType returns a boolean if a field has been set.

### GetCarbs

`func (o *Treatment) GetCarbs() float32`

GetCarbs returns the Carbs field if non-nil, zero value otherwise.

### GetCarbsOk

`func (o *Treatment) GetCarbsOk() (*float32, bool)`

GetCarbsOk returns a tuple with the Carbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarbs

`func (o *Treatment) SetCarbs(v float32)`

SetCarbs sets Carbs field to given value.

### HasCarbs

`func (o *Treatment) HasCarbs() bool`

HasCarbs returns a boolean if a field has been set.

### GetProtein

`func (o *Treatment) GetProtein() float32`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *Treatment) GetProteinOk() (*float32, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *Treatment) SetProtein(v float32)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *Treatment) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetFat

`func (o *Treatment) GetFat() float32`

GetFat returns the Fat field if non-nil, zero value otherwise.

### GetFatOk

`func (o *Treatment) GetFatOk() (*float32, bool)`

GetFatOk returns a tuple with the Fat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFat

`func (o *Treatment) SetFat(v float32)`

SetFat sets Fat field to given value.

### HasFat

`func (o *Treatment) HasFat() bool`

HasFat returns a boolean if a field has been set.

### GetInsulin

`func (o *Treatment) GetInsulin() float32`

GetInsulin returns the Insulin field if non-nil, zero value otherwise.

### GetInsulinOk

`func (o *Treatment) GetInsulinOk() (*float32, bool)`

GetInsulinOk returns a tuple with the Insulin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsulin

`func (o *Treatment) SetInsulin(v float32)`

SetInsulin sets Insulin field to given value.

### HasInsulin

`func (o *Treatment) HasInsulin() bool`

HasInsulin returns a boolean if a field has been set.

### GetUnits

`func (o *Treatment) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Treatment) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Treatment) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Treatment) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetTransmitterId

`func (o *Treatment) GetTransmitterId() string`

GetTransmitterId returns the TransmitterId field if non-nil, zero value otherwise.

### GetTransmitterIdOk

`func (o *Treatment) GetTransmitterIdOk() (*string, bool)`

GetTransmitterIdOk returns a tuple with the TransmitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterId

`func (o *Treatment) SetTransmitterId(v string)`

SetTransmitterId sets TransmitterId field to given value.

### HasTransmitterId

`func (o *Treatment) HasTransmitterId() bool`

HasTransmitterId returns a boolean if a field has been set.

### GetSensorCode

`func (o *Treatment) GetSensorCode() string`

GetSensorCode returns the SensorCode field if non-nil, zero value otherwise.

### GetSensorCodeOk

`func (o *Treatment) GetSensorCodeOk() (*string, bool)`

GetSensorCodeOk returns a tuple with the SensorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorCode

`func (o *Treatment) SetSensorCode(v string)`

SetSensorCode sets SensorCode field to given value.

### HasSensorCode

`func (o *Treatment) HasSensorCode() bool`

HasSensorCode returns a boolean if a field has been set.

### GetNotes

`func (o *Treatment) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Treatment) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Treatment) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Treatment) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEnteredBy

`func (o *Treatment) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *Treatment) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *Treatment) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *Treatment) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


