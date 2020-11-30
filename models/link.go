// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Link link
//
// swagger:model Link
type Link struct {

	// href
	Href string `json:"href,omitempty"`

	// rel
	Rel string `json:"rel,omitempty"`

	// type
	// Enum: [POST PUT GET DELETE]
	Type string `json:"type,omitempty"`
}

// Validate validates this link
func (m *Link) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var linkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["POST","PUT","GET","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkTypeTypePropEnum = append(linkTypeTypePropEnum, v)
	}
}

const (

	// LinkTypePOST captures enum value "POST"
	LinkTypePOST string = "POST"

	// LinkTypePUT captures enum value "PUT"
	LinkTypePUT string = "PUT"

	// LinkTypeGET captures enum value "GET"
	LinkTypeGET string = "GET"

	// LinkTypeDELETE captures enum value "DELETE"
	LinkTypeDELETE string = "DELETE"
)

// prop value enum
func (m *Link) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, linkTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Link) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Link) UnmarshalBinary(b []byte) error {
	var res Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
