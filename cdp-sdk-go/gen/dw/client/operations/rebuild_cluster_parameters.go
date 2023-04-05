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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewRebuildClusterParams creates a new RebuildClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRebuildClusterParams() *RebuildClusterParams {
	return &RebuildClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRebuildClusterParamsWithTimeout creates a new RebuildClusterParams object
// with the ability to set a timeout on a request.
func NewRebuildClusterParamsWithTimeout(timeout time.Duration) *RebuildClusterParams {
	return &RebuildClusterParams{
		timeout: timeout,
	}
}

// NewRebuildClusterParamsWithContext creates a new RebuildClusterParams object
// with the ability to set a context for a request.
func NewRebuildClusterParamsWithContext(ctx context.Context) *RebuildClusterParams {
	return &RebuildClusterParams{
		Context: ctx,
	}
}

// NewRebuildClusterParamsWithHTTPClient creates a new RebuildClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRebuildClusterParamsWithHTTPClient(client *http.Client) *RebuildClusterParams {
	return &RebuildClusterParams{
		HTTPClient: client,
	}
}

/*
RebuildClusterParams contains all the parameters to send to the API endpoint

	for the rebuild cluster operation.

	Typically these are written to a http.Request.
*/
type RebuildClusterParams struct {

	// Input.
	Input *models.RebuildClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rebuild cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebuildClusterParams) WithDefaults() *RebuildClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rebuild cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebuildClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rebuild cluster params
func (o *RebuildClusterParams) WithTimeout(timeout time.Duration) *RebuildClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebuild cluster params
func (o *RebuildClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebuild cluster params
func (o *RebuildClusterParams) WithContext(ctx context.Context) *RebuildClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebuild cluster params
func (o *RebuildClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebuild cluster params
func (o *RebuildClusterParams) WithHTTPClient(client *http.Client) *RebuildClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebuild cluster params
func (o *RebuildClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the rebuild cluster params
func (o *RebuildClusterParams) WithInput(input *models.RebuildClusterRequest) *RebuildClusterParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the rebuild cluster params
func (o *RebuildClusterParams) SetInput(input *models.RebuildClusterRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RebuildClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
