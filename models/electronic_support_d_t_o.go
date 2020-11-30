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

// ElectronicSupportDTO electronic support d t o
//
// swagger:model ElectronicSupportDTO
type ElectronicSupportDTO struct {

	// Bank ID
	BankID int32 `json:"bankId,omitempty"`

	// Name of the bank
	BankName string `json:"bankName,omitempty"`

	// Bank url for ordering electronic agreements for ElectronicSupportDTO of type PARTIAL.
	BankURL string `json:"bankUrl,omitempty"`

	// Type of electronic agreement creation is supported by this bank.COMPLETE: Supports creating the agreement towards AutoPay and Tripletex though the bank.PARTIAL: Supports creating the agreement towards AutoPay only.
	// Enum: [PARTIAL COMPLETE]
	Type string `json:"type,omitempty"`
}

// Validate validates this electronic support d t o
func (m *ElectronicSupportDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var electronicSupportDTOTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARTIAL","COMPLETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		electronicSupportDTOTypeTypePropEnum = append(electronicSupportDTOTypeTypePropEnum, v)
	}
}

const (

	// ElectronicSupportDTOTypePARTIAL captures enum value "PARTIAL"
	ElectronicSupportDTOTypePARTIAL string = "PARTIAL"

	// ElectronicSupportDTOTypeCOMPLETE captures enum value "COMPLETE"
	ElectronicSupportDTOTypeCOMPLETE string = "COMPLETE"
)

// prop value enum
func (m *ElectronicSupportDTO) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, electronicSupportDTOTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ElectronicSupportDTO) validateType(formats strfmt.Registry) error {

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
func (m *ElectronicSupportDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElectronicSupportDTO) UnmarshalBinary(b []byte) error {
	var res ElectronicSupportDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
