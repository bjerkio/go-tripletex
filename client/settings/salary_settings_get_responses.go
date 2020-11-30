// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// SalarySettingsGetReader is a Reader for the SalarySettingsGet structure.
type SalarySettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalarySettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalarySettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSalarySettingsGetOK creates a SalarySettingsGetOK with default headers values
func NewSalarySettingsGetOK() *SalarySettingsGetOK {
	return &SalarySettingsGetOK{}
}

/*SalarySettingsGetOK handles this case with default header values.

successful operation
*/
type SalarySettingsGetOK struct {
	Payload *models.ResponseWrapperSalarySettings
}

func (o *SalarySettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /salary/settings][%d] salarySettingsGetOK  %+v", 200, o.Payload)
}

func (o *SalarySettingsGetOK) GetPayload() *models.ResponseWrapperSalarySettings {
	return o.Payload
}

func (o *SalarySettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSalarySettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
