// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewTimesheetSettingsGetParams creates a new TimesheetSettingsGetParams object
// with the default values initialized.
func NewTimesheetSettingsGetParams() *TimesheetSettingsGetParams {
	var ()
	return &TimesheetSettingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetSettingsGetParamsWithTimeout creates a new TimesheetSettingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetSettingsGetParamsWithTimeout(timeout time.Duration) *TimesheetSettingsGetParams {
	var ()
	return &TimesheetSettingsGetParams{

		timeout: timeout,
	}
}

// NewTimesheetSettingsGetParamsWithContext creates a new TimesheetSettingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetSettingsGetParamsWithContext(ctx context.Context) *TimesheetSettingsGetParams {
	var ()
	return &TimesheetSettingsGetParams{

		Context: ctx,
	}
}

// NewTimesheetSettingsGetParamsWithHTTPClient creates a new TimesheetSettingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetSettingsGetParamsWithHTTPClient(client *http.Client) *TimesheetSettingsGetParams {
	var ()
	return &TimesheetSettingsGetParams{
		HTTPClient: client,
	}
}

/*TimesheetSettingsGetParams contains all the parameters to send to the API endpoint
for the timesheet settings get operation typically these are written to a http.Request
*/
type TimesheetSettingsGetParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet settings get params
func (o *TimesheetSettingsGetParams) WithTimeout(timeout time.Duration) *TimesheetSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet settings get params
func (o *TimesheetSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet settings get params
func (o *TimesheetSettingsGetParams) WithContext(ctx context.Context) *TimesheetSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet settings get params
func (o *TimesheetSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet settings get params
func (o *TimesheetSettingsGetParams) WithHTTPClient(client *http.Client) *TimesheetSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet settings get params
func (o *TimesheetSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the timesheet settings get params
func (o *TimesheetSettingsGetParams) WithFields(fields *string) *TimesheetSettingsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timesheet settings get params
func (o *TimesheetSettingsGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
