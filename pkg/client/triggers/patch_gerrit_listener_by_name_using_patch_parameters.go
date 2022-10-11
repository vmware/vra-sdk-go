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

// NewPatchGerritListenerByNameUsingPATCHParams creates a new PatchGerritListenerByNameUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchGerritListenerByNameUsingPATCHParams() *PatchGerritListenerByNameUsingPATCHParams {
	return &PatchGerritListenerByNameUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGerritListenerByNameUsingPATCHParamsWithTimeout creates a new PatchGerritListenerByNameUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewPatchGerritListenerByNameUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchGerritListenerByNameUsingPATCHParams {
	return &PatchGerritListenerByNameUsingPATCHParams{
		timeout: timeout,
	}
}

// NewPatchGerritListenerByNameUsingPATCHParamsWithContext creates a new PatchGerritListenerByNameUsingPATCHParams object
// with the ability to set a context for a request.
func NewPatchGerritListenerByNameUsingPATCHParamsWithContext(ctx context.Context) *PatchGerritListenerByNameUsingPATCHParams {
	return &PatchGerritListenerByNameUsingPATCHParams{
		Context: ctx,
	}
}

// NewPatchGerritListenerByNameUsingPATCHParamsWithHTTPClient creates a new PatchGerritListenerByNameUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchGerritListenerByNameUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchGerritListenerByNameUsingPATCHParams {
	return &PatchGerritListenerByNameUsingPATCHParams{
		HTTPClient: client,
	}
}

/*
PatchGerritListenerByNameUsingPATCHParams contains all the parameters to send to the API endpoint

	for the patch gerrit listener by name using p a t c h operation.

	Typically these are written to a http.Request.
*/
type PatchGerritListenerByNameUsingPATCHParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* GerritListenerPatch.

	   gerritListenerPatch
	*/
	GerritListenerPatch models.GerritListenerPatch

	/* Name.

	   name
	*/
	Name string

	/* Project.

	   project
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch gerrit listener by name using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGerritListenerByNameUsingPATCHParams) WithDefaults() *PatchGerritListenerByNameUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch gerrit listener by name using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGerritListenerByNameUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithContext(ctx context.Context) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithAuthorization(authorization string) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithAPIVersion(aPIVersion string) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritListenerPatch adds the gerritListenerPatch to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetGerritListenerPatch(gerritListenerPatch)
	return o
}

// SetGerritListenerPatch adds the gerritListenerPatch to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) {
	o.GerritListenerPatch = gerritListenerPatch
}

// WithName adds the name to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithName(name string) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) WithProject(project string) *PatchGerritListenerByNameUsingPATCHParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch gerrit listener by name using p a t c h params
func (o *PatchGerritListenerByNameUsingPATCHParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGerritListenerByNameUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.GerritListenerPatch); err != nil {
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