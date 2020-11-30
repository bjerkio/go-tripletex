// Code generated by go-swagger; DO NOT EDIT.

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewLedgerVoucherPutParams creates a new LedgerVoucherPutParams object
// with the default values initialized.
func NewLedgerVoucherPutParams() *LedgerVoucherPutParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &LedgerVoucherPutParams{
		SendToLedger: &sendToLedgerDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerVoucherPutParamsWithTimeout creates a new LedgerVoucherPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerVoucherPutParamsWithTimeout(timeout time.Duration) *LedgerVoucherPutParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &LedgerVoucherPutParams{
		SendToLedger: &sendToLedgerDefault,

		timeout: timeout,
	}
}

// NewLedgerVoucherPutParamsWithContext creates a new LedgerVoucherPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerVoucherPutParamsWithContext(ctx context.Context) *LedgerVoucherPutParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &LedgerVoucherPutParams{
		SendToLedger: &sendToLedgerDefault,

		Context: ctx,
	}
}

// NewLedgerVoucherPutParamsWithHTTPClient creates a new LedgerVoucherPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerVoucherPutParamsWithHTTPClient(client *http.Client) *LedgerVoucherPutParams {
	var (
		sendToLedgerDefault = bool(true)
	)
	return &LedgerVoucherPutParams{
		SendToLedger: &sendToLedgerDefault,
		HTTPClient:   client,
	}
}

/*LedgerVoucherPutParams contains all the parameters to send to the API endpoint
for the ledger voucher put operation typically these are written to a http.Request
*/
type LedgerVoucherPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Voucher
	/*ID
	  Element ID

	*/
	ID int32
	/*SendToLedger
	  Should the voucher be sent to ledger? Requires the "Advanced Voucher" permission.

	*/
	SendToLedger *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithTimeout(timeout time.Duration) *LedgerVoucherPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithContext(ctx context.Context) *LedgerVoucherPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithHTTPClient(client *http.Client) *LedgerVoucherPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithBody(body *models.Voucher) *LedgerVoucherPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetBody(body *models.Voucher) {
	o.Body = body
}

// WithID adds the id to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithID(id int32) *LedgerVoucherPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetID(id int32) {
	o.ID = id
}

// WithSendToLedger adds the sendToLedger to the ledger voucher put params
func (o *LedgerVoucherPutParams) WithSendToLedger(sendToLedger *bool) *LedgerVoucherPutParams {
	o.SetSendToLedger(sendToLedger)
	return o
}

// SetSendToLedger adds the sendToLedger to the ledger voucher put params
func (o *LedgerVoucherPutParams) SetSendToLedger(sendToLedger *bool) {
	o.SendToLedger = sendToLedger
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerVoucherPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.SendToLedger != nil {

		// query param sendToLedger
		var qrSendToLedger bool
		if o.SendToLedger != nil {
			qrSendToLedger = *o.SendToLedger
		}
		qSendToLedger := swag.FormatBool(qrSendToLedger)
		if qSendToLedger != "" {
			if err := r.SetQueryParam("sendToLedger", qSendToLedger); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
