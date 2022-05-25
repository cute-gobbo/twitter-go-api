/*
Tweets and Users

API Reference — Labs v2

API version: 2.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twitterapi

import (
	"encoding/json"
)

// UsageCapExceededProblem A problem that indicates that a usage cap has been exceeded.
type UsageCapExceededProblem struct {
	Type *string `json:"type,omitempty"`
	Period *string `json:"period,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Title string `json:"title"`
	Detail string `json:"detail"`
}

// NewUsageCapExceededProblem instantiates a new UsageCapExceededProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageCapExceededProblem(type_ string, title string, detail string) *UsageCapExceededProblem {
	this := UsageCapExceededProblem{}
	this.Type = &type_
	this.Title = title
	this.Detail = detail
	return &this
}

// NewUsageCapExceededProblemWithDefaults instantiates a new UsageCapExceededProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageCapExceededProblemWithDefaults() *UsageCapExceededProblem {
	this := UsageCapExceededProblem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UsageCapExceededProblem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCapExceededProblem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UsageCapExceededProblem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UsageCapExceededProblem) SetType(v string) {
	o.Type = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *UsageCapExceededProblem) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCapExceededProblem) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *UsageCapExceededProblem) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *UsageCapExceededProblem) SetPeriod(v string) {
	o.Period = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UsageCapExceededProblem) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCapExceededProblem) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UsageCapExceededProblem) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UsageCapExceededProblem) SetScope(v string) {
	o.Scope = &v
}

// GetTitle returns the Title field value
func (o *UsageCapExceededProblem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UsageCapExceededProblem) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UsageCapExceededProblem) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *UsageCapExceededProblem) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *UsageCapExceededProblem) GetDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *UsageCapExceededProblem) SetDetail(v string) {
	o.Detail = v
}

func (o UsageCapExceededProblem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableUsageCapExceededProblem struct {
	value *UsageCapExceededProblem
	isSet bool
}

func (v NullableUsageCapExceededProblem) Get() *UsageCapExceededProblem {
	return v.value
}

func (v *NullableUsageCapExceededProblem) Set(val *UsageCapExceededProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageCapExceededProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageCapExceededProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageCapExceededProblem(val *UsageCapExceededProblem) *NullableUsageCapExceededProblem {
	return &NullableUsageCapExceededProblem{value: val, isSet: true}
}

func (v NullableUsageCapExceededProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageCapExceededProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

