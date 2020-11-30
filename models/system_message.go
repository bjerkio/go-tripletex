// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemMessage system message
//
// swagger:model SystemMessage
type SystemMessage struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this system message
func (m *SystemMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemMessage) UnmarshalBinary(b []byte) error {
	var res SystemMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
