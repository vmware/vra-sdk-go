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

// NewGetKubernetesPropertiesUsingGETParams creates a new GetKubernetesPropertiesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesPropertiesUsingGETParams() *GetKubernetesPropertiesUsingGETParams {
	return &GetKubernetesPropertiesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesPropertiesUsingGETParamsWithTimeout creates a new GetKubernetesPropertiesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesPropertiesUsingGETParamsWithTimeout(timeout time.Duration) *GetKubernetesPropertiesUsingGETParams {
	return &GetKubernetesPropertiesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetKubernetesPropertiesUsingGETParamsWithContext creates a new GetKubernetesPropertiesUsingGETParams object
// with the ability to set a context for a request.
func NewGetKubernetesPropertiesUsingGETParamsWithContext(ctx context.Context) *GetKubernetesPropertiesUsingGETParams {
	return &GetKubernetesPropertiesUsingGETParams{
		Context: ctx,
	}
}

// NewGetKubernetesPropertiesUsingGETParamsWithHTTPClient creates a new GetKubernetesPropertiesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesPropertiesUsingGETParamsWithHTTPClient(client *http.Client) *GetKubernetesPropertiesUsingGETParams {
	return &GetKubernetesPropertiesUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesPropertiesUsingGETParams contains all the parameters to send to the API endpoint

	for the get kubernetes properties using g e t operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesPropertiesUsingGETParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes properties using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesPropertiesUsingGETParams) WithDefaults() *GetKubernetesPropertiesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes properties using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesPropertiesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) WithTimeout(timeout time.Duration) *GetKubernetesPropertiesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) WithContext(ctx context.Context) *GetKubernetesPropertiesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) WithHTTPClient(client *http.Client) *GetKubernetesPropertiesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) WithID(id strfmt.UUID) *GetKubernetesPropertiesUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes properties using get params
func (o *GetKubernetesPropertiesUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesPropertiesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
