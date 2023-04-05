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

// NewReplaceRecipesParams creates a new ReplaceRecipesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceRecipesParams() *ReplaceRecipesParams {
	return &ReplaceRecipesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRecipesParamsWithTimeout creates a new ReplaceRecipesParams object
// with the ability to set a timeout on a request.
func NewReplaceRecipesParamsWithTimeout(timeout time.Duration) *ReplaceRecipesParams {
	return &ReplaceRecipesParams{
		timeout: timeout,
	}
}

// NewReplaceRecipesParamsWithContext creates a new ReplaceRecipesParams object
// with the ability to set a context for a request.
func NewReplaceRecipesParamsWithContext(ctx context.Context) *ReplaceRecipesParams {
	return &ReplaceRecipesParams{
		Context: ctx,
	}
}

// NewReplaceRecipesParamsWithHTTPClient creates a new ReplaceRecipesParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceRecipesParamsWithHTTPClient(client *http.Client) *ReplaceRecipesParams {
	return &ReplaceRecipesParams{
		HTTPClient: client,
	}
}

/*
ReplaceRecipesParams contains all the parameters to send to the API endpoint

	for the replace recipes operation.

	Typically these are written to a http.Request.
*/
type ReplaceRecipesParams struct {

	// Input.
	Input *models.ReplaceRecipesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace recipes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRecipesParams) WithDefaults() *ReplaceRecipesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace recipes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRecipesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace recipes params
func (o *ReplaceRecipesParams) WithTimeout(timeout time.Duration) *ReplaceRecipesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace recipes params
func (o *ReplaceRecipesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace recipes params
func (o *ReplaceRecipesParams) WithContext(ctx context.Context) *ReplaceRecipesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace recipes params
func (o *ReplaceRecipesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace recipes params
func (o *ReplaceRecipesParams) WithHTTPClient(client *http.Client) *ReplaceRecipesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace recipes params
func (o *ReplaceRecipesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the replace recipes params
func (o *ReplaceRecipesParams) WithInput(input *models.ReplaceRecipesRequest) *ReplaceRecipesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the replace recipes params
func (o *ReplaceRecipesParams) SetInput(input *models.ReplaceRecipesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRecipesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
