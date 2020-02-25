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

// AWSLogsAsyncResponse struct for AWSLogsAsyncResponse
type AWSLogsAsyncResponse struct {
	Errors *[]AWSLogsAsyncResponseErrors `json:"errors,omitempty"`
	Status *string                       `json:"status,omitempty"`
}

// NewAWSLogsAsyncResponse instantiates a new AWSLogsAsyncResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSLogsAsyncResponse() *AWSLogsAsyncResponse {
	this := AWSLogsAsyncResponse{}
	return &this
}

// NewAWSLogsAsyncResponseWithDefaults instantiates a new AWSLogsAsyncResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSLogsAsyncResponseWithDefaults() *AWSLogsAsyncResponse {
	this := AWSLogsAsyncResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AWSLogsAsyncResponse) GetErrors() []AWSLogsAsyncResponseErrors {
	if o == nil || o.Errors == nil {
		var ret []AWSLogsAsyncResponseErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsAsyncResponse) GetErrorsOk() ([]AWSLogsAsyncResponseErrors, bool) {
	if o == nil || o.Errors == nil {
		var ret []AWSLogsAsyncResponseErrors
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AWSLogsAsyncResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AWSLogsAsyncResponseErrors and assigns it to the Errors field.
func (o *AWSLogsAsyncResponse) SetErrors(v []AWSLogsAsyncResponseErrors) {
	o.Errors = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AWSLogsAsyncResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsAsyncResponse) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AWSLogsAsyncResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AWSLogsAsyncResponse) SetStatus(v string) {
	o.Status = &v
}

func (o AWSLogsAsyncResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAWSLogsAsyncResponse struct {
	value *AWSLogsAsyncResponse
	isSet bool
}

func (v NullableAWSLogsAsyncResponse) Get() *AWSLogsAsyncResponse {
	return v.value
}

func (v NullableAWSLogsAsyncResponse) Set(val *AWSLogsAsyncResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSLogsAsyncResponse) IsSet() bool {
	return v.isSet
}

func (v NullableAWSLogsAsyncResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSLogsAsyncResponse(val *AWSLogsAsyncResponse) *NullableAWSLogsAsyncResponse {
	return &NullableAWSLogsAsyncResponse{value: val, isSet: true}
}

func (v NullableAWSLogsAsyncResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSLogsAsyncResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
