// Code generated by go-swagger; DO NOT EDIT.

package onboarding_plan_execution

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

// NewGetPlanExecutionByIDParams creates a new GetPlanExecutionByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlanExecutionByIDParams() *GetPlanExecutionByIDParams {
	return &GetPlanExecutionByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlanExecutionByIDParamsWithTimeout creates a new GetPlanExecutionByIDParams object
// with the ability to set a timeout on a request.
func NewGetPlanExecutionByIDParamsWithTimeout(timeout time.Duration) *GetPlanExecutionByIDParams {
	return &GetPlanExecutionByIDParams{
		timeout: timeout,
	}
}

// NewGetPlanExecutionByIDParamsWithContext creates a new GetPlanExecutionByIDParams object
// with the ability to set a context for a request.
func NewGetPlanExecutionByIDParamsWithContext(ctx context.Context) *GetPlanExecutionByIDParams {
	return &GetPlanExecutionByIDParams{
		Context: ctx,
	}
}

// NewGetPlanExecutionByIDParamsWithHTTPClient creates a new GetPlanExecutionByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlanExecutionByIDParamsWithHTTPClient(client *http.Client) *GetPlanExecutionByIDParams {
	return &GetPlanExecutionByIDParams{
		HTTPClient: client,
	}
}

/*
GetPlanExecutionByIDParams contains all the parameters to send to the API endpoint

	for the get plan execution by Id operation.

	Typically these are written to a http.Request.
*/
type GetPlanExecutionByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plan execution by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanExecutionByIDParams) WithDefaults() *GetPlanExecutionByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plan execution by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanExecutionByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) WithTimeout(timeout time.Duration) *GetPlanExecutionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) WithContext(ctx context.Context) *GetPlanExecutionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) WithHTTPClient(client *http.Client) *GetPlanExecutionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) WithID(id string) *GetPlanExecutionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get plan execution by Id params
func (o *GetPlanExecutionByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlanExecutionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
