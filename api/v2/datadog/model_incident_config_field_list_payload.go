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

// IncidentConfigFieldListPayload Represents the JSON API Payload of a List of Incident Config Field Items.
type IncidentConfigFieldListPayload struct {
	// The Incident config fields.
	Data []IncidentConfigField `json:"data"`
	// The User relationships.
	Included *[]UserJSONAPIResponse `json:"included,omitempty"`
	Meta     *QueryMetadata         `json:"meta,omitempty"`
}

// NewIncidentConfigFieldListPayload instantiates a new IncidentConfigFieldListPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentConfigFieldListPayload(data []IncidentConfigField) *IncidentConfigFieldListPayload {
	this := IncidentConfigFieldListPayload{}
	this.Data = data
	return &this
}

// NewIncidentConfigFieldListPayloadWithDefaults instantiates a new IncidentConfigFieldListPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentConfigFieldListPayloadWithDefaults() *IncidentConfigFieldListPayload {
	this := IncidentConfigFieldListPayload{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentConfigFieldListPayload) GetData() []IncidentConfigField {
	if o == nil {
		var ret []IncidentConfigField
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldListPayload) GetDataOk() (*[]IncidentConfigField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentConfigFieldListPayload) SetData(v []IncidentConfigField) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentConfigFieldListPayload) GetIncluded() []UserJSONAPIResponse {
	if o == nil || o.Included == nil {
		var ret []UserJSONAPIResponse
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentConfigFieldListPayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserJSONAPIResponse and assigns it to the Included field.
func (o *IncidentConfigFieldListPayload) SetIncluded(v []UserJSONAPIResponse) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentConfigFieldListPayload) GetMeta() QueryMetadata {
	if o == nil || o.Meta == nil {
		var ret QueryMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldListPayload) GetMetaOk() (*QueryMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentConfigFieldListPayload) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given QueryMetadata and assigns it to the Meta field.
func (o *IncidentConfigFieldListPayload) SetMeta(v QueryMetadata) {
	o.Meta = &v
}

func (o IncidentConfigFieldListPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentConfigFieldListPayload struct {
	value *IncidentConfigFieldListPayload
	isSet bool
}

func (v NullableIncidentConfigFieldListPayload) Get() *IncidentConfigFieldListPayload {
	return v.value
}

func (v *NullableIncidentConfigFieldListPayload) Set(val *IncidentConfigFieldListPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentConfigFieldListPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentConfigFieldListPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentConfigFieldListPayload(val *IncidentConfigFieldListPayload) *NullableIncidentConfigFieldListPayload {
	return &NullableIncidentConfigFieldListPayload{value: val, isSet: true}
}

func (v NullableIncidentConfigFieldListPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentConfigFieldListPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}