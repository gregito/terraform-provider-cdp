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

// NewCreateRecipeParams creates a new CreateRecipeParams object
// with the default values initialized.
func NewCreateRecipeParams() *CreateRecipeParams {
	var ()
	return &CreateRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRecipeParamsWithTimeout creates a new CreateRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRecipeParamsWithTimeout(timeout time.Duration) *CreateRecipeParams {
	var ()
	return &CreateRecipeParams{

		timeout: timeout,
	}
}

// NewCreateRecipeParamsWithContext creates a new CreateRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRecipeParamsWithContext(ctx context.Context) *CreateRecipeParams {
	var ()
	return &CreateRecipeParams{

		Context: ctx,
	}
}

// NewCreateRecipeParamsWithHTTPClient creates a new CreateRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRecipeParamsWithHTTPClient(client *http.Client) *CreateRecipeParams {
	var ()
	return &CreateRecipeParams{
		HTTPClient: client,
	}
}

/*CreateRecipeParams contains all the parameters to send to the API endpoint
for the create recipe operation typically these are written to a http.Request
*/
type CreateRecipeParams struct {

	/*Input*/
	Input *models.CreateRecipeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create recipe params
func (o *CreateRecipeParams) WithTimeout(timeout time.Duration) *CreateRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create recipe params
func (o *CreateRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create recipe params
func (o *CreateRecipeParams) WithContext(ctx context.Context) *CreateRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create recipe params
func (o *CreateRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create recipe params
func (o *CreateRecipeParams) WithHTTPClient(client *http.Client) *CreateRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create recipe params
func (o *CreateRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create recipe params
func (o *CreateRecipeParams) WithInput(input *models.CreateRecipeRequest) *CreateRecipeParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create recipe params
func (o *CreateRecipeParams) SetInput(input *models.CreateRecipeRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
