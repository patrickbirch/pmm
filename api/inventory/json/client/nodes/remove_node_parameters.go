// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveNodeParams creates a new RemoveNodeParams object
// with the default values initialized.
func NewRemoveNodeParams() *RemoveNodeParams {
	var ()
	return &RemoveNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveNodeParamsWithTimeout creates a new RemoveNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveNodeParamsWithTimeout(timeout time.Duration) *RemoveNodeParams {
	var ()
	return &RemoveNodeParams{

		timeout: timeout,
	}
}

// NewRemoveNodeParamsWithContext creates a new RemoveNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveNodeParamsWithContext(ctx context.Context) *RemoveNodeParams {
	var ()
	return &RemoveNodeParams{

		Context: ctx,
	}
}

// NewRemoveNodeParamsWithHTTPClient creates a new RemoveNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveNodeParamsWithHTTPClient(client *http.Client) *RemoveNodeParams {
	var ()
	return &RemoveNodeParams{
		HTTPClient: client,
	}
}

/*RemoveNodeParams contains all the parameters to send to the API endpoint
for the remove node operation typically these are written to a http.Request
*/
type RemoveNodeParams struct {

	/*Body*/
	Body RemoveNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove node params
func (o *RemoveNodeParams) WithTimeout(timeout time.Duration) *RemoveNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove node params
func (o *RemoveNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove node params
func (o *RemoveNodeParams) WithContext(ctx context.Context) *RemoveNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove node params
func (o *RemoveNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove node params
func (o *RemoveNodeParams) WithHTTPClient(client *http.Client) *RemoveNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove node params
func (o *RemoveNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove node params
func (o *RemoveNodeParams) WithBody(body RemoveNodeBody) *RemoveNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove node params
func (o *RemoveNodeParams) SetBody(body RemoveNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}