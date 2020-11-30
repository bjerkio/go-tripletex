// Code generated by go-swagger; DO NOT EDIT.

package payment_type_out

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment type out API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment type out API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LedgerPaymentTypeOutListPostList(params *LedgerPaymentTypeOutListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutListPostListCreated, error)

	LedgerPaymentTypeOutListPutList(params *LedgerPaymentTypeOutListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutListPutListOK, error)

	LedgerPaymentTypeOutDelete(params *LedgerPaymentTypeOutDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	LedgerPaymentTypeOutGet(params *LedgerPaymentTypeOutGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutGetOK, error)

	LedgerPaymentTypeOutPost(params *LedgerPaymentTypeOutPostParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutPostCreated, error)

	LedgerPaymentTypeOutPut(params *LedgerPaymentTypeOutPutParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutPutOK, error)

	LedgerPaymentTypeOutSearch(params *LedgerPaymentTypeOutSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LedgerPaymentTypeOutListPostList bs e t a create multiple payment types for outgoing payments at once
*/
func (a *Client) LedgerPaymentTypeOutListPostList(params *LedgerPaymentTypeOutListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOutList_postList",
		Method:             "POST",
		PathPattern:        "/ledger/paymentTypeOut/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOutList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerPaymentTypeOutListPutList bs e t a update multiple payment types for outgoing payments at once
*/
func (a *Client) LedgerPaymentTypeOutListPutList(params *LedgerPaymentTypeOutListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOutList_putList",
		Method:             "PUT",
		PathPattern:        "/ledger/paymentTypeOut/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOutList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerPaymentTypeOutDelete bs e t a delete payment type for outgoing payments by ID
*/
func (a *Client) LedgerPaymentTypeOutDelete(params *LedgerPaymentTypeOutDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOut_delete",
		Method:             "DELETE",
		PathPattern:        "/ledger/paymentTypeOut/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  LedgerPaymentTypeOutGet bs e t a get payment type for outgoing payments by ID
*/
func (a *Client) LedgerPaymentTypeOutGet(params *LedgerPaymentTypeOutGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOut_get",
		Method:             "GET",
		PathPattern:        "/ledger/paymentTypeOut/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOut_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerPaymentTypeOutPost bs e t a create new payment type for outgoing payments
*/
func (a *Client) LedgerPaymentTypeOutPost(params *LedgerPaymentTypeOutPostParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOut_post",
		Method:             "POST",
		PathPattern:        "/ledger/paymentTypeOut",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOut_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerPaymentTypeOutPut bs e t a update existing payment type for outgoing payments
*/
func (a *Client) LedgerPaymentTypeOutPut(params *LedgerPaymentTypeOutPutParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOut_put",
		Method:             "PUT",
		PathPattern:        "/ledger/paymentTypeOut/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOut_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerPaymentTypeOutSearch bs e t a gets payment types for outgoing payments

  This is an API endpoint for getting payment types for outgoing payments. This is equivalent to the section 'Outgoing Payments' under Accounts Settings in Tripletex. These are the payment types listed in supplier invoices, vat returns, salary payments and Tax/ENI
*/
func (a *Client) LedgerPaymentTypeOutSearch(params *LedgerPaymentTypeOutSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerPaymentTypeOutSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerPaymentTypeOutSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerPaymentTypeOut_search",
		Method:             "GET",
		PathPattern:        "/ledger/paymentTypeOut",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerPaymentTypeOutSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerPaymentTypeOutSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerPaymentTypeOut_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
