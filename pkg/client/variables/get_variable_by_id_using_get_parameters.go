// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewGetVariableByIDUsingGETParams creates a new GetVariableByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVariableByIDUsingGETParams() *GetVariableByIDUsingGETParams {
	return &GetVariableByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariableByIDUsingGETParamsWithTimeout creates a new GetVariableByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetVariableByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetVariableByIDUsingGETParams {
	return &GetVariableByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetVariableByIDUsingGETParamsWithContext creates a new GetVariableByIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetVariableByIDUsingGETParamsWithContext(ctx context.Context) *GetVariableByIDUsingGETParams {
	return &GetVariableByIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetVariableByIDUsingGETParamsWithHTTPClient creates a new GetVariableByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVariableByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetVariableByIDUsingGETParams {
	return &GetVariableByIDUsingGETParams{
		HTTPClient: client,
	}
}

/* GetVariableByIDUsingGETParams contains all the parameters to send to the API endpoint
   for the get variable by Id using g e t operation.

   Typically these are written to a http.Request.
*/
type GetVariableByIDUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Variable
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get variable by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableByIDUsingGETParams) WithDefaults() *GetVariableByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get variable by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetVariableByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithContext(ctx context.Context) *GetVariableByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetVariableByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithAuthorization(authorization string) *GetVariableByIDUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetVariableByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) WithID(id string) *GetVariableByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get variable by Id using get params
func (o *GetVariableByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariableByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
