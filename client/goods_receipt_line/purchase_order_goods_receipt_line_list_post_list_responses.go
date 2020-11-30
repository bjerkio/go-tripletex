// Code generated by go-swagger; DO NOT EDIT.

package goods_receipt_line

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// PurchaseOrderGoodsReceiptLineListPostListReader is a Reader for the PurchaseOrderGoodsReceiptLineListPostList structure.
type PurchaseOrderGoodsReceiptLineListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderGoodsReceiptLineListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPurchaseOrderGoodsReceiptLineListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurchaseOrderGoodsReceiptLineListPostListCreated creates a PurchaseOrderGoodsReceiptLineListPostListCreated with default headers values
func NewPurchaseOrderGoodsReceiptLineListPostListCreated() *PurchaseOrderGoodsReceiptLineListPostListCreated {
	return &PurchaseOrderGoodsReceiptLineListPostListCreated{}
}

/*PurchaseOrderGoodsReceiptLineListPostListCreated handles this case with default header values.

successfully created
*/
type PurchaseOrderGoodsReceiptLineListPostListCreated struct {
	Payload *models.ListResponseGoodsReceiptLine
}

func (o *PurchaseOrderGoodsReceiptLineListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /purchaseOrder/goodsReceiptLine/list][%d] purchaseOrderGoodsReceiptLineListPostListCreated  %+v", 201, o.Payload)
}

func (o *PurchaseOrderGoodsReceiptLineListPostListCreated) GetPayload() *models.ListResponseGoodsReceiptLine {
	return o.Payload
}

func (o *PurchaseOrderGoodsReceiptLineListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseGoodsReceiptLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
