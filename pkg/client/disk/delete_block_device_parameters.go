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
	"github.com/go-openapi/swag"
)

// NewDeleteBlockDeviceParams creates a new DeleteBlockDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlockDeviceParams() *DeleteBlockDeviceParams {
	return &DeleteBlockDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlockDeviceParamsWithTimeout creates a new DeleteBlockDeviceParams object
// with the ability to set a timeout on a request.
func NewDeleteBlockDeviceParamsWithTimeout(timeout time.Duration) *DeleteBlockDeviceParams {
	return &DeleteBlockDeviceParams{
		timeout: timeout,
	}
}

// NewDeleteBlockDeviceParamsWithContext creates a new DeleteBlockDeviceParams object
// with the ability to set a context for a request.
func NewDeleteBlockDeviceParamsWithContext(ctx context.Context) *DeleteBlockDeviceParams {
	return &DeleteBlockDeviceParams{
		Context: ctx,
	}
}

// NewDeleteBlockDeviceParamsWithHTTPClient creates a new DeleteBlockDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlockDeviceParamsWithHTTPClient(client *http.Client) *DeleteBlockDeviceParams {
	return &DeleteBlockDeviceParams{
		HTTPClient: client,
	}
}

/*
DeleteBlockDeviceParams contains all the parameters to send to the API endpoint

	for the delete block device operation.

	Typically these are written to a http.Request.
*/
type DeleteBlockDeviceParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ForceDelete.

	   Controls whether this is a force delete operation. If true, best effort is made for deleting this block device. Use with caution as force deleting may cause inconsistencies between the cloud provider and vRA.
	*/
	ForceDelete *bool

	/* ID.

	   The ID of the block device.
	*/
	ID string

	/* Purge.

	   Indicates if the disk has to be completely destroyed or should be kept in the system. Valid only for block devices with 'persistent' set to true.
	*/
	Purge *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete block device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlockDeviceParams) WithDefaults() *DeleteBlockDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete block device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlockDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete block device params
func (o *DeleteBlockDeviceParams) WithTimeout(timeout time.Duration) *DeleteBlockDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete block device params
func (o *DeleteBlockDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete block device params
func (o *DeleteBlockDeviceParams) WithContext(ctx context.Context) *DeleteBlockDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete block device params
func (o *DeleteBlockDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete block device params
func (o *DeleteBlockDeviceParams) WithHTTPClient(client *http.Client) *DeleteBlockDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete block device params
func (o *DeleteBlockDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete block device params
func (o *DeleteBlockDeviceParams) WithAPIVersion(aPIVersion *string) *DeleteBlockDeviceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete block device params
func (o *DeleteBlockDeviceParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithForceDelete adds the forceDelete to the delete block device params
func (o *DeleteBlockDeviceParams) WithForceDelete(forceDelete *bool) *DeleteBlockDeviceParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the delete block device params
func (o *DeleteBlockDeviceParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithID adds the id to the delete block device params
func (o *DeleteBlockDeviceParams) WithID(id string) *DeleteBlockDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete block device params
func (o *DeleteBlockDeviceParams) SetID(id string) {
	o.ID = id
}

// WithPurge adds the purge to the delete block device params
func (o *DeleteBlockDeviceParams) WithPurge(purge *bool) *DeleteBlockDeviceParams {
	o.SetPurge(purge)
	return o
}

// SetPurge adds the purge to the delete block device params
func (o *DeleteBlockDeviceParams) SetPurge(purge *bool) {
	o.Purge = purge
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlockDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ForceDelete != nil {

		// query param forceDelete
		var qrForceDelete bool

		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {

			if err := r.SetQueryParam("forceDelete", qForceDelete); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Purge != nil {

		// query param purge
		var qrPurge bool

		if o.Purge != nil {
			qrPurge = *o.Purge
		}
		qPurge := swag.FormatBool(qrPurge)
		if qPurge != "" {

			if err := r.SetQueryParam("purge", qPurge); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
