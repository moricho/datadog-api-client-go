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

// Team Represents a Team Item.
type Team struct {
	Attributes *TeamAttributes `json:"attributes,omitempty"`
	// Unique identifier that represents the Incident Config Field Choice.
	Id            *string                           `json:"id,omitempty"`
	Relationships *IncidentConfigFieldRelationships `json:"relationships,omitempty"`
	// JSONAPI Model Type
	Type *string `json:"type,omitempty"`
}

// NewTeam instantiates a new Team object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeam() *Team {
	this := Team{}
	var type_ string = "teams"
	this.Type = &type_
	return &this
}

// NewTeamWithDefaults instantiates a new Team object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamWithDefaults() *Team {
	this := Team{}
	var type_ string = "teams"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Team) GetAttributes() TeamAttributes {
	if o == nil || o.Attributes == nil {
		var ret TeamAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetAttributesOk() (*TeamAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Team) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TeamAttributes and assigns it to the Attributes field.
func (o *Team) SetAttributes(v TeamAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Team) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Team) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Team) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Team) GetRelationships() IncidentConfigFieldRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentConfigFieldRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetRelationshipsOk() (*IncidentConfigFieldRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Team) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given IncidentConfigFieldRelationships and assigns it to the Relationships field.
func (o *Team) SetRelationships(v IncidentConfigFieldRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Team) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Team) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Team) SetType(v string) {
	o.Type = &v
}

func (o Team) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTeam struct {
	value *Team
	isSet bool
}

func (v NullableTeam) Get() *Team {
	return v.value
}

func (v *NullableTeam) Set(val *Team) {
	v.value = val
	v.isSet = true
}

func (v NullableTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeam(val *Team) *NullableTeam {
	return &NullableTeam{value: val, isSet: true}
}

func (v NullableTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}