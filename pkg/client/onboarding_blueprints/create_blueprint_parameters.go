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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateBlueprintParams creates a new CreateBlueprintParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlueprintParams() *CreateBlueprintParams {
	return &CreateBlueprintParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintParamsWithTimeout creates a new CreateBlueprintParams object
// with the ability to set a timeout on a request.
func NewCreateBlueprintParamsWithTimeout(timeout time.Duration) *CreateBlueprintParams {
	return &CreateBlueprintParams{
		timeout: timeout,
	}
}

// NewCreateBlueprintParamsWithContext creates a new CreateBlueprintParams object
// with the ability to set a context for a request.
func NewCreateBlueprintParamsWithContext(ctx context.Context) *CreateBlueprintParams {
	return &CreateBlueprintParams{
		Context: ctx,
	}
}

// NewCreateBlueprintParamsWithHTTPClient creates a new CreateBlueprintParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlueprintParamsWithHTTPClient(client *http.Client) *CreateBlueprintParams {
	return &CreateBlueprintParams{
		HTTPClient: client,
	}
}

/*
CreateBlueprintParams contains all the parameters to send to the API endpoint

	for the create blueprint operation.

	Typically these are written to a http.Request.
*/
type CreateBlueprintParams struct {

	// Body.
	Body *models.OnboardingBlueprintRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blueprint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintParams) WithDefaults() *CreateBlueprintParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blueprint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blueprint params
func (o *CreateBlueprintParams) WithTimeout(timeout time.Duration) *CreateBlueprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint params
func (o *CreateBlueprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint params
func (o *CreateBlueprintParams) WithContext(ctx context.Context) *CreateBlueprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint params
func (o *CreateBlueprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint params
func (o *CreateBlueprintParams) WithHTTPClient(client *http.Client) *CreateBlueprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint params
func (o *CreateBlueprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blueprint params
func (o *CreateBlueprintParams) WithBody(body *models.OnboardingBlueprintRequest) *CreateBlueprintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blueprint params
func (o *CreateBlueprintParams) SetBody(body *models.OnboardingBlueprintRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
