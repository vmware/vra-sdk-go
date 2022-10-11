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

// NewGetGerritProjectsByIDUsingGETParams creates a new GetGerritProjectsByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGerritProjectsByIDUsingGETParams() *GetGerritProjectsByIDUsingGETParams {
	return &GetGerritProjectsByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGerritProjectsByIDUsingGETParamsWithTimeout creates a new GetGerritProjectsByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetGerritProjectsByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGerritProjectsByIDUsingGETParams {
	return &GetGerritProjectsByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetGerritProjectsByIDUsingGETParamsWithContext creates a new GetGerritProjectsByIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetGerritProjectsByIDUsingGETParamsWithContext(ctx context.Context) *GetGerritProjectsByIDUsingGETParams {
	return &GetGerritProjectsByIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetGerritProjectsByIDUsingGETParamsWithHTTPClient creates a new GetGerritProjectsByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGerritProjectsByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGerritProjectsByIDUsingGETParams {
	return &GetGerritProjectsByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetGerritProjectsByIDUsingGETParams contains all the parameters to send to the API endpoint

	for the get gerrit projects by Id using g e t operation.

	Typically these are written to a http.Request.
*/
type GetGerritProjectsByIDUsingGETParams struct {

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

// WithDefaults hydrates default values in the get gerrit projects by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGerritProjectsByIDUsingGETParams) WithDefaults() *GetGerritProjectsByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gerrit projects by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGerritProjectsByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGerritProjectsByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithContext(ctx context.Context) *GetGerritProjectsByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGerritProjectsByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithAuthorization(authorization string) *GetGerritProjectsByIDUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetGerritProjectsByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) WithID(id string) *GetGerritProjectsByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gerrit projects by Id using get params
func (o *GetGerritProjectsByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGerritProjectsByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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