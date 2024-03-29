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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateUsingPUTParams creates a new UpdateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUsingPUTParams() *UpdateUsingPUTParams {
	return &UpdateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsingPUTParamsWithTimeout creates a new UpdateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUsingPUTParams {
	return &UpdateUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateUsingPUTParamsWithContext creates a new UpdateUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateUsingPUTParamsWithContext(ctx context.Context) *UpdateUsingPUTParams {
	return &UpdateUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateUsingPUTParamsWithHTTPClient creates a new UpdateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUsingPUTParams {
	return &UpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateUsingPUTParams contains all the parameters to send to the API endpoint

	for the update using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateUsingPUTParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Plan.

	   plan
	*/
	Plan *models.ClusterPlan

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPUTParams) WithDefaults() *UpdateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update using p u t params
func (o *UpdateUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update using p u t params
func (o *UpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update using p u t params
func (o *UpdateUsingPUTParams) WithContext(ctx context.Context) *UpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update using p u t params
func (o *UpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update using p u t params
func (o *UpdateUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update using p u t params
func (o *UpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update using p u t params
func (o *UpdateUsingPUTParams) WithID(id strfmt.UUID) *UpdateUsingPUTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update using p u t params
func (o *UpdateUsingPUTParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPlan adds the plan to the update using p u t params
func (o *UpdateUsingPUTParams) WithPlan(plan *models.ClusterPlan) *UpdateUsingPUTParams {
	o.SetPlan(plan)
	return o
}

// SetPlan adds the plan to the update using p u t params
func (o *UpdateUsingPUTParams) SetPlan(plan *models.ClusterPlan) {
	o.Plan = plan
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.Plan != nil {
		if err := r.SetBodyParam(o.Plan); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
