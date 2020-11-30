// Code generated by go-swagger; DO NOT EDIT.

package unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProductUnitDeleteReader is a Reader for the ProductUnitDelete structure.
type ProductUnitDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductUnitDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProductUnitDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProductUnitDeleteDefault creates a ProductUnitDeleteDefault with default headers values
func NewProductUnitDeleteDefault(code int) *ProductUnitDeleteDefault {
	return &ProductUnitDeleteDefault{
		_statusCode: code,
	}
}

/*ProductUnitDeleteDefault handles this case with default header values.

successful operation
*/
type ProductUnitDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the product unit delete default response
func (o *ProductUnitDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProductUnitDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /product/unit/{id}][%d] ProductUnit_delete default ", o._statusCode)
}

func (o *ProductUnitDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
