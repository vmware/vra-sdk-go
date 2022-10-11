// Code generated by go-swagger; DO NOT EDIT.

package cluster_plans

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

// NewDeleteUsingDELETEParams creates a new DeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETEParams() *DeleteUsingDELETEParams {
	return &DeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETEParamsWithTimeout creates a new DeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETEParams {
	return &DeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETEParamsWithContext creates a new DeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteUsingDELETEParamsWithContext(ctx context.Context) *DeleteUsingDELETEParams {
	return &DeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteUsingDELETEParamsWithHTTPClient creates a new DeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETEParams {
	return &DeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteUsingDELETEParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETEParams) WithDefaults() *DeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) WithContext(ctx context.Context) *DeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) WithID(id strfmt.UUID) *DeleteUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete using d e l e t e params
func (o *DeleteUsingDELETEParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
