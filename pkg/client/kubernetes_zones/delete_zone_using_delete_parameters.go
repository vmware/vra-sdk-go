// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_zones

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

// NewDeleteZoneUsingDELETEParams creates a new DeleteZoneUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteZoneUsingDELETEParams() *DeleteZoneUsingDELETEParams {
	return &DeleteZoneUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteZoneUsingDELETEParamsWithTimeout creates a new DeleteZoneUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteZoneUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteZoneUsingDELETEParams {
	return &DeleteZoneUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteZoneUsingDELETEParamsWithContext creates a new DeleteZoneUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteZoneUsingDELETEParamsWithContext(ctx context.Context) *DeleteZoneUsingDELETEParams {
	return &DeleteZoneUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteZoneUsingDELETEParamsWithHTTPClient creates a new DeleteZoneUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteZoneUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteZoneUsingDELETEParams {
	return &DeleteZoneUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteZoneUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete zone using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteZoneUsingDELETEParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete zone using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteZoneUsingDELETEParams) WithDefaults() *DeleteZoneUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete zone using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteZoneUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteZoneUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) WithContext(ctx context.Context) *DeleteZoneUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteZoneUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) WithID(id strfmt.UUID) *DeleteZoneUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete zone using d e l e t e params
func (o *DeleteZoneUsingDELETEParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteZoneUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
