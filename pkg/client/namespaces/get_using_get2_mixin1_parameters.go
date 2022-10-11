// Code generated by go-swagger; DO NOT EDIT.

package namespaces

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

// NewGetUsingGET2Mixin1Params creates a new GetUsingGET2Mixin1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsingGET2Mixin1Params() *GetUsingGET2Mixin1Params {
	return &GetUsingGET2Mixin1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGET2Mixin1ParamsWithTimeout creates a new GetUsingGET2Mixin1Params object
// with the ability to set a timeout on a request.
func NewGetUsingGET2Mixin1ParamsWithTimeout(timeout time.Duration) *GetUsingGET2Mixin1Params {
	return &GetUsingGET2Mixin1Params{
		timeout: timeout,
	}
}

// NewGetUsingGET2Mixin1ParamsWithContext creates a new GetUsingGET2Mixin1Params object
// with the ability to set a context for a request.
func NewGetUsingGET2Mixin1ParamsWithContext(ctx context.Context) *GetUsingGET2Mixin1Params {
	return &GetUsingGET2Mixin1Params{
		Context: ctx,
	}
}

// NewGetUsingGET2Mixin1ParamsWithHTTPClient creates a new GetUsingGET2Mixin1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsingGET2Mixin1ParamsWithHTTPClient(client *http.Client) *GetUsingGET2Mixin1Params {
	return &GetUsingGET2Mixin1Params{
		HTTPClient: client,
	}
}

/*
GetUsingGET2Mixin1Params contains all the parameters to send to the API endpoint

	for the get using g e t 2 mixin1 operation.

	Typically these are written to a http.Request.
*/
type GetUsingGET2Mixin1Params struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get using g e t 2 mixin1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Mixin1Params) WithDefaults() *GetUsingGET2Mixin1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get using g e t 2 mixin1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Mixin1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) WithTimeout(timeout time.Duration) *GetUsingGET2Mixin1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) WithContext(ctx context.Context) *GetUsingGET2Mixin1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) WithHTTPClient(client *http.Client) *GetUsingGET2Mixin1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) WithID(id strfmt.UUID) *GetUsingGET2Mixin1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t 2 mixin1 params
func (o *GetUsingGET2Mixin1Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGET2Mixin1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
