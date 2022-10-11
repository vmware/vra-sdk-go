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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateBlockDeviceSnapshotParams creates a new CreateBlockDeviceSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlockDeviceSnapshotParams() *CreateBlockDeviceSnapshotParams {
	return &CreateBlockDeviceSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlockDeviceSnapshotParamsWithTimeout creates a new CreateBlockDeviceSnapshotParams object
// with the ability to set a timeout on a request.
func NewCreateBlockDeviceSnapshotParamsWithTimeout(timeout time.Duration) *CreateBlockDeviceSnapshotParams {
	return &CreateBlockDeviceSnapshotParams{
		timeout: timeout,
	}
}

// NewCreateBlockDeviceSnapshotParamsWithContext creates a new CreateBlockDeviceSnapshotParams object
// with the ability to set a context for a request.
func NewCreateBlockDeviceSnapshotParamsWithContext(ctx context.Context) *CreateBlockDeviceSnapshotParams {
	return &CreateBlockDeviceSnapshotParams{
		Context: ctx,
	}
}

// NewCreateBlockDeviceSnapshotParamsWithHTTPClient creates a new CreateBlockDeviceSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlockDeviceSnapshotParamsWithHTTPClient(client *http.Client) *CreateBlockDeviceSnapshotParams {
	return &CreateBlockDeviceSnapshotParams{
		HTTPClient: client,
	}
}

/*
CreateBlockDeviceSnapshotParams contains all the parameters to send to the API endpoint

	for the create block device snapshot operation.

	Typically these are written to a http.Request.
*/
type CreateBlockDeviceSnapshotParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Disk Snapshot Specification details
	*/
	Body *models.DiskSnapshotSpecification

	/* ID.

	   The ID of the block device.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create block device snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlockDeviceSnapshotParams) WithDefaults() *CreateBlockDeviceSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create block device snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlockDeviceSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithTimeout(timeout time.Duration) *CreateBlockDeviceSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithContext(ctx context.Context) *CreateBlockDeviceSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithHTTPClient(client *http.Client) *CreateBlockDeviceSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithAPIVersion(aPIVersion *string) *CreateBlockDeviceSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithBody(body *models.DiskSnapshotSpecification) *CreateBlockDeviceSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetBody(body *models.DiskSnapshotSpecification) {
	o.Body = body
}

// WithID adds the id to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) WithID(id string) *CreateBlockDeviceSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create block device snapshot params
func (o *CreateBlockDeviceSnapshotParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlockDeviceSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
