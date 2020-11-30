// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ContactSearchReader is a Reader for the ContactSearch structure.
type ContactSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContactSearchOK creates a ContactSearchOK with default headers values
func NewContactSearchOK() *ContactSearchOK {
	return &ContactSearchOK{}
}

/*ContactSearchOK handles this case with default header values.

successful operation
*/
type ContactSearchOK struct {
	Payload *models.ListResponseContact
}

func (o *ContactSearchOK) Error() string {
	return fmt.Sprintf("[GET /contact][%d] contactSearchOK  %+v", 200, o.Payload)
}

func (o *ContactSearchOK) GetPayload() *models.ListResponseContact {
	return o.Payload
}

func (o *ContactSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
