// Code generated by go-swagger; DO NOT EDIT.

package flavor_profile

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

// NewUpdateFlavorProfileParams creates a new UpdateFlavorProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFlavorProfileParams() *UpdateFlavorProfileParams {
	return &UpdateFlavorProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFlavorProfileParamsWithTimeout creates a new UpdateFlavorProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateFlavorProfileParamsWithTimeout(timeout time.Duration) *UpdateFlavorProfileParams {
	return &UpdateFlavorProfileParams{
		timeout: timeout,
	}
}

// NewUpdateFlavorProfileParamsWithContext creates a new UpdateFlavorProfileParams object
// with the ability to set a context for a request.
func NewUpdateFlavorProfileParamsWithContext(ctx context.Context) *UpdateFlavorProfileParams {
	return &UpdateFlavorProfileParams{
		Context: ctx,
	}
}

// NewUpdateFlavorProfileParamsWithHTTPClient creates a new UpdateFlavorProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFlavorProfileParamsWithHTTPClient(client *http.Client) *UpdateFlavorProfileParams {
	return &UpdateFlavorProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateFlavorProfileParams contains all the parameters to send to the API endpoint

	for the update flavor profile operation.

	Typically these are written to a http.Request.
*/
type UpdateFlavorProfileParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   FlavorProfile instance
	*/
	Body *models.UpdateFlavorProfileSpecification

	/* ID.

	   The ID of the flavor.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update flavor profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFlavorProfileParams) WithDefaults() *UpdateFlavorProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update flavor profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFlavorProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithTimeout(timeout time.Duration) *UpdateFlavorProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithContext(ctx context.Context) *UpdateFlavorProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithHTTPClient(client *http.Client) *UpdateFlavorProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithAPIVersion(aPIVersion *string) *UpdateFlavorProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithBody(body *models.UpdateFlavorProfileSpecification) *UpdateFlavorProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetBody(body *models.UpdateFlavorProfileSpecification) {
	o.Body = body
}

// WithID adds the id to the update flavor profile params
func (o *UpdateFlavorProfileParams) WithID(id string) *UpdateFlavorProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update flavor profile params
func (o *UpdateFlavorProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFlavorProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
