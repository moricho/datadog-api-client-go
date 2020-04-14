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

// IncidentInitialUserDefinedFieldChoices The initial choice selections.
type IncidentInitialUserDefinedFieldChoices struct {
	// The User Defined Choice ID associated with the Field.
	ChoiceId int64 `json:"choice_id"`
	// The User Defined Field ID associated with the choice.
	FieldId int64 `json:"field_id"`
}

// NewIncidentInitialUserDefinedFieldChoices instantiates a new IncidentInitialUserDefinedFieldChoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentInitialUserDefinedFieldChoices(choiceId int64, fieldId int64) *IncidentInitialUserDefinedFieldChoices {
	this := IncidentInitialUserDefinedFieldChoices{}
	this.ChoiceId = choiceId
	this.FieldId = fieldId
	return &this
}

// NewIncidentInitialUserDefinedFieldChoicesWithDefaults instantiates a new IncidentInitialUserDefinedFieldChoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentInitialUserDefinedFieldChoicesWithDefaults() *IncidentInitialUserDefinedFieldChoices {
	this := IncidentInitialUserDefinedFieldChoices{}
	return &this
}

// GetChoiceId returns the ChoiceId field value
func (o *IncidentInitialUserDefinedFieldChoices) GetChoiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ChoiceId
}

// GetChoiceIdOk returns a tuple with the ChoiceId field value
// and a boolean to check if the value has been set.
func (o *IncidentInitialUserDefinedFieldChoices) GetChoiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoiceId, true
}

// SetChoiceId sets field value
func (o *IncidentInitialUserDefinedFieldChoices) SetChoiceId(v int64) {
	o.ChoiceId = v
}

// GetFieldId returns the FieldId field value
func (o *IncidentInitialUserDefinedFieldChoices) GetFieldId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *IncidentInitialUserDefinedFieldChoices) GetFieldIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value
func (o *IncidentInitialUserDefinedFieldChoices) SetFieldId(v int64) {
	o.FieldId = v
}

func (o IncidentInitialUserDefinedFieldChoices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["choice_id"] = o.ChoiceId
	}
	if true {
		toSerialize["field_id"] = o.FieldId
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentInitialUserDefinedFieldChoices struct {
	value *IncidentInitialUserDefinedFieldChoices
	isSet bool
}

func (v NullableIncidentInitialUserDefinedFieldChoices) Get() *IncidentInitialUserDefinedFieldChoices {
	return v.value
}

func (v *NullableIncidentInitialUserDefinedFieldChoices) Set(val *IncidentInitialUserDefinedFieldChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentInitialUserDefinedFieldChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentInitialUserDefinedFieldChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentInitialUserDefinedFieldChoices(val *IncidentInitialUserDefinedFieldChoices) *NullableIncidentInitialUserDefinedFieldChoices {
	return &NullableIncidentInitialUserDefinedFieldChoices{value: val, isSet: true}
}

func (v NullableIncidentInitialUserDefinedFieldChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentInitialUserDefinedFieldChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}