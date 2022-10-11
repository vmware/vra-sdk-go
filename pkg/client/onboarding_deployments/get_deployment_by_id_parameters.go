// Code generated by go-swagger; DO NOT EDIT.

package onboarding_deployments

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

// NewGetDeploymentByIDParams creates a new GetDeploymentByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentByIDParams() *GetDeploymentByIDParams {
	return &GetDeploymentByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentByIDParamsWithTimeout creates a new GetDeploymentByIDParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentByIDParamsWithTimeout(timeout time.Duration) *GetDeploymentByIDParams {
	return &GetDeploymentByIDParams{
		timeout: timeout,
	}
}

// NewGetDeploymentByIDParamsWithContext creates a new GetDeploymentByIDParams object
// with the ability to set a context for a request.
func NewGetDeploymentByIDParamsWithContext(ctx context.Context) *GetDeploymentByIDParams {
	return &GetDeploymentByIDParams{
		Context: ctx,
	}
}

// NewGetDeploymentByIDParamsWithHTTPClient creates a new GetDeploymentByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentByIDParamsWithHTTPClient(client *http.Client) *GetDeploymentByIDParams {
	return &GetDeploymentByIDParams{
		HTTPClient: client,
	}
}

/*
GetDeploymentByIDParams contains all the parameters to send to the API endpoint

	for the get deployment by Id operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentByIDParams) WithDefaults() *GetDeploymentByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployment by Id params
func (o *GetDeploymentByIDParams) WithTimeout(timeout time.Duration) *GetDeploymentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment by Id params
func (o *GetDeploymentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment by Id params
func (o *GetDeploymentByIDParams) WithContext(ctx context.Context) *GetDeploymentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment by Id params
func (o *GetDeploymentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment by Id params
func (o *GetDeploymentByIDParams) WithHTTPClient(client *http.Client) *GetDeploymentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment by Id params
func (o *GetDeploymentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get deployment by Id params
func (o *GetDeploymentByIDParams) WithID(id string) *GetDeploymentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get deployment by Id params
func (o *GetDeploymentByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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