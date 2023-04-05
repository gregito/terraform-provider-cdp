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

// NewDisableS3GuardParams creates a new DisableS3GuardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableS3GuardParams() *DisableS3GuardParams {
	return &DisableS3GuardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableS3GuardParamsWithTimeout creates a new DisableS3GuardParams object
// with the ability to set a timeout on a request.
func NewDisableS3GuardParamsWithTimeout(timeout time.Duration) *DisableS3GuardParams {
	return &DisableS3GuardParams{
		timeout: timeout,
	}
}

// NewDisableS3GuardParamsWithContext creates a new DisableS3GuardParams object
// with the ability to set a context for a request.
func NewDisableS3GuardParamsWithContext(ctx context.Context) *DisableS3GuardParams {
	return &DisableS3GuardParams{
		Context: ctx,
	}
}

// NewDisableS3GuardParamsWithHTTPClient creates a new DisableS3GuardParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableS3GuardParamsWithHTTPClient(client *http.Client) *DisableS3GuardParams {
	return &DisableS3GuardParams{
		HTTPClient: client,
	}
}

/*
DisableS3GuardParams contains all the parameters to send to the API endpoint

	for the disable s3 guard operation.

	Typically these are written to a http.Request.
*/
type DisableS3GuardParams struct {

	// Input.
	Input *models.DisableS3GuardRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable s3 guard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableS3GuardParams) WithDefaults() *DisableS3GuardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable s3 guard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableS3GuardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable s3 guard params
func (o *DisableS3GuardParams) WithTimeout(timeout time.Duration) *DisableS3GuardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable s3 guard params
func (o *DisableS3GuardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable s3 guard params
func (o *DisableS3GuardParams) WithContext(ctx context.Context) *DisableS3GuardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable s3 guard params
func (o *DisableS3GuardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable s3 guard params
func (o *DisableS3GuardParams) WithHTTPClient(client *http.Client) *DisableS3GuardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable s3 guard params
func (o *DisableS3GuardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the disable s3 guard params
func (o *DisableS3GuardParams) WithInput(input *models.DisableS3GuardRequest) *DisableS3GuardParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the disable s3 guard params
func (o *DisableS3GuardParams) SetInput(input *models.DisableS3GuardRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DisableS3GuardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
