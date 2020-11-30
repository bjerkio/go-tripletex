// Code generated by go-swagger; DO NOT EDIT.

package hourly_cost_and_rate

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

// NewEmployeeHourlyCostAndRatePostParams creates a new EmployeeHourlyCostAndRatePostParams object
// with the default values initialized.
func NewEmployeeHourlyCostAndRatePostParams() *EmployeeHourlyCostAndRatePostParams {
	var ()
	return &EmployeeHourlyCostAndRatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeHourlyCostAndRatePostParamsWithTimeout creates a new EmployeeHourlyCostAndRatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeHourlyCostAndRatePostParamsWithTimeout(timeout time.Duration) *EmployeeHourlyCostAndRatePostParams {
	var ()
	return &EmployeeHourlyCostAndRatePostParams{

		timeout: timeout,
	}
}

// NewEmployeeHourlyCostAndRatePostParamsWithContext creates a new EmployeeHourlyCostAndRatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeHourlyCostAndRatePostParamsWithContext(ctx context.Context) *EmployeeHourlyCostAndRatePostParams {
	var ()
	return &EmployeeHourlyCostAndRatePostParams{

		Context: ctx,
	}
}

// NewEmployeeHourlyCostAndRatePostParamsWithHTTPClient creates a new EmployeeHourlyCostAndRatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeHourlyCostAndRatePostParamsWithHTTPClient(client *http.Client) *EmployeeHourlyCostAndRatePostParams {
	var ()
	return &EmployeeHourlyCostAndRatePostParams{
		HTTPClient: client,
	}
}

/*EmployeeHourlyCostAndRatePostParams contains all the parameters to send to the API endpoint
for the employee hourly cost and rate post operation typically these are written to a http.Request
*/
type EmployeeHourlyCostAndRatePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.HourlyCostAndRate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) WithTimeout(timeout time.Duration) *EmployeeHourlyCostAndRatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) WithContext(ctx context.Context) *EmployeeHourlyCostAndRatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) WithHTTPClient(client *http.Client) *EmployeeHourlyCostAndRatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) WithBody(body *models.HourlyCostAndRate) *EmployeeHourlyCostAndRatePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee hourly cost and rate post params
func (o *EmployeeHourlyCostAndRatePostParams) SetBody(body *models.HourlyCostAndRate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeHourlyCostAndRatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
