// Code generated by go-swagger; DO NOT EDIT.

package project_specific_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectHourlyRatesProjectSpecificRatesGetReader is a Reader for the ProjectHourlyRatesProjectSpecificRatesGet structure.
type ProjectHourlyRatesProjectSpecificRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectHourlyRatesProjectSpecificRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectHourlyRatesProjectSpecificRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectHourlyRatesProjectSpecificRatesGetOK creates a ProjectHourlyRatesProjectSpecificRatesGetOK with default headers values
func NewProjectHourlyRatesProjectSpecificRatesGetOK() *ProjectHourlyRatesProjectSpecificRatesGetOK {
	return &ProjectHourlyRatesProjectSpecificRatesGetOK{}
}

/*ProjectHourlyRatesProjectSpecificRatesGetOK handles this case with default header values.

successful operation
*/
type ProjectHourlyRatesProjectSpecificRatesGetOK struct {
	Payload *models.ResponseWrapperProjectSpecificRate
}

func (o *ProjectHourlyRatesProjectSpecificRatesGetOK) Error() string {
	return fmt.Sprintf("[GET /project/hourlyRates/projectSpecificRates/{id}][%d] projectHourlyRatesProjectSpecificRatesGetOK  %+v", 200, o.Payload)
}

func (o *ProjectHourlyRatesProjectSpecificRatesGetOK) GetPayload() *models.ResponseWrapperProjectSpecificRate {
	return o.Payload
}

func (o *ProjectHourlyRatesProjectSpecificRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectSpecificRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
