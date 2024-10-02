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

// checks if the V0040OpenapiDiagResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiDiagResp{}

// V0040OpenapiDiagResp struct for V0040OpenapiDiagResp
type V0040OpenapiDiagResp struct {
	Statistics V0040StatsMsg         `json:"statistics"`
	Meta       *V0040OpenapiMeta     `json:"meta,omitempty"`
	Errors     []V0040OpenapiError   `json:"errors,omitempty"`
	Warnings   []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiDiagResp V0040OpenapiDiagResp

// NewV0040OpenapiDiagResp instantiates a new V0040OpenapiDiagResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiDiagResp(statistics V0040StatsMsg) *V0040OpenapiDiagResp {
	this := V0040OpenapiDiagResp{}
	this.Statistics = statistics
	return &this
}

// NewV0040OpenapiDiagRespWithDefaults instantiates a new V0040OpenapiDiagResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiDiagRespWithDefaults() *V0040OpenapiDiagResp {
	this := V0040OpenapiDiagResp{}
	return &this
}

// GetStatistics returns the Statistics field value
func (o *V0040OpenapiDiagResp) GetStatistics() V0040StatsMsg {
	if o == nil {
		var ret V0040StatsMsg
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiDiagResp) GetStatisticsOk() (*V0040StatsMsg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// SetStatistics sets field value
func (o *V0040OpenapiDiagResp) SetStatistics(v V0040StatsMsg) {
	o.Statistics = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiDiagResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiDiagResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiDiagResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiDiagResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiDiagResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiDiagResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiDiagResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiDiagResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiDiagResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiDiagResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiDiagResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiDiagResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiDiagResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiDiagResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statistics"] = o.Statistics
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0040OpenapiDiagResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statistics",
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

	varV0040OpenapiDiagResp := _V0040OpenapiDiagResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiDiagResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiDiagResp(varV0040OpenapiDiagResp)

	return err
}

type NullableV0040OpenapiDiagResp struct {
	value *V0040OpenapiDiagResp
	isSet bool
}

func (v NullableV0040OpenapiDiagResp) Get() *V0040OpenapiDiagResp {
	return v.value
}

func (v *NullableV0040OpenapiDiagResp) Set(val *V0040OpenapiDiagResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiDiagResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiDiagResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiDiagResp(val *V0040OpenapiDiagResp) *NullableV0040OpenapiDiagResp {
	return &NullableV0040OpenapiDiagResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiDiagResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiDiagResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
