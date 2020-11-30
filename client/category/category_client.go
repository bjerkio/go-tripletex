// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new category API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for category API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CustomerCategoryGet(params *CustomerCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryGetOK, error)

	CustomerCategoryPost(params *CustomerCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryPostCreated, error)

	CustomerCategoryPut(params *CustomerCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryPutOK, error)

	CustomerCategorySearch(params *CustomerCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategorySearchOK, error)

	EmployeeCategoryListDeleteByIds(params *EmployeeCategoryListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	EmployeeCategoryListPostList(params *EmployeeCategoryListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryListPostListCreated, error)

	EmployeeCategoryListPutList(params *EmployeeCategoryListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryListPutListOK, error)

	EmployeeCategoryDelete(params *EmployeeCategoryDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	EmployeeCategoryGet(params *EmployeeCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryGetOK, error)

	EmployeeCategoryPost(params *EmployeeCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryPostCreated, error)

	EmployeeCategoryPut(params *EmployeeCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryPutOK, error)

	EmployeeCategorySearch(params *EmployeeCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategorySearchOK, error)

	ProjectCategoryGet(params *ProjectCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryGetOK, error)

	ProjectCategoryPost(params *ProjectCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryPostCreated, error)

	ProjectCategoryPut(params *ProjectCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryPutOK, error)

	ProjectCategorySearch(params *ProjectCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategorySearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CustomerCategoryGet finds customer supplier category by ID
*/
func (a *Client) CustomerCategoryGet(params *CustomerCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCategoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerCategory_get",
		Method:             "GET",
		PathPattern:        "/customer/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerCategoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerCategoryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomerCategory_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomerCategoryPost adds new customer supplier category
*/
func (a *Client) CustomerCategoryPost(params *CustomerCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCategoryPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerCategory_post",
		Method:             "POST",
		PathPattern:        "/customer/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerCategoryPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerCategoryPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomerCategory_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomerCategoryPut updates customer supplier category
*/
func (a *Client) CustomerCategoryPut(params *CustomerCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategoryPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCategoryPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerCategory_put",
		Method:             "PUT",
		PathPattern:        "/customer/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerCategoryPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerCategoryPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomerCategory_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomerCategorySearch finds customer supplier categories corresponding with sent data
*/
func (a *Client) CustomerCategorySearch(params *CustomerCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerCategorySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerCategorySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerCategory_search",
		Method:             "GET",
		PathPattern:        "/customer/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerCategorySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerCategorySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomerCategory_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategoryListDeleteByIds bs e t a delete multiple employee categories
*/
func (a *Client) EmployeeCategoryListDeleteByIds(params *EmployeeCategoryListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategoryList_deleteByIds",
		Method:             "DELETE",
		PathPattern:        "/employee/category/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryListDeleteByIdsReader{formats: a.formats},
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
  EmployeeCategoryListPostList bs e t a create new employee categories
*/
func (a *Client) EmployeeCategoryListPostList(params *EmployeeCategoryListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategoryList_postList",
		Method:             "POST",
		PathPattern:        "/employee/category/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategoryListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategoryList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategoryListPutList bs e t a update multiple employee categories
*/
func (a *Client) EmployeeCategoryListPutList(params *EmployeeCategoryListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategoryList_putList",
		Method:             "PUT",
		PathPattern:        "/employee/category/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategoryListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategoryList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategoryDelete bs e t a delete employee category by ID
*/
func (a *Client) EmployeeCategoryDelete(params *EmployeeCategoryDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategory_delete",
		Method:             "DELETE",
		PathPattern:        "/employee/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryDeleteReader{formats: a.formats},
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
  EmployeeCategoryGet bs e t a get employee category by ID
*/
func (a *Client) EmployeeCategoryGet(params *EmployeeCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategory_get",
		Method:             "GET",
		PathPattern:        "/employee/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategoryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategory_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategoryPost bs e t a create a new employee category
*/
func (a *Client) EmployeeCategoryPost(params *EmployeeCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategory_post",
		Method:             "POST",
		PathPattern:        "/employee/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategoryPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategory_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategoryPut bs e t a update employee category information
*/
func (a *Client) EmployeeCategoryPut(params *EmployeeCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategoryPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategoryPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategory_put",
		Method:             "PUT",
		PathPattern:        "/employee/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategoryPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategoryPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategory_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeCategorySearch bs e t a find employee category corresponding with sent data
*/
func (a *Client) EmployeeCategorySearch(params *EmployeeCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeCategorySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeCategorySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeCategory_search",
		Method:             "GET",
		PathPattern:        "/employee/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeCategorySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeCategorySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeCategory_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectCategoryGet finds project category by ID
*/
func (a *Client) ProjectCategoryGet(params *ProjectCategoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectCategoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectCategory_get",
		Method:             "GET",
		PathPattern:        "/project/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectCategoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectCategoryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectCategory_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectCategoryPost adds new project category
*/
func (a *Client) ProjectCategoryPost(params *ProjectCategoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectCategoryPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectCategory_post",
		Method:             "POST",
		PathPattern:        "/project/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectCategoryPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectCategoryPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectCategory_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectCategoryPut updates project category
*/
func (a *Client) ProjectCategoryPut(params *ProjectCategoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategoryPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectCategoryPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectCategory_put",
		Method:             "PUT",
		PathPattern:        "/project/category/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectCategoryPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectCategoryPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectCategory_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectCategorySearch finds project categories corresponding with sent data
*/
func (a *Client) ProjectCategorySearch(params *ProjectCategorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectCategorySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectCategorySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectCategory_search",
		Method:             "GET",
		PathPattern:        "/project/category",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectCategorySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectCategorySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectCategory_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
