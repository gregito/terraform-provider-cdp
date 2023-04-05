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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// NewGetEventsPayloadDataParams creates a new GetEventsPayloadDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEventsPayloadDataParams() *GetEventsPayloadDataParams {
	return &GetEventsPayloadDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsPayloadDataParamsWithTimeout creates a new GetEventsPayloadDataParams object
// with the ability to set a timeout on a request.
func NewGetEventsPayloadDataParamsWithTimeout(timeout time.Duration) *GetEventsPayloadDataParams {
	return &GetEventsPayloadDataParams{
		timeout: timeout,
	}
}

// NewGetEventsPayloadDataParamsWithContext creates a new GetEventsPayloadDataParams object
// with the ability to set a context for a request.
func NewGetEventsPayloadDataParamsWithContext(ctx context.Context) *GetEventsPayloadDataParams {
	return &GetEventsPayloadDataParams{
		Context: ctx,
	}
}

// NewGetEventsPayloadDataParamsWithHTTPClient creates a new GetEventsPayloadDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEventsPayloadDataParamsWithHTTPClient(client *http.Client) *GetEventsPayloadDataParams {
	return &GetEventsPayloadDataParams{
		HTTPClient: client,
	}
}

/*
GetEventsPayloadDataParams contains all the parameters to send to the API endpoint

	for the get events payload data operation.

	Typically these are written to a http.Request.
*/
type GetEventsPayloadDataParams struct {

	// Input.
	Input *models.GetEventsPayloadDataRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get events payload data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventsPayloadDataParams) WithDefaults() *GetEventsPayloadDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get events payload data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventsPayloadDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get events payload data params
func (o *GetEventsPayloadDataParams) WithTimeout(timeout time.Duration) *GetEventsPayloadDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events payload data params
func (o *GetEventsPayloadDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events payload data params
func (o *GetEventsPayloadDataParams) WithContext(ctx context.Context) *GetEventsPayloadDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events payload data params
func (o *GetEventsPayloadDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events payload data params
func (o *GetEventsPayloadDataParams) WithHTTPClient(client *http.Client) *GetEventsPayloadDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events payload data params
func (o *GetEventsPayloadDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get events payload data params
func (o *GetEventsPayloadDataParams) WithInput(input *models.GetEventsPayloadDataRequest) *GetEventsPayloadDataParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get events payload data params
func (o *GetEventsPayloadDataParams) SetInput(input *models.GetEventsPayloadDataRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsPayloadDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
