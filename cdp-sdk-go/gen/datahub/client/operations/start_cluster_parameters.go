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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// NewStartClusterParams creates a new StartClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartClusterParams() *StartClusterParams {
	return &StartClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartClusterParamsWithTimeout creates a new StartClusterParams object
// with the ability to set a timeout on a request.
func NewStartClusterParamsWithTimeout(timeout time.Duration) *StartClusterParams {
	return &StartClusterParams{
		timeout: timeout,
	}
}

// NewStartClusterParamsWithContext creates a new StartClusterParams object
// with the ability to set a context for a request.
func NewStartClusterParamsWithContext(ctx context.Context) *StartClusterParams {
	return &StartClusterParams{
		Context: ctx,
	}
}

// NewStartClusterParamsWithHTTPClient creates a new StartClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartClusterParamsWithHTTPClient(client *http.Client) *StartClusterParams {
	return &StartClusterParams{
		HTTPClient: client,
	}
}

/*
StartClusterParams contains all the parameters to send to the API endpoint

	for the start cluster operation.

	Typically these are written to a http.Request.
*/
type StartClusterParams struct {

	// Input.
	Input *models.StartClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartClusterParams) WithDefaults() *StartClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start cluster params
func (o *StartClusterParams) WithTimeout(timeout time.Duration) *StartClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start cluster params
func (o *StartClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start cluster params
func (o *StartClusterParams) WithContext(ctx context.Context) *StartClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start cluster params
func (o *StartClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start cluster params
func (o *StartClusterParams) WithHTTPClient(client *http.Client) *StartClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start cluster params
func (o *StartClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the start cluster params
func (o *StartClusterParams) WithInput(input *models.StartClusterRequest) *StartClusterParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the start cluster params
func (o *StartClusterParams) SetInput(input *models.StartClusterRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *StartClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
