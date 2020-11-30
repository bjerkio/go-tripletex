// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new match API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for match API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BankReconciliationMatchSuggestSuggest(params *BankReconciliationMatchSuggestSuggestParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchSuggestSuggestOK, error)

	BankReconciliationMatchDelete(params *BankReconciliationMatchDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	BankReconciliationMatchGet(params *BankReconciliationMatchGetParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchGetOK, error)

	BankReconciliationMatchPost(params *BankReconciliationMatchPostParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchPostCreated, error)

	BankReconciliationMatchPut(params *BankReconciliationMatchPutParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchPutOK, error)

	BankReconciliationMatchSearch(params *BankReconciliationMatchSearchParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BankReconciliationMatchSuggestSuggest bs e t a suggest matches for a bank reconciliation by ID
*/
func (a *Client) BankReconciliationMatchSuggestSuggest(params *BankReconciliationMatchSuggestSuggestParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchSuggestSuggestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchSuggestSuggestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatchSuggest_suggest",
		Method:             "PUT",
		PathPattern:        "/bank/reconciliation/match/:suggest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchSuggestSuggestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationMatchSuggestSuggestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationMatchSuggest_suggest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BankReconciliationMatchDelete bs e t a delete a bank reconciliation match by ID
*/
func (a *Client) BankReconciliationMatchDelete(params *BankReconciliationMatchDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatch_delete",
		Method:             "DELETE",
		PathPattern:        "/bank/reconciliation/match/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchDeleteReader{formats: a.formats},
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
  BankReconciliationMatchGet bs e t a get bank reconciliation match by ID
*/
func (a *Client) BankReconciliationMatchGet(params *BankReconciliationMatchGetParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatch_get",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/match/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationMatchGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationMatch_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BankReconciliationMatchPost bs e t a create a bank reconciliation match
*/
func (a *Client) BankReconciliationMatchPost(params *BankReconciliationMatchPostParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatch_post",
		Method:             "POST",
		PathPattern:        "/bank/reconciliation/match",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationMatchPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationMatch_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BankReconciliationMatchPut bs e t a update a bank reconciliation match by ID
*/
func (a *Client) BankReconciliationMatchPut(params *BankReconciliationMatchPutParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatch_put",
		Method:             "PUT",
		PathPattern:        "/bank/reconciliation/match/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationMatchPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationMatch_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BankReconciliationMatchSearch bs e t a find bank reconciliation match corresponding with sent data
*/
func (a *Client) BankReconciliationMatchSearch(params *BankReconciliationMatchSearchParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationMatchSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationMatchSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationMatch_search",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/match",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationMatchSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationMatchSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationMatch_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
