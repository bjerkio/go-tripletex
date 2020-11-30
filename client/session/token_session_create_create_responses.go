// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TokenSessionCreateCreateReader is a Reader for the TokenSessionCreateCreate structure.
type TokenSessionCreateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenSessionCreateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenSessionCreateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTokenSessionCreateCreateOK creates a TokenSessionCreateCreateOK with default headers values
func NewTokenSessionCreateCreateOK() *TokenSessionCreateCreateOK {
	return &TokenSessionCreateCreateOK{}
}

/*TokenSessionCreateCreateOK handles this case with default header values.

successful operation
*/
type TokenSessionCreateCreateOK struct {
	Payload *models.ResponseWrapperSessionToken
}

func (o *TokenSessionCreateCreateOK) Error() string {
	return fmt.Sprintf("[PUT /token/session/:create][%d] tokenSessionCreateCreateOK  %+v", 200, o.Payload)
}

func (o *TokenSessionCreateCreateOK) GetPayload() *models.ResponseWrapperSessionToken {
	return o.Payload
}

func (o *TokenSessionCreateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSessionToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
