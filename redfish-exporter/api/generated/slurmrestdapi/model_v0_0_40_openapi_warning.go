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

// checks if the V0040OpenapiWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiWarning{}

// V0040OpenapiWarning struct for V0040OpenapiWarning
type V0040OpenapiWarning struct {
	// Long form warning description
	Description *string `json:"description,omitempty"`
	// Source of warning or where warning was first detected
	Source *string `json:"source,omitempty"`
}

// NewV0040OpenapiWarning instantiates a new V0040OpenapiWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiWarning() *V0040OpenapiWarning {
	this := V0040OpenapiWarning{}
	return &this
}

// NewV0040OpenapiWarningWithDefaults instantiates a new V0040OpenapiWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiWarningWithDefaults() *V0040OpenapiWarning {
	this := V0040OpenapiWarning{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V0040OpenapiWarning) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWarning) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V0040OpenapiWarning) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V0040OpenapiWarning) SetDescription(v string) {
	o.Description = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V0040OpenapiWarning) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiWarning) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V0040OpenapiWarning) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V0040OpenapiWarning) SetSource(v string) {
	o.Source = &v
}

func (o V0040OpenapiWarning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableV0040OpenapiWarning struct {
	value *V0040OpenapiWarning
	isSet bool
}

func (v NullableV0040OpenapiWarning) Get() *V0040OpenapiWarning {
	return v.value
}

func (v *NullableV0040OpenapiWarning) Set(val *V0040OpenapiWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiWarning(val *V0040OpenapiWarning) *NullableV0040OpenapiWarning {
	return &NullableV0040OpenapiWarning{value: val, isSet: true}
}

func (v NullableV0040OpenapiWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
