// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewCreateMachineSnapshotParams creates a new CreateMachineSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMachineSnapshotParams() *CreateMachineSnapshotParams {
	return &CreateMachineSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMachineSnapshotParamsWithTimeout creates a new CreateMachineSnapshotParams object
// with the ability to set a timeout on a request.
func NewCreateMachineSnapshotParamsWithTimeout(timeout time.Duration) *CreateMachineSnapshotParams {
	return &CreateMachineSnapshotParams{
		timeout: timeout,
	}
}

// NewCreateMachineSnapshotParamsWithContext creates a new CreateMachineSnapshotParams object
// with the ability to set a context for a request.
func NewCreateMachineSnapshotParamsWithContext(ctx context.Context) *CreateMachineSnapshotParams {
	return &CreateMachineSnapshotParams{
		Context: ctx,
	}
}

// NewCreateMachineSnapshotParamsWithHTTPClient creates a new CreateMachineSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMachineSnapshotParamsWithHTTPClient(client *http.Client) *CreateMachineSnapshotParams {
	return &CreateMachineSnapshotParams{
		HTTPClient: client,
	}
}

/*
CreateMachineSnapshotParams contains all the parameters to send to the API endpoint

	for the create machine snapshot operation.

	Typically these are written to a http.Request.
*/
type CreateMachineSnapshotParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Snapshot Specification details
	*/
	Body *models.SnapshotSpecification

	/* ID.

	   The id of the Machine.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create machine snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineSnapshotParams) WithDefaults() *CreateMachineSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create machine snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithTimeout(timeout time.Duration) *CreateMachineSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithContext(ctx context.Context) *CreateMachineSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithHTTPClient(client *http.Client) *CreateMachineSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithAPIVersion(aPIVersion *string) *CreateMachineSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithBody(body *models.SnapshotSpecification) *CreateMachineSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetBody(body *models.SnapshotSpecification) {
	o.Body = body
}

// WithID adds the id to the create machine snapshot params
func (o *CreateMachineSnapshotParams) WithID(id string) *CreateMachineSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create machine snapshot params
func (o *CreateMachineSnapshotParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMachineSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
