// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// NewCustomerPostParams creates a new CustomerPostParams object
// with the default values initialized.
func NewCustomerPostParams() *CustomerPostParams {
	var ()
	return &CustomerPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPostParamsWithTimeout creates a new CustomerPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerPostParamsWithTimeout(timeout time.Duration) *CustomerPostParams {
	var ()
	return &CustomerPostParams{

		timeout: timeout,
	}
}

// NewCustomerPostParamsWithContext creates a new CustomerPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerPostParamsWithContext(ctx context.Context) *CustomerPostParams {
	var ()
	return &CustomerPostParams{

		Context: ctx,
	}
}

// NewCustomerPostParamsWithHTTPClient creates a new CustomerPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerPostParamsWithHTTPClient(client *http.Client) *CustomerPostParams {
	var ()
	return &CustomerPostParams{
		HTTPClient: client,
	}
}

/*CustomerPostParams contains all the parameters to send to the API endpoint
for the customer post operation typically these are written to a http.Request
*/
type CustomerPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Customer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer post params
func (o *CustomerPostParams) WithTimeout(timeout time.Duration) *CustomerPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer post params
func (o *CustomerPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer post params
func (o *CustomerPostParams) WithContext(ctx context.Context) *CustomerPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer post params
func (o *CustomerPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer post params
func (o *CustomerPostParams) WithHTTPClient(client *http.Client) *CustomerPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer post params
func (o *CustomerPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the customer post params
func (o *CustomerPostParams) WithBody(body *models.Customer) *CustomerPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the customer post params
func (o *CustomerPostParams) SetBody(body *models.Customer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
