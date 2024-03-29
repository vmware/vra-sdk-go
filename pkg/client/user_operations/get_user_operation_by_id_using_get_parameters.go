// Code generated by go-swagger; DO NOT EDIT.

package user_operations

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

// NewGetUserOperationByIDUsingGETParams creates a new GetUserOperationByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserOperationByIDUsingGETParams() *GetUserOperationByIDUsingGETParams {
	return &GetUserOperationByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserOperationByIDUsingGETParamsWithTimeout creates a new GetUserOperationByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetUserOperationByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetUserOperationByIDUsingGETParams {
	return &GetUserOperationByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetUserOperationByIDUsingGETParamsWithContext creates a new GetUserOperationByIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetUserOperationByIDUsingGETParamsWithContext(ctx context.Context) *GetUserOperationByIDUsingGETParams {
	return &GetUserOperationByIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetUserOperationByIDUsingGETParamsWithHTTPClient creates a new GetUserOperationByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserOperationByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetUserOperationByIDUsingGETParams {
	return &GetUserOperationByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetUserOperationByIDUsingGETParams contains all the parameters to send to the API endpoint

	for the get user operation by Id using g e t operation.

	Typically these are written to a http.Request.
*/
type GetUserOperationByIDUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the User Operation
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user operation by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserOperationByIDUsingGETParams) WithDefaults() *GetUserOperationByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user operation by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserOperationByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetUserOperationByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithContext(ctx context.Context) *GetUserOperationByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetUserOperationByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithAuthorization(authorization string) *GetUserOperationByIDUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetUserOperationByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) WithID(id string) *GetUserOperationByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user operation by Id using get params
func (o *GetUserOperationByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserOperationByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
