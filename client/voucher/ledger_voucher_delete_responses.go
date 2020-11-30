// Code generated by go-swagger; DO NOT EDIT.

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LedgerVoucherDeleteReader is a Reader for the LedgerVoucherDelete structure.
type LedgerVoucherDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLedgerVoucherDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLedgerVoucherDeleteDefault creates a LedgerVoucherDeleteDefault with default headers values
func NewLedgerVoucherDeleteDefault(code int) *LedgerVoucherDeleteDefault {
	return &LedgerVoucherDeleteDefault{
		_statusCode: code,
	}
}

/*LedgerVoucherDeleteDefault handles this case with default header values.

successful operation
*/
type LedgerVoucherDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the ledger voucher delete default response
func (o *LedgerVoucherDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LedgerVoucherDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ledger/voucher/{id}][%d] LedgerVoucher_delete default ", o._statusCode)
}

func (o *LedgerVoucherDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
