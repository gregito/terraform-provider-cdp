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

// NewDeleteRecipesParams creates a new DeleteRecipesParams object
// with the default values initialized.
func NewDeleteRecipesParams() *DeleteRecipesParams {
	var ()
	return &DeleteRecipesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecipesParamsWithTimeout creates a new DeleteRecipesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecipesParamsWithTimeout(timeout time.Duration) *DeleteRecipesParams {
	var ()
	return &DeleteRecipesParams{

		timeout: timeout,
	}
}

// NewDeleteRecipesParamsWithContext creates a new DeleteRecipesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecipesParamsWithContext(ctx context.Context) *DeleteRecipesParams {
	var ()
	return &DeleteRecipesParams{

		Context: ctx,
	}
}

// NewDeleteRecipesParamsWithHTTPClient creates a new DeleteRecipesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecipesParamsWithHTTPClient(client *http.Client) *DeleteRecipesParams {
	var ()
	return &DeleteRecipesParams{
		HTTPClient: client,
	}
}

/*DeleteRecipesParams contains all the parameters to send to the API endpoint
for the delete recipes operation typically these are written to a http.Request
*/
type DeleteRecipesParams struct {

	/*Input*/
	Input *models.DeleteRecipesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recipes params
func (o *DeleteRecipesParams) WithTimeout(timeout time.Duration) *DeleteRecipesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recipes params
func (o *DeleteRecipesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recipes params
func (o *DeleteRecipesParams) WithContext(ctx context.Context) *DeleteRecipesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recipes params
func (o *DeleteRecipesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recipes params
func (o *DeleteRecipesParams) WithHTTPClient(client *http.Client) *DeleteRecipesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recipes params
func (o *DeleteRecipesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete recipes params
func (o *DeleteRecipesParams) WithInput(input *models.DeleteRecipesRequest) *DeleteRecipesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete recipes params
func (o *DeleteRecipesParams) SetInput(input *models.DeleteRecipesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecipesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
