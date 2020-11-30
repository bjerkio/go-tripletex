// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrencyExchangeRate currency exchange rate
//
// swagger:model CurrencyExchangeRate
type CurrencyExchangeRate struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// date
	Date string `json:"date,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// rate
	// Minimum: 0
	Rate *float64 `json:"rate,omitempty"`

	// Source of exchange rates, i.e Norges Bank
	// Enum: [NORGES_BANK HALLONEN]
	Source string `json:"source,omitempty"`

	// source currency
	// Required: true
	SourceCurrency *Currency `json:"sourceCurrency"`

	// target currency
	// Required: true
	TargetCurrency *Currency `json:"targetCurrency"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this currency exchange rate
func (m *CurrencyExchangeRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrencyExchangeRate) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CurrencyExchangeRate) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := validate.Minimum("rate", "body", float64(*m.Rate), 0, false); err != nil {
		return err
	}

	return nil
}

var currencyExchangeRateTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NORGES_BANK","HALLONEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyExchangeRateTypeSourcePropEnum = append(currencyExchangeRateTypeSourcePropEnum, v)
	}
}

const (

	// CurrencyExchangeRateSourceNORGESBANK captures enum value "NORGES_BANK"
	CurrencyExchangeRateSourceNORGESBANK string = "NORGES_BANK"

	// CurrencyExchangeRateSourceHALLONEN captures enum value "HALLONEN"
	CurrencyExchangeRateSourceHALLONEN string = "HALLONEN"
)

// prop value enum
func (m *CurrencyExchangeRate) validateSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, currencyExchangeRateTypeSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CurrencyExchangeRate) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateSourceCurrency(formats strfmt.Registry) error {

	if err := validate.Required("sourceCurrency", "body", m.SourceCurrency); err != nil {
		return err
	}

	if m.SourceCurrency != nil {
		if err := m.SourceCurrency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceCurrency")
			}
			return err
		}
	}

	return nil
}

func (m *CurrencyExchangeRate) validateTargetCurrency(formats strfmt.Registry) error {

	if err := validate.Required("targetCurrency", "body", m.TargetCurrency); err != nil {
		return err
	}

	if m.TargetCurrency != nil {
		if err := m.TargetCurrency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetCurrency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyExchangeRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyExchangeRate) UnmarshalBinary(b []byte) error {
	var res CurrencyExchangeRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
