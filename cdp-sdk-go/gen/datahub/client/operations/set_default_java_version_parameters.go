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

// NewSetDefaultJavaVersionParams creates a new SetDefaultJavaVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDefaultJavaVersionParams() *SetDefaultJavaVersionParams {
	return &SetDefaultJavaVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDefaultJavaVersionParamsWithTimeout creates a new SetDefaultJavaVersionParams object
// with the ability to set a timeout on a request.
func NewSetDefaultJavaVersionParamsWithTimeout(timeout time.Duration) *SetDefaultJavaVersionParams {
	return &SetDefaultJavaVersionParams{
		timeout: timeout,
	}
}

// NewSetDefaultJavaVersionParamsWithContext creates a new SetDefaultJavaVersionParams object
// with the ability to set a context for a request.
func NewSetDefaultJavaVersionParamsWithContext(ctx context.Context) *SetDefaultJavaVersionParams {
	return &SetDefaultJavaVersionParams{
		Context: ctx,
	}
}

// NewSetDefaultJavaVersionParamsWithHTTPClient creates a new SetDefaultJavaVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDefaultJavaVersionParamsWithHTTPClient(client *http.Client) *SetDefaultJavaVersionParams {
	return &SetDefaultJavaVersionParams{
		HTTPClient: client,
	}
}

/*
SetDefaultJavaVersionParams contains all the parameters to send to the API endpoint

	for the set default java version operation.

	Typically these are written to a http.Request.
*/
type SetDefaultJavaVersionParams struct {

	// Input.
	Input *models.SetDefaultJavaVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set default java version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDefaultJavaVersionParams) WithDefaults() *SetDefaultJavaVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set default java version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDefaultJavaVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set default java version params
func (o *SetDefaultJavaVersionParams) WithTimeout(timeout time.Duration) *SetDefaultJavaVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set default java version params
func (o *SetDefaultJavaVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set default java version params
func (o *SetDefaultJavaVersionParams) WithContext(ctx context.Context) *SetDefaultJavaVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set default java version params
func (o *SetDefaultJavaVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set default java version params
func (o *SetDefaultJavaVersionParams) WithHTTPClient(client *http.Client) *SetDefaultJavaVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set default java version params
func (o *SetDefaultJavaVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the set default java version params
func (o *SetDefaultJavaVersionParams) WithInput(input *models.SetDefaultJavaVersionRequest) *SetDefaultJavaVersionParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the set default java version params
func (o *SetDefaultJavaVersionParams) SetInput(input *models.SetDefaultJavaVersionRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SetDefaultJavaVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
