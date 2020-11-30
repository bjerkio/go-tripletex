// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ActivityGetReader is a Reader for the ActivityGet structure.
type ActivityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActivityGetOK creates a ActivityGetOK with default headers values
func NewActivityGetOK() *ActivityGetOK {
	return &ActivityGetOK{}
}

/*ActivityGetOK handles this case with default header values.

successful operation
*/
type ActivityGetOK struct {
	Payload *models.ResponseWrapperActivity
}

func (o *ActivityGetOK) Error() string {
	return fmt.Sprintf("[GET /activity/{id}][%d] activityGetOK  %+v", 200, o.Payload)
}

func (o *ActivityGetOK) GetPayload() *models.ResponseWrapperActivity {
	return o.Payload
}

func (o *ActivityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
