// Code generated by go-swagger; DO NOT EDIT.

package voucher_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new voucher status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voucher status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	VoucherStatusGet(params *VoucherStatusGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusGetOK, error)

	VoucherStatusPost(params *VoucherStatusPostParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusPostCreated, error)

	VoucherStatusSearch(params *VoucherStatusSearchParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VoucherStatusGet bs e t a get voucher status by ID
*/
func (a *Client) VoucherStatusGet(params *VoucherStatusGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherStatusGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherStatus_get",
		Method:             "GET",
		PathPattern:        "/voucherStatus/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherStatusGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherStatusGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherStatus_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoucherStatusPost bs e t a post new voucher status
*/
func (a *Client) VoucherStatusPost(params *VoucherStatusPostParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherStatusPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherStatus_post",
		Method:             "POST",
		PathPattern:        "/voucherStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherStatusPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherStatusPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherStatus_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoucherStatusSearch bs e t a find voucher status corresponding with sent data the voucher status is used to coordinate integration processes requires setup done by tripletex currently supports debt collection
*/
func (a *Client) VoucherStatusSearch(params *VoucherStatusSearchParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherStatusSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherStatusSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherStatus_search",
		Method:             "GET",
		PathPattern:        "/voucherStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherStatusSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherStatusSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherStatus_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
