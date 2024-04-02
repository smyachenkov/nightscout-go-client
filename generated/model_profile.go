/*
Nightscout API

Own your DData with the Nightscout API

API version: 14.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Profile{}

// Profile struct for Profile
type Profile struct {
	// Internally assigned id
	Sens *int32 `json:"sens,omitempty"`
	// Internally assigned id
	Dia *int32 `json:"dia,omitempty"`
	// Internally assigned id
	Carbratio *int32 `json:"carbratio,omitempty"`
	// Internally assigned id
	CarbsHr *int32 `json:"carbs_hr,omitempty"`
	// Internally assigned id
	Id *string `json:"_id,omitempty"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile() *Profile {
	this := Profile{}
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetSens returns the Sens field value if set, zero value otherwise.
func (o *Profile) GetSens() int32 {
	if o == nil || IsNil(o.Sens) {
		var ret int32
		return ret
	}
	return *o.Sens
}

// GetSensOk returns a tuple with the Sens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetSensOk() (*int32, bool) {
	if o == nil || IsNil(o.Sens) {
		return nil, false
	}
	return o.Sens, true
}

// HasSens returns a boolean if a field has been set.
func (o *Profile) HasSens() bool {
	if o != nil && !IsNil(o.Sens) {
		return true
	}

	return false
}

// SetSens gets a reference to the given int32 and assigns it to the Sens field.
func (o *Profile) SetSens(v int32) {
	o.Sens = &v
}

// GetDia returns the Dia field value if set, zero value otherwise.
func (o *Profile) GetDia() int32 {
	if o == nil || IsNil(o.Dia) {
		var ret int32
		return ret
	}
	return *o.Dia
}

// GetDiaOk returns a tuple with the Dia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetDiaOk() (*int32, bool) {
	if o == nil || IsNil(o.Dia) {
		return nil, false
	}
	return o.Dia, true
}

// HasDia returns a boolean if a field has been set.
func (o *Profile) HasDia() bool {
	if o != nil && !IsNil(o.Dia) {
		return true
	}

	return false
}

// SetDia gets a reference to the given int32 and assigns it to the Dia field.
func (o *Profile) SetDia(v int32) {
	o.Dia = &v
}

// GetCarbratio returns the Carbratio field value if set, zero value otherwise.
func (o *Profile) GetCarbratio() int32 {
	if o == nil || IsNil(o.Carbratio) {
		var ret int32
		return ret
	}
	return *o.Carbratio
}

// GetCarbratioOk returns a tuple with the Carbratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetCarbratioOk() (*int32, bool) {
	if o == nil || IsNil(o.Carbratio) {
		return nil, false
	}
	return o.Carbratio, true
}

// HasCarbratio returns a boolean if a field has been set.
func (o *Profile) HasCarbratio() bool {
	if o != nil && !IsNil(o.Carbratio) {
		return true
	}

	return false
}

// SetCarbratio gets a reference to the given int32 and assigns it to the Carbratio field.
func (o *Profile) SetCarbratio(v int32) {
	o.Carbratio = &v
}

// GetCarbsHr returns the CarbsHr field value if set, zero value otherwise.
func (o *Profile) GetCarbsHr() int32 {
	if o == nil || IsNil(o.CarbsHr) {
		var ret int32
		return ret
	}
	return *o.CarbsHr
}

// GetCarbsHrOk returns a tuple with the CarbsHr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetCarbsHrOk() (*int32, bool) {
	if o == nil || IsNil(o.CarbsHr) {
		return nil, false
	}
	return o.CarbsHr, true
}

// HasCarbsHr returns a boolean if a field has been set.
func (o *Profile) HasCarbsHr() bool {
	if o != nil && !IsNil(o.CarbsHr) {
		return true
	}

	return false
}

// SetCarbsHr gets a reference to the given int32 and assigns it to the CarbsHr field.
func (o *Profile) SetCarbsHr(v int32) {
	o.CarbsHr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Profile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Profile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Profile) SetId(v string) {
	o.Id = &v
}

func (o Profile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sens) {
		toSerialize["sens"] = o.Sens
	}
	if !IsNil(o.Dia) {
		toSerialize["dia"] = o.Dia
	}
	if !IsNil(o.Carbratio) {
		toSerialize["carbratio"] = o.Carbratio
	}
	if !IsNil(o.CarbsHr) {
		toSerialize["carbs_hr"] = o.CarbsHr
	}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	return toSerialize, nil
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

