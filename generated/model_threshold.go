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

// checks if the Threshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Threshold{}

// Threshold struct for Threshold
type Threshold struct {
	// High BG range.
	BgHigh *int32 `json:"bg_high,omitempty"`
	// Top of target range.
	BgTargetTop *int32 `json:"bg_target_top,omitempty"`
	// Bottom of target range.
	BgTargetBottom *int32 `json:"bg_target_bottom,omitempty"`
	// Low BG range.
	BgLow *int32 `json:"bg_low,omitempty"`
}

// NewThreshold instantiates a new Threshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreshold() *Threshold {
	this := Threshold{}
	return &this
}

// NewThresholdWithDefaults instantiates a new Threshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdWithDefaults() *Threshold {
	this := Threshold{}
	return &this
}

// GetBgHigh returns the BgHigh field value if set, zero value otherwise.
func (o *Threshold) GetBgHigh() int32 {
	if o == nil || IsNil(o.BgHigh) {
		var ret int32
		return ret
	}
	return *o.BgHigh
}

// GetBgHighOk returns a tuple with the BgHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetBgHighOk() (*int32, bool) {
	if o == nil || IsNil(o.BgHigh) {
		return nil, false
	}
	return o.BgHigh, true
}

// HasBgHigh returns a boolean if a field has been set.
func (o *Threshold) HasBgHigh() bool {
	if o != nil && !IsNil(o.BgHigh) {
		return true
	}

	return false
}

// SetBgHigh gets a reference to the given int32 and assigns it to the BgHigh field.
func (o *Threshold) SetBgHigh(v int32) {
	o.BgHigh = &v
}

// GetBgTargetTop returns the BgTargetTop field value if set, zero value otherwise.
func (o *Threshold) GetBgTargetTop() int32 {
	if o == nil || IsNil(o.BgTargetTop) {
		var ret int32
		return ret
	}
	return *o.BgTargetTop
}

// GetBgTargetTopOk returns a tuple with the BgTargetTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetBgTargetTopOk() (*int32, bool) {
	if o == nil || IsNil(o.BgTargetTop) {
		return nil, false
	}
	return o.BgTargetTop, true
}

// HasBgTargetTop returns a boolean if a field has been set.
func (o *Threshold) HasBgTargetTop() bool {
	if o != nil && !IsNil(o.BgTargetTop) {
		return true
	}

	return false
}

// SetBgTargetTop gets a reference to the given int32 and assigns it to the BgTargetTop field.
func (o *Threshold) SetBgTargetTop(v int32) {
	o.BgTargetTop = &v
}

// GetBgTargetBottom returns the BgTargetBottom field value if set, zero value otherwise.
func (o *Threshold) GetBgTargetBottom() int32 {
	if o == nil || IsNil(o.BgTargetBottom) {
		var ret int32
		return ret
	}
	return *o.BgTargetBottom
}

// GetBgTargetBottomOk returns a tuple with the BgTargetBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetBgTargetBottomOk() (*int32, bool) {
	if o == nil || IsNil(o.BgTargetBottom) {
		return nil, false
	}
	return o.BgTargetBottom, true
}

// HasBgTargetBottom returns a boolean if a field has been set.
func (o *Threshold) HasBgTargetBottom() bool {
	if o != nil && !IsNil(o.BgTargetBottom) {
		return true
	}

	return false
}

// SetBgTargetBottom gets a reference to the given int32 and assigns it to the BgTargetBottom field.
func (o *Threshold) SetBgTargetBottom(v int32) {
	o.BgTargetBottom = &v
}

// GetBgLow returns the BgLow field value if set, zero value otherwise.
func (o *Threshold) GetBgLow() int32 {
	if o == nil || IsNil(o.BgLow) {
		var ret int32
		return ret
	}
	return *o.BgLow
}

// GetBgLowOk returns a tuple with the BgLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetBgLowOk() (*int32, bool) {
	if o == nil || IsNil(o.BgLow) {
		return nil, false
	}
	return o.BgLow, true
}

// HasBgLow returns a boolean if a field has been set.
func (o *Threshold) HasBgLow() bool {
	if o != nil && !IsNil(o.BgLow) {
		return true
	}

	return false
}

// SetBgLow gets a reference to the given int32 and assigns it to the BgLow field.
func (o *Threshold) SetBgLow(v int32) {
	o.BgLow = &v
}

func (o Threshold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Threshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BgHigh) {
		toSerialize["bg_high"] = o.BgHigh
	}
	if !IsNil(o.BgTargetTop) {
		toSerialize["bg_target_top"] = o.BgTargetTop
	}
	if !IsNil(o.BgTargetBottom) {
		toSerialize["bg_target_bottom"] = o.BgTargetBottom
	}
	if !IsNil(o.BgLow) {
		toSerialize["bg_low"] = o.BgLow
	}
	return toSerialize, nil
}

type NullableThreshold struct {
	value *Threshold
	isSet bool
}

func (v NullableThreshold) Get() *Threshold {
	return v.value
}

func (v *NullableThreshold) Set(val *Threshold) {
	v.value = val
	v.isSet = true
}

func (v NullableThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreshold(val *Threshold) *NullableThreshold {
	return &NullableThreshold{value: val, isSet: true}
}

func (v NullableThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


