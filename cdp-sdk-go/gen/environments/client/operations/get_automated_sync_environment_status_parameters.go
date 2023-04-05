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

// NewGetAutomatedSyncEnvironmentStatusParams creates a new GetAutomatedSyncEnvironmentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAutomatedSyncEnvironmentStatusParams() *GetAutomatedSyncEnvironmentStatusParams {
	return &GetAutomatedSyncEnvironmentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAutomatedSyncEnvironmentStatusParamsWithTimeout creates a new GetAutomatedSyncEnvironmentStatusParams object
// with the ability to set a timeout on a request.
func NewGetAutomatedSyncEnvironmentStatusParamsWithTimeout(timeout time.Duration) *GetAutomatedSyncEnvironmentStatusParams {
	return &GetAutomatedSyncEnvironmentStatusParams{
		timeout: timeout,
	}
}

// NewGetAutomatedSyncEnvironmentStatusParamsWithContext creates a new GetAutomatedSyncEnvironmentStatusParams object
// with the ability to set a context for a request.
func NewGetAutomatedSyncEnvironmentStatusParamsWithContext(ctx context.Context) *GetAutomatedSyncEnvironmentStatusParams {
	return &GetAutomatedSyncEnvironmentStatusParams{
		Context: ctx,
	}
}

// NewGetAutomatedSyncEnvironmentStatusParamsWithHTTPClient creates a new GetAutomatedSyncEnvironmentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAutomatedSyncEnvironmentStatusParamsWithHTTPClient(client *http.Client) *GetAutomatedSyncEnvironmentStatusParams {
	return &GetAutomatedSyncEnvironmentStatusParams{
		HTTPClient: client,
	}
}

/*
GetAutomatedSyncEnvironmentStatusParams contains all the parameters to send to the API endpoint

	for the get automated sync environment status operation.

	Typically these are written to a http.Request.
*/
type GetAutomatedSyncEnvironmentStatusParams struct {

	// Input.
	Input *models.GetAutomatedSyncEnvironmentStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get automated sync environment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAutomatedSyncEnvironmentStatusParams) WithDefaults() *GetAutomatedSyncEnvironmentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get automated sync environment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAutomatedSyncEnvironmentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) WithTimeout(timeout time.Duration) *GetAutomatedSyncEnvironmentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) WithContext(ctx context.Context) *GetAutomatedSyncEnvironmentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) WithHTTPClient(client *http.Client) *GetAutomatedSyncEnvironmentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) WithInput(input *models.GetAutomatedSyncEnvironmentStatusRequest) *GetAutomatedSyncEnvironmentStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get automated sync environment status params
func (o *GetAutomatedSyncEnvironmentStatusParams) SetInput(input *models.GetAutomatedSyncEnvironmentStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetAutomatedSyncEnvironmentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
