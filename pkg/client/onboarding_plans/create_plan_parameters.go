// Code generated by go-swagger; DO NOT EDIT.

package onboarding_plans

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

// NewCreatePlanParams creates a new CreatePlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePlanParams() *CreatePlanParams {
	return &CreatePlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePlanParamsWithTimeout creates a new CreatePlanParams object
// with the ability to set a timeout on a request.
func NewCreatePlanParamsWithTimeout(timeout time.Duration) *CreatePlanParams {
	return &CreatePlanParams{
		timeout: timeout,
	}
}

// NewCreatePlanParamsWithContext creates a new CreatePlanParams object
// with the ability to set a context for a request.
func NewCreatePlanParamsWithContext(ctx context.Context) *CreatePlanParams {
	return &CreatePlanParams{
		Context: ctx,
	}
}

// NewCreatePlanParamsWithHTTPClient creates a new CreatePlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePlanParamsWithHTTPClient(client *http.Client) *CreatePlanParams {
	return &CreatePlanParams{
		HTTPClient: client,
	}
}

/*
CreatePlanParams contains all the parameters to send to the API endpoint

	for the create plan operation.

	Typically these are written to a http.Request.
*/
type CreatePlanParams struct {

	// Body.
	Body *models.OnboardingPlanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlanParams) WithDefaults() *CreatePlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create plan params
func (o *CreatePlanParams) WithTimeout(timeout time.Duration) *CreatePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create plan params
func (o *CreatePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create plan params
func (o *CreatePlanParams) WithContext(ctx context.Context) *CreatePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create plan params
func (o *CreatePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create plan params
func (o *CreatePlanParams) WithHTTPClient(client *http.Client) *CreatePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create plan params
func (o *CreatePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create plan params
func (o *CreatePlanParams) WithBody(body *models.OnboardingPlanRequest) *CreatePlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create plan params
func (o *CreatePlanParams) SetBody(body *models.OnboardingPlanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
