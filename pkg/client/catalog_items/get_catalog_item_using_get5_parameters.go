// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

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

// NewGetCatalogItemUsingGET5Params creates a new GetCatalogItemUsingGET5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCatalogItemUsingGET5Params() *GetCatalogItemUsingGET5Params {
	return &GetCatalogItemUsingGET5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogItemUsingGET5ParamsWithTimeout creates a new GetCatalogItemUsingGET5Params object
// with the ability to set a timeout on a request.
func NewGetCatalogItemUsingGET5ParamsWithTimeout(timeout time.Duration) *GetCatalogItemUsingGET5Params {
	return &GetCatalogItemUsingGET5Params{
		timeout: timeout,
	}
}

// NewGetCatalogItemUsingGET5ParamsWithContext creates a new GetCatalogItemUsingGET5Params object
// with the ability to set a context for a request.
func NewGetCatalogItemUsingGET5ParamsWithContext(ctx context.Context) *GetCatalogItemUsingGET5Params {
	return &GetCatalogItemUsingGET5Params{
		Context: ctx,
	}
}

// NewGetCatalogItemUsingGET5ParamsWithHTTPClient creates a new GetCatalogItemUsingGET5Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetCatalogItemUsingGET5ParamsWithHTTPClient(client *http.Client) *GetCatalogItemUsingGET5Params {
	return &GetCatalogItemUsingGET5Params{
		HTTPClient: client,
	}
}

/*
GetCatalogItemUsingGET5Params contains all the parameters to send to the API endpoint

	for the get catalog item using g e t 5 operation.

	Typically these are written to a http.Request.
*/
type GetCatalogItemUsingGET5Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ExpandProjects.

	   Retrieves the 'projects' field of the catalog item
	*/
	ExpandProjects *bool

	/* ID.

	   Catalog item ID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get catalog item using g e t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogItemUsingGET5Params) WithDefaults() *GetCatalogItemUsingGET5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get catalog item using g e t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogItemUsingGET5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithTimeout(timeout time.Duration) *GetCatalogItemUsingGET5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithContext(ctx context.Context) *GetCatalogItemUsingGET5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithHTTPClient(client *http.Client) *GetCatalogItemUsingGET5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithAPIVersion(aPIVersion *string) *GetCatalogItemUsingGET5Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpandProjects adds the expandProjects to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithExpandProjects(expandProjects *bool) *GetCatalogItemUsingGET5Params {
	o.SetExpandProjects(expandProjects)
	return o
}

// SetExpandProjects adds the expandProjects to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetExpandProjects(expandProjects *bool) {
	o.ExpandProjects = expandProjects
}

// WithID adds the id to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) WithID(id strfmt.UUID) *GetCatalogItemUsingGET5Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get catalog item using g e t 5 params
func (o *GetCatalogItemUsingGET5Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogItemUsingGET5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExpandProjects != nil {

		// query param expandProjects
		var qrExpandProjects bool

		if o.ExpandProjects != nil {
			qrExpandProjects = *o.ExpandProjects
		}
		qExpandProjects := swag.FormatBool(qrExpandProjects)
		if qExpandProjects != "" {

			if err := r.SetQueryParam("expandProjects", qExpandProjects); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}