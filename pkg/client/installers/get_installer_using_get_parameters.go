// Code generated by go-swagger; DO NOT EDIT.

package installers

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

// NewGetInstallerUsingGETParams creates a new GetInstallerUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallerUsingGETParams() *GetInstallerUsingGETParams {
	return &GetInstallerUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallerUsingGETParamsWithTimeout creates a new GetInstallerUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetInstallerUsingGETParamsWithTimeout(timeout time.Duration) *GetInstallerUsingGETParams {
	return &GetInstallerUsingGETParams{
		timeout: timeout,
	}
}

// NewGetInstallerUsingGETParamsWithContext creates a new GetInstallerUsingGETParams object
// with the ability to set a context for a request.
func NewGetInstallerUsingGETParamsWithContext(ctx context.Context) *GetInstallerUsingGETParams {
	return &GetInstallerUsingGETParams{
		Context: ctx,
	}
}

// NewGetInstallerUsingGETParamsWithHTTPClient creates a new GetInstallerUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallerUsingGETParamsWithHTTPClient(client *http.Client) *GetInstallerUsingGETParams {
	return &GetInstallerUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetInstallerUsingGETParams contains all the parameters to send to the API endpoint

	for the get installer using g e t operation.

	Typically these are written to a http.Request.
*/
type GetInstallerUsingGETParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get installer using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallerUsingGETParams) WithDefaults() *GetInstallerUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get installer using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallerUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get installer using get params
func (o *GetInstallerUsingGETParams) WithTimeout(timeout time.Duration) *GetInstallerUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get installer using get params
func (o *GetInstallerUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get installer using get params
func (o *GetInstallerUsingGETParams) WithContext(ctx context.Context) *GetInstallerUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get installer using get params
func (o *GetInstallerUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get installer using get params
func (o *GetInstallerUsingGETParams) WithHTTPClient(client *http.Client) *GetInstallerUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get installer using get params
func (o *GetInstallerUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get installer using get params
func (o *GetInstallerUsingGETParams) WithID(id strfmt.UUID) *GetInstallerUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get installer using get params
func (o *GetInstallerUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallerUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
