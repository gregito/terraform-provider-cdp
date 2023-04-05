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

// NewListSSHPublicKeysParams creates a new ListSSHPublicKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSSHPublicKeysParams() *ListSSHPublicKeysParams {
	return &ListSSHPublicKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSSHPublicKeysParamsWithTimeout creates a new ListSSHPublicKeysParams object
// with the ability to set a timeout on a request.
func NewListSSHPublicKeysParamsWithTimeout(timeout time.Duration) *ListSSHPublicKeysParams {
	return &ListSSHPublicKeysParams{
		timeout: timeout,
	}
}

// NewListSSHPublicKeysParamsWithContext creates a new ListSSHPublicKeysParams object
// with the ability to set a context for a request.
func NewListSSHPublicKeysParamsWithContext(ctx context.Context) *ListSSHPublicKeysParams {
	return &ListSSHPublicKeysParams{
		Context: ctx,
	}
}

// NewListSSHPublicKeysParamsWithHTTPClient creates a new ListSSHPublicKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSSHPublicKeysParamsWithHTTPClient(client *http.Client) *ListSSHPublicKeysParams {
	return &ListSSHPublicKeysParams{
		HTTPClient: client,
	}
}

/*
ListSSHPublicKeysParams contains all the parameters to send to the API endpoint

	for the list Ssh public keys operation.

	Typically these are written to a http.Request.
*/
type ListSSHPublicKeysParams struct {

	// Input.
	Input *models.ListSSHPublicKeysRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list Ssh public keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSSHPublicKeysParams) WithDefaults() *ListSSHPublicKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list Ssh public keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSSHPublicKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) WithTimeout(timeout time.Duration) *ListSSHPublicKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) WithContext(ctx context.Context) *ListSSHPublicKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) WithHTTPClient(client *http.Client) *ListSSHPublicKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) WithInput(input *models.ListSSHPublicKeysRequest) *ListSSHPublicKeysParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list Ssh public keys params
func (o *ListSSHPublicKeysParams) SetInput(input *models.ListSSHPublicKeysRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListSSHPublicKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
