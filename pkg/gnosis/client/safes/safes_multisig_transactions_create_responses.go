// Code generated by go-swagger; DO NOT EDIT.

package safes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SafesMultisigTransactionsCreateReader is a Reader for the SafesMultisigTransactionsCreate structure.
type SafesMultisigTransactionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SafesMultisigTransactionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSafesMultisigTransactionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSafesMultisigTransactionsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSafesMultisigTransactionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSafesMultisigTransactionsCreateCreated creates a SafesMultisigTransactionsCreateCreated with default headers values
func NewSafesMultisigTransactionsCreateCreated() *SafesMultisigTransactionsCreateCreated {
	return &SafesMultisigTransactionsCreateCreated{}
}

/* SafesMultisigTransactionsCreateCreated describes a response with status code 201, with default header values.

Created or signature updated
*/
type SafesMultisigTransactionsCreateCreated struct {
}

func (o *SafesMultisigTransactionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /safes/{address}/multisig-transactions/][%d] safesMultisigTransactionsCreateCreated ", 201)
}

func (o *SafesMultisigTransactionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSafesMultisigTransactionsCreateBadRequest creates a SafesMultisigTransactionsCreateBadRequest with default headers values
func NewSafesMultisigTransactionsCreateBadRequest() *SafesMultisigTransactionsCreateBadRequest {
	return &SafesMultisigTransactionsCreateBadRequest{}
}

/* SafesMultisigTransactionsCreateBadRequest describes a response with status code 400, with default header values.

Invalid data
*/
type SafesMultisigTransactionsCreateBadRequest struct {
}

func (o *SafesMultisigTransactionsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /safes/{address}/multisig-transactions/][%d] safesMultisigTransactionsCreateBadRequest ", 400)
}

func (o *SafesMultisigTransactionsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSafesMultisigTransactionsCreateUnprocessableEntity creates a SafesMultisigTransactionsCreateUnprocessableEntity with default headers values
func NewSafesMultisigTransactionsCreateUnprocessableEntity() *SafesMultisigTransactionsCreateUnprocessableEntity {
	return &SafesMultisigTransactionsCreateUnprocessableEntity{}
}

/* SafesMultisigTransactionsCreateUnprocessableEntity describes a response with status code 422, with default header values.

Invalid ethereum address/User is not an owner/Invalid safeTxHash/Invalid signature/Nonce already executed/Sender is not an owner
*/
type SafesMultisigTransactionsCreateUnprocessableEntity struct {
}

func (o *SafesMultisigTransactionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /safes/{address}/multisig-transactions/][%d] safesMultisigTransactionsCreateUnprocessableEntity ", 422)
}

func (o *SafesMultisigTransactionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
