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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewDeleteDatalakeParams creates a new DeleteDatalakeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDatalakeParams() *DeleteDatalakeParams {
	return &DeleteDatalakeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatalakeParamsWithTimeout creates a new DeleteDatalakeParams object
// with the ability to set a timeout on a request.
func NewDeleteDatalakeParamsWithTimeout(timeout time.Duration) *DeleteDatalakeParams {
	return &DeleteDatalakeParams{
		timeout: timeout,
	}
}

// NewDeleteDatalakeParamsWithContext creates a new DeleteDatalakeParams object
// with the ability to set a context for a request.
func NewDeleteDatalakeParamsWithContext(ctx context.Context) *DeleteDatalakeParams {
	return &DeleteDatalakeParams{
		Context: ctx,
	}
}

// NewDeleteDatalakeParamsWithHTTPClient creates a new DeleteDatalakeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDatalakeParamsWithHTTPClient(client *http.Client) *DeleteDatalakeParams {
	return &DeleteDatalakeParams{
		HTTPClient: client,
	}
}

/*
DeleteDatalakeParams contains all the parameters to send to the API endpoint

	for the delete datalake operation.

	Typically these are written to a http.Request.
*/
type DeleteDatalakeParams struct {

	// Input.
	Input *models.DeleteDatalakeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatalakeParams) WithDefaults() *DeleteDatalakeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatalakeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete datalake params
func (o *DeleteDatalakeParams) WithTimeout(timeout time.Duration) *DeleteDatalakeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete datalake params
func (o *DeleteDatalakeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete datalake params
func (o *DeleteDatalakeParams) WithContext(ctx context.Context) *DeleteDatalakeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete datalake params
func (o *DeleteDatalakeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete datalake params
func (o *DeleteDatalakeParams) WithHTTPClient(client *http.Client) *DeleteDatalakeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete datalake params
func (o *DeleteDatalakeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete datalake params
func (o *DeleteDatalakeParams) WithInput(input *models.DeleteDatalakeRequest) *DeleteDatalakeParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete datalake params
func (o *DeleteDatalakeParams) SetInput(input *models.DeleteDatalakeRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatalakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
