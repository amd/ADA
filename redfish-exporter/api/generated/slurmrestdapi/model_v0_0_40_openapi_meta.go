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

// checks if the V0040OpenapiMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiMeta{}

// V0040OpenapiMeta struct for V0040OpenapiMeta
type V0040OpenapiMeta struct {
	Plugin  *V0041OpenapiMetaPlugin `json:"plugin,omitempty"`
	Client  *V0041OpenapiMetaClient `json:"client,omitempty"`
	Command []string                `json:"command,omitempty"`
	Slurm   *V0041OpenapiMetaSlurm  `json:"slurm,omitempty"`
}

// NewV0040OpenapiMeta instantiates a new V0040OpenapiMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiMeta() *V0040OpenapiMeta {
	this := V0040OpenapiMeta{}
	return &this
}

// NewV0040OpenapiMetaWithDefaults instantiates a new V0040OpenapiMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiMetaWithDefaults() *V0040OpenapiMeta {
	this := V0040OpenapiMeta{}
	return &this
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *V0040OpenapiMeta) GetPlugin() V0041OpenapiMetaPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret V0041OpenapiMetaPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMeta) GetPluginOk() (*V0041OpenapiMetaPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *V0040OpenapiMeta) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given V0041OpenapiMetaPlugin and assigns it to the Plugin field.
func (o *V0040OpenapiMeta) SetPlugin(v V0041OpenapiMetaPlugin) {
	o.Plugin = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *V0040OpenapiMeta) GetClient() V0041OpenapiMetaClient {
	if o == nil || IsNil(o.Client) {
		var ret V0041OpenapiMetaClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMeta) GetClientOk() (*V0041OpenapiMetaClient, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *V0040OpenapiMeta) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given V0041OpenapiMetaClient and assigns it to the Client field.
func (o *V0040OpenapiMeta) SetClient(v V0041OpenapiMetaClient) {
	o.Client = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V0040OpenapiMeta) GetCommand() []string {
	if o == nil || IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMeta) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V0040OpenapiMeta) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *V0040OpenapiMeta) SetCommand(v []string) {
	o.Command = v
}

// GetSlurm returns the Slurm field value if set, zero value otherwise.
func (o *V0040OpenapiMeta) GetSlurm() V0041OpenapiMetaSlurm {
	if o == nil || IsNil(o.Slurm) {
		var ret V0041OpenapiMetaSlurm
		return ret
	}
	return *o.Slurm
}

// GetSlurmOk returns a tuple with the Slurm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMeta) GetSlurmOk() (*V0041OpenapiMetaSlurm, bool) {
	if o == nil || IsNil(o.Slurm) {
		return nil, false
	}
	return o.Slurm, true
}

// HasSlurm returns a boolean if a field has been set.
func (o *V0040OpenapiMeta) HasSlurm() bool {
	if o != nil && !IsNil(o.Slurm) {
		return true
	}

	return false
}

// SetSlurm gets a reference to the given V0041OpenapiMetaSlurm and assigns it to the Slurm field.
func (o *V0040OpenapiMeta) SetSlurm(v V0041OpenapiMetaSlurm) {
	o.Slurm = &v
}

func (o V0040OpenapiMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Slurm) {
		toSerialize["slurm"] = o.Slurm
	}
	return toSerialize, nil
}

type NullableV0040OpenapiMeta struct {
	value *V0040OpenapiMeta
	isSet bool
}

func (v NullableV0040OpenapiMeta) Get() *V0040OpenapiMeta {
	return v.value
}

func (v *NullableV0040OpenapiMeta) Set(val *V0040OpenapiMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiMeta(val *V0040OpenapiMeta) *NullableV0040OpenapiMeta {
	return &NullableV0040OpenapiMeta{value: val, isSet: true}
}

func (v NullableV0040OpenapiMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
