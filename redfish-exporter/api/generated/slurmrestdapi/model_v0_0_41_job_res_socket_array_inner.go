/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.08.0-0rc1&openapi/slurmctld&openapi/v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestdapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the V0041JobResSocketArrayInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobResSocketArrayInner{}

// V0041JobResSocketArrayInner struct for V0041JobResSocketArrayInner
type V0041JobResSocketArrayInner struct {
	// Core index
	Index int32                       `json:"index"`
	Cores []V0041JobResCoreArrayInner `json:"cores"`
}

type _V0041JobResSocketArrayInner V0041JobResSocketArrayInner

// NewV0041JobResSocketArrayInner instantiates a new V0041JobResSocketArrayInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobResSocketArrayInner(index int32, cores []V0041JobResCoreArrayInner) *V0041JobResSocketArrayInner {
	this := V0041JobResSocketArrayInner{}
	this.Index = index
	this.Cores = cores
	return &this
}

// NewV0041JobResSocketArrayInnerWithDefaults instantiates a new V0041JobResSocketArrayInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobResSocketArrayInnerWithDefaults() *V0041JobResSocketArrayInner {
	this := V0041JobResSocketArrayInner{}
	return &this
}

// GetIndex returns the Index field value
func (o *V0041JobResSocketArrayInner) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *V0041JobResSocketArrayInner) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *V0041JobResSocketArrayInner) SetIndex(v int32) {
	o.Index = v
}

// GetCores returns the Cores field value
func (o *V0041JobResSocketArrayInner) GetCores() []V0041JobResCoreArrayInner {
	if o == nil {
		var ret []V0041JobResCoreArrayInner
		return ret
	}

	return o.Cores
}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
func (o *V0041JobResSocketArrayInner) GetCoresOk() ([]V0041JobResCoreArrayInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cores, true
}

// SetCores sets field value
func (o *V0041JobResSocketArrayInner) SetCores(v []V0041JobResCoreArrayInner) {
	o.Cores = v
}

func (o V0041JobResSocketArrayInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobResSocketArrayInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["cores"] = o.Cores
	return toSerialize, nil
}

func (o *V0041JobResSocketArrayInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"cores",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0041JobResSocketArrayInner := _V0041JobResSocketArrayInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041JobResSocketArrayInner)

	if err != nil {
		return err
	}

	*o = V0041JobResSocketArrayInner(varV0041JobResSocketArrayInner)

	return err
}

type NullableV0041JobResSocketArrayInner struct {
	value *V0041JobResSocketArrayInner
	isSet bool
}

func (v NullableV0041JobResSocketArrayInner) Get() *V0041JobResSocketArrayInner {
	return v.value
}

func (v *NullableV0041JobResSocketArrayInner) Set(val *V0041JobResSocketArrayInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobResSocketArrayInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobResSocketArrayInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobResSocketArrayInner(val *V0041JobResSocketArrayInner) *NullableV0041JobResSocketArrayInner {
	return &NullableV0041JobResSocketArrayInner{value: val, isSet: true}
}

func (v NullableV0041JobResSocketArrayInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobResSocketArrayInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
