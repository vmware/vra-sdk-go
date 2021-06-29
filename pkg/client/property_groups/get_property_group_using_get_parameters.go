// Code generated by go-swagger; DO NOT EDIT.

package property_groups

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

// NewGetPropertyGroupUsingGETParams creates a new GetPropertyGroupUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPropertyGroupUsingGETParams() *GetPropertyGroupUsingGETParams {
	return &GetPropertyGroupUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPropertyGroupUsingGETParamsWithTimeout creates a new GetPropertyGroupUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetPropertyGroupUsingGETParamsWithTimeout(timeout time.Duration) *GetPropertyGroupUsingGETParams {
	return &GetPropertyGroupUsingGETParams{
		timeout: timeout,
	}
}

// NewGetPropertyGroupUsingGETParamsWithContext creates a new GetPropertyGroupUsingGETParams object
// with the ability to set a context for a request.
func NewGetPropertyGroupUsingGETParamsWithContext(ctx context.Context) *GetPropertyGroupUsingGETParams {
	return &GetPropertyGroupUsingGETParams{
		Context: ctx,
	}
}

// NewGetPropertyGroupUsingGETParamsWithHTTPClient creates a new GetPropertyGroupUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPropertyGroupUsingGETParamsWithHTTPClient(client *http.Client) *GetPropertyGroupUsingGETParams {
	return &GetPropertyGroupUsingGETParams{
		HTTPClient: client,
	}
}

/* GetPropertyGroupUsingGETParams contains all the parameters to send to the API endpoint
   for the get property group using g e t operation.

   Typically these are written to a http.Request.
*/
type GetPropertyGroupUsingGETParams struct {

	/* DollarSelect.

	   Fields to include in content. Use * to get all fields.
	*/
	DollarSelect []string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* PropertyGroupID.

	   propertyGroupId

	   Format: uuid
	*/
	PropertyGroupID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get property group using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertyGroupUsingGETParams) WithDefaults() *GetPropertyGroupUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get property group using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertyGroupUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithTimeout(timeout time.Duration) *GetPropertyGroupUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithContext(ctx context.Context) *GetPropertyGroupUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithHTTPClient(client *http.Client) *GetPropertyGroupUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSelect adds the dollarSelect to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithDollarSelect(dollarSelect []string) *GetPropertyGroupUsingGETParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithAPIVersion adds the aPIVersion to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithAPIVersion(aPIVersion *string) *GetPropertyGroupUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithPropertyGroupID adds the propertyGroupID to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) WithPropertyGroupID(propertyGroupID strfmt.UUID) *GetPropertyGroupUsingGETParams {
	o.SetPropertyGroupID(propertyGroupID)
	return o
}

// SetPropertyGroupID adds the propertyGroupId to the get property group using get params
func (o *GetPropertyGroupUsingGETParams) SetPropertyGroupID(propertyGroupID strfmt.UUID) {
	o.PropertyGroupID = propertyGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPropertyGroupUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSelect != nil {

		// binding items for $select
		joinedDollarSelect := o.bindParamDollarSelect(reg)

		// query array param $select
		if err := r.SetQueryParam("$select", joinedDollarSelect...); err != nil {
			return err
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

	// path param propertyGroupId
	if err := r.SetPathParam("propertyGroupId", o.PropertyGroupID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPropertyGroupUsingGET binds the parameter $select
func (o *GetPropertyGroupUsingGETParams) bindParamDollarSelect(formats strfmt.Registry) []string {
	dollarSelectIR := o.DollarSelect

	var dollarSelectIC []string
	for _, dollarSelectIIR := range dollarSelectIR { // explode []string

		dollarSelectIIV := dollarSelectIIR // string as string
		dollarSelectIC = append(dollarSelectIC, dollarSelectIIV)
	}

	// items.CollectionFormat: "multi"
	dollarSelectIS := swag.JoinByFormat(dollarSelectIC, "multi")

	return dollarSelectIS
}
