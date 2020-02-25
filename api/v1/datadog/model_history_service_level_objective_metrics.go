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

// HistoryServiceLevelObjectiveMetrics A `metric` based SLO history response.
type HistoryServiceLevelObjectiveMetrics struct {
	Denominator HistoryServiceLevelObjectiveMetricsSeries `json:"denominator"`
	// The aggregated query interval for the series data. It's implicit based on the query time window.
	Interval int64 `json:"interval"`
	// Optional message if there are specific query issues/warnings.
	Message   *string                                   `json:"message,omitempty"`
	Numerator HistoryServiceLevelObjectiveMetricsSeries `json:"numerator"`
	// The combined numerator && denominator query CSV.
	Query string `json:"query"`
	// The series result type. This mimics `batch_query` response type
	ResType string `json:"res_type"`
	// The series response version type. This mimics `batch_query` response type
	RespVersion int64 `json:"resp_version"`
	// The query timestamps in epoch milliseconds
	Times []float64 `json:"times"`
}

// NewHistoryServiceLevelObjectiveMetrics instantiates a new HistoryServiceLevelObjectiveMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryServiceLevelObjectiveMetrics(denominator HistoryServiceLevelObjectiveMetricsSeries, interval int64, numerator HistoryServiceLevelObjectiveMetricsSeries, query string, resType string, respVersion int64, times []float64) *HistoryServiceLevelObjectiveMetrics {
	this := HistoryServiceLevelObjectiveMetrics{}
	this.Denominator = denominator
	this.Interval = interval
	this.Numerator = numerator
	this.Query = query
	this.ResType = resType
	this.RespVersion = respVersion
	this.Times = times
	return &this
}

// NewHistoryServiceLevelObjectiveMetricsWithDefaults instantiates a new HistoryServiceLevelObjectiveMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryServiceLevelObjectiveMetricsWithDefaults() *HistoryServiceLevelObjectiveMetrics {
	this := HistoryServiceLevelObjectiveMetrics{}
	return &this
}

// GetDenominator returns the Denominator field value
func (o *HistoryServiceLevelObjectiveMetrics) GetDenominator() HistoryServiceLevelObjectiveMetricsSeries {
	if o == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret
	}

	return o.Denominator
}

// SetDenominator sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetDenominator(v HistoryServiceLevelObjectiveMetricsSeries) {
	o.Denominator = v
}

// GetInterval returns the Interval field value
func (o *HistoryServiceLevelObjectiveMetrics) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Interval
}

// SetInterval sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetInterval(v int64) {
	o.Interval = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HistoryServiceLevelObjectiveMetrics) SetMessage(v string) {
	o.Message = &v
}

// GetNumerator returns the Numerator field value
func (o *HistoryServiceLevelObjectiveMetrics) GetNumerator() HistoryServiceLevelObjectiveMetricsSeries {
	if o == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret
	}

	return o.Numerator
}

// SetNumerator sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetNumerator(v HistoryServiceLevelObjectiveMetricsSeries) {
	o.Numerator = v
}

// GetQuery returns the Query field value
func (o *HistoryServiceLevelObjectiveMetrics) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// SetQuery sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetQuery(v string) {
	o.Query = v
}

// GetResType returns the ResType field value
func (o *HistoryServiceLevelObjectiveMetrics) GetResType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResType
}

// SetResType sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetResType(v string) {
	o.ResType = v
}

// GetRespVersion returns the RespVersion field value
func (o *HistoryServiceLevelObjectiveMetrics) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// SetRespVersion sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetRespVersion(v int64) {
	o.RespVersion = v
}

// GetTimes returns the Times field value
func (o *HistoryServiceLevelObjectiveMetrics) GetTimes() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Times
}

// SetTimes sets field value
func (o *HistoryServiceLevelObjectiveMetrics) SetTimes(v []float64) {
	o.Times = v
}

func (o HistoryServiceLevelObjectiveMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["denominator"] = o.Denominator
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["numerator"] = o.Numerator
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["res_type"] = o.ResType
	}
	if true {
		toSerialize["resp_version"] = o.RespVersion
	}
	if true {
		toSerialize["times"] = o.Times
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryServiceLevelObjectiveMetrics struct {
	value *HistoryServiceLevelObjectiveMetrics
	isSet bool
}

func (v NullableHistoryServiceLevelObjectiveMetrics) Get() *HistoryServiceLevelObjectiveMetrics {
	return v.value
}

func (v NullableHistoryServiceLevelObjectiveMetrics) Set(val *HistoryServiceLevelObjectiveMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryServiceLevelObjectiveMetrics) IsSet() bool {
	return v.isSet
}

func (v NullableHistoryServiceLevelObjectiveMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryServiceLevelObjectiveMetrics(val *HistoryServiceLevelObjectiveMetrics) *NullableHistoryServiceLevelObjectiveMetrics {
	return &NullableHistoryServiceLevelObjectiveMetrics{value: val, isSet: true}
}

func (v NullableHistoryServiceLevelObjectiveMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryServiceLevelObjectiveMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
