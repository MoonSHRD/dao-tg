// Code generated by go-swagger; DO NOT EDIT.

package safes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSafesAllTransactionsListParams creates a new SafesAllTransactionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSafesAllTransactionsListParams() *SafesAllTransactionsListParams {
	return &SafesAllTransactionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSafesAllTransactionsListParamsWithTimeout creates a new SafesAllTransactionsListParams object
// with the ability to set a timeout on a request.
func NewSafesAllTransactionsListParamsWithTimeout(timeout time.Duration) *SafesAllTransactionsListParams {
	return &SafesAllTransactionsListParams{
		timeout: timeout,
	}
}

// NewSafesAllTransactionsListParamsWithContext creates a new SafesAllTransactionsListParams object
// with the ability to set a context for a request.
func NewSafesAllTransactionsListParamsWithContext(ctx context.Context) *SafesAllTransactionsListParams {
	return &SafesAllTransactionsListParams{
		Context: ctx,
	}
}

// NewSafesAllTransactionsListParamsWithHTTPClient creates a new SafesAllTransactionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSafesAllTransactionsListParamsWithHTTPClient(client *http.Client) *SafesAllTransactionsListParams {
	return &SafesAllTransactionsListParams{
		HTTPClient: client,
	}
}

/* SafesAllTransactionsListParams contains all the parameters to send to the API endpoint
   for the safes all transactions list operation.

   Typically these are written to a http.Request.
*/
type SafesAllTransactionsListParams struct {

	// Address.
	Address string

	/* Executed.

	   If `True` only executed transactions are returned
	*/
	Executed *bool

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	/* Queued.

	   If `True` transactions with `nonce >= Safe current nonce` are also returned

	   Default: true
	*/
	Queued *bool

	/* Trusted.

	   If `True` just trusted transactions are shown (indexed, added by a delegate or with at least one confirmation)

	   Default: true
	*/
	Trusted *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the safes all transactions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SafesAllTransactionsListParams) WithDefaults() *SafesAllTransactionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the safes all transactions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SafesAllTransactionsListParams) SetDefaults() {
	var (
		executedDefault = bool(false)

		queuedDefault = bool(true)

		trustedDefault = bool(true)
	)

	val := SafesAllTransactionsListParams{
		Executed: &executedDefault,
		Queued:   &queuedDefault,
		Trusted:  &trustedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithTimeout(timeout time.Duration) *SafesAllTransactionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithContext(ctx context.Context) *SafesAllTransactionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithHTTPClient(client *http.Client) *SafesAllTransactionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithAddress(address string) *SafesAllTransactionsListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetAddress(address string) {
	o.Address = address
}

// WithExecuted adds the executed to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithExecuted(executed *bool) *SafesAllTransactionsListParams {
	o.SetExecuted(executed)
	return o
}

// SetExecuted adds the executed to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetExecuted(executed *bool) {
	o.Executed = executed
}

// WithLimit adds the limit to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithLimit(limit *int64) *SafesAllTransactionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithOffset(offset *int64) *SafesAllTransactionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithOrdering(ordering *string) *SafesAllTransactionsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithQueued adds the queued to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithQueued(queued *bool) *SafesAllTransactionsListParams {
	o.SetQueued(queued)
	return o
}

// SetQueued adds the queued to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetQueued(queued *bool) {
	o.Queued = queued
}

// WithTrusted adds the trusted to the safes all transactions list params
func (o *SafesAllTransactionsListParams) WithTrusted(trusted *bool) *SafesAllTransactionsListParams {
	o.SetTrusted(trusted)
	return o
}

// SetTrusted adds the trusted to the safes all transactions list params
func (o *SafesAllTransactionsListParams) SetTrusted(trusted *bool) {
	o.Trusted = trusted
}

// WriteToRequest writes these params to a swagger request
func (o *SafesAllTransactionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.Executed != nil {

		// query param executed
		var qrExecuted bool

		if o.Executed != nil {
			qrExecuted = *o.Executed
		}
		qExecuted := swag.FormatBool(qrExecuted)
		if qExecuted != "" {

			if err := r.SetQueryParam("executed", qExecuted); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Queued != nil {

		// query param queued
		var qrQueued bool

		if o.Queued != nil {
			qrQueued = *o.Queued
		}
		qQueued := swag.FormatBool(qrQueued)
		if qQueued != "" {

			if err := r.SetQueryParam("queued", qQueued); err != nil {
				return err
			}
		}
	}

	if o.Trusted != nil {

		// query param trusted
		var qrTrusted bool

		if o.Trusted != nil {
			qrTrusted = *o.Trusted
		}
		qTrusted := swag.FormatBool(qrTrusted)
		if qTrusted != "" {

			if err := r.SetQueryParam("trusted", qTrusted); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
