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

// NewListRuntimesParams creates a new ListRuntimesParams object
// with the default values initialized.
func NewListRuntimesParams() *ListRuntimesParams {
	var ()
	return &ListRuntimesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRuntimesParamsWithTimeout creates a new ListRuntimesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRuntimesParamsWithTimeout(timeout time.Duration) *ListRuntimesParams {
	var ()
	return &ListRuntimesParams{

		timeout: timeout,
	}
}

// NewListRuntimesParamsWithContext creates a new ListRuntimesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRuntimesParamsWithContext(ctx context.Context) *ListRuntimesParams {
	var ()
	return &ListRuntimesParams{

		Context: ctx,
	}
}

// NewListRuntimesParamsWithHTTPClient creates a new ListRuntimesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRuntimesParamsWithHTTPClient(client *http.Client) *ListRuntimesParams {
	var ()
	return &ListRuntimesParams{
		HTTPClient: client,
	}
}

/*ListRuntimesParams contains all the parameters to send to the API endpoint
for the list runtimes operation typically these are written to a http.Request
*/
type ListRuntimesParams struct {

	/*Input*/
	Input models.ListRuntimesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list runtimes params
func (o *ListRuntimesParams) WithTimeout(timeout time.Duration) *ListRuntimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list runtimes params
func (o *ListRuntimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list runtimes params
func (o *ListRuntimesParams) WithContext(ctx context.Context) *ListRuntimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list runtimes params
func (o *ListRuntimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list runtimes params
func (o *ListRuntimesParams) WithHTTPClient(client *http.Client) *ListRuntimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list runtimes params
func (o *ListRuntimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list runtimes params
func (o *ListRuntimesParams) WithInput(input models.ListRuntimesRequest) *ListRuntimesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list runtimes params
func (o *ListRuntimesParams) SetInput(input models.ListRuntimesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListRuntimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
