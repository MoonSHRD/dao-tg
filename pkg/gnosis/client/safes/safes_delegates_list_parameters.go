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

// NewSafesDelegatesListParams creates a new SafesDelegatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSafesDelegatesListParams() *SafesDelegatesListParams {
	return &SafesDelegatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSafesDelegatesListParamsWithTimeout creates a new SafesDelegatesListParams object
// with the ability to set a timeout on a request.
func NewSafesDelegatesListParamsWithTimeout(timeout time.Duration) *SafesDelegatesListParams {
	return &SafesDelegatesListParams{
		timeout: timeout,
	}
}

// NewSafesDelegatesListParamsWithContext creates a new SafesDelegatesListParams object
// with the ability to set a context for a request.
func NewSafesDelegatesListParamsWithContext(ctx context.Context) *SafesDelegatesListParams {
	return &SafesDelegatesListParams{
		Context: ctx,
	}
}

// NewSafesDelegatesListParamsWithHTTPClient creates a new SafesDelegatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSafesDelegatesListParamsWithHTTPClient(client *http.Client) *SafesDelegatesListParams {
	return &SafesDelegatesListParams{
		HTTPClient: client,
	}
}

/* SafesDelegatesListParams contains all the parameters to send to the API endpoint
   for the safes delegates list operation.

   Typically these are written to a http.Request.
*/
type SafesDelegatesListParams struct {

	// Address.
	Address string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the safes delegates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SafesDelegatesListParams) WithDefaults() *SafesDelegatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the safes delegates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SafesDelegatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the safes delegates list params
func (o *SafesDelegatesListParams) WithTimeout(timeout time.Duration) *SafesDelegatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the safes delegates list params
func (o *SafesDelegatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the safes delegates list params
func (o *SafesDelegatesListParams) WithContext(ctx context.Context) *SafesDelegatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the safes delegates list params
func (o *SafesDelegatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the safes delegates list params
func (o *SafesDelegatesListParams) WithHTTPClient(client *http.Client) *SafesDelegatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the safes delegates list params
func (o *SafesDelegatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the safes delegates list params
func (o *SafesDelegatesListParams) WithAddress(address string) *SafesDelegatesListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the safes delegates list params
func (o *SafesDelegatesListParams) SetAddress(address string) {
	o.Address = address
}

// WithLimit adds the limit to the safes delegates list params
func (o *SafesDelegatesListParams) WithLimit(limit *int64) *SafesDelegatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the safes delegates list params
func (o *SafesDelegatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the safes delegates list params
func (o *SafesDelegatesListParams) WithOffset(offset *int64) *SafesDelegatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the safes delegates list params
func (o *SafesDelegatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *SafesDelegatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
