// Code generated by go-swagger; DO NOT EDIT.

package supplier_invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// SupplierInvoiceGetReader is a Reader for the SupplierInvoiceGet structure.
type SupplierInvoiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplierInvoiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplierInvoiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSupplierInvoiceGetOK creates a SupplierInvoiceGetOK with default headers values
func NewSupplierInvoiceGetOK() *SupplierInvoiceGetOK {
	return &SupplierInvoiceGetOK{}
}

/*SupplierInvoiceGetOK handles this case with default header values.

successful operation
*/
type SupplierInvoiceGetOK struct {
	Payload *models.ResponseWrapperSupplierInvoice
}

func (o *SupplierInvoiceGetOK) Error() string {
	return fmt.Sprintf("[GET /supplierInvoice/{id}][%d] supplierInvoiceGetOK  %+v", 200, o.Payload)
}

func (o *SupplierInvoiceGetOK) GetPayload() *models.ResponseWrapperSupplierInvoice {
	return o.Payload
}

func (o *SupplierInvoiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSupplierInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
