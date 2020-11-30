// Code generated by go-swagger; DO NOT EDIT.

package department

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// DepartmentListPutListReader is a Reader for the DepartmentListPutList structure.
type DepartmentListPutListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DepartmentListPutListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDepartmentListPutListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDepartmentListPutListOK creates a DepartmentListPutListOK with default headers values
func NewDepartmentListPutListOK() *DepartmentListPutListOK {
	return &DepartmentListPutListOK{}
}

/*DepartmentListPutListOK handles this case with default header values.

successful operation
*/
type DepartmentListPutListOK struct {
	Payload *models.ListResponseDepartment
}

func (o *DepartmentListPutListOK) Error() string {
	return fmt.Sprintf("[PUT /department/list][%d] departmentListPutListOK  %+v", 200, o.Payload)
}

func (o *DepartmentListPutListOK) GetPayload() *models.ListResponseDepartment {
	return o.Payload
}

func (o *DepartmentListPutListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
