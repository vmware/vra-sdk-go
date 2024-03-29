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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateDockerRegistryWebhookByIDUsingPUTParams creates a new UpdateDockerRegistryWebhookByIDUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDockerRegistryWebhookByIDUsingPUTParams() *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	return &UpdateDockerRegistryWebhookByIDUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithTimeout creates a new UpdateDockerRegistryWebhookByIDUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	return &UpdateDockerRegistryWebhookByIDUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithContext creates a new UpdateDockerRegistryWebhookByIDUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithContext(ctx context.Context) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	return &UpdateDockerRegistryWebhookByIDUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithHTTPClient creates a new UpdateDockerRegistryWebhookByIDUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDockerRegistryWebhookByIDUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	return &UpdateDockerRegistryWebhookByIDUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateDockerRegistryWebhookByIDUsingPUTParams contains all the parameters to send to the API endpoint

	for the update docker registry webhook by Id using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateDockerRegistryWebhookByIDUsingPUTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* DockerRegistryWebHookSpec.

	   dockerRegistryWebHookSpec
	*/
	DockerRegistryWebHookSpec models.DockerRegistryWebHookSpec

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update docker registry webhook by Id using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithDefaults() *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update docker registry webhook by Id using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithContext(ctx context.Context) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithAuthorization(authorization string) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithAPIVersion(aPIVersion string) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDockerRegistryWebHookSpec adds the dockerRegistryWebHookSpec to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithDockerRegistryWebHookSpec(dockerRegistryWebHookSpec models.DockerRegistryWebHookSpec) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetDockerRegistryWebHookSpec(dockerRegistryWebHookSpec)
	return o
}

// SetDockerRegistryWebHookSpec adds the dockerRegistryWebHookSpec to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetDockerRegistryWebHookSpec(dockerRegistryWebHookSpec models.DockerRegistryWebHookSpec) {
	o.DockerRegistryWebHookSpec = dockerRegistryWebHookSpec
}

// WithID adds the id to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WithID(id string) *UpdateDockerRegistryWebhookByIDUsingPUTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update docker registry webhook by Id using p u t params
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDockerRegistryWebhookByIDUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.DockerRegistryWebHookSpec); err != nil {
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
