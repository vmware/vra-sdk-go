// Code generated by go-swagger; DO NOT EDIT.

package custom_naming

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

// NewGetCustomNameByProjectIDParams creates a new GetCustomNameByProjectIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomNameByProjectIDParams() *GetCustomNameByProjectIDParams {
	return &GetCustomNameByProjectIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomNameByProjectIDParamsWithTimeout creates a new GetCustomNameByProjectIDParams object
// with the ability to set a timeout on a request.
func NewGetCustomNameByProjectIDParamsWithTimeout(timeout time.Duration) *GetCustomNameByProjectIDParams {
	return &GetCustomNameByProjectIDParams{
		timeout: timeout,
	}
}

// NewGetCustomNameByProjectIDParamsWithContext creates a new GetCustomNameByProjectIDParams object
// with the ability to set a context for a request.
func NewGetCustomNameByProjectIDParamsWithContext(ctx context.Context) *GetCustomNameByProjectIDParams {
	return &GetCustomNameByProjectIDParams{
		Context: ctx,
	}
}

// NewGetCustomNameByProjectIDParamsWithHTTPClient creates a new GetCustomNameByProjectIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomNameByProjectIDParamsWithHTTPClient(client *http.Client) *GetCustomNameByProjectIDParams {
	return &GetCustomNameByProjectIDParams{
		HTTPClient: client,
	}
}

/*
GetCustomNameByProjectIDParams contains all the parameters to send to the API endpoint

	for the get custom name by project Id operation.

	Typically these are written to a http.Request.
*/
type GetCustomNameByProjectIDParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* ID.

	   Project id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get custom name by project Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomNameByProjectIDParams) WithDefaults() *GetCustomNameByProjectIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get custom name by project Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomNameByProjectIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) WithTimeout(timeout time.Duration) *GetCustomNameByProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) WithContext(ctx context.Context) *GetCustomNameByProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) WithHTTPClient(client *http.Client) *GetCustomNameByProjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) WithAPIVersion(aPIVersion string) *GetCustomNameByProjectIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) WithID(id string) *GetCustomNameByProjectIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get custom name by project Id params
func (o *GetCustomNameByProjectIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomNameByProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
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
