// Code generated by go-swagger; DO NOT EDIT.

package month

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TimesheetMonthGetReader is a Reader for the TimesheetMonthGet structure.
type TimesheetMonthGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetMonthGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetMonthGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetMonthGetOK creates a TimesheetMonthGetOK with default headers values
func NewTimesheetMonthGetOK() *TimesheetMonthGetOK {
	return &TimesheetMonthGetOK{}
}

/*TimesheetMonthGetOK handles this case with default header values.

successful operation
*/
type TimesheetMonthGetOK struct {
	Payload *models.ResponseWrapperMonthlyStatus
}

func (o *TimesheetMonthGetOK) Error() string {
	return fmt.Sprintf("[GET /timesheet/month/{id}][%d] timesheetMonthGetOK  %+v", 200, o.Payload)
}

func (o *TimesheetMonthGetOK) GetPayload() *models.ResponseWrapperMonthlyStatus {
	return o.Payload
}

func (o *TimesheetMonthGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperMonthlyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
