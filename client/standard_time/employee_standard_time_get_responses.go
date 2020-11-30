// Code generated by go-swagger; DO NOT EDIT.

package standard_time

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// EmployeeStandardTimeGetReader is a Reader for the EmployeeStandardTimeGet structure.
type EmployeeStandardTimeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeStandardTimeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeStandardTimeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeeStandardTimeGetOK creates a EmployeeStandardTimeGetOK with default headers values
func NewEmployeeStandardTimeGetOK() *EmployeeStandardTimeGetOK {
	return &EmployeeStandardTimeGetOK{}
}

/*EmployeeStandardTimeGetOK handles this case with default header values.

successful operation
*/
type EmployeeStandardTimeGetOK struct {
	Payload *models.ResponseWrapperStandardTime
}

func (o *EmployeeStandardTimeGetOK) Error() string {
	return fmt.Sprintf("[GET /employee/standardTime/{id}][%d] employeeStandardTimeGetOK  %+v", 200, o.Payload)
}

func (o *EmployeeStandardTimeGetOK) GetPayload() *models.ResponseWrapperStandardTime {
	return o.Payload
}

func (o *EmployeeStandardTimeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperStandardTime)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
