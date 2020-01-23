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

// NewPatchByNameUsingPATCHParams creates a new PatchByNameUsingPATCHParams object
// with the default values initialized.
func NewPatchByNameUsingPATCHParams() *PatchByNameUsingPATCHParams {
	var ()
	return &PatchByNameUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchByNameUsingPATCHParamsWithTimeout creates a new PatchByNameUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchByNameUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchByNameUsingPATCHParams {
	var ()
	return &PatchByNameUsingPATCHParams{

		timeout: timeout,
	}
}

// NewPatchByNameUsingPATCHParamsWithContext creates a new PatchByNameUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchByNameUsingPATCHParamsWithContext(ctx context.Context) *PatchByNameUsingPATCHParams {
	var ()
	return &PatchByNameUsingPATCHParams{

		Context: ctx,
	}
}

// NewPatchByNameUsingPATCHParamsWithHTTPClient creates a new PatchByNameUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchByNameUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchByNameUsingPATCHParams {
	var ()
	return &PatchByNameUsingPATCHParams{
		HTTPClient: client,
	}
}

/*PatchByNameUsingPATCHParams contains all the parameters to send to the API endpoint
for the patch by name using p a t c h operation typically these are written to a http.Request
*/
type PatchByNameUsingPATCHParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*DockerRegistryWebhookPatch
	  dockerRegistryWebhookPatch

	*/
	DockerRegistryWebhookPatch models.DockerRegistryWebhookPatch
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

// WithTimeout adds the timeout to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchByNameUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithContext(ctx context.Context) *PatchByNameUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchByNameUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithAPIVersion(aPIVersion string) *PatchByNameUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDockerRegistryWebhookPatch adds the dockerRegistryWebhookPatch to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithDockerRegistryWebhookPatch(dockerRegistryWebhookPatch models.DockerRegistryWebhookPatch) *PatchByNameUsingPATCHParams {
	o.SetDockerRegistryWebhookPatch(dockerRegistryWebhookPatch)
	return o
}

// SetDockerRegistryWebhookPatch adds the dockerRegistryWebhookPatch to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetDockerRegistryWebhookPatch(dockerRegistryWebhookPatch models.DockerRegistryWebhookPatch) {
	o.DockerRegistryWebhookPatch = dockerRegistryWebhookPatch
}

// WithName adds the name to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithName(name string) *PatchByNameUsingPATCHParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) WithProject(project string) *PatchByNameUsingPATCHParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch by name using p a t c h params
func (o *PatchByNameUsingPATCHParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchByNameUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.DockerRegistryWebhookPatch); err != nil {
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
