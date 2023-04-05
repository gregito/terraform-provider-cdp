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

// NewDescribeImageCatalogParams creates a new DescribeImageCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeImageCatalogParams() *DescribeImageCatalogParams {
	return &DescribeImageCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeImageCatalogParamsWithTimeout creates a new DescribeImageCatalogParams object
// with the ability to set a timeout on a request.
func NewDescribeImageCatalogParamsWithTimeout(timeout time.Duration) *DescribeImageCatalogParams {
	return &DescribeImageCatalogParams{
		timeout: timeout,
	}
}

// NewDescribeImageCatalogParamsWithContext creates a new DescribeImageCatalogParams object
// with the ability to set a context for a request.
func NewDescribeImageCatalogParamsWithContext(ctx context.Context) *DescribeImageCatalogParams {
	return &DescribeImageCatalogParams{
		Context: ctx,
	}
}

// NewDescribeImageCatalogParamsWithHTTPClient creates a new DescribeImageCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeImageCatalogParamsWithHTTPClient(client *http.Client) *DescribeImageCatalogParams {
	return &DescribeImageCatalogParams{
		HTTPClient: client,
	}
}

/*
DescribeImageCatalogParams contains all the parameters to send to the API endpoint

	for the describe image catalog operation.

	Typically these are written to a http.Request.
*/
type DescribeImageCatalogParams struct {

	// Input.
	Input *models.DescribeImageCatalogRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe image catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeImageCatalogParams) WithDefaults() *DescribeImageCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe image catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeImageCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe image catalog params
func (o *DescribeImageCatalogParams) WithTimeout(timeout time.Duration) *DescribeImageCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe image catalog params
func (o *DescribeImageCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe image catalog params
func (o *DescribeImageCatalogParams) WithContext(ctx context.Context) *DescribeImageCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe image catalog params
func (o *DescribeImageCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe image catalog params
func (o *DescribeImageCatalogParams) WithHTTPClient(client *http.Client) *DescribeImageCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe image catalog params
func (o *DescribeImageCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the describe image catalog params
func (o *DescribeImageCatalogParams) WithInput(input *models.DescribeImageCatalogRequest) *DescribeImageCatalogParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the describe image catalog params
func (o *DescribeImageCatalogParams) SetInput(input *models.DescribeImageCatalogRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeImageCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
