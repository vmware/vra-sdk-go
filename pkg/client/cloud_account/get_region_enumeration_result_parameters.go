// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewGetRegionEnumerationResultParams creates a new GetRegionEnumerationResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegionEnumerationResultParams() *GetRegionEnumerationResultParams {
	return &GetRegionEnumerationResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionEnumerationResultParamsWithTimeout creates a new GetRegionEnumerationResultParams object
// with the ability to set a timeout on a request.
func NewGetRegionEnumerationResultParamsWithTimeout(timeout time.Duration) *GetRegionEnumerationResultParams {
	return &GetRegionEnumerationResultParams{
		timeout: timeout,
	}
}

// NewGetRegionEnumerationResultParamsWithContext creates a new GetRegionEnumerationResultParams object
// with the ability to set a context for a request.
func NewGetRegionEnumerationResultParamsWithContext(ctx context.Context) *GetRegionEnumerationResultParams {
	return &GetRegionEnumerationResultParams{
		Context: ctx,
	}
}

// NewGetRegionEnumerationResultParamsWithHTTPClient creates a new GetRegionEnumerationResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegionEnumerationResultParamsWithHTTPClient(client *http.Client) *GetRegionEnumerationResultParams {
	return &GetRegionEnumerationResultParams{
		HTTPClient: client,
	}
}

/*
GetRegionEnumerationResultParams contains all the parameters to send to the API endpoint

	for the get region enumeration result operation.

	Typically these are written to a http.Request.
*/
type GetRegionEnumerationResultParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of enumeration response
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get region enumeration result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegionEnumerationResultParams) WithDefaults() *GetRegionEnumerationResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get region enumeration result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegionEnumerationResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) WithTimeout(timeout time.Duration) *GetRegionEnumerationResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) WithContext(ctx context.Context) *GetRegionEnumerationResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) WithHTTPClient(client *http.Client) *GetRegionEnumerationResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) WithAPIVersion(aPIVersion *string) *GetRegionEnumerationResultParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) WithID(id string) *GetRegionEnumerationResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get region enumeration result params
func (o *GetRegionEnumerationResultParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionEnumerationResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}