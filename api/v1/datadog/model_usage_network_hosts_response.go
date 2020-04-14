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

// UsageNetworkHostsResponse Response containting the number of active NPM hosts for each hour for a given organization.
type UsageNetworkHostsResponse struct {
	// Get hourly usage for NPM hosts.
	Usage *[]UsageNetworkHostsHour `json:"usage,omitempty"`
}

// NewUsageNetworkHostsResponse instantiates a new UsageNetworkHostsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageNetworkHostsResponse() *UsageNetworkHostsResponse {
	this := UsageNetworkHostsResponse{}
	return &this
}

// NewUsageNetworkHostsResponseWithDefaults instantiates a new UsageNetworkHostsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageNetworkHostsResponseWithDefaults() *UsageNetworkHostsResponse {
	this := UsageNetworkHostsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageNetworkHostsResponse) GetUsage() []UsageNetworkHostsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageNetworkHostsHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsResponse) GetUsageOk() (*[]UsageNetworkHostsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageNetworkHostsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageNetworkHostsHour and assigns it to the Usage field.
func (o *UsageNetworkHostsResponse) SetUsage(v []UsageNetworkHostsHour) {
	o.Usage = &v
}

func (o UsageNetworkHostsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageNetworkHostsResponse struct {
	value *UsageNetworkHostsResponse
	isSet bool
}

func (v NullableUsageNetworkHostsResponse) Get() *UsageNetworkHostsResponse {
	return v.value
}

func (v *NullableUsageNetworkHostsResponse) Set(val *UsageNetworkHostsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageNetworkHostsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageNetworkHostsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageNetworkHostsResponse(val *UsageNetworkHostsResponse) *NullableUsageNetworkHostsResponse {
	return &NullableUsageNetworkHostsResponse{value: val, isSet: true}
}

func (v NullableUsageNetworkHostsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageNetworkHostsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}