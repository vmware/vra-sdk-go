// Code generated by go-swagger; DO NOT EDIT.

package onboarding_blueprints

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

// NewDeleteBlueprintByIDParams creates a new DeleteBlueprintByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlueprintByIDParams() *DeleteBlueprintByIDParams {
	return &DeleteBlueprintByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlueprintByIDParamsWithTimeout creates a new DeleteBlueprintByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteBlueprintByIDParamsWithTimeout(timeout time.Duration) *DeleteBlueprintByIDParams {
	return &DeleteBlueprintByIDParams{
		timeout: timeout,
	}
}

// NewDeleteBlueprintByIDParamsWithContext creates a new DeleteBlueprintByIDParams object
// with the ability to set a context for a request.
func NewDeleteBlueprintByIDParamsWithContext(ctx context.Context) *DeleteBlueprintByIDParams {
	return &DeleteBlueprintByIDParams{
		Context: ctx,
	}
}

// NewDeleteBlueprintByIDParamsWithHTTPClient creates a new DeleteBlueprintByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlueprintByIDParamsWithHTTPClient(client *http.Client) *DeleteBlueprintByIDParams {
	return &DeleteBlueprintByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteBlueprintByIDParams contains all the parameters to send to the API endpoint

	for the delete blueprint by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteBlueprintByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete blueprint by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlueprintByIDParams) WithDefaults() *DeleteBlueprintByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete blueprint by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlueprintByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) WithTimeout(timeout time.Duration) *DeleteBlueprintByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) WithContext(ctx context.Context) *DeleteBlueprintByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) WithHTTPClient(client *http.Client) *DeleteBlueprintByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) WithID(id string) *DeleteBlueprintByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete blueprint by Id params
func (o *DeleteBlueprintByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlueprintByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
