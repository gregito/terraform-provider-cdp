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

// NewDescribeClusterTemplateParams creates a new DescribeClusterTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeClusterTemplateParams() *DescribeClusterTemplateParams {
	return &DescribeClusterTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeClusterTemplateParamsWithTimeout creates a new DescribeClusterTemplateParams object
// with the ability to set a timeout on a request.
func NewDescribeClusterTemplateParamsWithTimeout(timeout time.Duration) *DescribeClusterTemplateParams {
	return &DescribeClusterTemplateParams{
		timeout: timeout,
	}
}

// NewDescribeClusterTemplateParamsWithContext creates a new DescribeClusterTemplateParams object
// with the ability to set a context for a request.
func NewDescribeClusterTemplateParamsWithContext(ctx context.Context) *DescribeClusterTemplateParams {
	return &DescribeClusterTemplateParams{
		Context: ctx,
	}
}

// NewDescribeClusterTemplateParamsWithHTTPClient creates a new DescribeClusterTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeClusterTemplateParamsWithHTTPClient(client *http.Client) *DescribeClusterTemplateParams {
	return &DescribeClusterTemplateParams{
		HTTPClient: client,
	}
}

/*
DescribeClusterTemplateParams contains all the parameters to send to the API endpoint

	for the describe cluster template operation.

	Typically these are written to a http.Request.
*/
type DescribeClusterTemplateParams struct {

	// Input.
	Input *models.DescribeClusterTemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeClusterTemplateParams) WithDefaults() *DescribeClusterTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeClusterTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe cluster template params
func (o *DescribeClusterTemplateParams) WithTimeout(timeout time.Duration) *DescribeClusterTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe cluster template params
func (o *DescribeClusterTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe cluster template params
func (o *DescribeClusterTemplateParams) WithContext(ctx context.Context) *DescribeClusterTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe cluster template params
func (o *DescribeClusterTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe cluster template params
func (o *DescribeClusterTemplateParams) WithHTTPClient(client *http.Client) *DescribeClusterTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe cluster template params
func (o *DescribeClusterTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the describe cluster template params
func (o *DescribeClusterTemplateParams) WithInput(input *models.DescribeClusterTemplateRequest) *DescribeClusterTemplateParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the describe cluster template params
func (o *DescribeClusterTemplateParams) SetInput(input *models.DescribeClusterTemplateRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeClusterTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
