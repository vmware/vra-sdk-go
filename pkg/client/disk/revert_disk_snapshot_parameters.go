// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewRevertDiskSnapshotParams creates a new RevertDiskSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevertDiskSnapshotParams() *RevertDiskSnapshotParams {
	return &RevertDiskSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevertDiskSnapshotParamsWithTimeout creates a new RevertDiskSnapshotParams object
// with the ability to set a timeout on a request.
func NewRevertDiskSnapshotParamsWithTimeout(timeout time.Duration) *RevertDiskSnapshotParams {
	return &RevertDiskSnapshotParams{
		timeout: timeout,
	}
}

// NewRevertDiskSnapshotParamsWithContext creates a new RevertDiskSnapshotParams object
// with the ability to set a context for a request.
func NewRevertDiskSnapshotParamsWithContext(ctx context.Context) *RevertDiskSnapshotParams {
	return &RevertDiskSnapshotParams{
		Context: ctx,
	}
}

// NewRevertDiskSnapshotParamsWithHTTPClient creates a new RevertDiskSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevertDiskSnapshotParamsWithHTTPClient(client *http.Client) *RevertDiskSnapshotParams {
	return &RevertDiskSnapshotParams{
		HTTPClient: client,
	}
}

/*
RevertDiskSnapshotParams contains all the parameters to send to the API endpoint

	for the revert disk snapshot operation.

	Typically these are written to a http.Request.
*/
type RevertDiskSnapshotParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   Snapshot id to revert.
	*/
	QueryID string

	/* ID.

	   The id of the Disk.
	*/
	PathID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revert disk snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertDiskSnapshotParams) WithDefaults() *RevertDiskSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revert disk snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertDiskSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithTimeout(timeout time.Duration) *RevertDiskSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithContext(ctx context.Context) *RevertDiskSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithHTTPClient(client *http.Client) *RevertDiskSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithAPIVersion(aPIVersion *string) *RevertDiskSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithQueryID adds the id to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithQueryID(id string) *RevertDiskSnapshotParams {
	o.SetQueryID(id)
	return o
}

// SetQueryID adds the id to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetQueryID(id string) {
	o.QueryID = id
}

// WithPathID adds the id to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) WithPathID(id string) *RevertDiskSnapshotParams {
	o.SetPathID(id)
	return o
}

// SetPathID adds the id to the revert disk snapshot params
func (o *RevertDiskSnapshotParams) SetPathID(id string) {
	o.PathID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RevertDiskSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param id
	qrID := o.QueryID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.PathID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
