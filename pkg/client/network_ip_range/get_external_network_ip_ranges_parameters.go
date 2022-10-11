// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

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

// NewGetExternalNetworkIPRangesParams creates a new GetExternalNetworkIPRangesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalNetworkIPRangesParams() *GetExternalNetworkIPRangesParams {
	return &GetExternalNetworkIPRangesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalNetworkIPRangesParamsWithTimeout creates a new GetExternalNetworkIPRangesParams object
// with the ability to set a timeout on a request.
func NewGetExternalNetworkIPRangesParamsWithTimeout(timeout time.Duration) *GetExternalNetworkIPRangesParams {
	return &GetExternalNetworkIPRangesParams{
		timeout: timeout,
	}
}

// NewGetExternalNetworkIPRangesParamsWithContext creates a new GetExternalNetworkIPRangesParams object
// with the ability to set a context for a request.
func NewGetExternalNetworkIPRangesParamsWithContext(ctx context.Context) *GetExternalNetworkIPRangesParams {
	return &GetExternalNetworkIPRangesParams{
		Context: ctx,
	}
}

// NewGetExternalNetworkIPRangesParamsWithHTTPClient creates a new GetExternalNetworkIPRangesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalNetworkIPRangesParamsWithHTTPClient(client *http.Client) *GetExternalNetworkIPRangesParams {
	return &GetExternalNetworkIPRangesParams{
		HTTPClient: client,
	}
}

/*
GetExternalNetworkIPRangesParams contains all the parameters to send to the API endpoint

	for the get external network IP ranges operation.

	Typically these are written to a http.Request.
*/
type GetExternalNetworkIPRangesParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external network IP ranges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalNetworkIPRangesParams) WithDefaults() *GetExternalNetworkIPRangesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external network IP ranges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalNetworkIPRangesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) WithTimeout(timeout time.Duration) *GetExternalNetworkIPRangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) WithContext(ctx context.Context) *GetExternalNetworkIPRangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) WithHTTPClient(client *http.Client) *GetExternalNetworkIPRangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) WithAPIVersion(aPIVersion *string) *GetExternalNetworkIPRangesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get external network IP ranges params
func (o *GetExternalNetworkIPRangesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalNetworkIPRangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
