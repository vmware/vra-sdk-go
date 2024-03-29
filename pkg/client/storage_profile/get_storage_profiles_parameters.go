// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

// NewGetStorageProfilesParams creates a new GetStorageProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStorageProfilesParams() *GetStorageProfilesParams {
	return &GetStorageProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageProfilesParamsWithTimeout creates a new GetStorageProfilesParams object
// with the ability to set a timeout on a request.
func NewGetStorageProfilesParamsWithTimeout(timeout time.Duration) *GetStorageProfilesParams {
	return &GetStorageProfilesParams{
		timeout: timeout,
	}
}

// NewGetStorageProfilesParamsWithContext creates a new GetStorageProfilesParams object
// with the ability to set a context for a request.
func NewGetStorageProfilesParamsWithContext(ctx context.Context) *GetStorageProfilesParams {
	return &GetStorageProfilesParams{
		Context: ctx,
	}
}

// NewGetStorageProfilesParamsWithHTTPClient creates a new GetStorageProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStorageProfilesParamsWithHTTPClient(client *http.Client) *GetStorageProfilesParams {
	return &GetStorageProfilesParams{
		HTTPClient: client,
	}
}

/*
GetStorageProfilesParams contains all the parameters to send to the API endpoint

	for the get storage profiles operation.

	Typically these are written to a http.Request.
*/
type GetStorageProfilesParams struct {

	/* DollarFilter.

	   Add a filter to return limited results
	*/
	DollarFilter *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get storage profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageProfilesParams) WithDefaults() *GetStorageProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get storage profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get storage profiles params
func (o *GetStorageProfilesParams) WithTimeout(timeout time.Duration) *GetStorageProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage profiles params
func (o *GetStorageProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage profiles params
func (o *GetStorageProfilesParams) WithContext(ctx context.Context) *GetStorageProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage profiles params
func (o *GetStorageProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage profiles params
func (o *GetStorageProfilesParams) WithHTTPClient(client *http.Client) *GetStorageProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage profiles params
func (o *GetStorageProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get storage profiles params
func (o *GetStorageProfilesParams) WithDollarFilter(dollarFilter *string) *GetStorageProfilesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get storage profiles params
func (o *GetStorageProfilesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get storage profiles params
func (o *GetStorageProfilesParams) WithAPIVersion(aPIVersion *string) *GetStorageProfilesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get storage profiles params
func (o *GetStorageProfilesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string

		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {

			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
