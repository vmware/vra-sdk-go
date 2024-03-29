// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewUpdatePipelineByNameUsingPUTParams creates a new UpdatePipelineByNameUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePipelineByNameUsingPUTParams() *UpdatePipelineByNameUsingPUTParams {
	return &UpdatePipelineByNameUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePipelineByNameUsingPUTParamsWithTimeout creates a new UpdatePipelineByNameUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdatePipelineByNameUsingPUTParamsWithTimeout(timeout time.Duration) *UpdatePipelineByNameUsingPUTParams {
	return &UpdatePipelineByNameUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdatePipelineByNameUsingPUTParamsWithContext creates a new UpdatePipelineByNameUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdatePipelineByNameUsingPUTParamsWithContext(ctx context.Context) *UpdatePipelineByNameUsingPUTParams {
	return &UpdatePipelineByNameUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdatePipelineByNameUsingPUTParamsWithHTTPClient creates a new UpdatePipelineByNameUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePipelineByNameUsingPUTParamsWithHTTPClient(client *http.Client) *UpdatePipelineByNameUsingPUTParams {
	return &UpdatePipelineByNameUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdatePipelineByNameUsingPUTParams contains all the parameters to send to the API endpoint

	for the update pipeline by name using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdatePipelineByNameUsingPUTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Pipeline specification
	*/
	Body models.PipelineSpec

	/* Name.

	   The name of the Pipeline
	*/
	Name string

	/* Project.

	   The project the Pipeline belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update pipeline by name using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineByNameUsingPUTParams) WithDefaults() *UpdatePipelineByNameUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update pipeline by name using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineByNameUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithTimeout(timeout time.Duration) *UpdatePipelineByNameUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithContext(ctx context.Context) *UpdatePipelineByNameUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithHTTPClient(client *http.Client) *UpdatePipelineByNameUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithAuthorization(authorization string) *UpdatePipelineByNameUsingPUTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithAPIVersion(aPIVersion string) *UpdatePipelineByNameUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithBody(body models.PipelineSpec) *UpdatePipelineByNameUsingPUTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetBody(body models.PipelineSpec) {
	o.Body = body
}

// WithName adds the name to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithName(name string) *UpdatePipelineByNameUsingPUTParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) WithProject(project string) *UpdatePipelineByNameUsingPUTParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update pipeline by name using p u t params
func (o *UpdatePipelineByNameUsingPUTParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePipelineByNameUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
