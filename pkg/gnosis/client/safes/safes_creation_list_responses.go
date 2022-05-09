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

// SafesCreationListReader is a Reader for the SafesCreationList structure.
type SafesCreationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SafesCreationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSafesCreationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSafesCreationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSafesCreationListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSafesCreationListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSafesCreationListOK creates a SafesCreationListOK with default headers values
func NewSafesCreationListOK() *SafesCreationListOK {
	return &SafesCreationListOK{}
}

/* SafesCreationListOK describes a response with status code 200, with default header values.

SafesCreationListOK safes creation list o k
*/
type SafesCreationListOK struct {
	Payload *models.SafeCreationInfoResponse
}

func (o *SafesCreationListOK) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/creation/][%d] safesCreationListOK  %+v", 200, o.Payload)
}
func (o *SafesCreationListOK) GetPayload() *models.SafeCreationInfoResponse {
	return o.Payload
}

func (o *SafesCreationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SafeCreationInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSafesCreationListNotFound creates a SafesCreationListNotFound with default headers values
func NewSafesCreationListNotFound() *SafesCreationListNotFound {
	return &SafesCreationListNotFound{}
}

/* SafesCreationListNotFound describes a response with status code 404, with default header values.

Safe creation not found
*/
type SafesCreationListNotFound struct {
}

func (o *SafesCreationListNotFound) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/creation/][%d] safesCreationListNotFound ", 404)
}

func (o *SafesCreationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSafesCreationListUnprocessableEntity creates a SafesCreationListUnprocessableEntity with default headers values
func NewSafesCreationListUnprocessableEntity() *SafesCreationListUnprocessableEntity {
	return &SafesCreationListUnprocessableEntity{}
}

/* SafesCreationListUnprocessableEntity describes a response with status code 422, with default header values.

Owner address checksum not valid
*/
type SafesCreationListUnprocessableEntity struct {
}

func (o *SafesCreationListUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/creation/][%d] safesCreationListUnprocessableEntity ", 422)
}

func (o *SafesCreationListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSafesCreationListServiceUnavailable creates a SafesCreationListServiceUnavailable with default headers values
func NewSafesCreationListServiceUnavailable() *SafesCreationListServiceUnavailable {
	return &SafesCreationListServiceUnavailable{}
}

/* SafesCreationListServiceUnavailable describes a response with status code 503, with default header values.

Problem connecting to Ethereum network
*/
type SafesCreationListServiceUnavailable struct {
}

func (o *SafesCreationListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /safes/{address}/creation/][%d] safesCreationListServiceUnavailable ", 503)
}

func (o *SafesCreationListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
