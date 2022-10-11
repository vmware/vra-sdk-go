// Code generated by go-swagger; DO NOT EDIT.

package perspective_sync

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

// NewSyncPerspectiveGroupUsingPOST2Params creates a new SyncPerspectiveGroupUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSyncPerspectiveGroupUsingPOST2Params() *SyncPerspectiveGroupUsingPOST2Params {
	return &SyncPerspectiveGroupUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSyncPerspectiveGroupUsingPOST2ParamsWithTimeout creates a new SyncPerspectiveGroupUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewSyncPerspectiveGroupUsingPOST2ParamsWithTimeout(timeout time.Duration) *SyncPerspectiveGroupUsingPOST2Params {
	return &SyncPerspectiveGroupUsingPOST2Params{
		timeout: timeout,
	}
}

// NewSyncPerspectiveGroupUsingPOST2ParamsWithContext creates a new SyncPerspectiveGroupUsingPOST2Params object
// with the ability to set a context for a request.
func NewSyncPerspectiveGroupUsingPOST2ParamsWithContext(ctx context.Context) *SyncPerspectiveGroupUsingPOST2Params {
	return &SyncPerspectiveGroupUsingPOST2Params{
		Context: ctx,
	}
}

// NewSyncPerspectiveGroupUsingPOST2ParamsWithHTTPClient creates a new SyncPerspectiveGroupUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSyncPerspectiveGroupUsingPOST2ParamsWithHTTPClient(client *http.Client) *SyncPerspectiveGroupUsingPOST2Params {
	return &SyncPerspectiveGroupUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
SyncPerspectiveGroupUsingPOST2Params contains all the parameters to send to the API endpoint

	for the sync perspective group using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type SyncPerspectiveGroupUsingPOST2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sync perspective group using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncPerspectiveGroupUsingPOST2Params) WithDefaults() *SyncPerspectiveGroupUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sync perspective group using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncPerspectiveGroupUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) WithTimeout(timeout time.Duration) *SyncPerspectiveGroupUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) WithContext(ctx context.Context) *SyncPerspectiveGroupUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) WithHTTPClient(client *http.Client) *SyncPerspectiveGroupUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) WithAPIVersion(aPIVersion *string) *SyncPerspectiveGroupUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the sync perspective group using p o s t 2 params
func (o *SyncPerspectiveGroupUsingPOST2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *SyncPerspectiveGroupUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
