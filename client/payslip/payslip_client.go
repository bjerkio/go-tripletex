// Code generated by go-swagger; DO NOT EDIT.

package payslip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payslip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payslip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SalaryPayslipPdfDownloadPdf(params *SalaryPayslipPdfDownloadPdfParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipPdfDownloadPdfOK, error)

	SalaryPayslipGet(params *SalaryPayslipGetParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipGetOK, error)

	SalaryPayslipSearch(params *SalaryPayslipSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SalaryPayslipPdfDownloadPdf bs e t a find payslip p d f document by ID
*/
func (a *Client) SalaryPayslipPdfDownloadPdf(params *SalaryPayslipPdfDownloadPdfParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipPdfDownloadPdfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalaryPayslipPdfDownloadPdfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalaryPayslipPdf_downloadPdf",
		Method:             "GET",
		PathPattern:        "/salary/payslip/{id}/pdf",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalaryPayslipPdfDownloadPdfReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalaryPayslipPdfDownloadPdfOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalaryPayslipPdf_downloadPdf: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalaryPayslipGet bs e t a find payslip by ID
*/
func (a *Client) SalaryPayslipGet(params *SalaryPayslipGetParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalaryPayslipGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalaryPayslip_get",
		Method:             "GET",
		PathPattern:        "/salary/payslip/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalaryPayslipGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalaryPayslipGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalaryPayslip_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalaryPayslipSearch bs e t a find payslips corresponding with sent data
*/
func (a *Client) SalaryPayslipSearch(params *SalaryPayslipSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SalaryPayslipSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalaryPayslipSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalaryPayslip_search",
		Method:             "GET",
		PathPattern:        "/salary/payslip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalaryPayslipSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalaryPayslipSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalaryPayslip_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
