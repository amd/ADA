/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.08.0-0rc1&openapi/slurmctld&openapi/v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestdapi

import (
	"encoding/json"
)

// checks if the V0039MetaSlurmVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039MetaSlurmVersion{}

// V0039MetaSlurmVersion struct for V0039MetaSlurmVersion
type V0039MetaSlurmVersion struct {
	//
	Major *int32 `json:"major,omitempty"`
	//
	Micro *int32 `json:"micro,omitempty"`
	//
	Minor *int32 `json:"minor,omitempty"`
}

// NewV0039MetaSlurmVersion instantiates a new V0039MetaSlurmVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039MetaSlurmVersion() *V0039MetaSlurmVersion {
	this := V0039MetaSlurmVersion{}
	return &this
}

// NewV0039MetaSlurmVersionWithDefaults instantiates a new V0039MetaSlurmVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039MetaSlurmVersionWithDefaults() *V0039MetaSlurmVersion {
	this := V0039MetaSlurmVersion{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *V0039MetaSlurmVersion) GetMajor() int32 {
	if o == nil || IsNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039MetaSlurmVersion) GetMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *V0039MetaSlurmVersion) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *V0039MetaSlurmVersion) SetMajor(v int32) {
	o.Major = &v
}

// GetMicro returns the Micro field value if set, zero value otherwise.
func (o *V0039MetaSlurmVersion) GetMicro() int32 {
	if o == nil || IsNil(o.Micro) {
		var ret int32
		return ret
	}
	return *o.Micro
}

// GetMicroOk returns a tuple with the Micro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039MetaSlurmVersion) GetMicroOk() (*int32, bool) {
	if o == nil || IsNil(o.Micro) {
		return nil, false
	}
	return o.Micro, true
}

// HasMicro returns a boolean if a field has been set.
func (o *V0039MetaSlurmVersion) HasMicro() bool {
	if o != nil && !IsNil(o.Micro) {
		return true
	}

	return false
}

// SetMicro gets a reference to the given int32 and assigns it to the Micro field.
func (o *V0039MetaSlurmVersion) SetMicro(v int32) {
	o.Micro = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *V0039MetaSlurmVersion) GetMinor() int32 {
	if o == nil || IsNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039MetaSlurmVersion) GetMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *V0039MetaSlurmVersion) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *V0039MetaSlurmVersion) SetMinor(v int32) {
	o.Minor = &v
}

func (o V0039MetaSlurmVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039MetaSlurmVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Micro) {
		toSerialize["micro"] = o.Micro
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return toSerialize, nil
}

type NullableV0039MetaSlurmVersion struct {
	value *V0039MetaSlurmVersion
	isSet bool
}

func (v NullableV0039MetaSlurmVersion) Get() *V0039MetaSlurmVersion {
	return v.value
}

func (v *NullableV0039MetaSlurmVersion) Set(val *V0039MetaSlurmVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039MetaSlurmVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039MetaSlurmVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039MetaSlurmVersion(val *V0039MetaSlurmVersion) *NullableV0039MetaSlurmVersion {
	return &NullableV0039MetaSlurmVersion{value: val, isSet: true}
}

func (v NullableV0039MetaSlurmVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039MetaSlurmVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
