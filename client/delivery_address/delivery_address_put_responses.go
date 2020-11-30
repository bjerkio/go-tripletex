// Code generated by go-swagger; DO NOT EDIT.

package delivery_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// DeliveryAddressPutReader is a Reader for the DeliveryAddressPut structure.
type DeliveryAddressPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeliveryAddressPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeliveryAddressPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeliveryAddressPutOK creates a DeliveryAddressPutOK with default headers values
func NewDeliveryAddressPutOK() *DeliveryAddressPutOK {
	return &DeliveryAddressPutOK{}
}

/*DeliveryAddressPutOK handles this case with default header values.

successful operation
*/
type DeliveryAddressPutOK struct {
	Payload *models.ResponseWrapperDeliveryAddress
}

func (o *DeliveryAddressPutOK) Error() string {
	return fmt.Sprintf("[PUT /deliveryAddress/{id}][%d] deliveryAddressPutOK  %+v", 200, o.Payload)
}

func (o *DeliveryAddressPutOK) GetPayload() *models.ResponseWrapperDeliveryAddress {
	return o.Payload
}

func (o *DeliveryAddressPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDeliveryAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
