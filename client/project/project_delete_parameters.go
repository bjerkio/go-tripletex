// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewProjectDeleteParams creates a new ProjectDeleteParams object
// with the default values initialized.
func NewProjectDeleteParams() *ProjectDeleteParams {
	var ()
	return &ProjectDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectDeleteParamsWithTimeout creates a new ProjectDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectDeleteParamsWithTimeout(timeout time.Duration) *ProjectDeleteParams {
	var ()
	return &ProjectDeleteParams{

		timeout: timeout,
	}
}

// NewProjectDeleteParamsWithContext creates a new ProjectDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectDeleteParamsWithContext(ctx context.Context) *ProjectDeleteParams {
	var ()
	return &ProjectDeleteParams{

		Context: ctx,
	}
}

// NewProjectDeleteParamsWithHTTPClient creates a new ProjectDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectDeleteParamsWithHTTPClient(client *http.Client) *ProjectDeleteParams {
	var ()
	return &ProjectDeleteParams{
		HTTPClient: client,
	}
}

/*ProjectDeleteParams contains all the parameters to send to the API endpoint
for the project delete operation typically these are written to a http.Request
*/
type ProjectDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project delete params
func (o *ProjectDeleteParams) WithTimeout(timeout time.Duration) *ProjectDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project delete params
func (o *ProjectDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project delete params
func (o *ProjectDeleteParams) WithContext(ctx context.Context) *ProjectDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project delete params
func (o *ProjectDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project delete params
func (o *ProjectDeleteParams) WithHTTPClient(client *http.Client) *ProjectDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project delete params
func (o *ProjectDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the project delete params
func (o *ProjectDeleteParams) WithID(id int32) *ProjectDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project delete params
func (o *ProjectDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
