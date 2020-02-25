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

// ServiceLevelObjectivesBulkDeletedErrors struct for ServiceLevelObjectivesBulkDeletedErrors
type ServiceLevelObjectivesBulkDeletedErrors struct {
	// The ID of the service level objective object associated with this error.
	Id string `json:"id"`
	// The error message
	Message   string            `json:"message"`
	Timeframe SLOErrorTimeframe `json:"timeframe"`
}

// NewServiceLevelObjectivesBulkDeletedErrors instantiates a new ServiceLevelObjectivesBulkDeletedErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectivesBulkDeletedErrors(id string, message string, timeframe SLOErrorTimeframe) *ServiceLevelObjectivesBulkDeletedErrors {
	this := ServiceLevelObjectivesBulkDeletedErrors{}
	this.Id = id
	this.Message = message
	this.Timeframe = timeframe
	return &this
}

// NewServiceLevelObjectivesBulkDeletedErrorsWithDefaults instantiates a new ServiceLevelObjectivesBulkDeletedErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectivesBulkDeletedErrorsWithDefaults() *ServiceLevelObjectivesBulkDeletedErrors {
	this := ServiceLevelObjectivesBulkDeletedErrors{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetMessage(v string) {
	o.Message = v
}

// GetTimeframe returns the Timeframe field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetTimeframe() SLOErrorTimeframe {
	if o == nil {
		var ret SLOErrorTimeframe
		return ret
	}

	return o.Timeframe
}

// SetTimeframe sets field value
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetTimeframe(v SLOErrorTimeframe) {
	o.Timeframe = v
}

func (o ServiceLevelObjectivesBulkDeletedErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLevelObjectivesBulkDeletedErrors struct {
	value *ServiceLevelObjectivesBulkDeletedErrors
	isSet bool
}

func (v NullableServiceLevelObjectivesBulkDeletedErrors) Get() *ServiceLevelObjectivesBulkDeletedErrors {
	return v.value
}

func (v NullableServiceLevelObjectivesBulkDeletedErrors) Set(val *ServiceLevelObjectivesBulkDeletedErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectivesBulkDeletedErrors) IsSet() bool {
	return v.isSet
}

func (v NullableServiceLevelObjectivesBulkDeletedErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectivesBulkDeletedErrors(val *ServiceLevelObjectivesBulkDeletedErrors) *NullableServiceLevelObjectivesBulkDeletedErrors {
	return &NullableServiceLevelObjectivesBulkDeletedErrors{value: val, isSet: true}
}

func (v NullableServiceLevelObjectivesBulkDeletedErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectivesBulkDeletedErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
