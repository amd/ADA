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

// checks if the V0041OpenapiJobSubmitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobSubmitResponse{}

// V0041OpenapiJobSubmitResponse struct for V0041OpenapiJobSubmitResponse
type V0041OpenapiJobSubmitResponse struct {
	Result *V0041JobSubmitResponseMsg `json:"result,omitempty"`
	// submited JobId
	JobId *int32 `json:"job_id,omitempty"`
	// submited StepID
	StepId *string `json:"step_id,omitempty"`
	// job submision user message
	JobSubmitUserMsg *string                     `json:"job_submit_user_msg,omitempty"`
	Meta             *V0041OpenapiMeta           `json:"meta,omitempty"`
	Errors           []V0041OpenapiErrorsInner   `json:"errors,omitempty"`
	Warnings         []V0041OpenapiWarningsInner `json:"warnings,omitempty"`
}

// NewV0041OpenapiJobSubmitResponse instantiates a new V0041OpenapiJobSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobSubmitResponse() *V0041OpenapiJobSubmitResponse {
	this := V0041OpenapiJobSubmitResponse{}
	return &this
}

// NewV0041OpenapiJobSubmitResponseWithDefaults instantiates a new V0041OpenapiJobSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobSubmitResponseWithDefaults() *V0041OpenapiJobSubmitResponse {
	this := V0041OpenapiJobSubmitResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetResult() V0041JobSubmitResponseMsg {
	if o == nil || IsNil(o.Result) {
		var ret V0041JobSubmitResponseMsg
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetResultOk() (*V0041JobSubmitResponseMsg, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given V0041JobSubmitResponseMsg and assigns it to the Result field.
func (o *V0041OpenapiJobSubmitResponse) SetResult(v V0041JobSubmitResponseMsg) {
	o.Result = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0041OpenapiJobSubmitResponse) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0041OpenapiJobSubmitResponse) SetStepId(v string) {
	o.StepId = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *V0041OpenapiJobSubmitResponse) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetMeta() V0041OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetMetaOk() (*V0041OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiMeta and assigns it to the Meta field.
func (o *V0041OpenapiJobSubmitResponse) SetMeta(v V0041OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetErrors() []V0041OpenapiErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetErrorsOk() ([]V0041OpenapiErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiJobSubmitResponse) SetErrors(v []V0041OpenapiErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiJobSubmitResponse) GetWarnings() []V0041OpenapiWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobSubmitResponse) GetWarningsOk() ([]V0041OpenapiWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiJobSubmitResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiJobSubmitResponse) SetWarnings(v []V0041OpenapiWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiJobSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
	}
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

type NullableV0041OpenapiJobSubmitResponse struct {
	value *V0041OpenapiJobSubmitResponse
	isSet bool
}

func (v NullableV0041OpenapiJobSubmitResponse) Get() *V0041OpenapiJobSubmitResponse {
	return v.value
}

func (v *NullableV0041OpenapiJobSubmitResponse) Set(val *V0041OpenapiJobSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobSubmitResponse(val *V0041OpenapiJobSubmitResponse) *NullableV0041OpenapiJobSubmitResponse {
	return &NullableV0041OpenapiJobSubmitResponse{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
