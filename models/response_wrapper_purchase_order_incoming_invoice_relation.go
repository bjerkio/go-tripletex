// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperPurchaseOrderIncomingInvoiceRelation response wrapper purchase order incoming invoice relation
//
// swagger:model ResponseWrapperPurchaseOrderIncomingInvoiceRelation
type ResponseWrapperPurchaseOrderIncomingInvoiceRelation struct {

	// value
	Value *PurchaseOrderIncomingInvoiceRelation `json:"value,omitempty"`
}

// Validate validates this response wrapper purchase order incoming invoice relation
func (m *ResponseWrapperPurchaseOrderIncomingInvoiceRelation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperPurchaseOrderIncomingInvoiceRelation) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperPurchaseOrderIncomingInvoiceRelation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperPurchaseOrderIncomingInvoiceRelation) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperPurchaseOrderIncomingInvoiceRelation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
