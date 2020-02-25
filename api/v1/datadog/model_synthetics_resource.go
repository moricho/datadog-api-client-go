/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsResource struct for SyntheticsResource
type SyntheticsResource struct {
	Duration  *float64                `json:"duration,omitempty"`
	Method    *string                 `json:"method,omitempty"`
	Size      *int64                  `json:"size,omitempty"`
	Status    *int64                  `json:"status,omitempty"`
	Timestamp *float64                `json:"timestamp,omitempty"`
	TraceId   *string                 `json:"traceId,omitempty"`
	Type      *SyntheticsResourceType `json:"type,omitempty"`
	Url       *string                 `json:"url,omitempty"`
}

// NewSyntheticsResource instantiates a new SyntheticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsResource() *SyntheticsResource {
	this := SyntheticsResource{}
	return &this
}

// NewSyntheticsResourceWithDefaults instantiates a new SyntheticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsResourceWithDefaults() *SyntheticsResource {
	this := SyntheticsResource{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsResource) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetDurationOk() (float64, bool) {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret, false
	}
	return *o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsResource) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsResource) SetDuration(v float64) {
	o.Duration = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SyntheticsResource) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetMethodOk() (string, bool) {
	if o == nil || o.Method == nil {
		var ret string
		return ret, false
	}
	return *o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SyntheticsResource) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SyntheticsResource) SetMethod(v string) {
	o.Method = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SyntheticsResource) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetSizeOk() (int64, bool) {
	if o == nil || o.Size == nil {
		var ret int64
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SyntheticsResource) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SyntheticsResource) SetSize(v int64) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsResource) GetStatus() int64 {
	if o == nil || o.Status == nil {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetStatusOk() (int64, bool) {
	if o == nil || o.Status == nil {
		var ret int64
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsResource) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SyntheticsResource) SetStatus(v int64) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SyntheticsResource) GetTimestamp() float64 {
	if o == nil || o.Timestamp == nil {
		var ret float64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetTimestampOk() (float64, bool) {
	if o == nil || o.Timestamp == nil {
		var ret float64
		return ret, false
	}
	return *o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SyntheticsResource) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float64 and assigns it to the Timestamp field.
func (o *SyntheticsResource) SetTimestamp(v float64) {
	o.Timestamp = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SyntheticsResource) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetTraceIdOk() (string, bool) {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret, false
	}
	return *o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SyntheticsResource) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SyntheticsResource) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsResource) GetType() SyntheticsResourceType {
	if o == nil || o.Type == nil {
		var ret SyntheticsResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetTypeOk() (SyntheticsResourceType, bool) {
	if o == nil || o.Type == nil {
		var ret SyntheticsResourceType
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SyntheticsResourceType and assigns it to the Type field.
func (o *SyntheticsResource) SetType(v SyntheticsResourceType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsResource) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsResource) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsResource) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsResource) SetUrl(v string) {
	o.Url = &v
}

func (o SyntheticsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsResource struct {
	value *SyntheticsResource
	isSet bool
}

func (v NullableSyntheticsResource) Get() *SyntheticsResource {
	return v.value
}

func (v NullableSyntheticsResource) Set(val *SyntheticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsResource) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsResource(val *SyntheticsResource) *NullableSyntheticsResource {
	return &NullableSyntheticsResource{value: val, isSet: true}
}

func (v NullableSyntheticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
