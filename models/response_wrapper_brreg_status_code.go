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

// ResponseWrapperBrregStatusCode response wrapper brreg status code
//
// swagger:model ResponseWrapperBrregStatusCode
type ResponseWrapperBrregStatusCode struct {

	// value
	// Enum: [DENIED MANUAL_CHECK ACCEPTED]
	Value string `json:"value,omitempty"`
}

// Validate validates this response wrapper brreg status code
func (m *ResponseWrapperBrregStatusCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var responseWrapperBrregStatusCodeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DENIED","MANUAL_CHECK","ACCEPTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseWrapperBrregStatusCodeTypeValuePropEnum = append(responseWrapperBrregStatusCodeTypeValuePropEnum, v)
	}
}

const (

	// ResponseWrapperBrregStatusCodeValueDENIED captures enum value "DENIED"
	ResponseWrapperBrregStatusCodeValueDENIED string = "DENIED"

	// ResponseWrapperBrregStatusCodeValueMANUALCHECK captures enum value "MANUAL_CHECK"
	ResponseWrapperBrregStatusCodeValueMANUALCHECK string = "MANUAL_CHECK"

	// ResponseWrapperBrregStatusCodeValueACCEPTED captures enum value "ACCEPTED"
	ResponseWrapperBrregStatusCodeValueACCEPTED string = "ACCEPTED"
)

// prop value enum
func (m *ResponseWrapperBrregStatusCode) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responseWrapperBrregStatusCodeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponseWrapperBrregStatusCode) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperBrregStatusCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperBrregStatusCode) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperBrregStatusCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
