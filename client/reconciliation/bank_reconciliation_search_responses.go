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

// BankReconciliationSearchReader is a Reader for the BankReconciliationSearch structure.
type BankReconciliationSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankReconciliationSearchOK creates a BankReconciliationSearchOK with default headers values
func NewBankReconciliationSearchOK() *BankReconciliationSearchOK {
	return &BankReconciliationSearchOK{}
}

/*BankReconciliationSearchOK handles this case with default header values.

successful operation
*/
type BankReconciliationSearchOK struct {
	Payload *models.ListResponseBankReconciliation
}

func (o *BankReconciliationSearchOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation][%d] bankReconciliationSearchOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationSearchOK) GetPayload() *models.ListResponseBankReconciliation {
	return o.Payload
}

func (o *BankReconciliationSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankReconciliation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
