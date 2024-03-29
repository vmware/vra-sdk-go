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

// NewGetGerritTriggerByIDUsingGETParams creates a new GetGerritTriggerByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGerritTriggerByIDUsingGETParams() *GetGerritTriggerByIDUsingGETParams {
	return &GetGerritTriggerByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGerritTriggerByIDUsingGETParamsWithTimeout creates a new GetGerritTriggerByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetGerritTriggerByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGerritTriggerByIDUsingGETParams {
	return &GetGerritTriggerByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetGerritTriggerByIDUsingGETParamsWithContext creates a new GetGerritTriggerByIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetGerritTriggerByIDUsingGETParamsWithContext(ctx context.Context) *GetGerritTriggerByIDUsingGETParams {
	return &GetGerritTriggerByIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetGerritTriggerByIDUsingGETParamsWithHTTPClient creates a new GetGerritTriggerByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGerritTriggerByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGerritTriggerByIDUsingGETParams {
	return &GetGerritTriggerByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetGerritTriggerByIDUsingGETParams contains all the parameters to send to the API endpoint

	for the get gerrit trigger by Id using g e t operation.

	Typically these are written to a http.Request.
*/
type GetGerritTriggerByIDUsingGETParams struct {

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

// WithDefaults hydrates default values in the get gerrit trigger by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGerritTriggerByIDUsingGETParams) WithDefaults() *GetGerritTriggerByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gerrit trigger by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGerritTriggerByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGerritTriggerByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithContext(ctx context.Context) *GetGerritTriggerByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGerritTriggerByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithAuthorization(authorization string) *GetGerritTriggerByIDUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetGerritTriggerByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) WithID(id string) *GetGerritTriggerByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gerrit trigger by Id using get params
func (o *GetGerritTriggerByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGerritTriggerByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
