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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewDisableClouderaSSOLoginParams creates a new DisableClouderaSSOLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableClouderaSSOLoginParams() *DisableClouderaSSOLoginParams {
	return &DisableClouderaSSOLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableClouderaSSOLoginParamsWithTimeout creates a new DisableClouderaSSOLoginParams object
// with the ability to set a timeout on a request.
func NewDisableClouderaSSOLoginParamsWithTimeout(timeout time.Duration) *DisableClouderaSSOLoginParams {
	return &DisableClouderaSSOLoginParams{
		timeout: timeout,
	}
}

// NewDisableClouderaSSOLoginParamsWithContext creates a new DisableClouderaSSOLoginParams object
// with the ability to set a context for a request.
func NewDisableClouderaSSOLoginParamsWithContext(ctx context.Context) *DisableClouderaSSOLoginParams {
	return &DisableClouderaSSOLoginParams{
		Context: ctx,
	}
}

// NewDisableClouderaSSOLoginParamsWithHTTPClient creates a new DisableClouderaSSOLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableClouderaSSOLoginParamsWithHTTPClient(client *http.Client) *DisableClouderaSSOLoginParams {
	return &DisableClouderaSSOLoginParams{
		HTTPClient: client,
	}
}

/*
DisableClouderaSSOLoginParams contains all the parameters to send to the API endpoint

	for the disable cloudera s s o login operation.

	Typically these are written to a http.Request.
*/
type DisableClouderaSSOLoginParams struct {

	// Input.
	Input models.DisableClouderaSSOLoginRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable cloudera s s o login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableClouderaSSOLoginParams) WithDefaults() *DisableClouderaSSOLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable cloudera s s o login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableClouderaSSOLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) WithTimeout(timeout time.Duration) *DisableClouderaSSOLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) WithContext(ctx context.Context) *DisableClouderaSSOLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) WithHTTPClient(client *http.Client) *DisableClouderaSSOLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) WithInput(input models.DisableClouderaSSOLoginRequest) *DisableClouderaSSOLoginParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the disable cloudera s s o login params
func (o *DisableClouderaSSOLoginParams) SetInput(input models.DisableClouderaSSOLoginRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DisableClouderaSSOLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
