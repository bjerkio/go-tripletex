// Code generated by go-swagger; DO NOT EDIT.

package payment_type_out

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// LedgerPaymentTypeOutGetReader is a Reader for the LedgerPaymentTypeOutGet structure.
type LedgerPaymentTypeOutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerPaymentTypeOutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerPaymentTypeOutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLedgerPaymentTypeOutGetOK creates a LedgerPaymentTypeOutGetOK with default headers values
func NewLedgerPaymentTypeOutGetOK() *LedgerPaymentTypeOutGetOK {
	return &LedgerPaymentTypeOutGetOK{}
}

/*LedgerPaymentTypeOutGetOK handles this case with default header values.

successful operation
*/
type LedgerPaymentTypeOutGetOK struct {
	Payload *models.ResponseWrapperPaymentTypeOut
}

func (o *LedgerPaymentTypeOutGetOK) Error() string {
	return fmt.Sprintf("[GET /ledger/paymentTypeOut/{id}][%d] ledgerPaymentTypeOutGetOK  %+v", 200, o.Payload)
}

func (o *LedgerPaymentTypeOutGetOK) GetPayload() *models.ResponseWrapperPaymentTypeOut {
	return o.Payload
}

func (o *LedgerPaymentTypeOutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPaymentTypeOut)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
