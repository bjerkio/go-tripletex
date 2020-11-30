// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperInteger response wrapper integer
//
// swagger:model ResponseWrapperInteger
type ResponseWrapperInteger struct {

	// value
	Value int32 `json:"value,omitempty"`
}

// Validate validates this response wrapper integer
func (m *ResponseWrapperInteger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperInteger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperInteger) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperInteger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
