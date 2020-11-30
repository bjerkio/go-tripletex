// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewProductGroupPostParams creates a new ProductGroupPostParams object
// with the default values initialized.
func NewProductGroupPostParams() *ProductGroupPostParams {
	var ()
	return &ProductGroupPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupPostParamsWithTimeout creates a new ProductGroupPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupPostParamsWithTimeout(timeout time.Duration) *ProductGroupPostParams {
	var ()
	return &ProductGroupPostParams{

		timeout: timeout,
	}
}

// NewProductGroupPostParamsWithContext creates a new ProductGroupPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupPostParamsWithContext(ctx context.Context) *ProductGroupPostParams {
	var ()
	return &ProductGroupPostParams{

		Context: ctx,
	}
}

// NewProductGroupPostParamsWithHTTPClient creates a new ProductGroupPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupPostParamsWithHTTPClient(client *http.Client) *ProductGroupPostParams {
	var ()
	return &ProductGroupPostParams{
		HTTPClient: client,
	}
}

/*ProductGroupPostParams contains all the parameters to send to the API endpoint
for the product group post operation typically these are written to a http.Request
*/
type ProductGroupPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProductGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group post params
func (o *ProductGroupPostParams) WithTimeout(timeout time.Duration) *ProductGroupPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group post params
func (o *ProductGroupPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group post params
func (o *ProductGroupPostParams) WithContext(ctx context.Context) *ProductGroupPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group post params
func (o *ProductGroupPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group post params
func (o *ProductGroupPostParams) WithHTTPClient(client *http.Client) *ProductGroupPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group post params
func (o *ProductGroupPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product group post params
func (o *ProductGroupPostParams) WithBody(body *models.ProductGroup) *ProductGroupPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product group post params
func (o *ProductGroupPostParams) SetBody(body *models.ProductGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
