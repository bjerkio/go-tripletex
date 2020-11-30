// Code generated by go-swagger; DO NOT EDIT.

package division

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

// NewDivisionPostParams creates a new DivisionPostParams object
// with the default values initialized.
func NewDivisionPostParams() *DivisionPostParams {
	var ()
	return &DivisionPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDivisionPostParamsWithTimeout creates a new DivisionPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDivisionPostParamsWithTimeout(timeout time.Duration) *DivisionPostParams {
	var ()
	return &DivisionPostParams{

		timeout: timeout,
	}
}

// NewDivisionPostParamsWithContext creates a new DivisionPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDivisionPostParamsWithContext(ctx context.Context) *DivisionPostParams {
	var ()
	return &DivisionPostParams{

		Context: ctx,
	}
}

// NewDivisionPostParamsWithHTTPClient creates a new DivisionPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDivisionPostParamsWithHTTPClient(client *http.Client) *DivisionPostParams {
	var ()
	return &DivisionPostParams{
		HTTPClient: client,
	}
}

/*DivisionPostParams contains all the parameters to send to the API endpoint
for the division post operation typically these are written to a http.Request
*/
type DivisionPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Division

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the division post params
func (o *DivisionPostParams) WithTimeout(timeout time.Duration) *DivisionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the division post params
func (o *DivisionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the division post params
func (o *DivisionPostParams) WithContext(ctx context.Context) *DivisionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the division post params
func (o *DivisionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the division post params
func (o *DivisionPostParams) WithHTTPClient(client *http.Client) *DivisionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the division post params
func (o *DivisionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the division post params
func (o *DivisionPostParams) WithBody(body *models.Division) *DivisionPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the division post params
func (o *DivisionPostParams) SetBody(body *models.Division) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DivisionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
