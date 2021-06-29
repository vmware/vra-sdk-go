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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdatePropertyGroupUsingPUTParams creates a new UpdatePropertyGroupUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePropertyGroupUsingPUTParams() *UpdatePropertyGroupUsingPUTParams {
	return &UpdatePropertyGroupUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePropertyGroupUsingPUTParamsWithTimeout creates a new UpdatePropertyGroupUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdatePropertyGroupUsingPUTParamsWithTimeout(timeout time.Duration) *UpdatePropertyGroupUsingPUTParams {
	return &UpdatePropertyGroupUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdatePropertyGroupUsingPUTParamsWithContext creates a new UpdatePropertyGroupUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdatePropertyGroupUsingPUTParamsWithContext(ctx context.Context) *UpdatePropertyGroupUsingPUTParams {
	return &UpdatePropertyGroupUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdatePropertyGroupUsingPUTParamsWithHTTPClient creates a new UpdatePropertyGroupUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePropertyGroupUsingPUTParamsWithHTTPClient(client *http.Client) *UpdatePropertyGroupUsingPUTParams {
	return &UpdatePropertyGroupUsingPUTParams{
		HTTPClient: client,
	}
}

/* UpdatePropertyGroupUsingPUTParams contains all the parameters to send to the API endpoint
   for the update property group using p u t operation.

   Typically these are written to a http.Request.
*/
type UpdatePropertyGroupUsingPUTParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* PropertyGroup.

	   propertyGroup
	*/
	PropertyGroup *models.PropertyGroup

	/* PropertyGroupID.

	   propertyGroupId

	   Format: uuid
	*/
	PropertyGroupID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update property group using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePropertyGroupUsingPUTParams) WithDefaults() *UpdatePropertyGroupUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update property group using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePropertyGroupUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithTimeout(timeout time.Duration) *UpdatePropertyGroupUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithContext(ctx context.Context) *UpdatePropertyGroupUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithHTTPClient(client *http.Client) *UpdatePropertyGroupUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithAPIVersion(aPIVersion *string) *UpdatePropertyGroupUsingPUTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithPropertyGroup adds the propertyGroup to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithPropertyGroup(propertyGroup *models.PropertyGroup) *UpdatePropertyGroupUsingPUTParams {
	o.SetPropertyGroup(propertyGroup)
	return o
}

// SetPropertyGroup adds the propertyGroup to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetPropertyGroup(propertyGroup *models.PropertyGroup) {
	o.PropertyGroup = propertyGroup
}

// WithPropertyGroupID adds the propertyGroupID to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) WithPropertyGroupID(propertyGroupID strfmt.UUID) *UpdatePropertyGroupUsingPUTParams {
	o.SetPropertyGroupID(propertyGroupID)
	return o
}

// SetPropertyGroupID adds the propertyGroupId to the update property group using p u t params
func (o *UpdatePropertyGroupUsingPUTParams) SetPropertyGroupID(propertyGroupID strfmt.UUID) {
	o.PropertyGroupID = propertyGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePropertyGroupUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.PropertyGroup != nil {
		if err := r.SetBodyParam(o.PropertyGroup); err != nil {
			return err
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
