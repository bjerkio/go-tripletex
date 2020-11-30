// Code generated by go-swagger; DO NOT EDIT.

package payment_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// BankReconciliationPaymentTypeGetReader is a Reader for the BankReconciliationPaymentTypeGet structure.
type BankReconciliationPaymentTypeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationPaymentTypeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationPaymentTypeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankReconciliationPaymentTypeGetOK creates a BankReconciliationPaymentTypeGetOK with default headers values
func NewBankReconciliationPaymentTypeGetOK() *BankReconciliationPaymentTypeGetOK {
	return &BankReconciliationPaymentTypeGetOK{}
}

/*BankReconciliationPaymentTypeGetOK handles this case with default header values.

successful operation
*/
type BankReconciliationPaymentTypeGetOK struct {
	Payload *models.ResponseWrapperBankReconciliationPaymentType
}

func (o *BankReconciliationPaymentTypeGetOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/paymentType/{id}][%d] bankReconciliationPaymentTypeGetOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationPaymentTypeGetOK) GetPayload() *models.ResponseWrapperBankReconciliationPaymentType {
	return o.Payload
}

func (o *BankReconciliationPaymentTypeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliationPaymentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
