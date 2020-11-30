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

// LedgerVoucherSendToInboxSendToInboxReader is a Reader for the LedgerVoucherSendToInboxSendToInbox structure.
type LedgerVoucherSendToInboxSendToInboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherSendToInboxSendToInboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerVoucherSendToInboxSendToInboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLedgerVoucherSendToInboxSendToInboxOK creates a LedgerVoucherSendToInboxSendToInboxOK with default headers values
func NewLedgerVoucherSendToInboxSendToInboxOK() *LedgerVoucherSendToInboxSendToInboxOK {
	return &LedgerVoucherSendToInboxSendToInboxOK{}
}

/*LedgerVoucherSendToInboxSendToInboxOK handles this case with default header values.

successful operation
*/
type LedgerVoucherSendToInboxSendToInboxOK struct {
	Payload *models.ResponseWrapperVoucher
}

func (o *LedgerVoucherSendToInboxSendToInboxOK) Error() string {
	return fmt.Sprintf("[PUT /ledger/voucher/{id}/:sendToInbox][%d] ledgerVoucherSendToInboxSendToInboxOK  %+v", 200, o.Payload)
}

func (o *LedgerVoucherSendToInboxSendToInboxOK) GetPayload() *models.ResponseWrapperVoucher {
	return o.Payload
}

func (o *LedgerVoucherSendToInboxSendToInboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperVoucher)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
