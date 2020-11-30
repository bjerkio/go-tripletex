// Code generated by go-swagger; DO NOT EDIT.

package product_price

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
)

// NewProductProductPriceSearchParams creates a new ProductProductPriceSearchParams object
// with the default values initialized.
func NewProductProductPriceSearchParams() *ProductProductPriceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductProductPriceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewProductProductPriceSearchParamsWithTimeout creates a new ProductProductPriceSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductProductPriceSearchParamsWithTimeout(timeout time.Duration) *ProductProductPriceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductProductPriceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewProductProductPriceSearchParamsWithContext creates a new ProductProductPriceSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductProductPriceSearchParamsWithContext(ctx context.Context) *ProductProductPriceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductProductPriceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewProductProductPriceSearchParamsWithHTTPClient creates a new ProductProductPriceSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductProductPriceSearchParamsWithHTTPClient(client *http.Client) *ProductProductPriceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductProductPriceSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*ProductProductPriceSearchParams contains all the parameters to send to the API endpoint
for the product product price search operation typically these are written to a http.Request
*/
type ProductProductPriceSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*FromDate
	  From and including

	*/
	FromDate *string
	/*ProductID
	  Equals

	*/
	ProductID int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*ToDate
	  To and excluding

	*/
	ToDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product product price search params
func (o *ProductProductPriceSearchParams) WithTimeout(timeout time.Duration) *ProductProductPriceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product product price search params
func (o *ProductProductPriceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product product price search params
func (o *ProductProductPriceSearchParams) WithContext(ctx context.Context) *ProductProductPriceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product product price search params
func (o *ProductProductPriceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product product price search params
func (o *ProductProductPriceSearchParams) WithHTTPClient(client *http.Client) *ProductProductPriceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product product price search params
func (o *ProductProductPriceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the product product price search params
func (o *ProductProductPriceSearchParams) WithCount(count *int64) *ProductProductPriceSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the product product price search params
func (o *ProductProductPriceSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the product product price search params
func (o *ProductProductPriceSearchParams) WithFields(fields *string) *ProductProductPriceSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the product product price search params
func (o *ProductProductPriceSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the product product price search params
func (o *ProductProductPriceSearchParams) WithFrom(from *int64) *ProductProductPriceSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the product product price search params
func (o *ProductProductPriceSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithFromDate adds the fromDate to the product product price search params
func (o *ProductProductPriceSearchParams) WithFromDate(fromDate *string) *ProductProductPriceSearchParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the product product price search params
func (o *ProductProductPriceSearchParams) SetFromDate(fromDate *string) {
	o.FromDate = fromDate
}

// WithProductID adds the productID to the product product price search params
func (o *ProductProductPriceSearchParams) WithProductID(productID int32) *ProductProductPriceSearchParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the product product price search params
func (o *ProductProductPriceSearchParams) SetProductID(productID int32) {
	o.ProductID = productID
}

// WithSorting adds the sorting to the product product price search params
func (o *ProductProductPriceSearchParams) WithSorting(sorting *string) *ProductProductPriceSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the product product price search params
func (o *ProductProductPriceSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithToDate adds the toDate to the product product price search params
func (o *ProductProductPriceSearchParams) WithToDate(toDate *string) *ProductProductPriceSearchParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the product product price search params
func (o *ProductProductPriceSearchParams) SetToDate(toDate *string) {
	o.ToDate = toDate
}

// WriteToRequest writes these params to a swagger request
func (o *ProductProductPriceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.FromDate != nil {

		// query param fromDate
		var qrFromDate string
		if o.FromDate != nil {
			qrFromDate = *o.FromDate
		}
		qFromDate := qrFromDate
		if qFromDate != "" {
			if err := r.SetQueryParam("fromDate", qFromDate); err != nil {
				return err
			}
		}

	}

	// query param productId
	qrProductID := o.ProductID
	qProductID := swag.FormatInt32(qrProductID)
	if qProductID != "" {
		if err := r.SetQueryParam("productId", qProductID); err != nil {
			return err
		}
	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if o.ToDate != nil {

		// query param toDate
		var qrToDate string
		if o.ToDate != nil {
			qrToDate = *o.ToDate
		}
		qToDate := qrToDate
		if qToDate != "" {
			if err := r.SetQueryParam("toDate", qToDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
