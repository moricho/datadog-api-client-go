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

// UsersResponsePayload Response containing information about multiple users.
type UsersResponsePayload struct {
	// Array of returned users.
	Data *[]UserResponse `json:"data,omitempty"`
	// Array of objects related to the users.
	Included *[]UserResponseIncludedItem `json:"included,omitempty"`
	Meta     *ResponseMetaAttributes     `json:"meta,omitempty"`
}

// NewUsersResponsePayload instantiates a new UsersResponsePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersResponsePayload() *UsersResponsePayload {
	this := UsersResponsePayload{}
	return &this
}

// NewUsersResponsePayloadWithDefaults instantiates a new UsersResponsePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersResponsePayloadWithDefaults() *UsersResponsePayload {
	this := UsersResponsePayload{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsersResponsePayload) GetData() []UserResponse {
	if o == nil || o.Data == nil {
		var ret []UserResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponsePayload) GetDataOk() (*[]UserResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsersResponsePayload) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UserResponse and assigns it to the Data field.
func (o *UsersResponsePayload) SetData(v []UserResponse) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *UsersResponsePayload) GetIncluded() []UserResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []UserResponseIncludedItem
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponsePayload) GetIncludedOk() (*[]UserResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *UsersResponsePayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserResponseIncludedItem and assigns it to the Included field.
func (o *UsersResponsePayload) SetIncluded(v []UserResponseIncludedItem) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsersResponsePayload) GetMeta() ResponseMetaAttributes {
	if o == nil || o.Meta == nil {
		var ret ResponseMetaAttributes
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponsePayload) GetMetaOk() (*ResponseMetaAttributes, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsersResponsePayload) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetaAttributes and assigns it to the Meta field.
func (o *UsersResponsePayload) SetMeta(v ResponseMetaAttributes) {
	o.Meta = &v
}

func (o UsersResponsePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
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

type NullableUsersResponsePayload struct {
	value *UsersResponsePayload
	isSet bool
}

func (v NullableUsersResponsePayload) Get() *UsersResponsePayload {
	return v.value
}

func (v *NullableUsersResponsePayload) Set(val *UsersResponsePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersResponsePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersResponsePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersResponsePayload(val *UsersResponsePayload) *NullableUsersResponsePayload {
	return &NullableUsersResponsePayload{value: val, isSet: true}
}

func (v NullableUsersResponsePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersResponsePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}