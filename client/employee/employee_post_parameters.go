// Code generated by go-swagger; DO NOT EDIT.

package employee

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

// NewEmployeePostParams creates a new EmployeePostParams object
// with the default values initialized.
func NewEmployeePostParams() *EmployeePostParams {
	var ()
	return &EmployeePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeePostParamsWithTimeout creates a new EmployeePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeePostParamsWithTimeout(timeout time.Duration) *EmployeePostParams {
	var ()
	return &EmployeePostParams{

		timeout: timeout,
	}
}

// NewEmployeePostParamsWithContext creates a new EmployeePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeePostParamsWithContext(ctx context.Context) *EmployeePostParams {
	var ()
	return &EmployeePostParams{

		Context: ctx,
	}
}

// NewEmployeePostParamsWithHTTPClient creates a new EmployeePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeePostParamsWithHTTPClient(client *http.Client) *EmployeePostParams {
	var ()
	return &EmployeePostParams{
		HTTPClient: client,
	}
}

/*EmployeePostParams contains all the parameters to send to the API endpoint
for the employee post operation typically these are written to a http.Request
*/
type EmployeePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Employee

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee post params
func (o *EmployeePostParams) WithTimeout(timeout time.Duration) *EmployeePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee post params
func (o *EmployeePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee post params
func (o *EmployeePostParams) WithContext(ctx context.Context) *EmployeePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee post params
func (o *EmployeePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee post params
func (o *EmployeePostParams) WithHTTPClient(client *http.Client) *EmployeePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee post params
func (o *EmployeePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee post params
func (o *EmployeePostParams) WithBody(body *models.Employee) *EmployeePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee post params
func (o *EmployeePostParams) SetBody(body *models.Employee) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
