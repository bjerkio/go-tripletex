// Code generated by go-swagger; DO NOT EDIT.

package division

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new division API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for division API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DivisionListPostList(params *DivisionListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionListPostListCreated, error)

	DivisionListPutList(params *DivisionListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionListPutListOK, error)

	DivisionPost(params *DivisionPostParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionPostCreated, error)

	DivisionPut(params *DivisionPutParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionPutOK, error)

	DivisionSearch(params *DivisionSearchParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DivisionListPostList bs e t a create divisions
*/
func (a *Client) DivisionListPostList(params *DivisionListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDivisionListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DivisionList_postList",
		Method:             "POST",
		PathPattern:        "/division/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DivisionListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DivisionListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DivisionList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DivisionListPutList bs e t a update multiple divisions
*/
func (a *Client) DivisionListPutList(params *DivisionListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDivisionListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DivisionList_putList",
		Method:             "PUT",
		PathPattern:        "/division/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DivisionListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DivisionListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DivisionList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DivisionPost bs e t a create division
*/
func (a *Client) DivisionPost(params *DivisionPostParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDivisionPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Division_post",
		Method:             "POST",
		PathPattern:        "/division",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DivisionPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DivisionPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Division_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DivisionPut bs e t a update division information
*/
func (a *Client) DivisionPut(params *DivisionPutParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDivisionPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Division_put",
		Method:             "PUT",
		PathPattern:        "/division/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DivisionPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DivisionPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Division_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DivisionSearch bs e t a get divisions
*/
func (a *Client) DivisionSearch(params *DivisionSearchParams, authInfo runtime.ClientAuthInfoWriter) (*DivisionSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDivisionSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Division_search",
		Method:             "GET",
		PathPattern:        "/division",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DivisionSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DivisionSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Division_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
