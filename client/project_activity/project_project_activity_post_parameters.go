// Code generated by go-swagger; DO NOT EDIT.

package project_activity

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

// NewProjectProjectActivityPostParams creates a new ProjectProjectActivityPostParams object
// with the default values initialized.
func NewProjectProjectActivityPostParams() *ProjectProjectActivityPostParams {
	var ()
	return &ProjectProjectActivityPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectProjectActivityPostParamsWithTimeout creates a new ProjectProjectActivityPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectProjectActivityPostParamsWithTimeout(timeout time.Duration) *ProjectProjectActivityPostParams {
	var ()
	return &ProjectProjectActivityPostParams{

		timeout: timeout,
	}
}

// NewProjectProjectActivityPostParamsWithContext creates a new ProjectProjectActivityPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectProjectActivityPostParamsWithContext(ctx context.Context) *ProjectProjectActivityPostParams {
	var ()
	return &ProjectProjectActivityPostParams{

		Context: ctx,
	}
}

// NewProjectProjectActivityPostParamsWithHTTPClient creates a new ProjectProjectActivityPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectProjectActivityPostParamsWithHTTPClient(client *http.Client) *ProjectProjectActivityPostParams {
	var ()
	return &ProjectProjectActivityPostParams{
		HTTPClient: client,
	}
}

/*ProjectProjectActivityPostParams contains all the parameters to send to the API endpoint
for the project project activity post operation typically these are written to a http.Request
*/
type ProjectProjectActivityPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProjectActivity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project project activity post params
func (o *ProjectProjectActivityPostParams) WithTimeout(timeout time.Duration) *ProjectProjectActivityPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project project activity post params
func (o *ProjectProjectActivityPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project project activity post params
func (o *ProjectProjectActivityPostParams) WithContext(ctx context.Context) *ProjectProjectActivityPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project project activity post params
func (o *ProjectProjectActivityPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project project activity post params
func (o *ProjectProjectActivityPostParams) WithHTTPClient(client *http.Client) *ProjectProjectActivityPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project project activity post params
func (o *ProjectProjectActivityPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project project activity post params
func (o *ProjectProjectActivityPostParams) WithBody(body *models.ProjectActivity) *ProjectProjectActivityPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project project activity post params
func (o *ProjectProjectActivityPostParams) SetBody(body *models.ProjectActivity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectProjectActivityPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
