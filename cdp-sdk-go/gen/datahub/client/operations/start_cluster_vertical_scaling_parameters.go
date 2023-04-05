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

// NewStartClusterVerticalScalingParams creates a new StartClusterVerticalScalingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartClusterVerticalScalingParams() *StartClusterVerticalScalingParams {
	return &StartClusterVerticalScalingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartClusterVerticalScalingParamsWithTimeout creates a new StartClusterVerticalScalingParams object
// with the ability to set a timeout on a request.
func NewStartClusterVerticalScalingParamsWithTimeout(timeout time.Duration) *StartClusterVerticalScalingParams {
	return &StartClusterVerticalScalingParams{
		timeout: timeout,
	}
}

// NewStartClusterVerticalScalingParamsWithContext creates a new StartClusterVerticalScalingParams object
// with the ability to set a context for a request.
func NewStartClusterVerticalScalingParamsWithContext(ctx context.Context) *StartClusterVerticalScalingParams {
	return &StartClusterVerticalScalingParams{
		Context: ctx,
	}
}

// NewStartClusterVerticalScalingParamsWithHTTPClient creates a new StartClusterVerticalScalingParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartClusterVerticalScalingParamsWithHTTPClient(client *http.Client) *StartClusterVerticalScalingParams {
	return &StartClusterVerticalScalingParams{
		HTTPClient: client,
	}
}

/*
StartClusterVerticalScalingParams contains all the parameters to send to the API endpoint

	for the start cluster vertical scaling operation.

	Typically these are written to a http.Request.
*/
type StartClusterVerticalScalingParams struct {

	// Input.
	Input *models.StartClusterVerticalScalingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start cluster vertical scaling params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartClusterVerticalScalingParams) WithDefaults() *StartClusterVerticalScalingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start cluster vertical scaling params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartClusterVerticalScalingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) WithTimeout(timeout time.Duration) *StartClusterVerticalScalingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) WithContext(ctx context.Context) *StartClusterVerticalScalingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) WithHTTPClient(client *http.Client) *StartClusterVerticalScalingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) WithInput(input *models.StartClusterVerticalScalingRequest) *StartClusterVerticalScalingParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the start cluster vertical scaling params
func (o *StartClusterVerticalScalingParams) SetInput(input *models.StartClusterVerticalScalingRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *StartClusterVerticalScalingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
