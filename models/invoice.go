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

// Invoice invoice
//
// swagger:model Invoice
type Invoice struct {

	// In the company’s currency, typically NOK.
	// Read Only: true
	Amount float64 `json:"amount,omitempty"`

	// In the specified currency.
	// Read Only: true
	AmountCurrency float64 `json:"amountCurrency,omitempty"`

	// Amount excluding VAT (NOK).
	// Read Only: true
	AmountExcludingVat float64 `json:"amountExcludingVat,omitempty"`

	// Amount excluding VAT in the specified currency.
	// Read Only: true
	AmountExcludingVatCurrency float64 `json:"amountExcludingVatCurrency,omitempty"`

	// The amount outstanding based on the history collection, excluding reminders and any existing remits, in the invoice currency.
	// Read Only: true
	AmountOutstanding float64 `json:"amountOutstanding,omitempty"`

	// The amount outstanding based on the history collection and including the last reminder and any existing remits. This is the total invoice balance including reminders and remittances, in the invoice currency.
	// Read Only: true
	AmountOutstandingTotal float64 `json:"amountOutstandingTotal,omitempty"`

	// Amount of round off to nearest integer.
	// Read Only: true
	AmountRoundoff float64 `json:"amountRoundoff,omitempty"`

	// Amount of round off to nearest integer in the specified currency.
	// Read Only: true
	AmountRoundoffCurrency float64 `json:"amountRoundoffCurrency,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// Comment text for the specific invoice.
	Comment string `json:"comment,omitempty"`

	// The id of the original invoice if this is a credit note.
	// Read Only: true
	CreditedInvoice int32 `json:"creditedInvoice,omitempty"`

	// currency
	// Read Only: true
	Currency *Currency `json:"currency,omitempty"`

	// The invoice customer
	// Read Only: true
	Customer *Customer `json:"customer,omitempty"`

	// The delivery date.
	// Read Only: true
	DeliveryDate string `json:"deliveryDate,omitempty"`

	// [Deprecated] EHF (Peppol) send status. This only shows status for historic EHFs.
	// Enum: [DO_NOT_SEND SEND SENT SEND_FAILURE_RECIPIENT_NOT_FOUND]
	EhfSendStatus string `json:"ehfSendStatus,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// Comment text for the invoice. This was specified on the order as invoiceComment.
	// Read Only: true
	InvoiceComment string `json:"invoiceComment,omitempty"`

	// invoice date
	// Required: true
	InvoiceDate *string `json:"invoiceDate"`

	// invoice due date
	// Required: true
	InvoiceDueDate *string `json:"invoiceDueDate"`

	// If value is set to 0, the invoice number will be generated.
	// Minimum: 0
	InvoiceNumber *int32 `json:"invoiceNumber,omitempty"`

	// Invoice remarks - automatically stops reminder/notice of debt collection if specified.
	InvoiceRemarks string `json:"invoiceRemarks,omitempty"`

	// is approved
	// Read Only: true
	IsApproved *bool `json:"isApproved,omitempty"`

	// is charged
	// Read Only: true
	IsCharged *bool `json:"isCharged,omitempty"`

	// is credit note
	// Read Only: true
	IsCreditNote *bool `json:"isCreditNote,omitempty"`

	// is credited
	// Read Only: true
	IsCredited *bool `json:"isCredited,omitempty"`

	// KID - Kundeidentifikasjonsnummer.
	// Max Length: 25
	Kid string `json:"kid,omitempty"`

	// Related orders. Only one order per invoice is supported at the moment.
	// Required: true
	Orders []*Order `json:"orders"`

	// [BETA] Optional. Used to specify the prepaid amount of the invoice. The paid amount can be specified here, or as a parameter to the /invoice API endpoint.
	PaidAmount float64 `json:"paidAmount,omitempty"`

	// [BETA] Optional. Used to specify payment type for prepaid invoices. Payment type can be specified here, or as a parameter to the /invoice API endpoint.
	// Minimum: 0
	PaymentTypeID *int32 `json:"paymentTypeId,omitempty"`

	// The invoice postings, which includes a posting for the invoice with a positive amount, and one or more posting for the payments with negative amounts.
	// Read Only: true
	Postings []*Posting `json:"postings"`

	// ProjectInvoiceDetails contains additional information about the invoice, in particular invoices for projects. It contains information about the charged project, the fee amount, extra percent and amount, extra costs, travel expenses, invoice and project comments, akonto amount and values determining if extra costs, akonto and hours should be included. ProjectInvoiceDetails is an object which represents the relation between an invoice and a Project, Orderline and OrderOut object.
	// Read Only: true
	ProjectInvoiceDetails []*ProjectInvoiceDetails `json:"projectInvoiceDetails"`

	// Invoice debt collection and reminders.
	// Read Only: true
	Reminders []*Reminder `json:"reminders"`

	// The sum of all open remittances of the invoice. Remittances are reimbursement payments back to the customer and are therefore relevant to the bookkeeping of the invoice in the accounts.
	// Read Only: true
	SumRemits float64 `json:"sumRemits,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// The invoice voucher.
	// Read Only: true
	Voucher *Voucher `json:"voucher,omitempty"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEhfSendStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectInvoiceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReminders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoucher(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) validateChanges(formats strfmt.Registry) error {

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

func (m *Invoice) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

var invoiceTypeEhfSendStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_SEND","SEND","SENT","SEND_FAILURE_RECIPIENT_NOT_FOUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypeEhfSendStatusPropEnum = append(invoiceTypeEhfSendStatusPropEnum, v)
	}
}

const (

	// InvoiceEhfSendStatusDONOTSEND captures enum value "DO_NOT_SEND"
	InvoiceEhfSendStatusDONOTSEND string = "DO_NOT_SEND"

	// InvoiceEhfSendStatusSEND captures enum value "SEND"
	InvoiceEhfSendStatusSEND string = "SEND"

	// InvoiceEhfSendStatusSENT captures enum value "SENT"
	InvoiceEhfSendStatusSENT string = "SENT"

	// InvoiceEhfSendStatusSENDFAILURERECIPIENTNOTFOUND captures enum value "SEND_FAILURE_RECIPIENT_NOT_FOUND"
	InvoiceEhfSendStatusSENDFAILURERECIPIENTNOTFOUND string = "SEND_FAILURE_RECIPIENT_NOT_FOUND"
)

// prop value enum
func (m *Invoice) validateEhfSendStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypeEhfSendStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validateEhfSendStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EhfSendStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEhfSendStatusEnum("ehfSendStatus", "body", m.EhfSendStatus); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceDate(formats strfmt.Registry) error {

	if err := validate.Required("invoiceDate", "body", m.InvoiceDate); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceDueDate(formats strfmt.Registry) error {

	if err := validate.Required("invoiceDueDate", "body", m.InvoiceDueDate); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceNumber) { // not required
		return nil
	}

	if err := validate.MinimumInt("invoiceNumber", "body", int64(*m.InvoiceNumber), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateKid(formats strfmt.Registry) error {

	if swag.IsZero(m.Kid) { // not required
		return nil
	}

	if err := validate.MaxLength("kid", "body", string(m.Kid), 25); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateOrders(formats strfmt.Registry) error {

	if err := validate.Required("orders", "body", m.Orders); err != nil {
		return err
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validatePaymentTypeID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentTypeID) { // not required
		return nil
	}

	if err := validate.MinimumInt("paymentTypeId", "body", int64(*m.PaymentTypeID), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePostings(formats strfmt.Registry) error {

	if swag.IsZero(m.Postings) { // not required
		return nil
	}

	for i := 0; i < len(m.Postings); i++ {
		if swag.IsZero(m.Postings[i]) { // not required
			continue
		}

		if m.Postings[i] != nil {
			if err := m.Postings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validateProjectInvoiceDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectInvoiceDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectInvoiceDetails); i++ {
		if swag.IsZero(m.ProjectInvoiceDetails[i]) { // not required
			continue
		}

		if m.ProjectInvoiceDetails[i] != nil {
			if err := m.ProjectInvoiceDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectInvoiceDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validateReminders(formats strfmt.Registry) error {

	if swag.IsZero(m.Reminders) { // not required
		return nil
	}

	for i := 0; i < len(m.Reminders); i++ {
		if swag.IsZero(m.Reminders[i]) { // not required
			continue
		}

		if m.Reminders[i] != nil {
			if err := m.Reminders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reminders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validateVoucher(formats strfmt.Registry) error {

	if swag.IsZero(m.Voucher) { // not required
		return nil
	}

	if m.Voucher != nil {
		if err := m.Voucher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voucher")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
