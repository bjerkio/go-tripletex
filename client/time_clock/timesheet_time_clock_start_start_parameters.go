// Code generated by go-swagger; DO NOT EDIT.

package time_clock

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

// NewTimesheetTimeClockStartStartParams creates a new TimesheetTimeClockStartStartParams object
// with the default values initialized.
func NewTimesheetTimeClockStartStartParams() *TimesheetTimeClockStartStartParams {
	var (
		projectIDDefault = int32(0)
	)
	return &TimesheetTimeClockStartStartParams{
		ProjectID: &projectIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetTimeClockStartStartParamsWithTimeout creates a new TimesheetTimeClockStartStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetTimeClockStartStartParamsWithTimeout(timeout time.Duration) *TimesheetTimeClockStartStartParams {
	var (
		projectIDDefault = int32(0)
	)
	return &TimesheetTimeClockStartStartParams{
		ProjectID: &projectIDDefault,

		timeout: timeout,
	}
}

// NewTimesheetTimeClockStartStartParamsWithContext creates a new TimesheetTimeClockStartStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetTimeClockStartStartParamsWithContext(ctx context.Context) *TimesheetTimeClockStartStartParams {
	var (
		projectIdDefault = int32(0)
	)
	return &TimesheetTimeClockStartStartParams{
		ProjectID: &projectIdDefault,

		Context: ctx,
	}
}

// NewTimesheetTimeClockStartStartParamsWithHTTPClient creates a new TimesheetTimeClockStartStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetTimeClockStartStartParamsWithHTTPClient(client *http.Client) *TimesheetTimeClockStartStartParams {
	var (
		projectIdDefault = int32(0)
	)
	return &TimesheetTimeClockStartStartParams{
		ProjectID:  &projectIdDefault,
		HTTPClient: client,
	}
}

/*TimesheetTimeClockStartStartParams contains all the parameters to send to the API endpoint
for the timesheet time clock start start operation typically these are written to a http.Request
*/
type TimesheetTimeClockStartStartParams struct {

	/*ActivityID
	  Activity ID

	*/
	ActivityID int32
	/*Date
	  Optional. Default is today’s date

	*/
	Date *string
	/*EmployeeID
	  Employee ID. Defaults to ID of token owner.

	*/
	EmployeeID *int32
	/*ProjectID
	  Project ID

	*/
	ProjectID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithTimeout(timeout time.Duration) *TimesheetTimeClockStartStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithContext(ctx context.Context) *TimesheetTimeClockStartStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithHTTPClient(client *http.Client) *TimesheetTimeClockStartStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityID adds the activityID to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithActivityID(activityID int32) *TimesheetTimeClockStartStartParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetActivityID(activityID int32) {
	o.ActivityID = activityID
}

// WithDate adds the date to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithDate(date *string) *TimesheetTimeClockStartStartParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetDate(date *string) {
	o.Date = date
}

// WithEmployeeID adds the employeeID to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithEmployeeID(employeeID *int32) *TimesheetTimeClockStartStartParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithProjectID adds the projectID to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) WithProjectID(projectID *int32) *TimesheetTimeClockStartStartParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the timesheet time clock start start params
func (o *TimesheetTimeClockStartStartParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetTimeClockStartStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param activityId
	qrActivityID := o.ActivityID
	qActivityID := swag.FormatInt32(qrActivityID)
	if qActivityID != "" {
		if err := r.SetQueryParam("activityId", qActivityID); err != nil {
			return err
		}
	}

	if o.Date != nil {

		// query param date
		var qrDate string
		if o.Date != nil {
			qrDate = *o.Date
		}
		qDate := qrDate
		if qDate != "" {
			if err := r.SetQueryParam("date", qDate); err != nil {
				return err
			}
		}

	}

	if o.EmployeeID != nil {

		// query param employeeId
		var qrEmployeeID int32
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := swag.FormatInt32(qrEmployeeID)
		if qEmployeeID != "" {
			if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
				return err
			}
		}

	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
