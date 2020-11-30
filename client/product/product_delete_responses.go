// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProductDeleteReader is a Reader for the ProductDelete structure.
type ProductDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProductDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProductDeleteDefault creates a ProductDeleteDefault with default headers values
func NewProductDeleteDefault(code int) *ProductDeleteDefault {
	return &ProductDeleteDefault{
		_statusCode: code,
	}
}

/*ProductDeleteDefault handles this case with default header values.

successful operation
*/
type ProductDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the product delete default response
func (o *ProductDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProductDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /product/{id}][%d] Product_delete default ", o._statusCode)
}

func (o *ProductDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
