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

// NewManualGerritEventTriggerUsingPOSTParams creates a new ManualGerritEventTriggerUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManualGerritEventTriggerUsingPOSTParams() *ManualGerritEventTriggerUsingPOSTParams {
	return &ManualGerritEventTriggerUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManualGerritEventTriggerUsingPOSTParamsWithTimeout creates a new ManualGerritEventTriggerUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewManualGerritEventTriggerUsingPOSTParamsWithTimeout(timeout time.Duration) *ManualGerritEventTriggerUsingPOSTParams {
	return &ManualGerritEventTriggerUsingPOSTParams{
		timeout: timeout,
	}
}

// NewManualGerritEventTriggerUsingPOSTParamsWithContext creates a new ManualGerritEventTriggerUsingPOSTParams object
// with the ability to set a context for a request.
func NewManualGerritEventTriggerUsingPOSTParamsWithContext(ctx context.Context) *ManualGerritEventTriggerUsingPOSTParams {
	return &ManualGerritEventTriggerUsingPOSTParams{
		Context: ctx,
	}
}

// NewManualGerritEventTriggerUsingPOSTParamsWithHTTPClient creates a new ManualGerritEventTriggerUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewManualGerritEventTriggerUsingPOSTParamsWithHTTPClient(client *http.Client) *ManualGerritEventTriggerUsingPOSTParams {
	return &ManualGerritEventTriggerUsingPOSTParams{
		HTTPClient: client,
	}
}

/* ManualGerritEventTriggerUsingPOSTParams contains all the parameters to send to the API endpoint
   for the manual gerrit event trigger using p o s t operation.

   Typically these are written to a http.Request.
*/
type ManualGerritEventTriggerUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* GerritManualTrigger.

	   gerritManualTrigger
	*/
	GerritManualTrigger models.GerritManualTrigger

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the manual gerrit event trigger using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManualGerritEventTriggerUsingPOSTParams) WithDefaults() *ManualGerritEventTriggerUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the manual gerrit event trigger using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManualGerritEventTriggerUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithTimeout(timeout time.Duration) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithContext(ctx context.Context) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithHTTPClient(client *http.Client) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithAuthorization(authorization string) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithAPIVersion(aPIVersion string) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritManualTrigger adds the gerritManualTrigger to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) WithGerritManualTrigger(gerritManualTrigger models.GerritManualTrigger) *ManualGerritEventTriggerUsingPOSTParams {
	o.SetGerritManualTrigger(gerritManualTrigger)
	return o
}

// SetGerritManualTrigger adds the gerritManualTrigger to the manual gerrit event trigger using p o s t params
func (o *ManualGerritEventTriggerUsingPOSTParams) SetGerritManualTrigger(gerritManualTrigger models.GerritManualTrigger) {
	o.GerritManualTrigger = gerritManualTrigger
}

// WriteToRequest writes these params to a swagger request
func (o *ManualGerritEventTriggerUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.GerritManualTrigger); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
