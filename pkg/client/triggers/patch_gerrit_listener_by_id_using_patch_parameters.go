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

// NewPatchGerritListenerByIDUsingPATCHParams creates a new PatchGerritListenerByIDUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchGerritListenerByIDUsingPATCHParams() *PatchGerritListenerByIDUsingPATCHParams {
	return &PatchGerritListenerByIDUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGerritListenerByIDUsingPATCHParamsWithTimeout creates a new PatchGerritListenerByIDUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewPatchGerritListenerByIDUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchGerritListenerByIDUsingPATCHParams {
	return &PatchGerritListenerByIDUsingPATCHParams{
		timeout: timeout,
	}
}

// NewPatchGerritListenerByIDUsingPATCHParamsWithContext creates a new PatchGerritListenerByIDUsingPATCHParams object
// with the ability to set a context for a request.
func NewPatchGerritListenerByIDUsingPATCHParamsWithContext(ctx context.Context) *PatchGerritListenerByIDUsingPATCHParams {
	return &PatchGerritListenerByIDUsingPATCHParams{
		Context: ctx,
	}
}

// NewPatchGerritListenerByIDUsingPATCHParamsWithHTTPClient creates a new PatchGerritListenerByIDUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchGerritListenerByIDUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchGerritListenerByIDUsingPATCHParams {
	return &PatchGerritListenerByIDUsingPATCHParams{
		HTTPClient: client,
	}
}

/*
PatchGerritListenerByIDUsingPATCHParams contains all the parameters to send to the API endpoint

	for the patch gerrit listener by Id using p a t c h operation.

	Typically these are written to a http.Request.
*/
type PatchGerritListenerByIDUsingPATCHParams struct {

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

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch gerrit listener by Id using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGerritListenerByIDUsingPATCHParams) WithDefaults() *PatchGerritListenerByIDUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch gerrit listener by Id using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGerritListenerByIDUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithContext(ctx context.Context) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithAuthorization(authorization string) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithAPIVersion(aPIVersion string) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritListenerPatch adds the gerritListenerPatch to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetGerritListenerPatch(gerritListenerPatch)
	return o
}

// SetGerritListenerPatch adds the gerritListenerPatch to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) {
	o.GerritListenerPatch = gerritListenerPatch
}

// WithID adds the id to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) WithID(id string) *PatchGerritListenerByIDUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch gerrit listener by Id using p a t c h params
func (o *PatchGerritListenerByIDUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGerritListenerByIDUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}