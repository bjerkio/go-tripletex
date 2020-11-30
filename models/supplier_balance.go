// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupplierBalance supplier balance
//
// swagger:model SupplierBalance
type SupplierBalance struct {

	// balance change
	// Read Only: true
	BalanceChange float64 `json:"balanceChange,omitempty"`

	// balance in
	// Read Only: true
	BalanceIn float64 `json:"balanceIn,omitempty"`

	// Currencies that have been used prior to this period, for the given filter
	// Read Only: true
	BalanceInCurrencies []*Currency `json:"balanceInCurrencies"`

	// balance out
	// Read Only: true
	BalanceOut float64 `json:"balanceOut,omitempty"`

	// sum amount negative
	// Read Only: true
	SumAmountNegative float64 `json:"sumAmountNegative,omitempty"`

	// supplier
	Supplier *Supplier `json:"supplier,omitempty"`
}

// Validate validates this supplier balance
func (m *SupplierBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalanceInCurrencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupplierBalance) validateBalanceInCurrencies(formats strfmt.Registry) error {

	if swag.IsZero(m.BalanceInCurrencies) { // not required
		return nil
	}

	for i := 0; i < len(m.BalanceInCurrencies); i++ {
		if swag.IsZero(m.BalanceInCurrencies[i]) { // not required
			continue
		}

		if m.BalanceInCurrencies[i] != nil {
			if err := m.BalanceInCurrencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balanceInCurrencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SupplierBalance) validateSupplier(formats strfmt.Registry) error {

	if swag.IsZero(m.Supplier) { // not required
		return nil
	}

	if m.Supplier != nil {
		if err := m.Supplier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supplier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupplierBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupplierBalance) UnmarshalBinary(b []byte) error {
	var res SupplierBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
