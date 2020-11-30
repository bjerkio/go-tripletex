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

// LedgerVoucherOptionsOptionsReader is a Reader for the LedgerVoucherOptionsOptions structure.
type LedgerVoucherOptionsOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherOptionsOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerVoucherOptionsOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLedgerVoucherOptionsOptionsOK creates a LedgerVoucherOptionsOptionsOK with default headers values
func NewLedgerVoucherOptionsOptionsOK() *LedgerVoucherOptionsOptionsOK {
	return &LedgerVoucherOptionsOptionsOK{}
}

/*LedgerVoucherOptionsOptionsOK handles this case with default header values.

successful operation
*/
type LedgerVoucherOptionsOptionsOK struct {
	Payload *models.ResponseWrapperVoucherOptions
}

func (o *LedgerVoucherOptionsOptionsOK) Error() string {
	return fmt.Sprintf("[GET /ledger/voucher/{id}/options][%d] ledgerVoucherOptionsOptionsOK  %+v", 200, o.Payload)
}

func (o *LedgerVoucherOptionsOptionsOK) GetPayload() *models.ResponseWrapperVoucherOptions {
	return o.Payload
}

func (o *LedgerVoucherOptionsOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperVoucherOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
