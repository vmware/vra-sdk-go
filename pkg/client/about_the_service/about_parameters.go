// Code generated by go-swagger; DO NOT EDIT.

package about_the_service

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
)

// NewAboutParams creates a new AboutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAboutParams() *AboutParams {
	return &AboutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAboutParamsWithTimeout creates a new AboutParams object
// with the ability to set a timeout on a request.
func NewAboutParamsWithTimeout(timeout time.Duration) *AboutParams {
	return &AboutParams{
		timeout: timeout,
	}
}

// NewAboutParamsWithContext creates a new AboutParams object
// with the ability to set a context for a request.
func NewAboutParamsWithContext(ctx context.Context) *AboutParams {
	return &AboutParams{
		Context: ctx,
	}
}

// NewAboutParamsWithHTTPClient creates a new AboutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAboutParamsWithHTTPClient(client *http.Client) *AboutParams {
	return &AboutParams{
		HTTPClient: client,
	}
}

/*
AboutParams contains all the parameters to send to the API endpoint

	for the about operation.

	Typically these are written to a http.Request.
*/
type AboutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the about params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AboutParams) WithDefaults() *AboutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the about params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AboutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the about params
func (o *AboutParams) WithTimeout(timeout time.Duration) *AboutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the about params
func (o *AboutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the about params
func (o *AboutParams) WithContext(ctx context.Context) *AboutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the about params
func (o *AboutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the about params
func (o *AboutParams) WithHTTPClient(client *http.Client) *AboutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the about params
func (o *AboutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AboutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
