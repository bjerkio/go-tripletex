// Code generated by go-swagger; DO NOT EDIT.

package group_relation

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

// NewProductGroupRelationListPostListParams creates a new ProductGroupRelationListPostListParams object
// with the default values initialized.
func NewProductGroupRelationListPostListParams() *ProductGroupRelationListPostListParams {
	var ()
	return &ProductGroupRelationListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupRelationListPostListParamsWithTimeout creates a new ProductGroupRelationListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupRelationListPostListParamsWithTimeout(timeout time.Duration) *ProductGroupRelationListPostListParams {
	var ()
	return &ProductGroupRelationListPostListParams{

		timeout: timeout,
	}
}

// NewProductGroupRelationListPostListParamsWithContext creates a new ProductGroupRelationListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupRelationListPostListParamsWithContext(ctx context.Context) *ProductGroupRelationListPostListParams {
	var ()
	return &ProductGroupRelationListPostListParams{

		Context: ctx,
	}
}

// NewProductGroupRelationListPostListParamsWithHTTPClient creates a new ProductGroupRelationListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupRelationListPostListParamsWithHTTPClient(client *http.Client) *ProductGroupRelationListPostListParams {
	var ()
	return &ProductGroupRelationListPostListParams{
		HTTPClient: client,
	}
}

/*ProductGroupRelationListPostListParams contains all the parameters to send to the API endpoint
for the product group relation list post list operation typically these are written to a http.Request
*/
type ProductGroupRelationListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.ProductGroupRelation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) WithTimeout(timeout time.Duration) *ProductGroupRelationListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) WithContext(ctx context.Context) *ProductGroupRelationListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) WithHTTPClient(client *http.Client) *ProductGroupRelationListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) WithBody(body []*models.ProductGroupRelation) *ProductGroupRelationListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product group relation list post list params
func (o *ProductGroupRelationListPostListParams) SetBody(body []*models.ProductGroupRelation) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupRelationListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
