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

// NewUpdateSamlProviderParams creates a new UpdateSamlProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSamlProviderParams() *UpdateSamlProviderParams {
	return &UpdateSamlProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSamlProviderParamsWithTimeout creates a new UpdateSamlProviderParams object
// with the ability to set a timeout on a request.
func NewUpdateSamlProviderParamsWithTimeout(timeout time.Duration) *UpdateSamlProviderParams {
	return &UpdateSamlProviderParams{
		timeout: timeout,
	}
}

// NewUpdateSamlProviderParamsWithContext creates a new UpdateSamlProviderParams object
// with the ability to set a context for a request.
func NewUpdateSamlProviderParamsWithContext(ctx context.Context) *UpdateSamlProviderParams {
	return &UpdateSamlProviderParams{
		Context: ctx,
	}
}

// NewUpdateSamlProviderParamsWithHTTPClient creates a new UpdateSamlProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSamlProviderParamsWithHTTPClient(client *http.Client) *UpdateSamlProviderParams {
	return &UpdateSamlProviderParams{
		HTTPClient: client,
	}
}

/*
UpdateSamlProviderParams contains all the parameters to send to the API endpoint

	for the update saml provider operation.

	Typically these are written to a http.Request.
*/
type UpdateSamlProviderParams struct {

	// Input.
	Input *models.UpdateSamlProviderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update saml provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSamlProviderParams) WithDefaults() *UpdateSamlProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update saml provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSamlProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update saml provider params
func (o *UpdateSamlProviderParams) WithTimeout(timeout time.Duration) *UpdateSamlProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update saml provider params
func (o *UpdateSamlProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update saml provider params
func (o *UpdateSamlProviderParams) WithContext(ctx context.Context) *UpdateSamlProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update saml provider params
func (o *UpdateSamlProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update saml provider params
func (o *UpdateSamlProviderParams) WithHTTPClient(client *http.Client) *UpdateSamlProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update saml provider params
func (o *UpdateSamlProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update saml provider params
func (o *UpdateSamlProviderParams) WithInput(input *models.UpdateSamlProviderRequest) *UpdateSamlProviderParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update saml provider params
func (o *UpdateSamlProviderParams) SetInput(input *models.UpdateSamlProviderRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSamlProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
