// Code generated by go-swagger; DO NOT EDIT.

package transport_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TransportTypeGetReader is a Reader for the TransportTypeGet structure.
type TransportTypeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransportTypeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransportTypeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransportTypeGetOK creates a TransportTypeGetOK with default headers values
func NewTransportTypeGetOK() *TransportTypeGetOK {
	return &TransportTypeGetOK{}
}

/*TransportTypeGetOK handles this case with default header values.

successful operation
*/
type TransportTypeGetOK struct {
	Payload *models.ResponseWrapperTransportType
}

func (o *TransportTypeGetOK) Error() string {
	return fmt.Sprintf("[GET /transportType/{id}][%d] transportTypeGetOK  %+v", 200, o.Payload)
}

func (o *TransportTypeGetOK) GetPayload() *models.ResponseWrapperTransportType {
	return o.Payload
}

func (o *TransportTypeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTransportType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
