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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateEndpointByNameUsingPUTParams creates a new UpdateEndpointByNameUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEndpointByNameUsingPUTParams() *UpdateEndpointByNameUsingPUTParams {
	return &UpdateEndpointByNameUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEndpointByNameUsingPUTParamsWithTimeout creates a new UpdateEndpointByNameUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateEndpointByNameUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateEndpointByNameUsingPUTParams {
	return &UpdateEndpointByNameUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateEndpointByNameUsingPUTParamsWithContext creates a new UpdateEndpointByNameUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateEndpointByNameUsingPUTParamsWithContext(ctx context.Context) *UpdateEndpointByNameUsingPUTParams {
	return &UpdateEndpointByNameUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateEndpointByNameUsingPUTParamsWithHTTPClient creates a new UpdateEndpointByNameUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEndpointByNameUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateEndpointByNameUsingPUTParams {
	return &UpdateEndpointByNameUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateEndpointByNameUsingPUTParams contains all the parameters to send to the API endpoint

	for the update endpoint by name using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateEndpointByNameUsingPUTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Endpoint specification
	*/
	Body models.EndpointSpec

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

// WithDefaults hydrates default values in the update endpoint by name using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEndpointByNameUsingPUTParams) WithDefaults() *UpdateEndpointByNameUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update endpoint by name using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEndpointByNameUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateEndpointByNameUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithContext(ctx context.Context) *UpdateEndpointByNameUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateEndpointByNameUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithAuthorization(authorization string) *UpdateEndpointByNameUsingPUTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithAPIVersion(aPIVersion string) *UpdateEndpointByNameUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithBody(body models.EndpointSpec) *UpdateEndpointByNameUsingPUTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetBody(body models.EndpointSpec) {
	o.Body = body
}

// WithName adds the name to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithName(name string) *UpdateEndpointByNameUsingPUTParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) WithProject(project string) *UpdateEndpointByNameUsingPUTParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update endpoint by name using p u t params
func (o *UpdateEndpointByNameUsingPUTParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEndpointByNameUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
