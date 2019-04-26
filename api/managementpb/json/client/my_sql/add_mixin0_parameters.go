// Code generated by go-swagger; DO NOT EDIT.

package my_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddMixin0Params creates a new AddMixin0Params object
// with the default values initialized.
func NewAddMixin0Params() *AddMixin0Params {
	var ()
	return &AddMixin0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddMixin0ParamsWithTimeout creates a new AddMixin0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddMixin0ParamsWithTimeout(timeout time.Duration) *AddMixin0Params {
	var ()
	return &AddMixin0Params{

		timeout: timeout,
	}
}

// NewAddMixin0ParamsWithContext creates a new AddMixin0Params object
// with the default values initialized, and the ability to set a context for a request
func NewAddMixin0ParamsWithContext(ctx context.Context) *AddMixin0Params {
	var ()
	return &AddMixin0Params{

		Context: ctx,
	}
}

// NewAddMixin0ParamsWithHTTPClient creates a new AddMixin0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddMixin0ParamsWithHTTPClient(client *http.Client) *AddMixin0Params {
	var ()
	return &AddMixin0Params{
		HTTPClient: client,
	}
}

/*AddMixin0Params contains all the parameters to send to the API endpoint
for the add mixin0 operation typically these are written to a http.Request
*/
type AddMixin0Params struct {

	/*Body*/
	Body AddMixin0Body

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add mixin0 params
func (o *AddMixin0Params) WithTimeout(timeout time.Duration) *AddMixin0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add mixin0 params
func (o *AddMixin0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add mixin0 params
func (o *AddMixin0Params) WithContext(ctx context.Context) *AddMixin0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add mixin0 params
func (o *AddMixin0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add mixin0 params
func (o *AddMixin0Params) WithHTTPClient(client *http.Client) *AddMixin0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add mixin0 params
func (o *AddMixin0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add mixin0 params
func (o *AddMixin0Params) WithBody(body AddMixin0Body) *AddMixin0Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add mixin0 params
func (o *AddMixin0Params) SetBody(body AddMixin0Body) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddMixin0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
