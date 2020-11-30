// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// BankReconciliationPutReader is a Reader for the BankReconciliationPut structure.
type BankReconciliationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankReconciliationPutOK creates a BankReconciliationPutOK with default headers values
func NewBankReconciliationPutOK() *BankReconciliationPutOK {
	return &BankReconciliationPutOK{}
}

/*BankReconciliationPutOK handles this case with default header values.

successful operation
*/
type BankReconciliationPutOK struct {
	Payload *models.ResponseWrapperBankReconciliation
}

func (o *BankReconciliationPutOK) Error() string {
	return fmt.Sprintf("[PUT /bank/reconciliation/{id}][%d] bankReconciliationPutOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationPutOK) GetPayload() *models.ResponseWrapperBankReconciliation {
	return o.Payload
}

func (o *BankReconciliationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
