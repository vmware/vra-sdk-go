// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewReleaseBlueprintVersionUsingPOST1Params creates a new ReleaseBlueprintVersionUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReleaseBlueprintVersionUsingPOST1Params() *ReleaseBlueprintVersionUsingPOST1Params {
	return &ReleaseBlueprintVersionUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseBlueprintVersionUsingPOST1ParamsWithTimeout creates a new ReleaseBlueprintVersionUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewReleaseBlueprintVersionUsingPOST1ParamsWithTimeout(timeout time.Duration) *ReleaseBlueprintVersionUsingPOST1Params {
	return &ReleaseBlueprintVersionUsingPOST1Params{
		timeout: timeout,
	}
}

// NewReleaseBlueprintVersionUsingPOST1ParamsWithContext creates a new ReleaseBlueprintVersionUsingPOST1Params object
// with the ability to set a context for a request.
func NewReleaseBlueprintVersionUsingPOST1ParamsWithContext(ctx context.Context) *ReleaseBlueprintVersionUsingPOST1Params {
	return &ReleaseBlueprintVersionUsingPOST1Params{
		Context: ctx,
	}
}

// NewReleaseBlueprintVersionUsingPOST1ParamsWithHTTPClient creates a new ReleaseBlueprintVersionUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewReleaseBlueprintVersionUsingPOST1ParamsWithHTTPClient(client *http.Client) *ReleaseBlueprintVersionUsingPOST1Params {
	return &ReleaseBlueprintVersionUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
ReleaseBlueprintVersionUsingPOST1Params contains all the parameters to send to the API endpoint

	for the release blueprint version using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type ReleaseBlueprintVersionUsingPOST1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* BlueprintID.

	   blueprintId

	   Format: uuid
	*/
	BlueprintID strfmt.UUID

	/* Version.

	   version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the release blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithDefaults() *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the release blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithTimeout(timeout time.Duration) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithContext(ctx context.Context) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithHTTPClient(client *http.Client) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithAPIVersion(aPIVersion *string) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithBlueprintID(blueprintID strfmt.UUID) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WithVersion adds the version to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) WithVersion(version string) *ReleaseBlueprintVersionUsingPOST1Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the release blueprint version using p o s t 1 params
func (o *ReleaseBlueprintVersionUsingPOST1Params) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseBlueprintVersionUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
