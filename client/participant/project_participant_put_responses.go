// Code generated by go-swagger; DO NOT EDIT.

package participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectParticipantPutReader is a Reader for the ProjectParticipantPut structure.
type ProjectParticipantPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectParticipantPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectParticipantPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectParticipantPutOK creates a ProjectParticipantPutOK with default headers values
func NewProjectParticipantPutOK() *ProjectParticipantPutOK {
	return &ProjectParticipantPutOK{}
}

/*ProjectParticipantPutOK handles this case with default header values.

successful operation
*/
type ProjectParticipantPutOK struct {
	Payload *models.ResponseWrapperProjectParticipant
}

func (o *ProjectParticipantPutOK) Error() string {
	return fmt.Sprintf("[PUT /project/participant/{id}][%d] projectParticipantPutOK  %+v", 200, o.Payload)
}

func (o *ProjectParticipantPutOK) GetPayload() *models.ResponseWrapperProjectParticipant {
	return o.Payload
}

func (o *ProjectParticipantPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectParticipant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
