// Code generated by go-swagger; DO NOT EDIT.

package stocktaking

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

// NewInventoryStocktakingSearchParams creates a new InventoryStocktakingSearchParams object
// with the default values initialized.
func NewInventoryStocktakingSearchParams() *InventoryStocktakingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InventoryStocktakingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingSearchParamsWithTimeout creates a new InventoryStocktakingSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingSearchParamsWithTimeout(timeout time.Duration) *InventoryStocktakingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InventoryStocktakingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewInventoryStocktakingSearchParamsWithContext creates a new InventoryStocktakingSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingSearchParamsWithContext(ctx context.Context) *InventoryStocktakingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InventoryStocktakingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewInventoryStocktakingSearchParamsWithHTTPClient creates a new InventoryStocktakingSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingSearchParamsWithHTTPClient(client *http.Client) *InventoryStocktakingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InventoryStocktakingSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*InventoryStocktakingSearchParams contains all the parameters to send to the API endpoint
for the inventory stocktaking search operation typically these are written to a http.Request
*/
type InventoryStocktakingSearchParams struct {

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
	/*ID
	  List of IDs

	*/
	ID *string
	/*InventoryID
	  Equals

	*/
	InventoryID *int32
	/*IsCompleted
	  Equals

	*/
	IsCompleted *bool
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithTimeout(timeout time.Duration) *InventoryStocktakingSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithContext(ctx context.Context) *InventoryStocktakingSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithHTTPClient(client *http.Client) *InventoryStocktakingSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithCount(count *int64) *InventoryStocktakingSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithFields(fields *string) *InventoryStocktakingSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithFrom(from *int64) *InventoryStocktakingSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithID(id *string) *InventoryStocktakingSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetID(id *string) {
	o.ID = id
}

// WithInventoryID adds the inventoryID to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithInventoryID(inventoryID *int32) *InventoryStocktakingSearchParams {
	o.SetInventoryID(inventoryID)
	return o
}

// SetInventoryID adds the inventoryId to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetInventoryID(inventoryID *int32) {
	o.InventoryID = inventoryID
}

// WithIsCompleted adds the isCompleted to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithIsCompleted(isCompleted *bool) *InventoryStocktakingSearchParams {
	o.SetIsCompleted(isCompleted)
	return o
}

// SetIsCompleted adds the isCompleted to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetIsCompleted(isCompleted *bool) {
	o.IsCompleted = isCompleted
}

// WithSorting adds the sorting to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) WithSorting(sorting *string) *InventoryStocktakingSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the inventory stocktaking search params
func (o *InventoryStocktakingSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.InventoryID != nil {

		// query param inventoryId
		var qrInventoryID int32
		if o.InventoryID != nil {
			qrInventoryID = *o.InventoryID
		}
		qInventoryID := swag.FormatInt32(qrInventoryID)
		if qInventoryID != "" {
			if err := r.SetQueryParam("inventoryId", qInventoryID); err != nil {
				return err
			}
		}

	}

	if o.IsCompleted != nil {

		// query param isCompleted
		var qrIsCompleted bool
		if o.IsCompleted != nil {
			qrIsCompleted = *o.IsCompleted
		}
		qIsCompleted := swag.FormatBool(qrIsCompleted)
		if qIsCompleted != "" {
			if err := r.SetQueryParam("isCompleted", qIsCompleted); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
