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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateDeploymentWithResourcesParams creates a new CreateDeploymentWithResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeploymentWithResourcesParams() *CreateDeploymentWithResourcesParams {
	return &CreateDeploymentWithResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentWithResourcesParamsWithTimeout creates a new CreateDeploymentWithResourcesParams object
// with the ability to set a timeout on a request.
func NewCreateDeploymentWithResourcesParamsWithTimeout(timeout time.Duration) *CreateDeploymentWithResourcesParams {
	return &CreateDeploymentWithResourcesParams{
		timeout: timeout,
	}
}

// NewCreateDeploymentWithResourcesParamsWithContext creates a new CreateDeploymentWithResourcesParams object
// with the ability to set a context for a request.
func NewCreateDeploymentWithResourcesParamsWithContext(ctx context.Context) *CreateDeploymentWithResourcesParams {
	return &CreateDeploymentWithResourcesParams{
		Context: ctx,
	}
}

// NewCreateDeploymentWithResourcesParamsWithHTTPClient creates a new CreateDeploymentWithResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeploymentWithResourcesParamsWithHTTPClient(client *http.Client) *CreateDeploymentWithResourcesParams {
	return &CreateDeploymentWithResourcesParams{
		HTTPClient: client,
	}
}

/*
CreateDeploymentWithResourcesParams contains all the parameters to send to the API endpoint

	for the create deployment with resources operation.

	Typically these are written to a http.Request.
*/
type CreateDeploymentWithResourcesParams struct {

	// Body.
	Body *models.CreateDeploymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create deployment with resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentWithResourcesParams) WithDefaults() *CreateDeploymentWithResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create deployment with resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentWithResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) WithTimeout(timeout time.Duration) *CreateDeploymentWithResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) WithContext(ctx context.Context) *CreateDeploymentWithResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) WithHTTPClient(client *http.Client) *CreateDeploymentWithResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) WithBody(body *models.CreateDeploymentRequest) *CreateDeploymentWithResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create deployment with resources params
func (o *CreateDeploymentWithResourcesParams) SetBody(body *models.CreateDeploymentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentWithResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}