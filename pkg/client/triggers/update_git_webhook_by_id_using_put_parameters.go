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

// NewUpdateGitWebhookByIDUsingPUTParams creates a new UpdateGitWebhookByIDUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGitWebhookByIDUsingPUTParams() *UpdateGitWebhookByIDUsingPUTParams {
	return &UpdateGitWebhookByIDUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGitWebhookByIDUsingPUTParamsWithTimeout creates a new UpdateGitWebhookByIDUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateGitWebhookByIDUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateGitWebhookByIDUsingPUTParams {
	return &UpdateGitWebhookByIDUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateGitWebhookByIDUsingPUTParamsWithContext creates a new UpdateGitWebhookByIDUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateGitWebhookByIDUsingPUTParamsWithContext(ctx context.Context) *UpdateGitWebhookByIDUsingPUTParams {
	return &UpdateGitWebhookByIDUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateGitWebhookByIDUsingPUTParamsWithHTTPClient creates a new UpdateGitWebhookByIDUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGitWebhookByIDUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateGitWebhookByIDUsingPUTParams {
	return &UpdateGitWebhookByIDUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateGitWebhookByIDUsingPUTParams contains all the parameters to send to the API endpoint

	for the update git webhook by Id using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateGitWebhookByIDUsingPUTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* GitWebhookSpec.

	   gitWebhookSpec
	*/
	GitWebhookSpec models.GitWebhookSpec

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update git webhook by Id using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGitWebhookByIDUsingPUTParams) WithDefaults() *UpdateGitWebhookByIDUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update git webhook by Id using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGitWebhookByIDUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithContext(ctx context.Context) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithAuthorization(authorization string) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithAPIVersion(aPIVersion string) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGitWebhookSpec adds the gitWebhookSpec to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithGitWebhookSpec(gitWebhookSpec models.GitWebhookSpec) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetGitWebhookSpec(gitWebhookSpec)
	return o
}

// SetGitWebhookSpec adds the gitWebhookSpec to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetGitWebhookSpec(gitWebhookSpec models.GitWebhookSpec) {
	o.GitWebhookSpec = gitWebhookSpec
}

// WithID adds the id to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) WithID(id string) *UpdateGitWebhookByIDUsingPUTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update git webhook by Id using p u t params
func (o *UpdateGitWebhookByIDUsingPUTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGitWebhookByIDUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.GitWebhookSpec); err != nil {
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
