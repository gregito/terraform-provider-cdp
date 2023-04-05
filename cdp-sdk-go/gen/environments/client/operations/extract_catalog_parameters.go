// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewExtractCatalogParams creates a new ExtractCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtractCatalogParams() *ExtractCatalogParams {
	return &ExtractCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtractCatalogParamsWithTimeout creates a new ExtractCatalogParams object
// with the ability to set a timeout on a request.
func NewExtractCatalogParamsWithTimeout(timeout time.Duration) *ExtractCatalogParams {
	return &ExtractCatalogParams{
		timeout: timeout,
	}
}

// NewExtractCatalogParamsWithContext creates a new ExtractCatalogParams object
// with the ability to set a context for a request.
func NewExtractCatalogParamsWithContext(ctx context.Context) *ExtractCatalogParams {
	return &ExtractCatalogParams{
		Context: ctx,
	}
}

// NewExtractCatalogParamsWithHTTPClient creates a new ExtractCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtractCatalogParamsWithHTTPClient(client *http.Client) *ExtractCatalogParams {
	return &ExtractCatalogParams{
		HTTPClient: client,
	}
}

/*
ExtractCatalogParams contains all the parameters to send to the API endpoint

	for the extract catalog operation.

	Typically these are written to a http.Request.
*/
type ExtractCatalogParams struct {

	// Input.
	Input *models.ExtractCatalogRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extract catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtractCatalogParams) WithDefaults() *ExtractCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extract catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtractCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extract catalog params
func (o *ExtractCatalogParams) WithTimeout(timeout time.Duration) *ExtractCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extract catalog params
func (o *ExtractCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extract catalog params
func (o *ExtractCatalogParams) WithContext(ctx context.Context) *ExtractCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extract catalog params
func (o *ExtractCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extract catalog params
func (o *ExtractCatalogParams) WithHTTPClient(client *http.Client) *ExtractCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extract catalog params
func (o *ExtractCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the extract catalog params
func (o *ExtractCatalogParams) WithInput(input *models.ExtractCatalogRequest) *ExtractCatalogParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the extract catalog params
func (o *ExtractCatalogParams) SetInput(input *models.ExtractCatalogRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ExtractCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
