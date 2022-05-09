// Code generated by go-swagger; DO NOT EDIT.

package safes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MoonSHRD/dao-tg/pkg/gnosis/models"
)

// SafesIncomingTransfersListReader is a Reader for the SafesIncomingTransfersList structure.
type SafesIncomingTransfersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SafesIncomingTransfersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSafesIncomingTransfersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewSafesIncomingTransfersListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSafesIncomingTransfersListOK creates a SafesIncomingTransfersListOK with default headers values
func NewSafesIncomingTransfersListOK() *SafesIncomingTransfersListOK {
	return &SafesIncomingTransfersListOK{}
}

/* SafesIncomingTransfersListOK describes a response with status code 200, with default header values.

SafesIncomingTransfersListOK safes incoming transfers list o k
*/
type SafesIncomingTransfersListOK struct {
	Payload []*models.TransferWithTokenInfoResponse
}

func (o *SafesIncomingTransfersListOK) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/incoming-transfers/][%d] safesIncomingTransfersListOK  %+v", 200, o.Payload)
}
func (o *SafesIncomingTransfersListOK) GetPayload() []*models.TransferWithTokenInfoResponse {
	return o.Payload
}

func (o *SafesIncomingTransfersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSafesIncomingTransfersListUnprocessableEntity creates a SafesIncomingTransfersListUnprocessableEntity with default headers values
func NewSafesIncomingTransfersListUnprocessableEntity() *SafesIncomingTransfersListUnprocessableEntity {
	return &SafesIncomingTransfersListUnprocessableEntity{}
}

/* SafesIncomingTransfersListUnprocessableEntity describes a response with status code 422, with default header values.

Safe address checksum not valid
*/
type SafesIncomingTransfersListUnprocessableEntity struct {
}

func (o *SafesIncomingTransfersListUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/incoming-transfers/][%d] safesIncomingTransfersListUnprocessableEntity ", 422)
}

func (o *SafesIncomingTransfersListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
