// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// OrderListPostListReader is a Reader for the OrderListPostList structure.
type OrderListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrderListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrderListPostListCreated creates a OrderListPostListCreated with default headers values
func NewOrderListPostListCreated() *OrderListPostListCreated {
	return &OrderListPostListCreated{}
}

/*OrderListPostListCreated handles this case with default header values.

successfully created
*/
type OrderListPostListCreated struct {
	Payload *models.ListResponseOrder
}

func (o *OrderListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /order/list][%d] orderListPostListCreated  %+v", 201, o.Payload)
}

func (o *OrderListPostListCreated) GetPayload() *models.ListResponseOrder {
	return o.Payload
}

func (o *OrderListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
