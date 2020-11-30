// Code generated by go-swagger; DO NOT EDIT.

package period

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectPeriodInvoicedInvoicedReader is a Reader for the ProjectPeriodInvoicedInvoiced structure.
type ProjectPeriodInvoicedInvoicedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectPeriodInvoicedInvoicedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectPeriodInvoicedInvoicedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectPeriodInvoicedInvoicedOK creates a ProjectPeriodInvoicedInvoicedOK with default headers values
func NewProjectPeriodInvoicedInvoicedOK() *ProjectPeriodInvoicedInvoicedOK {
	return &ProjectPeriodInvoicedInvoicedOK{}
}

/*ProjectPeriodInvoicedInvoicedOK handles this case with default header values.

successful operation
*/
type ProjectPeriodInvoicedInvoicedOK struct {
	Payload *models.ResponseWrapperProjectPeriodInvoiced
}

func (o *ProjectPeriodInvoicedInvoicedOK) Error() string {
	return fmt.Sprintf("[GET /project/{id}/period/invoiced][%d] projectPeriodInvoicedInvoicedOK  %+v", 200, o.Payload)
}

func (o *ProjectPeriodInvoicedInvoicedOK) GetPayload() *models.ResponseWrapperProjectPeriodInvoiced {
	return o.Payload
}

func (o *ProjectPeriodInvoicedInvoicedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectPeriodInvoiced)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
