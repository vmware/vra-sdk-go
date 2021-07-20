// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewGetPipelineTilesUsingGETParams creates a new GetPipelineTilesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPipelineTilesUsingGETParams() *GetPipelineTilesUsingGETParams {
	return &GetPipelineTilesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineTilesUsingGETParamsWithTimeout creates a new GetPipelineTilesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetPipelineTilesUsingGETParamsWithTimeout(timeout time.Duration) *GetPipelineTilesUsingGETParams {
	return &GetPipelineTilesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetPipelineTilesUsingGETParamsWithContext creates a new GetPipelineTilesUsingGETParams object
// with the ability to set a context for a request.
func NewGetPipelineTilesUsingGETParamsWithContext(ctx context.Context) *GetPipelineTilesUsingGETParams {
	return &GetPipelineTilesUsingGETParams{
		Context: ctx,
	}
}

// NewGetPipelineTilesUsingGETParamsWithHTTPClient creates a new GetPipelineTilesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPipelineTilesUsingGETParamsWithHTTPClient(client *http.Client) *GetPipelineTilesUsingGETParams {
	return &GetPipelineTilesUsingGETParams{
		HTTPClient: client,
	}
}

/* GetPipelineTilesUsingGETParams contains all the parameters to send to the API endpoint
   for the get pipeline tiles using g e t operation.

   Typically these are written to a http.Request.
*/
type GetPipelineTilesUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pipeline tiles using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineTilesUsingGETParams) WithDefaults() *GetPipelineTilesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pipeline tiles using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineTilesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) WithTimeout(timeout time.Duration) *GetPipelineTilesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) WithContext(ctx context.Context) *GetPipelineTilesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) WithHTTPClient(client *http.Client) *GetPipelineTilesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) WithAuthorization(authorization string) *GetPipelineTilesUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) WithAPIVersion(aPIVersion string) *GetPipelineTilesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get pipeline tiles using get params
func (o *GetPipelineTilesUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineTilesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
