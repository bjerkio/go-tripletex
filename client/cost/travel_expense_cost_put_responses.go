// Code generated by go-swagger; DO NOT EDIT.

package cost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TravelExpenseCostPutReader is a Reader for the TravelExpenseCostPut structure.
type TravelExpenseCostPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseCostPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseCostPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseCostPutOK creates a TravelExpenseCostPutOK with default headers values
func NewTravelExpenseCostPutOK() *TravelExpenseCostPutOK {
	return &TravelExpenseCostPutOK{}
}

/*TravelExpenseCostPutOK handles this case with default header values.

successful operation
*/
type TravelExpenseCostPutOK struct {
	Payload *models.ResponseWrapperCost
}

func (o *TravelExpenseCostPutOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/cost/{id}][%d] travelExpenseCostPutOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseCostPutOK) GetPayload() *models.ResponseWrapperCost {
	return o.Payload
}

func (o *TravelExpenseCostPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
