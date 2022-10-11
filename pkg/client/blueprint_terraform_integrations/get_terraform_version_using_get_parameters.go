// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewGetTerraformVersionUsingGETParams creates a new GetTerraformVersionUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformVersionUsingGETParams() *GetTerraformVersionUsingGETParams {
	return &GetTerraformVersionUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformVersionUsingGETParamsWithTimeout creates a new GetTerraformVersionUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetTerraformVersionUsingGETParamsWithTimeout(timeout time.Duration) *GetTerraformVersionUsingGETParams {
	return &GetTerraformVersionUsingGETParams{
		timeout: timeout,
	}
}

// NewGetTerraformVersionUsingGETParamsWithContext creates a new GetTerraformVersionUsingGETParams object
// with the ability to set a context for a request.
func NewGetTerraformVersionUsingGETParamsWithContext(ctx context.Context) *GetTerraformVersionUsingGETParams {
	return &GetTerraformVersionUsingGETParams{
		Context: ctx,
	}
}

// NewGetTerraformVersionUsingGETParamsWithHTTPClient creates a new GetTerraformVersionUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformVersionUsingGETParamsWithHTTPClient(client *http.Client) *GetTerraformVersionUsingGETParams {
	return &GetTerraformVersionUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetTerraformVersionUsingGETParams contains all the parameters to send to the API endpoint

	for the get terraform version using g e t operation.

	Typically these are written to a http.Request.
*/
type GetTerraformVersionUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* VersionID.

	   versionId

	   Format: uuid
	*/
	VersionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform version using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformVersionUsingGETParams) WithDefaults() *GetTerraformVersionUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform version using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformVersionUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) WithTimeout(timeout time.Duration) *GetTerraformVersionUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) WithContext(ctx context.Context) *GetTerraformVersionUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) WithHTTPClient(client *http.Client) *GetTerraformVersionUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) WithAPIVersion(aPIVersion *string) *GetTerraformVersionUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithVersionID adds the versionID to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) WithVersionID(versionID strfmt.UUID) *GetTerraformVersionUsingGETParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get terraform version using get params
func (o *GetTerraformVersionUsingGETParams) SetVersionID(versionID strfmt.UUID) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformVersionUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
