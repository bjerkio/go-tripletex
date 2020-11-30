// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TravelExpenseUndeliverUndeliverReader is a Reader for the TravelExpenseUndeliverUndeliver structure.
type TravelExpenseUndeliverUndeliverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseUndeliverUndeliverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseUndeliverUndeliverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseUndeliverUndeliverOK creates a TravelExpenseUndeliverUndeliverOK with default headers values
func NewTravelExpenseUndeliverUndeliverOK() *TravelExpenseUndeliverUndeliverOK {
	return &TravelExpenseUndeliverUndeliverOK{}
}

/*TravelExpenseUndeliverUndeliverOK handles this case with default header values.

successful operation
*/
type TravelExpenseUndeliverUndeliverOK struct {
	Payload *models.ListResponseTravelExpense
}

func (o *TravelExpenseUndeliverUndeliverOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/:undeliver][%d] travelExpenseUndeliverUndeliverOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseUndeliverUndeliverOK) GetPayload() *models.ListResponseTravelExpense {
	return o.Payload
}

func (o *TravelExpenseUndeliverUndeliverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseTravelExpense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
