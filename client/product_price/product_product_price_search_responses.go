// Code generated by go-swagger; DO NOT EDIT.

package product_price

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProductProductPriceSearchReader is a Reader for the ProductProductPriceSearch structure.
type ProductProductPriceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductProductPriceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductProductPriceSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductProductPriceSearchOK creates a ProductProductPriceSearchOK with default headers values
func NewProductProductPriceSearchOK() *ProductProductPriceSearchOK {
	return &ProductProductPriceSearchOK{}
}

/*ProductProductPriceSearchOK handles this case with default header values.

successful operation
*/
type ProductProductPriceSearchOK struct {
	Payload *models.ListResponseProductPrice
}

func (o *ProductProductPriceSearchOK) Error() string {
	return fmt.Sprintf("[GET /product/productPrice][%d] productProductPriceSearchOK  %+v", 200, o.Payload)
}

func (o *ProductProductPriceSearchOK) GetPayload() *models.ListResponseProductPrice {
	return o.Payload
}

func (o *ProductProductPriceSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProductPrice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
