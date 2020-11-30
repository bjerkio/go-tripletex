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
)

// NewProjectProjectActivityListDeleteByIdsParams creates a new ProjectProjectActivityListDeleteByIdsParams object
// with the default values initialized.
func NewProjectProjectActivityListDeleteByIdsParams() *ProjectProjectActivityListDeleteByIdsParams {
	var ()
	return &ProjectProjectActivityListDeleteByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectProjectActivityListDeleteByIdsParamsWithTimeout creates a new ProjectProjectActivityListDeleteByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectProjectActivityListDeleteByIdsParamsWithTimeout(timeout time.Duration) *ProjectProjectActivityListDeleteByIdsParams {
	var ()
	return &ProjectProjectActivityListDeleteByIdsParams{

		timeout: timeout,
	}
}

// NewProjectProjectActivityListDeleteByIdsParamsWithContext creates a new ProjectProjectActivityListDeleteByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectProjectActivityListDeleteByIdsParamsWithContext(ctx context.Context) *ProjectProjectActivityListDeleteByIdsParams {
	var ()
	return &ProjectProjectActivityListDeleteByIdsParams{

		Context: ctx,
	}
}

// NewProjectProjectActivityListDeleteByIdsParamsWithHTTPClient creates a new ProjectProjectActivityListDeleteByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectProjectActivityListDeleteByIdsParamsWithHTTPClient(client *http.Client) *ProjectProjectActivityListDeleteByIdsParams {
	var ()
	return &ProjectProjectActivityListDeleteByIdsParams{
		HTTPClient: client,
	}
}

/*ProjectProjectActivityListDeleteByIdsParams contains all the parameters to send to the API endpoint
for the project project activity list delete by ids operation typically these are written to a http.Request
*/
type ProjectProjectActivityListDeleteByIdsParams struct {

	/*Ids
	  ID of the elements

	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) WithTimeout(timeout time.Duration) *ProjectProjectActivityListDeleteByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) WithContext(ctx context.Context) *ProjectProjectActivityListDeleteByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) WithHTTPClient(client *http.Client) *ProjectProjectActivityListDeleteByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) WithIds(ids string) *ProjectProjectActivityListDeleteByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the project project activity list delete by ids params
func (o *ProjectProjectActivityListDeleteByIdsParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectProjectActivityListDeleteByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
