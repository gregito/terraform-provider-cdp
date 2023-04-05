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

// NewGetClusterServiceStatusParams creates a new GetClusterServiceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterServiceStatusParams() *GetClusterServiceStatusParams {
	return &GetClusterServiceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterServiceStatusParamsWithTimeout creates a new GetClusterServiceStatusParams object
// with the ability to set a timeout on a request.
func NewGetClusterServiceStatusParamsWithTimeout(timeout time.Duration) *GetClusterServiceStatusParams {
	return &GetClusterServiceStatusParams{
		timeout: timeout,
	}
}

// NewGetClusterServiceStatusParamsWithContext creates a new GetClusterServiceStatusParams object
// with the ability to set a context for a request.
func NewGetClusterServiceStatusParamsWithContext(ctx context.Context) *GetClusterServiceStatusParams {
	return &GetClusterServiceStatusParams{
		Context: ctx,
	}
}

// NewGetClusterServiceStatusParamsWithHTTPClient creates a new GetClusterServiceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterServiceStatusParamsWithHTTPClient(client *http.Client) *GetClusterServiceStatusParams {
	return &GetClusterServiceStatusParams{
		HTTPClient: client,
	}
}

/*
GetClusterServiceStatusParams contains all the parameters to send to the API endpoint

	for the get cluster service status operation.

	Typically these are written to a http.Request.
*/
type GetClusterServiceStatusParams struct {

	// Input.
	Input *models.GetClusterServiceStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster service status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterServiceStatusParams) WithDefaults() *GetClusterServiceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster service status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterServiceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster service status params
func (o *GetClusterServiceStatusParams) WithTimeout(timeout time.Duration) *GetClusterServiceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster service status params
func (o *GetClusterServiceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster service status params
func (o *GetClusterServiceStatusParams) WithContext(ctx context.Context) *GetClusterServiceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster service status params
func (o *GetClusterServiceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster service status params
func (o *GetClusterServiceStatusParams) WithHTTPClient(client *http.Client) *GetClusterServiceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster service status params
func (o *GetClusterServiceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get cluster service status params
func (o *GetClusterServiceStatusParams) WithInput(input *models.GetClusterServiceStatusRequest) *GetClusterServiceStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get cluster service status params
func (o *GetClusterServiceStatusParams) SetInput(input *models.GetClusterServiceStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterServiceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
