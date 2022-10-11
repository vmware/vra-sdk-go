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
)

// NewResumeOrRerunDockerRegistryEventUsingPOSTParams creates a new ResumeOrRerunDockerRegistryEventUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeOrRerunDockerRegistryEventUsingPOSTParams() *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	return &ResumeOrRerunDockerRegistryEventUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithTimeout creates a new ResumeOrRerunDockerRegistryEventUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithTimeout(timeout time.Duration) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	return &ResumeOrRerunDockerRegistryEventUsingPOSTParams{
		timeout: timeout,
	}
}

// NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithContext creates a new ResumeOrRerunDockerRegistryEventUsingPOSTParams object
// with the ability to set a context for a request.
func NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithContext(ctx context.Context) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	return &ResumeOrRerunDockerRegistryEventUsingPOSTParams{
		Context: ctx,
	}
}

// NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithHTTPClient creates a new ResumeOrRerunDockerRegistryEventUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewResumeOrRerunDockerRegistryEventUsingPOSTParamsWithHTTPClient(client *http.Client) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	return &ResumeOrRerunDockerRegistryEventUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ResumeOrRerunDockerRegistryEventUsingPOSTParams contains all the parameters to send to the API endpoint

	for the resume or rerun docker registry event using p o s t operation.

	Typically these are written to a http.Request.
*/
type ResumeOrRerunDockerRegistryEventUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* Action.

	   Resume/Rerun
	*/
	Action string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resume or rerun docker registry event using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithDefaults() *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume or rerun docker registry event using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithTimeout(timeout time.Duration) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithContext(ctx context.Context) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithHTTPClient(client *http.Client) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithAuthorization(authorization string) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAction adds the action to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithAction(action string) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetAction(action string) {
	o.Action = action
}

// WithAPIVersion adds the aPIVersion to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithAPIVersion(aPIVersion string) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WithID(id string) *ResumeOrRerunDockerRegistryEventUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resume or rerun docker registry event using p o s t params
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeOrRerunDockerRegistryEventUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {

		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
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
