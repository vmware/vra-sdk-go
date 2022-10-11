// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewGetRegistryWebhookByIDParams creates a new GetRegistryWebhookByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistryWebhookByIDParams() *GetRegistryWebhookByIDParams {
	return &GetRegistryWebhookByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistryWebhookByIDParamsWithTimeout creates a new GetRegistryWebhookByIDParams object
// with the ability to set a timeout on a request.
func NewGetRegistryWebhookByIDParamsWithTimeout(timeout time.Duration) *GetRegistryWebhookByIDParams {
	return &GetRegistryWebhookByIDParams{
		timeout: timeout,
	}
}

// NewGetRegistryWebhookByIDParamsWithContext creates a new GetRegistryWebhookByIDParams object
// with the ability to set a context for a request.
func NewGetRegistryWebhookByIDParamsWithContext(ctx context.Context) *GetRegistryWebhookByIDParams {
	return &GetRegistryWebhookByIDParams{
		Context: ctx,
	}
}

// NewGetRegistryWebhookByIDParamsWithHTTPClient creates a new GetRegistryWebhookByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistryWebhookByIDParamsWithHTTPClient(client *http.Client) *GetRegistryWebhookByIDParams {
	return &GetRegistryWebhookByIDParams{
		HTTPClient: client,
	}
}

/*
GetRegistryWebhookByIDParams contains all the parameters to send to the API endpoint

	for the get registry webhook by ID operation.

	Typically these are written to a http.Request.
*/
type GetRegistryWebhookByIDParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registry webhook by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryWebhookByIDParams) WithDefaults() *GetRegistryWebhookByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registry webhook by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryWebhookByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithTimeout(timeout time.Duration) *GetRegistryWebhookByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithContext(ctx context.Context) *GetRegistryWebhookByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithHTTPClient(client *http.Client) *GetRegistryWebhookByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithAuthorization(authorization string) *GetRegistryWebhookByIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithAPIVersion(aPIVersion string) *GetRegistryWebhookByIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) WithID(id string) *GetRegistryWebhookByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get registry webhook by ID params
func (o *GetRegistryWebhookByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistryWebhookByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
