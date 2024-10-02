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

// checks if the V0041AssocSharesObjListInnerFairshare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041AssocSharesObjListInnerFairshare{}

// V0041AssocSharesObjListInnerFairshare struct for V0041AssocSharesObjListInnerFairshare
type V0041AssocSharesObjListInnerFairshare struct {
	// fairshare factor
	Factor *float64 `json:"factor,omitempty"`
	// fairshare factor at this level. stored on an assoc as a long double, but that is not needed for display in sshare
	Level *float64 `json:"level,omitempty"`
}

// NewV0041AssocSharesObjListInnerFairshare instantiates a new V0041AssocSharesObjListInnerFairshare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041AssocSharesObjListInnerFairshare() *V0041AssocSharesObjListInnerFairshare {
	this := V0041AssocSharesObjListInnerFairshare{}
	return &this
}

// NewV0041AssocSharesObjListInnerFairshareWithDefaults instantiates a new V0041AssocSharesObjListInnerFairshare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041AssocSharesObjListInnerFairshareWithDefaults() *V0041AssocSharesObjListInnerFairshare {
	this := V0041AssocSharesObjListInnerFairshare{}
	return &this
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInnerFairshare) GetFactor() float64 {
	if o == nil || IsNil(o.Factor) {
		var ret float64
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInnerFairshare) GetFactorOk() (*float64, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInnerFairshare) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given float64 and assigns it to the Factor field.
func (o *V0041AssocSharesObjListInnerFairshare) SetFactor(v float64) {
	o.Factor = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInnerFairshare) GetLevel() float64 {
	if o == nil || IsNil(o.Level) {
		var ret float64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInnerFairshare) GetLevelOk() (*float64, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInnerFairshare) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float64 and assigns it to the Level field.
func (o *V0041AssocSharesObjListInnerFairshare) SetLevel(v float64) {
	o.Level = &v
}

func (o V0041AssocSharesObjListInnerFairshare) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041AssocSharesObjListInnerFairshare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableV0041AssocSharesObjListInnerFairshare struct {
	value *V0041AssocSharesObjListInnerFairshare
	isSet bool
}

func (v NullableV0041AssocSharesObjListInnerFairshare) Get() *V0041AssocSharesObjListInnerFairshare {
	return v.value
}

func (v *NullableV0041AssocSharesObjListInnerFairshare) Set(val *V0041AssocSharesObjListInnerFairshare) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041AssocSharesObjListInnerFairshare) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041AssocSharesObjListInnerFairshare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041AssocSharesObjListInnerFairshare(val *V0041AssocSharesObjListInnerFairshare) *NullableV0041AssocSharesObjListInnerFairshare {
	return &NullableV0041AssocSharesObjListInnerFairshare{value: val, isSet: true}
}

func (v NullableV0041AssocSharesObjListInnerFairshare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041AssocSharesObjListInnerFairshare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
