// Code generated by go-swagger; DO NOT EDIT.

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// LedgerVoucherPostReader is a Reader for the LedgerVoucherPost structure.
type LedgerVoucherPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLedgerVoucherPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLedgerVoucherPostCreated creates a LedgerVoucherPostCreated with default headers values
func NewLedgerVoucherPostCreated() *LedgerVoucherPostCreated {
	return &LedgerVoucherPostCreated{}
}

/*LedgerVoucherPostCreated handles this case with default header values.

successfully created
*/
type LedgerVoucherPostCreated struct {
	Payload *models.ResponseWrapperVoucher
}

func (o *LedgerVoucherPostCreated) Error() string {
	return fmt.Sprintf("[POST /ledger/voucher][%d] ledgerVoucherPostCreated  %+v", 201, o.Payload)
}

func (o *LedgerVoucherPostCreated) GetPayload() *models.ResponseWrapperVoucher {
	return o.Payload
}

func (o *LedgerVoucherPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperVoucher)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
