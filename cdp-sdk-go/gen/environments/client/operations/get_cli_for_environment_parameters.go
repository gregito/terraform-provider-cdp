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

// NewGetCliForEnvironmentParams creates a new GetCliForEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCliForEnvironmentParams() *GetCliForEnvironmentParams {
	return &GetCliForEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCliForEnvironmentParamsWithTimeout creates a new GetCliForEnvironmentParams object
// with the ability to set a timeout on a request.
func NewGetCliForEnvironmentParamsWithTimeout(timeout time.Duration) *GetCliForEnvironmentParams {
	return &GetCliForEnvironmentParams{
		timeout: timeout,
	}
}

// NewGetCliForEnvironmentParamsWithContext creates a new GetCliForEnvironmentParams object
// with the ability to set a context for a request.
func NewGetCliForEnvironmentParamsWithContext(ctx context.Context) *GetCliForEnvironmentParams {
	return &GetCliForEnvironmentParams{
		Context: ctx,
	}
}

// NewGetCliForEnvironmentParamsWithHTTPClient creates a new GetCliForEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCliForEnvironmentParamsWithHTTPClient(client *http.Client) *GetCliForEnvironmentParams {
	return &GetCliForEnvironmentParams{
		HTTPClient: client,
	}
}

/*
GetCliForEnvironmentParams contains all the parameters to send to the API endpoint

	for the get cli for environment operation.

	Typically these are written to a http.Request.
*/
type GetCliForEnvironmentParams struct {

	// Input.
	Input *models.GetCliForEnvironmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cli for environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCliForEnvironmentParams) WithDefaults() *GetCliForEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cli for environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCliForEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cli for environment params
func (o *GetCliForEnvironmentParams) WithTimeout(timeout time.Duration) *GetCliForEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cli for environment params
func (o *GetCliForEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cli for environment params
func (o *GetCliForEnvironmentParams) WithContext(ctx context.Context) *GetCliForEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cli for environment params
func (o *GetCliForEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cli for environment params
func (o *GetCliForEnvironmentParams) WithHTTPClient(client *http.Client) *GetCliForEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cli for environment params
func (o *GetCliForEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get cli for environment params
func (o *GetCliForEnvironmentParams) WithInput(input *models.GetCliForEnvironmentRequest) *GetCliForEnvironmentParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get cli for environment params
func (o *GetCliForEnvironmentParams) SetInput(input *models.GetCliForEnvironmentRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetCliForEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
