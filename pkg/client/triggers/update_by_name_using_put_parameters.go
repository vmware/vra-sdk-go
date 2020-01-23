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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateByNameUsingPUTParams creates a new UpdateByNameUsingPUTParams object
// with the default values initialized.
func NewUpdateByNameUsingPUTParams() *UpdateByNameUsingPUTParams {
	var ()
	return &UpdateByNameUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateByNameUsingPUTParamsWithTimeout creates a new UpdateByNameUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateByNameUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateByNameUsingPUTParams {
	var ()
	return &UpdateByNameUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateByNameUsingPUTParamsWithContext creates a new UpdateByNameUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateByNameUsingPUTParamsWithContext(ctx context.Context) *UpdateByNameUsingPUTParams {
	var ()
	return &UpdateByNameUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateByNameUsingPUTParamsWithHTTPClient creates a new UpdateByNameUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateByNameUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateByNameUsingPUTParams {
	var ()
	return &UpdateByNameUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateByNameUsingPUTParams contains all the parameters to send to the API endpoint
for the update by name using p u t operation typically these are written to a http.Request
*/
type UpdateByNameUsingPUTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*DockerRegistryWebHookSpec
	  dockerRegistryWebHookSpec

	*/
	DockerRegistryWebHookSpec models.DockerRegistryWebHookSpec
	/*Name
	  name

	*/
	Name string
	/*Project
	  project

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateByNameUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithContext(ctx context.Context) *UpdateByNameUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateByNameUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithAPIVersion(aPIVersion string) *UpdateByNameUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDockerRegistryWebHookSpec adds the dockerRegistryWebHookSpec to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithDockerRegistryWebHookSpec(dockerRegistryWebHookSpec models.DockerRegistryWebHookSpec) *UpdateByNameUsingPUTParams {
	o.SetDockerRegistryWebHookSpec(dockerRegistryWebHookSpec)
	return o
}

// SetDockerRegistryWebHookSpec adds the dockerRegistryWebHookSpec to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetDockerRegistryWebHookSpec(dockerRegistryWebHookSpec models.DockerRegistryWebHookSpec) {
	o.DockerRegistryWebHookSpec = dockerRegistryWebHookSpec
}

// WithName adds the name to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithName(name string) *UpdateByNameUsingPUTParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) WithProject(project string) *UpdateByNameUsingPUTParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update by name using p u t params
func (o *UpdateByNameUsingPUTParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateByNameUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.DockerRegistryWebHookSpec); err != nil {
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
