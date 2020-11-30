// Code generated by go-swagger; DO NOT EDIT.

package cost

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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewTravelExpenseCostPostParams creates a new TravelExpenseCostPostParams object
// with the default values initialized.
func NewTravelExpenseCostPostParams() *TravelExpenseCostPostParams {
	var ()
	return &TravelExpenseCostPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseCostPostParamsWithTimeout creates a new TravelExpenseCostPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseCostPostParamsWithTimeout(timeout time.Duration) *TravelExpenseCostPostParams {
	var ()
	return &TravelExpenseCostPostParams{

		timeout: timeout,
	}
}

// NewTravelExpenseCostPostParamsWithContext creates a new TravelExpenseCostPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseCostPostParamsWithContext(ctx context.Context) *TravelExpenseCostPostParams {
	var ()
	return &TravelExpenseCostPostParams{

		Context: ctx,
	}
}

// NewTravelExpenseCostPostParamsWithHTTPClient creates a new TravelExpenseCostPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseCostPostParamsWithHTTPClient(client *http.Client) *TravelExpenseCostPostParams {
	var ()
	return &TravelExpenseCostPostParams{
		HTTPClient: client,
	}
}

/*TravelExpenseCostPostParams contains all the parameters to send to the API endpoint
for the travel expense cost post operation typically these are written to a http.Request
*/
type TravelExpenseCostPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Cost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense cost post params
func (o *TravelExpenseCostPostParams) WithTimeout(timeout time.Duration) *TravelExpenseCostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense cost post params
func (o *TravelExpenseCostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense cost post params
func (o *TravelExpenseCostPostParams) WithContext(ctx context.Context) *TravelExpenseCostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense cost post params
func (o *TravelExpenseCostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense cost post params
func (o *TravelExpenseCostPostParams) WithHTTPClient(client *http.Client) *TravelExpenseCostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense cost post params
func (o *TravelExpenseCostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the travel expense cost post params
func (o *TravelExpenseCostPostParams) WithBody(body *models.Cost) *TravelExpenseCostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the travel expense cost post params
func (o *TravelExpenseCostPostParams) SetBody(body *models.Cost) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseCostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
