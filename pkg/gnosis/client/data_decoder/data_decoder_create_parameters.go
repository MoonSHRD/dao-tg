// Code generated by go-swagger; DO NOT EDIT.

package data_decoder

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

	"github.com/MoonSHRD/dao-tg/pkg/gnosis/models"
)

// NewDataDecoderCreateParams creates a new DataDecoderCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDataDecoderCreateParams() *DataDecoderCreateParams {
	return &DataDecoderCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDataDecoderCreateParamsWithTimeout creates a new DataDecoderCreateParams object
// with the ability to set a timeout on a request.
func NewDataDecoderCreateParamsWithTimeout(timeout time.Duration) *DataDecoderCreateParams {
	return &DataDecoderCreateParams{
		timeout: timeout,
	}
}

// NewDataDecoderCreateParamsWithContext creates a new DataDecoderCreateParams object
// with the ability to set a context for a request.
func NewDataDecoderCreateParamsWithContext(ctx context.Context) *DataDecoderCreateParams {
	return &DataDecoderCreateParams{
		Context: ctx,
	}
}

// NewDataDecoderCreateParamsWithHTTPClient creates a new DataDecoderCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDataDecoderCreateParamsWithHTTPClient(client *http.Client) *DataDecoderCreateParams {
	return &DataDecoderCreateParams{
		HTTPClient: client,
	}
}

/* DataDecoderCreateParams contains all the parameters to send to the API endpoint
   for the data decoder create operation.

   Typically these are written to a http.Request.
*/
type DataDecoderCreateParams struct {

	// Data.
	Data *models.DataDecoder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the data decoder create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DataDecoderCreateParams) WithDefaults() *DataDecoderCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the data decoder create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DataDecoderCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the data decoder create params
func (o *DataDecoderCreateParams) WithTimeout(timeout time.Duration) *DataDecoderCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data decoder create params
func (o *DataDecoderCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data decoder create params
func (o *DataDecoderCreateParams) WithContext(ctx context.Context) *DataDecoderCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data decoder create params
func (o *DataDecoderCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data decoder create params
func (o *DataDecoderCreateParams) WithHTTPClient(client *http.Client) *DataDecoderCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data decoder create params
func (o *DataDecoderCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the data decoder create params
func (o *DataDecoderCreateParams) WithData(data *models.DataDecoder) *DataDecoderCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the data decoder create params
func (o *DataDecoderCreateParams) SetData(data *models.DataDecoder) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DataDecoderCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
