// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewGetByNameUsingGETParams creates a new GetByNameUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetByNameUsingGETParams() *GetByNameUsingGETParams {
	return &GetByNameUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetByNameUsingGETParamsWithTimeout creates a new GetByNameUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetByNameUsingGETParamsWithTimeout(timeout time.Duration) *GetByNameUsingGETParams {
	return &GetByNameUsingGETParams{
		timeout: timeout,
	}
}

// NewGetByNameUsingGETParamsWithContext creates a new GetByNameUsingGETParams object
// with the ability to set a context for a request.
func NewGetByNameUsingGETParamsWithContext(ctx context.Context) *GetByNameUsingGETParams {
	return &GetByNameUsingGETParams{
		Context: ctx,
	}
}

// NewGetByNameUsingGETParamsWithHTTPClient creates a new GetByNameUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetByNameUsingGETParamsWithHTTPClient(client *http.Client) *GetByNameUsingGETParams {
	return &GetByNameUsingGETParams{
		HTTPClient: client,
	}
}

/* GetByNameUsingGETParams contains all the parameters to send to the API endpoint
   for the get by name using g e t operation.

   Typically these are written to a http.Request.
*/
type GetByNameUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Name.

	   The name of the Endpoint
	*/
	Name string

	/* Project.

	   The project the Endpoint belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get by name using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByNameUsingGETParams) WithDefaults() *GetByNameUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get by name using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByNameUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get by name using get params
func (o *GetByNameUsingGETParams) WithTimeout(timeout time.Duration) *GetByNameUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by name using get params
func (o *GetByNameUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by name using get params
func (o *GetByNameUsingGETParams) WithContext(ctx context.Context) *GetByNameUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by name using get params
func (o *GetByNameUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by name using get params
func (o *GetByNameUsingGETParams) WithHTTPClient(client *http.Client) *GetByNameUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by name using get params
func (o *GetByNameUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get by name using get params
func (o *GetByNameUsingGETParams) WithAuthorization(authorization string) *GetByNameUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get by name using get params
func (o *GetByNameUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get by name using get params
func (o *GetByNameUsingGETParams) WithAPIVersion(aPIVersion string) *GetByNameUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get by name using get params
func (o *GetByNameUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the get by name using get params
func (o *GetByNameUsingGETParams) WithName(name string) *GetByNameUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get by name using get params
func (o *GetByNameUsingGETParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the get by name using get params
func (o *GetByNameUsingGETParams) WithProject(project string) *GetByNameUsingGETParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get by name using get params
func (o *GetByNameUsingGETParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetByNameUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
