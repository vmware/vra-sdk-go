// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

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

// NewCreateCustomIntegrationUsingPOSTParams creates a new CreateCustomIntegrationUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCustomIntegrationUsingPOSTParams() *CreateCustomIntegrationUsingPOSTParams {
	return &CreateCustomIntegrationUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomIntegrationUsingPOSTParamsWithTimeout creates a new CreateCustomIntegrationUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateCustomIntegrationUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateCustomIntegrationUsingPOSTParams {
	return &CreateCustomIntegrationUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateCustomIntegrationUsingPOSTParamsWithContext creates a new CreateCustomIntegrationUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateCustomIntegrationUsingPOSTParamsWithContext(ctx context.Context) *CreateCustomIntegrationUsingPOSTParams {
	return &CreateCustomIntegrationUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateCustomIntegrationUsingPOSTParamsWithHTTPClient creates a new CreateCustomIntegrationUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCustomIntegrationUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateCustomIntegrationUsingPOSTParams {
	return &CreateCustomIntegrationUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CreateCustomIntegrationUsingPOSTParams contains all the parameters to send to the API endpoint
   for the create custom integration using p o s t operation.

   Typically these are written to a http.Request.
*/
type CreateCustomIntegrationUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Custom Integration specification
	*/
	Body models.CustomIntegrationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create custom integration using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCustomIntegrationUsingPOSTParams) WithDefaults() *CreateCustomIntegrationUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create custom integration using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCustomIntegrationUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateCustomIntegrationUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithContext(ctx context.Context) *CreateCustomIntegrationUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateCustomIntegrationUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithAuthorization(authorization string) *CreateCustomIntegrationUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithAPIVersion(aPIVersion string) *CreateCustomIntegrationUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) WithBody(body models.CustomIntegrationSpec) *CreateCustomIntegrationUsingPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create custom integration using p o s t params
func (o *CreateCustomIntegrationUsingPOSTParams) SetBody(body models.CustomIntegrationSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomIntegrationUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
