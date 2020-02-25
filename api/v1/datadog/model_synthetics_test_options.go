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

// SyntheticsTestOptions struct for SyntheticsTestOptions
type SyntheticsTestOptions struct {
	AcceptSelfSigned  *bool                       `json:"accept_self_signed,omitempty"`
	DeviceIds         *[]SyntheticsDeviceID       `json:"device_ids,omitempty"`
	FollowRedirects   *bool                       `json:"follow_redirects,omitempty"`
	MinLocationFailed *int64                      `json:"min_location_failed,omitempty"`
	Retry             *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	TickEvery         *SyntheticsTickInterval     `json:"tick_every,omitempty"`
}

// NewSyntheticsTestOptions instantiates a new SyntheticsTestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestOptions() *SyntheticsTestOptions {
	this := SyntheticsTestOptions{}
	return &this
}

// NewSyntheticsTestOptionsWithDefaults instantiates a new SyntheticsTestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestOptionsWithDefaults() *SyntheticsTestOptions {
	this := SyntheticsTestOptions{}
	return &this
}

// GetAcceptSelfSigned returns the AcceptSelfSigned field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetAcceptSelfSigned() bool {
	if o == nil || o.AcceptSelfSigned == nil {
		var ret bool
		return ret
	}
	return *o.AcceptSelfSigned
}

// GetAcceptSelfSignedOk returns a tuple with the AcceptSelfSigned field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetAcceptSelfSignedOk() (bool, bool) {
	if o == nil || o.AcceptSelfSigned == nil {
		var ret bool
		return ret, false
	}
	return *o.AcceptSelfSigned, true
}

// HasAcceptSelfSigned returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasAcceptSelfSigned() bool {
	if o != nil && o.AcceptSelfSigned != nil {
		return true
	}

	return false
}

// SetAcceptSelfSigned gets a reference to the given bool and assigns it to the AcceptSelfSigned field.
func (o *SyntheticsTestOptions) SetAcceptSelfSigned(v bool) {
	o.AcceptSelfSigned = &v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetDeviceIds() []SyntheticsDeviceID {
	if o == nil || o.DeviceIds == nil {
		var ret []SyntheticsDeviceID
		return ret
	}
	return *o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetDeviceIdsOk() ([]SyntheticsDeviceID, bool) {
	if o == nil || o.DeviceIds == nil {
		var ret []SyntheticsDeviceID
		return ret, false
	}
	return *o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasDeviceIds() bool {
	if o != nil && o.DeviceIds != nil {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []SyntheticsDeviceID and assigns it to the DeviceIds field.
func (o *SyntheticsTestOptions) SetDeviceIds(v []SyntheticsDeviceID) {
	o.DeviceIds = &v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetFollowRedirects() bool {
	if o == nil || o.FollowRedirects == nil {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetFollowRedirectsOk() (bool, bool) {
	if o == nil || o.FollowRedirects == nil {
		var ret bool
		return ret, false
	}
	return *o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasFollowRedirects() bool {
	if o != nil && o.FollowRedirects != nil {
		return true
	}

	return false
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *SyntheticsTestOptions) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetMinLocationFailed returns the MinLocationFailed field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetMinLocationFailed() int64 {
	if o == nil || o.MinLocationFailed == nil {
		var ret int64
		return ret
	}
	return *o.MinLocationFailed
}

// GetMinLocationFailedOk returns a tuple with the MinLocationFailed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetMinLocationFailedOk() (int64, bool) {
	if o == nil || o.MinLocationFailed == nil {
		var ret int64
		return ret, false
	}
	return *o.MinLocationFailed, true
}

// HasMinLocationFailed returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasMinLocationFailed() bool {
	if o != nil && o.MinLocationFailed != nil {
		return true
	}

	return false
}

// SetMinLocationFailed gets a reference to the given int64 and assigns it to the MinLocationFailed field.
func (o *SyntheticsTestOptions) SetMinLocationFailed(v int64) {
	o.MinLocationFailed = &v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetRetryOk() (SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret, false
	}
	return *o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasRetry() bool {
	if o != nil && o.Retry != nil {
		return true
	}

	return false
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsTestOptions) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetTickEvery returns the TickEvery field value if set, zero value otherwise.
func (o *SyntheticsTestOptions) GetTickEvery() SyntheticsTickInterval {
	if o == nil || o.TickEvery == nil {
		var ret SyntheticsTickInterval
		return ret
	}
	return *o.TickEvery
}

// GetTickEveryOk returns a tuple with the TickEvery field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptions) GetTickEveryOk() (SyntheticsTickInterval, bool) {
	if o == nil || o.TickEvery == nil {
		var ret SyntheticsTickInterval
		return ret, false
	}
	return *o.TickEvery, true
}

// HasTickEvery returns a boolean if a field has been set.
func (o *SyntheticsTestOptions) HasTickEvery() bool {
	if o != nil && o.TickEvery != nil {
		return true
	}

	return false
}

// SetTickEvery gets a reference to the given SyntheticsTickInterval and assigns it to the TickEvery field.
func (o *SyntheticsTestOptions) SetTickEvery(v SyntheticsTickInterval) {
	o.TickEvery = &v
}

func (o SyntheticsTestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptSelfSigned != nil {
		toSerialize["accept_self_signed"] = o.AcceptSelfSigned
	}
	if o.DeviceIds != nil {
		toSerialize["device_ids"] = o.DeviceIds
	}
	if o.FollowRedirects != nil {
		toSerialize["follow_redirects"] = o.FollowRedirects
	}
	if o.MinLocationFailed != nil {
		toSerialize["min_location_failed"] = o.MinLocationFailed
	}
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	if o.TickEvery != nil {
		toSerialize["tick_every"] = o.TickEvery
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTestOptions struct {
	value *SyntheticsTestOptions
	isSet bool
}

func (v NullableSyntheticsTestOptions) Get() *SyntheticsTestOptions {
	return v.value
}

func (v NullableSyntheticsTestOptions) Set(val *SyntheticsTestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestOptions) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsTestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestOptions(val *SyntheticsTestOptions) *NullableSyntheticsTestOptions {
	return &NullableSyntheticsTestOptions{value: val, isSet: true}
}

func (v NullableSyntheticsTestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
