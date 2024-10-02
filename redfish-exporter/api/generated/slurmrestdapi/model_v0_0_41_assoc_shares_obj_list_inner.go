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

// checks if the V0041AssocSharesObjListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041AssocSharesObjListInner{}

// V0041AssocSharesObjListInner struct for V0041AssocSharesObjListInner
type V0041AssocSharesObjListInner struct {
	// assocation id
	Id *int32 `json:"id,omitempty"`
	// cluster name
	Cluster *string `json:"cluster,omitempty"`
	// share name
	Name *string `json:"name,omitempty"`
	// parent name
	Parent *string `json:"parent,omitempty"`
	// partition name
	Partition        *string                           `json:"partition,omitempty"`
	SharesNormalized *V0041Float64NoValStruct          `json:"shares_normalized,omitempty"`
	Shares           *V0041Uint32NoValStruct           `json:"shares,omitempty"`
	Tres             *V0041AssocSharesObjListInnerTres `json:"tres,omitempty"`
	// effective, normalized usage
	EffectiveUsage  *float64                 `json:"effective_usage,omitempty"`
	UsageNormalized *V0041Float64NoValStruct `json:"usage_normalized,omitempty"`
	// measure of tresbillableunits usage
	Usage     *int64                                 `json:"usage,omitempty"`
	Fairshare *V0041AssocSharesObjListInnerFairshare `json:"fairshare,omitempty"`
	// user or account association
	Type []string `json:"type,omitempty"`
}

// NewV0041AssocSharesObjListInner instantiates a new V0041AssocSharesObjListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041AssocSharesObjListInner() *V0041AssocSharesObjListInner {
	this := V0041AssocSharesObjListInner{}
	return &this
}

// NewV0041AssocSharesObjListInnerWithDefaults instantiates a new V0041AssocSharesObjListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041AssocSharesObjListInnerWithDefaults() *V0041AssocSharesObjListInner {
	this := V0041AssocSharesObjListInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0041AssocSharesObjListInner) SetId(v int32) {
	o.Id = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *V0041AssocSharesObjListInner) SetCluster(v string) {
	o.Cluster = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0041AssocSharesObjListInner) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *V0041AssocSharesObjListInner) SetParent(v string) {
	o.Parent = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0041AssocSharesObjListInner) SetPartition(v string) {
	o.Partition = &v
}

// GetSharesNormalized returns the SharesNormalized field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetSharesNormalized() V0041Float64NoValStruct {
	if o == nil || IsNil(o.SharesNormalized) {
		var ret V0041Float64NoValStruct
		return ret
	}
	return *o.SharesNormalized
}

// GetSharesNormalizedOk returns a tuple with the SharesNormalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetSharesNormalizedOk() (*V0041Float64NoValStruct, bool) {
	if o == nil || IsNil(o.SharesNormalized) {
		return nil, false
	}
	return o.SharesNormalized, true
}

// HasSharesNormalized returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasSharesNormalized() bool {
	if o != nil && !IsNil(o.SharesNormalized) {
		return true
	}

	return false
}

// SetSharesNormalized gets a reference to the given V0041Float64NoValStruct and assigns it to the SharesNormalized field.
func (o *V0041AssocSharesObjListInner) SetSharesNormalized(v V0041Float64NoValStruct) {
	o.SharesNormalized = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetShares() V0041Uint32NoValStruct {
	if o == nil || IsNil(o.Shares) {
		var ret V0041Uint32NoValStruct
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetSharesOk() (*V0041Uint32NoValStruct, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given V0041Uint32NoValStruct and assigns it to the Shares field.
func (o *V0041AssocSharesObjListInner) SetShares(v V0041Uint32NoValStruct) {
	o.Shares = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetTres() V0041AssocSharesObjListInnerTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0041AssocSharesObjListInnerTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetTresOk() (*V0041AssocSharesObjListInnerTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0041AssocSharesObjListInnerTres and assigns it to the Tres field.
func (o *V0041AssocSharesObjListInner) SetTres(v V0041AssocSharesObjListInnerTres) {
	o.Tres = &v
}

// GetEffectiveUsage returns the EffectiveUsage field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetEffectiveUsage() float64 {
	if o == nil || IsNil(o.EffectiveUsage) {
		var ret float64
		return ret
	}
	return *o.EffectiveUsage
}

// GetEffectiveUsageOk returns a tuple with the EffectiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetEffectiveUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.EffectiveUsage) {
		return nil, false
	}
	return o.EffectiveUsage, true
}

// HasEffectiveUsage returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasEffectiveUsage() bool {
	if o != nil && !IsNil(o.EffectiveUsage) {
		return true
	}

	return false
}

// SetEffectiveUsage gets a reference to the given float64 and assigns it to the EffectiveUsage field.
func (o *V0041AssocSharesObjListInner) SetEffectiveUsage(v float64) {
	o.EffectiveUsage = &v
}

// GetUsageNormalized returns the UsageNormalized field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetUsageNormalized() V0041Float64NoValStruct {
	if o == nil || IsNil(o.UsageNormalized) {
		var ret V0041Float64NoValStruct
		return ret
	}
	return *o.UsageNormalized
}

// GetUsageNormalizedOk returns a tuple with the UsageNormalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetUsageNormalizedOk() (*V0041Float64NoValStruct, bool) {
	if o == nil || IsNil(o.UsageNormalized) {
		return nil, false
	}
	return o.UsageNormalized, true
}

// HasUsageNormalized returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasUsageNormalized() bool {
	if o != nil && !IsNil(o.UsageNormalized) {
		return true
	}

	return false
}

// SetUsageNormalized gets a reference to the given V0041Float64NoValStruct and assigns it to the UsageNormalized field.
func (o *V0041AssocSharesObjListInner) SetUsageNormalized(v V0041Float64NoValStruct) {
	o.UsageNormalized = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetUsage() int64 {
	if o == nil || IsNil(o.Usage) {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *V0041AssocSharesObjListInner) SetUsage(v int64) {
	o.Usage = &v
}

// GetFairshare returns the Fairshare field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetFairshare() V0041AssocSharesObjListInnerFairshare {
	if o == nil || IsNil(o.Fairshare) {
		var ret V0041AssocSharesObjListInnerFairshare
		return ret
	}
	return *o.Fairshare
}

// GetFairshareOk returns a tuple with the Fairshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetFairshareOk() (*V0041AssocSharesObjListInnerFairshare, bool) {
	if o == nil || IsNil(o.Fairshare) {
		return nil, false
	}
	return o.Fairshare, true
}

// HasFairshare returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasFairshare() bool {
	if o != nil && !IsNil(o.Fairshare) {
		return true
	}

	return false
}

// SetFairshare gets a reference to the given V0041AssocSharesObjListInnerFairshare and assigns it to the Fairshare field.
func (o *V0041AssocSharesObjListInner) SetFairshare(v V0041AssocSharesObjListInnerFairshare) {
	o.Fairshare = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0041AssocSharesObjListInner) GetType() []string {
	if o == nil || IsNil(o.Type) {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041AssocSharesObjListInner) GetTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0041AssocSharesObjListInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *V0041AssocSharesObjListInner) SetType(v []string) {
	o.Type = v
}

func (o V0041AssocSharesObjListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041AssocSharesObjListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.SharesNormalized) {
		toSerialize["shares_normalized"] = o.SharesNormalized
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.EffectiveUsage) {
		toSerialize["effective_usage"] = o.EffectiveUsage
	}
	if !IsNil(o.UsageNormalized) {
		toSerialize["usage_normalized"] = o.UsageNormalized
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Fairshare) {
		toSerialize["fairshare"] = o.Fairshare
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV0041AssocSharesObjListInner struct {
	value *V0041AssocSharesObjListInner
	isSet bool
}

func (v NullableV0041AssocSharesObjListInner) Get() *V0041AssocSharesObjListInner {
	return v.value
}

func (v *NullableV0041AssocSharesObjListInner) Set(val *V0041AssocSharesObjListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041AssocSharesObjListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041AssocSharesObjListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041AssocSharesObjListInner(val *V0041AssocSharesObjListInner) *NullableV0041AssocSharesObjListInner {
	return &NullableV0041AssocSharesObjListInner{value: val, isSet: true}
}

func (v NullableV0041AssocSharesObjListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041AssocSharesObjListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
